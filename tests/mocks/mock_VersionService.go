// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	readme "github.com/liveoaklabs/readme-api-go-client/readme"
	mock "github.com/stretchr/testify/mock"
)

// MockVersionService is an autogenerated mock type for the VersionService type
type MockVersionService struct {
	mock.Mock
}

type MockVersionService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockVersionService) EXPECT() *MockVersionService_Expecter {
	return &MockVersionService_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: prams
func (_m *MockVersionService) Create(prams readme.VersionParams) (readme.Version, *readme.APIResponse, error) {
	ret := _m.Called(prams)

	var r0 readme.Version
	var r1 *readme.APIResponse
	var r2 error
	if rf, ok := ret.Get(0).(func(readme.VersionParams) (readme.Version, *readme.APIResponse, error)); ok {
		return rf(prams)
	}
	if rf, ok := ret.Get(0).(func(readme.VersionParams) readme.Version); ok {
		r0 = rf(prams)
	} else {
		r0 = ret.Get(0).(readme.Version)
	}

	if rf, ok := ret.Get(1).(func(readme.VersionParams) *readme.APIResponse); ok {
		r1 = rf(prams)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*readme.APIResponse)
		}
	}

	if rf, ok := ret.Get(2).(func(readme.VersionParams) error); ok {
		r2 = rf(prams)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockVersionService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockVersionService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - prams readme.VersionParams
func (_e *MockVersionService_Expecter) Create(prams interface{}) *MockVersionService_Create_Call {
	return &MockVersionService_Create_Call{Call: _e.mock.On("Create", prams)}
}

func (_c *MockVersionService_Create_Call) Run(run func(prams readme.VersionParams)) *MockVersionService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(readme.VersionParams))
	})
	return _c
}

func (_c *MockVersionService_Create_Call) Return(_a0 readme.Version, _a1 *readme.APIResponse, _a2 error) *MockVersionService_Create_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockVersionService_Create_Call) RunAndReturn(run func(readme.VersionParams) (readme.Version, *readme.APIResponse, error)) *MockVersionService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: version
func (_m *MockVersionService) Delete(version string) (bool, *readme.APIResponse, error) {
	ret := _m.Called(version)

	var r0 bool
	var r1 *readme.APIResponse
	var r2 error
	if rf, ok := ret.Get(0).(func(string) (bool, *readme.APIResponse, error)); ok {
		return rf(version)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(version)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) *readme.APIResponse); ok {
		r1 = rf(version)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*readme.APIResponse)
		}
	}

	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(version)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockVersionService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockVersionService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - version string
func (_e *MockVersionService_Expecter) Delete(version interface{}) *MockVersionService_Delete_Call {
	return &MockVersionService_Delete_Call{Call: _e.mock.On("Delete", version)}
}

func (_c *MockVersionService_Delete_Call) Run(run func(version string)) *MockVersionService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockVersionService_Delete_Call) Return(_a0 bool, _a1 *readme.APIResponse, _a2 error) *MockVersionService_Delete_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockVersionService_Delete_Call) RunAndReturn(run func(string) (bool, *readme.APIResponse, error)) *MockVersionService_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: version
func (_m *MockVersionService) Get(version string) (readme.Version, *readme.APIResponse, error) {
	ret := _m.Called(version)

	var r0 readme.Version
	var r1 *readme.APIResponse
	var r2 error
	if rf, ok := ret.Get(0).(func(string) (readme.Version, *readme.APIResponse, error)); ok {
		return rf(version)
	}
	if rf, ok := ret.Get(0).(func(string) readme.Version); ok {
		r0 = rf(version)
	} else {
		r0 = ret.Get(0).(readme.Version)
	}

	if rf, ok := ret.Get(1).(func(string) *readme.APIResponse); ok {
		r1 = rf(version)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*readme.APIResponse)
		}
	}

	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(version)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockVersionService_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockVersionService_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - version string
func (_e *MockVersionService_Expecter) Get(version interface{}) *MockVersionService_Get_Call {
	return &MockVersionService_Get_Call{Call: _e.mock.On("Get", version)}
}

