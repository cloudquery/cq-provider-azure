// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-azure/client/services (interfaces: ResClient,GroupsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	resources "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"
	gomock "github.com/golang/mock/gomock"
)

// MockResClient is a mock of ResClient interface.
type MockResClient struct {
	ctrl     *gomock.Controller
	recorder *MockResClientMockRecorder
}

// MockResClientMockRecorder is the mock recorder for MockResClient.
type MockResClientMockRecorder struct {
	mock *MockResClient
}

// NewMockResClient creates a new mock instance.
func NewMockResClient(ctrl *gomock.Controller) *MockResClient {
	mock := &MockResClient{ctrl: ctrl}
	mock.recorder = &MockResClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResClient) EXPECT() *MockResClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockResClient) List(arg0 context.Context, arg1, arg2 string, arg3 *int32) (resources.ListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(resources.ListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockResClientMockRecorder) List(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockResClient)(nil).List), arg0, arg1, arg2, arg3)
}

// MockGroupsClient is a mock of GroupsClient interface.
type MockGroupsClient struct {
	ctrl     *gomock.Controller
	recorder *MockGroupsClientMockRecorder
}

// MockGroupsClientMockRecorder is the mock recorder for MockGroupsClient.
type MockGroupsClientMockRecorder struct {
	mock *MockGroupsClient
}

// NewMockGroupsClient creates a new mock instance.
func NewMockGroupsClient(ctrl *gomock.Controller) *MockGroupsClient {
	mock := &MockGroupsClient{ctrl: ctrl}
	mock.recorder = &MockGroupsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupsClient) EXPECT() *MockGroupsClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockGroupsClient) List(arg0 context.Context, arg1 string, arg2 *int32) (resources.GroupListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(resources.GroupListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockGroupsClientMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockGroupsClient)(nil).List), arg0, arg1, arg2)
}
