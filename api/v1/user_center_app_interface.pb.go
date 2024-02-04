// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: user_center_app_interface.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/limes-cloud/resource/api/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AppInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AppId       uint32          `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	ParentId    uint32          `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Type        string          `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Title       string          `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Path        string          `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
	Method      string          `protobuf:"bytes,7,opt,name=method,proto3" json:"method,omitempty"`
	Description string          `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   uint32          `protobuf:"varint,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   uint32          `protobuf:"varint,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Children    []*AppInterface `protobuf:"bytes,11,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *AppInterface) Reset() {
	*x = AppInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_app_interface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppInterface) ProtoMessage() {}

func (x *AppInterface) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_app_interface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppInterface.ProtoReflect.Descriptor instead.
func (*AppInterface) Descriptor() ([]byte, []int) {
	return file_user_center_app_interface_proto_rawDescGZIP(), []int{0}
}

func (x *AppInterface) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AppInterface) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *AppInterface) GetParentId() uint32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *AppInterface) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AppInterface) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AppInterface) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *AppInterface) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *AppInterface) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AppInterface) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *AppInterface) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *AppInterface) GetChildren() []*AppInterface {
	if x != nil {
		return x.Children
	}
	return nil
}

type GetAppInterfaceTreeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId uint32 `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *GetAppInterfaceTreeRequest) Reset() {
	*x = GetAppInterfaceTreeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_app_interface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppInterfaceTreeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppInterfaceTreeRequest) ProtoMessage() {}

func (x *GetAppInterfaceTreeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_app_interface_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppInterfaceTreeRequest.ProtoReflect.Descriptor instead.
func (*GetAppInterfaceTreeRequest) Descriptor() ([]byte, []int) {
	return file_user_center_app_interface_proto_rawDescGZIP(), []int{1}
}

func (x *GetAppInterfaceTreeRequest) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

type GetAppInterfaceTreeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*AppInterface `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *GetAppInterfaceTreeReply) Reset() {
	*x = GetAppInterfaceTreeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_app_interface_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppInterfaceTreeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppInterfaceTreeReply) ProtoMessage() {}

func (x *GetAppInterfaceTreeReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_app_interface_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppInterfaceTreeReply.ProtoReflect.Descriptor instead.
func (*GetAppInterfaceTreeReply) Descriptor() ([]byte, []int) {
	return file_user_center_app_interface_proto_rawDescGZIP(), []int{2}
}

func (x *GetAppInterfaceTreeReply) GetList() []*AppInterface {
	if x != nil {
		return x.List
	}
	return nil
}

type AddAppInterfaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId       uint32  `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	ParentId    uint32  `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Type        string  `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Title       string  `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Path        *string `protobuf:"bytes,5,opt,name=path,proto3,oneof" json:"path,omitempty"`
	Method      *string `protobuf:"bytes,6,opt,name=method,proto3,oneof" json:"method,omitempty"`
	Description *string `protobuf:"bytes,7,opt,name=description,proto3,oneof" json:"description,omitempty"`
}

func (x *AddAppInterfaceRequest) Reset() {
	*x = AddAppInterfaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_app_interface_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAppInterfaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAppInterfaceRequest) ProtoMessage() {}

func (x *AddAppInterfaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_app_interface_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAppInterfaceRequest.ProtoReflect.Descriptor instead.
func (*AddAppInterfaceRequest) Descriptor() ([]byte, []int) {
	return file_user_center_app_interface_proto_rawDescGZIP(), []int{3}
}

func (x *AddAppInterfaceRequest) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *AddAppInterfaceRequest) GetParentId() uint32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *AddAppInterfaceRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AddAppInterfaceRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddAppInterfaceRequest) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *AddAppInterfaceRequest) GetMethod() string {
	if x != nil && x.Method != nil {
		return *x.Method
	}
	return ""
}

func (x *AddAppInterfaceRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

type AddAppInterfaceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddAppInterfaceReply) Reset() {
	*x = AddAppInterfaceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_app_interface_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAppInterfaceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAppInterfaceReply) ProtoMessage() {}

func (x *AddAppInterfaceReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_app_interface_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAppInterfaceReply.ProtoReflect.Descriptor instead.
func (*AddAppInterfaceReply) Descriptor() ([]byte, []int) {
	return file_user_center_app_interface_proto_rawDescGZIP(), []int{4}
}

func (x *AddAppInterfaceReply) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateAppInterfaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId    uint32  `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Type        string  `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Title       string  `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Path        *string `protobuf:"bytes,5,opt,name=path,proto3,oneof" json:"path,omitempty"`
	Method      *string `protobuf:"bytes,6,opt,name=method,proto3,oneof" json:"method,omitempty"`
	Description *string `protobuf:"bytes,7,opt,name=description,proto3,oneof" json:"description,omitempty"`
}

