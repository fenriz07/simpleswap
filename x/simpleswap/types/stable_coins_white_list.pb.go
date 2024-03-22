// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: simpleswap/simpleswap/stable_coins_white_list.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type StableCoinsWhiteList struct {
	Index     string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Coin      string `protobuf:"bytes,2,opt,name=coin,proto3" json:"coin,omitempty"`
	Available bool   `protobuf:"varint,3,opt,name=available,proto3" json:"available,omitempty"`
}

func (m *StableCoinsWhiteList) Reset()         { *m = StableCoinsWhiteList{} }
func (m *StableCoinsWhiteList) String() string { return proto.CompactTextString(m) }
func (*StableCoinsWhiteList) ProtoMessage()    {}
func (*StableCoinsWhiteList) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d61eac9a32eed23, []int{0}
}
func (m *StableCoinsWhiteList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StableCoinsWhiteList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StableCoinsWhiteList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StableCoinsWhiteList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StableCoinsWhiteList.Merge(m, src)
}
func (m *StableCoinsWhiteList) XXX_Size() int {
	return m.Size()
}
func (m *StableCoinsWhiteList) XXX_DiscardUnknown() {
	xxx_messageInfo_StableCoinsWhiteList.DiscardUnknown(m)
}

var xxx_messageInfo_StableCoinsWhiteList proto.InternalMessageInfo

func (m *StableCoinsWhiteList) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *StableCoinsWhiteList) GetCoin() string {
	if m != nil {
		return m.Coin
	}
	return ""
}

func (m *StableCoinsWhiteList) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func init() {
	proto.RegisterType((*StableCoinsWhiteList)(nil), "simpleswap.simpleswap.StableCoinsWhiteList")
}

func init() {
	proto.RegisterFile("simpleswap/simpleswap/stable_coins_white_list.proto", fileDescriptor_9d61eac9a32eed23)
}

var fileDescriptor_9d61eac9a32eed23 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2e, 0xce, 0xcc, 0x2d,
	0xc8, 0x49, 0x2d, 0x2e, 0x4f, 0x2c, 0xd0, 0x47, 0x66, 0x96, 0x24, 0x26, 0xe5, 0xa4, 0xc6, 0x27,
	0xe7, 0x67, 0xe6, 0x15, 0xc7, 0x97, 0x67, 0x64, 0x96, 0xa4, 0xc6, 0xe7, 0x64, 0x16, 0x97, 0xe8,
	0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89, 0x22, 0x54, 0xea, 0x21, 0x98, 0x4a, 0x71, 0x5c, 0x22,
	0xc1, 0x60, 0x7d, 0xce, 0x20, 0x6d, 0xe1, 0x20, 0x5d, 0x3e, 0x99, 0xc5, 0x25, 0x42, 0x22, 0x5c,
	0xac, 0x99, 0x79, 0x29, 0xa9, 0x15, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90,
	0x10, 0x17, 0x0b, 0xc8, 0x78, 0x09, 0x26, 0xb0, 0x20, 0x98, 0x2d, 0x24, 0xc3, 0xc5, 0x99, 0x58,
	0x96, 0x98, 0x99, 0x03, 0x32, 0x44, 0x82, 0x59, 0x81, 0x51, 0x83, 0x23, 0x08, 0x21, 0xe0, 0xe4,
	0x7d, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c,
	0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x86, 0xe9, 0x99, 0x25, 0x19,
	0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x69, 0xa9, 0x79, 0x45, 0x99, 0x55, 0x06, 0xe6, 0xc8,
	0xde, 0xa9, 0x40, 0xe6, 0x94, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0xbd, 0x62, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x7c, 0x62, 0x38, 0xe7, 0x01, 0x01, 0x00, 0x00,
}

func (m *StableCoinsWhiteList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StableCoinsWhiteList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StableCoinsWhiteList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Available {
		i--
		if m.Available {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Coin) > 0 {
		i -= len(m.Coin)
		copy(dAtA[i:], m.Coin)
		i = encodeVarintStableCoinsWhiteList(dAtA, i, uint64(len(m.Coin)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintStableCoinsWhiteList(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStableCoinsWhiteList(dAtA []byte, offset int, v uint64) int {
	offset -= sovStableCoinsWhiteList(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StableCoinsWhiteList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovStableCoinsWhiteList(uint64(l))
	}
	l = len(m.Coin)
	if l > 0 {
		n += 1 + l + sovStableCoinsWhiteList(uint64(l))
	}
	if m.Available {
		n += 2
	}
	return n
}

func sovStableCoinsWhiteList(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStableCoinsWhiteList(x uint64) (n int) {
	return sovStableCoinsWhiteList(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StableCoinsWhiteList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStableCoinsWhiteList
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
			return fmt.Errorf("proto: StableCoinsWhiteList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StableCoinsWhiteList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStableCoinsWhiteList
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
				return ErrInvalidLengthStableCoinsWhiteList
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStableCoinsWhiteList
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStableCoinsWhiteList
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
				return ErrInvalidLengthStableCoinsWhiteList
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStableCoinsWhiteList
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Available", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStableCoinsWhiteList
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Available = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipStableCoinsWhiteList(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStableCoinsWhiteList
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
func skipStableCoinsWhiteList(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStableCoinsWhiteList
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
					return 0, ErrIntOverflowStableCoinsWhiteList
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
					return 0, ErrIntOverflowStableCoinsWhiteList
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
				return 0, ErrInvalidLengthStableCoinsWhiteList
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStableCoinsWhiteList
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStableCoinsWhiteList
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStableCoinsWhiteList        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStableCoinsWhiteList          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStableCoinsWhiteList = fmt.Errorf("proto: unexpected end of group")
)
