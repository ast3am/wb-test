// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	"github.com/ast3am/wb-test/internal/models"
	"github.com/stretchr/testify/mock"
)

// cache is an autogenerated mock type for the cache type
type cache struct {
	mock.Mock
}

// GetOrdersById provides a mock function with given fields: Uid
func (_m *cache) GetOrdersById(Uid string) models.Orders {
	ret := _m.Called(Uid)

	var r0 models.Orders
	if rf, ok := ret.Get(0).(func(string) models.Orders); ok {
		r0 = rf(Uid)
	} else {
		r0 = ret.Get(0).(models.Orders)
	}

	return r0
}

type mockConstructorTestingTnewCache interface {
	mock.TestingT
	Cleanup(func())
}

// newCache creates a new instance of cache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCache(t mockConstructorTestingTnewCache) *cache {
	mock := &cache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
