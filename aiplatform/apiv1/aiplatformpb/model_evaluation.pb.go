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
// source: google/cloud/aiplatform/v1/model_evaluation.proto

package aiplatformpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A collection of metrics calculated by comparing Model's predictions on all of
// the test data against annotations from the test data.
type ModelEvaluation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the ModelEvaluation.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The display name of the ModelEvaluation.
	DisplayName string `protobuf:"bytes,10,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Points to a YAML file stored on Google Cloud Storage describing the
	// [metrics][google.cloud.aiplatform.v1.ModelEvaluation.metrics] of this
	// ModelEvaluation. The schema is defined as an OpenAPI 3.0.2 [Schema
	// Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	MetricsSchemaUri string `protobuf:"bytes,2,opt,name=metrics_schema_uri,json=metricsSchemaUri,proto3" json:"metrics_schema_uri,omitempty"`
	// Evaluation metrics of the Model. The schema of the metrics is stored in
	// [metrics_schema_uri][google.cloud.aiplatform.v1.ModelEvaluation.metrics_schema_uri]
	Metrics *structpb.Value `protobuf:"bytes,3,opt,name=metrics,proto3" json:"metrics,omitempty"`
	// Output only. Timestamp when this ModelEvaluation was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// All possible
	// [dimensions][google.cloud.aiplatform.v1.ModelEvaluationSlice.Slice.dimension]
	// of ModelEvaluationSlices. The dimensions can be used as the filter of the
	// [ModelService.ListModelEvaluationSlices][google.cloud.aiplatform.v1.ModelService.ListModelEvaluationSlices]
	// request, in the form of `slice.dimension = <dimension>`.
	SliceDimensions []string `protobuf:"bytes,5,rep,name=slice_dimensions,json=sliceDimensions,proto3" json:"slice_dimensions,omitempty"`
	// Points to a YAML file stored on Google Cloud Storage describing
	// [EvaluatedDataItemView.data_item_payload][] and
	// [EvaluatedAnnotation.data_item_payload][google.cloud.aiplatform.v1.EvaluatedAnnotation.data_item_payload].
	// The schema is defined as an OpenAPI 3.0.2 [Schema
	// Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	//
	// This field is not populated if there are neither EvaluatedDataItemViews nor
	// EvaluatedAnnotations under this ModelEvaluation.
	DataItemSchemaUri string `protobuf:"bytes,6,opt,name=data_item_schema_uri,json=dataItemSchemaUri,proto3" json:"data_item_schema_uri,omitempty"`
	// Points to a YAML file stored on Google Cloud Storage describing
	// [EvaluatedDataItemView.predictions][],
	// [EvaluatedDataItemView.ground_truths][],
	// [EvaluatedAnnotation.predictions][google.cloud.aiplatform.v1.EvaluatedAnnotation.predictions],
	// and
	// [EvaluatedAnnotation.ground_truths][google.cloud.aiplatform.v1.EvaluatedAnnotation.ground_truths].
	// The schema is defined as an OpenAPI 3.0.2 [Schema
	// Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	//
	// This field is not populated if there are neither EvaluatedDataItemViews nor
	// EvaluatedAnnotations under this ModelEvaluation.
	AnnotationSchemaUri string `protobuf:"bytes,7,opt,name=annotation_schema_uri,json=annotationSchemaUri,proto3" json:"annotation_schema_uri,omitempty"`
	// Aggregated explanation metrics for the Model's prediction output over the
	// data this ModelEvaluation uses. This field is populated only if the Model
	// is evaluated with explanations, and only for AutoML tabular Models.
	ModelExplanation *ModelExplanation `protobuf:"bytes,8,opt,name=model_explanation,json=modelExplanation,proto3" json:"model_explanation,omitempty"`
	// Describes the values of
	// [ExplanationSpec][google.cloud.aiplatform.v1.ExplanationSpec] that are used
	// for explaining the predicted values on the evaluated data.
	ExplanationSpecs []*ModelEvaluation_ModelEvaluationExplanationSpec `protobuf:"bytes,9,rep,name=explanation_specs,json=explanationSpecs,proto3" json:"explanation_specs,omitempty"`
	// The metadata of the ModelEvaluation.
	// For the ModelEvaluation uploaded from Managed Pipeline, metadata contains a
	// structured value with keys of "pipeline_job_id", "evaluation_dataset_type",
	// "evaluation_dataset_path", "row_based_metrics_path".
	Metadata *structpb.Value `protobuf:"bytes,11,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *ModelEvaluation) Reset() {
	*x = ModelEvaluation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_model_evaluation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelEvaluation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelEvaluation) ProtoMessage() {}

func (x *ModelEvaluation) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_model_evaluation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelEvaluation.ProtoReflect.Descriptor instead.
func (*ModelEvaluation) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_model_evaluation_proto_rawDescGZIP(), []int{0}
}

func (x *ModelEvaluation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModelEvaluation) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *ModelEvaluation) GetMetricsSchemaUri() string {
	if x != nil {
		return x.MetricsSchemaUri
	}
	return ""
}

func (x *ModelEvaluation) GetMetrics() *structpb.Value {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *ModelEvaluation) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ModelEvaluation) GetSliceDimensions() []string {
	if x != nil {
		return x.SliceDimensions
	}
	return nil
}

func (x *ModelEvaluation) GetDataItemSchemaUri() string {
	if x != nil {
		return x.DataItemSchemaUri
	}
	return ""
}

func (x *ModelEvaluation) GetAnnotationSchemaUri() string {
	if x != nil {
		return x.AnnotationSchemaUri
	}
	return ""
}

func (x *ModelEvaluation) GetModelExplanation() *ModelExplanation {
	if x != nil {
		return x.ModelExplanation
	}
	return nil
}

