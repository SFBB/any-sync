// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/any-sync/commonspace (interfaces: Space)
//
// Generated by this command:
//
//	mockgen -destination mock_commonspace/mock_commonspace.go github.com/anyproto/any-sync/commonspace Space
//

// Package mock_commonspace is a generated GoMock package.
package mock_commonspace

import (
	context "context"
	reflect "reflect"
	time "time"

	commonspace "github.com/anyproto/any-sync/commonspace"
	aclclient "github.com/anyproto/any-sync/commonspace/acl/aclclient"
	headsync "github.com/anyproto/any-sync/commonspace/headsync"
	syncacl "github.com/anyproto/any-sync/commonspace/object/acl/syncacl"
	treesyncer "github.com/anyproto/any-sync/commonspace/object/treesyncer"
	objecttreebuilder "github.com/anyproto/any-sync/commonspace/objecttreebuilder"
	spacestorage "github.com/anyproto/any-sync/commonspace/spacestorage"
	spacesyncproto "github.com/anyproto/any-sync/commonspace/spacesyncproto"
	objectmessages "github.com/anyproto/any-sync/commonspace/sync/objectsync/objectmessages"
	syncstatus "github.com/anyproto/any-sync/commonspace/syncstatus"
	peer "github.com/anyproto/any-sync/net/peer"
	gomock "go.uber.org/mock/gomock"
	drpc "storj.io/drpc"
)

// MockSpace is a mock of Space interface.
type MockSpace struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceMockRecorder
}

// MockSpaceMockRecorder is the mock recorder for MockSpace.
type MockSpaceMockRecorder struct {
	mock *MockSpace
}

// NewMockSpace creates a new mock instance.
func NewMockSpace(ctrl *gomock.Controller) *MockSpace {
	mock := &MockSpace{ctrl: ctrl}
	mock.recorder = &MockSpaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpace) EXPECT() *MockSpaceMockRecorder {
	return m.recorder
}

// Acl mocks base method.
func (m *MockSpace) Acl() syncacl.SyncAcl {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Acl")
	ret0, _ := ret[0].(syncacl.SyncAcl)
	return ret0
}

// Acl indicates an expected call of Acl.
func (mr *MockSpaceMockRecorder) Acl() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Acl", reflect.TypeOf((*MockSpace)(nil).Acl))
}

// AclClient mocks base method.
func (m *MockSpace) AclClient() aclclient.AclSpaceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AclClient")
	ret0, _ := ret[0].(aclclient.AclSpaceClient)
	return ret0
}

// AclClient indicates an expected call of AclClient.
func (mr *MockSpaceMockRecorder) AclClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AclClient", reflect.TypeOf((*MockSpace)(nil).AclClient))
}

// Close mocks base method.
func (m *MockSpace) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSpaceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSpace)(nil).Close))
}

// DebugAllHeads mocks base method.
func (m *MockSpace) DebugAllHeads() []headsync.TreeHeads {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DebugAllHeads")
	ret0, _ := ret[0].([]headsync.TreeHeads)
	return ret0
}

// DebugAllHeads indicates an expected call of DebugAllHeads.
func (mr *MockSpaceMockRecorder) DebugAllHeads() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DebugAllHeads", reflect.TypeOf((*MockSpace)(nil).DebugAllHeads))
}

