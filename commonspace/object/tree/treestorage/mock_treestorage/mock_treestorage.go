// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/any-sync/commonspace/object/tree/treestorage (interfaces: TreeStorage)
//
// Generated by this command:
//
//	mockgen -destination mock_treestorage/mock_treestorage.go github.com/anyproto/any-sync/commonspace/object/tree/treestorage TreeStorage
//

// Package mock_treestorage is a generated GoMock package.
package mock_treestorage

import (
	context "context"
	reflect "reflect"

	treechangeproto "github.com/anyproto/any-sync/commonspace/object/tree/treechangeproto"
	gomock "go.uber.org/mock/gomock"
)

// MockTreeStorage is a mock of TreeStorage interface.
type MockTreeStorage struct {
	ctrl     *gomock.Controller
	recorder *MockTreeStorageMockRecorder
	isgomock struct{}
}

// MockTreeStorageMockRecorder is the mock recorder for MockTreeStorage.
type MockTreeStorageMockRecorder struct {
	mock *MockTreeStorage
}

// NewMockTreeStorage creates a new mock instance.
func NewMockTreeStorage(ctrl *gomock.Controller) *MockTreeStorage {
	mock := &MockTreeStorage{ctrl: ctrl}
	mock.recorder = &MockTreeStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTreeStorage) EXPECT() *MockTreeStorageMockRecorder {
	return m.recorder
}

// AddRawChange mocks base method.
func (m *MockTreeStorage) AddRawChange(change *treechangeproto.RawTreeChangeWithId) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRawChange", change)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRawChange indicates an expected call of AddRawChange.
func (mr *MockTreeStorageMockRecorder) AddRawChange(change any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRawChange", reflect.TypeOf((*MockTreeStorage)(nil).AddRawChange), change)
}

// AddRawChangesSetHeads mocks base method.
func (m *MockTreeStorage) AddRawChangesSetHeads(changes []*treechangeproto.RawTreeChangeWithId, heads []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRawChangesSetHeads", changes, heads)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRawChangesSetHeads indicates an expected call of AddRawChangesSetHeads.
func (mr *MockTreeStorageMockRecorder) AddRawChangesSetHeads(changes, heads any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRawChangesSetHeads", reflect.TypeOf((*MockTreeStorage)(nil).AddRawChangesSetHeads), changes, heads)
}

// Delete mocks base method.
func (m *MockTreeStorage) Delete() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTreeStorageMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTreeStorage)(nil).Delete))
}

// GetAllChangeIds mocks base method.
func (m *MockTreeStorage) GetAllChangeIds() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllChangeIds")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllChangeIds indicates an expected call of GetAllChangeIds.
func (mr *MockTreeStorageMockRecorder) GetAllChangeIds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllChangeIds", reflect.TypeOf((*MockTreeStorage)(nil).GetAllChangeIds))
}

// GetAppendRawChange mocks base method.
func (m *MockTreeStorage) GetAppendRawChange(ctx context.Context, buf []byte, id string) (*treechangeproto.RawTreeChangeWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppendRawChange", ctx, buf, id)
	ret0, _ := ret[0].(*treechangeproto.RawTreeChangeWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppendRawChange indicates an expected call of GetAppendRawChange.
func (mr *MockTreeStorageMockRecorder) GetAppendRawChange(ctx, buf, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppendRawChange", reflect.TypeOf((*MockTreeStorage)(nil).GetAppendRawChange), ctx, buf, id)
}

// GetRawChange mocks base method.
func (m *MockTreeStorage) GetRawChange(ctx context.Context, id string) (*treechangeproto.RawTreeChangeWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawChange", ctx, id)
	ret0, _ := ret[0].(*treechangeproto.RawTreeChangeWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawChange indicates an expected call of GetRawChange.
func (mr *MockTreeStorageMockRecorder) GetRawChange(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawChange", reflect.TypeOf((*MockTreeStorage)(nil).GetRawChange), ctx, id)
}

// HasChange mocks base method.
func (m *MockTreeStorage) HasChange(ctx context.Context, id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasChange", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasChange indicates an expected call of HasChange.
func (mr *MockTreeStorageMockRecorder) HasChange(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasChange", reflect.TypeOf((*MockTreeStorage)(nil).HasChange), ctx, id)
}

// Heads mocks base method.
func (m *MockTreeStorage) Heads() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Heads")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Heads indicates an expected call of Heads.
func (mr *MockTreeStorageMockRecorder) Heads() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Heads", reflect.TypeOf((*MockTreeStorage)(nil).Heads))
}

// Id mocks base method.
func (m *MockTreeStorage) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockTreeStorageMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockTreeStorage)(nil).Id))
}

// Root mocks base method.
func (m *MockTreeStorage) Root() (*treechangeproto.RawTreeChangeWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Root")
	ret0, _ := ret[0].(*treechangeproto.RawTreeChangeWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Root indicates an expected call of Root.
func (mr *MockTreeStorageMockRecorder) Root() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Root", reflect.TypeOf((*MockTreeStorage)(nil).Root))
}

// SetHeads mocks base method.
func (m *MockTreeStorage) SetHeads(heads []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeads", heads)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeads indicates an expected call of SetHeads.
func (mr *MockTreeStorageMockRecorder) SetHeads(heads any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeads", reflect.TypeOf((*MockTreeStorage)(nil).SetHeads), heads)
}
