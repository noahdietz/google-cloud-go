// Copyright 2022 Google LLC
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
// source: google/cloud/discoveryengine/v1alpha/grounded_generation_service.proto

package discoveryenginepb

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
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Specification for the grounding check.
type CheckGroundingSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The threshold (in [0,1]) used for determining whether a fact must be
	// cited for a claim in the answer candidate. Choosing a higher threshold
	// will lead to fewer but very strong citations, while choosing a lower
	// threshold may lead to more but somewhat weaker citations. If unset, the
	// threshold will default to 0.6.
	CitationThreshold *float64 `protobuf:"fixed64,1,opt,name=citation_threshold,json=citationThreshold,proto3,oneof" json:"citation_threshold,omitempty"`
}

func (x *CheckGroundingSpec) Reset() {
	*x = CheckGroundingSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckGroundingSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckGroundingSpec) ProtoMessage() {}

func (x *CheckGroundingSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckGroundingSpec.ProtoReflect.Descriptor instead.
func (*CheckGroundingSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_rawDescGZIP(), []int{0}
}

func (x *CheckGroundingSpec) GetCitationThreshold() float64 {
	if x != nil && x.CitationThreshold != nil {
		return *x.CitationThreshold
	}
	return 0
}

// Request message for
// [GroundedGenerationService.CheckGrounding][google.cloud.discoveryengine.v1alpha.GroundedGenerationService.CheckGrounding]
// method.
type CheckGroundingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the grounding config, such as
	// `projects/*/locations/global/groundingConfigs/default_grounding_config`.
	GroundingConfig string `protobuf:"bytes,1,opt,name=grounding_config,json=groundingConfig,proto3" json:"grounding_config,omitempty"`
	// Answer candidate to check.
	AnswerCandidate string `protobuf:"bytes,2,opt,name=answer_candidate,json=answerCandidate,proto3" json:"answer_candidate,omitempty"`
	// List of facts for the grounding check.
	// We support up to 200 facts.
	Facts []*GroundingFact `protobuf:"bytes,3,rep,name=facts,proto3" json:"facts,omitempty"`
	// Configuration of the grounding check.
	GroundingSpec *CheckGroundingSpec `protobuf:"bytes,4,opt,name=grounding_spec,json=groundingSpec,proto3" json:"grounding_spec,omitempty"`
}

func (x *CheckGroundingRequest) Reset() {
	*x = CheckGroundingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckGroundingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckGroundingRequest) ProtoMessage() {}

func (x *CheckGroundingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckGroundingRequest.ProtoReflect.Descriptor instead.
func (*CheckGroundingRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_rawDescGZIP(), []int{1}
}

func (x *CheckGroundingRequest) GetGroundingConfig() string {
	if x != nil {
		return x.GroundingConfig
	}
	return ""
}

func (x *CheckGroundingRequest) GetAnswerCandidate() string {
	if x != nil {
		return x.AnswerCandidate
	}
	return ""
}

func (x *CheckGroundingRequest) GetFacts() []*GroundingFact {
	if x != nil {
		return x.Facts
	}
	return nil
}

func (x *CheckGroundingRequest) GetGroundingSpec() *CheckGroundingSpec {
	if x != nil {
		return x.GroundingSpec
	}
	return nil
}

// Response message for the
// [GroundedGenerationService.CheckGrounding][google.cloud.discoveryengine.v1alpha.GroundedGenerationService.CheckGrounding]
// method.
type CheckGroundingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The support score for the input answer candidate.
	// Higher the score, higher is the fraction of claims that are supported by
	// the provided facts. This is always set when a response is returned.
	SupportScore *float32 `protobuf:"fixed32,1,opt,name=support_score,json=supportScore,proto3,oneof" json:"support_score,omitempty"`
	// List of facts cited across all claims in the answer candidate.
	// These are derived from the facts supplied in the request.
	CitedChunks []*FactChunk `protobuf:"bytes,3,rep,name=cited_chunks,json=citedChunks,proto3" json:"cited_chunks,omitempty"`
	// Claim texts and citation info across all claims in the answer candidate.
	Claims []*CheckGroundingResponse_Claim `protobuf:"bytes,4,rep,name=claims,proto3" json:"claims,omitempty"`
}

func (x *CheckGroundingResponse) Reset() {
	*x = CheckGroundingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckGroundingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckGroundingResponse) ProtoMessage() {}

