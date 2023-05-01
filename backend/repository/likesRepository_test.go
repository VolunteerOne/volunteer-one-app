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

type LikesRepositoryUnitTestSuite struct {
	suite.Suite
	db          *sql.DB
	mock        sqlmock.Sqlmock
	err         error
	gormDB      *gorm.DB
	repo        LikesRepository
	object      models.Likes
	arrayObject []models.Likes
	paramID     string
}

func (suite *LikesRepositoryUnitTestSuite) SetupTest() {
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

	suite.repo = NewLikesRepository(suite.gormDB)
	suite.err = fmt.Errorf("error")
	suite.paramID = "1"
}

func (suite *LikesRepositoryUnitTestSuite) AfterTest(_, _ string) {
	if suite.err = suite.mock.ExpectationsWereMet(); suite.err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", suite.err)
	}
}

func TestLikesRepositoryUnitTestSuite(t *testing.T) {
	suite.Run(t, new(LikesRepositoryUnitTestSuite))
}

func (suite *LikesRepositoryUnitTestSuite) TestLikesRepository_CreateLike() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		uint(0), "").WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	if _, suite.err = suite.repo.CreateLike(suite.object); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *LikesRepositoryUnitTestSuite) TestLikesRepository_DeleteLike() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("UPDATE").WithArgs(sqlmock.AnyArg(), suite.object.ID).WillReturnError(suite.err)
	suite.mock.ExpectRollback()

	if suite.err = suite.repo.DeleteLike(suite.object); suite.err == nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *LikesRepositoryUnitTestSuite) TestLikesRepository_DeleteLike_Success() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("UPDATE").WithArgs(sqlmock.AnyArg(), suite.object.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	suite.object.DeletedAt = gorm.DeletedAt{}
	if suite.err = suite.repo.DeleteLike(suite.object); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *LikesRepositoryUnitTestSuite) TestLikesRepository_FindLike_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").
		WithArgs(suite.paramID).
		WillReturnError(suite.err)

	if _, suite.err = suite.repo.FindLike(suite.paramID); suite.err == nil {
		suite.T().Errorf("error was expected while updating stats: %s", suite.err)
	}
}

func (suite *LikesRepositoryUnitTestSuite) TestLikesRepository_FindLike_Success() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").
		WithArgs(suite.paramID).
		WillReturnRows(sqlmock.NewRows([]string{"PostsID", "Handle"}).AddRow(uint(0), ""))

	if _, suite.err = suite.repo.FindLike(suite.paramID); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *LikesRepositoryUnitTestSuite) TestLikesRepository_AllLikes_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").WillReturnError(suite.err)

	if _, suite.err = suite.repo.AllLikes(); suite.err == nil {
		suite.T().Errorf("error was expected while updating stats: %s", suite.err)
	}
}

func (suite *LikesRepositoryUnitTestSuite) TestLikesRepository_AllLikes_Success() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").
		WillReturnRows(sqlmock.NewRows([]string{"PostsID", "Handle"}).AddRow(uint(0), ""))

	if _, suite.err = suite.repo.AllLikes(); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *LikesRepositoryUnitTestSuite) TestLikesRepository_GetLikes_FindError() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").WithArgs(suite.paramID).WillReturnError(suite.err)

	if _, suite.err = suite.repo.GetLikes(suite.paramID); suite.err == nil {
		suite.T().Errorf("error was expected while updating stats: %s", suite.err)
	}
}

func (suite *LikesRepositoryUnitTestSuite) TestLikesRepository_GetLikes_CountError() {
	defer suite.db.Close()

	normalRows := sqlmock.NewRows([]string{"PostsID", "Handle"}).
		AddRow(uint(0), "").
		AddRow(uint(0), "").
		AddRow(uint(0), "")
	suite.mock.ExpectQuery("SELECT(.*)").WithArgs(suite.paramID).WillReturnRows(normalRows)
	suite.mock.ExpectQuery("SELECT(.*)").WithArgs(suite.paramID).WillReturnError(suite.err)

	if _, suite.err = suite.repo.GetLikes(suite.paramID); suite.err == nil {
		suite.T().Errorf("error was expected while updating stats: %s", suite.err)
	}
}

func (suite *LikesRepositoryUnitTestSuite) TestLikesRepository_GetLikes_Success() {
	defer suite.db.Close()

	normalRows := sqlmock.NewRows([]string{"PostsID", "Handle"}).
		AddRow(uint(0), "").
		AddRow(uint(0), "").
		AddRow(uint(0), "")
	rows := sqlmock.NewRows([]string{"count"}).AddRow(int64(3))

	suite.mock.ExpectQuery("SELECT(.*)").WithArgs(suite.paramID).WillReturnRows(normalRows)
	suite.mock.ExpectQuery("SELECT(.*)").WithArgs(suite.paramID).WillReturnRows(rows)

	if _, suite.err = suite.repo.GetLikes(suite.paramID); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}
