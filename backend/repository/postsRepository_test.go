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

type PostsRepositoryUnitTestSuite struct {
	suite.Suite
	db          *sql.DB
	mock        sqlmock.Sqlmock
	err         error
	gormDB      *gorm.DB
	repo        PostsRepository
	object      models.Posts
	arrayObject []models.Posts
	paramID     string
}

func (suite *PostsRepositoryUnitTestSuite) SetupTest() {
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

	suite.repo = NewPostsRepository(suite.gormDB)
	suite.err = fmt.Errorf("error")
	suite.paramID = "1"
}

func (suite *PostsRepositoryUnitTestSuite) AfterTest(_, _ string) {
	if suite.err = suite.mock.ExpectationsWereMet(); suite.err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", suite.err)
	}
}

func TestPostsRepositoryUnitTestSuite(t *testing.T) {
	suite.Run(t, new(PostsRepositoryUnitTestSuite))
}

func (suite *PostsRepositoryUnitTestSuite) TestPostsRepository_CreatePost() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		"", "", uint(0)).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	if _, suite.err = suite.repo.CreatePost(suite.object); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *PostsRepositoryUnitTestSuite) TestFriendRepository_RejectFriend_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("UPDATE").WithArgs(sqlmock.AnyArg(), suite.object.ID).WillReturnError(suite.err)
	suite.mock.ExpectRollback()

	if suite.err = suite.repo.DeletePost(suite.object); suite.err == nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *PostsRepositoryUnitTestSuite) TestPostsRepository_DeletePost_Success() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("UPDATE").WithArgs(sqlmock.AnyArg(), suite.object.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	suite.object.DeletedAt = gorm.DeletedAt{}
	if suite.err = suite.repo.DeletePost(suite.object); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *PostsRepositoryUnitTestSuite) TestFriendRepository_EditPost_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		"", "", uint(0)).WillReturnError(suite.err)
	suite.mock.ExpectRollback()

	if _, suite.err = suite.repo.EditPost(suite.object); suite.err == nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *PostsRepositoryUnitTestSuite) TestPostsRepository_EditPost_Success() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		"", "", uint(0)).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	if _, suite.err = suite.repo.EditPost(suite.object); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *PostsRepositoryUnitTestSuite) TestPostsRepository_FindPost_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").
		WithArgs(suite.paramID).
		WillReturnError(suite.err)

	if _, suite.err = suite.repo.FindPost(suite.paramID); suite.err == nil {
		suite.T().Errorf("error was expected while updating stats: %s", suite.err)
	}
}

func (suite *PostsRepositoryUnitTestSuite) TestPostsRepository_FindPost_Success() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").
		WithArgs(suite.paramID).
		WillReturnRows(sqlmock.NewRows([]string{"Handle", "PostDescription", "Likes"}).AddRow("", "", uint(0)))

	if _, suite.err = suite.repo.FindPost(suite.paramID); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *PostsRepositoryUnitTestSuite) TestPostsRepository_AllPosts_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").WillReturnError(suite.err)

	if _, suite.err = suite.repo.AllPosts(); suite.err == nil {
		suite.T().Errorf("error was expected while updating stats: %s", suite.err)
	}
}

func (suite *PostsRepositoryUnitTestSuite) TestPostsRepository_AllPosts_Success() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").
		WillReturnRows(sqlmock.NewRows([]string{"Handle", "PostDescription", "Likes"}).AddRow("", "", uint(0)))

	if _, suite.err = suite.repo.AllPosts(); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}
