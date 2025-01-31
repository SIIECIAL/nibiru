// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: perp/v2/genesis.proto

package v2

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// GenesisState defines the perp module's genesis state.
type GenesisState struct {
	Params    Params     `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Markets   []Market   `protobuf:"bytes,2,rep,name=markets,proto3" json:"markets"`
	Amms      []AMM      `protobuf:"bytes,3,rep,name=amms,proto3" json:"amms"`
	Positions []Position `protobuf:"bytes,4,rep,name=positions,proto3" json:"positions"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_8edcabc35f3cf683, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetMarkets() []Market {
	if m != nil {
		return m.Markets
	}
	return nil
}

func (m *GenesisState) GetAmms() []AMM {
	if m != nil {
		return m.Amms
	}
	return nil
}

func (m *GenesisState) GetPositions() []Position {
	if m != nil {
		return m.Positions
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "nibiru.perp.v2.GenesisState")
}

func init() { proto.RegisterFile("perp/v2/genesis.proto", fileDescriptor_8edcabc35f3cf683) }

var fileDescriptor_8edcabc35f3cf683 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0x93, 0xaf, 0x55, 0x3f, 0x11, 0x10, 0x43, 0x0a, 0x28, 0xaa, 0x90, 0xa9, 0x98, 0x58,
	0xea, 0xab, 0x06, 0xc4, 0xc4, 0x42, 0x19, 0x3a, 0x15, 0x21, 0xd8, 0xd8, 0x9c, 0xca, 0x72, 0x2d,
	0xb0, 0xaf, 0x15, 0xbb, 0x11, 0xbc, 0x05, 0x8f, 0xd5, 0xb1, 0x23, 0x13, 0x42, 0xed, 0xc4, 0x5b,
	0xa0, 0x38, 0xae, 0x10, 0x7f, 0x36, 0xeb, 0x9e, 0xf3, 0x3b, 0x3e, 0xf7, 0x26, 0xfb, 0x86, 0x97,
	0x06, 0xaa, 0x1c, 0x04, 0xd7, 0xdc, 0x4a, 0x4b, 0x4d, 0x89, 0x0e, 0xd3, 0x5d, 0x2d, 0x0b, 0x59,
	0xce, 0x69, 0xad, 0xd2, 0x2a, 0xef, 0xed, 0x09, 0x14, 0xe8, 0x25, 0xa8, 0x5f, 0x8d, 0xab, 0x77,
	0x28, 0x10, 0xc5, 0x23, 0x07, 0x66, 0x24, 0x30, 0xad, 0xd1, 0x31, 0x27, 0x51, 0x87, 0x8c, 0x1e,
	0x99, 0xa2, 0x55, 0x68, 0xa1, 0x60, 0x96, 0x43, 0x35, 0x2c, 0xb8, 0x63, 0x43, 0x98, 0xa2, 0xd4,
	0x41, 0xef, 0x6e, 0xbe, 0xb6, 0x8e, 0x39, 0xde, 0x0c, 0x8f, 0x3f, 0xe2, 0x64, 0x67, 0xdc, 0x54,
	0xb9, 0xab, 0xc7, 0xe9, 0x59, 0xd2, 0x31, 0xac, 0x64, 0xca, 0x66, 0x71, 0x3f, 0x3e, 0xd9, 0xce,
	0x0f, 0xe8, 0xf7, 0x6a, 0xf4, 0xc6, 0xab, 0xa3, 0xf6, 0xe2, 0xed, 0x28, 0xba, 0x0d, 0xde, 0xf4,
	0x3c, 0xf9, 0xaf, 0x58, 0xf9, 0xc0, 0x9d, 0xcd, 0xfe, 0xf5, 0x5b, 0x7f, 0x61, 0x13, 0x2f, 0x07,
	0x6c, 0x63, 0x4e, 0x07, 0x49, 0x9b, 0x29, 0x65, 0xb3, 0x96, 0x87, 0xba, 0x3f, 0xa1, 0xcb, 0xc9,
	0x24, 0x10, 0xde, 0x96, 0x5e, 0x24, 0x5b, 0x06, 0xad, 0xf4, 0x5b, 0x67, 0x6d, 0xcf, 0x64, 0xbf,
	0xfa, 0x05, 0x43, 0x00, 0xbf, 0x80, 0xd1, 0x78, 0xb1, 0x22, 0xf1, 0x72, 0x45, 0xe2, 0xf7, 0x15,
	0x89, 0x5f, 0xd6, 0x24, 0x5a, 0xae, 0x49, 0xf4, 0xba, 0x26, 0xd1, 0xfd, 0x40, 0x48, 0x37, 0x9b,
	0x17, 0x74, 0x8a, 0x0a, 0xae, 0x7d, 0xdc, 0xd5, 0x8c, 0x49, 0x0d, 0x4d, 0x34, 0x3c, 0x81, 0x3f,
	0x9d, 0x7b, 0x36, 0xdc, 0x42, 0x95, 0x17, 0x1d, 0x7f, 0xbb, 0xd3, 0xcf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x74, 0x8f, 0x37, 0x3b, 0xcd, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Positions) > 0 {
		for iNdEx := len(m.Positions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Positions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Amms) > 0 {
		for iNdEx := len(m.Amms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Markets) > 0 {
		for iNdEx := len(m.Markets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Markets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Markets) > 0 {
		for _, e := range m.Markets {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Amms) > 0 {
		for _, e := range m.Amms {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Positions) > 0 {
		for _, e := range m.Positions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Markets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Markets = append(m.Markets, Market{})
			if err := m.Markets[len(m.Markets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amms", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amms = append(m.Amms, AMM{})
			if err := m.Amms[len(m.Amms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Positions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Positions = append(m.Positions, Position{})
			if err := m.Positions[len(m.Positions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
