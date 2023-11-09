// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MockClientInterface is an autogenerated mock type for the ClientInterface type
type MockClientInterface struct {
	mock.Mock
}

type MockClientInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClientInterface) EXPECT() *MockClientInterface_Expecter {
	return &MockClientInterface_Expecter{mock: &_m.Mock}
}

// NewMockClientInterface creates a new instance of MockClientInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClientInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClientInterface {
	mock := &MockClientInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
