// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nitro/data.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type TransactionData struct {
	Slot            uint64   `protobuf:"varint,1,opt,name=slot,proto3" json:"slot"`
	Signature       string   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature"`
	IsVote          bool     `protobuf:"varint,3,opt,name=is_vote,json=isVote,proto3" json:"is_vote"`
	MessageType     uint64   `protobuf:"varint,4,opt,name=message_type,json=messageType,proto3" json:"message_type"`
	LegacyMessage   string   `protobuf:"bytes,5,opt,name=legacy_message,json=legacyMessage,proto3" json:"legacy_message"`
	V0LoadedMessage string   `protobuf:"bytes,6,opt,name=v0_loaded_message,json=v0LoadedMessage,proto3" json:"v0_loaded_message"`
	Signatures      []string `protobuf:"bytes,7,rep,name=signatures,proto3" json:"signatures"`
	MessageHash     string   `protobuf:"bytes,8,opt,name=message_hash,json=messageHash,proto3" json:"message_hash"`
	Meta            string   `protobuf:"bytes,9,opt,name=meta,proto3" json:"string"`
	WriteVersion    uint64   `protobuf:"varint,10,opt,name=write_version,json=writeVersion,proto3" json:"write_version"`
}

func (m *TransactionData) Reset()         { *m = TransactionData{} }
func (m *TransactionData) String() string { return proto.CompactTextString(m) }
func (*TransactionData) ProtoMessage()    {}
func (*TransactionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7cd900d11b0c050, []int{0}
}
func (m *TransactionData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransactionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionData.Merge(m, src)
}
func (m *TransactionData) XXX_Size() int {
	return m.Size()
}
func (m *TransactionData) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionData.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionData proto.InternalMessageInfo

func (m *TransactionData) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *TransactionData) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *TransactionData) GetIsVote() bool {
	if m != nil {
		return m.IsVote
	}
	return false
}

func (m *TransactionData) GetMessageType() uint64 {
	if m != nil {
		return m.MessageType
	}
	return 0
}

func (m *TransactionData) GetLegacyMessage() string {
	if m != nil {
		return m.LegacyMessage
	}
	return ""
}

func (m *TransactionData) GetV0LoadedMessage() string {
	if m != nil {
		return m.V0LoadedMessage
	}
	return ""
}

func (m *TransactionData) GetSignatures() []string {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *TransactionData) GetMessageHash() string {
	if m != nil {
		return m.MessageHash
	}
	return ""
}

func (m *TransactionData) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

func (m *TransactionData) GetWriteVersion() uint64 {
	if m != nil {
		return m.WriteVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*TransactionData)(nil), "seiprotocol.seichain.nitro.TransactionData")
}

func init() { proto.RegisterFile("nitro/data.proto", fileDescriptor_b7cd900d11b0c050) }