func (x *ModelEvaluation) GetExplanationSpecs() []*ModelEvaluation_ModelEvaluationExplanationSpec {
	if x != nil {
		return x.ExplanationSpecs
	}
	return nil
}

func (x *ModelEvaluation) GetMetadata() *structpb.Value {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type ModelEvaluation_ModelEvaluationExplanationSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Explanation type.
	//
	// For AutoML Image Classification models, possible values are:
	//
	//   - `image-integrated-gradients`
	//   - `image-xrai`
	ExplanationType string `protobuf:"bytes,1,opt,name=explanation_type,json=explanationType,proto3" json:"explanation_type,omitempty"`
	// Explanation spec details.
	ExplanationSpec *ExplanationSpec `protobuf:"bytes,2,opt,name=explanation_spec,json=explanationSpec,proto3" json:"explanation_spec,omitempty"`
}

func (x *ModelEvaluation_ModelEvaluationExplanationSpec) Reset() {
	*x = ModelEvaluation_ModelEvaluationExplanationSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_model_evaluation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelEvaluation_ModelEvaluationExplanationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelEvaluation_ModelEvaluationExplanationSpec) ProtoMessage() {}

func (x *ModelEvaluation_ModelEvaluationExplanationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_model_evaluation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelEvaluation_ModelEvaluationExplanationSpec.ProtoReflect.Descriptor instead.
func (*ModelEvaluation_ModelEvaluationExplanationSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_model_evaluation_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ModelEvaluation_ModelEvaluationExplanationSpec) GetExplanationType() string {
	if x != nil {
		return x.ExplanationType
	}
	return ""
}

func (x *ModelEvaluation_ModelEvaluationExplanationSpec) GetExplanationSpec() *ExplanationSpec {
	if x != nil {
		return x.ExplanationSpec
	}
	return nil
}

var File_google_cloud_aiplatform_v1_model_evaluation_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_model_evaluation_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x07, 0x0a, 0x0f, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x55, 0x72, 0x69, 0x12, 0x30, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x6c,
	0x69, 0x63, 0x65, 0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x44, 0x69, 0x6d, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x14, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x55, 0x72, 0x69, 0x12, 0x32, 0x0a, 0x15, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x75, 0x72, 0x69, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x55, 0x72, 0x69, 0x12, 0x59, 0x0a, 0x11, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x10, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x77, 0x0a, 0x11, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x4a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x70,
	0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x10, 0x65, 0x78,
	0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x73, 0x12, 0x32,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0xa3, 0x01, 0x0a, 0x1e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x45, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x56, 0x0a, 0x10, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0f, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x3a, 0x7f, 0xea, 0x41, 0x7c, 0x0a, 0x29, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x45, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x7d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x7b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x7d,
	0x2f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x65, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x42, 0xd2, 0x01, 0x0a, 0x1e, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0xea, 0x02,
	0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_model_evaluation_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_model_evaluation_proto_rawDescData = file_google_cloud_aiplatform_v1_model_evaluation_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_model_evaluation_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_model_evaluation_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_model_evaluation_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_model_evaluation_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_model_evaluation_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_model_evaluation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_aiplatform_v1_model_evaluation_proto_goTypes = []interface{}{
	(*ModelEvaluation)(nil),                                // 0: google.cloud.aiplatform.v1.ModelEvaluation
	(*ModelEvaluation_ModelEvaluationExplanationSpec)(nil), // 1: google.cloud.aiplatform.v1.ModelEvaluation.ModelEvaluationExplanationSpec
	(*structpb.Value)(nil),                                 // 2: google.protobuf.Value
	(*timestamppb.Timestamp)(nil),                          // 3: google.protobuf.Timestamp
	(*ModelExplanation)(nil),                               // 4: google.cloud.aiplatform.v1.ModelExplanation
	(*ExplanationSpec)(nil),                                // 5: google.cloud.aiplatform.v1.ExplanationSpec
}
var file_google_cloud_aiplatform_v1_model_evaluation_proto_depIdxs = []int32{
	2, // 0: google.cloud.aiplatform.v1.ModelEvaluation.metrics:type_name -> google.protobuf.Value
	3, // 1: google.cloud.aiplatform.v1.ModelEvaluation.create_time:type_name -> google.protobuf.Timestamp
	4, // 2: google.cloud.aiplatform.v1.ModelEvaluation.model_explanation:type_name -> google.cloud.aiplatform.v1.ModelExplanation
	1, // 3: google.cloud.aiplatform.v1.ModelEvaluation.explanation_specs:type_name -> google.cloud.aiplatform.v1.ModelEvaluation.ModelEvaluationExplanationSpec
	2, // 4: google.cloud.aiplatform.v1.ModelEvaluation.metadata:type_name -> google.protobuf.Value
	5, // 5: google.cloud.aiplatform.v1.ModelEvaluation.ModelEvaluationExplanationSpec.explanation_spec:type_name -> google.cloud.aiplatform.v1.ExplanationSpec
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1_model_evaluation_proto_init() }
func file_google_cloud_aiplatform_v1_model_evaluation_proto_init() {
	if File_google_cloud_aiplatform_v1_model_evaluation_proto != nil {
		return
	}
	file_google_cloud_aiplatform_v1_explanation_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1_model_evaluation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelEvaluation); i {
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
		file_google_cloud_aiplatform_v1_model_evaluation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelEvaluation_ModelEvaluationExplanationSpec); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1_model_evaluation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_model_evaluation_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_model_evaluation_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1_model_evaluation_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_model_evaluation_proto = out.File
	file_google_cloud_aiplatform_v1_model_evaluation_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_model_evaluation_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_model_evaluation_proto_depIdxs = nil
}
