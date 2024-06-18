package commonspace

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/anyproto/go-chash"
	"github.com/stretchr/testify/require"
	"storj.io/drpc"

	accountService "github.com/anyproto/any-sync/accountservice"
	"github.com/anyproto/any-sync/app"
	"github.com/anyproto/any-sync/app/ocache"
	"github.com/anyproto/any-sync/commonspace/config"
	"github.com/anyproto/any-sync/commonspace/credentialprovider"
	"github.com/anyproto/any-sync/commonspace/object/accountdata"
	"github.com/anyproto/any-sync/commonspace/object/acl/list"
	"github.com/anyproto/any-sync/commonspace/object/tree/objecttree"
	"github.com/anyproto/any-sync/commonspace/object/tree/treestorage"
	"github.com/anyproto/any-sync/commonspace/object/treemanager"
	"github.com/anyproto/any-sync/commonspace/object/treesyncer"
	"github.com/anyproto/any-sync/commonspace/objecttreebuilder"
	"github.com/anyproto/any-sync/commonspace/peermanager"
	"github.com/anyproto/any-sync/commonspace/spacestorage"
	"github.com/anyproto/any-sync/commonspace/spacesyncproto"
	"github.com/anyproto/any-sync/commonspace/sync/synctest"
	"github.com/anyproto/any-sync/commonspace/syncstatus"
	"github.com/anyproto/any-sync/consensus/consensusproto"
	"github.com/anyproto/any-sync/coordinator/coordinatorclient"
	"github.com/anyproto/any-sync/coordinator/coordinatorproto"
	"github.com/anyproto/any-sync/identityrepo/identityrepoproto"
	"github.com/anyproto/any-sync/net/peer"
	"github.com/anyproto/any-sync/net/pool"
	"github.com/anyproto/any-sync/net/rpc/rpctest"
	"github.com/anyproto/any-sync/net/streampool"
	"github.com/anyproto/any-sync/net/streampool/streamopener"
	"github.com/anyproto/any-sync/nodeconf"
	"github.com/anyproto/any-sync/testutil/accounttest"
	"github.com/anyproto/any-sync/util/crypto"
)

//
// Mock NodeConf implementation
//

type mockConf struct {
	id            string
	networkId     string
	configuration nodeconf.Configuration
}

func (m *mockConf) NetworkCompatibilityStatus() nodeconf.NetworkCompatibilityStatus {
	return nodeconf.NetworkCompatibilityStatusOk
}

func (m *mockConf) Init(a *app.App) (err error) {
	accountKeys := a.MustComponent(accountService.CName).(accountService.Service).Account()
	networkId := accountKeys.SignKey.GetPublic().Network()
	node := nodeconf.Node{
		PeerId:    accountKeys.PeerId,
		Addresses: []string{"127.0.0.1:4430"},
		Types:     []nodeconf.NodeType{nodeconf.NodeTypeTree},
	}
	m.id = networkId
	m.networkId = networkId
	m.configuration = nodeconf.Configuration{
		Id:           networkId,
		NetworkId:    networkId,
		Nodes:        []nodeconf.Node{node},
		CreationTime: time.Now(),
	}
	return nil
}

func (m *mockConf) Name() (name string) {
	return nodeconf.CName
}

func (m *mockConf) Run(ctx context.Context) (err error) {
	return nil
}

func (m *mockConf) Close(ctx context.Context) (err error) {
	return nil
}

func (m *mockConf) Id() string {
	return m.id
}

func (m *mockConf) Configuration() nodeconf.Configuration {
	return m.configuration
}

func (m *mockConf) NodeIds(spaceId string) []string {
	var nodeIds []string
	for _, node := range m.configuration.Nodes {
		nodeIds = append(nodeIds, node.PeerId)
	}
	return nodeIds
}

func (m *mockConf) IsResponsible(spaceId string) bool {
	return true
}

func (m *mockConf) FilePeers() []string {
	return nil
}

func (m *mockConf) ConsensusPeers() []string {
	return nil
}

func (m *mockConf) CoordinatorPeers() []string {
	return nil
}

func (m *mockConf) NamingNodePeers() []string {
	return nil
}

