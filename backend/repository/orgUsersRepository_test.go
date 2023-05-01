package repository

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupMockDB(t *testing.T) (*gorm.DB, *sql.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return gormDB, db, mock
}

// Mocks the DB and tries to successfully create an OrgUser.
func TestOrgUsersRepository_CreateOrgUser(t *testing.T) {
	gormDB, db, mock := setupMockDB(t)

	defer db.Close()

	// CREATE inserts into DB
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `org_users`")).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 7, 8, false, 9).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	// Establish repo using mock DB
	res := NewOrgUsersRepository(gormDB)

	// Fake user we are creating
	orgUser := models.OrgUsers{
		UsersID:        7,
		OrganizationID: 8,
		Role:           9,
	}

	// Create...
	_, err := res.CreateOrgUser(orgUser)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	err = mock.ExpectationsWereMet()

	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// Success for listing all OrgUsers in repo layer.
func TestOrgUsersRepository_ListAllOrgUsers(t *testing.T) {
	gormDB, db, mock := setupMockDB(t)

	defer db.Close()

	mockRows := sqlmock.NewRows([]string{"UsersID", "OrganizationID"}).
		AddRow(1, 2).
		AddRow(3, 4)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnRows(mockRows)

	// Establish repo using fake DB
	res := NewOrgUsersRepository(gormDB)

	_, err := res.ListAllOrgUsers()

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	err = mock.ExpectationsWereMet()

	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestOrgUsersRepository_FindOrgUser(t *testing.T) {
	gormDB, db, mock := setupMockDB(t)

	defer db.Close()

	// Fake user we are finding
	orgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 2,
		Role:           3,
	}

	mockRows := sqlmock.NewRows([]string{"UsersID", "OrganizationID", "Role"}).
		AddRow(orgUser.UsersID, orgUser.OrganizationID, orgUser.Role)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnRows(mockRows)

	// Establish repo using fake DB
	res := NewOrgUsersRepository(gormDB)

	_, err := res.FindOrgUser(orgUser.UsersID, orgUser.OrganizationID)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	err = mock.ExpectationsWereMet()

	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// Update OrgUser successfully in repo layer.
func TestOrgUsersRepository_UpdateOrgUser(t *testing.T) {
	gormDB, db, mock := setupMockDB(t)

	defer db.Close()

	// Fake user we are updating
	orgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 2,
		Role:           3,
	}

	mockRows := sqlmock.NewRows([]string{"UsersID", "OrganizationID", "Role"}).
		AddRow(orgUser.UsersID, orgUser.OrganizationID, orgUser.Role)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnRows(mockRows)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), orgUser.UsersID, orgUser.OrganizationID, false, orgUser.Role).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Establish repo using fake DB
	res := NewOrgUsersRepository(gormDB)

	_, err := res.UpdateOrgUser(orgUser.UsersID, orgUser.OrganizationID, orgUser.Role)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	err = mock.ExpectationsWereMet()

	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// Update OrgUser with expected failure.
func TestOrgUsersRepository_UpdateOrgUser_Failure(t *testing.T) {
	gormDB, db, mock := setupMockDB(t)

	defer db.Close()

	// Fake user we are updating
	orgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 2,
		Role:           3,
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnError(errors.New("we failed!!!!!"))

	mock.ExpectBegin()
	mock.ExpectExec("INSERT").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), orgUser.UsersID, orgUser.OrganizationID, false, orgUser.Role).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Establish repo using fake DB
	res := NewOrgUsersRepository(gormDB)

	_, err := res.UpdateOrgUser(orgUser.UsersID, orgUser.OrganizationID, orgUser.Role)

	if err == nil {
		t.Errorf("We expect an error here: %s", err)
	}

	err = mock.ExpectationsWereMet()

	if err == nil {
		t.Errorf("Error expected: %s", err)
	}
}

// Ensures a user is deleted from mock DB.
func TestOrgUsersRepository_DeleteOrgUser_Success(t *testing.T) {
	gormDB, db, mock := setupMockDB(t)

	defer db.Close()

	// Fake user we are updating
	orgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 2,
		Role:           3,
	}

	mockRows := sqlmock.NewRows([]string{"UsersID", "OrganizationID", "Role"}).
		AddRow(orgUser.UsersID, orgUser.OrganizationID, orgUser.Role)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnRows(mockRows)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs(sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Establish repo using fake DB
	res := NewOrgUsersRepository(gormDB)

	err := res.DeleteOrgUser(orgUser.UsersID, orgUser.OrganizationID)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	err = mock.ExpectationsWereMet()

	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// Expect error from DeleteOrgUser
func TestOrgUsersRepository_DeleteOrgUser_Failure(t *testing.T) {
	gormDB, db, mock := setupMockDB(t)

	defer db.Close()

	// Fake user we are updating
	orgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 2,
		Role:           3,
	}

	mockRows := sqlmock.NewRows([]string{"UsersID", "OrganizationID", "Role"}).
		AddRow(orgUser.UsersID, orgUser.OrganizationID, orgUser.Role)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnRows(mockRows)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs(sqlmock.AnyArg()).
		WillReturnError(errors.New("Oh noes :()"))
	mock.ExpectRollback()

	// Establish repo using fake DB
	res := NewOrgUsersRepository(gormDB)

	err := res.DeleteOrgUser(orgUser.UsersID, orgUser.OrganizationID)

	if err == nil {
		t.Errorf("Error was expected while updating stats: %s", err)
	}

	err = mock.ExpectationsWereMet()

	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// Expect error inside DeleteOrgUser
func TestOrgUsersRepository_DeleteOrgUser_Failure_Where(t *testing.T) {
	gormDB, db, mock := setupMockDB(t)

	defer db.Close()

	// Fake user we are updating
	orgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 2,
	}

	// Create error on purpose, expect failure
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnError(errors.New("Uh-oh spaghetti-o..."))

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs(sqlmock.AnyArg()).
		WithArgs(sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Establish repo using fake DB
	res := NewOrgUsersRepository(gormDB)

	err := res.DeleteOrgUser(orgUser.UsersID, orgUser.OrganizationID)

	if err == nil {
		t.Errorf("Error was expected while updating stats: %s", err)
	}

	err = mock.ExpectationsWereMet()

	if err == nil {
		t.Errorf("There was supposed to be an error: %s", err)
	}
}
