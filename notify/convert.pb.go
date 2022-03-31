// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: ziyunix/message/notify/convert.proto

package notify

import (
	core "github.com/ziyunix/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 转码通知
type Convert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编号
	Id int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// 状态
	Status core.Status `protobuf:"varint,4,opt,name=status,proto3,enum=ziyunix.core.Status" json:"status,omitempty"`
	// 发生时间
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Convert) Reset() {
	*x = Convert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ziyunix_message_notify_convert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Convert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Convert) ProtoMessage() {}

func (x *Convert) ProtoReflect() protoreflect.Message {
	mi := &file_ziyunix_message_notify_convert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Convert.ProtoReflect.Descriptor instead.
func (*Convert) Descriptor() ([]byte, []int) {
	return file_ziyunix_message_notify_convert_proto_rawDescGZIP(), []int{0}
}

func (x *Convert) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Convert) GetStatus() core.Status {
	if x != nil {
		return x.Status
	}
	return core.Status_STATUS_UNSPECIFIED
}

func (x *Convert) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_ziyunix_message_notify_convert_proto protoreflect.FileDescriptor

var file_ziyunix_message_notify_convert_proto_rawDesc = []byte{
	0x0a, 0x24, 0x7a, 0x69, 0x79, 0x75, 0x6e, 0x69, 0x78, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x7a, 0x69, 0x79, 0x75, 0x6e, 0x69, 0x78, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x7a, 0x69, 0x79, 0x75, 0x6e, 0x69, 0x78, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x7a, 0x69, 0x79, 0x75, 0x6e, 0x69, 0x78,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x48,
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x7a, 0x69, 0x79, 0x75, 0x6e, 0x69, 0x78, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x50, 0x01, 0x5a, 0x28,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x79, 0x75, 0x6e,
	0x69, 0x78, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x3b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ziyunix_message_notify_convert_proto_rawDescOnce sync.Once
	file_ziyunix_message_notify_convert_proto_rawDescData = file_ziyunix_message_notify_convert_proto_rawDesc
)

func file_ziyunix_message_notify_convert_proto_rawDescGZIP() []byte {
	file_ziyunix_message_notify_convert_proto_rawDescOnce.Do(func() {
		file_ziyunix_message_notify_convert_proto_rawDescData = protoimpl.X.CompressGZIP(file_ziyunix_message_notify_convert_proto_rawDescData)
	})
	return file_ziyunix_message_notify_convert_proto_rawDescData
}

var file_ziyunix_message_notify_convert_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ziyunix_message_notify_convert_proto_goTypes = []interface{}{
	(*Convert)(nil),               // 0: ziyunix.message.notify.Convert
	(core.Status)(0),              // 1: ziyunix.core.Status
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_ziyunix_message_notify_convert_proto_depIdxs = []int32{
	1, // 0: ziyunix.message.notify.Convert.status:type_name -> ziyunix.core.Status
	2, // 1: ziyunix.message.notify.Convert.timestamp:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ziyunix_message_notify_convert_proto_init() }
func file_ziyunix_message_notify_convert_proto_init() {
	if File_ziyunix_message_notify_convert_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ziyunix_message_notify_convert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Convert); i {
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
			RawDescriptor: file_ziyunix_message_notify_convert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ziyunix_message_notify_convert_proto_goTypes,
		DependencyIndexes: file_ziyunix_message_notify_convert_proto_depIdxs,
		MessageInfos:      file_ziyunix_message_notify_convert_proto_msgTypes,
	}.Build()
	File_ziyunix_message_notify_convert_proto = out.File
	file_ziyunix_message_notify_convert_proto_rawDesc = nil
	file_ziyunix_message_notify_convert_proto_goTypes = nil
	file_ziyunix_message_notify_convert_proto_depIdxs = nil
}
