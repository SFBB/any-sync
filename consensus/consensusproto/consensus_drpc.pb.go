// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.34
// source: consensus/consensusproto/protos/consensus.proto

package consensusproto

import (
	bytes "bytes"
	context "context"
	errors "errors"
	jsonpb "github.com/anyproto/protobuf/jsonpb"
	proto "github.com/anyproto/protobuf/proto"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_consensus_consensusproto_protos_consensus_proto struct{}

func (drpcEncoding_File_consensus_consensusproto_protos_consensus_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_consensus_consensusproto_protos_consensus_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_consensus_consensusproto_protos_consensus_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	var buf bytes.Buffer
	err := new(jsonpb.Marshaler).Marshal(&buf, msg.(proto.Message))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (drpcEncoding_File_consensus_consensusproto_protos_consensus_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return jsonpb.Unmarshal(bytes.NewReader(buf), msg.(proto.Message))
}

type DRPCConsensusClient interface {
	DRPCConn() drpc.Conn

	LogAdd(ctx context.Context, in *LogAddRequest) (*Ok, error)
	RecordAdd(ctx context.Context, in *RecordAddRequest) (*RawRecordWithId, error)
	LogWatch(ctx context.Context) (DRPCConsensus_LogWatchClient, error)
	LogDelete(ctx context.Context, in *LogDeleteRequest) (*Ok, error)
}

type drpcConsensusClient struct {
	cc drpc.Conn
}

func NewDRPCConsensusClient(cc drpc.Conn) DRPCConsensusClient {
	return &drpcConsensusClient{cc}
}

func (c *drpcConsensusClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcConsensusClient) LogAdd(ctx context.Context, in *LogAddRequest) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/consensusProto.Consensus/LogAdd", drpcEncoding_File_consensus_consensusproto_protos_consensus_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcConsensusClient) RecordAdd(ctx context.Context, in *RecordAddRequest) (*RawRecordWithId, error) {
	out := new(RawRecordWithId)
	err := c.cc.Invoke(ctx, "/consensusProto.Consensus/RecordAdd", drpcEncoding_File_consensus_consensusproto_protos_consensus_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcConsensusClient) LogWatch(ctx context.Context) (DRPCConsensus_LogWatchClient, error) {
	stream, err := c.cc.NewStream(ctx, "/consensusProto.Consensus/LogWatch", drpcEncoding_File_consensus_consensusproto_protos_consensus_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcConsensus_LogWatchClient{stream}
	return x, nil
}

type DRPCConsensus_LogWatchClient interface {
	drpc.Stream
	Send(*LogWatchRequest) error
	Recv() (*LogWatchEvent, error)
}

type drpcConsensus_LogWatchClient struct {
	drpc.Stream
}

func (x *drpcConsensus_LogWatchClient) GetStream() drpc.Stream {
	return x.Stream
}

func (x *drpcConsensus_LogWatchClient) Send(m *LogWatchRequest) error {
	return x.MsgSend(m, drpcEncoding_File_consensus_consensusproto_protos_consensus_proto{})
}

func (x *drpcConsensus_LogWatchClient) Recv() (*LogWatchEvent, error) {
	m := new(LogWatchEvent)
	if err := x.MsgRecv(m, drpcEncoding_File_consensus_consensusproto_protos_consensus_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcConsensus_LogWatchClient) RecvMsg(m *LogWatchEvent) error {
	return x.MsgRecv(m, drpcEncoding_File_consensus_consensusproto_protos_consensus_proto{})
}

func (c *drpcConsensusClient) LogDelete(ctx context.Context, in *LogDeleteRequest) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/consensusProto.Consensus/LogDelete", drpcEncoding_File_consensus_consensusproto_protos_consensus_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCConsensusServer interface {
	LogAdd(context.Context, *LogAddRequest) (*Ok, error)
	RecordAdd(context.Context, *RecordAddRequest) (*RawRecordWithId, error)
	LogWatch(DRPCConsensus_LogWatchStream) error
	LogDelete(context.Context, *LogDeleteRequest) (*Ok, error)
}

type DRPCConsensusUnimplementedServer struct{}

func (s *DRPCConsensusUnimplementedServer) LogAdd(context.Context, *LogAddRequest) (*Ok, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCConsensusUnimplementedServer) RecordAdd(context.Context, *RecordAddRequest) (*RawRecordWithId, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCConsensusUnimplementedServer) LogWatch(DRPCConsensus_LogWatchStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCConsensusUnimplementedServer) LogDelete(context.Context, *LogDeleteRequest) (*Ok, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCConsensusDescription struct{}

func (DRPCConsensusDescription) NumMethods() int { return 4 }

func (DRPCConsensusDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/consensusProto.Consensus/LogAdd", drpcEncoding_File_consensus_consensusproto_protos_consensus_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCConsensusServer).
					LogAdd(
						ctx,
						in1.(*LogAddRequest),
					)
			}, DRPCConsensusServer.LogAdd, true
	case 1:
		return "/consensusProto.Consensus/RecordAdd", drpcEncoding_File_consensus_consensusproto_protos_consensus_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCConsensusServer).
					RecordAdd(
						ctx,
						in1.(*RecordAddRequest),
					)
			}, DRPCConsensusServer.RecordAdd, true
	case 2:
		return "/consensusProto.Consensus/LogWatch", drpcEncoding_File_consensus_consensusproto_protos_consensus_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCConsensusServer).
					LogWatch(
						&drpcConsensus_LogWatchStream{in1.(drpc.Stream)},
					)
			}, DRPCConsensusServer.LogWatch, true
	case 3:
		return "/consensusProto.Consensus/LogDelete", drpcEncoding_File_consensus_consensusproto_protos_consensus_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCConsensusServer).
					LogDelete(
						ctx,
						in1.(*LogDeleteRequest),
					)
			}, DRPCConsensusServer.LogDelete, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterConsensus(mux drpc.Mux, impl DRPCConsensusServer) error {
	return mux.Register(impl, DRPCConsensusDescription{})
}

