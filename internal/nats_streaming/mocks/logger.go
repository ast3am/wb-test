// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// logger is an autogenerated mock type for the logger type
type logger struct {
	mock.Mock
}

// DebugMsg provides a mock function with given fields: msg
func (_m *logger) DebugMsg(msg string) {
	_m.Called(msg)
}

// ErrorMsg provides a mock function with given fields: msg, err
func (_m *logger) ErrorMsg(msg string, err error) {
	_m.Called(msg, err)
}

// InfoMsgf provides a mock function with given fields: format, v
func (_m *logger) InfoMsgf(format string, v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

type mockConstructorTestingTnewLogger interface {
	mock.TestingT
	Cleanup(func())
}

// newLogger creates a new instance of logger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLogger(t mockConstructorTestingTnewLogger) *logger {
	mock := &logger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