func (x *UpdateAppInterfaceRequest) Reset() {
	*x = UpdateAppInterfaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_app_interface_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAppInterfaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAppInterfaceRequest) ProtoMessage() {}

func (x *UpdateAppInterfaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_app_interface_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAppInterfaceRequest.ProtoReflect.Descriptor instead.
func (*UpdateAppInterfaceRequest) Descriptor() ([]byte, []int) {
	return file_user_center_app_interface_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateAppInterfaceRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateAppInterfaceRequest) GetParentId() uint32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *UpdateAppInterfaceRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UpdateAppInterfaceRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateAppInterfaceRequest) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *UpdateAppInterfaceRequest) GetMethod() string {
	if x != nil && x.Method != nil {
		return *x.Method
	}
	return ""
}

func (x *UpdateAppInterfaceRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

type DeleteAppInterfaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAppInterfaceRequest) Reset() {
	*x = DeleteAppInterfaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_app_interface_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAppInterfaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAppInterfaceRequest) ProtoMessage() {}

func (x *DeleteAppInterfaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_app_interface_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAppInterfaceRequest.ProtoReflect.Descriptor instead.
func (*DeleteAppInterfaceRequest) Descriptor() ([]byte, []int) {
	return file_user_center_app_interface_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteAppInterfaceRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_user_center_app_interface_proto protoreflect.FileDescriptor

var file_user_center_app_interface_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x70,
	0x70, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x02, 0x0a, 0x0c, 0x41, 0x70,
	0x70, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e,
	0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x3c, 0x0a, 0x1a, 0x47,
	0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x54, 0x72,
	0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x06, 0x61, 0x70, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02,
	0x20, 0x00, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x41, 0x70, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x54, 0x72, 0x65, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2d, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0x9b, 0x02, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12,
	0x24, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x00, 0x52, 0x08, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x26, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x97, 0x02, 0x0a, 0x19, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x24, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x00, 0x52, 0x08, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x34, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70,
	0x70, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_center_app_interface_proto_rawDescOnce sync.Once
	file_user_center_app_interface_proto_rawDescData = file_user_center_app_interface_proto_rawDesc
)

func file_user_center_app_interface_proto_rawDescGZIP() []byte {
	file_user_center_app_interface_proto_rawDescOnce.Do(func() {
		file_user_center_app_interface_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_center_app_interface_proto_rawDescData)
	})
	return file_user_center_app_interface_proto_rawDescData
}

var file_user_center_app_interface_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_user_center_app_interface_proto_goTypes = []interface{}{
	(*AppInterface)(nil),               // 0: user_center.AppInterface
	(*GetAppInterfaceTreeRequest)(nil), // 1: user_center.GetAppInterfaceTreeRequest
	(*GetAppInterfaceTreeReply)(nil),   // 2: user_center.GetAppInterfaceTreeReply
	(*AddAppInterfaceRequest)(nil),     // 3: user_center.AddAppInterfaceRequest
	(*AddAppInterfaceReply)(nil),       // 4: user_center.AddAppInterfaceReply
	(*UpdateAppInterfaceRequest)(nil),  // 5: user_center.UpdateAppInterfaceRequest
	(*DeleteAppInterfaceRequest)(nil),  // 6: user_center.DeleteAppInterfaceRequest
}
var file_user_center_app_interface_proto_depIdxs = []int32{
	0, // 0: user_center.AppInterface.children:type_name -> user_center.AppInterface
	0, // 1: user_center.GetAppInterfaceTreeReply.list:type_name -> user_center.AppInterface
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_center_app_interface_proto_init() }
func file_user_center_app_interface_proto_init() {
	if File_user_center_app_interface_proto != nil {
		return
	}
	file_user_center_channel_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_center_app_interface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppInterface); i {
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
		file_user_center_app_interface_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppInterfaceTreeRequest); i {
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
		file_user_center_app_interface_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppInterfaceTreeReply); i {
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
		file_user_center_app_interface_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAppInterfaceRequest); i {
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
		file_user_center_app_interface_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAppInterfaceReply); i {
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
		file_user_center_app_interface_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAppInterfaceRequest); i {
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
		file_user_center_app_interface_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAppInterfaceRequest); i {
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
	file_user_center_app_interface_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_user_center_app_interface_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_center_app_interface_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_center_app_interface_proto_goTypes,
		DependencyIndexes: file_user_center_app_interface_proto_depIdxs,
		MessageInfos:      file_user_center_app_interface_proto_msgTypes,
	}.Build()
	File_user_center_app_interface_proto = out.File
	file_user_center_app_interface_proto_rawDesc = nil
	file_user_center_app_interface_proto_goTypes = nil
	file_user_center_app_interface_proto_depIdxs = nil
}