// DeleteTree mocks base method.
func (m *MockSpace) DeleteTree(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTree", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTree indicates an expected call of DeleteTree.
func (mr *MockSpaceMockRecorder) DeleteTree(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTree", reflect.TypeOf((*MockSpace)(nil).DeleteTree), arg0, arg1)
}

// Description mocks base method.
func (m *MockSpace) Description(arg0 context.Context) (commonspace.SpaceDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Description", arg0)
	ret0, _ := ret[0].(commonspace.SpaceDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Description indicates an expected call of Description.
func (mr *MockSpaceMockRecorder) Description(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Description", reflect.TypeOf((*MockSpace)(nil).Description), arg0)
}

// GetNodePeers mocks base method.
func (m *MockSpace) GetNodePeers(arg0 context.Context) ([]peer.Peer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodePeers", arg0)
	ret0, _ := ret[0].([]peer.Peer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodePeers indicates an expected call of GetNodePeers.
func (mr *MockSpaceMockRecorder) GetNodePeers(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodePeers", reflect.TypeOf((*MockSpace)(nil).GetNodePeers), arg0)
}

// HandleMessage mocks base method.
func (m *MockSpace) HandleMessage(arg0 context.Context, arg1 *objectmessages.HeadUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleMessage indicates an expected call of HandleMessage.
func (mr *MockSpaceMockRecorder) HandleMessage(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMessage", reflect.TypeOf((*MockSpace)(nil).HandleMessage), arg0, arg1)
}

// HandleRangeRequest mocks base method.
func (m *MockSpace) HandleRangeRequest(arg0 context.Context, arg1 *spacesyncproto.HeadSyncRequest) (*spacesyncproto.HeadSyncResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleRangeRequest", arg0, arg1)
	ret0, _ := ret[0].(*spacesyncproto.HeadSyncResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleRangeRequest indicates an expected call of HandleRangeRequest.
func (mr *MockSpaceMockRecorder) HandleRangeRequest(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleRangeRequest", reflect.TypeOf((*MockSpace)(nil).HandleRangeRequest), arg0, arg1)
}

// HandleStreamSyncRequest mocks base method.
func (m *MockSpace) HandleStreamSyncRequest(arg0 context.Context, arg1 *spacesyncproto.ObjectSyncMessage, arg2 drpc.Stream) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleStreamSyncRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleStreamSyncRequest indicates an expected call of HandleStreamSyncRequest.
func (mr *MockSpaceMockRecorder) HandleStreamSyncRequest(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleStreamSyncRequest", reflect.TypeOf((*MockSpace)(nil).HandleStreamSyncRequest), arg0, arg1, arg2)
}

// Id mocks base method.
func (m *MockSpace) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockSpaceMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockSpace)(nil).Id))
}

// Init mocks base method.
func (m *MockSpace) Init(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockSpaceMockRecorder) Init(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockSpace)(nil).Init), arg0)
}

// Storage mocks base method.
func (m *MockSpace) Storage() spacestorage.SpaceStorage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Storage")
	ret0, _ := ret[0].(spacestorage.SpaceStorage)
	return ret0
}

// Storage indicates an expected call of Storage.
func (mr *MockSpaceMockRecorder) Storage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Storage", reflect.TypeOf((*MockSpace)(nil).Storage))
}

// StoredIds mocks base method.
func (m *MockSpace) StoredIds() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoredIds")
	ret0, _ := ret[0].([]string)
	return ret0
}

// StoredIds indicates an expected call of StoredIds.
func (mr *MockSpaceMockRecorder) StoredIds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoredIds", reflect.TypeOf((*MockSpace)(nil).StoredIds))
}

// SyncStatus mocks base method.
func (m *MockSpace) SyncStatus() syncstatus.StatusUpdater {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStatus")
	ret0, _ := ret[0].(syncstatus.StatusUpdater)
	return ret0
}

// SyncStatus indicates an expected call of SyncStatus.
func (mr *MockSpaceMockRecorder) SyncStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStatus", reflect.TypeOf((*MockSpace)(nil).SyncStatus))
}

// TreeBuilder mocks base method.
func (m *MockSpace) TreeBuilder() objecttreebuilder.TreeBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TreeBuilder")
	ret0, _ := ret[0].(objecttreebuilder.TreeBuilder)
	return ret0
}

// TreeBuilder indicates an expected call of TreeBuilder.
func (mr *MockSpaceMockRecorder) TreeBuilder() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TreeBuilder", reflect.TypeOf((*MockSpace)(nil).TreeBuilder))
}

// TreeSyncer mocks base method.
func (m *MockSpace) TreeSyncer() treesyncer.TreeSyncer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TreeSyncer")
	ret0, _ := ret[0].(treesyncer.TreeSyncer)
	return ret0
}

// TreeSyncer indicates an expected call of TreeSyncer.
func (mr *MockSpaceMockRecorder) TreeSyncer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TreeSyncer", reflect.TypeOf((*MockSpace)(nil).TreeSyncer))
}

// TryClose mocks base method.
func (m *MockSpace) TryClose(arg0 time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TryClose", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TryClose indicates an expected call of TryClose.
func (mr *MockSpaceMockRecorder) TryClose(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TryClose", reflect.TypeOf((*MockSpace)(nil).TryClose), arg0)
}
