// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	console "github.com/goravel/framework/contracts/console"
	command "github.com/goravel/framework/contracts/console/command"

	mock "github.com/stretchr/testify/mock"
)

// Command is an autogenerated mock type for the Command type
type Command struct {
	mock.Mock
}

type Command_Expecter struct {
	mock *mock.Mock
}

func (_m *Command) EXPECT() *Command_Expecter {
	return &Command_Expecter{mock: &_m.Mock}
}

// Description provides a mock function with given fields:
func (_m *Command) Description() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Description")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Command_Description_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Description'
type Command_Description_Call struct {
	*mock.Call
}

// Description is a helper method to define mock.On call
func (_e *Command_Expecter) Description() *Command_Description_Call {
	return &Command_Description_Call{Call: _e.mock.On("Description")}
}

func (_c *Command_Description_Call) Run(run func()) *Command_Description_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Command_Description_Call) Return(_a0 string) *Command_Description_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Command_Description_Call) RunAndReturn(run func() string) *Command_Description_Call {
	_c.Call.Return(run)
	return _c
}

// Extend provides a mock function with given fields:
func (_m *Command) Extend() command.Extend {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Extend")
	}

	var r0 command.Extend
	if rf, ok := ret.Get(0).(func() command.Extend); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(command.Extend)
	}

	return r0
}

// Command_Extend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Extend'
type Command_Extend_Call struct {
	*mock.Call
}

// Extend is a helper method to define mock.On call
func (_e *Command_Expecter) Extend() *Command_Extend_Call {
	return &Command_Extend_Call{Call: _e.mock.On("Extend")}
}

func (_c *Command_Extend_Call) Run(run func()) *Command_Extend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Command_Extend_Call) Return(_a0 command.Extend) *Command_Extend_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Command_Extend_Call) RunAndReturn(run func() command.Extend) *Command_Extend_Call {
	_c.Call.Return(run)
	return _c
}

// Handle provides a mock function with given fields: ctx
func (_m *Command) Handle(ctx console.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(console.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Command_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type Command_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
//   - ctx console.Context
func (_e *Command_Expecter) Handle(ctx interface{}) *Command_Handle_Call {
	return &Command_Handle_Call{Call: _e.mock.On("Handle", ctx)}
}

func (_c *Command_Handle_Call) Run(run func(ctx console.Context)) *Command_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(console.Context))
	})
	return _c
}

func (_c *Command_Handle_Call) Return(_a0 error) *Command_Handle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Command_Handle_Call) RunAndReturn(run func(console.Context) error) *Command_Handle_Call {
	_c.Call.Return(run)
	return _c
}

// Signature provides a mock function with given fields:
func (_m *Command) Signature() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Signature")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Command_Signature_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Signature'
type Command_Signature_Call struct {
	*mock.Call
}

// Signature is a helper method to define mock.On call
func (_e *Command_Expecter) Signature() *Command_Signature_Call {
	return &Command_Signature_Call{Call: _e.mock.On("Signature")}
}

func (_c *Command_Signature_Call) Run(run func()) *Command_Signature_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Command_Signature_Call) Return(_a0 string) *Command_Signature_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Command_Signature_Call) RunAndReturn(run func() string) *Command_Signature_Call {
	_c.Call.Return(run)
	return _c
}

// NewCommand creates a new instance of Command. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCommand(t interface {
	mock.TestingT
	Cleanup(func())
}) *Command {
	mock := &Command{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