type DRPCConsensus_LogAddStream interface {
	drpc.Stream
	SendAndClose(*Ok) error
}

type drpcConsensus_LogAddStream struct {
	drpc.Stream
}

func (x *drpcConsensus_LogAddStream) SendAndClose(m *Ok) error {
	if err := x.MsgSend(m, drpcEncoding_File_consensus_consensusproto_protos_consensus_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCConsensus_RecordAddStream interface {
	drpc.Stream
	SendAndClose(*RawRecordWithId) error
}

type drpcConsensus_RecordAddStream struct {
	drpc.Stream
}

func (x *drpcConsensus_RecordAddStream) SendAndClose(m *RawRecordWithId) error {
	if err := x.MsgSend(m, drpcEncoding_File_consensus_consensusproto_protos_consensus_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCConsensus_LogWatchStream interface {
	drpc.Stream
	Send(*LogWatchEvent) error
	Recv() (*LogWatchRequest, error)
}

type drpcConsensus_LogWatchStream struct {
	drpc.Stream
}

func (x *drpcConsensus_LogWatchStream) Send(m *LogWatchEvent) error {
	return x.MsgSend(m, drpcEncoding_File_consensus_consensusproto_protos_consensus_proto{})
}

func (x *drpcConsensus_LogWatchStream) Recv() (*LogWatchRequest, error) {
	m := new(LogWatchRequest)
	if err := x.MsgRecv(m, drpcEncoding_File_consensus_consensusproto_protos_consensus_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcConsensus_LogWatchStream) RecvMsg(m *LogWatchRequest) error {
	return x.MsgRecv(m, drpcEncoding_File_consensus_consensusproto_protos_consensus_proto{})
}

type DRPCConsensus_LogDeleteStream interface {
	drpc.Stream
	SendAndClose(*Ok) error
}

type drpcConsensus_LogDeleteStream struct {
	drpc.Stream
}

func (x *drpcConsensus_LogDeleteStream) SendAndClose(m *Ok) error {
	if err := x.MsgSend(m, drpcEncoding_File_consensus_consensusproto_protos_consensus_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}
