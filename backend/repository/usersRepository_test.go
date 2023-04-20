package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestUsersRepository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	// choose insert and mock the args
	// will return result has just random
	mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(),
		sqlmock.AnyArg(), "", "", "", "", "", "", "", 0, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	res := NewUsersRepository(gormDB)

	var user models.Users

	// now we execute our method
	if user, err = res.CreateUser(user); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUsersRepository_DeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()

	res := NewUsersRepository(gormDB)

	// Give model some data
	var user models.Users
	user.ID = 7 // This is what actually matters, since it's how we target the row to delete
	user.FirstName = "Deletus"
	user.LastName = "Testus"

	//
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `deleted_at`=? WHERE `users`.`id` = ?")).WithArgs(sqlmock.AnyArg(), 7).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// now we execute our method
	if user, err = res.DeleteUser(user); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
