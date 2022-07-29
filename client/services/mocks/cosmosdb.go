// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-azure/client/services (interfaces: CosmosDBAccountClient,CosmosDBSQLClient,CosmosDBMongoDBClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	documentdb "github.com/Azure/azure-sdk-for-go/services/preview/cosmos-db/mgmt/2021-11-15-preview/documentdb"
	gomock "github.com/golang/mock/gomock"
)

// MockCosmosDBAccountClient is a mock of CosmosDBAccountClient interface.
type MockCosmosDBAccountClient struct {
	ctrl     *gomock.Controller
	recorder *MockCosmosDBAccountClientMockRecorder
}

// MockCosmosDBAccountClientMockRecorder is the mock recorder for MockCosmosDBAccountClient.
type MockCosmosDBAccountClientMockRecorder struct {
	mock *MockCosmosDBAccountClient
}

// NewMockCosmosDBAccountClient creates a new mock instance.
func NewMockCosmosDBAccountClient(ctrl *gomock.Controller) *MockCosmosDBAccountClient {
	mock := &MockCosmosDBAccountClient{ctrl: ctrl}
	mock.recorder = &MockCosmosDBAccountClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCosmosDBAccountClient) EXPECT() *MockCosmosDBAccountClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockCosmosDBAccountClient) List(arg0 context.Context) (documentdb.DatabaseAccountsListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(documentdb.DatabaseAccountsListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockCosmosDBAccountClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCosmosDBAccountClient)(nil).List), arg0)
}

// MockCosmosDBSQLClient is a mock of CosmosDBSQLClient interface.
type MockCosmosDBSQLClient struct {
	ctrl     *gomock.Controller
	recorder *MockCosmosDBSQLClientMockRecorder
}

// MockCosmosDBSQLClientMockRecorder is the mock recorder for MockCosmosDBSQLClient.
type MockCosmosDBSQLClientMockRecorder struct {
	mock *MockCosmosDBSQLClient
}

// NewMockCosmosDBSQLClient creates a new mock instance.
func NewMockCosmosDBSQLClient(ctrl *gomock.Controller) *MockCosmosDBSQLClient {
	mock := &MockCosmosDBSQLClient{ctrl: ctrl}
	mock.recorder = &MockCosmosDBSQLClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCosmosDBSQLClient) EXPECT() *MockCosmosDBSQLClientMockRecorder {
	return m.recorder
}

// ListSQLDatabases mocks base method.
func (m *MockCosmosDBSQLClient) ListSQLDatabases(arg0 context.Context, arg1, arg2 string) (documentdb.SQLDatabaseListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSQLDatabases", arg0, arg1, arg2)
	ret0, _ := ret[0].(documentdb.SQLDatabaseListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSQLDatabases indicates an expected call of ListSQLDatabases.
func (mr *MockCosmosDBSQLClientMockRecorder) ListSQLDatabases(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSQLDatabases", reflect.TypeOf((*MockCosmosDBSQLClient)(nil).ListSQLDatabases), arg0, arg1, arg2)
}

// MockCosmosDBMongoDBClient is a mock of CosmosDBMongoDBClient interface.
type MockCosmosDBMongoDBClient struct {
	ctrl     *gomock.Controller
	recorder *MockCosmosDBMongoDBClientMockRecorder
}

// MockCosmosDBMongoDBClientMockRecorder is the mock recorder for MockCosmosDBMongoDBClient.
type MockCosmosDBMongoDBClientMockRecorder struct {
	mock *MockCosmosDBMongoDBClient
}

// NewMockCosmosDBMongoDBClient creates a new mock instance.
func NewMockCosmosDBMongoDBClient(ctrl *gomock.Controller) *MockCosmosDBMongoDBClient {
	mock := &MockCosmosDBMongoDBClient{ctrl: ctrl}
	mock.recorder = &MockCosmosDBMongoDBClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCosmosDBMongoDBClient) EXPECT() *MockCosmosDBMongoDBClientMockRecorder {
	return m.recorder
}

// ListMongoDBDatabases mocks base method.
func (m *MockCosmosDBMongoDBClient) ListMongoDBDatabases(arg0 context.Context, arg1, arg2 string) (documentdb.MongoDBDatabaseListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMongoDBDatabases", arg0, arg1, arg2)
	ret0, _ := ret[0].(documentdb.MongoDBDatabaseListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMongoDBDatabases indicates an expected call of ListMongoDBDatabases.
func (mr *MockCosmosDBMongoDBClientMockRecorder) ListMongoDBDatabases(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMongoDBDatabases", reflect.TypeOf((*MockCosmosDBMongoDBClient)(nil).ListMongoDBDatabases), arg0, arg1, arg2)
}
