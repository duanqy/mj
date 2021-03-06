// Code generated by protoc-gen-go.
// source: coin.proto
// DO NOT EDIT!

package protocol

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

// 更新玩家的金币等经济资源，当玩家的经济资源任何时候有变动，服务器主动推送
type SResource struct {
	Code             *uint32  `protobuf:"varint,1,opt,name=code,def=5000" json:"code,omitempty"`
	List             []*ResVo `protobuf:"bytes,2,rep,name=list" json:"list,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SResource) Reset()         { *m = SResource{} }
func (m *SResource) String() string { return proto.CompactTextString(m) }
func (*SResource) ProtoMessage()    {}

const Default_SResource_Code uint32 = 5000

func (m *SResource) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_SResource_Code
}

func (m *SResource) GetList() []*ResVo {
	if m != nil {
		return m.List
	}
	return nil
}

type ResVo struct {
	Id               *uint32 `protobuf:"varint,2,req,name=id" json:"id,omitempty"`
	Count            *int32  `protobuf:"varint,3,req,name=count" json:"count,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ResVo) Reset()         { *m = ResVo{} }
func (m *ResVo) String() string { return proto.CompactTextString(m) }
func (*ResVo) ProtoMessage()    {}

func (m *ResVo) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *ResVo) GetCount() int32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

// 获取虚拟货币
type CGetCurrency struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=5001" json:"code,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CGetCurrency) Reset()         { *m = CGetCurrency{} }
func (m *CGetCurrency) String() string { return proto.CompactTextString(m) }
func (*CGetCurrency) ProtoMessage()    {}

const Default_CGetCurrency_Code uint32 = 5001

func (m *CGetCurrency) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_CGetCurrency_Code
}

type SGetCurrency struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=5001" json:"code,omitempty"`
	Coin             *uint32 `protobuf:"varint,2,req,name=coin" json:"coin,omitempty"`
	Diamond          *uint32 `protobuf:"varint,3,req,name=diamond" json:"diamond,omitempty"`
	Exchange         *uint32 `protobuf:"varint,4,req,name=exchange" json:"exchange,omitempty"`
	Ticket           *uint32 `protobuf:"varint,5,req,name=ticket" json:"ticket,omitempty"`
	Roomcard         *uint32 `protobuf:"varint,6,req,name=roomcard" json:"roomcard,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SGetCurrency) Reset()         { *m = SGetCurrency{} }
func (m *SGetCurrency) String() string { return proto.CompactTextString(m) }
func (*SGetCurrency) ProtoMessage()    {}

const Default_SGetCurrency_Code uint32 = 5001

func (m *SGetCurrency) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_SGetCurrency_Code
}

func (m *SGetCurrency) GetCoin() uint32 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *SGetCurrency) GetDiamond() uint32 {
	if m != nil && m.Diamond != nil {
		return *m.Diamond
	}
	return 0
}

func (m *SGetCurrency) GetExchange() uint32 {
	if m != nil && m.Exchange != nil {
		return *m.Exchange
	}
	return 0
}

func (m *SGetCurrency) GetTicket() uint32 {
	if m != nil && m.Ticket != nil {
		return *m.Ticket
	}
	return 0
}

func (m *SGetCurrency) GetRoomcard() uint32 {
	if m != nil && m.Roomcard != nil {
		return *m.Roomcard
	}
	return 0
}

func init() {
}
