package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SQLMockSuite struct {
	suite.Suite
	db     *sql.DB
	mock   sqlmock.Sqlmock
	err    error
	gormDB *gorm.DB
	repo   LoginRepository
}

// ran before each test
func (suite *SQLMockSuite) SetupTest() {
	// set up a mock db for each test
	suite.db, suite.mock, suite.err = sqlmock.New()
	suite.gormDB, suite.err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      suite.db,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if suite.err != nil {
		suite.T().Fatalf("an error '%s' was not expected when opening a stub database connection", suite.err)
	}
	// instantiate new repo for each test
	suite.repo = NewLoginRepository(suite.gormDB)
}

// ran after each test
func (suite *SQLMockSuite) AfterTest(_, _ string) {
	// we make sure that all expectations were met
	if suite.err = suite.mock.ExpectationsWereMet(); suite.err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", suite.err)
	}
}

// run all the tests in the suite
func TestSQLMockSuite(t *testing.T) {
	suite.Run(t, new(SQLMockSuite))
}

func (suite *SQLMockSuite) TestLoginRepository_FindUserFromEmail() {
	defer suite.db.Close()

	email := "test@user.com"

	mockRows := sqlmock.NewRows([]string{"Email", "Password"}).
		AddRow(email, "fakepassword")

	suite.mock.ExpectQuery("SELECT(.*)").
		WithArgs(email).
		WillReturnRows(mockRows)

	var user models.Users

	// now we execute our method
	if user, suite.err = suite.repo.FindUserFromEmail(email, user); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *SQLMockSuite) TestLoginRepository_SaveResetCodeToUser() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(),
		sqlmock.AnyArg(), "", "", "", "", "", "", "", 0, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	var user models.Users

	fakeCode, _ := uuid.Parse("00000000-0000-0000-0000-000000000000")

	// now we execute our method
	if suite.err = suite.repo.SaveResetCodeToUser(fakeCode, user); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *SQLMockSuite) TestLoginRepository_ChangePassword() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(),
		sqlmock.AnyArg(), "", "", "", "", "", "", "", 0, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	var user models.Users

	// now we execute our method
	if suite.err = suite.repo.ChangePassword([]byte(""), user); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *SQLMockSuite) TestLoginRepository_FindTokenFromID() {
	defer suite.db.Close()

	mockRows := sqlmock.NewRows([]string{"RefreshToken", "UsersID"}).
		AddRow("", uint(0))

	suite.mock.ExpectQuery("SELECT(.*)").
		WithArgs(uint(0)).
		WillReturnRows(mockRows)

	var deleg models.Delegations

	// now we execute our method
	if deleg, suite.err = suite.repo.FindTokenFromID(uint(0), deleg); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *SQLMockSuite) TestLoginRepository_SaveRefreshTokenFail() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(),
		sqlmock.AnyArg(), "", uint(0)).
		WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("error")))
	suite.mock.ExpectCommit()

	var deleg models.Delegations

	// now we execute our method
	if suite.err = suite.repo.SaveRefreshToken(uint(0), "", deleg); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}