func (m *mockConf) PaymentProcessingNodePeers() []string {
	return nil
}

func (m *mockConf) PeerAddresses(peerId string) (addrs []string, ok bool) {
	if peerId == m.configuration.Nodes[0].PeerId {
		return m.configuration.Nodes[0].Addresses, true
	}
	return nil, false
}

func (m *mockConf) CHash() chash.CHash {
	return nil
}

func (m *mockConf) Partition(spaceId string) (part int) {
	return 0
}

func (m *mockConf) NodeTypes(nodeId string) []nodeconf.NodeType {
	if nodeId == m.configuration.Nodes[0].PeerId {
		return m.configuration.Nodes[0].Types
	}
	return nil
}

//
// Mock PeerManager
//

type mockPeerManager struct {
}

func (p *mockPeerManager) Init(a *app.App) (err error) {
	return nil
}

func (p *mockPeerManager) Name() (name string) {
	return peermanager.CName
}

func (p *mockPeerManager) SendPeer(ctx context.Context, peerId string, msg *spacesyncproto.ObjectSyncMessage) (err error) {
	return nil
}

func (p *mockPeerManager) Broadcast(ctx context.Context, msg *spacesyncproto.ObjectSyncMessage) (err error) {
	return nil
}

func (p *mockPeerManager) GetResponsiblePeers(ctx context.Context) (peers []peer.Peer, err error) {
	return nil, nil
}

func (p *mockPeerManager) GetNodePeers(ctx context.Context) (peers []peer.Peer, err error) {
	return nil, nil
}

//
// Mock PeerManagerProvider
//

type mockPeerManagerProvider struct {
}

func (m *mockPeerManagerProvider) Init(a *app.App) (err error) {
	return nil
}

func (m *mockPeerManagerProvider) Name() (name string) {
	return peermanager.CName
}

func (m *mockPeerManagerProvider) NewPeerManager(ctx context.Context, spaceId string) (sm peermanager.PeerManager, err error) {
	return synctest.NewCounterPeerManager(), nil
}

//
// Mock StatusServiceProvider
//

type mockStatusServiceProvider struct {
}

func (m *mockStatusServiceProvider) Init(a *app.App) (err error) {
	return nil
}

func (m *mockStatusServiceProvider) Name() (name string) {
	return syncstatus.CName
}

func (m *mockStatusServiceProvider) NewStatusService() syncstatus.StatusService {
	return syncstatus.NewNoOpSyncStatus()
}

//
// Mock Pool
//

type mockPool struct {
}

func (m *mockPool) AddPeer(ctx context.Context, p peer.Peer) (err error) {
	return nil
}

func (m *mockPool) Init(a *app.App) (err error) {
	return nil
}

func (m *mockPool) Name() (name string) {
	return pool.CName
}

func (m *mockPool) Get(ctx context.Context, id string) (peer.Peer, error) {
	return nil, fmt.Errorf("no such peer")
}

func (m *mockPool) Dial(ctx context.Context, id string) (peer.Peer, error) {
	return nil, fmt.Errorf("can't dial peer")
}

func (m *mockPool) GetOneOf(ctx context.Context, peerIds []string) (peer.Peer, error) {
	return nil, fmt.Errorf("can't dial peer")
}

func (m *mockPool) DialOneOf(ctx context.Context, peerIds []string) (peer.Peer, error) {
	return nil, fmt.Errorf("can't dial peer")
}

//
// Mock Config
//

type mockConfig struct {
}

func (m *mockConfig) Init(a *app.App) (err error) {
	return nil
}

func (m *mockConfig) Name() (name string) {
	return "config"
}

func (m *mockConfig) GetSpace() config.Config {
	return config.Config{
		GCTTL:                60,
		SyncPeriod:           5,
		KeepTreeDataInMemory: true,
	}
}

func (m *mockConfig) GetStreamConfig() streampool.StreamConfig {
	return streampool.StreamConfig{
		SendQueueSize:    100,
		DialQueueWorkers: 100,
		DialQueueSize:    100,
	}
}

//
// Mock TreeManager
//

type noOpSyncer struct {
}

func (n noOpSyncer) Init() {
}

