// Copyright 2023 Google LLC
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
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: google/cloud/aiplatform/v1beta1/reasoning_engine_execution_service.proto

package aiplatformpb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for [ReasoningEngineExecutionService.Query][].
type QueryReasoningEngineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the ReasoningEngine resource to use.
	// Format:
	// `projects/{project}/locations/{location}/reasoningEngines/{reasoning_engine}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. Input content provided by users in JSON object format. Examples
	// include text query, function calling parameters, media bytes, etc.
	Input *structpb.Struct `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *QueryReasoningEngineRequest) Reset() {
	*x = QueryReasoningEngineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryReasoningEngineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryReasoningEngineRequest) ProtoMessage() {}

func (x *QueryReasoningEngineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryReasoningEngineRequest.ProtoReflect.Descriptor instead.
func (*QueryReasoningEngineRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_rawDescGZIP(), []int{0}
}

func (x *QueryReasoningEngineRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *QueryReasoningEngineRequest) GetInput() *structpb.Struct {
	if x != nil {
		return x.Input
	}
	return nil
}

// Response message for [ReasoningEngineExecutionService.Query][]
type QueryReasoningEngineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Response provided by users in JSON object format.
	Output *structpb.Value `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *QueryReasoningEngineResponse) Reset() {
	*x = QueryReasoningEngineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryReasoningEngineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryReasoningEngineResponse) ProtoMessage() {}

func (x *QueryReasoningEngineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryReasoningEngineResponse.ProtoReflect.Descriptor instead.
func (*QueryReasoningEngineResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_rawDescGZIP(), []int{1}
}

func (x *QueryReasoningEngineResponse) GetOutput() *structpb.Value {
	if x != nil {
		return x.Output
	}
	return nil
}

var File_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_rawDesc = []byte{
	0x0a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a,
	0x1b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x31, 0xe0, 0x41, 0x02, 0xfa,
	0x41, 0x2b, 0x0a, 0x29, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x01,
	0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x4e, 0x0a, 0x1c, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0xd2, 0x02, 0x0a, 0x1f, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xdf, 0x01, 0x0a, 0x14,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x12, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x69, 0x6e, 0x67, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x4a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x44, 0x3a, 0x01, 0x2a, 0x22, 0x3f, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x2a, 0x2f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x4d, 0xca,
	0x41, 0x19, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0xfb, 0x01, 0x0a,
	0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x42, 0x24, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70,
	0x62, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x42, 0x65,
	0x74, 0x61, 0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_goTypes = []interface{}{
	(*QueryReasoningEngineRequest)(nil),  // 0: google.cloud.aiplatform.v1beta1.QueryReasoningEngineRequest
	(*QueryReasoningEngineResponse)(nil), // 1: google.cloud.aiplatform.v1beta1.QueryReasoningEngineResponse
	(*structpb.Struct)(nil),              // 2: google.protobuf.Struct
	(*structpb.Value)(nil),               // 3: google.protobuf.Value
}
var file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_depIdxs = []int32{
	2, // 0: google.cloud.aiplatform.v1beta1.QueryReasoningEngineRequest.input:type_name -> google.protobuf.Struct
	3, // 1: google.cloud.aiplatform.v1beta1.QueryReasoningEngineResponse.output:type_name -> google.protobuf.Value
	0, // 2: google.cloud.aiplatform.v1beta1.ReasoningEngineExecutionService.QueryReasoningEngine:input_type -> google.cloud.aiplatform.v1beta1.QueryReasoningEngineRequest
	1, // 3: google.cloud.aiplatform.v1beta1.ReasoningEngineExecutionService.QueryReasoningEngine:output_type -> google.cloud.aiplatform.v1beta1.QueryReasoningEngineResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_init() }
func file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryReasoningEngineRequest); i {
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
		file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryReasoningEngineResponse); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto = out.File
	file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_reasoning_engine_execution_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReasoningEngineExecutionServiceClient is the client API for ReasoningEngineExecutionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReasoningEngineExecutionServiceClient interface {
	// Queries using a reasoning engine.
	QueryReasoningEngine(ctx context.Context, in *QueryReasoningEngineRequest, opts ...grpc.CallOption) (*QueryReasoningEngineResponse, error)
}

type reasoningEngineExecutionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReasoningEngineExecutionServiceClient(cc grpc.ClientConnInterface) ReasoningEngineExecutionServiceClient {
	return &reasoningEngineExecutionServiceClient{cc}
}

func (c *reasoningEngineExecutionServiceClient) QueryReasoningEngine(ctx context.Context, in *QueryReasoningEngineRequest, opts ...grpc.CallOption) (*QueryReasoningEngineResponse, error) {
	out := new(QueryReasoningEngineResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.aiplatform.v1beta1.ReasoningEngineExecutionService/QueryReasoningEngine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReasoningEngineExecutionServiceServer is the server API for ReasoningEngineExecutionService service.
type ReasoningEngineExecutionServiceServer interface {
	// Queries using a reasoning engine.
	QueryReasoningEngine(context.Context, *QueryReasoningEngineRequest) (*QueryReasoningEngineResponse, error)
}

// UnimplementedReasoningEngineExecutionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedReasoningEngineExecutionServiceServer struct {
}

func (*UnimplementedReasoningEngineExecutionServiceServer) QueryReasoningEngine(context.Context, *QueryReasoningEngineRequest) (*QueryReasoningEngineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryReasoningEngine not implemented")
}

func RegisterReasoningEngineExecutionServiceServer(s *grpc.Server, srv ReasoningEngineExecutionServiceServer) {
	s.RegisterService(&_ReasoningEngineExecutionService_serviceDesc, srv)
}

func _ReasoningEngineExecutionService_QueryReasoningEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryReasoningEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReasoningEngineExecutionServiceServer).QueryReasoningEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.aiplatform.v1beta1.ReasoningEngineExecutionService/QueryReasoningEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReasoningEngineExecutionServiceServer).QueryReasoningEngine(ctx, req.(*QueryReasoningEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReasoningEngineExecutionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.aiplatform.v1beta1.ReasoningEngineExecutionService",
	HandlerType: (*ReasoningEngineExecutionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryReasoningEngine",
			Handler:    _ReasoningEngineExecutionService_QueryReasoningEngine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/aiplatform/v1beta1/reasoning_engine_execution_service.proto",
}
