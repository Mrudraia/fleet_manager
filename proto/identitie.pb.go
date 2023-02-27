// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: proto/identitie.proto

package fleet_manager

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

type IdentitieIdentityType int32

const (
	Identitie_AWS    IdentitieIdentityType = 0
	Identitie_RedHat IdentitieIdentityType = 1
)

// Enum value maps for IdentitieIdentityType.
var (
	IdentitieIdentityType_name = map[int32]string{
		0: "AWS",
		1: "RedHat",
	}
	IdentitieIdentityType_value = map[string]int32{
		"AWS":    0,
		"RedHat": 1,
	}
)

func (x IdentitieIdentityType) Enum() *IdentitieIdentityType {
	p := new(IdentitieIdentityType)
	*p = x
	return p
}

func (x IdentitieIdentityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IdentitieIdentityType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_identitie_proto_enumTypes[0].Descriptor()
}

func (IdentitieIdentityType) Type() protoreflect.EnumType {
	return &file_proto_identitie_proto_enumTypes[0]
}

func (x IdentitieIdentityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IdentitieIdentityType.Descriptor instead.
func (IdentitieIdentityType) EnumDescriptor() ([]byte, []int) {
	return file_proto_identitie_proto_rawDescGZIP(), []int{0, 0}
}

type Identitie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *Identitie) Reset() {
	*x = Identitie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identitie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identitie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identitie) ProtoMessage() {}

func (x *Identitie) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identitie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identitie.ProtoReflect.Descriptor instead.
func (*Identitie) Descriptor() ([]byte, []int) {
	return file_proto_identitie_proto_rawDescGZIP(), []int{0}
}

func (x *Identitie) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type CreateIdentitieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identitie *Identitie `protobuf:"bytes,1,opt,name=identitie,proto3" json:"identitie,omitempty"`
}

func (x *CreateIdentitieRequest) Reset() {
	*x = CreateIdentitieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identitie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIdentitieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIdentitieRequest) ProtoMessage() {}

func (x *CreateIdentitieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identitie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIdentitieRequest.ProtoReflect.Descriptor instead.
func (*CreateIdentitieRequest) Descriptor() ([]byte, []int) {
	return file_proto_identitie_proto_rawDescGZIP(), []int{1}
}

func (x *CreateIdentitieRequest) GetIdentitie() *Identitie {
	if x != nil {
		return x.Identitie
	}
	return nil
}

type CreateIdentitieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identitie *Identitie `protobuf:"bytes,1,opt,name=identitie,proto3" json:"identitie,omitempty"`
}

func (x *CreateIdentitieResponse) Reset() {
	*x = CreateIdentitieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identitie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIdentitieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIdentitieResponse) ProtoMessage() {}

func (x *CreateIdentitieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identitie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIdentitieResponse.ProtoReflect.Descriptor instead.
func (*CreateIdentitieResponse) Descriptor() ([]byte, []int) {
	return file_proto_identitie_proto_rawDescGZIP(), []int{2}
}

func (x *CreateIdentitieResponse) GetIdentitie() *Identitie {
	if x != nil {
		return x.Identitie
	}
	return nil
}

type ReadIdentitieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *ReadIdentitieRequest) Reset() {
	*x = ReadIdentitieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identitie_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadIdentitieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadIdentitieRequest) ProtoMessage() {}

