// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/zookeeper-data.proto

/*
Package model is a generated protocol buffer package.

It is generated from these files:
	model/zookeeper-data.proto
	model/zookeeper-proto.proto
	model/zookeeper-quorum.proto
	model/zookeeper-txn.proto
	model/zookeeper-persistance.proto

It has these top-level messages:
	Id
	ACL
	Stat
	StatPersisted
	StatPersistedV1
	ConnectRequest
	ConnectResponse
	SetWatches
	RequestHeader
	MultiHeader
	AuthPacket
	ReplyHeader
	GetDataRequest
	SetDataRequest
	SetDataResponse
	GetSASLRequest
	SetSASLRequest
	SetSASLResponse
	CreateRequest
	DeleteRequest
	GetChildrenRequest
	GetChildren2Request
	CheckVersionRequest
	GetMaxChildrenRequest
	GetMaxChildrenResponse
	SetMaxChildrenRequest
	SyncRequest
	SyncResponse
	GetACLRequest
	SetACLRequest
	SetACLResponse
	WatcherEvent
	ErrorResponse
	CreateResponse
	ExistsRequest
	ExistsResponse
	GetDataResponse
	GetChildrenResponse
	GetChildren2Response
	GetACLResponse
	LearnerInfo
	QuorumPacket
	TxnHeader
	CreateTxnV0
	CreateTxn
	DeleteTxn
	SetDataTxn
	CheckVersionTxn
	SetACLTxn
	SetMaxChildrenTxn
	CreateSessionTxn
	ErrorTxn
	Txn
	MultiTxn
	FileHandler
*/
package model

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

type Id struct {
	Scheme string `protobuf:"bytes,1,opt,name=scheme" json:"scheme,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *Id) Reset()                    { *m = Id{} }
func (m *Id) String() string            { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()               {}
func (*Id) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Id) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *Id) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ACL struct {
	Perms int32 `protobuf:"varint,1,opt,name=perms" json:"perms,omitempty"`
	Id    *Id   `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *ACL) Reset()                    { *m = ACL{} }
func (m *ACL) String() string            { return proto.CompactTextString(m) }
func (*ACL) ProtoMessage()               {}
func (*ACL) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ACL) GetPerms() int32 {
	if m != nil {
		return m.Perms
	}
	return 0
}

func (m *ACL) GetId() *Id {
	if m != nil {
		return m.Id
	}
	return nil
}

// information shared with the client
type Stat struct {
	Czxid          int64 `protobuf:"varint,1,opt,name=czxid" json:"czxid,omitempty"`
	Mzxid          int64 `protobuf:"varint,2,opt,name=mzxid" json:"mzxid,omitempty"`
	Ctime          int64 `protobuf:"varint,3,opt,name=ctime" json:"ctime,omitempty"`
	Mtime          int64 `protobuf:"varint,4,opt,name=mtime" json:"mtime,omitempty"`
	Version        int32 `protobuf:"varint,5,opt,name=version" json:"version,omitempty"`
	Cversion       int32 `protobuf:"varint,6,opt,name=cversion" json:"cversion,omitempty"`
	Aversion       int32 `protobuf:"varint,7,opt,name=aversion" json:"aversion,omitempty"`
	EphemeralOwner int64 `protobuf:"varint,8,opt,name=ephemeralOwner" json:"ephemeralOwner,omitempty"`
	DataLength     int32 `protobuf:"varint,9,opt,name=dataLength" json:"dataLength,omitempty"`
	NumChildren    int32 `protobuf:"varint,10,opt,name=numChildren" json:"numChildren,omitempty"`
	Pzxid          int64 `protobuf:"varint,11,opt,name=pzxid" json:"pzxid,omitempty"`
}

func (m *Stat) Reset()                    { *m = Stat{} }
func (m *Stat) String() string            { return proto.CompactTextString(m) }
func (*Stat) ProtoMessage()               {}
func (*Stat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Stat) GetCzxid() int64 {
	if m != nil {
		return m.Czxid
	}
	return 0
}

func (m *Stat) GetMzxid() int64 {
	if m != nil {
		return m.Mzxid
	}
	return 0
}

func (m *Stat) GetCtime() int64 {
	if m != nil {
		return m.Ctime
	}
	return 0
}

func (m *Stat) GetMtime() int64 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

func (m *Stat) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Stat) GetCversion() int32 {
	if m != nil {
		return m.Cversion
	}
	return 0
}

func (m *Stat) GetAversion() int32 {
	if m != nil {
		return m.Aversion
	}
	return 0
}

func (m *Stat) GetEphemeralOwner() int64 {
	if m != nil {
		return m.EphemeralOwner
	}
	return 0
}

