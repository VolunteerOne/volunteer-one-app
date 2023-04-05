package mocks

import (
	models "github.com/VolunteerOne/volunteer-one-app/backend/models"
	mock "github.com/stretchr/testify/mock"
)

type UsersService struct {
	mock.Mock
}



func (_m *UsersService) CreateUser(_a0 models.Users) (models.Users, error) {
	ret := _m.Called(_a0)

	var r0 models.Users
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Users) (models.Users, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(models.Users) models.Users); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.Users)
	}

	if rf, ok := ret.Get(1).(func(models.Users) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// func (_m *UsersService) UpdateUser(_a0 models.Users) (models.Users, error) {
	
// }

// func (_m *UsersService) DeleteUser(_a0 models.Users) (models.Users, error) {
	
// }

type mockConstructorTestingTNewUsersService interface {
	mock.TestingT
	Cleanup(func())
}

func NewUsersService(t mockConstructorTestingTNewUsersService) *UsersService {
	mock := &UsersService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}
