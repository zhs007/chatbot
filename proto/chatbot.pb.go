// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chatbot.proto

package chatbotpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ChatAppType - chat app type
type ChatAppType int32

const (
	ChatAppType_CAT_TELEGRAM ChatAppType = 0
	ChatAppType_CAT_COOLQQ   ChatAppType = 1
)

var ChatAppType_name = map[int32]string{
	0: "CAT_TELEGRAM",
	1: "CAT_COOLQQ",
}
var ChatAppType_value = map[string]int32{
	"CAT_TELEGRAM": 0,
	"CAT_COOLQQ":   1,
}

func (x ChatAppType) String() string {
	return proto.EnumName(ChatAppType_name, int32(x))
}
func (ChatAppType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_chatbot_bcb94d0b7321ef21, []int{0}
}

// UserAppInfo - user app info
type UserAppInfo struct {
	App                  ChatAppType `protobuf:"varint,1,opt,name=app,proto3,enum=chatbotpb.ChatAppType" json:"app,omitempty"`
	Appuid               string      `protobuf:"bytes,2,opt,name=appuid,proto3" json:"appuid,omitempty"`
	Appuname             string      `protobuf:"bytes,3,opt,name=appuname,proto3" json:"appuname,omitempty"`
	Chatnums             int32       `protobuf:"varint,4,opt,name=chatnums,proto3" json:"chatnums,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UserAppInfo) Reset()         { *m = UserAppInfo{} }
func (m *UserAppInfo) String() string { return proto.CompactTextString(m) }
func (*UserAppInfo) ProtoMessage()    {}
func (*UserAppInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatbot_bcb94d0b7321ef21, []int{0}
}
func (m *UserAppInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAppInfo.Unmarshal(m, b)
}
func (m *UserAppInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAppInfo.Marshal(b, m, deterministic)
}
func (dst *UserAppInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAppInfo.Merge(dst, src)
}
func (m *UserAppInfo) XXX_Size() int {
	return xxx_messageInfo_UserAppInfo.Size(m)
}
func (m *UserAppInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAppInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserAppInfo proto.InternalMessageInfo

func (m *UserAppInfo) GetApp() ChatAppType {
	if m != nil {
		return m.App
	}
	return ChatAppType_CAT_TELEGRAM
}

func (m *UserAppInfo) GetAppuid() string {
	if m != nil {
		return m.Appuid
	}
	return ""
}

func (m *UserAppInfo) GetAppuname() string {
	if m != nil {
		return m.Appuname
	}
	return ""
}

func (m *UserAppInfo) GetChatnums() int32 {
	if m != nil {
		return m.Chatnums
	}
	return 0
}

// UserInfo - user info
type UserInfo struct {
	Uid                  int64          `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Apps                 []*UserAppInfo `protobuf:"bytes,3,rep,name=apps,proto3" json:"apps,omitempty"`
	Money                int64          `protobuf:"varint,10,opt,name=money,proto3" json:"money,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatbot_bcb94d0b7321ef21, []int{1}
}
func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (dst *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(dst, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetApps() []*UserAppInfo {
	if m != nil {
		return m.Apps
	}
	return nil
}

func (m *UserInfo) GetMoney() int64 {
	if m != nil {
		return m.Money
	}
	return 0
}

// RegisterAppService - register app service
type RegisterAppService struct {
	Token                string      `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	AppType              ChatAppType `protobuf:"varint,2,opt,name=appType,proto3,enum=chatbotpb.ChatAppType" json:"appType,omitempty"`
	Username             string      `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RegisterAppService) Reset()         { *m = RegisterAppService{} }
func (m *RegisterAppService) String() string { return proto.CompactTextString(m) }
func (*RegisterAppService) ProtoMessage()    {}
func (*RegisterAppService) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatbot_bcb94d0b7321ef21, []int{2}
}
func (m *RegisterAppService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterAppService.Unmarshal(m, b)
}
func (m *RegisterAppService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterAppService.Marshal(b, m, deterministic)
}
func (dst *RegisterAppService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterAppService.Merge(dst, src)
}
func (m *RegisterAppService) XXX_Size() int {
	return xxx_messageInfo_RegisterAppService.Size(m)
}
func (m *RegisterAppService) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterAppService.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterAppService proto.InternalMessageInfo

func (m *RegisterAppService) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RegisterAppService) GetAppType() ChatAppType {
	if m != nil {
		return m.AppType
	}
	return ChatAppType_CAT_TELEGRAM
}

func (m *RegisterAppService) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

// ReplyRegisterAppService - reply RegisterAppService
type ReplyRegisterAppService struct {
	AppType              ChatAppType `protobuf:"varint,1,opt,name=appType,proto3,enum=chatbotpb.ChatAppType" json:"appType,omitempty"`
	Error                string      `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReplyRegisterAppService) Reset()         { *m = ReplyRegisterAppService{} }
