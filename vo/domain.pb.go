// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: tech/storezhang/dnsabre/core/vo/domain.proto

package vo

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

// 域名
type Domain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 标识
	Id int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// 名称
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Domain) Reset() {
	*x = Domain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tech_storezhang_dnsabre_core_vo_domain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Domain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Domain) ProtoMessage() {}

func (x *Domain) ProtoReflect() protoreflect.Message {
	mi := &file_tech_storezhang_dnsabre_core_vo_domain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Domain.ProtoReflect.Descriptor instead.
func (*Domain) Descriptor() ([]byte, []int) {
	return file_tech_storezhang_dnsabre_core_vo_domain_proto_rawDescGZIP(), []int{0}
}

func (x *Domain) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Domain) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_tech_storezhang_dnsabre_core_vo_domain_proto protoreflect.FileDescriptor

var file_tech_storezhang_dnsabre_core_vo_domain_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e,
	0x67, 0x2f, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x6f, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e,
	0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x6f, 0x22,
	0x2c, 0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x42, 0x0a,
	0x1f, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67,
	0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x6f,
	0x50, 0x01, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x6f, 0x3b, 0x76,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tech_storezhang_dnsabre_core_vo_domain_proto_rawDescOnce sync.Once
	file_tech_storezhang_dnsabre_core_vo_domain_proto_rawDescData = file_tech_storezhang_dnsabre_core_vo_domain_proto_rawDesc
)

func file_tech_storezhang_dnsabre_core_vo_domain_proto_rawDescGZIP() []byte {
	file_tech_storezhang_dnsabre_core_vo_domain_proto_rawDescOnce.Do(func() {
		file_tech_storezhang_dnsabre_core_vo_domain_proto_rawDescData = protoimpl.X.CompressGZIP(file_tech_storezhang_dnsabre_core_vo_domain_proto_rawDescData)
	})
	return file_tech_storezhang_dnsabre_core_vo_domain_proto_rawDescData
}

var file_tech_storezhang_dnsabre_core_vo_domain_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tech_storezhang_dnsabre_core_vo_domain_proto_goTypes = []interface{}{
	(*Domain)(nil), // 0: tech.storezhang.dnsabre.core.vo.Domain
}
var file_tech_storezhang_dnsabre_core_vo_domain_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tech_storezhang_dnsabre_core_vo_domain_proto_init() }
func file_tech_storezhang_dnsabre_core_vo_domain_proto_init() {
	if File_tech_storezhang_dnsabre_core_vo_domain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tech_storezhang_dnsabre_core_vo_domain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Domain); i {
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
			RawDescriptor: file_tech_storezhang_dnsabre_core_vo_domain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tech_storezhang_dnsabre_core_vo_domain_proto_goTypes,
		DependencyIndexes: file_tech_storezhang_dnsabre_core_vo_domain_proto_depIdxs,
		MessageInfos:      file_tech_storezhang_dnsabre_core_vo_domain_proto_msgTypes,
	}.Build()
	File_tech_storezhang_dnsabre_core_vo_domain_proto = out.File
	file_tech_storezhang_dnsabre_core_vo_domain_proto_rawDesc = nil
	file_tech_storezhang_dnsabre_core_vo_domain_proto_goTypes = nil
	file_tech_storezhang_dnsabre_core_vo_domain_proto_depIdxs = nil
}
