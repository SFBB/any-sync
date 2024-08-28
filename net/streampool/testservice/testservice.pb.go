// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: net/streampool/testservice/protos/testservice.proto

package testservice

import (
	fmt "fmt"
	proto "github.com/anyproto/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type StreamMessage struct {
	ReqData string `protobuf:"bytes,1,opt,name=reqData,proto3" json:"reqData,omitempty"`
}

func (m *StreamMessage) Reset()         { *m = StreamMessage{} }
func (m *StreamMessage) String() string { return proto.CompactTextString(m) }
func (*StreamMessage) ProtoMessage()    {}
func (*StreamMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c28d5a3a78be18f, []int{0}
}
func (m *StreamMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StreamMessage) XXX_MarshalAppend(b []byte, newLen int) ([]byte, error) {
	b = b[:newLen]
	_, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (m *StreamMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMessage.Merge(m, src)
}
func (m *StreamMessage) XXX_Size() int {
	return m.Size()
}
func (m *StreamMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMessage proto.InternalMessageInfo

func (m *StreamMessage) GetReqData() string {
	if m != nil {
		return m.ReqData
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamMessage)(nil), "testService.StreamMessage")
}

func init() {
	proto.RegisterFile("net/streampool/testservice/protos/testservice.proto", fileDescriptor_1c28d5a3a78be18f)
}

var fileDescriptor_1c28d5a3a78be18f = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xce, 0x4b, 0x2d, 0xd1,
	0x2f, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0x2d, 0xc8, 0xcf, 0xcf, 0xd1, 0x2f, 0x49, 0x2d, 0x2e, 0x29,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x2f, 0x46, 0x16, 0xd2,
	0x03, 0x0b, 0x09, 0x71, 0x83, 0x84, 0x82, 0x21, 0x42, 0x4a, 0x9a, 0x5c, 0xbc, 0xc1, 0x60, 0xfd,
	0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x42, 0x12, 0x5c, 0xec, 0x45, 0xa9, 0x85, 0x2e, 0x89,
	0x25, 0x89, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x51, 0x00, 0x17, 0x4b, 0x48,
	0x6a, 0x71, 0x89, 0x90, 0x07, 0x17, 0x17, 0x88, 0x86, 0x68, 0x13, 0x92, 0xd2, 0x43, 0x32, 0x4e,
	0x0f, 0xc5, 0x2c, 0x29, 0x3c, 0x72, 0x1a, 0x8c, 0x06, 0x8c, 0x4e, 0x26, 0x27, 0x1e, 0xc9, 0x31,
	0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb,
	0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x25, 0x85, 0xdb, 0x63, 0x49, 0x6c, 0x60, 0x6f, 0x18, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x59, 0x8d, 0x93, 0xfd, 0x00, 0x00, 0x00,
}

func (m *StreamMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StreamMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ReqData) > 0 {
		i -= len(m.ReqData)
		copy(dAtA[i:], m.ReqData)
		i = encodeVarintTestservice(dAtA, i, uint64(len(m.ReqData)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTestservice(dAtA []byte, offset int, v uint64) int {
	offset -= sovTestservice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StreamMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ReqData)
	if l > 0 {
		n += 1 + l + sovTestservice(uint64(l))
	}
	return n
}

func sovTestservice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTestservice(x uint64) (n int) {
	return sovTestservice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StreamMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTestservice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StreamMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReqData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestservice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTestservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTestservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReqData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTestservice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTestservice
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTestservice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTestservice
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTestservice
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTestservice
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTestservice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTestservice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTestservice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTestservice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTestservice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTestservice = fmt.Errorf("proto: unexpected end of group")
)
