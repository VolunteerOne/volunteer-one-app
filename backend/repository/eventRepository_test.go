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
	"time"
)

type EventRepositoryUnitTestSuite struct {
	suite.Suite
	db           *sql.DB
	mock         sqlmock.Sqlmock
	err          error
	gormDB       *gorm.DB
	repo         EventRepository
	object       models.Event
	arrayObject  []models.Event
	paramID      string
	expectedExec *sqlmock.Rows
}

func (suite *EventRepositoryUnitTestSuite) SetupTest() {
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

	suite.repo = NewEventRepository(suite.gormDB)
	suite.err = fmt.Errorf("error")
	suite.paramID = "1"

	suite.expectedExec = sqlmock.NewRows([]string{"OrganizationID", "Name", "Address", "Date", "Description",
		"Interests", "Skills", "GoodFor", "CauseAreas", "Requirements"})
	suite.expectedExec.AddRow(uint(0), "", "", time.Now(), "", "", "", "", "", "")
}

func (suite *EventRepositoryUnitTestSuite) AfterTest(_, _ string) {
	if suite.err = suite.mock.ExpectationsWereMet(); suite.err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", suite.err)
	}
}

func TestEventRepositoryUnitTestSuite(t *testing.T) {
	suite.Run(t, new(EventRepositoryUnitTestSuite))
}

func (suite *EventRepositoryUnitTestSuite) TestEventRepository_CreateEvent_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), uint(0),
		"", "", sqlmock.AnyArg(), "", "", "", "", "", "").WillReturnError(suite.err)
	suite.mock.ExpectRollback()

	if _, suite.err = suite.repo.CreateEvent(suite.object); suite.err == nil {
		suite.T().Errorf("error was expected while updating stats: %s", suite.err)
	}
}

func (suite *EventRepositoryUnitTestSuite) TestEventRepository_CreateEvent_Success() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), uint(0),
		"", "", sqlmock.AnyArg(), "", "", "", "", "", "").
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	if _, suite.err = suite.repo.CreateEvent(suite.object); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *EventRepositoryUnitTestSuite) TestEventRepository_DeleteEvent_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("UPDATE").WithArgs(sqlmock.AnyArg(), suite.object.ID).WillReturnError(suite.err)
	suite.mock.ExpectRollback()

	if suite.err = suite.repo.DeleteEvent(suite.object); suite.err == nil {
		suite.T().Errorf("error was expected while updating stats: %s", suite.err)
	}
}

func (suite *EventRepositoryUnitTestSuite) TestEventRepository_DeleteEvent_Success() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("UPDATE").WithArgs(sqlmock.AnyArg(), suite.object.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	if suite.err = suite.repo.DeleteEvent(suite.object); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *EventRepositoryUnitTestSuite) TestEventRepository_GetEventById_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").
		WithArgs(suite.paramID).
		WillReturnError(suite.err)

	if _, suite.err = suite.repo.GetEventById(suite.paramID); suite.err == nil {
		suite.T().Errorf("error was expected while updating stats: %s", suite.err)
	}
}

func (suite *EventRepositoryUnitTestSuite) TestEventRepository_GetEventById_Success() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").
		WithArgs(suite.paramID).
		WillReturnRows(suite.expectedExec)

	if _, suite.err = suite.repo.GetEventById(suite.paramID); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *EventRepositoryUnitTestSuite) TestEventRepository_GetEvents_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").WillReturnError(suite.err)

	if _, suite.err = suite.repo.GetEvents(); suite.err == nil {
		suite.T().Errorf("error was expected while updating stats: %s", suite.err)
	}
}

func (suite *EventRepositoryUnitTestSuite) TestEventRepository_GetEvents_Success() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").
		WillReturnRows(suite.expectedExec)

	if _, suite.err = suite.repo.GetEvents(); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *EventRepositoryUnitTestSuite) TestEventRepository_UpdateEvent_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), uint(0),
		"", "", sqlmock.AnyArg(), "", "", "", "", "", "").WillReturnError(suite.err)
	suite.mock.ExpectRollback()

	if _, suite.err = suite.repo.UpdateEvent(suite.object); suite.err == nil {
		suite.T().Errorf("error was expected while updating stats: %s", suite.err)
	}
}

func (suite *EventRepositoryUnitTestSuite) TestEventRepository_UpdateEvent_Success() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), uint(0),
		"", "", sqlmock.AnyArg(), "", "", "", "", "", "").
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	if _, suite.err = suite.repo.UpdateEvent(suite.object); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}
