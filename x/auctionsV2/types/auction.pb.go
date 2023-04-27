// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/auctionsV2/v1beta1/auction.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type Auctions struct {
	AuctionId       uint64                                  `protobuf:"varint,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty" yaml:"auction_id"`
	CollateralToken github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,2,opt,name=collateral_token,json=collateralToken,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"collateral_token" yaml:"collateral_token"`
	// cosmos.base.v1beta1.Coin outflow_token_current_amount = 3 [
	//   (gogoproto.nullable) = false,
	//   (gogoproto.moretags) = "yaml:\"outflow_token_current_amount\"",
	//   (gogoproto.casttype) = "github.com/cosmos/cosmos-sdk/types.Coin"
	// ];
	DebtToken github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,4,opt,name=debt_token,json=debtToken,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"debt_token" yaml:"inflow_token_target_amount"`
	// cosmos.base.v1beta1.Coin inflow_token_current_amount = 5 [
	//   (gogoproto.nullable) = false,
	//   (gogoproto.moretags) = "yaml:\"inflow_token_current_amount\"",
	//   (gogoproto.casttype) = "github.com/cosmos/cosmos-sdk/types.Coin"
	// ];
	ActiveBiddingId uint64                                 `protobuf:"varint,6,opt,name=active_bidding_id,json=activeBiddingId,proto3" json:"active_bidding_id,omitempty" yaml:"active_bidding_id"`
	BiddingIds      []*BidOwnerMapping                     `protobuf:"bytes,7,rep,name=bidding_ids,json=biddingIds,proto3" json:"bidding_ids,omitempty" yaml:"bidding_ids"`
	BidFactor       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=bid_factor,json=bidFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"bid_factor" yaml:"bid_factor"`
	// price indicator only for dutch auctions
	CollateralTokenInitialPrice github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,9,opt,name=collateral_token_initial_price,json=collateralTokenInitialPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"collateral_token_initial_price" yaml:"outflow_token_initial_price"`
	CollateralTokenCurrentPrice github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,10,opt,name=collateral_token_current_price,json=collateralTokenCurrentPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"collateral_token_current_price" yaml:"outflow_token_current_price"`
	CollateralTokenEndPrice     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,11,opt,name=collateral_token_end_price,json=collateralTokenEndPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"collateral_token_end_price" yaml:"outflow_token_end_price"`
	DebtTokenCurrentPrice       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,12,opt,name=debt_token_current_price,json=debtTokenCurrentPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"debt_token_current_price" yaml:"inflow_token_current_price"`
	LockedVaultId               uint64                                 `protobuf:"varint,13,opt,name=locked_vault_id,json=lockedVaultId,proto3" json:"locked_vault_id,omitempty" yaml:"locked_vault_id"`
	StartTime                   time.Time                              `protobuf:"bytes,15,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
	EndTime                     time.Time                              `protobuf:"bytes,16,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time" yaml:"end_time"`
	AppId                       uint64                                 `protobuf:"varint,18,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty" yaml:"app_id"`
	AuctionType                 bool                                   `protobuf:"varint,19,opt,name=auction_type,json=auctionType,proto3" json:"auction_type,omitempty" yaml:"auction_type"`
}

func (m *Auctions) Reset()         { *m = Auctions{} }
func (m *Auctions) String() string { return proto.CompactTextString(m) }
func (*Auctions) ProtoMessage()    {}
func (*Auctions) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ee47f5a405fa8ba, []int{0}
}
func (m *Auctions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Auctions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Auctions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Auctions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auctions.Merge(m, src)
}
func (m *Auctions) XXX_Size() int {
	return m.Size()
}
func (m *Auctions) XXX_DiscardUnknown() {
	xxx_messageInfo_Auctions.DiscardUnknown(m)
}

var xxx_messageInfo_Auctions proto.InternalMessageInfo

func (m *Auctions) GetAuctionId() uint64 {
	if m != nil {
		return m.AuctionId
	}
	return 0
}

func (m *Auctions) GetCollateralToken() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.CollateralToken
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

func (m *Auctions) GetDebtToken() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.DebtToken
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

