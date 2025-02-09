// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/dualstaking/delegate.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type Delegation struct {
	Provider  string     `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	ChainID   string     `protobuf:"bytes,2,opt,name=chainID,proto3" json:"chainID,omitempty"`
	Delegator string     `protobuf:"bytes,3,opt,name=delegator,proto3" json:"delegator,omitempty"`
	Amount    types.Coin `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount"`
}

func (m *Delegation) Reset()         { *m = Delegation{} }
func (m *Delegation) String() string { return proto.CompactTextString(m) }
func (*Delegation) ProtoMessage()    {}
func (*Delegation) Descriptor() ([]byte, []int) {
	return fileDescriptor_547eac7f30bf94d4, []int{0}
}
func (m *Delegation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Delegation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Delegation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Delegation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Delegation.Merge(m, src)
}
func (m *Delegation) XXX_Size() int {
	return m.Size()
}
func (m *Delegation) XXX_DiscardUnknown() {
	xxx_messageInfo_Delegation.DiscardUnknown(m)
}

var xxx_messageInfo_Delegation proto.InternalMessageInfo

func (m *Delegation) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *Delegation) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *Delegation) GetDelegator() string {
	if m != nil {
		return m.Delegator
	}
	return ""
}

func (m *Delegation) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

type Delegator struct {
	Providers []string `protobuf:"bytes,1,rep,name=providers,proto3" json:"providers,omitempty"`
}

func (m *Delegator) Reset()         { *m = Delegator{} }
func (m *Delegator) String() string { return proto.CompactTextString(m) }
func (*Delegator) ProtoMessage()    {}
func (*Delegator) Descriptor() ([]byte, []int) {
	return fileDescriptor_547eac7f30bf94d4, []int{1}
}
func (m *Delegator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Delegator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Delegator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Delegator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Delegator.Merge(m, src)
}
func (m *Delegator) XXX_Size() int {
	return m.Size()
}
func (m *Delegator) XXX_DiscardUnknown() {
	xxx_messageInfo_Delegator.DiscardUnknown(m)
}

var xxx_messageInfo_Delegator proto.InternalMessageInfo

func (m *Delegator) GetProviders() []string {
	if m != nil {
		return m.Providers
	}
	return nil
}

func init() {
	proto.RegisterType((*Delegation)(nil), "lavanet.lava.dualstaking.Delegation")
	proto.RegisterType((*Delegator)(nil), "lavanet.lava.dualstaking.Delegator")
}

func init() {
	proto.RegisterFile("lavanet/lava/dualstaking/delegate.proto", fileDescriptor_547eac7f30bf94d4)
}

