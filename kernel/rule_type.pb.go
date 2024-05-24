// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: tech/storezhang/dnsabre/core/kernel/rule_type.proto

package kernel

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

// 规则类型
type RuleType int32

const (
	// 未知，不要使用
	RuleType_RULE_TYPE_UNSPECIFIED RuleType = 0
	RuleType_RULE_TYPE_HOSTS       RuleType = 1
)

// Enum value maps for RuleType.
var (
	RuleType_name = map[int32]string{
		0: "RULE_TYPE_UNSPECIFIED",
		1: "RULE_TYPE_HOSTS",
	}
	RuleType_value = map[string]int32{
		"RULE_TYPE_UNSPECIFIED": 0,
		"RULE_TYPE_HOSTS":       1,
	}
)

func (x RuleType) Enum() *RuleType {
	p := new(RuleType)
	*p = x
	return p
}

func (x RuleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RuleType) Descriptor() protoreflect.EnumDescriptor {
	return file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_enumTypes[0].Descriptor()
}

func (RuleType) Type() protoreflect.EnumType {
	return &file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_enumTypes[0]
}

func (x RuleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RuleType.Descriptor instead.
func (RuleType) EnumDescriptor() ([]byte, []int) {
	return file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_rawDescGZIP(), []int{0}
}

var File_tech_storezhang_dnsabre_core_kernel_rule_type_proto protoreflect.FileDescriptor

var file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_rawDesc = []byte{
	0x0a, 0x33, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e,
	0x67, 0x2f, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6b,
	0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2a, 0x3a, 0x0a, 0x08, 0x52, 0x75,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x55, 0x4c, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x55, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48,
	0x4f, 0x53, 0x54, 0x53, 0x10, 0x01, 0x42, 0x4e, 0x0a, 0x23, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72,
	0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x50, 0x01, 0x5a,
	0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6e, 0x73, 0x61,
	0x62, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x3b,
	0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_rawDescOnce sync.Once
	file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_rawDescData = file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_rawDesc
)

func file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_rawDescGZIP() []byte {
	file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_rawDescOnce.Do(func() {
		file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_rawDescData)
	})
	return file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_rawDescData
}

var file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_goTypes = []interface{}{
	(RuleType)(0), // 0: tech.storezhang.dnsabre.core.kernel.RuleType
}
var file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_init() }
func file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_init() {
	if File_tech_storezhang_dnsabre_core_kernel_rule_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_goTypes,
		DependencyIndexes: file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_depIdxs,
		EnumInfos:         file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_enumTypes,
	}.Build()
	File_tech_storezhang_dnsabre_core_kernel_rule_type_proto = out.File
	file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_rawDesc = nil
	file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_goTypes = nil
	file_tech_storezhang_dnsabre_core_kernel_rule_type_proto_depIdxs = nil
}