func (m *Auctions) GetActiveBiddingId() uint64 {
	if m != nil {
		return m.ActiveBiddingId
	}
	return 0
}

func (m *Auctions) GetBiddingIds() []*BidOwnerMapping {
	if m != nil {
		return m.BiddingIds
	}
	return nil
}

func (m *Auctions) GetLockedVaultId() uint64 {
	if m != nil {
		return m.LockedVaultId
	}
	return 0
}

func (m *Auctions) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *Auctions) GetEndTime() time.Time {
	if m != nil {
		return m.EndTime
	}
	return time.Time{}
}

func (m *Auctions) GetAppId() uint64 {
	if m != nil {
		return m.AppId
	}
	return 0
}

func (m *Auctions) GetAuctionType() bool {
	if m != nil {
		return m.AuctionType
	}
	return false
}

type BidOwnerMapping struct {
	BidId    uint64 `protobuf:"varint,1,opt,name=bid_id,json=bidId,proto3" json:"bid_id,omitempty"`
	BidOwner string `protobuf:"bytes,2,opt,name=bid_owner,json=bidOwner,proto3" json:"bid_owner,omitempty"`
}

func (m *BidOwnerMapping) Reset()         { *m = BidOwnerMapping{} }
func (m *BidOwnerMapping) String() string { return proto.CompactTextString(m) }
func (*BidOwnerMapping) ProtoMessage()    {}
func (*BidOwnerMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ee47f5a405fa8ba, []int{1}
}
func (m *BidOwnerMapping) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BidOwnerMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BidOwnerMapping.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BidOwnerMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidOwnerMapping.Merge(m, src)
}
func (m *BidOwnerMapping) XXX_Size() int {
	return m.Size()
}
func (m *BidOwnerMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_BidOwnerMapping.DiscardUnknown(m)
}

var xxx_messageInfo_BidOwnerMapping proto.InternalMessageInfo

func (m *BidOwnerMapping) GetBidId() uint64 {
	if m != nil {
		return m.BidId
	}
	return 0
}

func (m *BidOwnerMapping) GetBidOwner() string {
	if m != nil {
		return m.BidOwner
	}
	return ""
}

type AuctionParams struct {
	AuctionDurationSeconds uint64 `protobuf:"varint,1,opt,name=auction_duration_seconds,json=auctionDurationSeconds,proto3" json:"auction_duration_seconds,omitempty" yaml:"auction_duration_seconds"`
}

func (m *AuctionParams) Reset()         { *m = AuctionParams{} }
func (m *AuctionParams) String() string { return proto.CompactTextString(m) }
func (*AuctionParams) ProtoMessage()    {}
func (*AuctionParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ee47f5a405fa8ba, []int{2}
}
func (m *AuctionParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuctionParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuctionParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuctionParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuctionParams.Merge(m, src)
}
func (m *AuctionParams) XXX_Size() int {
	return m.Size()
}
func (m *AuctionParams) XXX_DiscardUnknown() {
	xxx_messageInfo_AuctionParams.DiscardUnknown(m)
}

var xxx_messageInfo_AuctionParams proto.InternalMessageInfo

func (m *AuctionParams) GetAuctionDurationSeconds() uint64 {
	if m != nil {
		return m.AuctionDurationSeconds
	}
	return 0
}

func init() {
	proto.RegisterType((*Auctions)(nil), "comdex.auctionsV2.v1beta1.Auctions")
	proto.RegisterType((*BidOwnerMapping)(nil), "comdex.auctionsV2.v1beta1.bidOwnerMapping")
	proto.RegisterType((*AuctionParams)(nil), "comdex.auctionsV2.v1beta1.AuctionParams")
}

func init() {
	proto.RegisterFile("comdex/auctionsV2/v1beta1/auction.proto", fileDescriptor_8ee47f5a405fa8ba)
}