func (n noOpSyncer) SyncAll(ctx context.Context, peerId string, existing, missing []string) error {
	return nil
}

func (n noOpSyncer) Close() error {
	return nil
}

type mockTreeSyncer struct {
}

func (m mockTreeSyncer) ShouldSync(peerId string) bool {
	return false
}

func (m mockTreeSyncer) Init(a *app.App) (err error) {
	return nil
}

func (m mockTreeSyncer) Name() (name string) {
	return treesyncer.CName
}

func (m mockTreeSyncer) Run(ctx context.Context) (err error) {
	return nil
}

func (m mockTreeSyncer) Close(ctx context.Context) (err error) {
	return nil
}

func (m mockTreeSyncer) StartSync() {
}

func (m mockTreeSyncer) StopSync() {
}

func (m mockTreeSyncer) SyncAll(ctx context.Context, peerId string, existing, missing []string) error {
	return nil
}

type mockTreeManager struct {
	mx          sync.Mutex
	space       Space
	spaceId     string
	cache       ocache.OCache
	spaceGetter *RpcServer
	accService  accountService.Service
	deletedIds  []string
	markedIds   []string
	treesToPut  map[string]treestorage.TreeStorageCreatePayload
	wait        bool
	waitLoad    chan struct{}
}

func newMockTreeManager(spaceId string) *mockTreeManager {
	return &mockTreeManager{
		spaceId:    spaceId,
		treesToPut: make(map[string]treestorage.TreeStorageCreatePayload),
	}
}

func (t *mockTreeManager) MarkTreeDeleted(ctx context.Context, spaceId, treeId string) error {
	t.mx.Lock()
	defer t.mx.Unlock()
	t.markedIds = append(t.markedIds, treeId)
	return nil
}

func (t *mockTreeManager) getSpace(ctx context.Context) (Space, error) {
	if t.space != nil {
		return t.space, nil
	}
	return t.spaceGetter.GetSpace(ctx, t.spaceId)
}

func (t *mockTreeManager) Init(a *app.App) (err error) {
	t.cache = ocache.New(func(ctx context.Context, id string) (value ocache.Object, err error) {
		if t.wait {
			<-t.waitLoad
		}
		sp, err := t.getSpace(ctx)
		if err != nil {
			return nil, err
		}
		t.mx.Lock()
		if tr, ok := t.treesToPut[id]; ok {
			delete(t.treesToPut, id)
			t.mx.Unlock()
			return sp.TreeBuilder().PutTree(ctx, tr, nil)
		}
		t.mx.Unlock()
		return sp.TreeBuilder().BuildTree(ctx, id, objecttreebuilder.BuildTreeOpts{})
	},
		ocache.WithGCPeriod(time.Minute),
		ocache.WithTTL(time.Duration(60)*time.Minute))
	t.spaceGetter = a.MustComponent(RpcName).(*RpcServer)
	t.accService = a.MustComponent(accountService.CName).(accountService.Service)
	return nil
}

func (t *mockTreeManager) Name() (name string) {
	return treemanager.CName
}

func (t *mockTreeManager) Run(ctx context.Context) (err error) {
	return nil
}

func (t *mockTreeManager) Close(ctx context.Context) (err error) {
	return t.cache.Close()
}

func (t *mockTreeManager) GetTree(ctx context.Context, spaceId, treeId string) (objecttree.ObjectTree, error) {
	val, err := t.cache.Get(ctx, treeId)
	if err != nil {
		return nil, err
	}
	return val.(objecttree.ObjectTree), nil
}

func (t *mockTreeManager) CreateTree(ctx context.Context, spaceId string) (objecttree.ObjectTree, error) {
	sp, err := t.getSpace(ctx)
	if err != nil {
		return nil, err
	}
	rnd := []byte(fmt.Sprint(rand.Uint32()))
	payload := objecttree.ObjectTreeCreatePayload{
		PrivKey:     t.accService.Account().SignKey,
		ChangeType:  "change",
		SpaceId:     spaceId,
		IsEncrypted: true,
		Seed:        rnd,
		Timestamp:   time.Now().UnixNano(),
	}
	createPayload, err := sp.TreeBuilder().CreateTree(ctx, payload)
	if err != nil {
		return nil, err
	}
	return t.PutTree(ctx, createPayload)
}

