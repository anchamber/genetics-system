// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: system.proto

package proto

import (
	proto "github.com/anchamber/genetics-api/proto"
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

type SystemType int32

const (
	SystemType_UNKNOWN     SystemType = 0
	SystemType_GLASS       SystemType = 1
	SystemType_TECHNIPLAST SystemType = 2
)

// Enum value maps for SystemType.
var (
	SystemType_name = map[int32]string{
		0: "UNKNOWN",
		1: "GLASS",
		2: "TECHNIPLAST",
	}
	SystemType_value = map[string]int32{
		"UNKNOWN":     0,
		"GLASS":       1,
		"TECHNIPLAST": 2,
	}
)

func (x SystemType) Enum() *SystemType {
	p := new(SystemType)
	*p = x
	return p
}

func (x SystemType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SystemType) Descriptor() protoreflect.EnumDescriptor {
	return file_system_proto_enumTypes[0].Descriptor()
}

func (SystemType) Type() protoreflect.EnumType {
	return &file_system_proto_enumTypes[0]
}

func (x SystemType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SystemType.Descriptor instead.
func (SystemType) EnumDescriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{0}
}

type System struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *System) Reset() {
	*x = System{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *System) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*System) ProtoMessage() {}

func (x *System) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use System.ProtoReflect.Descriptor instead.
func (*System) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{0}
}

type Systems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Systems []*System `protobuf:"bytes,1,rep,name=systems,proto3" json:"systems,omitempty"`
}

func (x *Systems) Reset() {
	*x = Systems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Systems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Systems) ProtoMessage() {}

func (x *Systems) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Systems.ProtoReflect.Descriptor instead.
func (*Systems) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{1}
}

func (x *Systems) GetSystems() []*System {
	if x != nil {
		return x.Systems
	}
	return nil
}

type GetSystemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filters     []*proto.Filter   `protobuf:"bytes,1,rep,name=filters,proto3" json:"filters,omitempty"`
	Pageination *proto.Pagination `protobuf:"bytes,2,opt,name=pageination,proto3" json:"pageination,omitempty"`
}

func (x *GetSystemsRequest) Reset() {
	*x = GetSystemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSystemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSystemsRequest) ProtoMessage() {}

func (x *GetSystemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSystemsRequest.ProtoReflect.Descriptor instead.
func (*GetSystemsRequest) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{2}
}

func (x *GetSystemsRequest) GetFilters() []*proto.Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *GetSystemsRequest) GetPageination() *proto.Pagination {
	if x != nil {
		return x.Pageination
	}
	return nil
}

type GetSystemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetSystemRequest) Reset() {
	*x = GetSystemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSystemRequest) ProtoMessage() {}

func (x *GetSystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSystemRequest.ProtoReflect.Descriptor instead.
func (*GetSystemRequest) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{3}
}

func (x *GetSystemRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SystemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location         string     `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Type             SystemType `protobuf:"varint,4,opt,name=type,proto3,enum=genetics_system.SystemType" json:"type,omitempty"`
	Responsible      string     `protobuf:"bytes,5,opt,name=responsible,proto3" json:"responsible,omitempty"`
	CleaningInterval int32      `protobuf:"varint,6,opt,name=cleaning_interval,json=cleaningInterval,proto3" json:"cleaning_interval,omitempty"`
	LastCleaned      int64      `protobuf:"varint,7,opt,name=last_cleaned,json=lastCleaned,proto3" json:"last_cleaned,omitempty"`
}

func (x *SystemResponse) Reset() {
	*x = SystemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemResponse) ProtoMessage() {}

func (x *SystemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemResponse.ProtoReflect.Descriptor instead.
func (*SystemResponse) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{4}
}

func (x *SystemResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SystemResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SystemResponse) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *SystemResponse) GetType() SystemType {
	if x != nil {
		return x.Type
	}
	return SystemType_UNKNOWN
}

func (x *SystemResponse) GetResponsible() string {
	if x != nil {
		return x.Responsible
	}
	return ""
}

func (x *SystemResponse) GetCleaningInterval() int32 {
	if x != nil {
		return x.CleaningInterval
	}
	return 0
}

func (x *SystemResponse) GetLastCleaned() int64 {
	if x != nil {
		return x.LastCleaned
	}
	return 0
}

type CreateSystemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location         string     `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Type             SystemType `protobuf:"varint,3,opt,name=type,proto3,enum=genetics_system.SystemType" json:"type,omitempty"`
	Responsible      string     `protobuf:"bytes,4,opt,name=responsible,proto3" json:"responsible,omitempty"`
	CleaningInterval int32      `protobuf:"varint,5,opt,name=cleaning_interval,json=cleaningInterval,proto3" json:"cleaning_interval,omitempty"`
	LastCleaned      int64      `protobuf:"varint,6,opt,name=last_cleaned,json=lastCleaned,proto3" json:"last_cleaned,omitempty"`
}

