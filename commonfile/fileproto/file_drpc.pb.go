// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.34
// source: commonfile/fileproto/protos/file.proto

package fileproto

import (
	bytes "bytes"
	context "context"
	errors "errors"
	jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_commonfile_fileproto_protos_file_proto struct{}

func (drpcEncoding_File_commonfile_fileproto_protos_file_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_commonfile_fileproto_protos_file_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_commonfile_fileproto_protos_file_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	var buf bytes.Buffer
	err := new(jsonpb.Marshaler).Marshal(&buf, msg.(proto.Message))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (drpcEncoding_File_commonfile_fileproto_protos_file_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return jsonpb.Unmarshal(bytes.NewReader(buf), msg.(proto.Message))
}

type DRPCFileClient interface {
	DRPCConn() drpc.Conn

	BlockGet(ctx context.Context, in *BlockGetRequest) (*BlockGetResponse, error)
	BlockPush(ctx context.Context, in *BlockPushRequest) (*Ok, error)
	BlocksCheck(ctx context.Context, in *BlocksCheckRequest) (*BlocksCheckResponse, error)
	BlocksBind(ctx context.Context, in *BlocksBindRequest) (*Ok, error)
	FilesDelete(ctx context.Context, in *FilesDeleteRequest) (*FilesDeleteResponse, error)
	FilesInfo(ctx context.Context, in *FilesInfoRequest) (*FilesInfoResponse, error)
	FilesGet(ctx context.Context, in *FilesGetRequest) (DRPCFile_FilesGetClient, error)
	Check(ctx context.Context, in *CheckRequest) (*CheckResponse, error)
	SpaceInfo(ctx context.Context, in *SpaceInfoRequest) (*SpaceInfoResponse, error)
	AccountInfo(ctx context.Context, in *AccountInfoRequest) (*AccountInfoResponse, error)
	AccountLimitSet(ctx context.Context, in *AccountLimitSetRequest) (*Ok, error)
	SpaceLimitSet(ctx context.Context, in *SpaceLimitSetRequest) (*Ok, error)
}

type drpcFileClient struct {
	cc drpc.Conn
}

func NewDRPCFileClient(cc drpc.Conn) DRPCFileClient {
	return &drpcFileClient{cc}
}

func (c *drpcFileClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcFileClient) BlockGet(ctx context.Context, in *BlockGetRequest) (*BlockGetResponse, error) {
	out := new(BlockGetResponse)
	err := c.cc.Invoke(ctx, "/filesync.File/BlockGet", drpcEncoding_File_commonfile_fileproto_protos_file_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFileClient) BlockPush(ctx context.Context, in *BlockPushRequest) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/filesync.File/BlockPush", drpcEncoding_File_commonfile_fileproto_protos_file_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFileClient) BlocksCheck(ctx context.Context, in *BlocksCheckRequest) (*BlocksCheckResponse, error) {
	out := new(BlocksCheckResponse)
	err := c.cc.Invoke(ctx, "/filesync.File/BlocksCheck", drpcEncoding_File_commonfile_fileproto_protos_file_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFileClient) BlocksBind(ctx context.Context, in *BlocksBindRequest) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/filesync.File/BlocksBind", drpcEncoding_File_commonfile_fileproto_protos_file_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFileClient) FilesDelete(ctx context.Context, in *FilesDeleteRequest) (*FilesDeleteResponse, error) {
	out := new(FilesDeleteResponse)
	err := c.cc.Invoke(ctx, "/filesync.File/FilesDelete", drpcEncoding_File_commonfile_fileproto_protos_file_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFileClient) FilesInfo(ctx context.Context, in *FilesInfoRequest) (*FilesInfoResponse, error) {
	out := new(FilesInfoResponse)
	err := c.cc.Invoke(ctx, "/filesync.File/FilesInfo", drpcEncoding_File_commonfile_fileproto_protos_file_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFileClient) FilesGet(ctx context.Context, in *FilesGetRequest) (DRPCFile_FilesGetClient, error) {
	stream, err := c.cc.NewStream(ctx, "/filesync.File/FilesGet", drpcEncoding_File_commonfile_fileproto_protos_file_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcFile_FilesGetClient{stream}
	if err := x.MsgSend(in, drpcEncoding_File_commonfile_fileproto_protos_file_proto{}); err != nil {
		return nil, err
	}
	if err := x.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DRPCFile_FilesGetClient interface {
	drpc.Stream
	Recv() (*FilesGetResponse, error)
}

type drpcFile_FilesGetClient struct {
	drpc.Stream
}

func (x *drpcFile_FilesGetClient) GetStream() drpc.Stream {
	return x.Stream
}

func (x *drpcFile_FilesGetClient) Recv() (*FilesGetResponse, error) {
	m := new(FilesGetResponse)
	if err := x.MsgRecv(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcFile_FilesGetClient) RecvMsg(m *FilesGetResponse) error {
	return x.MsgRecv(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{})
}

func (c *drpcFileClient) Check(ctx context.Context, in *CheckRequest) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/filesync.File/Check", drpcEncoding_File_commonfile_fileproto_protos_file_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFileClient) SpaceInfo(ctx context.Context, in *SpaceInfoRequest) (*SpaceInfoResponse, error) {
	out := new(SpaceInfoResponse)
	err := c.cc.Invoke(ctx, "/filesync.File/SpaceInfo", drpcEncoding_File_commonfile_fileproto_protos_file_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFileClient) AccountInfo(ctx context.Context, in *AccountInfoRequest) (*AccountInfoResponse, error) {
	out := new(AccountInfoResponse)
	err := c.cc.Invoke(ctx, "/filesync.File/AccountInfo", drpcEncoding_File_commonfile_fileproto_protos_file_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFileClient) AccountLimitSet(ctx context.Context, in *AccountLimitSetRequest) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/filesync.File/AccountLimitSet", drpcEncoding_File_commonfile_fileproto_protos_file_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFileClient) SpaceLimitSet(ctx context.Context, in *SpaceLimitSetRequest) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/filesync.File/SpaceLimitSet", drpcEncoding_File_commonfile_fileproto_protos_file_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCFileServer interface {
	BlockGet(context.Context, *BlockGetRequest) (*BlockGetResponse, error)
	BlockPush(context.Context, *BlockPushRequest) (*Ok, error)
	BlocksCheck(context.Context, *BlocksCheckRequest) (*BlocksCheckResponse, error)
	BlocksBind(context.Context, *BlocksBindRequest) (*Ok, error)
	FilesDelete(context.Context, *FilesDeleteRequest) (*FilesDeleteResponse, error)
	FilesInfo(context.Context, *FilesInfoRequest) (*FilesInfoResponse, error)
	FilesGet(*FilesGetRequest, DRPCFile_FilesGetStream) error
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	SpaceInfo(context.Context, *SpaceInfoRequest) (*SpaceInfoResponse, error)
	AccountInfo(context.Context, *AccountInfoRequest) (*AccountInfoResponse, error)
	AccountLimitSet(context.Context, *AccountLimitSetRequest) (*Ok, error)
	SpaceLimitSet(context.Context, *SpaceLimitSetRequest) (*Ok, error)
}

type DRPCFileUnimplementedServer struct{}

func (s *DRPCFileUnimplementedServer) BlockGet(context.Context, *BlockGetRequest) (*BlockGetResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFileUnimplementedServer) BlockPush(context.Context, *BlockPushRequest) (*Ok, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFileUnimplementedServer) BlocksCheck(context.Context, *BlocksCheckRequest) (*BlocksCheckResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFileUnimplementedServer) BlocksBind(context.Context, *BlocksBindRequest) (*Ok, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFileUnimplementedServer) FilesDelete(context.Context, *FilesDeleteRequest) (*FilesDeleteResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFileUnimplementedServer) FilesInfo(context.Context, *FilesInfoRequest) (*FilesInfoResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFileUnimplementedServer) FilesGet(*FilesGetRequest, DRPCFile_FilesGetStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFileUnimplementedServer) Check(context.Context, *CheckRequest) (*CheckResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFileUnimplementedServer) SpaceInfo(context.Context, *SpaceInfoRequest) (*SpaceInfoResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFileUnimplementedServer) AccountInfo(context.Context, *AccountInfoRequest) (*AccountInfoResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFileUnimplementedServer) AccountLimitSet(context.Context, *AccountLimitSetRequest) (*Ok, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFileUnimplementedServer) SpaceLimitSet(context.Context, *SpaceLimitSetRequest) (*Ok, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCFileDescription struct{}

func (DRPCFileDescription) NumMethods() int { return 12 }

func (DRPCFileDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/filesync.File/BlockGet", drpcEncoding_File_commonfile_fileproto_protos_file_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFileServer).
					BlockGet(
						ctx,
						in1.(*BlockGetRequest),
					)
			}, DRPCFileServer.BlockGet, true
	case 1:
		return "/filesync.File/BlockPush", drpcEncoding_File_commonfile_fileproto_protos_file_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFileServer).
					BlockPush(
						ctx,
						in1.(*BlockPushRequest),
					)
			}, DRPCFileServer.BlockPush, true
	case 2:
		return "/filesync.File/BlocksCheck", drpcEncoding_File_commonfile_fileproto_protos_file_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFileServer).
					BlocksCheck(
						ctx,
						in1.(*BlocksCheckRequest),
					)
			}, DRPCFileServer.BlocksCheck, true
	case 3:
		return "/filesync.File/BlocksBind", drpcEncoding_File_commonfile_fileproto_protos_file_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFileServer).
					BlocksBind(
						ctx,
						in1.(*BlocksBindRequest),
					)
			}, DRPCFileServer.BlocksBind, true
	case 4:
		return "/filesync.File/FilesDelete", drpcEncoding_File_commonfile_fileproto_protos_file_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFileServer).
					FilesDelete(
						ctx,
						in1.(*FilesDeleteRequest),
					)
			}, DRPCFileServer.FilesDelete, true
	case 5:
		return "/filesync.File/FilesInfo", drpcEncoding_File_commonfile_fileproto_protos_file_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFileServer).
					FilesInfo(
						ctx,
						in1.(*FilesInfoRequest),
					)
			}, DRPCFileServer.FilesInfo, true
	case 6:
		return "/filesync.File/FilesGet", drpcEncoding_File_commonfile_fileproto_protos_file_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCFileServer).
					FilesGet(
						in1.(*FilesGetRequest),
						&drpcFile_FilesGetStream{in2.(drpc.Stream)},
					)
			}, DRPCFileServer.FilesGet, true
	case 7:
		return "/filesync.File/Check", drpcEncoding_File_commonfile_fileproto_protos_file_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFileServer).
					Check(
						ctx,
						in1.(*CheckRequest),
					)
			}, DRPCFileServer.Check, true
	case 8:
		return "/filesync.File/SpaceInfo", drpcEncoding_File_commonfile_fileproto_protos_file_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFileServer).
					SpaceInfo(
						ctx,
						in1.(*SpaceInfoRequest),
					)
			}, DRPCFileServer.SpaceInfo, true
	case 9:
		return "/filesync.File/AccountInfo", drpcEncoding_File_commonfile_fileproto_protos_file_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFileServer).
					AccountInfo(
						ctx,
						in1.(*AccountInfoRequest),
					)
			}, DRPCFileServer.AccountInfo, true
	case 10:
		return "/filesync.File/AccountLimitSet", drpcEncoding_File_commonfile_fileproto_protos_file_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFileServer).
					AccountLimitSet(
						ctx,
						in1.(*AccountLimitSetRequest),
					)
			}, DRPCFileServer.AccountLimitSet, true
	case 11:
		return "/filesync.File/SpaceLimitSet", drpcEncoding_File_commonfile_fileproto_protos_file_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFileServer).
					SpaceLimitSet(
						ctx,
						in1.(*SpaceLimitSetRequest),
					)
			}, DRPCFileServer.SpaceLimitSet, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterFile(mux drpc.Mux, impl DRPCFileServer) error {
	return mux.Register(impl, DRPCFileDescription{})
}

