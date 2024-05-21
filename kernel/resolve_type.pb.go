// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: tech/storezhang/dnsabre/core/kernel/resolve_type.proto

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

// 解析类型
type ResolveType int32

const (
	// 未知，不要使用
	ResolveType_RESOLVE_TYPE_UNSPECIFIED ResolveType = 0
	ResolveType_RESOLVE_TYPE_A           ResolveType = 1
	ResolveType_RESOLVE_TYPE_AAAA        ResolveType = 2
	ResolveType_RESOLVE_TYPE_CNAME       ResolveType = 3
	ResolveType_RESOLVE_TYPE_PTR         ResolveType = 4
	ResolveType_RESOLVE_TYPE_MX          ResolveType = 5
	ResolveType_RESOLVE_TYPE_TXT         ResolveType = 6
	ResolveType_RESOLVE_TYPE_SRV         ResolveType = 7
	ResolveType_RESOLVE_TYPE_NS          ResolveType = 8
	ResolveType_RESOLVE_TYPE_SOA         ResolveType = 9
)

// Enum value maps for ResolveType.
var (
	ResolveType_name = map[int32]string{
		0: "RESOLVE_TYPE_UNSPECIFIED",
		1: "RESOLVE_TYPE_A",
		2: "RESOLVE_TYPE_AAAA",
		3: "RESOLVE_TYPE_CNAME",
		4: "RESOLVE_TYPE_PTR",
		5: "RESOLVE_TYPE_MX",
		6: "RESOLVE_TYPE_TXT",
		7: "RESOLVE_TYPE_SRV",
		8: "RESOLVE_TYPE_NS",
		9: "RESOLVE_TYPE_SOA",
	}
	ResolveType_value = map[string]int32{
		"RESOLVE_TYPE_UNSPECIFIED": 0,
		"RESOLVE_TYPE_A":           1,
		"RESOLVE_TYPE_AAAA":        2,
		"RESOLVE_TYPE_CNAME":       3,
		"RESOLVE_TYPE_PTR":         4,
		"RESOLVE_TYPE_MX":          5,
		"RESOLVE_TYPE_TXT":         6,
		"RESOLVE_TYPE_SRV":         7,
		"RESOLVE_TYPE_NS":          8,
		"RESOLVE_TYPE_SOA":         9,
	}
)

func (x ResolveType) Enum() *ResolveType {
	p := new(ResolveType)
	*p = x
	return p
}

func (x ResolveType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResolveType) Descriptor() protoreflect.EnumDescriptor {
	return file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_enumTypes[0].Descriptor()
}

func (ResolveType) Type() protoreflect.EnumType {
	return &file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_enumTypes[0]
}

func (x ResolveType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResolveType.Descriptor instead.
func (ResolveType) EnumDescriptor() ([]byte, []int) {
	return file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_rawDescGZIP(), []int{0}
}

var File_tech_storezhang_dnsabre_core_kernel_resolve_type_proto protoreflect.FileDescriptor

var file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_rawDesc = []byte{
	0x0a, 0x36, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e,
	0x67, 0x2f, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6b,
	0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72,
	0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2a, 0xf0, 0x01,
	0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a,
	0x18, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x52,
	0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x10, 0x01, 0x12,
	0x15, 0x0a, 0x11, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x41, 0x41, 0x41, 0x41, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x03, 0x12, 0x14,
	0x0a, 0x10, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50,
	0x54, 0x52, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x58, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x53,
	0x4f, 0x4c, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x58, 0x54, 0x10, 0x06, 0x12,
	0x14, 0x0a, 0x10, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x52, 0x56, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x53, 0x10, 0x08, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x45,
	0x53, 0x4f, 0x4c, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4f, 0x41, 0x10, 0x09,
	0x42, 0x4e, 0x0a, 0x23, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68,
	0x61, 0x6e, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x50, 0x01, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x3b, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_rawDescOnce sync.Once
	file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_rawDescData = file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_rawDesc
)

func file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_rawDescGZIP() []byte {
	file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_rawDescOnce.Do(func() {
		file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_rawDescData)
	})
	return file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_rawDescData
}

var file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_goTypes = []interface{}{
	(ResolveType)(0), // 0: tech.storezhang.dnsabre.core.kernel.ResolveType
}
var file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_init() }
func file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_init() {
	if File_tech_storezhang_dnsabre_core_kernel_resolve_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_goTypes,
		DependencyIndexes: file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_depIdxs,
		EnumInfos:         file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_enumTypes,
	}.Build()
	File_tech_storezhang_dnsabre_core_kernel_resolve_type_proto = out.File
	file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_rawDesc = nil
	file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_goTypes = nil
	file_tech_storezhang_dnsabre_core_kernel_resolve_type_proto_depIdxs = nil
}