func (x *CheckGroundingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckGroundingResponse.ProtoReflect.Descriptor instead.
func (*CheckGroundingResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_rawDescGZIP(), []int{2}
}

func (x *CheckGroundingResponse) GetSupportScore() float32 {
	if x != nil && x.SupportScore != nil {
		return *x.SupportScore
	}
	return 0
}

func (x *CheckGroundingResponse) GetCitedChunks() []*FactChunk {
	if x != nil {
		return x.CitedChunks
	}
	return nil
}

func (x *CheckGroundingResponse) GetClaims() []*CheckGroundingResponse_Claim {
	if x != nil {
		return x.Claims
	}
	return nil
}

// Text and citation info for a claim in the answer candidate.
type CheckGroundingResponse_Claim struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Position indicating the start of the claim in the answer candidate,
	// measured in bytes.
	StartPos *int32 `protobuf:"varint,1,opt,name=start_pos,json=startPos,proto3,oneof" json:"start_pos,omitempty"`
	// Position indicating the end of the claim in the answer candidate,
	// exclusive.
	EndPos *int32 `protobuf:"varint,2,opt,name=end_pos,json=endPos,proto3,oneof" json:"end_pos,omitempty"`
	// Text for the claim in the answer candidate. Always provided regardless of
	// whether citations or anti-citations are found.
	ClaimText string `protobuf:"bytes,3,opt,name=claim_text,json=claimText,proto3" json:"claim_text,omitempty"`
	// A list of indices (into 'cited_chunks') specifying the citations
	// associated with the claim. For instance [1,3,4] means that
	// cited_chunks[1], cited_chunks[3], cited_chunks[4] are the facts cited
	// supporting for the claim. A citation to a fact indicates that the claim
	// is supported by the fact.
	CitationIndices []int32 `protobuf:"varint,4,rep,packed,name=citation_indices,json=citationIndices,proto3" json:"citation_indices,omitempty"`
}

func (x *CheckGroundingResponse_Claim) Reset() {
	*x = CheckGroundingResponse_Claim{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckGroundingResponse_Claim) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckGroundingResponse_Claim) ProtoMessage() {}

func (x *CheckGroundingResponse_Claim) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckGroundingResponse_Claim.ProtoReflect.Descriptor instead.
func (*CheckGroundingResponse_Claim) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *CheckGroundingResponse_Claim) GetStartPos() int32 {
	if x != nil && x.StartPos != nil {
		return *x.StartPos
	}
	return 0
}

func (x *CheckGroundingResponse_Claim) GetEndPos() int32 {
	if x != nil && x.EndPos != nil {
		return *x.EndPos
	}
	return 0
}

func (x *CheckGroundingResponse_Claim) GetClaimText() string {
	if x != nil {
		return x.ClaimText
	}
	return ""
}

func (x *CheckGroundingResponse_Claim) GetCitationIndices() []int32 {
	if x != nil {
		return x.CitationIndices
	}
	return nil
}

var File_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto protoreflect.FileDescriptor

var file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_rawDesc = []byte{
	0x0a, 0x46, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x5f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x47, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x65, 0x63, 0x12, 0x32, 0x0a,
	0x12, 0x63, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x11, 0x63, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x88, 0x01,
	0x01, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x63, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x22, 0xd1, 0x02, 0x0a, 0x15, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x47, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x61, 0x0a, 0x10, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0xe0, 0x41,
	0x02, 0xfa, 0x41, 0x30, 0x0a, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x0f, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f,
	0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x49, 0x0a, 0x05, 0x66, 0x61, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x46, 0x61, 0x63, 0x74, 0x52, 0x05, 0x66, 0x61, 0x63, 0x74, 0x73, 0x12, 0x5f, 0x0a, 0x0e, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x47, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0d, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x65, 0x63, 0x22, 0xb2, 0x03, 0x0a,
	0x16, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x47, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0d, 0x73, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00,
	0x52, 0x0c, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x52, 0x0a, 0x0c, 0x63, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x46,
	0x61, 0x63, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x0b, 0x63, 0x69, 0x74, 0x65, 0x64, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x5a, 0x0a, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x47, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x73, 0x1a, 0xab, 0x01, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0x20, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00,
	0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a,
	0x07, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01,
	0x52, 0x06, 0x65, 0x6e, 0x64, 0x50, 0x6f, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x54, 0x65, 0x78, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x64, 0x69, 0x63, 0x65, 0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x70, 0x6f, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x32, 0xd5, 0x02, 0x0a, 0x19, 0x47, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0xe3, 0x01, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x47, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x12, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x47,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x47, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x56, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x50, 0x3a, 0x01, 0x2a, 0x22, 0x4b, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x7b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x2a, 0x7d, 0x3a,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x1a, 0x52, 0xca, 0x41, 0x1e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0xaa, 0x02, 0x0a, 0x28, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x1e, 0x47, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x52, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0xa2, 0x02, 0x0f, 0x44,
	0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0xaa, 0x02,
	0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x31,
	0x41, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xea, 0x02, 0x27, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_rawDescOnce sync.Once
	file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_rawDescData = file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_rawDesc
)

