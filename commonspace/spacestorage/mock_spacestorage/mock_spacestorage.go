// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/any-sync/commonspace/spacestorage (interfaces: SpaceStorage)
//
// Generated by this command:
//
//	mockgen -destination mock_spacestorage/mock_spacestorage.go github.com/anyproto/any-sync/commonspace/spacestorage SpaceStorage
//

// Package mock_spacestorage is a generated GoMock package.
package mock_spacestorage

import (
	context "context"
	reflect "reflect"

	anystore "github.com/anyproto/any-store"
	app "github.com/anyproto/any-sync/app"
	headstorage "github.com/anyproto/any-sync/commonspace/headsync/headstorage"
	statestorage "github.com/anyproto/any-sync/commonspace/headsync/statestorage"
	list "github.com/anyproto/any-sync/commonspace/object/acl/list"
	objecttree "github.com/anyproto/any-sync/commonspace/object/tree/objecttree"
	treestorage "github.com/anyproto/any-sync/commonspace/object/tree/treestorage"
	gomock "go.uber.org/mock/gomock"
)

// MockSpaceStorage is a mock of SpaceStorage interface.
type MockSpaceStorage struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceStorageMockRecorder
}

// MockSpaceStorageMockRecorder is the mock recorder for MockSpaceStorage.
type MockSpaceStorageMockRecorder struct {
	mock *MockSpaceStorage
}

// NewMockSpaceStorage creates a new mock instance.
func NewMockSpaceStorage(ctrl *gomock.Controller) *MockSpaceStorage {
	mock := &MockSpaceStorage{ctrl: ctrl}
	mock.recorder = &MockSpaceStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceStorage) EXPECT() *MockSpaceStorageMockRecorder {
	return m.recorder
}

// AclStorage mocks base method.
func (m *MockSpaceStorage) AclStorage() (list.Storage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AclStorage")
	ret0, _ := ret[0].(list.Storage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AclStorage indicates an expected call of AclStorage.
func (mr *MockSpaceStorageMockRecorder) AclStorage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AclStorage", reflect.TypeOf((*MockSpaceStorage)(nil).AclStorage))
}

// AnyStore mocks base method.
func (m *MockSpaceStorage) AnyStore() anystore.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnyStore")
	ret0, _ := ret[0].(anystore.DB)
	return ret0
}

// AnyStore indicates an expected call of AnyStore.
func (mr *MockSpaceStorageMockRecorder) AnyStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnyStore", reflect.TypeOf((*MockSpaceStorage)(nil).AnyStore))
}

// Close mocks base method.
func (m *MockSpaceStorage) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSpaceStorageMockRecorder) Close(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSpaceStorage)(nil).Close), arg0)
}

// CreateTreeStorage mocks base method.
func (m *MockSpaceStorage) CreateTreeStorage(arg0 context.Context, arg1 treestorage.TreeStorageCreatePayload) (objecttree.Storage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTreeStorage", arg0, arg1)
	ret0, _ := ret[0].(objecttree.Storage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTreeStorage indicates an expected call of CreateTreeStorage.
func (mr *MockSpaceStorageMockRecorder) CreateTreeStorage(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTreeStorage", reflect.TypeOf((*MockSpaceStorage)(nil).CreateTreeStorage), arg0, arg1)
}

// HeadStorage mocks base method.
func (m *MockSpaceStorage) HeadStorage() headstorage.HeadStorage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeadStorage")
	ret0, _ := ret[0].(headstorage.HeadStorage)
	return ret0
}

// HeadStorage indicates an expected call of HeadStorage.
func (mr *MockSpaceStorageMockRecorder) HeadStorage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadStorage", reflect.TypeOf((*MockSpaceStorage)(nil).HeadStorage))
}

// Id mocks base method.
func (m *MockSpaceStorage) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockSpaceStorageMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockSpaceStorage)(nil).Id))
}

// Init mocks base method.
func (m *MockSpaceStorage) Init(arg0 *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockSpaceStorageMockRecorder) Init(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockSpaceStorage)(nil).Init), arg0)
}

// Name mocks base method.
func (m *MockSpaceStorage) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockSpaceStorageMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockSpaceStorage)(nil).Name))
}

// Run mocks base method.
func (m *MockSpaceStorage) Run(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockSpaceStorageMockRecorder) Run(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockSpaceStorage)(nil).Run), arg0)
}

// StateStorage mocks base method.
func (m *MockSpaceStorage) StateStorage() statestorage.StateStorage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateStorage")
	ret0, _ := ret[0].(statestorage.StateStorage)
	return ret0
}

// StateStorage indicates an expected call of StateStorage.
func (mr *MockSpaceStorageMockRecorder) StateStorage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateStorage", reflect.TypeOf((*MockSpaceStorage)(nil).StateStorage))
}

// TreeStorage mocks base method.
func (m *MockSpaceStorage) TreeStorage(arg0 context.Context, arg1 string) (objecttree.Storage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TreeStorage", arg0, arg1)
	ret0, _ := ret[0].(objecttree.Storage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TreeStorage indicates an expected call of TreeStorage.
func (mr *MockSpaceStorageMockRecorder) TreeStorage(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TreeStorage", reflect.TypeOf((*MockSpaceStorage)(nil).TreeStorage), arg0, arg1)
}
