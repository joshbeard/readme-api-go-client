// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	readme "github.com/liveoaklabs/readme-api-go-client/readme"
	mock "github.com/stretchr/testify/mock"
)

// MockCategoryService is an autogenerated mock type for the CategoryService type
type MockCategoryService struct {
	mock.Mock
}

type MockCategoryService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCategoryService) EXPECT() *MockCategoryService_Expecter {
	return &MockCategoryService_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: response, params, options
func (_m *MockCategoryService) Create(response interface{}, params readme.CategoryParams, options ...readme.RequestOptions) (*readme.APIResponse, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, response, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *readme.APIResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, readme.CategoryParams, ...readme.RequestOptions) (*readme.APIResponse, error)); ok {
		return rf(response, params, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, readme.CategoryParams, ...readme.RequestOptions) *readme.APIResponse); ok {
		r0 = rf(response, params, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*readme.APIResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, readme.CategoryParams, ...readme.RequestOptions) error); ok {
		r1 = rf(response, params, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCategoryService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockCategoryService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - response interface{}
//   - params readme.CategoryParams
//   - options ...readme.RequestOptions
func (_e *MockCategoryService_Expecter) Create(response interface{}, params interface{}, options ...interface{}) *MockCategoryService_Create_Call {
	return &MockCategoryService_Create_Call{Call: _e.mock.On("Create",
		append([]interface{}{response, params}, options...)...)}
}

func (_c *MockCategoryService_Create_Call) Run(run func(response interface{}, params readme.CategoryParams, options ...readme.RequestOptions)) *MockCategoryService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]readme.RequestOptions, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(readme.RequestOptions)
			}
		}
		run(args[0].(interface{}), args[1].(readme.CategoryParams), variadicArgs...)
	})
	return _c
}