func (t *mockTreeManager) PutTree(ctx context.Context, payload treestorage.TreeStorageCreatePayload) (objecttree.ObjectTree, error) {
	t.mx.Lock()
	t.treesToPut[payload.RootRawChange.Id] = payload
	t.mx.Unlock()
	return t.GetTree(ctx, t.spaceId, payload.RootRawChange.Id)
}

func (t *mockTreeManager) DeleteTree(ctx context.Context, spaceId, treeId string) (err error) {
	tr, err := t.GetTree(ctx, spaceId, treeId)
	if err != nil {
		return
	}
	err = tr.Delete()
	if err != nil {
		return
	}
	t.deletedIds = append(t.deletedIds, treeId)
	_, err = t.cache.Remove(ctx, treeId)
	return nil
}

type mockCoordinatorClient struct {
}

func (m mockCoordinatorClient) StatusCheckMany(ctx context.Context, spaceIds []string) (statuses []*coordinatorproto.SpaceStatusPayload, limits *coordinatorproto.AccountLimits, err error) {
	return
}

func (m mockCoordinatorClient) SpaceMakeShareable(ctx context.Context, spaceId string) (err error) {
	return
}

func (m mockCoordinatorClient) SpaceMakeUnshareable(ctx context.Context, spaceId, aclId string) (err error) {
	return
}

func (m mockCoordinatorClient) AccountLimitsSet(ctx context.Context, req *coordinatorproto.AccountLimitsSetRequest) error {
	return nil
}

func (m mockCoordinatorClient) SpaceDelete(ctx context.Context, spaceId string, conf *coordinatorproto.DeletionConfirmPayloadWithSignature) (err error) {
	return
}

func (m mockCoordinatorClient) AccountDelete(ctx context.Context, conf *coordinatorproto.DeletionConfirmPayloadWithSignature) (timestamp int64, err error) {
	return
}

func (m mockCoordinatorClient) AccountRevertDeletion(ctx context.Context) (err error) {
	return
}

func (m mockCoordinatorClient) StatusCheck(ctx context.Context, spaceId string) (status *coordinatorproto.SpaceStatusPayload, err error) {
	return
}

func (m mockCoordinatorClient) SpaceSign(ctx context.Context, payload coordinatorclient.SpaceSignPayload) (receipt *coordinatorproto.SpaceReceiptWithSignature, err error) {
	return
}

func (m mockCoordinatorClient) NetworkConfiguration(ctx context.Context, currentId string) (*coordinatorproto.NetworkConfigurationResponse, error) {
	return nil, nil
}

func (m mockCoordinatorClient) DeletionLog(ctx context.Context, lastRecordId string, limit int) (records []*coordinatorproto.DeletionLogRecord, err error) {
	return
}

func (m mockCoordinatorClient) IdentityRepoPut(ctx context.Context, identity string, data []*identityrepoproto.Data) (err error) {
	return
}

func (m mockCoordinatorClient) IdentityRepoGet(ctx context.Context, identities []string, kinds []string) (res []*identityrepoproto.DataWithIdentity, err error) {
	return
}

func (m mockCoordinatorClient) AclAddRecord(ctx context.Context, spaceId string, rec *consensusproto.RawRecord) (res *consensusproto.RawRecordWithId, err error) {
	return
}

func (m mockCoordinatorClient) AclGetRecords(ctx context.Context, spaceId, aclHead string) (res []*consensusproto.RawRecordWithId, err error) {
	return
}

func (m mockCoordinatorClient) Init(a *app.App) (err error) {
	return
}

func (m mockCoordinatorClient) Name() (name string) {
	return coordinatorclient.CName
}

//
// Stream opener
//

func newStreamOpener(spaceId string) streamopener.StreamOpener {
	return &streamOpener{spaceId: spaceId}
}

type streamOpener struct {
	spaceId string
}

func (c *streamOpener) Init(a *app.App) (err error) {
	return nil
}

func (c *streamOpener) Name() (name string) {
	return streamopener.CName
}

