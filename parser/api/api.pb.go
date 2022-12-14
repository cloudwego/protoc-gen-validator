// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: api.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FieldRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Const       *string     `protobuf:"bytes,1,opt,name=const" json:"const,omitempty"`
	Lt          *string     `protobuf:"bytes,2,opt,name=lt" json:"lt,omitempty"`
	Le          *string     `protobuf:"bytes,3,opt,name=le" json:"le,omitempty"`
	Gt          *string     `protobuf:"bytes,4,opt,name=gt" json:"gt,omitempty"`
	Ge          *string     `protobuf:"bytes,5,opt,name=ge" json:"ge,omitempty"`
	In          []string    `protobuf:"bytes,6,rep,name=in" json:"in,omitempty"`
	NotIn       []string    `protobuf:"bytes,7,rep,name=not_in,json=notIn" json:"not_in,omitempty"`
	Len         *string     `protobuf:"bytes,8,opt,name=len" json:"len,omitempty"`
	MinSize     *string     `protobuf:"bytes,9,opt,name=min_size,json=minSize" json:"min_size,omitempty"`
	MaxSize     *string     `protobuf:"bytes,10,opt,name=max_size,json=maxSize" json:"max_size,omitempty"`
	Pattern     *string     `protobuf:"bytes,11,opt,name=pattern" json:"pattern,omitempty"`
	Prefix      *string     `protobuf:"bytes,12,opt,name=prefix" json:"prefix,omitempty"`
	Suffix      *string     `protobuf:"bytes,13,opt,name=suffix" json:"suffix,omitempty"`
	Contains    *string     `protobuf:"bytes,14,opt,name=contains" json:"contains,omitempty"`
	NotContains *string     `protobuf:"bytes,15,opt,name=not_contains,json=notContains" json:"not_contains,omitempty"`
	DefinedOnly *string     `protobuf:"bytes,16,opt,name=defined_only,json=definedOnly" json:"defined_only,omitempty"`
	NoSparse    *string     `protobuf:"bytes,17,opt,name=no_sparse,json=noSparse" json:"no_sparse,omitempty"`
	Key         *FieldRules `protobuf:"bytes,18,opt,name=key" json:"key,omitempty"`
	Value       *FieldRules `protobuf:"bytes,19,opt,name=value" json:"value,omitempty"`
	Elem        *FieldRules `protobuf:"bytes,20,opt,name=elem" json:"elem,omitempty"`
	Skip        *string     `protobuf:"bytes,21,opt,name=skip" json:"skip,omitempty"`
	Required    *string     `protobuf:"bytes,22,opt,name=required" json:"required,omitempty"`
	NotNil      *string     `protobuf:"bytes,23,opt,name=not_nil,json=notNil" json:"not_nil,omitempty"`
	Assert      *string     `protobuf:"bytes,24,opt,name=assert" json:"assert,omitempty"`
}

func (x *FieldRules) Reset() {
	*x = FieldRules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldRules) ProtoMessage() {}

func (x *FieldRules) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldRules.ProtoReflect.Descriptor instead.
func (*FieldRules) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *FieldRules) GetConst() string {
	if x != nil && x.Const != nil {
		return *x.Const
	}
	return ""
}

func (x *FieldRules) GetLt() string {
	if x != nil && x.Lt != nil {
		return *x.Lt
	}
	return ""
}

func (x *FieldRules) GetLe() string {
	if x != nil && x.Le != nil {
		return *x.Le
	}
	return ""
}

func (x *FieldRules) GetGt() string {
	if x != nil && x.Gt != nil {
		return *x.Gt
	}
	return ""
}

func (x *FieldRules) GetGe() string {
	if x != nil && x.Ge != nil {
		return *x.Ge
	}
	return ""
}

func (x *FieldRules) GetIn() []string {
	if x != nil {
		return x.In
	}
	return nil
}

func (x *FieldRules) GetNotIn() []string {
	if x != nil {
		return x.NotIn
	}
	return nil
}

func (x *FieldRules) GetLen() string {
	if x != nil && x.Len != nil {
		return *x.Len
	}
	return ""
}