var fileDescriptor_8ee47f5a405fa8ba = []byte{
	// 862 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4d, 0x6f, 0xdc, 0x44,
	0x18, 0x8e, 0xa1, 0x4d, 0x77, 0x67, 0x13, 0x6d, 0xe3, 0x92, 0xd4, 0xdd, 0x80, 0xbd, 0x0c, 0x52,
	0xbb, 0x42, 0xaa, 0xad, 0x86, 0x9e, 0xb8, 0xe1, 0xb4, 0x88, 0x45, 0x02, 0x22, 0x77, 0x55, 0x21,
	0x24, 0x64, 0x8d, 0x3d, 0xb3, 0xcb, 0x28, 0xb6, 0xc7, 0xb2, 0xc7, 0x29, 0xf9, 0x0f, 0x80, 0x72,
	0xe5, 0x87, 0xf0, 0x1f, 0x7a, 0xec, 0x11, 0x71, 0x30, 0x28, 0xf9, 0x07, 0x3e, 0x72, 0x42, 0xf3,
	0xe1, 0xb8, 0xeb, 0x06, 0x95, 0x48, 0x9c, 0xec, 0x79, 0x3f, 0x9e, 0xe7, 0x99, 0x77, 0xe6, 0x7d,
	0x07, 0x3c, 0x88, 0x59, 0x8a, 0xc9, 0x8f, 0x1e, 0xaa, 0x62, 0x4e, 0x59, 0x56, 0x3e, 0x3f, 0xf0,
	0x4e, 0x1e, 0x45, 0x84, 0xa3, 0x47, 0xad, 0xc9, 0xcd, 0x0b, 0xc6, 0x99, 0x79, 0x4f, 0x05, 0xba,
	0x5d, 0xa0, 0xab, 0x03, 0x27, 0xef, 0xad, 0xd8, 0x8a, 0xc9, 0x28, 0x4f, 0xfc, 0xa9, 0x84, 0x89,
	0xb3, 0x62, 0x6c, 0x95, 0x10, 0x4f, 0xae, 0xa2, 0x6a, 0xe9, 0x71, 0x9a, 0x92, 0x92, 0xa3, 0x34,
	0xd7, 0x01, 0x76, 0xcc, 0xca, 0x94, 0x95, 0x5e, 0x84, 0x4a, 0x72, 0x49, 0x1a, 0x33, 0xaa, 0x19,
	0xe1, 0x6f, 0x5b, 0x60, 0xf0, 0x99, 0x66, 0x33, 0x1f, 0x03, 0xa0, 0x99, 0x43, 0x8a, 0x2d, 0x63,
	0x6a, 0xcc, 0x6e, 0xf8, 0xbb, 0x4d, 0xed, 0xec, 0x9c, 0xa2, 0x34, 0xf9, 0x14, 0x76, 0x3e, 0x18,
	0x0c, 0xf5, 0x62, 0x8e, 0xcd, 0x33, 0x03, 0xdc, 0x8e, 0x59, 0x92, 0x20, 0x4e, 0x0a, 0x94, 0x84,
	0x9c, 0x1d, 0x93, 0xcc, 0x7a, 0x67, 0x6a, 0xcc, 0x46, 0x07, 0xf7, 0x5c, 0x45, 0xef, 0x0a, 0xfa,
	0x76, 0x2b, 0xee, 0x21, 0xa3, 0x99, 0xff, 0xe5, 0xcb, 0xda, 0xd9, 0x68, 0x6a, 0xe7, 0xae, 0xc2,
	0xee, 0x03, 0xc0, 0xbf, 0x6b, 0xe7, 0xc1, 0x8a, 0xf2, 0x1f, 0xaa, 0xc8, 0x8d, 0x59, 0xea, 0xe9,
	0x6d, 0xa8, 0xcf, 0xc3, 0x12, 0x1f, 0x7b, 0xfc, 0x34, 0x27, 0xa5, 0xc4, 0x0a, 0xc6, 0x5d, 0xf6,
	0x42, 0x24, 0x9b, 0xbf, 0x18, 0x00, 0x60, 0x12, 0x71, 0x2d, 0xe6, 0xc6, 0xdb, 0xc4, 0x2c, 0xb4,
	0x98, 0x0f, 0x95, 0x18, 0x9a, 0x2d, 0x13, 0xf6, 0x42, 0x25, 0x87, 0x1c, 0x15, 0x2b, 0xc2, 0x43,
	0x94, 0xb2, 0x2a, 0xe3, 0xd7, 0x92, 0x35, 0x14, 0x12, 0x94, 0xa0, 0x2f, 0xc0, 0x0e, 0x8a, 0x39,
	0x3d, 0x21, 0x61, 0x44, 0x31, 0xa6, 0xd9, 0x4a, 0x14, 0x78, 0x53, 0x16, 0xf8, 0xfd, 0xa6, 0x76,
	0x2c, 0x5d, 0xe0, 0x7e, 0x08, 0x0c, 0xc6, 0xca, 0xe6, 0x2b, 0xd3, 0x1c, 0x9b, 0x31, 0x18, 0x75,
	0xfe, 0xd2, 0xba, 0x35, 0x7d, 0x77, 0x36, 0x3a, 0xf8, 0xd8, 0xfd, 0xd7, 0x8b, 0xe3, 0x46, 0x14,
	0x7f, 0xf3, 0x22, 0x23, 0xc5, 0x57, 0x28, 0xcf, 0x69, 0xb6, 0xf2, 0xf7, 0x9a, 0xda, 0x31, 0x15,
	0xdf, 0x6b, 0x40, 0x30, 0x00, 0x51, 0xcb, 0x51, 0x9a, 0x11, 0x10, 0xab, 0x70, 0x89, 0x62, 0xce,
	0x0a, 0x6b, 0x30, 0x35, 0x66, 0x43, 0xff, 0x50, 0xd4, 0xe8, 0x8f, 0xda, 0xb9, 0xff, 0x1f, 0xb6,
	0xff, 0x84, 0xc4, 0xdd, 0xb5, 0xe9, 0x90, 0x60, 0x30, 0x8c, 0x28, 0xfe, 0x5c, 0xfe, 0x9b, 0xbf,
	0x1a, 0xc0, 0xee, 0x9f, 0x7a, 0x48, 0x33, 0xca, 0x29, 0x4a, 0xc2, 0xbc, 0xa0, 0x31, 0xb1, 0x86,
	0x92, 0x78, 0x71, 0x6d, 0x62, 0xa8, 0x88, 0x59, 0xc5, 0x5f, 0x3b, 0xc7, 0x35, 0x68, 0x18, 0xec,
	0xf7, 0xee, 0xcc, 0x5c, 0xb9, 0x8f, 0x84, 0xf7, 0x6a, 0x6d, 0x71, 0x55, 0x14, 0x24, 0xe3, 0x5a,
	0x1b, 0xf8, 0x3f, 0xb5, 0xad, 0x41, 0xbf, 0xa9, 0xed, 0x50, 0xb9, 0x95, 0xb6, 0x9f, 0x0d, 0x30,
	0x79, 0x43, 0x1b, 0xc9, 0xb0, 0xd6, 0x35, 0x92, 0xba, 0x8e, 0xae, 0xad, 0xcb, 0xbe, 0x4a, 0xd7,
	0x25, 0x2c, 0x0c, 0xee, 0xf6, 0x34, 0x3d, 0xcd, 0xb0, 0xd2, 0xf3, 0x93, 0x01, 0xac, 0xae, 0xd7,
	0x7a, 0x55, 0xda, 0x92, 0x6a, 0x9e, 0x5d, 0x5b, 0xcd, 0x55, 0x8d, 0xd8, 0x2b, 0xd2, 0xee, 0x65,
	0x77, 0xad, 0x95, 0xc7, 0x07, 0xe3, 0x84, 0xc5, 0xc7, 0x04, 0x87, 0x27, 0xa8, 0x4a, 0xb8, 0xe8,
	0xb3, 0x6d, 0xd9, 0x67, 0x93, 0xa6, 0x76, 0xf6, 0x14, 0x6c, 0x2f, 0x00, 0x06, 0xdb, 0xca, 0xf2,
	0x5c, 0x18, 0xe6, 0xd8, 0xfc, 0x16, 0x80, 0x92, 0xa3, 0x82, 0x87, 0x62, 0x9a, 0x5a, 0x63, 0x39,
	0x3d, 0x26, 0xae, 0x1a, 0xb5, 0x6e, 0x3b, 0x6a, 0xdd, 0x45, 0x3b, 0x6a, 0xfd, 0x0f, 0xf4, 0xf8,
	0xd0, 0x17, 0xbe, 0xcb, 0x85, 0x67, 0x7f, 0x3a, 0x46, 0x30, 0x94, 0x06, 0x11, 0x6e, 0x06, 0x60,
	0x20, 0x6a, 0x2a, 0x71, 0x6f, 0xbf, 0x15, 0x77, 0x5f, 0xe3, 0x8e, 0x15, 0x6e, 0x9b, 0xa9, 0x50,
	0x6f, 0x91, 0x0c, 0x4b, 0xcc, 0x19, 0xd8, 0x44, 0x79, 0x2e, 0x36, 0x6a, 0xca, 0x8d, 0xee, 0x34,
	0xb5, 0xb3, 0xad, 0x07, 0x8a, 0xb4, 0xc3, 0xe0, 0x26, 0xca, 0xf3, 0x39, 0x36, 0xe7, 0x60, 0xab,
	0x9d, 0xe1, 0xa2, 0xd8, 0xd6, 0x9d, 0xa9, 0x31, 0x1b, 0xf8, 0xf7, 0xcf, 0x6b, 0x67, 0xa4, 0xdf,
	0x80, 0xc5, 0x69, 0x4e, 0x9a, 0xda, 0xb9, 0xb3, 0x3e, 0xf0, 0x45, 0x30, 0x0c, 0x46, 0xa8, 0x8b,
	0x81, 0x4f, 0xc1, 0xb8, 0x37, 0x58, 0xcc, 0x5d, 0xb0, 0x29, 0x5a, 0xbd, 0x7d, 0x39, 0x82, 0x9b,
	0x11, 0xc5, 0x73, 0x6c, 0xee, 0x03, 0xd1, 0xf4, 0x21, 0x13, 0xa1, 0xf2, 0x59, 0x18, 0x06, 0x83,
	0x36, 0x15, 0x66, 0x60, 0x5b, 0x33, 0x1f, 0xa1, 0x02, 0xa5, 0xa5, 0xf9, 0x3d, 0xb0, 0x5a, 0x56,
	0x5c, 0x15, 0x48, 0xfe, 0x94, 0x24, 0x66, 0x19, 0x2e, 0xf5, 0x83, 0xf4, 0x51, 0x53, 0x3b, 0xce,
	0xba, 0xbe, 0x7e, 0x24, 0x0c, 0xf6, 0xb4, 0xeb, 0x89, 0xf6, 0x3c, 0x53, 0x0e, 0xff, 0xeb, 0x97,
	0xe7, 0xb6, 0xf1, 0xea, 0xdc, 0x36, 0xfe, 0x3a, 0xb7, 0x8d, 0xb3, 0x0b, 0x7b, 0xe3, 0xd5, 0x85,
	0xbd, 0xf1, 0xfb, 0x85, 0xbd, 0xf1, 0xdd, 0xe3, 0xb5, 0xbb, 0x29, 0x86, 0xe9, 0x43, 0xb6, 0x5c,
	0xd2, 0x98, 0xa2, 0x44, 0xaf, 0xbd, 0xb5, 0x07, 0x5c, 0xde, 0xd6, 0x68, 0x53, 0x9e, 0xda, 0x27,
	0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x92, 0x52, 0xdc, 0xe3, 0xe2, 0x07, 0x00, 0x00,
}

