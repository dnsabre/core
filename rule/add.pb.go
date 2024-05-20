// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: tech/storezhang/dnsabre/core/rule/add.proto

package rule

import (
	kernel "github.com/dnsabre/core/kernel"
	vo "github.com/dnsabre/core/vo"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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
	// 目标
	// @gotags: validate:"omitempty,url"
	Target string `protobuf:"bytes,4,opt,name=target,proto3" json:"target,omitempty" validate:"omitempty,url"`
	// 类型
	// @gotags: default:"1" validate:"omitempty,oneof=1"
	Type kernel.RuleType `protobuf:"varint,5,opt,name=type,proto3,enum=tech.storezhang.dnsabre.core.kernel.RuleType" json:"type,omitempty" default:"1" validate:"omitempty,oneof=1"`
	// 更新间隔
	// 默认为5个小时
	// @gotags: default:"{'seconds': 86400}"
	Interval *durationpb.Duration `protobuf:"bytes,6,opt,name=interval,proto3" json:"interval,omitempty" default:"{'seconds': 86400}"`
	// 信息
	// @gotags: default:"{'source': 1}"
	Info *Info `protobuf:"bytes,10,opt,name=info,proto3" json:"info,omitempty" default:"{'source': 1}"`
}

func (x *AddReq) Reset() {
	*x = AddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tech_storezhang_dnsabre_core_rule_add_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReq) ProtoMessage() {}

func (x *AddReq) ProtoReflect() protoreflect.Message {
	mi := &file_tech_storezhang_dnsabre_core_rule_add_proto_msgTypes[0]
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
	return file_tech_storezhang_dnsabre_core_rule_add_proto_rawDescGZIP(), []int{0}
}

func (x *AddReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddReq) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *AddReq) GetType() kernel.RuleType {
	if x != nil {
		return x.Type
	}
	return kernel.RuleType(0)
}

func (x *AddReq) GetInterval() *durationpb.Duration {
	if x != nil {
		return x.Interval
	}
	return nil
}

func (x *AddReq) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

// 创建响应
type AddRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rule *vo.Rule `protobuf:"bytes,3,opt,name=rule,proto3" json:"rule,omitempty"`
}

func (x *AddRsp) Reset() {
	*x = AddRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tech_storezhang_dnsabre_core_rule_add_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRsp) ProtoMessage() {}

func (x *AddRsp) ProtoReflect() protoreflect.Message {
	mi := &file_tech_storezhang_dnsabre_core_rule_add_proto_msgTypes[1]
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
	return file_tech_storezhang_dnsabre_core_rule_add_proto_rawDescGZIP(), []int{1}
}

func (x *AddRsp) GetRule() *vo.Rule {
	if x != nil {
		return x.Rule
	}
	return nil
}

var File_tech_storezhang_dnsabre_core_rule_add_proto protoreflect.FileDescriptor

var file_tech_storezhang_dnsabre_core_rule_add_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e,
	0x67, 0x2f, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72,
	0x75, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x74,
	0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x64,
	0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x75, 0x6c, 0x65,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x33, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e,
	0x67, 0x2f, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6b,
	0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2f, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a,
	0x68, 0x61, 0x6e, 0x67, 0x2f, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x6f, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xeb, 0x01, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x41, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x12, 0x3b, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67,
	0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x75,
	0x6c, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x43, 0x0a,
	0x06, 0x41, 0x64, 0x64, 0x52, 0x73, 0x70, 0x12, 0x39, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x6f, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x75,
	0x6c, 0x65, 0x42, 0x48, 0x0a, 0x21, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x3b, 0x72, 0x75, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tech_storezhang_dnsabre_core_rule_add_proto_rawDescOnce sync.Once
	file_tech_storezhang_dnsabre_core_rule_add_proto_rawDescData = file_tech_storezhang_dnsabre_core_rule_add_proto_rawDesc
)

func file_tech_storezhang_dnsabre_core_rule_add_proto_rawDescGZIP() []byte {
	file_tech_storezhang_dnsabre_core_rule_add_proto_rawDescOnce.Do(func() {
		file_tech_storezhang_dnsabre_core_rule_add_proto_rawDescData = protoimpl.X.CompressGZIP(file_tech_storezhang_dnsabre_core_rule_add_proto_rawDescData)
	})
	return file_tech_storezhang_dnsabre_core_rule_add_proto_rawDescData
}

var file_tech_storezhang_dnsabre_core_rule_add_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tech_storezhang_dnsabre_core_rule_add_proto_goTypes = []interface{}{
	(*AddReq)(nil),              // 0: tech.storezhang.dnsabre.core.rule.AddReq
	(*AddRsp)(nil),              // 1: tech.storezhang.dnsabre.core.rule.AddRsp
	(kernel.RuleType)(0),        // 2: tech.storezhang.dnsabre.core.kernel.RuleType
	(*durationpb.Duration)(nil), // 3: google.protobuf.Duration
	(*Info)(nil),                // 4: tech.storezhang.dnsabre.core.rule.Info
	(*vo.Rule)(nil),             // 5: tech.storezhang.dnsabre.core.vo.Rule
}
var file_tech_storezhang_dnsabre_core_rule_add_proto_depIdxs = []int32{
	2, // 0: tech.storezhang.dnsabre.core.rule.AddReq.type:type_name -> tech.storezhang.dnsabre.core.kernel.RuleType
	3, // 1: tech.storezhang.dnsabre.core.rule.AddReq.interval:type_name -> google.protobuf.Duration
	4, // 2: tech.storezhang.dnsabre.core.rule.AddReq.info:type_name -> tech.storezhang.dnsabre.core.rule.Info
	5, // 3: tech.storezhang.dnsabre.core.rule.AddRsp.rule:type_name -> tech.storezhang.dnsabre.core.vo.Rule
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tech_storezhang_dnsabre_core_rule_add_proto_init() }
func file_tech_storezhang_dnsabre_core_rule_add_proto_init() {
	if File_tech_storezhang_dnsabre_core_rule_add_proto != nil {
		return
	}
	file_tech_storezhang_dnsabre_core_rule_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tech_storezhang_dnsabre_core_rule_add_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_tech_storezhang_dnsabre_core_rule_add_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_tech_storezhang_dnsabre_core_rule_add_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tech_storezhang_dnsabre_core_rule_add_proto_goTypes,
		DependencyIndexes: file_tech_storezhang_dnsabre_core_rule_add_proto_depIdxs,
		MessageInfos:      file_tech_storezhang_dnsabre_core_rule_add_proto_msgTypes,
	}.Build()
	File_tech_storezhang_dnsabre_core_rule_add_proto = out.File
	file_tech_storezhang_dnsabre_core_rule_add_proto_rawDesc = nil
	file_tech_storezhang_dnsabre_core_rule_add_proto_goTypes = nil
	file_tech_storezhang_dnsabre_core_rule_add_proto_depIdxs = nil
}