func (x *FieldRules) GetMinSize() string {
	if x != nil && x.MinSize != nil {
		return *x.MinSize
	}
	return ""
}

func (x *FieldRules) GetMaxSize() string {
	if x != nil && x.MaxSize != nil {
		return *x.MaxSize
	}
	return ""
}

func (x *FieldRules) GetPattern() string {
	if x != nil && x.Pattern != nil {
		return *x.Pattern
	}
	return ""
}

func (x *FieldRules) GetPrefix() string {
	if x != nil && x.Prefix != nil {
		return *x.Prefix
	}
	return ""
}

func (x *FieldRules) GetSuffix() string {
	if x != nil && x.Suffix != nil {
		return *x.Suffix
	}
	return ""
}

func (x *FieldRules) GetContains() string {
	if x != nil && x.Contains != nil {
		return *x.Contains
	}
	return ""
}

func (x *FieldRules) GetNotContains() string {
	if x != nil && x.NotContains != nil {
		return *x.NotContains
	}
	return ""
}

func (x *FieldRules) GetDefinedOnly() string {
	if x != nil && x.DefinedOnly != nil {
		return *x.DefinedOnly
	}
	return ""
}

func (x *FieldRules) GetNoSparse() string {
	if x != nil && x.NoSparse != nil {
		return *x.NoSparse
	}
	return ""
}

func (x *FieldRules) GetKey() *FieldRules {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *FieldRules) GetValue() *FieldRules {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *FieldRules) GetElem() *FieldRules {
	if x != nil {
		return x.Elem
	}
	return nil
}

func (x *FieldRules) GetSkip() string {
	if x != nil && x.Skip != nil {
		return *x.Skip
	}
	return ""
}

func (x *FieldRules) GetRequired() string {
	if x != nil && x.Required != nil {
		return *x.Required
	}
	return ""
}

func (x *FieldRules) GetNotNil() string {
	if x != nil && x.NotNil != nil {
		return *x.NotNil
	}
	return ""
}

func (x *FieldRules) GetAssert() string {
	if x != nil && x.Assert != nil {
		return *x.Assert
	}
	return ""
}

