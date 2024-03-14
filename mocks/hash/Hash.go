// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Hash is an autogenerated mock type for the Hash type
type Hash struct {
	mock.Mock
}

type Hash_Expecter struct {
	mock *mock.Mock
}

func (_m *Hash) EXPECT() *Hash_Expecter {
	return &Hash_Expecter{mock: &_m.Mock}
}

// Check provides a mock function with given fields: value, hashedValue
func (_m *Hash) Check(value string, hashedValue string) bool {
	ret := _m.Called(value, hashedValue)

	if len(ret) == 0 {
		panic("no return value specified for Check")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(value, hashedValue)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Hash_Check_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Check'
type Hash_Check_Call struct {
	*mock.Call
}

// Check is a helper method to define mock.On call
//   - value string
//   - hashedValue string
func (_e *Hash_Expecter) Check(value interface{}, hashedValue interface{}) *Hash_Check_Call {
	return &Hash_Check_Call{Call: _e.mock.On("Check", value, hashedValue)}
}

func (_c *Hash_Check_Call) Run(run func(value string, hashedValue string)) *Hash_Check_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Hash_Check_Call) Return(_a0 bool) *Hash_Check_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Hash_Check_Call) RunAndReturn(run func(string, string) bool) *Hash_Check_Call {
	_c.Call.Return(run)
	return _c
}

// Make provides a mock function with given fields: value
func (_m *Hash) Make(value string) (string, error) {
	ret := _m.Called(value)

	if len(ret) == 0 {
		panic("no return value specified for Make")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(value)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Hash_Make_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Make'
type Hash_Make_Call struct {
	*mock.Call
}

// Make is a helper method to define mock.On call
//   - value string
func (_e *Hash_Expecter) Make(value interface{}) *Hash_Make_Call {
	return &Hash_Make_Call{Call: _e.mock.On("Make", value)}
}

func (_c *Hash_Make_Call) Run(run func(value string)) *Hash_Make_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Hash_Make_Call) Return(_a0 string, _a1 error) *Hash_Make_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Hash_Make_Call) RunAndReturn(run func(string) (string, error)) *Hash_Make_Call {
	_c.Call.Return(run)
	return _c
}

// NeedsRehash provides a mock function with given fields: hashedValue
func (_m *Hash) NeedsRehash(hashedValue string) bool {
	ret := _m.Called(hashedValue)

	if len(ret) == 0 {
		panic("no return value specified for NeedsRehash")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(hashedValue)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Hash_NeedsRehash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NeedsRehash'
type Hash_NeedsRehash_Call struct {
	*mock.Call
}

// NeedsRehash is a helper method to define mock.On call
//   - hashedValue string
func (_e *Hash_Expecter) NeedsRehash(hashedValue interface{}) *Hash_NeedsRehash_Call {
	return &Hash_NeedsRehash_Call{Call: _e.mock.On("NeedsRehash", hashedValue)}
}

func (_c *Hash_NeedsRehash_Call) Run(run func(hashedValue string)) *Hash_NeedsRehash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Hash_NeedsRehash_Call) Return(_a0 bool) *Hash_NeedsRehash_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Hash_NeedsRehash_Call) RunAndReturn(run func(string) bool) *Hash_NeedsRehash_Call {
	_c.Call.Return(run)
	return _c
}

// NewHash creates a new instance of Hash. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHash(t interface {
	mock.TestingT
	Cleanup(func())
}) *Hash {
	mock := &Hash{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
