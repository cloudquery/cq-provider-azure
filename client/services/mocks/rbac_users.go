// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-azure/client/services (interfaces: RBACUsersClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	graphrbac "github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	gomock "github.com/golang/mock/gomock"
)

// MockRBACUsersClient is a mock of RBACUsersClient interface.
type MockRBACUsersClient struct {
	ctrl     *gomock.Controller
	recorder *MockRBACUsersClientMockRecorder
}

// MockRBACUsersClientMockRecorder is the mock recorder for MockRBACUsersClient.
type MockRBACUsersClientMockRecorder struct {
	mock *MockRBACUsersClient
}

// NewMockRBACUsersClient creates a new mock instance.
func NewMockRBACUsersClient(ctrl *gomock.Controller) *MockRBACUsersClient {
	mock := &MockRBACUsersClient{ctrl: ctrl}
	mock.recorder = &MockRBACUsersClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRBACUsersClient) EXPECT() *MockRBACUsersClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockRBACUsersClient) List(arg0 context.Context, arg1, arg2 string) (graphrbac.UserListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(graphrbac.UserListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRBACUsersClientMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRBACUsersClient)(nil).List), arg0, arg1, arg2)
}