func (m *Stat) GetDataLength() int32 {
	if m != nil {
		return m.DataLength
	}
	return 0
}

func (m *Stat) GetNumChildren() int32 {
	if m != nil {
		return m.NumChildren
	}
	return 0
}

func (m *Stat) GetPzxid() int64 {
	if m != nil {
		return m.Pzxid
	}
	return 0
}

// information explicitly stored by the server persistently
type StatPersisted struct {
	Czxid          int64 `protobuf:"varint,1,opt,name=czxid" json:"czxid,omitempty"`
	Mzxid          int64 `protobuf:"varint,2,opt,name=mzxid" json:"mzxid,omitempty"`
	Ctime          int64 `protobuf:"varint,3,opt,name=ctime" json:"ctime,omitempty"`
	Mtime          int64 `protobuf:"varint,4,opt,name=mtime" json:"mtime,omitempty"`
	Version        int32 `protobuf:"varint,5,opt,name=version" json:"version,omitempty"`
	Cversion       int32 `protobuf:"varint,6,opt,name=cversion" json:"cversion,omitempty"`
	Aversion       int32 `protobuf:"varint,7,opt,name=aversion" json:"aversion,omitempty"`
	EphemeralOwner int64 `protobuf:"varint,8,opt,name=ephemeralOwner" json:"ephemeralOwner,omitempty"`
	Pzxid          int64 `protobuf:"varint,9,opt,name=pzxid" json:"pzxid,omitempty"`
}

func (m *StatPersisted) Reset()                    { *m = StatPersisted{} }
func (m *StatPersisted) String() string            { return proto.CompactTextString(m) }
func (*StatPersisted) ProtoMessage()               {}
func (*StatPersisted) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StatPersisted) GetCzxid() int64 {
	if m != nil {
		return m.Czxid
	}
	return 0
}

func (m *StatPersisted) GetMzxid() int64 {
	if m != nil {
		return m.Mzxid
	}
	return 0
}

func (m *StatPersisted) GetCtime() int64 {
	if m != nil {
		return m.Ctime
	}
	return 0
}

func (m *StatPersisted) GetMtime() int64 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

func (m *StatPersisted) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *StatPersisted) GetCversion() int32 {
	if m != nil {
		return m.Cversion
	}
	return 0
}

func (m *StatPersisted) GetAversion() int32 {
	if m != nil {
		return m.Aversion
	}
	return 0
}

func (m *StatPersisted) GetEphemeralOwner() int64 {
	if m != nil {
		return m.EphemeralOwner
	}
	return 0
}

func (m *StatPersisted) GetPzxid() int64 {
	if m != nil {
		return m.Pzxid
	}
	return 0
}

// information explicitly stored by the version 1 database of servers
type StatPersistedV1 struct {
	Czxid          int64 `protobuf:"varint,1,opt,name=czxid" json:"czxid,omitempty"`
	Mzxid          int64 `protobuf:"varint,2,opt,name=mzxid" json:"mzxid,omitempty"`
	Ctime          int64 `protobuf:"varint,3,opt,name=ctime" json:"ctime,omitempty"`
	Mtime          int64 `protobuf:"varint,4,opt,name=mtime" json:"mtime,omitempty"`
	Version        int32 `protobuf:"varint,5,opt,name=version" json:"version,omitempty"`
	Cversion       int32 `protobuf:"varint,6,opt,name=cversion" json:"cversion,omitempty"`
	Aversion       int32 `protobuf:"varint,7,opt,name=aversion" json:"aversion,omitempty"`
	EphemeralOwner int64 `protobuf:"varint,8,opt,name=ephemeralOwner" json:"ephemeralOwner,omitempty"`
}

func (m *StatPersistedV1) Reset()                    { *m = StatPersistedV1{} }
func (m *StatPersistedV1) String() string            { return proto.CompactTextString(m) }
func (*StatPersistedV1) ProtoMessage()               {}
func (*StatPersistedV1) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StatPersistedV1) GetCzxid() int64 {
	if m != nil {
		return m.Czxid
	}
	return 0
}

func (m *StatPersistedV1) GetMzxid() int64 {
	if m != nil {
		return m.Mzxid
	}
	return 0
}

func (m *StatPersistedV1) GetCtime() int64 {
	if m != nil {
		return m.Ctime
	}
	return 0
}

func (m *StatPersistedV1) GetMtime() int64 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

func (m *StatPersistedV1) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *StatPersistedV1) GetCversion() int32 {
	if m != nil {
		return m.Cversion
	}
	return 0
}

func (m *StatPersistedV1) GetAversion() int32 {
	if m != nil {
		return m.Aversion
	}
	return 0
}

func (m *StatPersistedV1) GetEphemeralOwner() int64 {
	if m != nil {
		return m.EphemeralOwner
	}
	return 0
}

