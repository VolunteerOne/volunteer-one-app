package service

import (
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/repository"
)

type LoginService interface {
	FindUserFromEmail(string, models.User) (models.User, error)
}

type loginService struct {
	loginRepository repository.LoginRepository
}

// Instantiated in router.go
func NewLoginService(r repository.LoginRepository) LoginService {
	return loginService{
		loginRepository: r,
	}
}

func (l loginService) FindUserFromEmail(email string, user models.User) (models.User, error) {
    return l.loginRepository.FindUserFromEmail(email, user)
}
