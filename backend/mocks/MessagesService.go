// Code generated by mockery v2.26.0. DO NOT EDIT.

package mocks

import (
	models "github.com/VolunteerOne/volunteer-one-app/backend/models"
	mock "github.com/stretchr/testify/mock"
)

// MessagesService is an autogenerated mock type for the MessagesService type
type MessagesService struct {
	mock.Mock
}

// CreateMessage provides a mock function with given fields: _a0
func (_m *MessagesService) CreateMessage(_a0 models.Messages) (models.Messages, error) {
	ret := _m.Called(_a0)

	var r0 models.Messages
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Messages) (models.Messages, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(models.Messages) models.Messages); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.Messages)
	}

	if rf, ok := ret.Get(1).(func(models.Messages) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMessage provides a mock function with given fields: _a0
func (_m *MessagesService) DeleteMessage(_a0 uint) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindMessage provides a mock function with given fields: _a0
func (_m *MessagesService) FindMessage(_a0 uint) (models.Messages, error) {
	ret := _m.Called(_a0)

	var r0 models.Messages
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (models.Messages, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uint) models.Messages); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.Messages)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllMessagesForUser provides a mock function with given fields: _a0
func (_m *MessagesService) ListAllMessagesForUser(_a0 uint) ([]models.Messages, error) {
	ret := _m.Called(_a0)

	var r0 []models.Messages
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]models.Messages, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uint) []models.Messages); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Messages)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMessageReadStatus provides a mock function with given fields: _a0, _a1
func (_m *MessagesService) UpdateMessageReadStatus(_a0 uint, _a1 bool) (models.Messages, error) {
	ret := _m.Called(_a0, _a1)

	var r0 models.Messages
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, bool) (models.Messages, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(uint, bool) models.Messages); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(models.Messages)
	}

	if rf, ok := ret.Get(1).(func(uint, bool) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMessagesService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMessagesService creates a new instance of MessagesService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMessagesService(t mockConstructorTestingTNewMessagesService) *MessagesService {
	mock := &MessagesService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}