func (_c *MockCategoryService_Create_Call) Return(_a0 *readme.APIResponse, _a1 error) *MockCategoryService_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCategoryService_Create_Call) RunAndReturn(run func(interface{}, readme.CategoryParams, ...readme.RequestOptions) (*readme.APIResponse, error)) *MockCategoryService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: slug, options
func (_m *MockCategoryService) Delete(slug string, options ...readme.RequestOptions) (bool, *readme.APIResponse, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, slug)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	var r1 *readme.APIResponse
	var r2 error
	if rf, ok := ret.Get(0).(func(string, ...readme.RequestOptions) (bool, *readme.APIResponse, error)); ok {
		return rf(slug, options...)
	}
	if rf, ok := ret.Get(0).(func(string, ...readme.RequestOptions) bool); ok {
		r0 = rf(slug, options...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, ...readme.RequestOptions) *readme.APIResponse); ok {
		r1 = rf(slug, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*readme.APIResponse)
		}
	}

	if rf, ok := ret.Get(2).(func(string, ...readme.RequestOptions) error); ok {
		r2 = rf(slug, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockCategoryService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockCategoryService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - slug string
//   - options ...readme.RequestOptions
func (_e *MockCategoryService_Expecter) Delete(slug interface{}, options ...interface{}) *MockCategoryService_Delete_Call {
	return &MockCategoryService_Delete_Call{Call: _e.mock.On("Delete",
		append([]interface{}{slug}, options...)...)}
}

func (_c *MockCategoryService_Delete_Call) Run(run func(slug string, options ...readme.RequestOptions)) *MockCategoryService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]readme.RequestOptions, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(readme.RequestOptions)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockCategoryService_Delete_Call) Return(_a0 bool, _a1 *readme.APIResponse, _a2 error) *MockCategoryService_Delete_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockCategoryService_Delete_Call) RunAndReturn(run func(string, ...readme.RequestOptions) (bool, *readme.APIResponse, error)) *MockCategoryService_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: category, options
func (_m *MockCategoryService) Get(category string, options ...readme.RequestOptions) (readme.Category, *readme.APIResponse, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, category)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 readme.Category
	var r1 *readme.APIResponse
	var r2 error
	if rf, ok := ret.Get(0).(func(string, ...readme.RequestOptions) (readme.Category, *readme.APIResponse, error)); ok {
		return rf(category, options...)
	}
	if rf, ok := ret.Get(0).(func(string, ...readme.RequestOptions) readme.Category); ok {
		r0 = rf(category, options...)
	} else {
		r0 = ret.Get(0).(readme.Category)
	}

	if rf, ok := ret.Get(1).(func(string, ...readme.RequestOptions) *readme.APIResponse); ok {
		r1 = rf(category, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*readme.APIResponse)
		}
	}

	if rf, ok := ret.Get(2).(func(string, ...readme.RequestOptions) error); ok {
		r2 = rf(category, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockCategoryService_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockCategoryService_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - category string
//   - options ...readme.RequestOptions
func (_e *MockCategoryService_Expecter) Get(category interface{}, options ...interface{}) *MockCategoryService_Get_Call {
	return &MockCategoryService_Get_Call{Call: _e.mock.On("Get",
		append([]interface{}{category}, options...)...)}
}

func (_c *MockCategoryService_Get_Call) Run(run func(category string, options ...readme.RequestOptions)) *MockCategoryService_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]readme.RequestOptions, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(readme.RequestOptions)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockCategoryService_Get_Call) Return(_a0 readme.Category, _a1 *readme.APIResponse, _a2 error) *MockCategoryService_Get_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockCategoryService_Get_Call) RunAndReturn(run func(string, ...readme.RequestOptions) (readme.Category, *readme.APIResponse, error)) *MockCategoryService_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields: options
func (_m *MockCategoryService) GetAll(options ...readme.RequestOptions) ([]readme.Category, *readme.APIResponse, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []readme.Category
	var r1 *readme.APIResponse
	var r2 error
	if rf, ok := ret.Get(0).(func(...readme.RequestOptions) ([]readme.Category, *readme.APIResponse, error)); ok {
		return rf(options...)
	}
	if rf, ok := ret.Get(0).(func(...readme.RequestOptions) []readme.Category); ok {
		r0 = rf(options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]readme.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(...readme.RequestOptions) *readme.APIResponse); ok {
		r1 = rf(options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*readme.APIResponse)
		}
	}

	if rf, ok := ret.Get(2).(func(...readme.RequestOptions) error); ok {
		r2 = rf(options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockCategoryService_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type MockCategoryService_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//   - options ...readme.RequestOptions
func (_e *MockCategoryService_Expecter) GetAll(options ...interface{}) *MockCategoryService_GetAll_Call {
	return &MockCategoryService_GetAll_Call{Call: _e.mock.On("GetAll",
		append([]interface{}{}, options...)...)}
}

func (_c *MockCategoryService_GetAll_Call) Run(run func(options ...readme.RequestOptions)) *MockCategoryService_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]readme.RequestOptions, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(readme.RequestOptions)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockCategoryService_GetAll_Call) Return(_a0 []readme.Category, _a1 *readme.APIResponse, _a2 error) *MockCategoryService_GetAll_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockCategoryService_GetAll_Call) RunAndReturn(run func(...readme.RequestOptions) ([]readme.Category, *readme.APIResponse, error)) *MockCategoryService_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetDocs provides a mock function with given fields: slug, options
func (_m *MockCategoryService) GetDocs(slug string, options ...readme.RequestOptions) ([]readme.CategoryDocs, *readme.APIResponse, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, slug)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []readme.CategoryDocs
	var r1 *readme.APIResponse
	var r2 error
	if rf, ok := ret.Get(0).(func(string, ...readme.RequestOptions) ([]readme.CategoryDocs, *readme.APIResponse, error)); ok {
		return rf(slug, options...)
	}
	if rf, ok := ret.Get(0).(func(string, ...readme.RequestOptions) []readme.CategoryDocs); ok {
		r0 = rf(slug, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]readme.CategoryDocs)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...readme.RequestOptions) *readme.APIResponse); ok {
		r1 = rf(slug, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*readme.APIResponse)
		}
	}

	if rf, ok := ret.Get(2).(func(string, ...readme.RequestOptions) error); ok {
		r2 = rf(slug, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockCategoryService_GetDocs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDocs'
type MockCategoryService_GetDocs_Call struct {
	*mock.Call
}

// GetDocs is a helper method to define mock.On call
//   - slug string
//   - options ...readme.RequestOptions
func (_e *MockCategoryService_Expecter) GetDocs(slug interface{}, options ...interface{}) *MockCategoryService_GetDocs_Call {
	return &MockCategoryService_GetDocs_Call{Call: _e.mock.On("GetDocs",
		append([]interface{}{slug}, options...)...)}
}

func (_c *MockCategoryService_GetDocs_Call) Run(run func(slug string, options ...readme.RequestOptions)) *MockCategoryService_GetDocs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]readme.RequestOptions, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(readme.RequestOptions)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockCategoryService_GetDocs_Call) Return(_a0 []readme.CategoryDocs, _a1 *readme.APIResponse, _a2 error) *MockCategoryService_GetDocs_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockCategoryService_GetDocs_Call) RunAndReturn(run func(string, ...readme.RequestOptions) ([]readme.CategoryDocs, *readme.APIResponse, error)) *MockCategoryService_GetDocs_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: slug, params, options
func (_m *MockCategoryService) Update(slug string, params readme.CategoryParams, options ...readme.RequestOptions) (readme.Category, *readme.APIResponse, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, slug, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 readme.Category
	var r1 *readme.APIResponse
	var r2 error
	if rf, ok := ret.Get(0).(func(string, readme.CategoryParams, ...readme.RequestOptions) (readme.Category, *readme.APIResponse, error)); ok {
		return rf(slug, params, options...)
	}
	if rf, ok := ret.Get(0).(func(string, readme.CategoryParams, ...readme.RequestOptions) readme.Category); ok {
		r0 = rf(slug, params, options...)
	} else {
		r0 = ret.Get(0).(readme.Category)
	}

	if rf, ok := ret.Get(1).(func(string, readme.CategoryParams, ...readme.RequestOptions) *readme.APIResponse); ok {
		r1 = rf(slug, params, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*readme.APIResponse)
		}
	}

	if rf, ok := ret.Get(2).(func(string, readme.CategoryParams, ...readme.RequestOptions) error); ok {
		r2 = rf(slug, params, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockCategoryService_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockCategoryService_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - slug string
//   - params readme.CategoryParams
//   - options ...readme.RequestOptions
func (_e *MockCategoryService_Expecter) Update(slug interface{}, params interface{}, options ...interface{}) *MockCategoryService_Update_Call {
	return &MockCategoryService_Update_Call{Call: _e.mock.On("Update",
		append([]interface{}{slug, params}, options...)...)}
}

func (_c *MockCategoryService_Update_Call) Run(run func(slug string, params readme.CategoryParams, options ...readme.RequestOptions)) *MockCategoryService_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]readme.RequestOptions, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(readme.RequestOptions)
			}
		}
		run(args[0].(string), args[1].(readme.CategoryParams), variadicArgs...)
	})
	return _c
}

func (_c *MockCategoryService_Update_Call) Return(_a0 readme.Category, _a1 *readme.APIResponse, _a2 error) *MockCategoryService_Update_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockCategoryService_Update_Call) RunAndReturn(run func(string, readme.CategoryParams, ...readme.RequestOptions) (readme.Category, *readme.APIResponse, error)) *MockCategoryService_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCategoryService creates a new instance of MockCategoryService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCategoryService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCategoryService {
	mock := &MockCategoryService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
