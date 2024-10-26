// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/any-sync/commonspace/object/acl/liststorage (interfaces: ListStorage)
//
// Generated by this command:
//
//	mockgen -destination mock_liststorage/mock_liststorage.go github.com/anyproto/any-sync/commonspace/object/acl/liststorage ListStorage
//

// Package mock_liststorage is a generated GoMock package.
package mock_liststorage

import (
	context "context"
	reflect "reflect"

	consensusproto "github.com/anyproto/any-sync/consensus/consensusproto"
	gomock "go.uber.org/mock/gomock"
)

// MockListStorage is a mock of ListStorage interface.
type MockListStorage struct {
	ctrl     *gomock.Controller
	recorder *MockListStorageMockRecorder
	isgomock struct{}
}

// MockListStorageMockRecorder is the mock recorder for MockListStorage.
type MockListStorageMockRecorder struct {
	mock *MockListStorage
}

// NewMockListStorage creates a new mock instance.
func NewMockListStorage(ctrl *gomock.Controller) *MockListStorage {
	mock := &MockListStorage{ctrl: ctrl}
	mock.recorder = &MockListStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListStorage) EXPECT() *MockListStorageMockRecorder {
	return m.recorder
}

// AddRawRecord mocks base method.
func (m *MockListStorage) AddRawRecord(ctx context.Context, rec *consensusproto.RawRecordWithId) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRawRecord", ctx, rec)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRawRecord indicates an expected call of AddRawRecord.
func (mr *MockListStorageMockRecorder) AddRawRecord(ctx, rec any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRawRecord", reflect.TypeOf((*MockListStorage)(nil).AddRawRecord), ctx, rec)
}

// GetRawRecord mocks base method.
func (m *MockListStorage) GetRawRecord(ctx context.Context, id string) (*consensusproto.RawRecordWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawRecord", ctx, id)
	ret0, _ := ret[0].(*consensusproto.RawRecordWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawRecord indicates an expected call of GetRawRecord.
func (mr *MockListStorageMockRecorder) GetRawRecord(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawRecord", reflect.TypeOf((*MockListStorage)(nil).GetRawRecord), ctx, id)
}

// Head mocks base method.
func (m *MockListStorage) Head() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Head")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Head indicates an expected call of Head.
func (mr *MockListStorageMockRecorder) Head() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Head", reflect.TypeOf((*MockListStorage)(nil).Head))
}

// Id mocks base method.
func (m *MockListStorage) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockListStorageMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockListStorage)(nil).Id))
}

// Root mocks base method.
func (m *MockListStorage) Root() (*consensusproto.RawRecordWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Root")
	ret0, _ := ret[0].(*consensusproto.RawRecordWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Root indicates an expected call of Root.
func (mr *MockListStorageMockRecorder) Root() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Root", reflect.TypeOf((*MockListStorage)(nil).Root))
}

// SetHead mocks base method.
func (m *MockListStorage) SetHead(headId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHead", headId)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHead indicates an expected call of SetHead.
func (mr *MockListStorageMockRecorder) SetHead(headId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHead", reflect.TypeOf((*MockListStorage)(nil).SetHead), headId)
}
