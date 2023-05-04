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

type OrganizationRepositoryUnitTestSuite struct {
	suite.Suite
	db          *sql.DB
	mock        sqlmock.Sqlmock
	err         error
	gormDB      *gorm.DB
	repo        OrganizationRepository
	object      models.Organization
	arrayObject []models.Organization
	paramID     string
}

func (suite *OrganizationRepositoryUnitTestSuite) SetupTest() {
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

	suite.repo = NewOrganizationRepository(suite.gormDB)
	suite.err = fmt.Errorf("error")
	suite.paramID = "1"
}

func (suite *OrganizationRepositoryUnitTestSuite) AfterTest(_, _ string) {
	if suite.err = suite.mock.ExpectationsWereMet(); suite.err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", suite.err)
	}
}

func TestOrganizationRepositoryUnitTestSuite(t *testing.T) {
	suite.Run(t, new(OrganizationRepositoryUnitTestSuite))
}

func (suite *OrganizationRepositoryUnitTestSuite) TestOrganizationRepository_CreateOrganization_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		"", "", false, "").WillReturnError(suite.err)
	suite.mock.ExpectRollback()

	if _, suite.err = suite.repo.CreateOrganization(suite.object); suite.err == nil {
		suite.T().Errorf("error was expected while updating stats: %s", suite.err)
	}
}

func (suite *OrganizationRepositoryUnitTestSuite) TestOrganizationRepository_CreateOrganization_Success() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		"", "", false, "").WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	if _, suite.err = suite.repo.CreateOrganization(suite.object); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *OrganizationRepositoryUnitTestSuite) TestOrganizationRepository_GetOrganizations_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").WillReturnError(suite.err)

	if _, suite.err = suite.repo.GetOrganizations(); suite.err == nil {
		suite.T().Errorf("error was expected while updating stats: %s", suite.err)
	}
}

func (suite *OrganizationRepositoryUnitTestSuite) TestOrganizationRepository_GetOrganizations_Success() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").
		WillReturnRows(sqlmock.NewRows([]string{"Name", "Description", "Verified", "Interests"}).AddRow("", "", false, ""))

	if _, suite.err = suite.repo.GetOrganizations(); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *OrganizationRepositoryUnitTestSuite) TestOrganizationRepository_GetOrganizationById_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").
		WithArgs(suite.paramID).
		WillReturnError(suite.err)

	if _, suite.err = suite.repo.GetOrganizationById(suite.paramID); suite.err == nil {
		suite.T().Errorf("error was expected while updating stats: %s", suite.err)
	}
}

func (suite *OrganizationRepositoryUnitTestSuite) TestOrganizationRepository_GetOrganizationById_Success() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").
		WithArgs(suite.paramID).
		WillReturnRows(sqlmock.NewRows([]string{"Name", "Description", "Verified", "Interests"}).AddRow("", "", false, ""))

	if _, suite.err = suite.repo.GetOrganizationById(suite.paramID); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *OrganizationRepositoryUnitTestSuite) TestOrganizationRepository_UpdateOrganization_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		"", "", false, "").WillReturnError(suite.err)
	suite.mock.ExpectRollback()

	if _, suite.err = suite.repo.UpdateOrganization(suite.object); suite.err == nil {
		suite.T().Errorf("error was expected while updating stats: %s", suite.err)
	}
}

func (suite *OrganizationRepositoryUnitTestSuite) TestOrganizationRepository_UpdateOrganization_Success() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		"", "", false, "").WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	if _, suite.err = suite.repo.UpdateOrganization(suite.object); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *OrganizationRepositoryUnitTestSuite) TestOrganizationRepository_DeleteOrganization_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("UPDATE").WithArgs(sqlmock.AnyArg(), suite.object.ID).WillReturnError(suite.err)
	suite.mock.ExpectRollback()

	if suite.err = suite.repo.DeleteOrganization(suite.object); suite.err == nil {
		suite.T().Errorf("error was expected while updating stats: %s", suite.err)
	}
}

func (suite *OrganizationRepositoryUnitTestSuite) TestOrganizationRepository_DeleteOrganization_Success() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("UPDATE").WithArgs(sqlmock.AnyArg(), suite.object.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	if suite.err = suite.repo.DeleteOrganization(suite.object); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}
