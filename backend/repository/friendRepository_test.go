package repository

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type FriendRepositoryUnitTestSuite struct {
	suite.Suite
	db                 *sql.DB
	mock               sqlmock.Sqlmock
	err                error
	gormDB             *gorm.DB
	repo               FriendRepository
	friendsObject      models.Friend
	arrayFriendsObject []models.Friend
	relationshipBit    string
}

func (suite *FriendRepositoryUnitTestSuite) SetupTest() {
	suite.db, suite.mock, suite.err = sqlmock.New()
	if suite.err != nil {
		suite.T().Fatalf("an error '%s' was not expected when opening a stub database connection", suite.err)
	}

	suite.gormDB, suite.err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      suite.db,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if suite.err != nil {
		suite.T().Fatalf("an error '%s' was not expected when opening a stub database connection", suite.err)
	}

	suite.repo = NewFriendRepository(suite.gormDB)
	suite.relationshipBit = "pending"
	suite.err = fmt.Errorf("error")
}

func (suite *FriendRepositoryUnitTestSuite) AfterTest(_, _ string) {
	if suite.err = suite.mock.ExpectationsWereMet(); suite.err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", suite.err)
	}
}

func TestSQLMockSuite(t *testing.T) {
	suite.Run(t, new(FriendRepositoryUnitTestSuite))
}

func (suite *FriendRepositoryUnitTestSuite) TestFriendRepository_CreateFriend() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		"", "", suite.relationshipBit).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	if _, suite.err = suite.repo.CreateFriend(suite.friendsObject); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *FriendRepositoryUnitTestSuite) TestFriendRepository_AcceptFriend_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		"", "", suite.relationshipBit).WillReturnError(suite.err)
	suite.mock.ExpectRollback()

	if _, suite.err = suite.repo.AcceptFriend(suite.friendsObject); suite.err == nil {
		suite.T().Errorf("error was expected while updating stats: %s", suite.err)
	}
}

func (suite *FriendRepositoryUnitTestSuite) TestFriendRepository_AcceptFriend_Success() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		"", "", suite.relationshipBit).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	if _, suite.err = suite.repo.AcceptFriend(suite.friendsObject); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *FriendRepositoryUnitTestSuite) TestFriendRepository_RejectFriend_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("UPDATE").WithArgs(sqlmock.AnyArg(), suite.friendsObject.ID).WillReturnError(suite.err)
	suite.mock.ExpectRollback()

	if suite.err = suite.repo.RejectFriend(suite.friendsObject); suite.err == nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *FriendRepositoryUnitTestSuite) TestFriendRepository_RejectFriend_Success() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("UPDATE").WithArgs(sqlmock.AnyArg(), suite.friendsObject.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	suite.friendsObject.DeletedAt = gorm.DeletedAt{}
	if suite.err = suite.repo.RejectFriend(suite.friendsObject); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *FriendRepositoryUnitTestSuite) TestFriendRepository_OneFriend_Fail() {
	defer suite.db.Close()

	id := "1"
	suite.mock.ExpectQuery("SELECT(.*)").
		WithArgs(id).
		WillReturnError(suite.err)

	if _, suite.err = suite.repo.OneFriend(id); suite.err == nil {
		suite.T().Errorf("error was expected while updating stats: %s", suite.err)
	}
}

func (suite *FriendRepositoryUnitTestSuite) TestFriendRepository_OneFriend_Success() {
	defer suite.db.Close()

	id := "1"
	suite.mock.ExpectQuery("SELECT(.*)").
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"FriendOneHandle", "FriendTwoHandle", "RelationshipBit"}).
			AddRow("", "", ""))

	if _, suite.err = suite.repo.OneFriend(id); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *FriendRepositoryUnitTestSuite) TestFriendRepository_GetFriends_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").WillReturnError(suite.err)

	if _, suite.err = suite.repo.GetFriends(); suite.err == nil {
		suite.T().Errorf("error was expected while updating stats: %s", suite.err)
	}
}

func (suite *FriendRepositoryUnitTestSuite) TestFriendRepository_GetFriends_Success() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").
		WillReturnRows(sqlmock.NewRows([]string{"FriendOneHandle", "FriendTwoHandle", "RelationshipBit"}).AddRow("", "", ""))

	if _, suite.err = suite.repo.GetFriends(); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}
