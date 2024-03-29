// protoc生成代码兼容proto2的语法

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: base.proto

package psm

import (
	context "context"
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

type EnumType2 int32

const (
	EnumType2_TWEET2   EnumType2 = 0
	EnumType2_RETWEET2 EnumType2 = 1
)

// Enum value maps for EnumType2.
var (
	EnumType2_name = map[int32]string{
		0: "TWEET2",
		1: "RETWEET2",
	}
	EnumType2_value = map[string]int32{
		"TWEET2":   0,
		"RETWEET2": 1,
	}
)

func (x EnumType2) Enum() *EnumType2 {
	p := new(EnumType2)
	*p = x
	return p
}

func (x EnumType2) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumType2) Descriptor() protoreflect.EnumDescriptor {
	return file_base_proto_enumTypes[0].Descriptor()
}

func (EnumType2) Type() protoreflect.EnumType {
	return &file_base_proto_enumTypes[0]
}

func (x EnumType2) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumType2.Descriptor instead.
func (EnumType2) EnumDescriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

var File_base_proto protoreflect.FileDescriptor

var file_base_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x70, 0x73,
	0x6d, 0x2a, 0x25, 0x0a, 0x09, 0x45, 0x6e, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x32, 0x12, 0x0a,
	0x0a, 0x06, 0x54, 0x57, 0x45, 0x45, 0x54, 0x32, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45,
	0x54, 0x57, 0x45, 0x45, 0x54, 0x32, 0x10, 0x01, 0x42, 0x15, 0x5a, 0x13, 0x61, 0x2f, 0x62, 0x2f,
	0x63, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x73, 0x6d, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_proto_rawDescOnce sync.Once
	file_base_proto_rawDescData = file_base_proto_rawDesc
)

func file_base_proto_rawDescGZIP() []byte {
	file_base_proto_rawDescOnce.Do(func() {
		file_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_proto_rawDescData)
	})
	return file_base_proto_rawDescData
}

var file_base_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_base_proto_goTypes = []interface{}{
	(EnumType2)(0), // 0: psm.EnumType2
}
var file_base_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_base_proto_init() }
func file_base_proto_init() {
	if File_base_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_base_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_proto_goTypes,
		DependencyIndexes: file_base_proto_depIdxs,
		EnumInfos:         file_base_proto_enumTypes,
	}.Build()
	File_base_proto = out.File
	file_base_proto_rawDesc = nil
	file_base_proto_goTypes = nil
	file_base_proto_depIdxs = nil
}

var _ context.Context
