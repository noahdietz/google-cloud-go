// Copyright 2021 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: google/cloud/automl/v1/text_segment.proto

package automlpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A contiguous part of a text (string), assuming it has an UTF-8 NFC encoding.
type TextSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The content of the TextSegment.
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// Required. Zero-based character index of the first character of the text
	// segment (counting characters from the beginning of the text).
	StartOffset int64 `protobuf:"varint,1,opt,name=start_offset,json=startOffset,proto3" json:"start_offset,omitempty"`
	// Required. Zero-based character index of the first character past the end of
	// the text segment (counting character from the beginning of the text).
	// The character at the end_offset is NOT included in the text segment.
	EndOffset int64 `protobuf:"varint,2,opt,name=end_offset,json=endOffset,proto3" json:"end_offset,omitempty"`
}

func (x *TextSegment) Reset() {
	*x = TextSegment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1_text_segment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextSegment) ProtoMessage() {}

func (x *TextSegment) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1_text_segment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextSegment.ProtoReflect.Descriptor instead.
func (*TextSegment) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1_text_segment_proto_rawDescGZIP(), []int{0}
}

func (x *TextSegment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *TextSegment) GetStartOffset() int64 {
	if x != nil {
		return x.StartOffset
	}
	return 0
}

func (x *TextSegment) GetEndOffset() int64 {
	if x != nil {
		return x.EndOffset
	}
	return 0
}

var File_google_cloud_automl_v1_text_segment_proto protoreflect.FileDescriptor

var file_google_cloud_automl_v1_text_segment_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x65,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c,
	0x2e, 0x76, 0x31, 0x22, 0x69, 0x0a, 0x0b, 0x54, 0x65, 0x78, 0x74, 0x53, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x42, 0xbc,
	0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x54,
	0x65, 0x78, 0x74, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0xaa,
	0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41,
	0x75, 0x74, 0x6f, 0x4d, 0x4c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x6c, 0x5c, 0x56,
	0x31, 0xea, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x4c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_automl_v1_text_segment_proto_rawDescOnce sync.Once
	file_google_cloud_automl_v1_text_segment_proto_rawDescData = file_google_cloud_automl_v1_text_segment_proto_rawDesc
)

func file_google_cloud_automl_v1_text_segment_proto_rawDescGZIP() []byte {
	file_google_cloud_automl_v1_text_segment_proto_rawDescOnce.Do(func() {
		file_google_cloud_automl_v1_text_segment_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_automl_v1_text_segment_proto_rawDescData)
	})
	return file_google_cloud_automl_v1_text_segment_proto_rawDescData
}

var file_google_cloud_automl_v1_text_segment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_automl_v1_text_segment_proto_goTypes = []interface{}{
	(*TextSegment)(nil), // 0: google.cloud.automl.v1.TextSegment
}
var file_google_cloud_automl_v1_text_segment_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_automl_v1_text_segment_proto_init() }
func file_google_cloud_automl_v1_text_segment_proto_init() {
	if File_google_cloud_automl_v1_text_segment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_automl_v1_text_segment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextSegment); i {
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
			RawDescriptor: file_google_cloud_automl_v1_text_segment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_automl_v1_text_segment_proto_goTypes,
		DependencyIndexes: file_google_cloud_automl_v1_text_segment_proto_depIdxs,
		MessageInfos:      file_google_cloud_automl_v1_text_segment_proto_msgTypes,
	}.Build()
	File_google_cloud_automl_v1_text_segment_proto = out.File
	file_google_cloud_automl_v1_text_segment_proto_rawDesc = nil
	file_google_cloud_automl_v1_text_segment_proto_goTypes = nil
	file_google_cloud_automl_v1_text_segment_proto_depIdxs = nil
}
