// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// DB is an autogenerated mock type for the DB type
type DB struct {
	mock.Mock
}

type DB_Expecter struct {
	mock *mock.Mock
}

func (_m *DB) EXPECT() *DB_Expecter {
	return &DB_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *DB) Close() error {
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

// DB_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type DB_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *DB_Expecter) Close() *DB_Close_Call {
	return &DB_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *DB_Close_Call) Run(run func()) *DB_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DB_Close_Call) Return(_a0 error) *DB_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_Close_Call) RunAndReturn(run func() error) *DB_Close_Call {
	_c.Call.Return(run)
	return _c
}

// DoContext provides a mock function with given fields: ctx, commandName, args
func (_m *DB) DoContext(ctx context.Context, commandName string, args ...interface{}) (interface{}, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, commandName)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DoContext")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (interface{}, error)); ok {
		return rf(ctx, commandName, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) interface{}); ok {
		r0 = rf(ctx, commandName, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, commandName, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_DoContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DoContext'
type DB_DoContext_Call struct {
	*mock.Call
}

// DoContext is a helper method to define mock.On call
//   - ctx context.Context
//   - commandName string
//   - args ...interface{}
func (_e *DB_Expecter) DoContext(ctx interface{}, commandName interface{}, args ...interface{}) *DB_DoContext_Call {
	return &DB_DoContext_Call{Call: _e.mock.On("DoContext",
		append([]interface{}{ctx, commandName}, args...)...)}
}

func (_c *DB_DoContext_Call) Run(run func(ctx context.Context, commandName string, args ...interface{})) *DB_DoContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *DB_DoContext_Call) Return(reply interface{}, err error) *DB_DoContext_Call {
	_c.Call.Return(reply, err)
	return _c
}

func (_c *DB_DoContext_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (interface{}, error)) *DB_DoContext_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields: reply, err
func (_m *DB) String(reply interface{}, err error) (string, error) {
	ret := _m.Called(reply, err)

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, error) (string, error)); ok {
		return rf(reply, err)
	}
	if rf, ok := ret.Get(0).(func(interface{}, error) string); ok {
		r0 = rf(reply, err)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(interface{}, error) error); ok {
		r1 = rf(reply, err)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type DB_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
//   - reply interface{}
//   - err error
func (_e *DB_Expecter) String(reply interface{}, err interface{}) *DB_String_Call {
	return &DB_String_Call{Call: _e.mock.On("String", reply, err)}
}

func (_c *DB_String_Call) Run(run func(reply interface{}, err error)) *DB_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}), args[1].(error))
	})
	return _c
}

func (_c *DB_String_Call) Return(_a0 string, _a1 error) *DB_String_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_String_Call) RunAndReturn(run func(interface{}, error) (string, error)) *DB_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewDB creates a new instance of DB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *DB {
	mock := &DB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