func (c *streamOpener) OpenStream(ctx context.Context, p peer.Peer) (stream drpc.Stream, tags []string, err error) {
	conn, err := p.AcquireDrpcConn(ctx)
	if err != nil {
		return
	}
	objectStream, err := spacesyncproto.NewDRPCSpaceSyncClient(conn).ObjectSyncStream(ctx)
	if err != nil {
		return
	}
	var msg = &spacesyncproto.SpaceSubscription{
		SpaceIds: []string{c.spaceId},
		Action:   spacesyncproto.SpaceSubscriptionAction_Subscribe,
	}
	payload, err := msg.Marshal()
	if err != nil {
		return
	}
	if err = objectStream.Send(&spacesyncproto.ObjectSyncMessage{
		Payload: payload,
		SpaceId: c.spaceId,
	}); err != nil {
		return
	}
	return &failingStream{objectStream, false}, nil, nil
}

//
// Space fixture
//

type spaceFixture struct {
	ctx                  context.Context
	app                  *app.App
	config               *mockConfig
	account              accountService.Service
	configurationService nodeconf.Service
	storageProvider      spacestorage.SpaceStorageProvider
	peerManagerProvider  peermanager.PeerManagerProvider
	streamOpener         streamopener.StreamOpener
	credentialProvider   credentialprovider.CredentialProvider
	treeManager          *mockTreeManager
	pool                 *mockPool
	spaceService         SpaceService
	process              *spaceProcess
	cancelFunc           context.CancelFunc
}

func newFixture(t *testing.T) *spaceFixture {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	fx := &spaceFixture{
		cancelFunc:           cancel,
		config:               &mockConfig{},
		app:                  &app.App{},
		account:              &accounttest.AccountTestService{},
		configurationService: &mockConf{},
		storageProvider:      spacestorage.NewInMemorySpaceStorageProvider(),
		peerManagerProvider:  &mockPeerManagerProvider{},
		treeManager:          &mockTreeManager{waitLoad: make(chan struct{})},
		pool:                 &mockPool{},
		spaceService:         New(),
	}
	fx.app.Register(fx.account).
		Register(fx.config).
		Register(credentialprovider.NewNoOp()).
		Register(&mockStatusServiceProvider{}).
		Register(mockCoordinatorClient{}).
		Register(fx.configurationService).
		Register(fx.storageProvider).
		Register(fx.peerManagerProvider).
		Register(fx.treeManager).
		Register(fx.spaceService)
	err := fx.app.Start(ctx)
	if err != nil {
		fx.cancelFunc()
	}
	require.NoError(t, err)
	return fx
}

func newFixtureWithData(t *testing.T, spaceId string, keys *accountdata.AccountKeys, peerPool *synctest.PeerGlobalPool, provider *spacestorage.InMemorySpaceStorageProvider) *spaceFixture {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	fx := &spaceFixture{
		ctx:                  ctx,
		cancelFunc:           cancel,
		app:                  &app.App{},
		config:               &mockConfig{},
		account:              accounttest.NewWithAcc(keys),
		configurationService: &mockConf{},
		storageProvider:      provider,
		streamOpener:         newStreamOpener(spaceId),
		peerManagerProvider:  &mockPeerManagerProvider{},
		treeManager:          newMockTreeManager(spaceId),
		pool:                 &mockPool{},
		spaceService:         New(),
		process:              newSpaceProcess(spaceId),
	}
	fx.app.Register(fx.account).
		Register(fx.config).
		Register(peerPool).
		Register(rpctest.NewTestServer()).
		Register(synctest.NewPeerProvider(keys.PeerId)).
		Register(pool.New()).
		Register(credentialprovider.NewNoOp()).
		Register(&mockStatusServiceProvider{}).
		Register(mockCoordinatorClient{}).
		Register(fx.configurationService).
		Register(fx.storageProvider).
		Register(fx.peerManagerProvider).
		Register(fx.treeManager).
		Register(fx.spaceService).
		Register(NewRpcServer()).
		Register(fx.process)
	err := fx.app.Start(ctx)
	if err != nil {
		fx.cancelFunc()
	}
	require.NoError(t, err)
	return fx
}

type multiPeerFixture struct {
	peerFixtures []*spaceFixture
}

