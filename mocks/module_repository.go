// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/odpf/entropy/domain"
	mock "github.com/stretchr/testify/mock"
)

// ModuleRepository is an autogenerated mock type for the ModuleRepository type
type ModuleRepository struct {
	mock.Mock
}

type ModuleRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *ModuleRepository) EXPECT() *ModuleRepository_Expecter {
	return &ModuleRepository_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: id
func (_m *ModuleRepository) Get(id string) (domain.Module, error) {
	ret := _m.Called(id)

	var r0 domain.Module
	if rf, ok := ret.Get(0).(func(string) domain.Module); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Module)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModuleRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ModuleRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - id string
func (_e *ModuleRepository_Expecter) Get(id interface{}) *ModuleRepository_Get_Call {
	return &ModuleRepository_Get_Call{Call: _e.mock.On("Get", id)}
}

func (_c *ModuleRepository_Get_Call) Run(run func(id string)) *ModuleRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ModuleRepository_Get_Call) Return(_a0 domain.Module, _a1 error) *ModuleRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Register provides a mock function with given fields: module
func (_m *ModuleRepository) Register(module domain.Module) error {
	ret := _m.Called(module)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Module) error); ok {
		r0 = rf(module)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ModuleRepository_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type ModuleRepository_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//  - module domain.Module
func (_e *ModuleRepository_Expecter) Register(module interface{}) *ModuleRepository_Register_Call {
	return &ModuleRepository_Register_Call{Call: _e.mock.On("Register", module)}
}

func (_c *ModuleRepository_Register_Call) Run(run func(module domain.Module)) *ModuleRepository_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.Module))
	})
	return _c
}

func (_c *ModuleRepository_Register_Call) Return(_a0 error) *ModuleRepository_Register_Call {
	_c.Call.Return(_a0)
	return _c
}