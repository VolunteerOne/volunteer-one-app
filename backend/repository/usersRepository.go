package repository

import (
	"log"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"gorm.io/gorm"
)

type UsersRepository interface {
	CreateUser(user models.Users) (models.Users, error)
	OneUser(id string, user models.Users) (models.Users, error)
	UpdateUser(user models.Users) (models.Users, error)
	DeleteUser(user models.Users) (models.Users, error)
}

type usersRepository struct {
	DB *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return usersRepository{
		DB: db,
	}
}

// Add the user to the DB
func (u usersRepository) CreateUser(user models.Users) (models.Users, error) {
	log.Println("[UsersRepository] Create user...")

	err := u.DB.Create(&user).Error

	return user, err
}


// Add the user to the DB
func (u usersRepository) OneUser(id string, user models.Users) (models.Users, error) {
	log.Println("[UsersRepository] One user...")

	err := u.DB.First(&user, id).Error

	return user, err
}

// Update user to the DB
func (u usersRepository) UpdateUser(user models.Users) (models.Users, error) {
	log.Println("[UsersRepository] Update User...")

	err := u.DB.Update("Update", &user).Error

	return user, err
}

// Delete user from the DB
func (u usersRepository) DeleteUser(user models.Users) (models.Users, error) {
	log.Println("[UsersRepository] Delete User...")

	err := u.DB.Delete(&user).Error

	return user, err
}
