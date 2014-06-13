// Code generated by protoc-gen-go.
// source: pb.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb.proto

It has these top-level messages:
	PbClientLogin
	PbServerAcceptLogin
	PbClientLogout
	PbClientPing
	PbC2CTextChat
*/
package pb

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

// 客户端登陆聊天服务器
type PbClientLogin struct {
	Uuid             *string  `protobuf:"bytes,1,req,name=uuid" json:"uuid,omitempty"`
	Version          *float32 `protobuf:"fixed32,2,req,name=version" json:"version,omitempty"`
	Timestamp        *int64   `protobuf:"varint,3,req,name=timestamp" json:"timestamp,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *PbClientLogin) Reset()         { *m = PbClientLogin{} }
func (m *PbClientLogin) String() string { return proto.CompactTextString(m) }
func (*PbClientLogin) ProtoMessage()    {}

func (m *PbClientLogin) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *PbClientLogin) GetVersion() float32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *PbClientLogin) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

// 服务器通知客户端是否接受连接
type PbServerAcceptLogin struct {
	Login            *bool   `protobuf:"varint,1,req,name=login" json:"login,omitempty"`
	TipsMsg          *string `protobuf:"bytes,2,req,name=tips_msg" json:"tips_msg,omitempty"`
	Timestamp        *int64  `protobuf:"varint,3,req,name=timestamp" json:"timestamp,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PbServerAcceptLogin) Reset()         { *m = PbServerAcceptLogin{} }
func (m *PbServerAcceptLogin) String() string { return proto.CompactTextString(m) }
func (*PbServerAcceptLogin) ProtoMessage()    {}

func (m *PbServerAcceptLogin) GetLogin() bool {
	if m != nil && m.Login != nil {
		return *m.Login
	}
	return false
}

func (m *PbServerAcceptLogin) GetTipsMsg() string {
	if m != nil && m.TipsMsg != nil {
		return *m.TipsMsg
	}
	return ""
}

func (m *PbServerAcceptLogin) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

// 客户端下线通知服务器
type PbClientLogout struct {
	Logout           *bool  `protobuf:"varint,1,req,name=logout" json:"logout,omitempty"`
	Timestamp        *int64 `protobuf:"varint,2,req,name=timestamp" json:"timestamp,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PbClientLogout) Reset()         { *m = PbClientLogout{} }
func (m *PbClientLogout) String() string { return proto.CompactTextString(m) }
func (*PbClientLogout) ProtoMessage()    {}

func (m *PbClientLogout) GetLogout() bool {
	if m != nil && m.Logout != nil {
		return *m.Logout
	}
	return false
}

func (m *PbClientLogout) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

// 客户端心跳包
type PbClientPing struct {
	Ping             *bool  `protobuf:"varint,1,req,name=ping" json:"ping,omitempty"`
	Timestamp        *int64 `protobuf:"varint,2,req,name=timestamp" json:"timestamp,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PbClientPing) Reset()         { *m = PbClientPing{} }
func (m *PbClientPing) String() string { return proto.CompactTextString(m) }
func (*PbClientPing) ProtoMessage()    {}

func (m *PbClientPing) GetPing() bool {
	if m != nil && m.Ping != nil {
		return *m.Ping
	}
	return false
}

func (m *PbClientPing) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

// 客户端之间聊天
type PbC2CTextChat struct {
	FromUuid         *string `protobuf:"bytes,1,req,name=from_uuid" json:"from_uuid,omitempty"`
	ToUuid           *string `protobuf:"bytes,2,req,name=to_uuid" json:"to_uuid,omitempty"`
	TextMsg          *string `protobuf:"bytes,3,req,name=text_msg" json:"text_msg,omitempty"`
	Timestamp        *int64  `protobuf:"varint,4,req,name=timestamp" json:"timestamp,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PbC2CTextChat) Reset()         { *m = PbC2CTextChat{} }
func (m *PbC2CTextChat) String() string { return proto.CompactTextString(m) }
func (*PbC2CTextChat) ProtoMessage()    {}

func (m *PbC2CTextChat) GetFromUuid() string {
	if m != nil && m.FromUuid != nil {
		return *m.FromUuid
	}
	return ""
}

func (m *PbC2CTextChat) GetToUuid() string {
	if m != nil && m.ToUuid != nil {
		return *m.ToUuid
	}
	return ""
}

func (m *PbC2CTextChat) GetTextMsg() string {
	if m != nil && m.TextMsg != nil {
		return *m.TextMsg
	}
	return ""
}

func (m *PbC2CTextChat) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func init() {
}