// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: notyfier.proto

package api_pb

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//GetNotificationRequest is a request for getting a chats by user ids
type GetNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid string `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
}

func (x *GetNotificationRequest) Reset() {
	*x = GetNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notyfier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationRequest) ProtoMessage() {}

func (x *GetNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notyfier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationRequest.ProtoReflect.Descriptor instead.
func (*GetNotificationRequest) Descriptor() ([]byte, []int) {
	return file_notyfier_proto_rawDescGZIP(), []int{0}
}

func (x *GetNotificationRequest) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

//GetNotificationResponse is a response for getting a chats by user ids
type GetNotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer string `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *GetNotificationResponse) Reset() {
	*x = GetNotificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notyfier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationResponse) ProtoMessage() {}

func (x *GetNotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notyfier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationResponse.ProtoReflect.Descriptor instead.
func (*GetNotificationResponse) Descriptor() ([]byte, []int) {
	return file_notyfier_proto_rawDescGZIP(), []int{1}
}

func (x *GetNotificationResponse) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

//DeleteChatsRequest is a request to delete chats
type DeleteChatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *DeleteChatsRequest) Reset() {
	*x = DeleteChatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notyfier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChatsRequest) ProtoMessage() {}

func (x *DeleteChatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notyfier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChatsRequest.ProtoReflect.Descriptor instead.
func (*DeleteChatsRequest) Descriptor() ([]byte, []int) {
	return file_notyfier_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteChatsRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

//DeleteChatsResponse is a response to delete chats
type DeleteChatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeletedChats []string `protobuf:"bytes,1,rep,name=deleted_chats,json=deletedChats,proto3" json:"deleted_chats,omitempty"`
}

func (x *DeleteChatsResponse) Reset() {
	*x = DeleteChatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notyfier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChatsResponse) ProtoMessage() {}

func (x *DeleteChatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notyfier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChatsResponse.ProtoReflect.Descriptor instead.
func (*DeleteChatsResponse) Descriptor() ([]byte, []int) {
	return file_notyfier_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteChatsResponse) GetDeletedChats() []string {
	if x != nil {
		return x.DeletedChats
	}
	return nil
}

//CreateChats
type CreateChatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExternalUserIdToInterval map[string]int32 `protobuf:"bytes,1,rep,name=ExternalUserIdToInterval,proto3" json:"ExternalUserIdToInterval,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *CreateChatsRequest) Reset() {
	*x = CreateChatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notyfier_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatsRequest) ProtoMessage() {}

func (x *CreateChatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notyfier_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatsRequest.ProtoReflect.Descriptor instead.
func (*CreateChatsRequest) Descriptor() ([]byte, []int) {
	return file_notyfier_proto_rawDescGZIP(), []int{4}
}

func (x *CreateChatsRequest) GetExternalUserIdToInterval() map[string]int32 {
	if x != nil {
		return x.ExternalUserIdToInterval
	}
	return nil
}

type CreateChatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExternalUserIdToChatUrl map[string]string `protobuf:"bytes,1,rep,name=ExternalUserIdToChatUrl,proto3" json:"ExternalUserIdToChatUrl,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CreateChatsResponse) Reset() {
	*x = CreateChatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notyfier_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatsResponse) ProtoMessage() {}

func (x *CreateChatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notyfier_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatsResponse.ProtoReflect.Descriptor instead.
func (*CreateChatsResponse) Descriptor() ([]byte, []int) {
	return file_notyfier_proto_rawDescGZIP(), []int{5}
}

func (x *CreateChatsResponse) GetExternalUserIdToChatUrl() map[string]string {
	if x != nil {
		return x.ExternalUserIdToChatUrl
	}
	return nil
}

type UserInfoDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExternalUserId string `protobuf:"bytes,1,opt,name=ExternalUserId,proto3" json:"ExternalUserId,omitempty"`
	Address        string `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
}

func (x *UserInfoDto) Reset() {
	*x = UserInfoDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notyfier_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoDto) ProtoMessage() {}

func (x *UserInfoDto) ProtoReflect() protoreflect.Message {
	mi := &file_notyfier_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoDto.ProtoReflect.Descriptor instead.
func (*UserInfoDto) Descriptor() ([]byte, []int) {
	return file_notyfier_proto_rawDescGZIP(), []int{6}
}

func (x *UserInfoDto) GetExternalUserId() string {
	if x != nil {
		return x.ExternalUserId
	}
	return ""
}

func (x *UserInfoDto) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_notyfier_proto protoreflect.FileDescriptor

var file_notyfier_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x79, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x61, 0x63, 0x68, 0x69, 0x72, 0x69, 0x6b, 0x6f, 0x76, 0x61, 0x2e, 0x65, 0x6d, 0x65,
	0x72, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x30, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x2f, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x3a, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x43, 0x68,
	0x61, 0x74, 0x73, 0x22, 0xe6, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x82, 0x01, 0x0a, 0x18, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x54, 0x6f, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x46, 0x2e,
	0x61, 0x63, 0x68, 0x69, 0x72, 0x69, 0x6b, 0x6f, 0x76, 0x61, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x54, 0x6f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x18, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x54, 0x6f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x1a,
	0x4b, 0x0a, 0x1d, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x54, 0x6f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe4, 0x01, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x80, 0x01, 0x0a, 0x17, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x54, 0x6f, 0x43, 0x68, 0x61, 0x74, 0x55, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x72, 0x69, 0x6b,
	0x6f, 0x76, 0x61, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x54, 0x6f, 0x43, 0x68, 0x61, 0x74, 0x55, 0x72, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x17,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x54, 0x6f,
	0x43, 0x68, 0x61, 0x74, 0x55, 0x72, 0x6c, 0x1a, 0x4a, 0x0a, 0x1c, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x54, 0x6f, 0x43, 0x68, 0x61, 0x74, 0x55,
	0x72, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x4f, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x44,
	0x74, 0x6f, 0x12, 0x26, 0x0a, 0x0e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x32, 0x9f, 0x03, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x79, 0x66, 0x69, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7b, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x73, 0x12, 0x28, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x72, 0x69,
	0x6b, 0x6f, 0x76, 0x61, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x72, 0x69, 0x6b, 0x6f, 0x76, 0x61, 0x2e, 0x65,
	0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x7b, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x74, 0x73, 0x12, 0x28, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x72, 0x69, 0x6b, 0x6f, 0x76,
	0x61, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x61, 0x63, 0x68, 0x69, 0x72, 0x69, 0x6b, 0x6f, 0x76, 0x61, 0x2e, 0x65, 0x6d, 0x65, 0x72,
	0x67, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x11, 0x22, 0x0c, 0x2f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x73, 0x3a,
	0x01, 0x2a, 0x12, 0x91, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x72, 0x69, 0x6b,
	0x6f, 0x76, 0x61, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x72, 0x69, 0x6b, 0x6f, 0x76,
	0x61, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x47, 0x65,
	0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x75,
	0x73, 0x65, 0x72, 0x69, 0x64, 0x7d, 0x42, 0x16, 0x5a, 0x14, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65,
	0x6e, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notyfier_proto_rawDescOnce sync.Once
	file_notyfier_proto_rawDescData = file_notyfier_proto_rawDesc
)

func file_notyfier_proto_rawDescGZIP() []byte {
	file_notyfier_proto_rawDescOnce.Do(func() {
		file_notyfier_proto_rawDescData = protoimpl.X.CompressGZIP(file_notyfier_proto_rawDescData)
	})
	return file_notyfier_proto_rawDescData
}

var file_notyfier_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_notyfier_proto_goTypes = []interface{}{
	(*GetNotificationRequest)(nil),  // 0: achirikova.emergence.GetNotificationRequest
	(*GetNotificationResponse)(nil), // 1: achirikova.emergence.GetNotificationResponse
	(*DeleteChatsRequest)(nil),      // 2: achirikova.emergence.DeleteChatsRequest
	(*DeleteChatsResponse)(nil),     // 3: achirikova.emergence.DeleteChatsResponse
	(*CreateChatsRequest)(nil),      // 4: achirikova.emergence.CreateChatsRequest
	(*CreateChatsResponse)(nil),     // 5: achirikova.emergence.CreateChatsResponse
	(*UserInfoDto)(nil),             // 6: achirikova.emergence.UserInfoDto
	nil,                             // 7: achirikova.emergence.CreateChatsRequest.ExternalUserIdToIntervalEntry
	nil,                             // 8: achirikova.emergence.CreateChatsResponse.ExternalUserIdToChatUrlEntry
}
var file_notyfier_proto_depIdxs = []int32{
	7, // 0: achirikova.emergence.CreateChatsRequest.ExternalUserIdToInterval:type_name -> achirikova.emergence.CreateChatsRequest.ExternalUserIdToIntervalEntry
	8, // 1: achirikova.emergence.CreateChatsResponse.ExternalUserIdToChatUrl:type_name -> achirikova.emergence.CreateChatsResponse.ExternalUserIdToChatUrlEntry
	4, // 2: achirikova.emergence.NotyfierService.CreateChats:input_type -> achirikova.emergence.CreateChatsRequest
	2, // 3: achirikova.emergence.NotyfierService.DeleteChats:input_type -> achirikova.emergence.DeleteChatsRequest
	0, // 4: achirikova.emergence.NotyfierService.GetNotification:input_type -> achirikova.emergence.GetNotificationRequest
	5, // 5: achirikova.emergence.NotyfierService.CreateChats:output_type -> achirikova.emergence.CreateChatsResponse
	3, // 6: achirikova.emergence.NotyfierService.DeleteChats:output_type -> achirikova.emergence.DeleteChatsResponse
	1, // 7: achirikova.emergence.NotyfierService.GetNotification:output_type -> achirikova.emergence.GetNotificationResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_notyfier_proto_init() }
func file_notyfier_proto_init() {
	if File_notyfier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notyfier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotificationRequest); i {
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
		file_notyfier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotificationResponse); i {
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
		file_notyfier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteChatsRequest); i {
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
		file_notyfier_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteChatsResponse); i {
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
		file_notyfier_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChatsRequest); i {
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
		file_notyfier_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChatsResponse); i {
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
		file_notyfier_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoDto); i {
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
			RawDescriptor: file_notyfier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notyfier_proto_goTypes,
		DependencyIndexes: file_notyfier_proto_depIdxs,
		MessageInfos:      file_notyfier_proto_msgTypes,
	}.Build()
	File_notyfier_proto = out.File
	file_notyfier_proto_rawDesc = nil
	file_notyfier_proto_goTypes = nil
	file_notyfier_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NotyfierServiceClient is the client API for NotyfierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotyfierServiceClient interface {
	CreateChats(ctx context.Context, in *CreateChatsRequest, opts ...grpc.CallOption) (*CreateChatsResponse, error)
	// RPC for deleting a chats by external user ids
	DeleteChats(ctx context.Context, in *DeleteChatsRequest, opts ...grpc.CallOption) (*DeleteChatsResponse, error)
	// get rpc for getting a chats by user ids
	GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...grpc.CallOption) (*GetNotificationResponse, error)
}

type notyfierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotyfierServiceClient(cc grpc.ClientConnInterface) NotyfierServiceClient {
	return &notyfierServiceClient{cc}
}

func (c *notyfierServiceClient) CreateChats(ctx context.Context, in *CreateChatsRequest, opts ...grpc.CallOption) (*CreateChatsResponse, error) {
	out := new(CreateChatsResponse)
	err := c.cc.Invoke(ctx, "/achirikova.emergence.NotyfierService/CreateChats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notyfierServiceClient) DeleteChats(ctx context.Context, in *DeleteChatsRequest, opts ...grpc.CallOption) (*DeleteChatsResponse, error) {
	out := new(DeleteChatsResponse)
	err := c.cc.Invoke(ctx, "/achirikova.emergence.NotyfierService/DeleteChats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notyfierServiceClient) GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...grpc.CallOption) (*GetNotificationResponse, error) {
	out := new(GetNotificationResponse)
	err := c.cc.Invoke(ctx, "/achirikova.emergence.NotyfierService/GetNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotyfierServiceServer is the server API for NotyfierService service.
type NotyfierServiceServer interface {
	CreateChats(context.Context, *CreateChatsRequest) (*CreateChatsResponse, error)
	// RPC for deleting a chats by external user ids
	DeleteChats(context.Context, *DeleteChatsRequest) (*DeleteChatsResponse, error)
	// get rpc for getting a chats by user ids
	GetNotification(context.Context, *GetNotificationRequest) (*GetNotificationResponse, error)
}

// UnimplementedNotyfierServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNotyfierServiceServer struct {
}

func (*UnimplementedNotyfierServiceServer) CreateChats(context.Context, *CreateChatsRequest) (*CreateChatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChats not implemented")
}
func (*UnimplementedNotyfierServiceServer) DeleteChats(context.Context, *DeleteChatsRequest) (*DeleteChatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChats not implemented")
}
func (*UnimplementedNotyfierServiceServer) GetNotification(context.Context, *GetNotificationRequest) (*GetNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotification not implemented")
}

func RegisterNotyfierServiceServer(s *grpc.Server, srv NotyfierServiceServer) {
	s.RegisterService(&_NotyfierService_serviceDesc, srv)
}

func _NotyfierService_CreateChats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotyfierServiceServer).CreateChats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/achirikova.emergence.NotyfierService/CreateChats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotyfierServiceServer).CreateChats(ctx, req.(*CreateChatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotyfierService_DeleteChats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotyfierServiceServer).DeleteChats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/achirikova.emergence.NotyfierService/DeleteChats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotyfierServiceServer).DeleteChats(ctx, req.(*DeleteChatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotyfierService_GetNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotyfierServiceServer).GetNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/achirikova.emergence.NotyfierService/GetNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotyfierServiceServer).GetNotification(ctx, req.(*GetNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotyfierService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "achirikova.emergence.NotyfierService",
	HandlerType: (*NotyfierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChats",
			Handler:    _NotyfierService_CreateChats_Handler,
		},
		{
			MethodName: "DeleteChats",
			Handler:    _NotyfierService_DeleteChats_Handler,
		},
		{
			MethodName: "GetNotification",
			Handler:    _NotyfierService_GetNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notyfier.proto",
}
