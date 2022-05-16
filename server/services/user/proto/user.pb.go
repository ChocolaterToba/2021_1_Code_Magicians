// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: user.proto

package proto

import (
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

type Role int32

const (
	Role_Undefined Role = 0
	Role_Client    Role = 1
	Role_Manager   Role = 2
	Role_Admin     Role = 3
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "Undefined",
		1: "Client",
		2: "Manager",
		3: "Admin",
	}
	Role_value = map[string]int32{
		"Undefined": 0,
		"Client":    1,
		"Manager":   2,
		"Admin":     3,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_user_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_user_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

type UserReg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Email     string `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *UserReg) Reset() {
	*x = UserReg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReg) ProtoMessage() {}

func (x *UserReg) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReg.ProtoReflect.Descriptor instead.
func (*UserReg) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserReg) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserReg) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserReg) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserReg) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserReg) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UserEditInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    uint64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Email     string `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *UserEditInput) Reset() {
	*x = UserEditInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEditInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEditInput) ProtoMessage() {}

func (x *UserEditInput) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEditInput.ProtoReflect.Descriptor instead.
func (*UserEditInput) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserEditInput) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UserEditInput) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserEditInput) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserEditInput) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserEditInput) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

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
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAuth) ProtoMessage() {}

func (x *UserAuth) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
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
	return file_user_proto_rawDescGZIP(), []int{2}
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

type UserOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID     uint64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Username   string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Email      string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	AvatarPath string `protobuf:"bytes,4,opt,name=AvatarPath,proto3" json:"AvatarPath,omitempty"`
	FirstName  string `protobuf:"bytes,5,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName   string `protobuf:"bytes,6,opt,name=LastName,proto3" json:"LastName,omitempty"`
}

func (x *UserOutput) Reset() {
	*x = UserOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOutput) ProtoMessage() {}

func (x *UserOutput) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserOutput.ProtoReflect.Descriptor instead.
func (*UserOutput) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserOutput) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UserOutput) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserOutput) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserOutput) GetAvatarPath() string {
	if x != nil {
		return x.AvatarPath
	}
	return ""
}

func (x *UserOutput) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserOutput) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type UsersListOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserOutput `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
}

func (x *UsersListOutput) Reset() {
	*x = UsersListOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersListOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersListOutput) ProtoMessage() {}

func (x *UsersListOutput) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersListOutput.ProtoReflect.Descriptor instead.
func (*UsersListOutput) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *UsersListOutput) GetUsers() []*UserOutput {
	if x != nil {
		return x.Users
	}
	return nil
}

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
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
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserID) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type Username struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *Username) Reset() {
	*x = Username{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Username) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Username) ProtoMessage() {}

func (x *Username) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Username.ProtoReflect.Descriptor instead.
func (*Username) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *Username) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type SearchInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyWords string `protobuf:"bytes,1,opt,name=keyWords,proto3" json:"keyWords,omitempty"`
}

func (x *SearchInput) Reset() {
	*x = SearchInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchInput) ProtoMessage() {}

func (x *SearchInput) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchInput.ProtoReflect.Descriptor instead.
func (*SearchInput) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *SearchInput) GetKeyWords() string {
	if x != nil {
		return x.KeyWords
	}
	return ""
}

type UpdateAvatarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*UpdateAvatarRequest_UserId
	//	*UpdateAvatarRequest_Filename
	//	*UpdateAvatarRequest_Chunk
	Data isUpdateAvatarRequest_Data `protobuf_oneof:"data"`
}

func (x *UpdateAvatarRequest) Reset() {
	*x = UpdateAvatarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAvatarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAvatarRequest) ProtoMessage() {}

func (x *UpdateAvatarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAvatarRequest.ProtoReflect.Descriptor instead.
func (*UpdateAvatarRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (m *UpdateAvatarRequest) GetData() isUpdateAvatarRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UpdateAvatarRequest) GetUserId() uint64 {
	if x, ok := x.GetData().(*UpdateAvatarRequest_UserId); ok {
		return x.UserId
	}
	return 0
}

func (x *UpdateAvatarRequest) GetFilename() string {
	if x, ok := x.GetData().(*UpdateAvatarRequest_Filename); ok {
		return x.Filename
	}
	return ""
}

func (x *UpdateAvatarRequest) GetChunk() []byte {
	if x, ok := x.GetData().(*UpdateAvatarRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isUpdateAvatarRequest_Data interface {
	isUpdateAvatarRequest_Data()
}

type UpdateAvatarRequest_UserId struct {
	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3,oneof"`
}

type UpdateAvatarRequest_Filename struct {
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3,oneof"`
}

type UpdateAvatarRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,3,opt,name=chunk,proto3,oneof"`
}

func (*UpdateAvatarRequest_UserId) isUpdateAvatarRequest_Data() {}

func (*UpdateAvatarRequest_Filename) isUpdateAvatarRequest_Data() {}

func (*UpdateAvatarRequest_Chunk) isUpdateAvatarRequest_Data() {}

type AddRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Role   Role   `protobuf:"varint,2,opt,name=role,proto3,enum=user.Role" json:"role,omitempty"`
}

func (x *AddRoleRequest) Reset() {
	*x = AddRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoleRequest) ProtoMessage() {}

func (x *AddRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoleRequest.ProtoReflect.Descriptor instead.
func (*AddRoleRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{9}
}

func (x *AddRoleRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddRoleRequest) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_Undefined
}

type GetRolesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roles []Role `protobuf:"varint,1,rep,packed,name=roles,proto3,enum=user.Role" json:"roles,omitempty"`
}

