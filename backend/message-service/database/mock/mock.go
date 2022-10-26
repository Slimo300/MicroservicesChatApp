// Code generated by mockery v2.8.0. DO NOT EDIT.

package mock

import (
	time "time"

	models "github.com/Slimo300/MicroservicesChatApp/backend/message-service/models"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockMessageDB is an autogenerated mock type for the DBLayer type
type MockMessageDB struct {
	mock.Mock
}

// AddMessage provides a mock function with given fields: userID, groupID, nick, text, when
func (_m *MockMessageDB) AddMessage(userID uuid.UUID, groupID uuid.UUID, nick string, text string, when time.Time) error {
	ret := _m.Called(userID, groupID, nick, text, when)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, uuid.UUID, string, string, time.Time) error); ok {
		r0 = rf(userID, groupID, nick, text, when)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMessageForEveryone provides a mock function with given fields: userID, messageID
func (_m *MockMessageDB) DeleteMessageForEveryone(userID uuid.UUID, messageID uuid.UUID) (models.Message, error) {
	ret := _m.Called(userID, messageID)

	var r0 models.Message
	if rf, ok := ret.Get(0).(func(uuid.UUID, uuid.UUID) models.Message); ok {
		r0 = rf(userID, messageID)
	} else {
		r0 = ret.Get(0).(models.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID, uuid.UUID) error); ok {
		r1 = rf(userID, messageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMessageForYourself provides a mock function with given fields: userID, messageID
func (_m *MockMessageDB) DeleteMessageForYourself(userID uuid.UUID, messageID uuid.UUID) (models.Message, error) {
	ret := _m.Called(userID, messageID)

	var r0 models.Message
	if rf, ok := ret.Get(0).(func(uuid.UUID, uuid.UUID) models.Message); ok {
		r0 = rf(userID, messageID)
	} else {
		r0 = ret.Get(0).(models.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID, uuid.UUID) error); ok {
		r1 = rf(userID, messageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroupMessages provides a mock function with given fields: userID, groupID, offset, num
func (_m *MockMessageDB) GetGroupMessages(userID uuid.UUID, groupID uuid.UUID, offset int, num int) ([]models.Message, error) {
	ret := _m.Called(userID, groupID, offset, num)

	var r0 []models.Message
	if rf, ok := ret.Get(0).(func(uuid.UUID, uuid.UUID, int, int) []models.Message); ok {
		r0 = rf(userID, groupID, offset, num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID, uuid.UUID, int, int) error); ok {
		r1 = rf(userID, groupID, offset, num)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
