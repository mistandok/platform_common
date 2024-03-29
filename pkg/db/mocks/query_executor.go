// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	db "github.com/mistandok/platform_common/pkg/db"
	mock "github.com/stretchr/testify/mock"

	pgconn "github.com/jackc/pgx/v5/pgconn"

	pgx "github.com/jackc/pgx/v5"
)

// QueryExecutor is an autogenerated mock type for the QueryExecutor type
type QueryExecutor struct {
	mock.Mock
}

type QueryExecutor_Expecter struct {
	mock *mock.Mock
}

func (_m *QueryExecutor) EXPECT() *QueryExecutor_Expecter {
	return &QueryExecutor_Expecter{mock: &_m.Mock}
}

// ExecContext provides a mock function with given fields: ctx, q, args
func (_m *QueryExecutor) ExecContext(ctx context.Context, q db.Query, args ...interface{}) (pgconn.CommandTag, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, q)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExecContext")
	}

	var r0 pgconn.CommandTag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.Query, ...interface{}) (pgconn.CommandTag, error)); ok {
		return rf(ctx, q, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.Query, ...interface{}) pgconn.CommandTag); ok {
		r0 = rf(ctx, q, args...)
	} else {
		r0 = ret.Get(0).(pgconn.CommandTag)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.Query, ...interface{}) error); ok {
		r1 = rf(ctx, q, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryExecutor_ExecContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecContext'
type QueryExecutor_ExecContext_Call struct {
	*mock.Call
}

// ExecContext is a helper method to define mock.On call
//   - ctx context.Context
//   - q db.Query
//   - args ...interface{}
func (_e *QueryExecutor_Expecter) ExecContext(ctx interface{}, q interface{}, args ...interface{}) *QueryExecutor_ExecContext_Call {
	return &QueryExecutor_ExecContext_Call{Call: _e.mock.On("ExecContext",
		append([]interface{}{ctx, q}, args...)...)}
}

func (_c *QueryExecutor_ExecContext_Call) Run(run func(ctx context.Context, q db.Query, args ...interface{})) *QueryExecutor_ExecContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(db.Query), variadicArgs...)
	})
	return _c
}

func (_c *QueryExecutor_ExecContext_Call) Return(_a0 pgconn.CommandTag, _a1 error) *QueryExecutor_ExecContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QueryExecutor_ExecContext_Call) RunAndReturn(run func(context.Context, db.Query, ...interface{}) (pgconn.CommandTag, error)) *QueryExecutor_ExecContext_Call {
	_c.Call.Return(run)
	return _c
}

// QueryContext provides a mock function with given fields: ctx, q, args
func (_m *QueryExecutor) QueryContext(ctx context.Context, q db.Query, args ...interface{}) (pgx.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, q)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryContext")
	}

	var r0 pgx.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.Query, ...interface{}) (pgx.Rows, error)); ok {
		return rf(ctx, q, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.Query, ...interface{}) pgx.Rows); ok {
		r0 = rf(ctx, q, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.Query, ...interface{}) error); ok {
		r1 = rf(ctx, q, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryExecutor_QueryContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryContext'
type QueryExecutor_QueryContext_Call struct {
	*mock.Call
}

// QueryContext is a helper method to define mock.On call
//   - ctx context.Context
//   - q db.Query
//   - args ...interface{}
func (_e *QueryExecutor_Expecter) QueryContext(ctx interface{}, q interface{}, args ...interface{}) *QueryExecutor_QueryContext_Call {
	return &QueryExecutor_QueryContext_Call{Call: _e.mock.On("QueryContext",
		append([]interface{}{ctx, q}, args...)...)}
}

func (_c *QueryExecutor_QueryContext_Call) Run(run func(ctx context.Context, q db.Query, args ...interface{})) *QueryExecutor_QueryContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(db.Query), variadicArgs...)
	})
	return _c
}

func (_c *QueryExecutor_QueryContext_Call) Return(_a0 pgx.Rows, _a1 error) *QueryExecutor_QueryContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QueryExecutor_QueryContext_Call) RunAndReturn(run func(context.Context, db.Query, ...interface{}) (pgx.Rows, error)) *QueryExecutor_QueryContext_Call {
	_c.Call.Return(run)
	return _c
}

// QueryRowContext provides a mock function with given fields: ctx, q, args
func (_m *QueryExecutor) QueryRowContext(ctx context.Context, q db.Query, args ...interface{}) pgx.Row {
	var _ca []interface{}
	_ca = append(_ca, ctx, q)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryRowContext")
	}

	var r0 pgx.Row
	if rf, ok := ret.Get(0).(func(context.Context, db.Query, ...interface{}) pgx.Row); ok {
		r0 = rf(ctx, q, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Row)
		}
	}

	return r0
}

// QueryExecutor_QueryRowContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryRowContext'
type QueryExecutor_QueryRowContext_Call struct {
	*mock.Call
}

// QueryRowContext is a helper method to define mock.On call
//   - ctx context.Context
//   - q db.Query
//   - args ...interface{}
func (_e *QueryExecutor_Expecter) QueryRowContext(ctx interface{}, q interface{}, args ...interface{}) *QueryExecutor_QueryRowContext_Call {
	return &QueryExecutor_QueryRowContext_Call{Call: _e.mock.On("QueryRowContext",
		append([]interface{}{ctx, q}, args...)...)}
}

func (_c *QueryExecutor_QueryRowContext_Call) Run(run func(ctx context.Context, q db.Query, args ...interface{})) *QueryExecutor_QueryRowContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(db.Query), variadicArgs...)
	})
	return _c
}

func (_c *QueryExecutor_QueryRowContext_Call) Return(_a0 pgx.Row) *QueryExecutor_QueryRowContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryExecutor_QueryRowContext_Call) RunAndReturn(run func(context.Context, db.Query, ...interface{}) pgx.Row) *QueryExecutor_QueryRowContext_Call {
	_c.Call.Return(run)
	return _c
}

// NewQueryExecutor creates a new instance of QueryExecutor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQueryExecutor(t interface {
	mock.TestingT
	Cleanup(func())
}) *QueryExecutor {
	mock := &QueryExecutor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