func file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_rawDescGZIP() []byte {
	file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_rawDescData)
	})
	return file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_rawDescData
}

var file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_goTypes = []interface{}{
	(*CheckGroundingSpec)(nil),           // 0: google.cloud.discoveryengine.v1alpha.CheckGroundingSpec
	(*CheckGroundingRequest)(nil),        // 1: google.cloud.discoveryengine.v1alpha.CheckGroundingRequest
	(*CheckGroundingResponse)(nil),       // 2: google.cloud.discoveryengine.v1alpha.CheckGroundingResponse
	(*CheckGroundingResponse_Claim)(nil), // 3: google.cloud.discoveryengine.v1alpha.CheckGroundingResponse.Claim
	(*GroundingFact)(nil),                // 4: google.cloud.discoveryengine.v1alpha.GroundingFact
	(*FactChunk)(nil),                    // 5: google.cloud.discoveryengine.v1alpha.FactChunk
}
var file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_depIdxs = []int32{
	4, // 0: google.cloud.discoveryengine.v1alpha.CheckGroundingRequest.facts:type_name -> google.cloud.discoveryengine.v1alpha.GroundingFact
	0, // 1: google.cloud.discoveryengine.v1alpha.CheckGroundingRequest.grounding_spec:type_name -> google.cloud.discoveryengine.v1alpha.CheckGroundingSpec
	5, // 2: google.cloud.discoveryengine.v1alpha.CheckGroundingResponse.cited_chunks:type_name -> google.cloud.discoveryengine.v1alpha.FactChunk
	3, // 3: google.cloud.discoveryengine.v1alpha.CheckGroundingResponse.claims:type_name -> google.cloud.discoveryengine.v1alpha.CheckGroundingResponse.Claim
	1, // 4: google.cloud.discoveryengine.v1alpha.GroundedGenerationService.CheckGrounding:input_type -> google.cloud.discoveryengine.v1alpha.CheckGroundingRequest
	2, // 5: google.cloud.discoveryengine.v1alpha.GroundedGenerationService.CheckGrounding:output_type -> google.cloud.discoveryengine.v1alpha.CheckGroundingResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_init() }
func file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_init() {
	if File_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto != nil {
		return
	}
	file_google_cloud_discoveryengine_v1alpha_grounding_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckGroundingSpec); i {
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
		file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckGroundingRequest); i {
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
		file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckGroundingResponse); i {
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
		file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckGroundingResponse_Claim); i {
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
	file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_msgTypes,
	}.Build()
	File_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto = out.File
	file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_rawDesc = nil
	file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_goTypes = nil
	file_google_cloud_discoveryengine_v1alpha_grounded_generation_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GroundedGenerationServiceClient is the client API for GroundedGenerationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GroundedGenerationServiceClient interface {
	// Performs a grounding check.
	CheckGrounding(ctx context.Context, in *CheckGroundingRequest, opts ...grpc.CallOption) (*CheckGroundingResponse, error)
}

type groundedGenerationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroundedGenerationServiceClient(cc grpc.ClientConnInterface) GroundedGenerationServiceClient {
	return &groundedGenerationServiceClient{cc}
}

func (c *groundedGenerationServiceClient) CheckGrounding(ctx context.Context, in *CheckGroundingRequest, opts ...grpc.CallOption) (*CheckGroundingResponse, error) {
	out := new(CheckGroundingResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.discoveryengine.v1alpha.GroundedGenerationService/CheckGrounding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroundedGenerationServiceServer is the server API for GroundedGenerationService service.
type GroundedGenerationServiceServer interface {
	// Performs a grounding check.
	CheckGrounding(context.Context, *CheckGroundingRequest) (*CheckGroundingResponse, error)
}

// UnimplementedGroundedGenerationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGroundedGenerationServiceServer struct {
}

func (*UnimplementedGroundedGenerationServiceServer) CheckGrounding(context.Context, *CheckGroundingRequest) (*CheckGroundingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckGrounding not implemented")
}

func RegisterGroundedGenerationServiceServer(s *grpc.Server, srv GroundedGenerationServiceServer) {
	s.RegisterService(&_GroundedGenerationService_serviceDesc, srv)
}

func _GroundedGenerationService_CheckGrounding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckGroundingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroundedGenerationServiceServer).CheckGrounding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.discoveryengine.v1alpha.GroundedGenerationService/CheckGrounding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroundedGenerationServiceServer).CheckGrounding(ctx, req.(*CheckGroundingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GroundedGenerationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.discoveryengine.v1alpha.GroundedGenerationService",
	HandlerType: (*GroundedGenerationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckGrounding",
			Handler:    _GroundedGenerationService_CheckGrounding_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/discoveryengine/v1alpha/grounded_generation_service.proto",
}
