// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/liquidity/v1beta1/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/duration"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the parameters for the liquidity module.
type Params struct {
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_babec35f52b1356c, []int{0}
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

// Params defines the parameters for the liquidity module.
type GenericParams struct {
	BatchSize                uint64                                   `protobuf:"varint,1,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	TickPrecision            uint64                                   `protobuf:"varint,2,opt,name=tick_precision,json=tickPrecision,proto3" json:"tick_precision,omitempty"`
	FeeCollectorAddress      string                                   `protobuf:"bytes,3,opt,name=fee_collector_address,json=feeCollectorAddress,proto3" json:"fee_collector_address,omitempty"`
	DustCollectorAddress     string                                   `protobuf:"bytes,4,opt,name=dust_collector_address,json=dustCollectorAddress,proto3" json:"dust_collector_address,omitempty"`
	MinInitialPoolCoinSupply github_com_cosmos_cosmos_sdk_types.Int   `protobuf:"bytes,5,opt,name=min_initial_pool_coin_supply,json=minInitialPoolCoinSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_initial_pool_coin_supply"`
	PairCreationFee          github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=pair_creation_fee,json=pairCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"pair_creation_fee"`
	PoolCreationFee          github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,7,rep,name=pool_creation_fee,json=poolCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"pool_creation_fee"`
	MinInitialDepositAmount  github_com_cosmos_cosmos_sdk_types.Int   `protobuf:"bytes,8,opt,name=min_initial_deposit_amount,json=minInitialDepositAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_initial_deposit_amount"`
	MaxPriceLimitRatio       github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,9,opt,name=max_price_limit_ratio,json=maxPriceLimitRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_price_limit_ratio"`
	MaxOrderLifespan         time.Duration                            `protobuf:"bytes,10,opt,name=max_order_lifespan,json=maxOrderLifespan,proto3,stdduration" json:"max_order_lifespan"`
	SwapFeeRate              github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,11,opt,name=swap_fee_rate,json=swapFeeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"swap_fee_rate"`
	WithdrawFeeRate          github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,12,opt,name=withdraw_fee_rate,json=withdrawFeeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"withdraw_fee_rate"`
	DepositExtraGas          github_com_cosmos_cosmos_sdk_types.Gas   `protobuf:"varint,13,opt,name=deposit_extra_gas,json=depositExtraGas,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Gas" json:"deposit_extra_gas"`
	WithdrawExtraGas         github_com_cosmos_cosmos_sdk_types.Gas   `protobuf:"varint,14,opt,name=withdraw_extra_gas,json=withdrawExtraGas,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Gas" json:"withdraw_extra_gas"`
	OrderExtraGas            github_com_cosmos_cosmos_sdk_types.Gas   `protobuf:"varint,15,opt,name=order_extra_gas,json=orderExtraGas,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Gas" json:"order_extra_gas"`
	SwapFeeDistrDenom        string                                   `protobuf:"bytes,16,opt,name=swap_fee_distr_denom,json=swapFeeDistrDenom,proto3" json:"swap_fee_distr_denom,omitempty"`
	SwapFeeBurnRate          github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,17,opt,name=swap_fee_burn_rate,json=swapFeeBurnRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"swap_fee_burn_rate"`
	AppId                    uint64                                   `protobuf:"varint,18,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (m *GenericParams) Reset()         { *m = GenericParams{} }
func (m *GenericParams) String() string { return proto.CompactTextString(m) }
func (*GenericParams) ProtoMessage()    {}
func (*GenericParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_babec35f52b1356c, []int{1}
}
func (m *GenericParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenericParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenericParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenericParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericParams.Merge(m, src)
}
func (m *GenericParams) XXX_Size() int {
	return m.Size()
}
func (m *GenericParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericParams.DiscardUnknown(m)
}

var xxx_messageInfo_GenericParams proto.InternalMessageInfo

var fileDescriptor_babec35f52b1356c = []byte{
	// 739 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xc1, 0x6e, 0xeb, 0x44,
	0x14, 0x86, 0x63, 0xee, 0xbd, 0xa1, 0x9d, 0x90, 0x9b, 0x66, 0x68, 0xc1, 0x54, 0xe0, 0x44, 0x57,
	0x2a, 0xca, 0xa6, 0x36, 0x6d, 0x79, 0x81, 0xa6, 0xa1, 0x55, 0xa5, 0x4a, 0xa4, 0xae, 0xc4, 0xa2,
	0x20, 0x8d, 0x26, 0xf6, 0x49, 0x3a, 0xaa, 0xed, 0x19, 0x66, 0xc6, 0x24, 0xed, 0x53, 0xb0, 0xe4,
	0x19, 0xe0, 0x45, 0xba, 0xec, 0x12, 0xb1, 0x68, 0xa1, 0x7d, 0x11, 0x34, 0x63, 0x3b, 0x09, 0x82,
	0x05, 0x8d, 0x74, 0x57, 0x89, 0xe7, 0x9c, 0xff, 0xfb, 0xc7, 0xe7, 0xf7, 0xd8, 0x68, 0x27, 0xe2,
	0x69, 0x0c, 0xb3, 0x20, 0x61, 0x3f, 0xe6, 0x2c, 0x66, 0xfa, 0x26, 0xf8, 0x69, 0x6f, 0x04, 0x9a,
	0xee, 0x05, 0x82, 0x4a, 0x9a, 0x2a, 0x5f, 0x48, 0xae, 0x39, 0x76, 0x8b, 0x36, 0x7f, 0xde, 0xe6,
	0x97, 0x6d, 0xdb, 0x9b, 0x13, 0x3e, 0xe1, 0xb6, 0x29, 0x30, 0xff, 0x8a, 0xfe, 0x6d, 0x2f, 0xe2,
	0x2a, 0xe5, 0x2a, 0x18, 0x51, 0x05, 0x73, 0x62, 0xc4, 0x59, 0x56, 0xd5, 0x27, 0x9c, 0x4f, 0x12,
	0x08, 0xec, 0xd5, 0x28, 0x1f, 0x07, 0x71, 0x2e, 0xa9, 0x66, 0xbc, 0xac, 0xbf, 0x5b, 0x43, 0xf5,
	0xa1, 0xf5, 0x7f, 0xf7, 0x5b, 0x03, 0x35, 0x4f, 0x20, 0x03, 0xc9, 0xa2, 0x62, 0x05, 0x7f, 0x81,
	0xd0, 0x88, 0xea, 0xe8, 0x8a, 0x28, 0x76, 0x0b, 0xae, 0xd3, 0x75, 0x7a, 0xaf, 0xc3, 0x75, 0xbb,
	0x72, 0xc1, 0x6e, 0x01, 0xef, 0xa0, 0xb7, 0x9a, 0x45, 0xd7, 0x44, 0x48, 0x88, 0x98, 0x62, 0x3c,
	0x73, 0x3f, 0xb0, 0x2d, 0x4d, 0xb3, 0x3a, 0xac, 0x16, 0xf1, 0x3e, 0xda, 0x1a, 0x03, 0x90, 0x88,
	0x27, 0x09, 0x44, 0x9a, 0x4b, 0x42, 0xe3, 0x58, 0x82, 0x52, 0xee, 0xab, 0xae, 0xd3, 0x5b, 0x0f,
	0x3f, 0x1e, 0x03, 0x1c, 0x55, 0xb5, 0xc3, 0xa2, 0x84, 0xbf, 0x46, 0x9f, 0xc4, 0xb9, 0xd2, 0xff,
	0x21, 0x7a, 0x6d, 0x45, 0x9b, 0xa6, 0xfa, 0x2f, 0x55, 0x86, 0x3e, 0x4f, 0x59, 0x46, 0x58, 0xc6,
	0x34, 0xa3, 0x09, 0x11, 0x9c, 0x27, 0xc4, 0x8c, 0x82, 0xa8, 0x5c, 0x88, 0xe4, 0xc6, 0x7d, 0x63,
	0xb4, 0x7d, 0xff, 0xee, 0xa1, 0x53, 0xfb, 0xe3, 0xa1, 0xf3, 0xe5, 0x84, 0xe9, 0xab, 0x7c, 0xe4,
	0x47, 0x3c, 0x0d, 0xca, 0x21, 0x16, 0x3f, 0xbb, 0x2a, 0xbe, 0x0e, 0xf4, 0x8d, 0x00, 0xe5, 0x9f,
	0x66, 0x3a, 0x74, 0x53, 0x96, 0x9d, 0x16, 0xc8, 0x21, 0xe7, 0xc9, 0x11, 0x67, 0xd9, 0x85, 0xe5,
	0xe1, 0x29, 0x6a, 0x0b, 0xca, 0x24, 0x89, 0x24, 0xd8, 0x91, 0x92, 0x31, 0x80, 0x5b, 0xef, 0xbe,
	0xea, 0x35, 0xf6, 0x3f, 0xf3, 0x0b, 0x96, 0x6f, 0x72, 0xa9, 0x22, 0xf4, 0x8d, 0xb6, 0xff, 0x95,
	0xf1, 0xff, 0xf5, 0xb1, 0xd3, 0xfb, 0x1f, 0xfe, 0x46, 0xa0, 0xc2, 0x96, 0x71, 0x39, 0x2a, 0x4d,
	0x8e, 0x01, 0xac, 0xb1, 0xbd, 0xb9, 0x65, 0xe3, 0x0f, 0xdf, 0x87, 0xb1, 0xb9, 0xe1, 0x25, 0xe3,
	0x6b, 0xb4, 0xbd, 0x3c, 0xe1, 0x18, 0x04, 0x57, 0x4c, 0x13, 0x9a, 0xf2, 0x3c, 0xd3, 0xee, 0xda,
	0x4a, 0xf3, 0xfd, 0x74, 0x31, 0xdf, 0x41, 0xc1, 0x3b, 0xb4, 0x38, 0x4c, 0xd1, 0x56, 0x4a, 0x67,
	0x44, 0x48, 0x16, 0x01, 0x49, 0x58, 0xca, 0x34, 0xb1, 0x8f, 0xae, 0xbb, 0xfe, 0x62, 0x9f, 0x01,
	0x44, 0x21, 0x4e, 0xe9, 0x6c, 0x68, 0x58, 0x67, 0x06, 0x15, 0x1a, 0x12, 0x3e, 0x47, 0x66, 0x95,
	0x70, 0x19, 0x83, 0x24, 0x09, 0x1b, 0x83, 0x12, 0x34, 0x73, 0x51, 0xd7, 0xb1, 0x93, 0x2c, 0x8e,
	0x8e, 0x5f, 0x1d, 0x1d, 0x7f, 0x50, 0x1e, 0x9d, 0xfe, 0x9a, 0xb1, 0xfe, 0xe5, 0xb1, 0xe3, 0x84,
	0x1b, 0x29, 0x9d, 0x7d, 0x6b, 0xd4, 0x67, 0xa5, 0x18, 0x87, 0xa8, 0xa9, 0xa6, 0x54, 0x98, 0x48,
	0xcc, 0x76, 0xc1, 0x6d, 0xac, 0xb4, 0xdb, 0x86, 0x81, 0x1c, 0x03, 0x84, 0x54, 0x03, 0xbe, 0x44,
	0xed, 0x29, 0xd3, 0x57, 0xb1, 0xa4, 0xd3, 0x05, 0xf7, 0xa3, 0x95, 0xb8, 0xad, 0x0a, 0xb4, 0xc4,
	0xae, 0x62, 0x84, 0x99, 0x96, 0x94, 0x4c, 0xa8, 0x72, 0x9b, 0xe6, 0x20, 0xbf, 0x88, 0x7d, 0x42,
	0x55, 0xd8, 0x2a, 0x41, 0xdf, 0x18, 0xce, 0x09, 0x55, 0xf8, 0x07, 0x84, 0xe7, 0xfb, 0x5e, 0xc0,
	0xdf, 0xae, 0x04, 0xdf, 0xa8, 0x48, 0x73, 0xfa, 0x77, 0xa8, 0x55, 0x04, 0xb7, 0x40, 0xb7, 0x56,
	0x42, 0x37, 0x2d, 0x66, 0xce, 0x0d, 0xd0, 0xe6, 0x3c, 0xc1, 0x98, 0x29, 0x2d, 0x49, 0x0c, 0x19,
	0x4f, 0xdd, 0x0d, 0xfb, 0xea, 0x69, 0x97, 0xc1, 0x0c, 0x4c, 0x65, 0x60, 0x0a, 0xf8, 0x7b, 0x84,
	0xe7, 0x82, 0x51, 0x2e, 0xb3, 0x22, 0x9f, 0xf6, 0x6a, 0xf9, 0x94, 0xf8, 0x7e, 0x2e, 0x33, 0x9b,
	0xcf, 0x16, 0xaa, 0x53, 0x21, 0x08, 0x8b, 0x5d, 0x6c, 0xdf, 0xae, 0x6f, 0xa8, 0x10, 0xa7, 0x71,
	0xff, 0xfc, 0xee, 0x2f, 0xaf, 0x76, 0xf7, 0xe4, 0x39, 0xf7, 0x4f, 0x9e, 0xf3, 0xe7, 0x93, 0xe7,
	0xfc, 0xfc, 0xec, 0xd5, 0xee, 0x9f, 0xbd, 0xda, 0xef, 0xcf, 0x5e, 0xed, 0xf2, 0xe0, 0x1f, 0x6e,
	0xe6, 0x83, 0xb2, 0xcb, 0xc7, 0x63, 0x16, 0x31, 0x9a, 0x94, 0xd7, 0xc1, 0xf2, 0x97, 0xc8, 0xda,
	0x8f, 0xea, 0xf6, 0x41, 0x3f, 0xf8, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x61, 0x18, 0xba, 0xbf, 0xaa,
	0x06, 0x00, 0x00,
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
	return len(dAtA) - i, nil
}

func (m *GenericParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenericParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenericParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AppId != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.AppId))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x90
	}
	{
		size := m.SwapFeeBurnRate.Size()
		i -= size
		if _, err := m.SwapFeeBurnRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x8a
	if len(m.SwapFeeDistrDenom) > 0 {
		i -= len(m.SwapFeeDistrDenom)
		copy(dAtA[i:], m.SwapFeeDistrDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.SwapFeeDistrDenom)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if m.OrderExtraGas != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.OrderExtraGas))
		i--
		dAtA[i] = 0x78
	}
	if m.WithdrawExtraGas != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.WithdrawExtraGas))
		i--
		dAtA[i] = 0x70
	}
	if m.DepositExtraGas != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DepositExtraGas))
		i--
		dAtA[i] = 0x68
	}
	{
		size := m.WithdrawFeeRate.Size()
		i -= size
		if _, err := m.WithdrawFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size := m.SwapFeeRate.Size()
		i -= size
		if _, err := m.SwapFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.MaxOrderLifespan, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.MaxOrderLifespan):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x52
	{
		size := m.MaxPriceLimitRatio.Size()
		i -= size
		if _, err := m.MaxPriceLimitRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.MinInitialDepositAmount.Size()
		i -= size
		if _, err := m.MinInitialDepositAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if len(m.PoolCreationFee) > 0 {
		for iNdEx := len(m.PoolCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.PairCreationFee) > 0 {
		for iNdEx := len(m.PairCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PairCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	{
		size := m.MinInitialPoolCoinSupply.Size()
		i -= size
		if _, err := m.MinInitialPoolCoinSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.DustCollectorAddress) > 0 {
		i -= len(m.DustCollectorAddress)
		copy(dAtA[i:], m.DustCollectorAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DustCollectorAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.FeeCollectorAddress) > 0 {
		i -= len(m.FeeCollectorAddress)
		copy(dAtA[i:], m.FeeCollectorAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.FeeCollectorAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.TickPrecision != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TickPrecision))
		i--
		dAtA[i] = 0x10
	}
	if m.BatchSize != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BatchSize))
		i--
		dAtA[i] = 0x8
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
	return n
}

func (m *GenericParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BatchSize != 0 {
		n += 1 + sovParams(uint64(m.BatchSize))
	}
	if m.TickPrecision != 0 {
		n += 1 + sovParams(uint64(m.TickPrecision))
	}
	l = len(m.FeeCollectorAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.DustCollectorAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.MinInitialPoolCoinSupply.Size()
	n += 1 + l + sovParams(uint64(l))
	if len(m.PairCreationFee) > 0 {
		for _, e := range m.PairCreationFee {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.PoolCreationFee) > 0 {
		for _, e := range m.PoolCreationFee {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = m.MinInitialDepositAmount.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxPriceLimitRatio.Size()
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.MaxOrderLifespan)
	n += 1 + l + sovParams(uint64(l))
	l = m.SwapFeeRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.WithdrawFeeRate.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.DepositExtraGas != 0 {
		n += 1 + sovParams(uint64(m.DepositExtraGas))
	}
	if m.WithdrawExtraGas != 0 {
		n += 1 + sovParams(uint64(m.WithdrawExtraGas))
	}
	if m.OrderExtraGas != 0 {
		n += 1 + sovParams(uint64(m.OrderExtraGas))
	}
	l = len(m.SwapFeeDistrDenom)
	if l > 0 {
		n += 2 + l + sovParams(uint64(l))
	}
	l = m.SwapFeeBurnRate.Size()
	n += 2 + l + sovParams(uint64(l))
	if m.AppId != 0 {
		n += 2 + sovParams(uint64(m.AppId))
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
func (m *GenericParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GenericParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenericParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchSize", wireType)
			}
			m.BatchSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BatchSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickPrecision", wireType)
			}
			m.TickPrecision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TickPrecision |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeCollectorAddress", wireType)
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
			m.FeeCollectorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DustCollectorAddress", wireType)
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
			m.DustCollectorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinInitialPoolCoinSupply", wireType)
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
			if err := m.MinInitialPoolCoinSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PairCreationFee = append(m.PairCreationFee, types.Coin{})
			if err := m.PairCreationFee[len(m.PairCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolCreationFee = append(m.PoolCreationFee, types.Coin{})
			if err := m.PoolCreationFee[len(m.PoolCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinInitialDepositAmount", wireType)
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
			if err := m.MinInitialDepositAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPriceLimitRatio", wireType)
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
			if err := m.MaxPriceLimitRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxOrderLifespan", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.MaxOrderLifespan, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFeeRate", wireType)
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
			if err := m.SwapFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawFeeRate", wireType)
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
			if err := m.WithdrawFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositExtraGas", wireType)
			}
			m.DepositExtraGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DepositExtraGas |= github_com_cosmos_cosmos_sdk_types.Gas(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawExtraGas", wireType)
			}
			m.WithdrawExtraGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WithdrawExtraGas |= github_com_cosmos_cosmos_sdk_types.Gas(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderExtraGas", wireType)
			}
			m.OrderExtraGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderExtraGas |= github_com_cosmos_cosmos_sdk_types.Gas(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFeeDistrDenom", wireType)
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
			m.SwapFeeDistrDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFeeBurnRate", wireType)
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
			if err := m.SwapFeeBurnRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 18:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppId", wireType)
			}
			m.AppId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppId |= uint64(b&0x7F) << shift
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