var fileDescriptor_b7cd900d11b0c050 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcf, 0x8a, 0xdb, 0x30,
	0x10, 0xc6, 0xe3, 0x26, 0xcd, 0x26, 0xda, 0xcd, 0xfe, 0x11, 0x2d, 0x88, 0xa5, 0x58, 0xa1, 0xf4,
	0x10, 0x28, 0xb5, 0x17, 0x16, 0x0a, 0x3d, 0xd6, 0xf4, 0xd0, 0x42, 0x7b, 0x11, 0xcb, 0x1e, 0x7a,
	0x31, 0x5a, 0x47, 0xd8, 0x02, 0xc7, 0x0a, 0x1e, 0xad, 0xdb, 0xbc, 0x45, 0x1f, 0xab, 0xc7, 0x3d,
	0xf6, 0x24, 0x4a, 0x72, 0xd3, 0xb1, 0x4f, 0x50, 0x32, 0xde, 0xfc, 0x6b, 0x2f, 0x66, 0xbe, 0xdf,
	0x37, 0x63, 0xf9, 0xb3, 0x86, 0x9c, 0x57, 0xda, 0xd6, 0x26, 0x9e, 0x4a, 0x2b, 0xa3, 0x79, 0x6d,
	0xac, 0xa1, 0x97, 0xa0, 0x34, 0x56, 0x99, 0x29, 0x23, 0x50, 0x3a, 0x2b, 0xa4, 0xae, 0x22, 0x6c,
	0xbb, 0x7c, 0x96, 0x9b, 0xdc, 0xa0, 0x19, 0xaf, 0xab, 0x76, 0xe2, 0xe5, 0x9f, 0x2e, 0x39, 0xbb,
	0xa9, 0x65, 0x05, 0x32, 0xb3, 0xda, 0x54, 0x1f, 0xa4, 0x95, 0xf4, 0x05, 0xe9, 0x41, 0x69, 0x2c,
	0x0b, 0xc6, 0xc1, 0xa4, 0x97, 0x0c, 0xbc, 0xe3, 0xa8, 0x05, 0x3e, 0xe9, 0x6b, 0x32, 0x04, 0x9d,
	0x57, 0xd2, 0xde, 0xd7, 0x8a, 0x3d, 0x19, 0x07, 0x93, 0x61, 0x32, 0xf2, 0x8e, 0xef, 0xa0, 0xd8,
	0x95, 0xf4, 0x15, 0x39, 0xd2, 0x90, 0x36, 0xc6, 0x2a, 0xd6, 0x1d, 0x07, 0x93, 0x41, 0x72, 0xec,
	0x1d, 0xdf, 0x20, 0xd1, 0xd7, 0x70, 0x6b, 0xac, 0xa2, 0xd7, 0xe4, 0x64, 0xa6, 0x00, 0x64, 0xae,
	0x52, 0xbb, 0x98, 0x2b, 0xd6, 0xc3, 0x83, 0xcf, 0xbd, 0xe3, 0x07, 0x5c, 0x1c, 0x3f, 0xaa, 0x9b,
	0xc5, 0x5c, 0xd1, 0x77, 0xe4, 0xb4, 0x54, 0xb9, 0xcc, 0x16, 0xe9, 0x23, 0x65, 0x4f, 0xf1, 0x63,
	0xa8, 0x77, 0xfc, 0x1f, 0x47, 0x8c, 0x5a, 0xfd, 0xa5, 0x95, 0xf4, 0x3d, 0xb9, 0x68, 0xae, 0xd2,
	0xd2, 0xc8, 0xa9, 0x9a, 0x6e, 0xa7, 0xfb, 0x38, 0xfd, 0xdc, 0x3b, 0xfe, 0xbf, 0x29, 0xce, 0x9a,
	0xab, 0xcf, 0x48, 0x36, 0xaf, 0x88, 0x08, 0xd9, 0xa6, 0x04, 0x76, 0x34, 0xee, 0x4e, 0x86, 0xc9,
	0xa9, 0x77, 0x7c, 0x8f, 0x8a, 0xbd, 0x7a, 0x3f, 0x62, 0x21, 0xa1, 0x60, 0x03, 0x3c, 0xed, 0x20,
	0xe2, 0x9a, 0x6f, 0x23, 0x7e, 0x94, 0x50, 0xd0, 0x90, 0xf4, 0x66, 0xca, 0x4a, 0x36, 0xc4, 0x66,
	0xe2, 0x1d, 0xef, 0x83, 0xad, 0x75, 0x95, 0x0b, 0xe4, 0xf4, 0x2d, 0x19, 0x7d, 0xab, 0xb5, 0x55,
	0x69, 0xa3, 0x6a, 0xd0, 0xa6, 0x62, 0x04, 0x7f, 0xdc, 0x85, 0x77, 0xfc, 0xd0, 0x10, 0x27, 0x28,
	0x6f, 0x5b, 0x95, 0x7c, 0xfa, 0xb9, 0x0c, 0x83, 0x87, 0x65, 0x18, 0xfc, 0x5e, 0x86, 0xc1, 0x8f,
	0x55, 0xd8, 0x79, 0x58, 0x85, 0x9d, 0x5f, 0xab, 0xb0, 0xf3, 0x35, 0xce, 0xb5, 0x2d, 0xee, 0xef,
	0xa2, 0xcc, 0xcc, 0x62, 0x50, 0xfa, 0xcd, 0x66, 0x99, 0x50, 0xe0, 0x36, 0xc5, 0xdf, 0xe3, 0x76,
	0xed, 0xd6, 0x37, 0x02, 0x77, 0x7d, 0xec, 0xb8, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xba, 0x78,
	0x13, 0x5a, 0x8c, 0x02, 0x00, 0x00,
}

