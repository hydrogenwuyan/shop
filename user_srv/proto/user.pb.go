// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_srv/proto/user.proto

package shop_user_srv

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 用户数据
type UserInfo struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	CreateTime           int64    `protobuf:"varint,2,opt,name=createTime,proto3" json:"createTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_615921daf335c569, []int{0}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserInfo) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

// 错误码
type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_615921daf335c569, []int{1}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

// 用户信息请求
type CSUserInfo struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CSUserInfo) Reset()         { *m = CSUserInfo{} }
func (m *CSUserInfo) String() string { return proto.CompactTextString(m) }
func (*CSUserInfo) ProtoMessage()    {}
func (*CSUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_615921daf335c569, []int{2}
}

func (m *CSUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CSUserInfo.Unmarshal(m, b)
}
func (m *CSUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CSUserInfo.Marshal(b, m, deterministic)
}
func (m *CSUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CSUserInfo.Merge(m, src)
}
func (m *CSUserInfo) XXX_Size() int {
	return xxx_messageInfo_CSUserInfo.Size(m)
}
func (m *CSUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CSUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CSUserInfo proto.InternalMessageInfo

func (m *CSUserInfo) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CSUserInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SCUserInfo struct {
	Error                *Error    `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *UserInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SCUserInfo) Reset()         { *m = SCUserInfo{} }
func (m *SCUserInfo) String() string { return proto.CompactTextString(m) }
func (*SCUserInfo) ProtoMessage()    {}
func (*SCUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_615921daf335c569, []int{3}
}

func (m *SCUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SCUserInfo.Unmarshal(m, b)
}
func (m *SCUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SCUserInfo.Marshal(b, m, deterministic)
}
func (m *SCUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SCUserInfo.Merge(m, src)
}
func (m *SCUserInfo) XXX_Size() int {
	return xxx_messageInfo_SCUserInfo.Size(m)
}
func (m *SCUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SCUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SCUserInfo proto.InternalMessageInfo

func (m *SCUserInfo) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *SCUserInfo) GetInfo() *UserInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "shop.user.srv.UserInfo")
	proto.RegisterType((*Error)(nil), "shop.user.srv.Error")
	proto.RegisterType((*CSUserInfo)(nil), "shop.user.srv.CSUserInfo")
	proto.RegisterType((*SCUserInfo)(nil), "shop.user.srv.SCUserInfo")
}

func init() { proto.RegisterFile("user_srv/proto/user.proto", fileDescriptor_615921daf335c569) }

var fileDescriptor_615921daf335c569 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0x4d, 0x4f, 0x02, 0x31,
	0x14, 0x14, 0xdd, 0x25, 0x38, 0xc4, 0x4b, 0xe3, 0x07, 0x70, 0x30, 0xa6, 0x27, 0xa3, 0x49, 0x49,
	0x96, 0x3f, 0x60, 0x20, 0x1e, 0x38, 0x19, 0x8b, 0xc6, 0xa3, 0x59, 0xe9, 0x23, 0x6e, 0xa2, 0x74,
	0xf3, 0xba, 0x60, 0xfc, 0xf7, 0xa6, 0x8f, 0x05, 0x74, 0x2f, 0xde, 0xde, 0xcc, 0x9b, 0x99, 0xce,
	0x4b, 0xd1, 0x5f, 0x05, 0xe2, 0xd7, 0xc0, 0xeb, 0x61, 0xc9, 0xbe, 0xf2, 0xc3, 0x08, 0x8d, 0x8c,
	0xea, 0x24, 0xbc, 0xfb, 0xd2, 0x08, 0x11, 0x78, 0xad, 0xc7, 0xe8, 0x3c, 0x07, 0xe2, 0xe9, 0x72,
	0xe1, 0xd5, 0x39, 0xda, 0x91, 0x9f, 0xba, 0x5e, 0xeb, 0xaa, 0x75, 0x7d, 0x64, 0x6b, 0xa4, 0x2e,
	0x81, 0x39, 0x53, 0x5e, 0xd1, 0x53, 0xf1, 0x49, 0xbd, 0x43, 0xd9, 0xfd, 0x62, 0xf4, 0x08, 0xe9,
	0x3d, 0xb3, 0x67, 0xa5, 0x90, 0xcc, 0xbd, 0x23, 0xb1, 0xa7, 0x56, 0xe6, 0x18, 0xea, 0xa8, 0xca,
	0x8b, 0x0f, 0x31, 0x1e, 0xdb, 0x1a, 0xe9, 0x3b, 0x60, 0x32, 0xfb, 0xf7, 0xe9, 0x01, 0x3a, 0x65,
	0x1e, 0xc2, 0x97, 0x67, 0x57, 0xfb, 0x77, 0x58, 0x13, 0x30, 0x9b, 0xec, 0x12, 0x6e, 0x90, 0x52,
	0x2c, 0x21, 0x01, 0xdd, 0xec, 0xd4, 0xfc, 0xb9, 0xd3, 0x48, 0x41, 0xbb, 0x91, 0xa8, 0x5b, 0x24,
	0xc5, 0x72, 0xe1, 0x25, 0xb1, 0x9b, 0x5d, 0x34, 0xa4, 0xdb, 0x48, 0x2b, 0xa2, 0xec, 0x05, 0x49,
	0x64, 0xd4, 0x03, 0xce, 0x1e, 0x57, 0xc4, 0xdf, 0xdb, 0xf5, 0x78, 0x33, 0x39, 0xd5, 0x6f, 0xf8,
	0xf7, 0x67, 0x0d, 0x9a, 0xab, 0x7d, 0x5f, 0x7d, 0xf0, 0xd6, 0x96, 0x0f, 0x19, 0xfd, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x67, 0xd3, 0x88, 0x2a, 0xad, 0x01, 0x00, 0x00,
}
