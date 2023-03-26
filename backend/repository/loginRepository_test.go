package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestLoginRepository_FindUserFromEmail(t *testing.T) {
	email := "test@user.com"
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

	mockRows := sqlmock.NewRows([]string{"Email", "Password"}).
		AddRow(email, "fakepassword")

	mock.ExpectQuery("SELECT(.*)").
		WithArgs(email).
		WillReturnRows(mockRows)

	var user models.User
	res := NewLoginRepository(gormDB)

	// now we execute our method
	if user, err = res.FindUserFromEmail(email, user); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
