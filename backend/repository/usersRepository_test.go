package repository

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type UsersRepositoryUnitTestSuite struct {
	suite.Suite
	db           *sql.DB
	mock         sqlmock.Sqlmock
	err          error
	gormDB       *gorm.DB
	repo         UsersRepository
	object       models.Users
	arrayObject  []models.Users
	paramID      string
	expectedExec *sqlmock.Rows
	fakeUUID     uuid.UUID
}

func (suite *UsersRepositoryUnitTestSuite) SetupTest() {
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

	suite.repo = NewUsersRepository(suite.gormDB)
	suite.err = fmt.Errorf("error")
	suite.paramID = "1"
	suite.fakeUUID, _ = uuid.Parse("00000000-0000-0000-0000-000000000000")

	suite.expectedExec = sqlmock.NewRows([]string{"Handle", "Email", "Password", "Birthdate", "FirstName",
		"LastName", "Interests", "Verified", "ResetCode"})
	suite.expectedExec.AddRow("", "", "", "", "", "", "", uint(0), suite.fakeUUID)
}

func (suite *UsersRepositoryUnitTestSuite) AfterTest(_, _ string) {
	if suite.err = suite.mock.ExpectationsWereMet(); suite.err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", suite.err)
	}
}

func TestUsersRepositoryUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UsersRepositoryUnitTestSuite))
}

func (suite *UsersRepositoryUnitTestSuite) TestUsersRepository_CreateUser_Success() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		"", "", "", "", "", "", "", uint(0), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	if _, suite.err = suite.repo.CreateUser(suite.object); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *UsersRepositoryUnitTestSuite) TestUsersRepository_GetUsersById_Success() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").
		WithArgs(suite.paramID).
		WillReturnRows(suite.expectedExec)

	if _, suite.err = suite.repo.OneUser(suite.paramID, suite.object); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *UsersRepositoryUnitTestSuite) TestUsersRepository_UpdateUsers_Success() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		"", "", "", "", "", "", "", uint(0), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	if _, suite.err = suite.repo.UpdateUser(suite.object); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *UsersRepositoryUnitTestSuite) TestUsersRepository_DeleteUsers_Success() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("UPDATE").WithArgs(sqlmock.AnyArg(), suite.object.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	if _, suite.err = suite.repo.DeleteUser(suite.object); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}
