// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	context "context"

	resource "github.com/odpf/entropy/core/resource"
	mock "github.com/stretchr/testify/mock"
)

// ResourceRepository is an autogenerated mock type for the Repository type
type ResourceRepository struct {
	mock.Mock
}

type ResourceRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *ResourceRepository) EXPECT() *ResourceRepository_Expecter {
	return &ResourceRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, r, hooks
func (_m *ResourceRepository) Create(ctx context.Context, r resource.Resource, hooks ...resource.MutationHook) error {
	_va := make([]interface{}, len(hooks))
	for _i := range hooks {
		_va[_i] = hooks[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, r)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, resource.Resource, ...resource.MutationHook) error); ok {
		r0 = rf(ctx, r, hooks...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResourceRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ResourceRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//  - ctx context.Context
//  - r resource.Resource
//  - hooks ...resource.MutationHook
func (_e *ResourceRepository_Expecter) Create(ctx interface{}, r interface{}, hooks ...interface{}) *ResourceRepository_Create_Call {
	return &ResourceRepository_Create_Call{Call: _e.mock.On("Create",
		append([]interface{}{ctx, r}, hooks...)...)}
}

func (_c *ResourceRepository_Create_Call) Run(run func(ctx context.Context, r resource.Resource, hooks ...resource.MutationHook)) *ResourceRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]resource.MutationHook, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(resource.MutationHook)
			}
		}
		run(args[0].(context.Context), args[1].(resource.Resource), variadicArgs...)
	})
	return _c
}

func (_c *ResourceRepository_Create_Call) Return(_a0 error) *ResourceRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

// Delete provides a mock function with given fields: ctx, urn, hooks
func (_m *ResourceRepository) Delete(ctx context.Context, urn string, hooks ...resource.MutationHook) error {
	_va := make([]interface{}, len(hooks))
	for _i := range hooks {
		_va[_i] = hooks[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, urn)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...resource.MutationHook) error); ok {
		r0 = rf(ctx, urn, hooks...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResourceRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ResourceRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//  - ctx context.Context
//  - urn string
//  - hooks ...resource.MutationHook
func (_e *ResourceRepository_Expecter) Delete(ctx interface{}, urn interface{}, hooks ...interface{}) *ResourceRepository_Delete_Call {
	return &ResourceRepository_Delete_Call{Call: _e.mock.On("Delete",
		append([]interface{}{ctx, urn}, hooks...)...)}
}

func (_c *ResourceRepository_Delete_Call) Run(run func(ctx context.Context, urn string, hooks ...resource.MutationHook)) *ResourceRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]resource.MutationHook, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(resource.MutationHook)
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *ResourceRepository_Delete_Call) Return(_a0 error) *ResourceRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// DoPending provides a mock function with given fields: ctx, fn
func (_m *ResourceRepository) DoPending(ctx context.Context, fn resource.PendingHandler) error {
	ret := _m.Called(ctx, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, resource.PendingHandler) error); ok {
		r0 = rf(ctx, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResourceRepository_DoPending_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DoPending'
type ResourceRepository_DoPending_Call struct {
	*mock.Call
}

// DoPending is a helper method to define mock.On call
//  - ctx context.Context
//  - fn resource.PendingHandler
func (_e *ResourceRepository_Expecter) DoPending(ctx interface{}, fn interface{}) *ResourceRepository_DoPending_Call {
	return &ResourceRepository_DoPending_Call{Call: _e.mock.On("DoPending", ctx, fn)}
}

func (_c *ResourceRepository_DoPending_Call) Run(run func(ctx context.Context, fn resource.PendingHandler)) *ResourceRepository_DoPending_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(resource.PendingHandler))
	})
	return _c
}

func (_c *ResourceRepository_DoPending_Call) Return(_a0 error) *ResourceRepository_DoPending_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetByURN provides a mock function with given fields: ctx, urn
func (_m *ResourceRepository) GetByURN(ctx context.Context, urn string) (*resource.Resource, error) {
	ret := _m.Called(ctx, urn)

	var r0 *resource.Resource
	if rf, ok := ret.Get(0).(func(context.Context, string) *resource.Resource); ok {
		r0 = rf(ctx, urn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resource.Resource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, urn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceRepository_GetByURN_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByURN'
type ResourceRepository_GetByURN_Call struct {
	*mock.Call
}

// GetByURN is a helper method to define mock.On call
//  - ctx context.Context
//  - urn string
func (_e *ResourceRepository_Expecter) GetByURN(ctx interface{}, urn interface{}) *ResourceRepository_GetByURN_Call {
	return &ResourceRepository_GetByURN_Call{Call: _e.mock.On("GetByURN", ctx, urn)}
}

func (_c *ResourceRepository_GetByURN_Call) Run(run func(ctx context.Context, urn string)) *ResourceRepository_GetByURN_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ResourceRepository_GetByURN_Call) Return(_a0 *resource.Resource, _a1 error) *ResourceRepository_GetByURN_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// List provides a mock function with given fields: ctx, filter
func (_m *ResourceRepository) List(ctx context.Context, filter map[string]string) ([]*resource.Resource, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*resource.Resource
	if rf, ok := ret.Get(0).(func(context.Context, map[string]string) []*resource.Resource); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*resource.Resource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, map[string]string) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceRepository_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type ResourceRepository_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//  - ctx context.Context
//  - filter map[string]string
func (_e *ResourceRepository_Expecter) List(ctx interface{}, filter interface{}) *ResourceRepository_List_Call {
	return &ResourceRepository_List_Call{Call: _e.mock.On("List", ctx, filter)}
}

func (_c *ResourceRepository_List_Call) Run(run func(ctx context.Context, filter map[string]string)) *ResourceRepository_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(map[string]string))
	})
	return _c
}

func (_c *ResourceRepository_List_Call) Return(_a0 []*resource.Resource, _a1 error) *ResourceRepository_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Update provides a mock function with given fields: ctx, r, hooks
func (_m *ResourceRepository) Update(ctx context.Context, r resource.Resource, hooks ...resource.MutationHook) error {
	_va := make([]interface{}, len(hooks))
	for _i := range hooks {
		_va[_i] = hooks[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, r)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, resource.Resource, ...resource.MutationHook) error); ok {
		r0 = rf(ctx, r, hooks...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResourceRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ResourceRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//  - ctx context.Context
//  - r resource.Resource
//  - hooks ...resource.MutationHook
func (_e *ResourceRepository_Expecter) Update(ctx interface{}, r interface{}, hooks ...interface{}) *ResourceRepository_Update_Call {
	return &ResourceRepository_Update_Call{Call: _e.mock.On("Update",
		append([]interface{}{ctx, r}, hooks...)...)}
}

func (_c *ResourceRepository_Update_Call) Run(run func(ctx context.Context, r resource.Resource, hooks ...resource.MutationHook)) *ResourceRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]resource.MutationHook, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(resource.MutationHook)
			}
		}
		run(args[0].(context.Context), args[1].(resource.Resource), variadicArgs...)
	})
	return _c
}

func (_c *ResourceRepository_Update_Call) Return(_a0 error) *ResourceRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}