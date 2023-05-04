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

func setupMockDBMessages(t *testing.T) (*gorm.DB, *sql.DB, sqlmock.Sqlmock) {
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

// Mocks the DB and tries to successfully create a message.
func TestMessagesRepository_CreateMessage(t *testing.T) {
	gormDB, db, mock := setupMockDBMessages(t)

	defer db.Close()

	// CREATE inserts into DB
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `messages`")).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 1, 2, "ooo", "ahh", false).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	// Establish repo using mock DB
	res := NewMessagesRepository(gormDB)

	// Fake message we are creating
	message := models.Messages{
		FromUsersID: 1,
		ToUsersID:   2,
		Subject:     "ooo",
		Message:     "ahh",
	}

	// Create...
	_, err := res.CreateMessage(message)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	err = mock.ExpectationsWereMet()

	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// Success for listing all messages for specific user in repo layer.
func TestMessagesRepository_ListAllMessagesForUser(t *testing.T) {
	gormDB, db, mock := setupMockDBMessages(t)

	defer db.Close()

	mockRows := sqlmock.NewRows([]string{"FromUsersID", "ToUsersID", "Subject", "Message"}).
		AddRow(1, 2, "sub", "msg").
		AddRow(3, 2, "sub2", "msg2")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnRows(mockRows)

	// Establish repo using fake DB
	res := NewMessagesRepository(gormDB)

	_, err := res.ListAllMessagesForUser(uint(2))

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	err = mock.ExpectationsWereMet()

	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestMessagesRepository_FindMessage(t *testing.T) {
	gormDB, db, mock := setupMockDBMessages(t)

	defer db.Close()

	// Fake message we are finding
	message := models.Messages{
		FromUsersID: 16,
		ToUsersID:   17,
		Subject:     "You are 16, going on 17...",
		Message:     "Baby, it's time to think.",
	}

	mockRows := sqlmock.NewRows([]string{"FromUsersID", "ToUsersID", "Subject", "Message"}).
		AddRow(message.FromUsersID, message.ToUsersID, message.Subject, message.Message)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnRows(mockRows)

	// Establish repo using fake DB
	res := NewMessagesRepository(gormDB)

	_, err := res.FindMessage(uint(111))

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	err = mock.ExpectationsWereMet()

	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// Update UpdateMessageReadStatus successfully in repo layer.
func TestMessagesRepository_UpdateMessageReadStatus(t *testing.T) {
	gormDB, db, mock := setupMockDBMessages(t)

	defer db.Close()

	// Fake message we are updating
	message := models.Messages{
		FromUsersID: 33,
		ToUsersID:   300000001,
		Subject:     "what what",
		Message:     "what what what",
		Read:        true,
	}

	mockRows := sqlmock.NewRows([]string{"FromUsersID", "ToUsersID", "Subject", "Message", "Read"}).
		AddRow(message.FromUsersID, message.ToUsersID, message.Subject, message.Message, message.Read)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnRows(mockRows)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT").
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			message.FromUsersID,
			message.ToUsersID,
			message.Subject,
			message.Message,
			message.Read).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Establish repo using fake DB
	res := NewMessagesRepository(gormDB)

	_, err := res.UpdateMessageReadStatus(uint(12), message.Read)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	err = mock.ExpectationsWereMet()

	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// UpdateMessageReadStatus with expected failure.
func TestMessagesRepository_UpdateOrgUser_Failure(t *testing.T) {
	gormDB, db, mock := setupMockDBMessages(t)

	defer db.Close()

	// Fake message we are updating
	message := models.Messages{
		FromUsersID: 33,
		ToUsersID:   300000001,
		Subject:     "what what",
		Message:     "what what what",
		Read:        true,
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnError(errors.New("we failed!!!!!"))

	mock.ExpectBegin()
	mock.ExpectExec("INSERT").
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			message.FromUsersID,
			message.ToUsersID,
			message.Subject,
			message.Message,
			message.Read).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Establish repo using fake DB
	res := NewMessagesRepository(gormDB)

	_, err := res.UpdateMessageReadStatus(uint(999), message.Read)

	if err == nil {
		t.Errorf("We expect an error here: %s", err)
	}

	err = mock.ExpectationsWereMet()

	if err == nil {
		t.Errorf("Error expected: %s", err)
	}
}

// Ensures a message is deleted from mock DB.
func TestMessagesRepository_DeleteMessage_Success(t *testing.T) {
	gormDB, db, mock := setupMockDBMessages(t)

	defer db.Close()

	// Fake message we are deleting
	message := models.Messages{
		FromUsersID: 111,
		ToUsersID:   232,
		Subject:     "new volunteer, who dis",
		Message:     "idk u",
	}

	mockRows := sqlmock.NewRows([]string{"ID", "FromUsersID", "ToUsersID", "Subject", "Message"}).
		AddRow(uint(765), message.FromUsersID, message.ToUsersID, message.Subject, message.Message)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnRows(mockRows)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs(sqlmock.AnyArg(), uint(765)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Establish repo using fake DB
	res := NewMessagesRepository(gormDB)

	err := res.DeleteMessage(uint(765))

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	err = mock.ExpectationsWereMet()

	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// Expect error from DeleteOrgUser
func TestMessagesRepository_DeleteMessage_FindFailure(t *testing.T) {
	gormDB, db, mock := setupMockDBMessages(t)

	defer db.Close()

	// Fake user we are deleting
	message := models.Messages{
		FromUsersID: 702,
		ToUsersID:   808,
		Subject:     "what's poppin'",
		Message:     "brand new whip just hopped in",
	}

	mockRows := sqlmock.NewRows([]string{"ID", "FromUsersID", "ToUsersID", "Subject", "Message"}).
		AddRow(uint(123), message.FromUsersID, message.ToUsersID, message.Subject, message.Message)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnRows(mockRows)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs(sqlmock.AnyArg(), uint(123)).
		WillReturnError(errors.New("Oh noes :()"))
	mock.ExpectRollback()

	// Establish repo using fake DB
	res := NewMessagesRepository(gormDB)

	err := res.DeleteMessage(uint(123))

	if err == nil {
		t.Errorf("Error was expected while updating stats: %s", err)
	}

	err = mock.ExpectationsWereMet()

	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// Expect error inside DeleteOrgUser
func TestMessagesRepository_DeleteMessage_DeleteFailure(t *testing.T) {
	gormDB, db, mock := setupMockDBMessages(t)

	defer db.Close()

	// Create error on purpose, expect failure
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnError(errors.New("Uh-oh spaghetti-o..."))

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs(sqlmock.AnyArg()).
		WithArgs(uint(725)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Establish repo using fake DB
	res := NewMessagesRepository(gormDB)

	err := res.DeleteMessage(uint(725))

	if err == nil {
		t.Errorf("Error was expected while updating stats: %s", err)
	}

	err = mock.ExpectationsWereMet()

	if err == nil {
		t.Errorf("There was supposed to be an error: %s", err)
	}
}
