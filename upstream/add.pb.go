// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: tech/storezhang/dnsabre/core/upstream/add.proto

package upstream

import (
	vo "github.com/dnsabre/core/vo"
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

// 创建请求
type AddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 名称
	// @gotags: validate:"required"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" validate:"required"`
	// 主机
	// @gotags: validate:"required,hostname|ip"
	Hostname string `protobuf:"bytes,4,opt,name=hostname,proto3" json:"hostname,omitempty" validate:"required,hostname|ip"`
	// 端口
	// @gotags: validate:"omitempty,min=1,max=65535"
	Port int32 `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty" validate:"omitempty,min=1,max=65535"`
}

func (x *AddReq) Reset() {
	*x = AddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tech_storezhang_dnsabre_core_upstream_add_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReq) ProtoMessage() {}

func (x *AddReq) ProtoReflect() protoreflect.Message {
	mi := &file_tech_storezhang_dnsabre_core_upstream_add_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReq.ProtoReflect.Descriptor instead.
func (*AddReq) Descriptor() ([]byte, []int) {
	return file_tech_storezhang_dnsabre_core_upstream_add_proto_rawDescGZIP(), []int{0}
}

func (x *AddReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddReq) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *AddReq) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

// 创建响应
type AddRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Upstream *vo.Upstream `protobuf:"bytes,3,opt,name=upstream,proto3" json:"upstream,omitempty"`
}

func (x *AddRsp) Reset() {
	*x = AddRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tech_storezhang_dnsabre_core_upstream_add_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRsp) ProtoMessage() {}

func (x *AddRsp) ProtoReflect() protoreflect.Message {
	mi := &file_tech_storezhang_dnsabre_core_upstream_add_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRsp.ProtoReflect.Descriptor instead.
func (*AddRsp) Descriptor() ([]byte, []int) {
	return file_tech_storezhang_dnsabre_core_upstream_add_proto_rawDescGZIP(), []int{1}
}

func (x *AddRsp) GetUpstream() *vo.Upstream {
	if x != nil {
		return x.Upstream
	}
	return nil
}

var File_tech_storezhang_dnsabre_core_upstream_add_proto protoreflect.FileDescriptor

var file_tech_storezhang_dnsabre_core_upstream_add_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e,
	0x67, 0x2f, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x75,
	0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x61, 0x64, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x25, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61,
	0x6e, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2f, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72,
	0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x6f, 0x2f, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x4f, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x73, 0x70,
	0x12, 0x45, 0x0a, 0x08, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a,
	0x68, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x6f, 0x2e, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x08, 0x75,
	0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x54, 0x0a, 0x25, 0x74, 0x65, 0x63, 0x68, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62,
	0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x75, 0x70, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x3b, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tech_storezhang_dnsabre_core_upstream_add_proto_rawDescOnce sync.Once
	file_tech_storezhang_dnsabre_core_upstream_add_proto_rawDescData = file_tech_storezhang_dnsabre_core_upstream_add_proto_rawDesc
)

func file_tech_storezhang_dnsabre_core_upstream_add_proto_rawDescGZIP() []byte {
	file_tech_storezhang_dnsabre_core_upstream_add_proto_rawDescOnce.Do(func() {
		file_tech_storezhang_dnsabre_core_upstream_add_proto_rawDescData = protoimpl.X.CompressGZIP(file_tech_storezhang_dnsabre_core_upstream_add_proto_rawDescData)
	})
	return file_tech_storezhang_dnsabre_core_upstream_add_proto_rawDescData
}

var file_tech_storezhang_dnsabre_core_upstream_add_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tech_storezhang_dnsabre_core_upstream_add_proto_goTypes = []interface{}{
	(*AddReq)(nil),      // 0: tech.storezhang.dnsabre.core.upstream.AddReq
	(*AddRsp)(nil),      // 1: tech.storezhang.dnsabre.core.upstream.AddRsp
	(*vo.Upstream)(nil), // 2: tech.storezhang.dnsabre.core.vo.Upstream
}
var file_tech_storezhang_dnsabre_core_upstream_add_proto_depIdxs = []int32{
	2, // 0: tech.storezhang.dnsabre.core.upstream.AddRsp.upstream:type_name -> tech.storezhang.dnsabre.core.vo.Upstream
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tech_storezhang_dnsabre_core_upstream_add_proto_init() }
func file_tech_storezhang_dnsabre_core_upstream_add_proto_init() {
	if File_tech_storezhang_dnsabre_core_upstream_add_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tech_storezhang_dnsabre_core_upstream_add_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReq); i {
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
		file_tech_storezhang_dnsabre_core_upstream_add_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRsp); i {
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
			RawDescriptor: file_tech_storezhang_dnsabre_core_upstream_add_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tech_storezhang_dnsabre_core_upstream_add_proto_goTypes,
		DependencyIndexes: file_tech_storezhang_dnsabre_core_upstream_add_proto_depIdxs,
		MessageInfos:      file_tech_storezhang_dnsabre_core_upstream_add_proto_msgTypes,
	}.Build()
	File_tech_storezhang_dnsabre_core_upstream_add_proto = out.File
	file_tech_storezhang_dnsabre_core_upstream_add_proto_rawDesc = nil
	file_tech_storezhang_dnsabre_core_upstream_add_proto_goTypes = nil
	file_tech_storezhang_dnsabre_core_upstream_add_proto_depIdxs = nil
}
