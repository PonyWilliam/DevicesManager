// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.9.0
// source: proto/borrow.proto

package borrow

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

type Borrow_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerId     int64  `protobuf:"varint,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`             //借出人信息
	ProductId    int64  `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`          //借出产品信息
	BorrowTime   int64  `protobuf:"varint,3,opt,name=borrow_time,json=borrowTime,proto3" json:"borrow_time,omitempty"`       //借出时间
	ScheduleTime int64  `protobuf:"varint,4,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"` //预计归还时间
	ReturnTime   int64  `protobuf:"varint,5,opt,name=return_time,json=returnTime,proto3" json:"return_time,omitempty"`
	Description  string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"` //备注信息,转借等等说明
}

func (x *Borrow_Request) Reset() {
	*x = Borrow_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_borrow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Borrow_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Borrow_Request) ProtoMessage() {}

func (x *Borrow_Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_borrow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Borrow_Request.ProtoReflect.Descriptor instead.
func (*Borrow_Request) Descriptor() ([]byte, []int) {
	return file_proto_borrow_proto_rawDescGZIP(), []int{0}
}

func (x *Borrow_Request) GetWorkerId() int64 {
	if x != nil {
		return x.WorkerId
	}
	return 0
}

func (x *Borrow_Request) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *Borrow_Request) GetBorrowTime() int64 {
	if x != nil {
		return x.BorrowTime
	}
	return 0
}

func (x *Borrow_Request) GetScheduleTime() int64 {
	if x != nil {
		return x.ScheduleTime
	}
	return 0
}

func (x *Borrow_Request) GetReturnTime() int64 {
	if x != nil {
		return x.ReturnTime
	}
	return 0
}

func (x *Borrow_Request) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Returns_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Returns_Request) Reset() {
	*x = Returns_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_borrow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Returns_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Returns_Request) ProtoMessage() {}

func (x *Returns_Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_borrow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Returns_Request.ProtoReflect.Descriptor instead.
func (*Returns_Request) Descriptor() ([]byte, []int) {
	return file_proto_borrow_proto_rawDescGZIP(), []int{1}
}

func (x *Returns_Request) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ToOtherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`   //物品已有记录
	Wid int64 `protobuf:"varint,2,opt,name=wid,proto3" json:"wid,omitempty"` //其它人的id，方便变更
}

func (x *ToOtherRequest) Reset() {
	*x = ToOtherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_borrow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToOtherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToOtherRequest) ProtoMessage() {}

func (x *ToOtherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_borrow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToOtherRequest.ProtoReflect.Descriptor instead.
func (*ToOtherRequest) Descriptor() ([]byte, []int) {
	return file_proto_borrow_proto_rawDescGZIP(), []int{2}
}

func (x *ToOtherRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ToOtherRequest) GetWid() int64 {
	if x != nil {
		return x.Wid
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_borrow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_borrow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_borrow_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_borrow_proto protoreflect.FileDescriptor

var file_proto_borrow_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x22, 0xd5, 0x01, 0x0a,
	0x0e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62,
	0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x21, 0x0a, 0x0f, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x5f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x10, 0x74, 0x6f, 0x5f, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x77,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x77, 0x69, 0x64, 0x22, 0x3c, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xa8, 0x01, 0x0a, 0x06,
	0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x12, 0x32, 0x0a, 0x06, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x12, 0x16, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f,
	0x77, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x73, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x35, 0x0a, 0x07, 0x54, 0x6f, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x62, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x2e, 0x74, 0x6f, 0x5f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_borrow_proto_rawDescOnce sync.Once
	file_proto_borrow_proto_rawDescData = file_proto_borrow_proto_rawDesc
)

func file_proto_borrow_proto_rawDescGZIP() []byte {
	file_proto_borrow_proto_rawDescOnce.Do(func() {
		file_proto_borrow_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_borrow_proto_rawDescData)
	})
	return file_proto_borrow_proto_rawDescData
}

var file_proto_borrow_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_borrow_proto_goTypes = []interface{}{
	(*Borrow_Request)(nil),  // 0: borrow.Borrow_Request
	(*Returns_Request)(nil), // 1: borrow.Returns_Request
	(*ToOtherRequest)(nil),  // 2: borrow.to_other_request
	(*Response)(nil),        // 3: borrow.Response
}
var file_proto_borrow_proto_depIdxs = []int32{
	0, // 0: borrow.Borrow.Borrow:input_type -> borrow.Borrow_Request
	1, // 1: borrow.Borrow.Return:input_type -> borrow.Returns_Request
	2, // 2: borrow.Borrow.ToOther:input_type -> borrow.to_other_request
	3, // 3: borrow.Borrow.Borrow:output_type -> borrow.Response
	3, // 4: borrow.Borrow.Return:output_type -> borrow.Response
	3, // 5: borrow.Borrow.ToOther:output_type -> borrow.Response
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_borrow_proto_init() }
func file_proto_borrow_proto_init() {
	if File_proto_borrow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_borrow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Borrow_Request); i {
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
		file_proto_borrow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Returns_Request); i {
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
		file_proto_borrow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToOtherRequest); i {
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
		file_proto_borrow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_proto_borrow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_borrow_proto_goTypes,
		DependencyIndexes: file_proto_borrow_proto_depIdxs,
		MessageInfos:      file_proto_borrow_proto_msgTypes,
	}.Build()
	File_proto_borrow_proto = out.File
	file_proto_borrow_proto_rawDesc = nil
	file_proto_borrow_proto_goTypes = nil
	file_proto_borrow_proto_depIdxs = nil
}
