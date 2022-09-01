// Code generated by mockery v2.14.0. DO NOT EDIT.

package events

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	mock "github.com/stretchr/testify/mock"
)

// mockTbAPI is an autogenerated mock type for the tbAPI type
type mockTbAPI struct {
	mock.Mock
}

// GetChat provides a mock function with given fields: config
func (_m *mockTbAPI) GetChat(config tgbotapi.ChatInfoConfig) (tgbotapi.Chat, error) {
	ret := _m.Called(config)

	var r0 tgbotapi.Chat
	if rf, ok := ret.Get(0).(func(tgbotapi.ChatInfoConfig) tgbotapi.Chat); ok {
		r0 = rf(config)
	} else {
		r0 = ret.Get(0).(tgbotapi.Chat)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(tgbotapi.ChatInfoConfig) error); ok {
		r1 = rf(config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUpdatesChan provides a mock function with given fields: config
func (_m *mockTbAPI) GetUpdatesChan(config tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel {
	ret := _m.Called(config)

	var r0 tgbotapi.UpdatesChannel
	if rf, ok := ret.Get(0).(func(tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel); ok {
		r0 = rf(config)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tgbotapi.UpdatesChannel)
		}
	}

	return r0
}

// Send provides a mock function with given fields: c
func (_m *mockTbAPI) Send(c tgbotapi.Chattable) (tgbotapi.Message, error) {
	ret := _m.Called(c)

	var r0 tgbotapi.Message
	if rf, ok := ret.Get(0).(func(tgbotapi.Chattable) tgbotapi.Message); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(tgbotapi.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(tgbotapi.Chattable) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewMockTbAPI interface {
	mock.TestingT
	Cleanup(func())
}

// newMockTbAPI creates a new instance of mockTbAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockTbAPI(t mockConstructorTestingTnewMockTbAPI) *mockTbAPI {
	mock := &mockTbAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