type DRPCFile_BlockGetStream interface {
	drpc.Stream
	SendAndClose(*BlockGetResponse) error
}

type drpcFile_BlockGetStream struct {
	drpc.Stream
}

func (x *drpcFile_BlockGetStream) SendAndClose(m *BlockGetResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFile_BlockPushStream interface {
	drpc.Stream
	SendAndClose(*Ok) error
}

type drpcFile_BlockPushStream struct {
	drpc.Stream
}

func (x *drpcFile_BlockPushStream) SendAndClose(m *Ok) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFile_BlocksCheckStream interface {
	drpc.Stream
	SendAndClose(*BlocksCheckResponse) error
}

type drpcFile_BlocksCheckStream struct {
	drpc.Stream
}

func (x *drpcFile_BlocksCheckStream) SendAndClose(m *BlocksCheckResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFile_BlocksBindStream interface {
	drpc.Stream
	SendAndClose(*Ok) error
}

type drpcFile_BlocksBindStream struct {
	drpc.Stream
}

func (x *drpcFile_BlocksBindStream) SendAndClose(m *Ok) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFile_FilesDeleteStream interface {
	drpc.Stream
	SendAndClose(*FilesDeleteResponse) error
}

type drpcFile_FilesDeleteStream struct {
	drpc.Stream
}

func (x *drpcFile_FilesDeleteStream) SendAndClose(m *FilesDeleteResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFile_FilesInfoStream interface {
	drpc.Stream
	SendAndClose(*FilesInfoResponse) error
}

type drpcFile_FilesInfoStream struct {
	drpc.Stream
}

func (x *drpcFile_FilesInfoStream) SendAndClose(m *FilesInfoResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFile_FilesGetStream interface {
	drpc.Stream
	Send(*FilesGetResponse) error
}

type drpcFile_FilesGetStream struct {
	drpc.Stream
}

func (x *drpcFile_FilesGetStream) Send(m *FilesGetResponse) error {
	return x.MsgSend(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{})
}

type DRPCFile_CheckStream interface {
	drpc.Stream
	SendAndClose(*CheckResponse) error
}

type drpcFile_CheckStream struct {
	drpc.Stream
}

func (x *drpcFile_CheckStream) SendAndClose(m *CheckResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFile_SpaceInfoStream interface {
	drpc.Stream
	SendAndClose(*SpaceInfoResponse) error
}

type drpcFile_SpaceInfoStream struct {
	drpc.Stream
}

func (x *drpcFile_SpaceInfoStream) SendAndClose(m *SpaceInfoResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFile_AccountInfoStream interface {
	drpc.Stream
	SendAndClose(*AccountInfoResponse) error
}

type drpcFile_AccountInfoStream struct {
	drpc.Stream
}

func (x *drpcFile_AccountInfoStream) SendAndClose(m *AccountInfoResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFile_AccountLimitSetStream interface {
	drpc.Stream
	SendAndClose(*Ok) error
}

type drpcFile_AccountLimitSetStream struct {
	drpc.Stream
}

func (x *drpcFile_AccountLimitSetStream) SendAndClose(m *Ok) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFile_SpaceLimitSetStream interface {
	drpc.Stream
	SendAndClose(*Ok) error
}

type drpcFile_SpaceLimitSetStream struct {
	drpc.Stream
}

func (x *drpcFile_SpaceLimitSetStream) SendAndClose(m *Ok) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}
