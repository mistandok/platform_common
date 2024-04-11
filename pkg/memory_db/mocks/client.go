// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	memory_db "github.com/mistandok/platform_common/pkg/memory_db"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_Expecter struct {
	mock *mock.Mock
}

func (_m *Client) EXPECT() *Client_Expecter {
	return &Client_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *Client) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type Client_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *Client_Expecter) Close() *Client_Close_Call {
	return &Client_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *Client_Close_Call) Run(run func()) *Client_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_Close_Call) Return(_a0 error) *Client_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Close_Call) RunAndReturn(run func() error) *Client_Close_Call {
	_c.Call.Return(run)
	return _c
}

// DB provides a mock function with given fields:
func (_m *Client) DB() memory_db.DB {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DB")
	}

	var r0 memory_db.DB
	if rf, ok := ret.Get(0).(func() memory_db.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(memory_db.DB)
		}
	}

	return r0
}

// Client_DB_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DB'
type Client_DB_Call struct {
	*mock.Call
}

// DB is a helper method to define mock.On call
func (_e *Client_Expecter) DB() *Client_DB_Call {
	return &Client_DB_Call{Call: _e.mock.On("DB")}
}

func (_c *Client_DB_Call) Run(run func()) *Client_DB_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_DB_Call) Return(_a0 memory_db.DB) *Client_DB_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_DB_Call) RunAndReturn(run func() memory_db.DB) *Client_DB_Call {
	_c.Call.Return(run)
	return _c
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