func (_c *MockVersionService_Get_Call) Run(run func(version string)) *MockVersionService_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockVersionService_Get_Call) Return(_a0 readme.Version, _a1 *readme.APIResponse, _a2 error) *MockVersionService_Get_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockVersionService_Get_Call) RunAndReturn(run func(string) (readme.Version, *readme.APIResponse, error)) *MockVersionService_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields:
func (_m *MockVersionService) GetAll() ([]readme.VersionSummary, *readme.APIResponse, error) {
	ret := _m.Called()

	var r0 []readme.VersionSummary
	var r1 *readme.APIResponse
	var r2 error
	if rf, ok := ret.Get(0).(func() ([]readme.VersionSummary, *readme.APIResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []readme.VersionSummary); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]readme.VersionSummary)
		}
	}

	if rf, ok := ret.Get(1).(func() *readme.APIResponse); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*readme.APIResponse)
		}
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockVersionService_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type MockVersionService_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
func (_e *MockVersionService_Expecter) GetAll() *MockVersionService_GetAll_Call {
	return &MockVersionService_GetAll_Call{Call: _e.mock.On("GetAll")}
}

func (_c *MockVersionService_GetAll_Call) Run(run func()) *MockVersionService_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockVersionService_GetAll_Call) Return(_a0 []readme.VersionSummary, _a1 *readme.APIResponse, _a2 error) *MockVersionService_GetAll_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockVersionService_GetAll_Call) RunAndReturn(run func() ([]readme.VersionSummary, *readme.APIResponse, error)) *MockVersionService_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetVersion provides a mock function with given fields: version
func (_m *MockVersionService) GetVersion(version string) (string, error) {
	ret := _m.Called(version)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(version)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(version)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVersionService_GetVersion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVersion'
type MockVersionService_GetVersion_Call struct {
	*mock.Call
}

// GetVersion is a helper method to define mock.On call
//   - version string
func (_e *MockVersionService_Expecter) GetVersion(version interface{}) *MockVersionService_GetVersion_Call {
	return &MockVersionService_GetVersion_Call{Call: _e.mock.On("GetVersion", version)}
}

func (_c *MockVersionService_GetVersion_Call) Run(run func(version string)) *MockVersionService_GetVersion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockVersionService_GetVersion_Call) Return(_a0 string, _a1 error) *MockVersionService_GetVersion_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVersionService_GetVersion_Call) RunAndReturn(run func(string) (string, error)) *MockVersionService_GetVersion_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: version, params
func (_m *MockVersionService) Update(version string, params readme.VersionParams) (readme.Version, *readme.APIResponse, error) {
	ret := _m.Called(version, params)

	var r0 readme.Version
	var r1 *readme.APIResponse
	var r2 error
	if rf, ok := ret.Get(0).(func(string, readme.VersionParams) (readme.Version, *readme.APIResponse, error)); ok {
		return rf(version, params)
	}
	if rf, ok := ret.Get(0).(func(string, readme.VersionParams) readme.Version); ok {
		r0 = rf(version, params)
	} else {
		r0 = ret.Get(0).(readme.Version)
	}

	if rf, ok := ret.Get(1).(func(string, readme.VersionParams) *readme.APIResponse); ok {
		r1 = rf(version, params)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*readme.APIResponse)
		}
	}

	if rf, ok := ret.Get(2).(func(string, readme.VersionParams) error); ok {
		r2 = rf(version, params)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockVersionService_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockVersionService_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - version string
//   - params readme.VersionParams
func (_e *MockVersionService_Expecter) Update(version interface{}, params interface{}) *MockVersionService_Update_Call {
	return &MockVersionService_Update_Call{Call: _e.mock.On("Update", version, params)}
}

func (_c *MockVersionService_Update_Call) Run(run func(version string, params readme.VersionParams)) *MockVersionService_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(readme.VersionParams))
	})
	return _c
}

func (_c *MockVersionService_Update_Call) Return(_a0 readme.Version, _a1 *readme.APIResponse, _a2 error) *MockVersionService_Update_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockVersionService_Update_Call) RunAndReturn(run func(string, readme.VersionParams) (readme.Version, *readme.APIResponse, error)) *MockVersionService_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockVersionService creates a new instance of MockVersionService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockVersionService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockVersionService {
	mock := &MockVersionService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}