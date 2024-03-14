// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// Lock is an autogenerated mock type for the Lock type
type Lock struct {
	mock.Mock
}

type Lock_Expecter struct {
	mock *mock.Mock
}

func (_m *Lock) EXPECT() *Lock_Expecter {
	return &Lock_Expecter{mock: &_m.Mock}
}

// Block provides a mock function with given fields: t, callback
func (_m *Lock) Block(t time.Duration, callback ...func()) bool {
	_va := make([]interface{}, len(callback))
	for _i := range callback {
		_va[_i] = callback[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, t)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Block")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(time.Duration, ...func()) bool); ok {
		r0 = rf(t, callback...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Lock_Block_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Block'
type Lock_Block_Call struct {
	*mock.Call
}

// Block is a helper method to define mock.On call
//   - t time.Duration
//   - callback ...func()
func (_e *Lock_Expecter) Block(t interface{}, callback ...interface{}) *Lock_Block_Call {
	return &Lock_Block_Call{Call: _e.mock.On("Block",
		append([]interface{}{t}, callback...)...)}
}

func (_c *Lock_Block_Call) Run(run func(t time.Duration, callback ...func())) *Lock_Block_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(), len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(func())
			}
		}
		run(args[0].(time.Duration), variadicArgs...)
	})
	return _c
}

func (_c *Lock_Block_Call) Return(_a0 bool) *Lock_Block_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Lock_Block_Call) RunAndReturn(run func(time.Duration, ...func()) bool) *Lock_Block_Call {
	_c.Call.Return(run)
	return _c
}

// ForceRelease provides a mock function with given fields:
func (_m *Lock) ForceRelease() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ForceRelease")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Lock_ForceRelease_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForceRelease'
type Lock_ForceRelease_Call struct {
	*mock.Call
}

// ForceRelease is a helper method to define mock.On call
func (_e *Lock_Expecter) ForceRelease() *Lock_ForceRelease_Call {
	return &Lock_ForceRelease_Call{Call: _e.mock.On("ForceRelease")}
}

func (_c *Lock_ForceRelease_Call) Run(run func()) *Lock_ForceRelease_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Lock_ForceRelease_Call) Return(_a0 bool) *Lock_ForceRelease_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Lock_ForceRelease_Call) RunAndReturn(run func() bool) *Lock_ForceRelease_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: callback
func (_m *Lock) Get(callback ...func()) bool {
	_va := make([]interface{}, len(callback))
	for _i := range callback {
		_va[_i] = callback[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(...func()) bool); ok {
		r0 = rf(callback...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Lock_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Lock_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - callback ...func()
func (_e *Lock_Expecter) Get(callback ...interface{}) *Lock_Get_Call {
	return &Lock_Get_Call{Call: _e.mock.On("Get",
		append([]interface{}{}, callback...)...)}
}

func (_c *Lock_Get_Call) Run(run func(callback ...func())) *Lock_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(), len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(func())
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Lock_Get_Call) Return(_a0 bool) *Lock_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Lock_Get_Call) RunAndReturn(run func(...func()) bool) *Lock_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Release provides a mock function with given fields:
func (_m *Lock) Release() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Release")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Lock_Release_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Release'
type Lock_Release_Call struct {
	*mock.Call
}

// Release is a helper method to define mock.On call
func (_e *Lock_Expecter) Release() *Lock_Release_Call {
	return &Lock_Release_Call{Call: _e.mock.On("Release")}
}

func (_c *Lock_Release_Call) Run(run func()) *Lock_Release_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Lock_Release_Call) Return(_a0 bool) *Lock_Release_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Lock_Release_Call) RunAndReturn(run func() bool) *Lock_Release_Call {
	_c.Call.Return(run)
	return _c
}

// NewLock creates a new instance of Lock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLock(t interface {
	mock.TestingT
	Cleanup(func())
}) *Lock {
	mock := &Lock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
