// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vcoa/vnft/class_data.proto

package types

import (
	fmt "fmt"
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

type ClassData struct {
	Metadata JsonInput `protobuf:"bytes,1,opt,name=metadata,proto3,customtype=JsonInput" json:"metadata"`
}

func (m *ClassData) Reset()         { *m = ClassData{} }
func (m *ClassData) String() string { return proto.CompactTextString(m) }
func (*ClassData) ProtoMessage()    {}
func (*ClassData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a155c220e0171849, []int{0}
}
func (m *ClassData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClassData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClassData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClassData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassData.Merge(m, src)
}
func (m *ClassData) XXX_Size() int {
	return m.Size()
}
func (m *ClassData) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassData.DiscardUnknown(m)
}

var xxx_messageInfo_ClassData proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ClassData)(nil), "vcoa.vnft.ClassData")
}

func init() { proto.RegisterFile("vcoa/vnft/class_data.proto", fileDescriptor_a155c220e0171849) }

var fileDescriptor_a155c220e0171849 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x4b, 0xce, 0x4f,
	0xd4, 0x2f, 0xcb, 0x4b, 0x2b, 0xd1, 0x4f, 0xce, 0x49, 0x2c, 0x2e, 0x8e, 0x4f, 0x49, 0x2c, 0x49,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0xc9, 0xe9, 0x81, 0xe4, 0xa4, 0x44, 0xd2,
	0xf3, 0xd3, 0xf3, 0xc1, 0xa2, 0xfa, 0x20, 0x16, 0x44, 0x81, 0x94, 0x18, 0x42, 0x73, 0x41, 0x62,
	0x51, 0x62, 0x6e, 0x31, 0x44, 0x5c, 0xc9, 0x8a, 0x8b, 0xd3, 0x19, 0x64, 0x98, 0x4b, 0x62, 0x49,
	0xa2, 0x90, 0x2e, 0x17, 0x47, 0x6e, 0x6a, 0x49, 0x22, 0xc8, 0x5c, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0x1e, 0x27, 0xc1, 0x13, 0xf7, 0xe4, 0x19, 0x6e, 0xdd, 0x93, 0xe7, 0xf4, 0x2a, 0xce, 0xcf, 0xf3,
	0xcc, 0x2b, 0x28, 0x2d, 0x09, 0x82, 0x2b, 0x71, 0xd2, 0x3e, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23,
	0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6,
	0x63, 0x39, 0x86, 0x28, 0x41, 0xb0, 0x6d, 0x15, 0x10, 0xfb, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93,
	0xd8, 0xc0, 0xf6, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xba, 0x61, 0x5e, 0xc6, 0x00,
	0x00, 0x00,
}

func (m *ClassData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClassData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Metadata.Size()
		i -= size
		if _, err := m.Metadata.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintClassData(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintClassData(dAtA []byte, offset int, v uint64) int {
	offset -= sovClassData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClassData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Metadata.Size()
	n += 1 + l + sovClassData(uint64(l))
	return n
}

func sovClassData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClassData(x uint64) (n int) {
	return sovClassData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClassData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClassData
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
			return fmt.Errorf("proto: ClassData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthClassData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthClassData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClassData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClassData
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
func skipClassData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClassData
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
					return 0, ErrIntOverflowClassData
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
					return 0, ErrIntOverflowClassData
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
				return 0, ErrInvalidLengthClassData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClassData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClassData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClassData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClassData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClassData = fmt.Errorf("proto: unexpected end of group")
)