func (m *TransactionData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransactionData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.WriteVersion != 0 {
		i = encodeVarintData(dAtA, i, uint64(m.WriteVersion))
		i--
		dAtA[i] = 0x50
	}
	if len(m.Meta) > 0 {
		i -= len(m.Meta)
		copy(dAtA[i:], m.Meta)
		i = encodeVarintData(dAtA, i, uint64(len(m.Meta)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.MessageHash) > 0 {
		i -= len(m.MessageHash)
		copy(dAtA[i:], m.MessageHash)
		i = encodeVarintData(dAtA, i, uint64(len(m.MessageHash)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Signatures) > 0 {
		for iNdEx := len(m.Signatures) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signatures[iNdEx])
			copy(dAtA[i:], m.Signatures[iNdEx])
			i = encodeVarintData(dAtA, i, uint64(len(m.Signatures[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.V0LoadedMessage) > 0 {
		i -= len(m.V0LoadedMessage)
		copy(dAtA[i:], m.V0LoadedMessage)
		i = encodeVarintData(dAtA, i, uint64(len(m.V0LoadedMessage)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.LegacyMessage) > 0 {
		i -= len(m.LegacyMessage)
		copy(dAtA[i:], m.LegacyMessage)
		i = encodeVarintData(dAtA, i, uint64(len(m.LegacyMessage)))
		i--
		dAtA[i] = 0x2a
	}
	if m.MessageType != 0 {
		i = encodeVarintData(dAtA, i, uint64(m.MessageType))
		i--
		dAtA[i] = 0x20
	}
	if m.IsVote {
		i--
		if m.IsVote {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintData(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x12
	}
	if m.Slot != 0 {
		i = encodeVarintData(dAtA, i, uint64(m.Slot))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintData(dAtA []byte, offset int, v uint64) int {
	offset -= sovData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TransactionData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Slot != 0 {
		n += 1 + sovData(uint64(m.Slot))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovData(uint64(l))
	}
	if m.IsVote {
		n += 2
	}
	if m.MessageType != 0 {
		n += 1 + sovData(uint64(m.MessageType))
	}
	l = len(m.LegacyMessage)
	if l > 0 {
		n += 1 + l + sovData(uint64(l))
	}
	l = len(m.V0LoadedMessage)
	if l > 0 {
		n += 1 + l + sovData(uint64(l))
	}
	if len(m.Signatures) > 0 {
		for _, s := range m.Signatures {
			l = len(s)
			n += 1 + l + sovData(uint64(l))
		}
	}
	l = len(m.MessageHash)
	if l > 0 {
		n += 1 + l + sovData(uint64(l))
	}
	l = len(m.Meta)
	if l > 0 {
		n += 1 + l + sovData(uint64(l))
	}
	if m.WriteVersion != 0 {
		n += 1 + sovData(uint64(m.WriteVersion))
	}
	return n
}

func sovData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozData(x uint64) (n int) {
	return sovData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TransactionData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowData
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
			return fmt.Errorf("proto: TransactionData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slot", wireType)
			}
			m.Slot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Slot |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsVote", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
			m.IsVote = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageType", wireType)
			}
			m.MessageType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MessageType |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LegacyMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LegacyMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field V0LoadedMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.V0LoadedMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signatures = append(m.Signatures, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Meta = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WriteVersion", wireType)
			}
			m.WriteVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WriteVersion |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthData
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
func skipData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowData
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
					return 0, ErrIntOverflowData
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
					return 0, ErrIntOverflowData
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
				return 0, ErrInvalidLengthData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupData = fmt.Errorf("proto: unexpected end of group")
)