func (m *ReplyRegisterAppService) String() string { return proto.CompactTextString(m) }
func (*ReplyRegisterAppService) ProtoMessage()    {}
func (*ReplyRegisterAppService) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatbot_bcb94d0b7321ef21, []int{3}
}
func (m *ReplyRegisterAppService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyRegisterAppService.Unmarshal(m, b)
}
func (m *ReplyRegisterAppService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyRegisterAppService.Marshal(b, m, deterministic)
}
func (dst *ReplyRegisterAppService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyRegisterAppService.Merge(dst, src)
}
func (m *ReplyRegisterAppService) XXX_Size() int {
	return xxx_messageInfo_ReplyRegisterAppService.Size(m)
}
func (m *ReplyRegisterAppService) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyRegisterAppService.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyRegisterAppService proto.InternalMessageInfo

func (m *ReplyRegisterAppService) GetAppType() ChatAppType {
	if m != nil {
		return m.AppType
	}
	return ChatAppType_CAT_TELEGRAM
}

func (m *ReplyRegisterAppService) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*UserAppInfo)(nil), "chatbotpb.UserAppInfo")
	proto.RegisterType((*UserInfo)(nil), "chatbotpb.UserInfo")
	proto.RegisterType((*RegisterAppService)(nil), "chatbotpb.RegisterAppService")
	proto.RegisterType((*ReplyRegisterAppService)(nil), "chatbotpb.ReplyRegisterAppService")
	proto.RegisterEnum("chatbotpb.ChatAppType", ChatAppType_name, ChatAppType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatBotServiceClient is the client API for ChatBotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatBotServiceClient interface {
	// registerAppService - register app service
	RegisterAppService(ctx context.Context, in *RegisterAppService, opts ...grpc.CallOption) (*ReplyRegisterAppService, error)
}

type chatBotServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatBotServiceClient(cc *grpc.ClientConn) ChatBotServiceClient {
	return &chatBotServiceClient{cc}
}

func (c *chatBotServiceClient) RegisterAppService(ctx context.Context, in *RegisterAppService, opts ...grpc.CallOption) (*ReplyRegisterAppService, error) {
	out := new(ReplyRegisterAppService)
	err := c.cc.Invoke(ctx, "/chatbotpb.ChatBotService/registerAppService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatBotServiceServer is the server API for ChatBotService service.
type ChatBotServiceServer interface {
	// registerAppService - register app service
	RegisterAppService(context.Context, *RegisterAppService) (*ReplyRegisterAppService, error)
}

func RegisterChatBotServiceServer(s *grpc.Server, srv ChatBotServiceServer) {
	s.RegisterService(&_ChatBotService_serviceDesc, srv)
}

func _ChatBotService_RegisterAppService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAppService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatBotServiceServer).RegisterAppService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatbotpb.ChatBotService/RegisterAppService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatBotServiceServer).RegisterAppService(ctx, req.(*RegisterAppService))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatBotService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chatbotpb.ChatBotService",
	HandlerType: (*ChatBotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "registerAppService",
			Handler:    _ChatBotService_RegisterAppService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chatbot.proto",
}

func init() { proto.RegisterFile("chatbot.proto", fileDescriptor_chatbot_bcb94d0b7321ef21) }

var fileDescriptor_chatbot_bcb94d0b7321ef21 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xd1, 0x6a, 0xea, 0x40,
	0x10, 0x86, 0xdd, 0x13, 0xf5, 0xe8, 0x78, 0x8e, 0x84, 0x41, 0x3c, 0x41, 0x38, 0x10, 0x72, 0x15,
	0xbc, 0xb0, 0xc5, 0x3e, 0x41, 0x2a, 0x52, 0x0a, 0x16, 0x71, 0x6b, 0x2f, 0x7a, 0x55, 0x56, 0xbb,
	0xad, 0x62, 0xcd, 0x0e, 0x9b, 0xb5, 0xd4, 0x27, 0xe8, 0x6b, 0x97, 0xdd, 0xa8, 0x44, 0xb4, 0xd0,
	0xbb, 0xf9, 0x33, 0x7f, 0xe6, 0xff, 0x66, 0x58, 0xf8, 0x3b, 0x5f, 0x08, 0x33, 0x53, 0xa6, 0x47,
	0x5a, 0x19, 0x85, 0xf5, 0x9d, 0xa4, 0x59, 0xf4, 0xc9, 0xa0, 0xf1, 0x90, 0x49, 0x9d, 0x10, 0xdd,
	0xa6, 0x2f, 0x0a, 0x63, 0xf0, 0x04, 0x51, 0xc0, 0x42, 0x16, 0x37, 0xfb, 0xed, 0xde, 0xc1, 0xd8,
	0x1b, 0x2c, 0x84, 0x49, 0x88, 0xa6, 0x5b, 0x92, 0xdc, 0x5a, 0xb0, 0x0d, 0x55, 0x41, 0xb4, 0x59,
	0x3e, 0x07, 0xbf, 0x42, 0x16, 0xd7, 0xf9, 0x4e, 0x61, 0x07, 0x6a, 0xb6, 0x4a, 0xc5, 0x5a, 0x06,
	0x9e, 0xeb, 0x1c, 0xb4, 0xed, 0xd9, 0x89, 0xe9, 0x66, 0x9d, 0x05, 0xe5, 0x90, 0xc5, 0x15, 0x7e,
	0xd0, 0x91, 0x86, 0x9a, 0x05, 0x71, 0x14, 0x3e, 0x78, 0x76, 0xb0, 0xa5, 0xf0, 0xb8, 0x2d, 0x11,
	0xa1, 0xec, 0x26, 0xe6, 0x59, 0xae, 0xc6, 0x2e, 0x94, 0x05, 0x51, 0x16, 0x78, 0xa1, 0x17, 0x37,
	0x8e, 0x60, 0x0b, 0x1b, 0x71, 0xe7, 0xc1, 0x16, 0x54, 0xd6, 0x2a, 0x95, 0xdb, 0x00, 0xdc, 0xcc,
	0x5c, 0x44, 0x1f, 0x80, 0x5c, 0xbe, 0x2e, 0x33, 0xe3, 0xec, 0xf7, 0x52, 0xbf, 0x2f, 0xe7, 0xd2,
	0x7a, 0x8d, 0x5a, 0xc9, 0xd4, 0xe5, 0xd7, 0x79, 0x2e, 0xf0, 0x12, 0x7e, 0x8b, 0x7c, 0x7f, 0x07,
	0xf1, 0xfd, 0x75, 0xf6, 0x36, 0xbb, 0xed, 0x26, 0x93, 0xba, 0x78, 0x89, 0xbd, 0x8e, 0x04, 0xfc,
	0xe3, 0x92, 0xde, 0xb6, 0x67, 0xe2, 0x0b, 0x41, 0xec, 0x67, 0x41, 0x2d, 0xa8, 0x48, 0xad, 0x95,
	0xde, 0x5d, 0x27, 0x17, 0xdd, 0x0b, 0x68, 0x14, 0xdc, 0xe8, 0xc3, 0x9f, 0x41, 0x32, 0x7d, 0x9a,
	0x0e, 0x47, 0xc3, 0x1b, 0x9e, 0xdc, 0xf9, 0x25, 0x6c, 0x02, 0xd8, 0x2f, 0x83, 0xf1, 0x78, 0x34,
	0x99, 0xf8, 0xac, 0xbf, 0x82, 0xa6, 0xfd, 0xe1, 0x5a, 0x99, 0x3d, 0xca, 0x23, 0xa0, 0x3e, 0x05,
	0xfc, 0x5f, 0xe0, 0x39, 0xe5, 0xef, 0x44, 0x47, 0xed, 0xb3, 0x3b, 0x46, 0xa5, 0x59, 0xd5, 0x3d,
	0xc5, 0xab, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xc2, 0x1b, 0xb9, 0x9b, 0x02, 0x00, 0x00,
}