var file_api_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50101,
		Name:          "api.raw_body",
		Tag:           "bytes,50101,opt,name=raw_body",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50102,
		Name:          "api.query",
		Tag:           "bytes,50102,opt,name=query",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50103,
		Name:          "api.header",
		Tag:           "bytes,50103,opt,name=header",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50104,
		Name:          "api.cookie",
		Tag:           "bytes,50104,opt,name=cookie",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50105,
		Name:          "api.body",
		Tag:           "bytes,50105,opt,name=body",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50106,
		Name:          "api.path",
		Tag:           "bytes,50106,opt,name=path",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50107,
		Name:          "api.vd",
		Tag:           "bytes,50107,opt,name=vd",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50108,
		Name:          "api.form",
		Tag:           "bytes,50108,opt,name=form",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51001,
		Name:          "api.go_tag",
		Tag:           "bytes,51001,opt,name=go_tag",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50109,
		Name:          "api.js_conv",
		Tag:           "bytes,50109,opt,name=js_conv",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldRules)(nil),
		Field:         50110,
		Name:          "api.vt",
		Tag:           "bytes,50110,opt,name=vt",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50201,
		Name:          "api.get",
		Tag:           "bytes,50201,opt,name=get",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50202,
		Name:          "api.post",
		Tag:           "bytes,50202,opt,name=post",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50203,
		Name:          "api.put",
		Tag:           "bytes,50203,opt,name=put",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50204,
		Name:          "api.delete",
		Tag:           "bytes,50204,opt,name=delete",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50205,
		Name:          "api.patch",
		Tag:           "bytes,50205,opt,name=patch",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50206,
		Name:          "api.options",
		Tag:           "bytes,50206,opt,name=options",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50207,
		Name:          "api.head",
		Tag:           "bytes,50207,opt,name=head",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50208,
		Name:          "api.any",
		Tag:           "bytes,50208,opt,name=any",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50301,
		Name:          "api.gen_path",
		Tag:           "bytes,50301,opt,name=gen_path",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50302,
		Name:          "api.api_version",
		Tag:           "bytes,50302,opt,name=api_version",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50303,
		Name:          "api.tag",
		Tag:           "bytes,50303,opt,name=tag",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50304,
		Name:          "api.name",
		Tag:           "bytes,50304,opt,name=name",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50305,
		Name:          "api.api_level",
		Tag:           "bytes,50305,opt,name=api_level",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50306,
		Name:          "api.serializer",
		Tag:           "bytes,50306,opt,name=serializer",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50307,
		Name:          "api.param",
		Tag:           "bytes,50307,opt,name=param",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50308,
		Name:          "api.baseurl",
		Tag:           "bytes,50308,opt,name=baseurl",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50309,
		Name:          "api.handler_path",
		Tag:           "bytes,50309,opt,name=handler_path",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         50401,
		Name:          "api.http_code",
		Tag:           "varint,50401,opt,name=http_code",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*FieldRules)(nil),
		Field:         50111,
		Name:          "api.msg_vt",
		Tag:           "bytes,50111,opt,name=msg_vt",
		Filename:      "api.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string raw_body = 50101;
	E_RawBody = &file_api_proto_extTypes[0]
	// optional string query = 50102;
	E_Query = &file_api_proto_extTypes[1]
	// optional string header = 50103;
	E_Header = &file_api_proto_extTypes[2]
	// optional string cookie = 50104;
	E_Cookie = &file_api_proto_extTypes[3]
	// optional string body = 50105;
	E_Body = &file_api_proto_extTypes[4]
	// optional string path = 50106;
	E_Path = &file_api_proto_extTypes[5]
	// optional string vd = 50107;
	E_Vd = &file_api_proto_extTypes[6]
	// optional string form = 50108;
	E_Form = &file_api_proto_extTypes[7]
	// optional string go_tag = 51001;
	E_GoTag = &file_api_proto_extTypes[8]
	// optional string js_conv = 50109;
	E_JsConv = &file_api_proto_extTypes[9]
	// optional api.FieldRules vt = 50110;
	E_Vt = &file_api_proto_extTypes[10]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional string get = 50201;
	E_Get = &file_api_proto_extTypes[11]
	// optional string post = 50202;
	E_Post = &file_api_proto_extTypes[12]
	// optional string put = 50203;
	E_Put = &file_api_proto_extTypes[13]
	// optional string delete = 50204;
	E_Delete = &file_api_proto_extTypes[14]
	// optional string patch = 50205;
	E_Patch = &file_api_proto_extTypes[15]
	// optional string options = 50206;
	E_Options = &file_api_proto_extTypes[16]
	// optional string head = 50207;
	E_Head = &file_api_proto_extTypes[17]
	// optional string any = 50208;
	E_Any = &file_api_proto_extTypes[18]
	// optional string gen_path = 50301;
	E_GenPath = &file_api_proto_extTypes[19] // The path specified by the user when the client code is generated, with a higher priority than api_version
	// optional string api_version = 50302;
	E_ApiVersion = &file_api_proto_extTypes[20] // Specify the value of the :version variable in path when the client code is generated
	// optional string tag = 50303;
	E_Tag = &file_api_proto_extTypes[21] // rpc tag, can be multiple, separated by commas
	// optional string name = 50304;
	E_Name = &file_api_proto_extTypes[22] // Name of rpc
	// optional string api_level = 50305;
	E_ApiLevel = &file_api_proto_extTypes[23] // Interface Level
	// optional string serializer = 50306;
	E_Serializer = &file_api_proto_extTypes[24] // Serialization method
	// optional string param = 50307;
	E_Param = &file_api_proto_extTypes[25] // Whether client requests take public parameters
	// optional string baseurl = 50308;
	E_Baseurl = &file_api_proto_extTypes[26] // Baseurl used in ttnet routing
	// optional string handler_path = 50309;
	E_HandlerPath = &file_api_proto_extTypes[27] // handler_path specifies the path to generate the method
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional int32 http_code = 50401;
	E_HttpCode = &file_api_proto_extTypes[28]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional api.FieldRules msg_vt = 50111;
	E_MsgVt = &file_api_proto_extTypes[29]
)

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xea, 0x04, 0x0a, 0x0a, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x75, 0x6c, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6c, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x67, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x67, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x67, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x6e, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x5f, 0x69,
	0x6e, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x49, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x65, 0x6e,
	0x12, 0x19, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x61, 0x78, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x61, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66,
	0x69, 0x78, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c,
	0x6e, 0x6f, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6e, 0x6f, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x4f, 0x6e,
	0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x5f, 0x73, 0x70, 0x61, 0x72, 0x73, 0x65, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x53, 0x70, 0x61, 0x72, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x75, 0x6c,
	0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x65, 0x6c, 0x65,
	0x6d, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x04, 0x65, 0x6c, 0x65, 0x6d, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6b,
	0x69, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x6e, 0x6f, 0x74, 0x5f, 0x6e, 0x69, 0x6c, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x6f, 0x74, 0x4e, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x73, 0x73, 0x65, 0x72,
	0x74, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x3a,
	0x3a, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb5, 0x87, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x61, 0x77, 0x42, 0x6f, 0x64, 0x79, 0x3a, 0x35, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xb6, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x3a, 0x37, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb7, 0x87, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x3a, 0x37, 0x0a, 0x06, 0x63,
	0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb8, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f,
	0x6f, 0x6b, 0x69, 0x65, 0x3a, 0x33, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb9, 0x87, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x3a, 0x33, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xba, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x3a, 0x2f,
	0x0a, 0x02, 0x76, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xbb, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x76, 0x64, 0x3a,
	0x33, 0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbc, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x36, 0x0a, 0x06, 0x67, 0x6f, 0x5f, 0x74, 0x61, 0x67, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb9, 0x8e,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x6f, 0x54, 0x61, 0x67, 0x3a, 0x38, 0x0a, 0x07,
	0x6a, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbd, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6a, 0x73, 0x43, 0x6f, 0x6e, 0x76, 0x3a, 0x40, 0x0a, 0x02, 0x76, 0x74, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbe, 0x87, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52,
	0x75, 0x6c, 0x65, 0x73, 0x52, 0x02, 0x76, 0x74, 0x3a, 0x32, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12,
	0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x99, 0x88, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x65, 0x74, 0x3a, 0x34, 0x0a, 0x04,
	0x70, 0x6f, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9a, 0x88, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f,
	0x73, 0x74, 0x3a, 0x32, 0x0a, 0x03, 0x70, 0x75, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9b, 0x88, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x70, 0x75, 0x74, 0x3a, 0x38, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x9c, 0x88, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x3a, 0x36, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9d, 0x88, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x3a, 0x3a, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x9e, 0x88, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x34, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64, 0x12, 0x1e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9f, 0x88, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x3a, 0x32, 0x0a, 0x03, 0x61, 0x6e,
	0x79, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xa0, 0x88, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x6e, 0x79, 0x3a, 0x3b,
	0x0a, 0x08, 0x67, 0x65, 0x6e, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfd, 0x88, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x67, 0x65, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x3a, 0x41, 0x0a, 0x0b, 0x61,
	0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfe, 0x88, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x32,
	0x0a, 0x03, 0x74, 0x61, 0x67, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xff, 0x88, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74,
	0x61, 0x67, 0x3a, 0x34, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x80, 0x89, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x3d, 0x0a, 0x09, 0x61, 0x70, 0x69, 0x5f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x81, 0x89, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x70, 0x69, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x3a, 0x40, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x82, 0x89, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x3a, 0x36, 0x0a, 0x05, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x83, 0x89, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x3a, 0x3a, 0x0a, 0x07, 0x62, 0x61, 0x73, 0x65, 0x75, 0x72, 0x6c, 0x12, 0x1e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x84, 0x89, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x75, 0x72, 0x6c, 0x3a, 0x43, 0x0a,
	0x0c, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x85, 0x89,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x50, 0x61,
	0x74, 0x68, 0x3a, 0x40, 0x0a, 0x09, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xe1, 0x89, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x68, 0x74, 0x74, 0x70,
	0x43, 0x6f, 0x64, 0x65, 0x3a, 0x49, 0x0a, 0x06, 0x6d, 0x73, 0x67, 0x5f, 0x76, 0x74, 0x12, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xbf, 0x87, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x56, 0x74, 0x42,
	0x06, 0x5a, 0x04, 0x2f, 0x61, 0x70, 0x69,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_proto_goTypes = []interface{}{
	(*FieldRules)(nil),                    // 0: api.FieldRules
	(*descriptorpb.FieldOptions)(nil),     // 1: google.protobuf.FieldOptions
	(*descriptorpb.MethodOptions)(nil),    // 2: google.protobuf.MethodOptions
	(*descriptorpb.EnumValueOptions)(nil), // 3: google.protobuf.EnumValueOptions
	(*descriptorpb.MessageOptions)(nil),   // 4: google.protobuf.MessageOptions
}
var file_api_proto_depIdxs = []int32{
	0,  // 0: api.FieldRules.key:type_name -> api.FieldRules
	0,  // 1: api.FieldRules.value:type_name -> api.FieldRules
	0,  // 2: api.FieldRules.elem:type_name -> api.FieldRules
	1,  // 3: api.raw_body:extendee -> google.protobuf.FieldOptions
	1,  // 4: api.query:extendee -> google.protobuf.FieldOptions
	1,  // 5: api.header:extendee -> google.protobuf.FieldOptions
	1,  // 6: api.cookie:extendee -> google.protobuf.FieldOptions
	1,  // 7: api.body:extendee -> google.protobuf.FieldOptions
	1,  // 8: api.path:extendee -> google.protobuf.FieldOptions
	1,  // 9: api.vd:extendee -> google.protobuf.FieldOptions
	1,  // 10: api.form:extendee -> google.protobuf.FieldOptions
	1,  // 11: api.go_tag:extendee -> google.protobuf.FieldOptions
	1,  // 12: api.js_conv:extendee -> google.protobuf.FieldOptions
	1,  // 13: api.vt:extendee -> google.protobuf.FieldOptions
	2,  // 14: api.get:extendee -> google.protobuf.MethodOptions
	2,  // 15: api.post:extendee -> google.protobuf.MethodOptions
	2,  // 16: api.put:extendee -> google.protobuf.MethodOptions
	2,  // 17: api.delete:extendee -> google.protobuf.MethodOptions
	2,  // 18: api.patch:extendee -> google.protobuf.MethodOptions
	2,  // 19: api.options:extendee -> google.protobuf.MethodOptions
	2,  // 20: api.head:extendee -> google.protobuf.MethodOptions
	2,  // 21: api.any:extendee -> google.protobuf.MethodOptions
	2,  // 22: api.gen_path:extendee -> google.protobuf.MethodOptions
	2,  // 23: api.api_version:extendee -> google.protobuf.MethodOptions
	2,  // 24: api.tag:extendee -> google.protobuf.MethodOptions
	2,  // 25: api.name:extendee -> google.protobuf.MethodOptions
	2,  // 26: api.api_level:extendee -> google.protobuf.MethodOptions
	2,  // 27: api.serializer:extendee -> google.protobuf.MethodOptions
	2,  // 28: api.param:extendee -> google.protobuf.MethodOptions
	2,  // 29: api.baseurl:extendee -> google.protobuf.MethodOptions
	2,  // 30: api.handler_path:extendee -> google.protobuf.MethodOptions
	3,  // 31: api.http_code:extendee -> google.protobuf.EnumValueOptions
	4,  // 32: api.msg_vt:extendee -> google.protobuf.MessageOptions
	0,  // 33: api.vt:type_name -> api.FieldRules
	0,  // 34: api.msg_vt:type_name -> api.FieldRules
	35, // [35:35] is the sub-list for method output_type
	35, // [35:35] is the sub-list for method input_type
	33, // [33:35] is the sub-list for extension type_name
	3,  // [3:33] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldRules); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 30,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
		ExtensionInfos:    file_api_proto_extTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
