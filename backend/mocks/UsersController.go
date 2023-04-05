package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

type UsersController struct {
	mock.Mock
}

func (_m *UsersController) Create(c *gin.Context) {
	_m.Called(c)
}

func (_m *UsersController) All(c *gin.Context) {
	_m.Called(c)
}

func (_m *UsersController) One(c *gin.Context) {
	_m.Called(c)
}

func (_m *UsersController) Update(c *gin.Context) {
	_m.Called(c)
}

func (_m *UsersController) Delete(c *gin.Context) {
	_m.Called(c)
}

type mockConstructorTestingTNewUsersController interface {
	mock.TestingT
	Cleanup(func())
}

func NewUsersController(t mockConstructorTestingTNewUsersController) *UsersController {
	mock := &UsersController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}
