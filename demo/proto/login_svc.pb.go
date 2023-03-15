// Code generated by protoc-gen-go. DO NOT EDIT.
// source: login_svc.proto

/*
Package sproto is a generated protocol buffer package.

It is generated from these files:
	login_svc.proto

It has these top-level messages:
	UserLoginReq
	GetSessionReq
	OnLogin
*/
package sproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserLoginReq struct {
	GateAddr   string `protobuf:"bytes,1,opt,name=GateAddr" json:"GateAddr,omitempty"`
	GateClient uint64 `protobuf:"varint,2,opt,name=GateClient" json:"GateClient,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Password   string `protobuf:"bytes,4,opt,name=Password" json:"Password,omitempty"`
}

func (m *UserLoginReq) Reset()                    { *m = UserLoginReq{} }
func (m *UserLoginReq) String() string            { return proto.CompactTextString(m) }
func (*UserLoginReq) ProtoMessage()               {}
func (*UserLoginReq) Descriptor() ([]byte, []int) { return fileDescriptor_login_svc, []int{0} }

func (m *UserLoginReq) GetGateAddr() string {
	if m != nil {
		return m.GateAddr
	}
	return ""
}

func (m *UserLoginReq) GetGateClient() uint64 {
	if m != nil {
		return m.GateClient
	}
	return 0
}

func (m *UserLoginReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserLoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type GetSessionReq struct {
	UserName string `protobuf:"bytes,1,opt,name=UserName" json:"UserName,omitempty"`
}

func (m *GetSessionReq) Reset()                    { *m = GetSessionReq{} }
func (m *GetSessionReq) String() string            { return proto.CompactTextString(m) }
func (*GetSessionReq) ProtoMessage()               {}
func (*GetSessionReq) Descriptor() ([]byte, []int) { return fileDescriptor_login_svc, []int{1} }

func (m *GetSessionReq) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type OnLogin struct {
	// Broadcast为广播消息标识
	Broadcast  bool   `protobuf:"varint,1,opt,name=Broadcast" json:"Broadcast,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	NameHash   uint32 `protobuf:"varint,3,opt,name=NameHash" json:"NameHash,omitempty"`
	OwnerAddr  string `protobuf:"bytes,4,opt,name=OwnerAddr" json:"OwnerAddr,omitempty"`
	Token      string `protobuf:"bytes,5,opt,name=Token" json:"Token,omitempty"`
	GateAddr   string `protobuf:"bytes,6,opt,name=GateAddr" json:"GateAddr,omitempty"`
	GateClient uint64 `protobuf:"varint,7,opt,name=GateClient" json:"GateClient,omitempty"`
}

func (m *OnLogin) Reset()                    { *m = OnLogin{} }
func (m *OnLogin) String() string            { return proto.CompactTextString(m) }
func (*OnLogin) ProtoMessage()               {}
func (*OnLogin) Descriptor() ([]byte, []int) { return fileDescriptor_login_svc, []int{2} }

func (m *OnLogin) GetBroadcast() bool {
	if m != nil {
		return m.Broadcast
	}
	return false
}

func (m *OnLogin) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OnLogin) GetNameHash() uint32 {
	if m != nil {
		return m.NameHash
	}
	return 0
}

func (m *OnLogin) GetOwnerAddr() string {
	if m != nil {
		return m.OwnerAddr
	}
	return ""
}

func (m *OnLogin) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *OnLogin) GetGateAddr() string {
	if m != nil {
		return m.GateAddr
	}
	return ""
}

func (m *OnLogin) GetGateClient() uint64 {
	if m != nil {
		return m.GateClient
	}
	return 0
}

func init() {
	proto.RegisterType((*UserLoginReq)(nil), "sproto.UserLoginReq")
	proto.RegisterType((*GetSessionReq)(nil), "sproto.GetSessionReq")
	proto.RegisterType((*OnLogin)(nil), "sproto.OnLogin")
}

func init() { proto.RegisterFile("login_svc.proto", fileDescriptor_login_svc) }

var fileDescriptor_login_svc = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0xc9, 0x4f, 0xcf,
	0xcc, 0x8b, 0x2f, 0x2e, 0x4b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2b, 0x06, 0xd3,
	0x4a, 0x55, 0x5c, 0x3c, 0xa1, 0xc5, 0xa9, 0x45, 0x3e, 0x20, 0xe9, 0xa0, 0xd4, 0x42, 0x21, 0x29,
	0x2e, 0x0e, 0xf7, 0xc4, 0x92, 0x54, 0xc7, 0x94, 0x94, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x38, 0x5f, 0x48, 0x8e, 0x8b, 0x0b, 0xc4, 0x76, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x91, 0x60,
	0x52, 0x60, 0xd4, 0x60, 0x09, 0x42, 0x12, 0x11, 0x12, 0xe2, 0x62, 0xf1, 0x4b, 0xcc, 0x4d, 0x95,
	0x60, 0x06, 0xeb, 0x03, 0xb3, 0x41, 0xe6, 0x05, 0x24, 0x16, 0x17, 0x97, 0xe7, 0x17, 0xa5, 0x48,
	0xb0, 0x40, 0xcc, 0x83, 0xf1, 0x95, 0xb4, 0xb9, 0x78, 0xdd, 0x53, 0x4b, 0x82, 0x53, 0x8b, 0x8b,
	0x33, 0xf3, 0x61, 0x96, 0x83, 0x1c, 0x03, 0x36, 0x04, 0x6a, 0x39, 0x8c, 0xaf, 0x74, 0x9c, 0x91,
	0x8b, 0xdd, 0x3f, 0x0f, 0xec, 0x4e, 0x21, 0x19, 0x2e, 0x4e, 0xa7, 0xa2, 0xfc, 0xc4, 0x94, 0xe4,
	0xc4, 0xe2, 0x12, 0xb0, 0x42, 0x8e, 0x20, 0x84, 0x00, 0xdc, 0x19, 0x4c, 0xa8, 0xce, 0x00, 0xd1,
	0x1e, 0x89, 0xc5, 0x19, 0x60, 0xe7, 0xf1, 0x06, 0xc1, 0xf9, 0x20, 0xd3, 0xfc, 0xcb, 0xf3, 0x52,
	0x8b, 0xc0, 0x7e, 0x86, 0xb8, 0x11, 0x21, 0x20, 0x24, 0xc2, 0xc5, 0x1a, 0x92, 0x9f, 0x9d, 0x9a,
	0x27, 0xc1, 0x0a, 0x96, 0x81, 0x70, 0x50, 0x82, 0x89, 0x0d, 0x6f, 0x30, 0xb1, 0xa3, 0x07, 0x53,
	0x12, 0x1b, 0x38, 0xe4, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x21, 0x0d, 0x6f, 0x4a, 0x94,
	0x01, 0x00, 0x00,
}
