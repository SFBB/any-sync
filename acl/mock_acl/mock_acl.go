// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/any-sync/acl (interfaces: AclService)
//
// Generated by this command:
//
//	mockgen -destination mock_acl/mock_acl.go github.com/anyproto/any-sync/acl AclService
//

// Package mock_acl is a generated GoMock package.
package mock_acl

import (
	context "context"
	reflect "reflect"

	acl "github.com/anyproto/any-sync/acl"
	app "github.com/anyproto/any-sync/app"
	list "github.com/anyproto/any-sync/commonspace/object/acl/list"
	consensusproto "github.com/anyproto/any-sync/consensus/consensusproto"
	crypto "github.com/anyproto/any-sync/util/crypto"
	gomock "go.uber.org/mock/gomock"
)

// MockAclService is a mock of AclService interface.
type MockAclService struct {
	ctrl     *gomock.Controller
	recorder *MockAclServiceMockRecorder
}

// MockAclServiceMockRecorder is the mock recorder for MockAclService.
type MockAclServiceMockRecorder struct {
	mock *MockAclService
}

// NewMockAclService creates a new mock instance.
func NewMockAclService(ctrl *gomock.Controller) *MockAclService {
	mock := &MockAclService{ctrl: ctrl}
	mock.recorder = &MockAclServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAclService) EXPECT() *MockAclServiceMockRecorder {
	return m.recorder
}

// AddRecord mocks base method.
func (m *MockAclService) AddRecord(arg0 context.Context, arg1 string, arg2 *consensusproto.RawRecord, arg3 acl.Limits) (*consensusproto.RawRecordWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRecord", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*consensusproto.RawRecordWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddRecord indicates an expected call of AddRecord.
func (mr *MockAclServiceMockRecorder) AddRecord(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRecord", reflect.TypeOf((*MockAclService)(nil).AddRecord), arg0, arg1, arg2, arg3)
}

// Close mocks base method.
func (m *MockAclService) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockAclServiceMockRecorder) Close(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAclService)(nil).Close), arg0)
}

// HasRecord mocks base method.
func (m *MockAclService) HasRecord(arg0 context.Context, arg1, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasRecord", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasRecord indicates an expected call of HasRecord.
func (mr *MockAclServiceMockRecorder) HasRecord(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasRecord", reflect.TypeOf((*MockAclService)(nil).HasRecord), arg0, arg1, arg2)
}

// Init mocks base method.
func (m *MockAclService) Init(arg0 *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockAclServiceMockRecorder) Init(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockAclService)(nil).Init), arg0)
}

// Name mocks base method.
func (m *MockAclService) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockAclServiceMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAclService)(nil).Name))
}

// OwnerPubKey mocks base method.
func (m *MockAclService) OwnerPubKey(arg0 context.Context, arg1 string) (crypto.PubKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OwnerPubKey", arg0, arg1)
	ret0, _ := ret[0].(crypto.PubKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OwnerPubKey indicates an expected call of OwnerPubKey.
func (mr *MockAclServiceMockRecorder) OwnerPubKey(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OwnerPubKey", reflect.TypeOf((*MockAclService)(nil).OwnerPubKey), arg0, arg1)
}

// Permissions mocks base method.
func (m *MockAclService) Permissions(arg0 context.Context, arg1 crypto.PubKey, arg2 string) (list.AclPermissions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Permissions", arg0, arg1, arg2)
	ret0, _ := ret[0].(list.AclPermissions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Permissions indicates an expected call of Permissions.
func (mr *MockAclServiceMockRecorder) Permissions(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Permissions", reflect.TypeOf((*MockAclService)(nil).Permissions), arg0, arg1, arg2)
}

// ReadState mocks base method.
func (m *MockAclService) ReadState(arg0 context.Context, arg1 string, arg2 func(*list.AclState) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadState", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadState indicates an expected call of ReadState.
func (mr *MockAclServiceMockRecorder) ReadState(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadState", reflect.TypeOf((*MockAclService)(nil).ReadState), arg0, arg1, arg2)
}

// RecordsAfter mocks base method.
func (m *MockAclService) RecordsAfter(arg0 context.Context, arg1, arg2 string) ([]*consensusproto.RawRecordWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordsAfter", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*consensusproto.RawRecordWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecordsAfter indicates an expected call of RecordsAfter.
func (mr *MockAclServiceMockRecorder) RecordsAfter(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordsAfter", reflect.TypeOf((*MockAclService)(nil).RecordsAfter), arg0, arg1, arg2)
}

// Run mocks base method.
func (m *MockAclService) Run(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockAclServiceMockRecorder) Run(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockAclService)(nil).Run), arg0)
}