func (x *GetRolesResponse) Reset() {
	*x = GetRolesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRolesResponse) ProtoMessage() {}

func (x *GetRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRolesResponse.ProtoReflect.Descriptor instead.
func (*GetRolesResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{10}
}

func (x *GetRolesResponse) GetRoles() []Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{11}
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x91, 0x01, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x12, 0x1a,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x93, 0x01, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x45,
	0x64, 0x69, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x42, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0xb0, 0x01, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x26, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x1a,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x26, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x6e, 0x0a,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x49, 0x0a,
	0x0e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x34, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x2a, 0x39, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x0d, 0x0a, 0x09, 0x55, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x10, 0x03, 0x32, 0x9f, 0x03, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x1a, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x08, 0x45, 0x64, 0x69, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x45, 0x64, 0x69, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x28, 0x01, 0x12, 0x2f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x44, 0x12, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x1a, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x10, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x30,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00,
	0x12, 0x32, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x0c, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x1f, 0x5a, 0x1d, 0x70, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73,
	0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_user_proto_goTypes = []interface{}{
	(Role)(0),                   // 0: user.Role
	(*UserReg)(nil),             // 1: user.UserReg
	(*UserEditInput)(nil),       // 2: user.UserEditInput
	(*UserAuth)(nil),            // 3: user.UserAuth
	(*UserOutput)(nil),          // 4: user.UserOutput
	(*UsersListOutput)(nil),     // 5: user.UsersListOutput
	(*UserID)(nil),              // 6: user.UserID
	(*Username)(nil),            // 7: user.Username
	(*SearchInput)(nil),         // 8: user.SearchInput
	(*UpdateAvatarRequest)(nil), // 9: user.UpdateAvatarRequest
	(*AddRoleRequest)(nil),      // 10: user.AddRoleRequest
	(*GetRolesResponse)(nil),    // 11: user.GetRolesResponse
	(*Empty)(nil),               // 12: user.Empty
}
var file_user_proto_depIdxs = []int32{
	4,  // 0: user.UsersListOutput.Users:type_name -> user.UserOutput
	0,  // 1: user.AddRoleRequest.role:type_name -> user.Role
	0,  // 2: user.GetRolesResponse.roles:type_name -> user.Role
	1,  // 3: user.User.CreateUser:input_type -> user.UserReg
	2,  // 4: user.User.EditUser:input_type -> user.UserEditInput
	9,  // 5: user.User.UpdateAvatar:input_type -> user.UpdateAvatarRequest
	6,  // 6: user.User.GetUserByID:input_type -> user.UserID
	7,  // 7: user.User.GetUserByUsername:input_type -> user.Username
	12, // 8: user.User.GetUsers:input_type -> user.Empty
	6,  // 9: user.User.GetRoles:input_type -> user.UserID
	10, // 10: user.User.AddRole:input_type -> user.AddRoleRequest
	6,  // 11: user.User.CreateUser:output_type -> user.UserID
	12, // 12: user.User.EditUser:output_type -> user.Empty
	12, // 13: user.User.UpdateAvatar:output_type -> user.Empty
	4,  // 14: user.User.GetUserByID:output_type -> user.UserOutput
	4,  // 15: user.User.GetUserByUsername:output_type -> user.UserOutput
	5,  // 16: user.User.GetUsers:output_type -> user.UsersListOutput
	11, // 17: user.User.GetRoles:output_type -> user.GetRolesResponse
	12, // 18: user.User.AddRole:output_type -> user.Empty
	11, // [11:19] is the sub-list for method output_type
	3,  // [3:11] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReg); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEditInput); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserOutput); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersListOutput); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Username); i {
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
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchInput); i {
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
		file_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAvatarRequest); i {
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
		file_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRoleRequest); i {
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
		file_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRolesResponse); i {
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
		file_user_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
	file_user_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*UpdateAvatarRequest_UserId)(nil),
		(*UpdateAvatarRequest_Filename)(nil),
		(*UpdateAvatarRequest_Chunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		EnumInfos:         file_user_proto_enumTypes,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
