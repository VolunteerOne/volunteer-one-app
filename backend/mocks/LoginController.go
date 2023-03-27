// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// LoginController is an autogenerated mock type for the LoginController type
type LoginController struct {
	mock.Mock
}

// Login provides a mock function with given fields: c
func (_m *LoginController) Login(c *gin.Context) {
	_m.Called(c)
}

type mockConstructorTestingTNewLoginController interface {
	mock.TestingT
	Cleanup(func())
}

// NewLoginController creates a new instance of LoginController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLoginController(t mockConstructorTestingTNewLoginController) *LoginController {
	mock := &LoginController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