func (x *ReadIdentitieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identitie_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadIdentitieRequest.ProtoReflect.Descriptor instead.
func (*ReadIdentitieRequest) Descriptor() ([]byte, []int) {
	return file_proto_identitie_proto_rawDescGZIP(), []int{3}
}

func (x *ReadIdentitieRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type ReadIdentitieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identitie *Identitie `protobuf:"bytes,1,opt,name=identitie,proto3" json:"identitie,omitempty"`
}

func (x *ReadIdentitieResponse) Reset() {
	*x = ReadIdentitieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identitie_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadIdentitieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadIdentitieResponse) ProtoMessage() {}

func (x *ReadIdentitieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identitie_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadIdentitieResponse.ProtoReflect.Descriptor instead.
func (*ReadIdentitieResponse) Descriptor() ([]byte, []int) {
	return file_proto_identitie_proto_rawDescGZIP(), []int{4}
}

func (x *ReadIdentitieResponse) GetIdentitie() *Identitie {
	if x != nil {
		return x.Identitie
	}
	return nil
}

type ReadIdentitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReadIdentitiesRequest) Reset() {
	*x = ReadIdentitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identitie_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadIdentitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadIdentitiesRequest) ProtoMessage() {}

func (x *ReadIdentitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identitie_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadIdentitiesRequest.ProtoReflect.Descriptor instead.
func (*ReadIdentitiesRequest) Descriptor() ([]byte, []int) {
	return file_proto_identitie_proto_rawDescGZIP(), []int{5}
}

type ReadIdentitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identitie []*Identitie `protobuf:"bytes,1,rep,name=identitie,proto3" json:"identitie,omitempty"`
}

func (x *ReadIdentitiesResponse) Reset() {
	*x = ReadIdentitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identitie_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadIdentitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadIdentitiesResponse) ProtoMessage() {}

func (x *ReadIdentitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identitie_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadIdentitiesResponse.ProtoReflect.Descriptor instead.
func (*ReadIdentitiesResponse) Descriptor() ([]byte, []int) {
	return file_proto_identitie_proto_rawDescGZIP(), []int{6}
}

func (x *ReadIdentitiesResponse) GetIdentitie() []*Identitie {
	if x != nil {
		return x.Identitie
	}
	return nil
}

type UpdateIdentitieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identitie *Identitie `protobuf:"bytes,1,opt,name=Identitie,proto3" json:"Identitie,omitempty"`
}

func (x *UpdateIdentitieRequest) Reset() {
	*x = UpdateIdentitieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identitie_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateIdentitieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIdentitieRequest) ProtoMessage() {}

func (x *UpdateIdentitieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identitie_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIdentitieRequest.ProtoReflect.Descriptor instead.
func (*UpdateIdentitieRequest) Descriptor() ([]byte, []int) {
	return file_proto_identitie_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateIdentitieRequest) GetIdentitie() *Identitie {
	if x != nil {
		return x.Identitie
	}
	return nil
}

type UpdateIdentitieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identitie *Identitie `protobuf:"bytes,1,opt,name=Identitie,proto3" json:"Identitie,omitempty"`
}

func (x *UpdateIdentitieResponse) Reset() {
	*x = UpdateIdentitieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identitie_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateIdentitieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIdentitieResponse) ProtoMessage() {}

func (x *UpdateIdentitieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identitie_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIdentitieResponse.ProtoReflect.Descriptor instead.
func (*UpdateIdentitieResponse) Descriptor() ([]byte, []int) {
	return file_proto_identitie_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateIdentitieResponse) GetIdentitie() *Identitie {
	if x != nil {
		return x.Identitie
	}
	return nil
}

type DeleteIdentitieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *DeleteIdentitieRequest) Reset() {
	*x = DeleteIdentitieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identitie_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteIdentitieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteIdentitieRequest) ProtoMessage() {}

func (x *DeleteIdentitieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identitie_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteIdentitieRequest.ProtoReflect.Descriptor instead.
func (*DeleteIdentitieRequest) Descriptor() ([]byte, []int) {
	return file_proto_identitie_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteIdentitieRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type DeleteIdentitieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteIdentitieResponse) Reset() {
	*x = DeleteIdentitieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identitie_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteIdentitieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteIdentitieResponse) ProtoMessage() {}

func (x *DeleteIdentitieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identitie_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteIdentitieResponse.ProtoReflect.Descriptor instead.
func (*DeleteIdentitieResponse) Descriptor() ([]byte, []int) {
	return file_proto_identitie_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteIdentitieResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type IdentitieCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IdentitieCredentials) Reset() {
	*x = IdentitieCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identitie_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentitieCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentitieCredentials) ProtoMessage() {}

func (x *IdentitieCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identitie_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentitieCredentials.ProtoReflect.Descriptor instead.
func (*IdentitieCredentials) Descriptor() ([]byte, []int) {
	return file_proto_identitie_proto_rawDescGZIP(), []int{0, 0}
}

var File_proto_identitie_proto protoreflect.FileDescriptor

var file_proto_identitie_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f,
	0x0a, 0x09, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x0d, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x22, 0x24, 0x0a, 0x0d, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x57,
	0x53, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x65, 0x64, 0x48, 0x61, 0x74, 0x10, 0x01, 0x22,
	0x48, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x52, 0x09,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x22, 0x49, 0x0a, 0x17, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x52, 0x09, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x22, 0x35, 0x0a, 0x14, 0x52, 0x65, 0x61, 0x64, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x15, 0x52,
	0x65, 0x61, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x52, 0x09, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x61, 0x64, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x48, 0x0a,
	0x16, 0x52, 0x65, 0x61, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x52, 0x09, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x22, 0x48, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2e, 0x0a, 0x09, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x52, 0x09, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x22, 0x49, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x52, 0x09, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x22, 0x37, 0x0a, 0x16,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xab, 0x03, 0x0a, 0x10, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x52, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x52, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x72, 0x75, 0x64, 0x72, 0x61, 0x69, 0x61, 0x2f,
	0x66, 0x6c, 0x65, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_identitie_proto_rawDescOnce sync.Once
	file_proto_identitie_proto_rawDescData = file_proto_identitie_proto_rawDesc
)

func file_proto_identitie_proto_rawDescGZIP() []byte {
	file_proto_identitie_proto_rawDescOnce.Do(func() {
		file_proto_identitie_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_identitie_proto_rawDescData)
	})
	return file_proto_identitie_proto_rawDescData
}

var file_proto_identitie_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_identitie_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_identitie_proto_goTypes = []interface{}{
	(IdentitieIdentityType)(0),      // 0: proto.Identitie.identity_type
	(*Identitie)(nil),               // 1: proto.Identitie
	(*CreateIdentitieRequest)(nil),  // 2: proto.CreateIdentitieRequest
	(*CreateIdentitieResponse)(nil), // 3: proto.CreateIdentitieResponse
	(*ReadIdentitieRequest)(nil),    // 4: proto.ReadIdentitieRequest
	(*ReadIdentitieResponse)(nil),   // 5: proto.ReadIdentitieResponse
	(*ReadIdentitiesRequest)(nil),   // 6: proto.ReadIdentitiesRequest
	(*ReadIdentitiesResponse)(nil),  // 7: proto.ReadIdentitiesResponse
	(*UpdateIdentitieRequest)(nil),  // 8: proto.UpdateIdentitieRequest
	(*UpdateIdentitieResponse)(nil), // 9: proto.UpdateIdentitieResponse
	(*DeleteIdentitieRequest)(nil),  // 10: proto.DeleteIdentitieRequest
	(*DeleteIdentitieResponse)(nil), // 11: proto.DeleteIdentitieResponse
	(*IdentitieCredentials)(nil),    // 12: proto.Identitie.credentials
}
var file_proto_identitie_proto_depIdxs = []int32{
	1,  // 0: proto.CreateIdentitieRequest.identitie:type_name -> proto.Identitie
	1,  // 1: proto.CreateIdentitieResponse.identitie:type_name -> proto.Identitie
	1,  // 2: proto.ReadIdentitieResponse.identitie:type_name -> proto.Identitie
	1,  // 3: proto.ReadIdentitiesResponse.identitie:type_name -> proto.Identitie
	1,  // 4: proto.UpdateIdentitieRequest.Identitie:type_name -> proto.Identitie
	1,  // 5: proto.UpdateIdentitieResponse.Identitie:type_name -> proto.Identitie
	2,  // 6: proto.IdentitieService.CreateIdentitie:input_type -> proto.CreateIdentitieRequest
	4,  // 7: proto.IdentitieService.GetIdentitie:input_type -> proto.ReadIdentitieRequest
	6,  // 8: proto.IdentitieService.GetIdentities:input_type -> proto.ReadIdentitiesRequest
	8,  // 9: proto.IdentitieService.UpdateIdentitie:input_type -> proto.UpdateIdentitieRequest
	10, // 10: proto.IdentitieService.DeleteIdentitie:input_type -> proto.DeleteIdentitieRequest
	3,  // 11: proto.IdentitieService.CreateIdentitie:output_type -> proto.CreateIdentitieResponse
	5,  // 12: proto.IdentitieService.GetIdentitie:output_type -> proto.ReadIdentitieResponse
	7,  // 13: proto.IdentitieService.GetIdentities:output_type -> proto.ReadIdentitiesResponse
	9,  // 14: proto.IdentitieService.UpdateIdentitie:output_type -> proto.UpdateIdentitieResponse
	11, // 15: proto.IdentitieService.DeleteIdentitie:output_type -> proto.DeleteIdentitieResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_identitie_proto_init() }
func file_proto_identitie_proto_init() {
	if File_proto_identitie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_identitie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identitie); i {
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
		file_proto_identitie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIdentitieRequest); i {
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
		file_proto_identitie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIdentitieResponse); i {
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
		file_proto_identitie_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadIdentitieRequest); i {
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
		file_proto_identitie_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadIdentitieResponse); i {
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
		file_proto_identitie_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadIdentitiesRequest); i {
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
		file_proto_identitie_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadIdentitiesResponse); i {
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
		file_proto_identitie_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateIdentitieRequest); i {
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
		file_proto_identitie_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateIdentitieResponse); i {
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
		file_proto_identitie_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteIdentitieRequest); i {
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
		file_proto_identitie_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteIdentitieResponse); i {
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
		file_proto_identitie_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentitieCredentials); i {
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
			RawDescriptor: file_proto_identitie_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_identitie_proto_goTypes,
		DependencyIndexes: file_proto_identitie_proto_depIdxs,
		EnumInfos:         file_proto_identitie_proto_enumTypes,
		MessageInfos:      file_proto_identitie_proto_msgTypes,
	}.Build()
	File_proto_identitie_proto = out.File
	file_proto_identitie_proto_rawDesc = nil
	file_proto_identitie_proto_goTypes = nil
	file_proto_identitie_proto_depIdxs = nil
}