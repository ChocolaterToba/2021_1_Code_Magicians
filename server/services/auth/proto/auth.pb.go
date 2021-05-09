// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: auth.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *UserAuth) Reset() {
	*x = UserAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAuth) ProtoMessage() {}

func (x *UserAuth) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAuth.ProtoReflect.Descriptor instead.
func (*UserAuth) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *UserAuth) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserAuth) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *UserID) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type Cookie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionValue string `protobuf:"bytes,1,opt,name=sessionValue,proto3" json:"sessionValue,omitempty"`
}

func (x *Cookie) Reset() {
	*x = Cookie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cookie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cookie) ProtoMessage() {}

func (x *Cookie) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cookie.ProtoReflect.Descriptor instead.
func (*Cookie) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *Cookie) GetSessionValue() string {
	if x != nil {
		return x.SessionValue
	}
	return ""
}

type CookieInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID       int32   `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	SessionValue *Cookie `protobuf:"bytes,2,opt,name=sessionValue,proto3" json:"sessionValue,omitempty"`
}

func (x *CookieInfo) Reset() {
	*x = CookieInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CookieInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CookieInfo) ProtoMessage() {}

func (x *CookieInfo) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CookieInfo.ProtoReflect.Descriptor instead.
func (*CookieInfo) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

func (x *CookieInfo) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CookieInfo) GetSessionValue() *Cookie {
	if x != nil {
		return x.SessionValue
	}
	return nil
}

type CheckCookieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CookieInfo *CookieInfo `protobuf:"bytes,1,opt,name=cookieInfo,proto3" json:"cookieInfo,omitempty"`
	IsCookie   bool        `protobuf:"varint,2,opt,name=isCookie,proto3" json:"isCookie,omitempty"`
}

func (x *CheckCookieResponse) Reset() {
	*x = CheckCookieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckCookieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckCookieResponse) ProtoMessage() {}

func (x *CheckCookieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckCookieResponse.ProtoReflect.Descriptor instead.
func (*CheckCookieResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{4}
}

func (x *CheckCookieResponse) GetCookieInfo() *CookieInfo {
	if x != nil {
		return x.CookieInfo
	}
	return nil
}

func (x *CheckCookieResponse) GetIsCookie() bool {
	if x != nil {
		return x.IsCookie
	}
	return false
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{5}
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x22, 0x42, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1a,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x1a, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x22, 0x2c, 0x0a, 0x06, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x56, 0x0a, 0x0a, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x63, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x0a, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x69,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x9c, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x2f, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x10, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00,
	0x12, 0x29, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0c,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x0b, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0b, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x0c, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_auth_proto_goTypes = []interface{}{
	(*UserAuth)(nil),            // 0: auth.UserAuth
	(*UserID)(nil),              // 1: auth.UserID
	(*Cookie)(nil),              // 2: auth.Cookie
	(*CookieInfo)(nil),          // 3: auth.CookieInfo
	(*CheckCookieResponse)(nil), // 4: auth.CheckCookieResponse
	(*Error)(nil),               // 5: auth.Error
}
var file_auth_proto_depIdxs = []int32{
	2, // 0: auth.CookieInfo.sessionValue:type_name -> auth.Cookie
	3, // 1: auth.CheckCookieResponse.cookieInfo:type_name -> auth.CookieInfo
	0, // 2: auth.Auth.LoginUser:input_type -> auth.UserAuth
	1, // 3: auth.Auth.LogoutUser:input_type -> auth.UserID
	2, // 4: auth.Auth.CheckCookie:input_type -> auth.Cookie
	3, // 5: auth.Auth.LoginUser:output_type -> auth.CookieInfo
	5, // 6: auth.Auth.LogoutUser:output_type -> auth.Error
	4, // 7: auth.Auth.CheckCookie:output_type -> auth.CheckCookieResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAuth); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cookie); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CookieInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckCookieResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	LoginUser(ctx context.Context, in *UserAuth, opts ...grpc.CallOption) (*CookieInfo, error)
	LogoutUser(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Error, error)
	CheckCookie(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*CheckCookieResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) LoginUser(ctx context.Context, in *UserAuth, opts ...grpc.CallOption) (*CookieInfo, error) {
	out := new(CookieInfo)
	err := c.cc.Invoke(ctx, "/auth.Auth/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) LogoutUser(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/auth.Auth/LogoutUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CheckCookie(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*CheckCookieResponse, error) {
	out := new(CheckCookieResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/CheckCookie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	LoginUser(context.Context, *UserAuth) (*CookieInfo, error)
	LogoutUser(context.Context, *UserID) (*Error, error)
	CheckCookie(context.Context, *Cookie) (*CheckCookieResponse, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) LoginUser(context.Context, *UserAuth) (*CookieInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (*UnimplementedAuthServer) LogoutUser(context.Context, *UserID) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogoutUser not implemented")
}
func (*UnimplementedAuthServer) CheckCookie(context.Context, *Cookie) (*CheckCookieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckCookie not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).LoginUser(ctx, req.(*UserAuth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_LogoutUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).LogoutUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/LogoutUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).LogoutUser(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CheckCookie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cookie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CheckCookie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/CheckCookie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CheckCookie(ctx, req.(*Cookie))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginUser",
			Handler:    _Auth_LoginUser_Handler,
		},
		{
			MethodName: "LogoutUser",
			Handler:    _Auth_LogoutUser_Handler,
		},
		{
			MethodName: "CheckCookie",
			Handler:    _Auth_CheckCookie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