func init() {
	proto.RegisterType((*Id)(nil), "model.Id")
	proto.RegisterType((*ACL)(nil), "model.ACL")
	proto.RegisterType((*Stat)(nil), "model.Stat")
	proto.RegisterType((*StatPersisted)(nil), "model.StatPersisted")
	proto.RegisterType((*StatPersistedV1)(nil), "model.StatPersistedV1")
}

func init() { proto.RegisterFile("model/zookeeper-data.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x93, 0xcd, 0x4a, 0xfb, 0x40,
	0x14, 0xc5, 0xc9, 0xb4, 0x69, 0x9b, 0x5b, 0xfe, 0xfd, 0xc3, 0x20, 0x32, 0x76, 0x21, 0xa5, 0x0b,
	0x71, 0xa1, 0x15, 0x15, 0xdc, 0x4b, 0x57, 0x85, 0x82, 0x12, 0xc1, 0x7d, 0xcc, 0x5c, 0xcc, 0x60,
	0x26, 0x13, 0x26, 0xe3, 0x07, 0x7d, 0x2a, 0xdf, 0xc9, 0x87, 0x70, 0x2b, 0x73, 0xf3, 0x41, 0xf4,
	0x11, 0xba, 0x3c, 0xbf, 0x73, 0x66, 0x32, 0xe7, 0x40, 0x60, 0xae, 0x8d, 0xc4, 0xfc, 0x62, 0x67,
	0xcc, 0x0b, 0x62, 0x89, 0xf6, 0x5c, 0x26, 0x2e, 0x59, 0x95, 0xd6, 0x38, 0xc3, 0x43, 0xf2, 0x96,
	0x67, 0xc0, 0x36, 0x92, 0x1f, 0xc2, 0xa8, 0x4a, 0x33, 0xd4, 0x28, 0x82, 0x45, 0x70, 0x1a, 0xc5,
	0x8d, 0xe2, 0x33, 0x60, 0x4a, 0x0a, 0x46, 0x8c, 0x29, 0xb9, 0xbc, 0x81, 0xc1, 0xed, 0x7a, 0xcb,
	0x0f, 0x20, 0x2c, 0xd1, 0xea, 0x8a, 0xd2, 0x61, 0x5c, 0x0b, 0x7e, 0xd4, 0x85, 0xa7, 0x57, 0xd1,
	0x8a, 0xae, 0x5f, 0x6d, 0x24, 0x9d, 0xfb, 0x64, 0x30, 0x7c, 0x70, 0x89, 0xf3, 0x27, 0xd3, 0xdd,
	0x87, 0x92, 0x74, 0x72, 0x10, 0xd7, 0xc2, 0x53, 0x4d, 0x94, 0xd5, 0x54, 0xb7, 0x34, 0x75, 0x4a,
	0xa3, 0x18, 0x34, 0x59, 0x2f, 0x28, 0x4b, 0x74, 0xd8, 0x64, 0x89, 0x0a, 0x18, 0xbf, 0xa1, 0xad,
	0x94, 0x29, 0x44, 0x48, 0x6f, 0x6a, 0x25, 0x9f, 0xc3, 0x24, 0x6d, 0xad, 0x11, 0x59, 0x9d, 0xf6,
	0x5e, 0xd2, 0x7a, 0xe3, 0xda, 0x6b, 0x35, 0x3f, 0x81, 0x19, 0x96, 0x7e, 0x04, 0x9b, 0xe4, 0x77,
	0xef, 0x05, 0x5a, 0x31, 0xa1, 0x0f, 0xfe, 0xa1, 0xfc, 0x18, 0xc0, 0xaf, 0xba, 0xc5, 0xe2, 0xd9,
	0x65, 0x22, 0xa2, 0x5b, 0x7a, 0x84, 0x2f, 0x60, 0x5a, 0xbc, 0xea, 0x75, 0xa6, 0x72, 0x69, 0xb1,
	0x10, 0x40, 0x81, 0x3e, 0xa2, 0x35, 0xa9, 0xfd, 0xb4, 0x6e, 0x44, 0x62, 0xf9, 0x1d, 0xc0, 0x3f,
	0x3f, 0xd9, 0xbd, 0x7f, 0x4f, 0xe5, 0x50, 0xee, 0xe1, 0x76, 0x5d, 0xf3, 0xa8, 0xdf, 0xfc, 0x2b,
	0x80, 0xff, 0xbf, 0x9a, 0x3f, 0x5e, 0xee, 0x5f, 0xf7, 0xa7, 0x11, 0xfd, 0x86, 0xd7, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x82, 0xed, 0x9f, 0x5a, 0xa4, 0x03, 0x00, 0x00,
}