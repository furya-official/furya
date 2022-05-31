// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/auction/v1beta1/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Params struct {
	LiquidationPenaltyPercent       string                                 `protobuf:"bytes,1,opt,name=liquidation_penalty_percent,json=liquidationPenaltyPercent,proto3" json:"liquidation_penalty_percent,omitempty" yaml:"liquidation_penalty_percent"`
	AuctionDiscountPercent          string                                 `protobuf:"bytes,2,opt,name=auction_discount_percent,json=auctionDiscountPercent,proto3" json:"auction_discount_percent,omitempty" yaml:"auction_discount_percent"`
	AuctionDurationSeconds          uint64                                 `protobuf:"varint,3,opt,name=auction_duration_seconds,json=auctionDurationSeconds,proto3" json:"auction_duration_seconds,omitempty" yaml:"auction_duration_seconds"`
	DebtMintTokenDecreasePercentage github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=debt_mint_token_decrease_percentage,json=debtMintTokenDecreasePercentage,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"debt_mint_token_decrease_percentage" yaml:"debt_mint_token_decrease_percentage"`
	Buffer                          github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=buffer,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"buffer" yaml:"buffer"`
	Cusp                            github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=cusp,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"cusp" yaml:"cusp"`
	Tau                             github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=tau,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"tau" yaml:"tau"`
	DutchDecreasePercentage         github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=dutchDecreasePercentage,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"dutchDecreasePercentage" yaml:"decreasePercent"`
	Chost                           github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,9,opt,name=chost,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"chost" yaml:"chost"`
	Step                            github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,10,opt,name=step,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"step" yaml:"step"`
	PriceFunctionType               uint64                                 `protobuf:"varint,11,opt,name=priceFunctionType,proto3" json:"priceFunctionType,omitempty" yaml:"priceFunctionType"`
	SurplusId                       uint64                                 `protobuf:"varint,12,opt,name=surplus_id,json=surplusId,proto3" json:"surplus_id,omitempty" yaml:"priceFunctionType"`
	DebtId                          uint64                                 `protobuf:"varint,13,opt,name=debt_id,json=debtId,proto3" json:"debt_id,omitempty" yaml:"priceFunctionType"`
	DutchId                         uint64                                 `protobuf:"varint,14,opt,name=dutch_id,json=dutchId,proto3" json:"dutch_id,omitempty" yaml:"priceFunctionType"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_4370eec7f59a9a46, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetLiquidationPenaltyPercent() string {
	if m != nil {
		return m.LiquidationPenaltyPercent
	}
	return ""
}

func (m *Params) GetAuctionDiscountPercent() string {
	if m != nil {
		return m.AuctionDiscountPercent
	}
	return ""
}

func (m *Params) GetAuctionDurationSeconds() uint64 {
	if m != nil {
		return m.AuctionDurationSeconds
	}
	return 0
}

func (m *Params) GetPriceFunctionType() uint64 {
	if m != nil {
		return m.PriceFunctionType
	}
	return 0
}

func (m *Params) GetSurplusId() uint64 {
	if m != nil {
		return m.SurplusId
	}
	return 0
}

func (m *Params) GetDebtId() uint64 {
	if m != nil {
		return m.DebtId
	}
	return 0
}

func (m *Params) GetDutchId() uint64 {
	if m != nil {
		return m.DutchId
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "comdex.auction.v1beta1.Params")
}

func init() {
	proto.RegisterFile("comdex/auction/v1beta1/params.proto", fileDescriptor_4370eec7f59a9a46)
}

