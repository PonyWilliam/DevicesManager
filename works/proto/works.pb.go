// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.9.0
// source: proto/works.proto

//
//ID int64 `gorm:"primary_key;auto_increment;not_null" json:"id"`
//Name string `json:"name"`
//Nums string `json:"nums"`//工号
//Sex string `json:"sex"` //性别
//Level int64 `json:"level"`//等级
//Score int64 `json:"score"`//信誉分
//Place string `json:"place"`//住址
//Telephone string `json:"telephone"`//电话
//Mail string `json:"mail"`
//Description string `json:"description"`//补充描述
//ISWork bool `json:"is_work"`//是否在职

package works

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"` //未加密的password
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_works_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_works_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_proto_works_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code bool   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_works_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_works_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_proto_works_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetCode() bool {
	if x != nil {
		return x.Code
	}
	return false
}

func (x *LoginResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Request_Null struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Request_Null) Reset() {
	*x = Request_Null{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_works_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Null) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Null) ProtoMessage() {}

func (x *Request_Null) ProtoReflect() protoreflect.Message {
	mi := &file_proto_works_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Null.ProtoReflect.Descriptor instead.
func (*Request_Null) Descriptor() ([]byte, []int) {
	return file_proto_works_proto_rawDescGZIP(), []int{2}
}

type Request_Workers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Nums        string `protobuf:"bytes,2,opt,name=Nums,proto3" json:"Nums,omitempty"`
	Sex         string `protobuf:"bytes,3,opt,name=Sex,proto3" json:"Sex,omitempty"`
	Level       int64  `protobuf:"varint,4,opt,name=Level,proto3" json:"Level,omitempty"`
	Score       int64  `protobuf:"varint,5,opt,name=Score,proto3" json:"Score,omitempty"`
	Place       string `protobuf:"bytes,6,opt,name=Place,proto3" json:"Place,omitempty"`
	Telephone   string `protobuf:"bytes,7,opt,name=Telephone,proto3" json:"Telephone,omitempty"`
	Mail        string `protobuf:"bytes,8,opt,name=Mail,proto3" json:"Mail,omitempty"`
	Description string `protobuf:"bytes,9,opt,name=Description,proto3" json:"Description,omitempty"`
	ISWork      bool   `protobuf:"varint,10,opt,name=ISWork,proto3" json:"ISWork,omitempty"`
	Username    string `protobuf:"bytes,11,opt,name=username,proto3" json:"username,omitempty"`
	Password    string `protobuf:"bytes,12,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Request_Workers) Reset() {
	*x = Request_Workers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_works_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Workers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Workers) ProtoMessage() {}

func (x *Request_Workers) ProtoReflect() protoreflect.Message {
	mi := &file_proto_works_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Workers.ProtoReflect.Descriptor instead.
func (*Request_Workers) Descriptor() ([]byte, []int) {
	return file_proto_works_proto_rawDescGZIP(), []int{3}
}

func (x *Request_Workers) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Request_Workers) GetNums() string {
	if x != nil {
		return x.Nums
	}
	return ""
}

func (x *Request_Workers) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *Request_Workers) GetLevel() int64 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Request_Workers) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Request_Workers) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *Request_Workers) GetTelephone() string {
	if x != nil {
		return x.Telephone
	}
	return ""
}

func (x *Request_Workers) GetMail() string {
	if x != nil {
		return x.Mail
	}
	return ""
}

func (x *Request_Workers) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Request_Workers) GetISWork() bool {
	if x != nil {
		return x.ISWork
	}
	return false
}

func (x *Request_Workers) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Request_Workers) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Response_CreateWorker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response_CreateWorker) Reset() {
	*x = Response_CreateWorker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_works_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_CreateWorker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_CreateWorker) ProtoMessage() {}

func (x *Response_CreateWorker) ProtoReflect() protoreflect.Message {
	mi := &file_proto_works_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_CreateWorker.ProtoReflect.Descriptor instead.
func (*Response_CreateWorker) Descriptor() ([]byte, []int) {
	return file_proto_works_proto_rawDescGZIP(), []int{4}
}

func (x *Response_CreateWorker) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Response_CreateWorker) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Request_Workers_ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Request_Workers_ID) Reset() {
	*x = Request_Workers_ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_works_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Workers_ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Workers_ID) ProtoMessage() {}

func (x *Request_Workers_ID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_works_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Workers_ID.ProtoReflect.Descriptor instead.
func (*Request_Workers_ID) Descriptor() ([]byte, []int) {
	return file_proto_works_proto_rawDescGZIP(), []int{5}
}

func (x *Request_Workers_ID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Request_Workers_Nums struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nums int64 `protobuf:"varint,1,opt,name=nums,proto3" json:"nums,omitempty"`
}

func (x *Request_Workers_Nums) Reset() {
	*x = Request_Workers_Nums{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_works_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Workers_Nums) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Workers_Nums) ProtoMessage() {}

func (x *Request_Workers_Nums) ProtoReflect() protoreflect.Message {
	mi := &file_proto_works_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Workers_Nums.ProtoReflect.Descriptor instead.
func (*Request_Workers_Nums) Descriptor() ([]byte, []int) {
	return file_proto_works_proto_rawDescGZIP(), []int{6}
}

func (x *Request_Workers_Nums) GetNums() int64 {
	if x != nil {
		return x.Nums
	}
	return 0
}

type Request_Workers_Name struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Request_Workers_Name) Reset() {
	*x = Request_Workers_Name{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_works_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Workers_Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Workers_Name) ProtoMessage() {}

func (x *Request_Workers_Name) ProtoReflect() protoreflect.Message {
	mi := &file_proto_works_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Workers_Name.ProtoReflect.Descriptor instead.
func (*Request_Workers_Name) Descriptor() ([]byte, []int) {
	return file_proto_works_proto_rawDescGZIP(), []int{7}
}

func (x *Request_Workers_Name) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Response_Workers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response_Workers) Reset() {
	*x = Response_Workers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_works_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_Workers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_Workers) ProtoMessage() {}

func (x *Response_Workers) ProtoReflect() protoreflect.Message {
	mi := &file_proto_works_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_Workers.ProtoReflect.Descriptor instead.
func (*Response_Workers) Descriptor() ([]byte, []int) {
	return file_proto_works_proto_rawDescGZIP(), []int{8}
}

func (x *Response_Workers) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Response_Worker_Show struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Worker *Response_Workers_Info `protobuf:"bytes,1,opt,name=worker,proto3" json:"worker,omitempty"`
}

func (x *Response_Worker_Show) Reset() {
	*x = Response_Worker_Show{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_works_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_Worker_Show) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_Worker_Show) ProtoMessage() {}

func (x *Response_Worker_Show) ProtoReflect() protoreflect.Message {
	mi := &file_proto_works_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_Worker_Show.ProtoReflect.Descriptor instead.
func (*Response_Worker_Show) Descriptor() ([]byte, []int) {
	return file_proto_works_proto_rawDescGZIP(), []int{9}
}

func (x *Response_Worker_Show) GetWorker() *Response_Workers_Info {
	if x != nil {
		return x.Worker
	}
	return nil
}

type Response_Workers_Show struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workers []*Response_Workers_Info `protobuf:"bytes,1,rep,name=workers,proto3" json:"workers,omitempty"`
}

func (x *Response_Workers_Show) Reset() {
	*x = Response_Workers_Show{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_works_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_Workers_Show) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_Workers_Show) ProtoMessage() {}

func (x *Response_Workers_Show) ProtoReflect() protoreflect.Message {
	mi := &file_proto_works_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_Workers_Show.ProtoReflect.Descriptor instead.
func (*Response_Workers_Show) Descriptor() ([]byte, []int) {
	return file_proto_works_proto_rawDescGZIP(), []int{10}
}

func (x *Response_Workers_Show) GetWorkers() []*Response_Workers_Info {
	if x != nil {
		return x.Workers
	}
	return nil
}

type Response_Workers_Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Nums        string `protobuf:"bytes,3,opt,name=Nums,proto3" json:"Nums,omitempty"`
	Sex         string `protobuf:"bytes,4,opt,name=Sex,proto3" json:"Sex,omitempty"`
	Level       int64  `protobuf:"varint,5,opt,name=Level,proto3" json:"Level,omitempty"`
	Score       int64  `protobuf:"varint,6,opt,name=Score,proto3" json:"Score,omitempty"`
	Place       string `protobuf:"bytes,7,opt,name=Place,proto3" json:"Place,omitempty"`
	Telephone   string `protobuf:"bytes,8,opt,name=Telephone,proto3" json:"Telephone,omitempty"`
	Mail        string `protobuf:"bytes,9,opt,name=Mail,proto3" json:"Mail,omitempty"`
	Description string `protobuf:"bytes,10,opt,name=Description,proto3" json:"Description,omitempty"`
	ISWork      bool   `protobuf:"varint,11,opt,name=ISWork,proto3" json:"ISWork,omitempty"`
}

func (x *Response_Workers_Info) Reset() {
	*x = Response_Workers_Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_works_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_Workers_Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_Workers_Info) ProtoMessage() {}

func (x *Response_Workers_Info) ProtoReflect() protoreflect.Message {
	mi := &file_proto_works_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_Workers_Info.ProtoReflect.Descriptor instead.
func (*Response_Workers_Info) Descriptor() ([]byte, []int) {
	return file_proto_works_proto_rawDescGZIP(), []int{11}
}

func (x *Response_Workers_Info) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Response_Workers_Info) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Response_Workers_Info) GetNums() string {
	if x != nil {
		return x.Nums
	}
	return ""
}

func (x *Response_Workers_Info) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *Response_Workers_Info) GetLevel() int64 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Response_Workers_Info) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Response_Workers_Info) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *Response_Workers_Info) GetTelephone() string {
	if x != nil {
		return x.Telephone
	}
	return ""
}

func (x *Response_Workers_Info) GetMail() string {
	if x != nil {
		return x.Mail
	}
	return ""
}

func (x *Response_Workers_Info) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Response_Workers_Info) GetISWork() bool {
	if x != nil {
		return x.ISWork
	}
	return false
}

var File_proto_works_proto protoreflect.FileDescriptor

var file_proto_works_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x22, 0x3e, 0x0a, 0x0c, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x35, 0x0a, 0x0d, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x0e, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x4e, 0x75, 0x6c,
	0x6c, 0x22, 0xb1, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x75, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x75, 0x6d, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x53, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x65, 0x78, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4d,
	0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x53, 0x57, 0x6f, 0x72, 0x6b, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x49, 0x53, 0x57, 0x6f, 0x72, 0x6b, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x41, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x5f, 0x49, 0x44, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2a,
	0x0a, 0x14, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x73, 0x5f, 0x4e, 0x75, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x2a, 0x0a, 0x14, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x5f, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x5f, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x4c, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x34, 0x0a, 0x06,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x73, 0x5f, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x22, 0x4f, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x5f, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x36, 0x0a, 0x07, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x73, 0x5f, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x73, 0x22, 0x8f, 0x02, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x5f, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x75, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x75, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x53, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x65, 0x6c,
	0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x65,
	0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x49, 0x53, 0x57, 0x6f, 0x72, 0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x49,
	0x53, 0x57, 0x6f, 0x72, 0x6b, 0x32, 0xb7, 0x04, 0x0a, 0x05, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x12,
	0x44, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12,
	0x16, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x1a, 0x1c, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x1a, 0x1c, 0x2e,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12,
	0x19, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x5f, 0x49, 0x44, 0x1a, 0x17, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x73, 0x12, 0x48, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x19, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x5f, 0x49, 0x44,
	0x1a, 0x1b, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x4c, 0x0a,
	0x10, 0x46, 0x69, 0x6e, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x42, 0x79, 0x4e, 0x75, 0x6d,
	0x73, 0x12, 0x1b, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x5f, 0x4e, 0x75, 0x6d, 0x73, 0x1a, 0x1b,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x4d, 0x0a, 0x10, 0x46,
	0x69, 0x6e, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x5f, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x1c, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x73, 0x5f, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x3c, 0x0a, 0x07, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x13, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x4e, 0x75, 0x6c, 0x6c, 0x1a, 0x1c, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x73, 0x5f, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x35, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x53, 0x75, 0x6d, 0x12, 0x13, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0d, 0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_works_proto_rawDescOnce sync.Once
	file_proto_works_proto_rawDescData = file_proto_works_proto_rawDesc
)

func file_proto_works_proto_rawDescGZIP() []byte {
	file_proto_works_proto_rawDescOnce.Do(func() {
		file_proto_works_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_works_proto_rawDescData)
	})
	return file_proto_works_proto_rawDescData
}

var file_proto_works_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_works_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),          // 0: works.LoginRequest
	(*LoginResponse)(nil),         // 1: works.LoginResponse
	(*Request_Null)(nil),          // 2: works.Request_Null
	(*Request_Workers)(nil),       // 3: works.Request_Workers
	(*Response_CreateWorker)(nil), // 4: works.Response_CreateWorker
	(*Request_Workers_ID)(nil),    // 5: works.Request_Workers_ID
	(*Request_Workers_Nums)(nil),  // 6: works.Request_Workers_Nums
	(*Request_Workers_Name)(nil),  // 7: works.Request_Workers_Name
	(*Response_Workers)(nil),      // 8: works.Response_Workers
	(*Response_Worker_Show)(nil),  // 9: works.Response_Worker_Show
	(*Response_Workers_Show)(nil), // 10: works.Response_Workers_Show
	(*Response_Workers_Info)(nil), // 11: works.Response_Workers_Info
}
var file_proto_works_proto_depIdxs = []int32{
	11, // 0: works.Response_Worker_Show.worker:type_name -> works.Response_Workers_Info
	11, // 1: works.Response_Workers_Show.workers:type_name -> works.Response_Workers_Info
	3,  // 2: works.Works.CreateWorker:input_type -> works.Request_Workers
	3,  // 3: works.Works.UpdateWorker:input_type -> works.Request_Workers
	5,  // 4: works.Works.DeleteWorkerByID:input_type -> works.Request_Workers_ID
	5,  // 5: works.Works.FindWorkerByID:input_type -> works.Request_Workers_ID
	6,  // 6: works.Works.FindWorkerByNums:input_type -> works.Request_Workers_Nums
	7,  // 7: works.Works.FindWorkerByName:input_type -> works.Request_Workers_Name
	2,  // 8: works.Works.FindAll:input_type -> works.Request_Null
	0,  // 9: works.Works.CheckSum:input_type -> works.LoginRequest
	4,  // 10: works.Works.CreateWorker:output_type -> works.Response_CreateWorker
	4,  // 11: works.Works.UpdateWorker:output_type -> works.Response_CreateWorker
	8,  // 12: works.Works.DeleteWorkerByID:output_type -> works.Response_Workers
	9,  // 13: works.Works.FindWorkerByID:output_type -> works.Response_Worker_Show
	9,  // 14: works.Works.FindWorkerByNums:output_type -> works.Response_Worker_Show
	10, // 15: works.Works.FindWorkerByName:output_type -> works.Response_Workers_Show
	10, // 16: works.Works.FindAll:output_type -> works.Response_Workers_Show
	1,  // 17: works.Works.CheckSum:output_type -> works.LoginResponse
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_works_proto_init() }
func file_proto_works_proto_init() {
	if File_proto_works_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_works_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_proto_works_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_proto_works_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Null); i {
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
		file_proto_works_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Workers); i {
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
		file_proto_works_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_CreateWorker); i {
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
		file_proto_works_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Workers_ID); i {
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
		file_proto_works_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Workers_Nums); i {
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
		file_proto_works_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Workers_Name); i {
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
		file_proto_works_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_Workers); i {
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
		file_proto_works_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_Worker_Show); i {
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
		file_proto_works_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_Workers_Show); i {
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
		file_proto_works_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_Workers_Info); i {
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
			RawDescriptor: file_proto_works_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_works_proto_goTypes,
		DependencyIndexes: file_proto_works_proto_depIdxs,
		MessageInfos:      file_proto_works_proto_msgTypes,
	}.Build()
	File_proto_works_proto = out.File
	file_proto_works_proto_rawDesc = nil
	file_proto_works_proto_goTypes = nil
	file_proto_works_proto_depIdxs = nil
}