var fileDescriptor_547eac7f30bf94d4 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x63, 0x5a, 0x15, 0x62, 0xb6, 0x88, 0xc1, 0x54, 0x95, 0xa9, 0xba, 0x50, 0x84, 0x64,
	0xab, 0x30, 0xb0, 0x97, 0x30, 0xb0, 0x76, 0x64, 0x73, 0x12, 0x2b, 0xb5, 0x48, 0x7c, 0x51, 0xec,
	0x44, 0xf0, 0x16, 0xac, 0xbc, 0x51, 0xc7, 0x8e, 0x4c, 0x08, 0x25, 0x2f, 0x82, 0xf2, 0xa7, 0x81,
	0x4e, 0xe7, 0xbb, 0xfb, 0xe4, 0xfb, 0xf4, 0xc3, 0xd7, 0x89, 0x28, 0x85, 0x96, 0x96, 0x37, 0x95,
	0x47, 0x85, 0x48, 0x8c, 0x15, 0xaf, 0x4a, 0xc7, 0x3c, 0x92, 0x89, 0x8c, 0x85, 0x95, 0x2c, 0xcb,
	0xc1, 0x82, 0x47, 0x7a, 0x90, 0x35, 0x95, 0xfd, 0x03, 0xa7, 0x17, 0x31, 0xc4, 0xd0, 0x42, 0xbc,
	0x79, 0x75, 0xfc, 0x94, 0x86, 0x60, 0x52, 0x30, 0x3c, 0x10, 0x46, 0xf2, 0x72, 0x15, 0x48, 0x2b,
	0x56, 0x3c, 0x04, 0xa5, 0xbb, 0xfd, 0xe2, 0x13, 0x61, 0xec, 0x77, 0x27, 0x14, 0x68, 0x6f, 0x8a,
	0xcf, 0xb2, 0x1c, 0x4a, 0x15, 0xc9, 0x9c, 0xa0, 0x39, 0x5a, 0xba, 0x9b, 0xa1, 0xf7, 0x08, 0x3e,
	0x0d, 0xb7, 0x42, 0xe9, 0x67, 0x9f, 0x9c, 0xb4, 0xab, 0x43, 0xeb, 0xcd, 0xb0, 0xdb, 0x6b, 0x42,
	0x4e, 0x46, 0xed, 0xee, 0x6f, 0xe0, 0x3d, 0xe0, 0x89, 0x48, 0xa1, 0xd0, 0x96, 0x8c, 0xe7, 0x68,
	0x79, 0x7e, 0x77, 0xc9, 0x3a, 0x27, 0xd6, 0x38, 0xb1, 0xde, 0x89, 0x3d, 0x82, 0xd2, 0xeb, 0xf1,
	0xee, 0xfb, 0xca, 0xd9, 0xf4, 0xf8, 0xe2, 0x06, 0xbb, 0xfe, 0xf0, 0xcb, 0x0c, 0xbb, 0x07, 0x13,
	0x43, 0xd0, 0x7c, 0xd4, 0xdc, 0x18, 0x06, 0xeb, 0xa7, 0x5d, 0x45, 0xd1, 0xbe, 0xa2, 0xe8, 0xa7,
	0xa2, 0xe8, 0xa3, 0xa6, 0xce, 0xbe, 0xa6, 0xce, 0x57, 0x4d, 0x9d, 0x97, 0xdb, 0x58, 0xd9, 0x6d,
	0x11, 0xb0, 0x10, 0x52, 0x7e, 0x14, 0xf2, 0xdb, 0x51, 0xcc, 0xf6, 0x3d, 0x93, 0x26, 0x98, 0xb4,
	0xa1, 0xdc, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x95, 0x40, 0x43, 0x8f, 0x01, 0x00, 0x00,
}

func (m *Delegation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Delegation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Delegation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDelegate(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Delegator) > 0 {
		i -= len(m.Delegator)
		copy(dAtA[i:], m.Delegator)
		i = encodeVarintDelegate(dAtA, i, uint64(len(m.Delegator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintDelegate(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintDelegate(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Delegator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Delegator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Delegator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Providers) > 0 {
		for iNdEx := len(m.Providers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Providers[iNdEx])
			copy(dAtA[i:], m.Providers[iNdEx])
			i = encodeVarintDelegate(dAtA, i, uint64(len(m.Providers[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintDelegate(dAtA []byte, offset int, v uint64) int {
	offset -= sovDelegate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Delegation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovDelegate(uint64(l))
	}
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovDelegate(uint64(l))
	}
	l = len(m.Delegator)
	if l > 0 {
		n += 1 + l + sovDelegate(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovDelegate(uint64(l))
	return n
}

func (m *Delegator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Providers) > 0 {
		for _, s := range m.Providers {
			l = len(s)
			n += 1 + l + sovDelegate(uint64(l))
		}
	}
	return n
}

func sovDelegate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDelegate(x uint64) (n int) {
	return sovDelegate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Delegation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDelegate
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
			return fmt.Errorf("proto: Delegation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Delegation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegate
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
				return ErrInvalidLengthDelegate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegate
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
				return ErrInvalidLengthDelegate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegate
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
				return ErrInvalidLengthDelegate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDelegate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDelegate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDelegate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDelegate
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
func (m *Delegator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDelegate
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
			return fmt.Errorf("proto: Delegator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Delegator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Providers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegate
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
				return ErrInvalidLengthDelegate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Providers = append(m.Providers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDelegate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDelegate
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
func skipDelegate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDelegate
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
					return 0, ErrIntOverflowDelegate
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
					return 0, ErrIntOverflowDelegate
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
				return 0, ErrInvalidLengthDelegate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDelegate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDelegate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDelegate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDelegate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDelegate = fmt.Errorf("proto: unexpected end of group")
)