var fileDescriptor_4370eec7f59a9a46 = []byte{
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdd, 0x6a, 0x13, 0x4d,
	0x18, 0xc7, 0xb3, 0x6f, 0xdb, 0xb4, 0x99, 0xb6, 0x2f, 0x38, 0x48, 0x1d, 0x3f, 0xc8, 0x96, 0x29,
	0x94, 0x22, 0x34, 0x4b, 0x11, 0x11, 0xfc, 0x84, 0x52, 0xc4, 0x28, 0x4a, 0x5c, 0x0b, 0x82, 0x28,
	0xcb, 0x66, 0x66, 0x92, 0x0e, 0x4d, 0x76, 0xd6, 0x9d, 0x19, 0x31, 0xa7, 0x5e, 0x81, 0xd7, 0xe1,
	0x95, 0xf4, 0xb0, 0x87, 0xe2, 0xc1, 0x22, 0xed, 0x1d, 0xe4, 0x02, 0x44, 0xe6, 0xa3, 0xb5, 0x35,
	0xad, 0x64, 0x8f, 0x92, 0xf0, 0xfc, 0x9e, 0xdf, 0xf3, 0xe4, 0xbf, 0xb3, 0x03, 0xd6, 0x88, 0x18,
	0x52, 0xf6, 0x39, 0x4a, 0x35, 0x51, 0x5c, 0x64, 0xd1, 0xa7, 0xad, 0x2e, 0x53, 0xe9, 0x56, 0x94,
	0xa7, 0x45, 0x3a, 0x94, 0xad, 0xbc, 0x10, 0x4a, 0xc0, 0x15, 0x07, 0xb5, 0x3c, 0xd4, 0xf2, 0xd0,
	0x8d, 0xab, 0x7d, 0xd1, 0x17, 0x16, 0x89, 0xcc, 0x37, 0x47, 0xe3, 0x5f, 0x0d, 0x50, 0xef, 0xd8,
	0x76, 0xd8, 0x03, 0x37, 0x07, 0xfc, 0xa3, 0xe6, 0x34, 0x35, 0x7d, 0x49, 0xce, 0xb2, 0x74, 0xa0,
	0x46, 0x49, 0xce, 0x0a, 0xc2, 0x32, 0x85, 0x82, 0xd5, 0x60, 0xa3, 0xb1, 0xbd, 0x3e, 0x2e, 0x43,
	0x3c, 0x4a, 0x87, 0x83, 0xfb, 0xf8, 0x1f, 0x30, 0x8e, 0xaf, 0x9f, 0xa9, 0x76, 0x5c, 0xb1, 0xe3,
	0x6a, 0xf0, 0x03, 0x40, 0x7e, 0xb7, 0x84, 0x72, 0x49, 0x84, 0xce, 0xd4, 0xe9, 0x90, 0xff, 0xec,
	0x90, 0xb5, 0x71, 0x19, 0x86, 0x6e, 0xc8, 0x65, 0x24, 0x8e, 0x57, 0x7c, 0x69, 0xc7, 0x57, 0x2e,
	0xd2, 0xeb, 0xc2, 0xad, 0x27, 0x19, 0x11, 0x19, 0x95, 0x68, 0x66, 0x35, 0xd8, 0x98, 0xbd, 0x50,
	0xff, 0x17, 0x79, 0x46, 0xef, 0x2b, 0x6f, 0x5c, 0x01, 0x7e, 0x0b, 0xc0, 0x1a, 0x65, 0x5d, 0x95,
	0x0c, 0x79, 0xa6, 0x12, 0x25, 0xf6, 0x59, 0x96, 0x50, 0x46, 0x0a, 0x96, 0x4a, 0x76, 0xb2, 0x5c,
	0xda, 0x67, 0x68, 0xd6, 0xfe, 0x93, 0xf7, 0x07, 0x65, 0x58, 0xfb, 0x51, 0x86, 0xeb, 0x7d, 0xae,
	0xf6, 0x74, 0xb7, 0x45, 0xc4, 0x30, 0x22, 0x42, 0x0e, 0x85, 0xf4, 0x1f, 0x9b, 0x92, 0xee, 0x47,
	0x6a, 0x94, 0x33, 0xd9, 0xda, 0x61, 0x64, 0x5c, 0x86, 0xb7, 0xdd, 0x62, 0x53, 0x8c, 0xc0, 0x71,
	0x68, 0xa8, 0x97, 0x3c, 0x53, 0xbb, 0x86, 0xd9, 0xf1, 0x48, 0xe7, 0x94, 0x80, 0x6f, 0x41, 0xbd,
	0xab, 0x7b, 0x3d, 0x56, 0xa0, 0x39, 0xbb, 0xce, 0x93, 0xca, 0xeb, 0x2c, 0xbb, 0x75, 0x9c, 0x05,
	0xc7, 0x5e, 0x07, 0x5f, 0x83, 0x59, 0xa2, 0x65, 0x8e, 0xea, 0x56, 0xfb, 0xa8, 0xb2, 0x76, 0xd1,
	0x69, 0x8d, 0x03, 0xc7, 0x56, 0x05, 0x5f, 0x81, 0x19, 0x95, 0x6a, 0x34, 0x6f, 0x8d, 0x0f, 0x2b,
	0x18, 0xdb, 0x99, 0x1a, 0x97, 0x21, 0x70, 0x46, 0x95, 0x6a, 0x1c, 0x1b, 0x11, 0xfc, 0x12, 0x80,
	0x6b, 0x54, 0x2b, 0xb2, 0x37, 0x99, 0x0b, 0x5a, 0xb0, 0x43, 0x9e, 0x55, 0x5e, 0x7b, 0xe5, 0xe4,
	0xe1, 0x9c, 0x33, 0xe2, 0xf8, 0xb2, 0x41, 0x70, 0x17, 0xcc, 0x91, 0x3d, 0x21, 0x15, 0x6a, 0xd8,
	0x89, 0x8f, 0x2b, 0x4f, 0x5c, 0xf2, 0x41, 0x19, 0x09, 0x8e, 0x9d, 0xcc, 0xa4, 0x2f, 0x15, 0xcb,
	0x11, 0xa8, 0x9c, 0xbe, 0xcb, 0xca, 0xa7, 0x6f, 0x1c, 0x38, 0xb6, 0x2a, 0xf8, 0x1c, 0x5c, 0xc9,
	0x0b, 0x4e, 0xd8, 0x53, 0x9d, 0xd9, 0x63, 0xbf, 0x3b, 0xca, 0x19, 0x5a, 0xb4, 0xaf, 0xcb, 0xad,
	0x71, 0x19, 0x22, 0xd7, 0x31, 0x81, 0xe0, 0x78, 0xb2, 0x0d, 0x3e, 0x00, 0x40, 0xea, 0x22, 0x1f,
	0x68, 0x99, 0x70, 0x8a, 0x96, 0xa6, 0x90, 0x34, 0x3c, 0xdf, 0xa6, 0xf0, 0x2e, 0x98, 0xb7, 0x67,
	0x9f, 0x53, 0xb4, 0x3c, 0x45, 0x67, 0xdd, 0xc0, 0x6d, 0x0a, 0xef, 0x81, 0x05, 0xfb, 0x0c, 0x4c,
	0xdf, 0xff, 0x53, 0xf4, 0xcd, 0x5b, 0xba, 0x4d, 0xb7, 0x5f, 0x1c, 0x1c, 0x35, 0x83, 0xc3, 0xa3,
	0x66, 0xf0, 0xf3, 0xa8, 0x19, 0x7c, 0x3d, 0x6e, 0xd6, 0x0e, 0x8f, 0x9b, 0xb5, 0xef, 0xc7, 0xcd,
	0xda, 0xbb, 0xad, 0x73, 0x79, 0x9a, 0x3b, 0x75, 0x53, 0xf4, 0x7a, 0x9c, 0xf0, 0x74, 0xe0, 0x7f,
	0x47, 0x7f, 0xae, 0x62, 0x1b, 0x6f, 0xb7, 0x6e, 0x2f, 0xd5, 0x3b, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xc8, 0xe3, 0xe2, 0xdf, 0xa9, 0x05, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DutchId != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DutchId))
		i--
		dAtA[i] = 0x70
	}
	if m.DebtId != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DebtId))
		i--
		dAtA[i] = 0x68
	}
	if m.SurplusId != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SurplusId))
		i--
		dAtA[i] = 0x60
	}
	if m.PriceFunctionType != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.PriceFunctionType))
		i--
		dAtA[i] = 0x58
	}
	{
		size := m.Step.Size()
		i -= size
		if _, err := m.Step.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.Chost.Size()
		i -= size
		if _, err := m.Chost.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.DutchDecreasePercentage.Size()
		i -= size
		if _, err := m.DutchDecreasePercentage.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.Tau.Size()
		i -= size
		if _, err := m.Tau.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.Cusp.Size()
		i -= size
		if _, err := m.Cusp.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.Buffer.Size()
		i -= size
		if _, err := m.Buffer.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.DebtMintTokenDecreasePercentage.Size()
		i -= size
		if _, err := m.DebtMintTokenDecreasePercentage.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.AuctionDurationSeconds != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.AuctionDurationSeconds))
		i--
		dAtA[i] = 0x18
	}
	if len(m.AuctionDiscountPercent) > 0 {
		i -= len(m.AuctionDiscountPercent)
		copy(dAtA[i:], m.AuctionDiscountPercent)
		i = encodeVarintParams(dAtA, i, uint64(len(m.AuctionDiscountPercent)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.LiquidationPenaltyPercent) > 0 {
		i -= len(m.LiquidationPenaltyPercent)
		copy(dAtA[i:], m.LiquidationPenaltyPercent)
		i = encodeVarintParams(dAtA, i, uint64(len(m.LiquidationPenaltyPercent)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.LiquidationPenaltyPercent)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.AuctionDiscountPercent)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.AuctionDurationSeconds != 0 {
		n += 1 + sovParams(uint64(m.AuctionDurationSeconds))
	}
	l = m.DebtMintTokenDecreasePercentage.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.Buffer.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.Cusp.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.Tau.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.DutchDecreasePercentage.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.Chost.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.Step.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.PriceFunctionType != 0 {
		n += 1 + sovParams(uint64(m.PriceFunctionType))
	}
	if m.SurplusId != 0 {
		n += 1 + sovParams(uint64(m.SurplusId))
	}
	if m.DebtId != 0 {
		n += 1 + sovParams(uint64(m.DebtId))
	}
	if m.DutchId != 0 {
		n += 1 + sovParams(uint64(m.DutchId))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationPenaltyPercent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LiquidationPenaltyPercent = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionDiscountPercent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuctionDiscountPercent = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionDurationSeconds", wireType)
			}
			m.AuctionDurationSeconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionDurationSeconds |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebtMintTokenDecreasePercentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DebtMintTokenDecreasePercentage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Buffer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Buffer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cusp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Cusp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tau", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Tau.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DutchDecreasePercentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DutchDecreasePercentage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chost", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Chost.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Step", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Step.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceFunctionType", wireType)
			}
			m.PriceFunctionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PriceFunctionType |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SurplusId", wireType)
			}
			m.SurplusId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SurplusId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebtId", wireType)
			}
			m.DebtId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DebtId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DutchId", wireType)
			}
			m.DutchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DutchId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)