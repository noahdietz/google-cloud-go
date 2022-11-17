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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: google/cloud/run/v2/revision_template.proto

package runpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RevisionTemplate describes the data a revision should have when created from
// a template.
type RevisionTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique name for the revision. If this field is omitted, it will be
	// automatically generated based on the Service name.
	Revision string `protobuf:"bytes,1,opt,name=revision,proto3" json:"revision,omitempty"`
	// KRM-style labels for the resource.
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// KRM-style annotations for the resource.
	Annotations map[string]string `protobuf:"bytes,3,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Scaling settings for this Revision.
	Scaling *RevisionScaling `protobuf:"bytes,4,opt,name=scaling,proto3" json:"scaling,omitempty"`
	// VPC Access configuration to use for this Revision. For more information,
	// visit https://cloud.google.com/run/docs/configuring/connecting-vpc.
	VpcAccess *VpcAccess `protobuf:"bytes,6,opt,name=vpc_access,json=vpcAccess,proto3" json:"vpc_access,omitempty"`
	// Max allowed time for an instance to respond to a request.
	Timeout *durationpb.Duration `protobuf:"bytes,8,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Email address of the IAM service account associated with the revision of
	// the service. The service account represents the identity of the running
	// revision, and determines what permissions the revision has. If not
	// provided, the revision will use the project's default service account.
	ServiceAccount string `protobuf:"bytes,9,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
	// Holds the single container that defines the unit of execution for this
	// Revision.
	Containers []*Container `protobuf:"bytes,10,rep,name=containers,proto3" json:"containers,omitempty"`
	// A list of Volumes to make available to containers.
	Volumes []*Volume `protobuf:"bytes,11,rep,name=volumes,proto3" json:"volumes,omitempty"`
	// The sandbox environment to host this Revision.
	ExecutionEnvironment ExecutionEnvironment `protobuf:"varint,13,opt,name=execution_environment,json=executionEnvironment,proto3,enum=google.cloud.run.v2.ExecutionEnvironment" json:"execution_environment,omitempty"`
	// A reference to a customer managed encryption key (CMEK) to use to encrypt
	// this container image. For more information, go to
	// https://cloud.google.com/run/docs/securing/using-cmek
	EncryptionKey string `protobuf:"bytes,14,opt,name=encryption_key,json=encryptionKey,proto3" json:"encryption_key,omitempty"`
	// Sets the maximum number of requests that each serving instance can receive.
	MaxInstanceRequestConcurrency int32 `protobuf:"varint,15,opt,name=max_instance_request_concurrency,json=maxInstanceRequestConcurrency,proto3" json:"max_instance_request_concurrency,omitempty"`
}

func (x *RevisionTemplate) Reset() {
	*x = RevisionTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v2_revision_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevisionTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevisionTemplate) ProtoMessage() {}

func (x *RevisionTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v2_revision_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevisionTemplate.ProtoReflect.Descriptor instead.
func (*RevisionTemplate) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v2_revision_template_proto_rawDescGZIP(), []int{0}
}

func (x *RevisionTemplate) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *RevisionTemplate) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *RevisionTemplate) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *RevisionTemplate) GetScaling() *RevisionScaling {
	if x != nil {
		return x.Scaling
	}
	return nil
}

func (x *RevisionTemplate) GetVpcAccess() *VpcAccess {
	if x != nil {
		return x.VpcAccess
	}
	return nil
}

func (x *RevisionTemplate) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *RevisionTemplate) GetServiceAccount() string {
	if x != nil {
		return x.ServiceAccount
	}
	return ""
}

func (x *RevisionTemplate) GetContainers() []*Container {
	if x != nil {
		return x.Containers
	}
	return nil
}

func (x *RevisionTemplate) GetVolumes() []*Volume {
	if x != nil {
		return x.Volumes
	}
	return nil
}

func (x *RevisionTemplate) GetExecutionEnvironment() ExecutionEnvironment {
	if x != nil {
		return x.ExecutionEnvironment
	}
	return ExecutionEnvironment_EXECUTION_ENVIRONMENT_UNSPECIFIED
}

func (x *RevisionTemplate) GetEncryptionKey() string {
	if x != nil {
		return x.EncryptionKey
	}
	return ""
}

func (x *RevisionTemplate) GetMaxInstanceRequestConcurrency() int32 {
	if x != nil {
		return x.MaxInstanceRequestConcurrency
	}
	return 0
}

var File_google_cloud_run_v2_revision_template_proto protoreflect.FileDescriptor

var file_google_cloud_run_v2_revision_template_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x75, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e,
	0x76, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x75, 0x6e, 0x2f,
	0x76, 0x32, 0x2f, 0x6b, 0x38, 0x73, 0x2e, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x75, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x07, 0x0a, 0x10,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x12, 0x3c, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x20, 0xfa, 0x41, 0x1d, 0x0a, 0x1b, 0x72, 0x75, 0x6e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x49,
	0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75,
	0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x58, 0x0a, 0x0b, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75,
	0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x3e, 0x0a, 0x07, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x73, 0x63, 0x61, 0x6c,
	0x69, 0x6e, 0x67, 0x12, 0x3d, 0x0a, 0x0a, 0x76, 0x70, 0x63, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x70,
	0x63, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x09, 0x76, 0x70, 0x63, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x12, 0x35, 0x0a, 0x07, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x07,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x12, 0x5e, 0x0a, 0x15, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x14, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4d, 0x0a, 0x0e, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x26, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6b, 0x6d, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x52, 0x0d, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x47, 0x0a, 0x20, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x1d, 0x6d, 0x61, 0x78, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x1a,
	0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x6a, 0x0a, 0x17, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72,
	0x75, 0x6e, 0x2e, 0x76, 0x32, 0x42, 0x15, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x75, 0x6e, 0x2f,
	0x76, 0x32, 0x3b, 0x72, 0x75, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_run_v2_revision_template_proto_rawDescOnce sync.Once
	file_google_cloud_run_v2_revision_template_proto_rawDescData = file_google_cloud_run_v2_revision_template_proto_rawDesc
)

func file_google_cloud_run_v2_revision_template_proto_rawDescGZIP() []byte {
	file_google_cloud_run_v2_revision_template_proto_rawDescOnce.Do(func() {
		file_google_cloud_run_v2_revision_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_run_v2_revision_template_proto_rawDescData)
	})
	return file_google_cloud_run_v2_revision_template_proto_rawDescData
}

var file_google_cloud_run_v2_revision_template_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_run_v2_revision_template_proto_goTypes = []interface{}{
	(*RevisionTemplate)(nil),    // 0: google.cloud.run.v2.RevisionTemplate
	nil,                         // 1: google.cloud.run.v2.RevisionTemplate.LabelsEntry
	nil,                         // 2: google.cloud.run.v2.RevisionTemplate.AnnotationsEntry
	(*RevisionScaling)(nil),     // 3: google.cloud.run.v2.RevisionScaling
	(*VpcAccess)(nil),           // 4: google.cloud.run.v2.VpcAccess
	(*durationpb.Duration)(nil), // 5: google.protobuf.Duration
	(*Container)(nil),           // 6: google.cloud.run.v2.Container
	(*Volume)(nil),              // 7: google.cloud.run.v2.Volume
	(ExecutionEnvironment)(0),   // 8: google.cloud.run.v2.ExecutionEnvironment
}
var file_google_cloud_run_v2_revision_template_proto_depIdxs = []int32{
	1, // 0: google.cloud.run.v2.RevisionTemplate.labels:type_name -> google.cloud.run.v2.RevisionTemplate.LabelsEntry
	2, // 1: google.cloud.run.v2.RevisionTemplate.annotations:type_name -> google.cloud.run.v2.RevisionTemplate.AnnotationsEntry
	3, // 2: google.cloud.run.v2.RevisionTemplate.scaling:type_name -> google.cloud.run.v2.RevisionScaling
	4, // 3: google.cloud.run.v2.RevisionTemplate.vpc_access:type_name -> google.cloud.run.v2.VpcAccess
	5, // 4: google.cloud.run.v2.RevisionTemplate.timeout:type_name -> google.protobuf.Duration
	6, // 5: google.cloud.run.v2.RevisionTemplate.containers:type_name -> google.cloud.run.v2.Container
	7, // 6: google.cloud.run.v2.RevisionTemplate.volumes:type_name -> google.cloud.run.v2.Volume
	8, // 7: google.cloud.run.v2.RevisionTemplate.execution_environment:type_name -> google.cloud.run.v2.ExecutionEnvironment
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_cloud_run_v2_revision_template_proto_init() }
func file_google_cloud_run_v2_revision_template_proto_init() {
	if File_google_cloud_run_v2_revision_template_proto != nil {
		return
	}
	file_google_cloud_run_v2_k8s_min_proto_init()
	file_google_cloud_run_v2_vendor_settings_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_run_v2_revision_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevisionTemplate); i {
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
			RawDescriptor: file_google_cloud_run_v2_revision_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_run_v2_revision_template_proto_goTypes,
		DependencyIndexes: file_google_cloud_run_v2_revision_template_proto_depIdxs,
		MessageInfos:      file_google_cloud_run_v2_revision_template_proto_msgTypes,
	}.Build()
	File_google_cloud_run_v2_revision_template_proto = out.File
	file_google_cloud_run_v2_revision_template_proto_rawDesc = nil
	file_google_cloud_run_v2_revision_template_proto_goTypes = nil
	file_google_cloud_run_v2_revision_template_proto_depIdxs = nil
}
