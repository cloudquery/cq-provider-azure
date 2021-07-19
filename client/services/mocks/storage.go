// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-azure/client/services (interfaces: StorageAccountClient,StorageContainerClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	storage "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
	gomock "github.com/golang/mock/gomock"
)

// MockStorageAccountClient is a mock of StorageAccountClient interface.
type MockStorageAccountClient struct {
	ctrl     *gomock.Controller
	recorder *MockStorageAccountClientMockRecorder
}

// MockStorageAccountClientMockRecorder is the mock recorder for MockStorageAccountClient.
type MockStorageAccountClientMockRecorder struct {
	mock *MockStorageAccountClient
}

// NewMockStorageAccountClient creates a new mock instance.
func NewMockStorageAccountClient(ctrl *gomock.Controller) *MockStorageAccountClient {
	mock := &MockStorageAccountClient{ctrl: ctrl}
	mock.recorder = &MockStorageAccountClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageAccountClient) EXPECT() *MockStorageAccountClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockStorageAccountClient) List(arg0 context.Context) (storage.AccountListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(storage.AccountListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockStorageAccountClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStorageAccountClient)(nil).List), arg0)
}

// MockStorageContainerClient is a mock of StorageContainerClient interface.
type MockStorageContainerClient struct {
	ctrl     *gomock.Controller
	recorder *MockStorageContainerClientMockRecorder
}

// MockStorageContainerClientMockRecorder is the mock recorder for MockStorageContainerClient.
type MockStorageContainerClientMockRecorder struct {
	mock *MockStorageContainerClient
}

// NewMockStorageContainerClient creates a new mock instance.
func NewMockStorageContainerClient(ctrl *gomock.Controller) *MockStorageContainerClient {
	mock := &MockStorageContainerClient{ctrl: ctrl}
	mock.recorder = &MockStorageContainerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageContainerClient) EXPECT() *MockStorageContainerClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockStorageContainerClient) List(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 storage.ListContainersInclude) (storage.ListContainerItemsPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(storage.ListContainerItemsPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockStorageContainerClientMockRecorder) List(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStorageContainerClient)(nil).List), arg0, arg1, arg2, arg3, arg4, arg5)
}
