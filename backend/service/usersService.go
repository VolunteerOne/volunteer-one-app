package service

import (
	"log"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/repository"
)

type UsersService interface {
	CreateUser(user models.Users) (models.Users, error)
	OneUser(id string, user models.Users) (models.Users, error)
	UpdateUser(user models.Users) (models.Users, error)
	DeleteUser(user models.Users) (models.Users, error)
}

type usersService struct {
	usersRepository repository.UsersRepository
}

func NewUsersService(r repository.UsersRepository) UsersService {
	return usersService{
		usersRepository: r,
	}
}

func (u usersService) CreateUser(user models.Users) (models.Users, error) {
	log.Println("[UsersService] Create user...")

	return u.usersRepository.CreateUser(user)
}


func (u usersService) OneUser(id string, user models.Users) (models.Users, error) {
	log.Println("[UsersService] Get One User...")

	return u.usersRepository.OneUser(id, user)
}

func (u usersService) UpdateUser(user models.Users) (models.Users, error) {
	log.Println("[UsersService] Update User...")

	return u.usersRepository.UpdateUser(user)
}

func (u usersService) DeleteUser(user models.Users) (models.Users, error) {
	log.Println("[UsersService] Delete User...")

	return u.usersRepository.DeleteUser(user)
}
