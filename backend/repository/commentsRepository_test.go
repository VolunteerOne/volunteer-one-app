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

type CommentsRepositoryUnitTestSuite struct {
	suite.Suite
	db          *sql.DB
	mock        sqlmock.Sqlmock
	err         error
	gormDB      *gorm.DB
	repo        CommentsRepository
	object      models.Comments
	arrayObject []models.Comments
	paramID     string
}

func (suite *CommentsRepositoryUnitTestSuite) SetupTest() {
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

	suite.repo = NewCommentsRepository(suite.gormDB)
	suite.err = fmt.Errorf("error")
	suite.paramID = "1"
}

func (suite *CommentsRepositoryUnitTestSuite) AfterTest(_, _ string) {
	if suite.err = suite.mock.ExpectationsWereMet(); suite.err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", suite.err)
	}
}

func TestCommentsRepositoryUnitTestSuite(t *testing.T) {
	suite.Run(t, new(CommentsRepositoryUnitTestSuite))
}

func (suite *CommentsRepositoryUnitTestSuite) TestCommentsRepository_CreateComment() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		uint(0), "", "").WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	if _, suite.err = suite.repo.CreateComment(suite.object); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *CommentsRepositoryUnitTestSuite) TestCommentsRepository_DeleteComment_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("UPDATE").WithArgs(sqlmock.AnyArg(), suite.object.ID).WillReturnError(suite.err)
	suite.mock.ExpectRollback()

	if suite.err = suite.repo.DeleteComment(suite.object); suite.err == nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *CommentsRepositoryUnitTestSuite) TestCommentsRepository_DeleteComment_Success() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("UPDATE").WithArgs(sqlmock.AnyArg(), suite.object.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	suite.object.DeletedAt = gorm.DeletedAt{}
	if suite.err = suite.repo.DeleteComment(suite.object); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *CommentsRepositoryUnitTestSuite) TestCommentsRepository_EditComment_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		uint(0), "", "").WillReturnError(suite.err)
	suite.mock.ExpectRollback()

	if _, suite.err = suite.repo.EditComment(suite.object); suite.err == nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *CommentsRepositoryUnitTestSuite) TestCommentRepository_EditComment_Success() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		uint(0), "", "").WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	if _, suite.err = suite.repo.EditComment(suite.object); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *CommentsRepositoryUnitTestSuite) TestCommentsRepository_FindComment_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").
		WithArgs(suite.paramID).
		WillReturnError(suite.err)

	if _, suite.err = suite.repo.FindComment(suite.paramID); suite.err == nil {
		suite.T().Errorf("error was expected while updating stats: %s", suite.err)
	}
}

func (suite *CommentsRepositoryUnitTestSuite) TestCommentRepository_FindComment_Success() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").
		WithArgs(suite.paramID).
		WillReturnRows(sqlmock.NewRows([]string{"Handle", "PostDescription", "Likes"}).AddRow("", "", uint(0)))

	if _, suite.err = suite.repo.FindComment(suite.paramID); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}

func (suite *CommentsRepositoryUnitTestSuite) TestCommentsRepository_AllComments_Fail() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").WillReturnError(suite.err)

	if _, suite.err = suite.repo.AllComments(); suite.err == nil {
		suite.T().Errorf("error was expected while updating stats: %s", suite.err)
	}
}

func (suite *CommentsRepositoryUnitTestSuite) TestCommentsRepository_AllComments_Success() {
	defer suite.db.Close()

	suite.mock.ExpectQuery("SELECT(.*)").
		WillReturnRows(sqlmock.NewRows([]string{"Handle", "PostDescription", "Likes"}).AddRow("", "", uint(0)))

	if _, suite.err = suite.repo.AllComments(); suite.err != nil {
		suite.T().Errorf("error was not expected while updating stats: %s", suite.err)
	}
}