func (x *CreateSystemRequest) Reset() {
	*x = CreateSystemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSystemRequest) ProtoMessage() {}

func (x *CreateSystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSystemRequest.ProtoReflect.Descriptor instead.
func (*CreateSystemRequest) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{5}
}

func (x *CreateSystemRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSystemRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *CreateSystemRequest) GetType() SystemType {
	if x != nil {
		return x.Type
	}
	return SystemType_UNKNOWN
}

func (x *CreateSystemRequest) GetResponsible() string {
	if x != nil {
		return x.Responsible
	}
	return ""
}

func (x *CreateSystemRequest) GetCleaningInterval() int32 {
	if x != nil {
		return x.CleaningInterval
	}
	return 0
}

func (x *CreateSystemRequest) GetLastCleaned() int64 {
	if x != nil {
		return x.LastCleaned
	}
	return 0
}

type CreateSystemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateSystemResponse) Reset() {
	*x = CreateSystemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSystemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSystemResponse) ProtoMessage() {}

func (x *CreateSystemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSystemResponse.ProtoReflect.Descriptor instead.
func (*CreateSystemResponse) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{6}
}

type UpdateSystemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location         string     `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Type             SystemType `protobuf:"varint,3,opt,name=type,proto3,enum=genetics_system.SystemType" json:"type,omitempty"`
	Responsible      string     `protobuf:"bytes,4,opt,name=responsible,proto3" json:"responsible,omitempty"`
	CleaningInterval int32      `protobuf:"varint,5,opt,name=cleaning_interval,json=cleaningInterval,proto3" json:"cleaning_interval,omitempty"`
	LastCleaned      int64      `protobuf:"varint,6,opt,name=last_cleaned,json=lastCleaned,proto3" json:"last_cleaned,omitempty"`
}

func (x *UpdateSystemRequest) Reset() {
	*x = UpdateSystemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSystemRequest) ProtoMessage() {}

func (x *UpdateSystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSystemRequest.ProtoReflect.Descriptor instead.
func (*UpdateSystemRequest) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateSystemRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateSystemRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *UpdateSystemRequest) GetType() SystemType {
	if x != nil {
		return x.Type
	}
	return SystemType_UNKNOWN
}

func (x *UpdateSystemRequest) GetResponsible() string {
	if x != nil {
		return x.Responsible
	}
	return ""
}

func (x *UpdateSystemRequest) GetCleaningInterval() int32 {
	if x != nil {
		return x.CleaningInterval
	}
	return 0
}

func (x *UpdateSystemRequest) GetLastCleaned() int64 {
	if x != nil {
		return x.LastCleaned
	}
	return 0
}

type UpdateSystemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateSystemResponse) Reset() {
	*x = UpdateSystemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSystemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSystemResponse) ProtoMessage() {}

func (x *UpdateSystemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSystemResponse.ProtoReflect.Descriptor instead.
func (*UpdateSystemResponse) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{8}
}

type DeleteSystemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteSystemRequest) Reset() {
	*x = DeleteSystemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSystemRequest) ProtoMessage() {}

func (x *DeleteSystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSystemRequest.ProtoReflect.Descriptor instead.
func (*DeleteSystemRequest) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteSystemRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteSystemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSystemResponse) Reset() {
	*x = DeleteSystemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSystemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSystemResponse) ProtoMessage() {}

func (x *DeleteSystemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSystemResponse.ProtoReflect.Descriptor instead.
func (*DeleteSystemResponse) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{10}
}

var File_system_proto protoreflect.FileDescriptor

var file_system_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x67, 0x65, 0x6e, 0x65, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a,
	0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x08, 0x0a, 0x06, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x22, 0x3c, 0x0a, 0x07, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x31, 0x0a, 0x07, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x6d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x31,
	0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x26, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xf3, 0x01, 0x0a, 0x0e, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12,
	0x2b, 0x0a, 0x11, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x63, 0x6c, 0x65, 0x61,
	0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x65, 0x64, 0x22,
	0xe8, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x74, 0x69, 0x63, 0x73,
	0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6c,
	0x65, 0x61, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x69, 0x6e, 0x67, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x63, 0x6c, 0x65, 0x61, 0x6e, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c,
	0x61, 0x73, 0x74, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x65, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0xe8, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x74,
	0x69, 0x63, 0x73, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x2b, 0x0a,
	0x11, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x65, 0x64, 0x22, 0x16, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x35, 0x0a, 0x0a, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x4c, 0x41, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0f,
	0x0a, 0x0b, 0x54, 0x45, 0x43, 0x48, 0x4e, 0x49, 0x50, 0x4c, 0x41, 0x53, 0x54, 0x10, 0x02, 0x32,
	0xd6, 0x03, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x55, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x22, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x51, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x21, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x74, 0x69, 0x63, 0x73,
	0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x74,
	0x69, 0x63, 0x73, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x24, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x24, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0c, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x24, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x63, 0x68, 0x61, 0x6d, 0x62, 0x65, 0x72,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x74, 0x69, 0x63, 0x73, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_system_proto_rawDescOnce sync.Once
	file_system_proto_rawDescData = file_system_proto_rawDesc
)

func file_system_proto_rawDescGZIP() []byte {
	file_system_proto_rawDescOnce.Do(func() {
		file_system_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_proto_rawDescData)
	})
	return file_system_proto_rawDescData
}

var file_system_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_system_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_system_proto_goTypes = []interface{}{
	(SystemType)(0),              // 0: genetics_system.SystemType
	(*System)(nil),               // 1: genetics_system.System
	(*Systems)(nil),              // 2: genetics_system.Systems
	(*GetSystemsRequest)(nil),    // 3: genetics_system.GetSystemsRequest
	(*GetSystemRequest)(nil),     // 4: genetics_system.GetSystemRequest
	(*SystemResponse)(nil),       // 5: genetics_system.SystemResponse
	(*CreateSystemRequest)(nil),  // 6: genetics_system.CreateSystemRequest
	(*CreateSystemResponse)(nil), // 7: genetics_system.CreateSystemResponse
	(*UpdateSystemRequest)(nil),  // 8: genetics_system.UpdateSystemRequest
	(*UpdateSystemResponse)(nil), // 9: genetics_system.UpdateSystemResponse
	(*DeleteSystemRequest)(nil),  // 10: genetics_system.DeleteSystemRequest
	(*DeleteSystemResponse)(nil), // 11: genetics_system.DeleteSystemResponse
	(*proto.Filter)(nil),         // 12: api.Filter
	(*proto.Pagination)(nil),     // 13: api.Pagination
}
var file_system_proto_depIdxs = []int32{
	1,  // 0: genetics_system.Systems.systems:type_name -> genetics_system.System
	12, // 1: genetics_system.GetSystemsRequest.filters:type_name -> api.Filter
	13, // 2: genetics_system.GetSystemsRequest.pageination:type_name -> api.Pagination
	0,  // 3: genetics_system.SystemResponse.type:type_name -> genetics_system.SystemType
	0,  // 4: genetics_system.CreateSystemRequest.type:type_name -> genetics_system.SystemType
	0,  // 5: genetics_system.UpdateSystemRequest.type:type_name -> genetics_system.SystemType
	3,  // 6: genetics_system.SystemService.GetSystems:input_type -> genetics_system.GetSystemsRequest
	4,  // 7: genetics_system.SystemService.GetSystem:input_type -> genetics_system.GetSystemRequest
	6,  // 8: genetics_system.SystemService.CreateSystem:input_type -> genetics_system.CreateSystemRequest
	8,  // 9: genetics_system.SystemService.UpdateSystem:input_type -> genetics_system.UpdateSystemRequest
	10, // 10: genetics_system.SystemService.DeleteSystem:input_type -> genetics_system.DeleteSystemRequest
	5,  // 11: genetics_system.SystemService.GetSystems:output_type -> genetics_system.SystemResponse
	5,  // 12: genetics_system.SystemService.GetSystem:output_type -> genetics_system.SystemResponse
	7,  // 13: genetics_system.SystemService.CreateSystem:output_type -> genetics_system.CreateSystemResponse
	9,  // 14: genetics_system.SystemService.UpdateSystem:output_type -> genetics_system.UpdateSystemResponse
	11, // 15: genetics_system.SystemService.DeleteSystem:output_type -> genetics_system.DeleteSystemResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_system_proto_init() }
func file_system_proto_init() {
	if File_system_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_system_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*System); i {
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
		file_system_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Systems); i {
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
		file_system_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSystemsRequest); i {
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
		file_system_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSystemRequest); i {
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
		file_system_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemResponse); i {
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
		file_system_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSystemRequest); i {
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
		file_system_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSystemResponse); i {
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
		file_system_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSystemRequest); i {
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
		file_system_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSystemResponse); i {
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
		file_system_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSystemRequest); i {
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
		file_system_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSystemResponse); i {
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
			RawDescriptor: file_system_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_system_proto_goTypes,
		DependencyIndexes: file_system_proto_depIdxs,
		EnumInfos:         file_system_proto_enumTypes,
		MessageInfos:      file_system_proto_msgTypes,
	}.Build()
	File_system_proto = out.File
	file_system_proto_rawDesc = nil
	file_system_proto_goTypes = nil
	file_system_proto_depIdxs = nil
}