func (m *Auctions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Auctions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Auctions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AuctionType {
		i--
		if m.AuctionType {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x98
	}
	if m.AppId != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.AppId))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x90
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintAuction(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x82
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintAuction(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x7a
	if m.LockedVaultId != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.LockedVaultId))
		i--
		dAtA[i] = 0x68
	}
	{
		size := m.DebtTokenCurrentPrice.Size()
		i -= size
		if _, err := m.DebtTokenCurrentPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size := m.CollateralTokenEndPrice.Size()
		i -= size
		if _, err := m.CollateralTokenEndPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.CollateralTokenCurrentPrice.Size()
		i -= size
		if _, err := m.CollateralTokenCurrentPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.CollateralTokenInitialPrice.Size()
		i -= size
		if _, err := m.CollateralTokenInitialPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.BidFactor.Size()
		i -= size
		if _, err := m.BidFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if len(m.BiddingIds) > 0 {
		for iNdEx := len(m.BiddingIds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BiddingIds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAuction(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.ActiveBiddingId != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.ActiveBiddingId))
		i--
		dAtA[i] = 0x30
	}
	{
		size, err := m.DebtToken.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.CollateralToken.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.AuctionId != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.AuctionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BidOwnerMapping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BidOwnerMapping) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BidOwnerMapping) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BidOwner) > 0 {
		i -= len(m.BidOwner)
		copy(dAtA[i:], m.BidOwner)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.BidOwner)))
		i--
		dAtA[i] = 0x12
	}
	if m.BidId != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.BidId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AuctionParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuctionParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuctionParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AuctionDurationSeconds != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.AuctionDurationSeconds))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuction(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Auctions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AuctionId != 0 {
		n += 1 + sovAuction(uint64(m.AuctionId))
	}
	l = m.CollateralToken.Size()
	n += 1 + l + sovAuction(uint64(l))
	l = m.DebtToken.Size()
	n += 1 + l + sovAuction(uint64(l))
	if m.ActiveBiddingId != 0 {
		n += 1 + sovAuction(uint64(m.ActiveBiddingId))
	}
	if len(m.BiddingIds) > 0 {
		for _, e := range m.BiddingIds {
			l = e.Size()
			n += 1 + l + sovAuction(uint64(l))
		}
	}
	l = m.BidFactor.Size()
	n += 1 + l + sovAuction(uint64(l))
	l = m.CollateralTokenInitialPrice.Size()
	n += 1 + l + sovAuction(uint64(l))
	l = m.CollateralTokenCurrentPrice.Size()
	n += 1 + l + sovAuction(uint64(l))
	l = m.CollateralTokenEndPrice.Size()
	n += 1 + l + sovAuction(uint64(l))
	l = m.DebtTokenCurrentPrice.Size()
	n += 1 + l + sovAuction(uint64(l))
	if m.LockedVaultId != 0 {
		n += 1 + sovAuction(uint64(m.LockedVaultId))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovAuction(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 2 + l + sovAuction(uint64(l))
	if m.AppId != 0 {
		n += 2 + sovAuction(uint64(m.AppId))
	}
	if m.AuctionType {
		n += 3
	}
	return n
}

func (m *BidOwnerMapping) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BidId != 0 {
		n += 1 + sovAuction(uint64(m.BidId))
	}
	l = len(m.BidOwner)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	return n
}

func (m *AuctionParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AuctionDurationSeconds != 0 {
		n += 1 + sovAuction(uint64(m.AuctionDurationSeconds))
	}
	return n
}

func sovAuction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuction(x uint64) (n int) {
	return sovAuction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Auctions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: Auctions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Auctions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionId", wireType)
			}
			m.AuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralToken", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CollateralToken.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebtToken", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DebtToken.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveBiddingId", wireType)
			}
			m.ActiveBiddingId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActiveBiddingId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BiddingIds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BiddingIds = append(m.BiddingIds, &BidOwnerMapping{})
			if err := m.BiddingIds[len(m.BiddingIds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BidFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralTokenInitialPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CollateralTokenInitialPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralTokenCurrentPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CollateralTokenCurrentPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralTokenEndPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CollateralTokenEndPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebtTokenCurrentPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DebtTokenCurrentPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockedVaultId", wireType)
			}
			m.LockedVaultId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LockedVaultId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
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
					return ErrIntOverflowAuction
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
		case 19:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionType", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
			m.AuctionType = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
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
func (m *BidOwnerMapping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: bidOwnerMapping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: bidOwnerMapping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidId", wireType)
			}
			m.BidId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BidId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BidOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
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
func (m *AuctionParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: AuctionParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuctionParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionDurationSeconds", wireType)
			}
			m.AuctionDurationSeconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
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
func skipAuction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
				return 0, ErrInvalidLengthAuction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuction = fmt.Errorf("proto: unexpected end of group")
)