func (m *multiPeerFixture) Close() {
	for _, fx := range m.peerFixtures {
		fx.app.Close(context.Background())
	}
}

func newMultiPeerFixture(t *testing.T, peerNum int) *multiPeerFixture {
	keys, err := accountdata.NewRandom()
	require.NoError(t, err)
	masterKey, _, err := crypto.GenerateRandomEd25519KeyPair()
	require.NoError(t, err)
	metaKey, _, err := crypto.GenerateRandomEd25519KeyPair()
	require.NoError(t, err)
	readKey := crypto.NewAES()
	meta := []byte("account")
	payload := SpaceCreatePayload{
		SigningKey:     keys.SignKey,
		SpaceType:      "space",
		ReplicationKey: 10,
		SpacePayload:   nil,
		MasterKey:      masterKey,
		ReadKey:        readKey,
		MetadataKey:    metaKey,
		Metadata:       meta,
	}
	createSpace, err := storagePayloadForSpaceCreate(payload)
	require.NoError(t, err)
	executor := list.NewExternalKeysAclExecutor(createSpace.SpaceHeaderWithId.Id, keys, meta, createSpace.AclWithId)
	cmds := []string{
		"0.init::0",
		"0.invite::invId",
	}
	for i := 1; i < peerNum; i++ {
		cmds = append(cmds, fmt.Sprintf("%d.join::invId", i))
		cmds = append(cmds, fmt.Sprintf("0.approve::%d,rw", i))
	}
	for _, cmd := range cmds {
		err := executor.Execute(cmd)
		require.NoError(t, err, cmd)
	}
	var (
		allKeys    []*accountdata.AccountKeys
		allRecords []*consensusproto.RawRecordWithId
		providers  []*spacestorage.InMemorySpaceStorageProvider
		peerIds    []string
	)
	allRecords, err = executor.ActualAccounts()["0"].Acl.RecordsAfter(context.Background(), "")
	require.NoError(t, err)
	for i := 0; i < peerNum; i++ {
		allKeys = append(allKeys, executor.ActualAccounts()[fmt.Sprint(i)].Keys)
		peerIds = append(peerIds, executor.ActualAccounts()[fmt.Sprint(i)].Keys.PeerId)
		provider := spacestorage.NewInMemorySpaceStorageProvider().(*spacestorage.InMemorySpaceStorageProvider)
		providers = append(providers, provider)
		spaceStore, err := provider.CreateSpaceStorage(createSpace)
		require.NoError(t, err)
		listStorage, err := spaceStore.AclStorage()
		require.NoError(t, err)
		for _, rec := range allRecords {
			err = listStorage.AddRawRecord(context.Background(), rec)
			require.NoError(t, err)
		}
		err = listStorage.SetHead(allRecords[len(allRecords)-1].Id)
		require.NoError(t, err)
	}
	peerPool := synctest.NewPeerGlobalPool(peerIds)
	peerPool.MakePeers()
	var peerFixtures []*spaceFixture
	for i := 0; i < peerNum; i++ {
		fx := newFixtureWithData(t, createSpace.SpaceHeaderWithId.Id, allKeys[i], peerPool, providers[i])
		peerFixtures = append(peerFixtures, fx)
	}
	return &multiPeerFixture{peerFixtures: peerFixtures}
}

func Test_Sync(t *testing.T) {
	mpFixture := newMultiPeerFixture(t, 3)
	time.Sleep(5 * time.Second)
	for _, fx := range mpFixture.peerFixtures {
		err := fx.process.Close(context.Background())
		require.NoError(t, err)
	}
	time.Sleep(5 * time.Second)
	var hashes []string
	for _, fx := range mpFixture.peerFixtures {
		sp, err := fx.app.MustComponent(RpcName).(*RpcServer).GetSpace(context.Background(), fx.process.spaceId)
		require.NoError(t, err)
		spaceHash, err := sp.Storage().ReadSpaceHash()
		require.NoError(t, err)
		hashes = append(hashes, spaceHash)
	}
	for i := 1; i < len(hashes); i++ {
		require.Equal(t, hashes[0], hashes[i])
	}
	mpFixture.Close()
}
