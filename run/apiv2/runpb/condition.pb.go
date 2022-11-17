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
// source: google/cloud/run/v2/condition.proto

package runpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents the possible Condition states.
type Condition_State int32

const (
	// The default value. This value is used if the state is omitted.
	Condition_STATE_UNSPECIFIED Condition_State = 0
	// Transient state: Reconciliation has not started yet.
	Condition_CONDITION_PENDING Condition_State = 1
	// Transient state: reconciliation is still in progress.
	Condition_CONDITION_RECONCILING Condition_State = 2
	// Terminal state: Reconciliation did not succeed.
	Condition_CONDITION_FAILED Condition_State = 3
	// Terminal state: Reconciliation completed successfully.
	Condition_CONDITION_SUCCEEDED Condition_State = 4
)

// Enum value maps for Condition_State.
var (
	Condition_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CONDITION_PENDING",
		2: "CONDITION_RECONCILING",
		3: "CONDITION_FAILED",
		4: "CONDITION_SUCCEEDED",
	}
	Condition_State_value = map[string]int32{
		"STATE_UNSPECIFIED":     0,
		"CONDITION_PENDING":     1,
		"CONDITION_RECONCILING": 2,
		"CONDITION_FAILED":      3,
		"CONDITION_SUCCEEDED":   4,
	}
)

func (x Condition_State) Enum() *Condition_State {
	p := new(Condition_State)
	*p = x
	return p
}

func (x Condition_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Condition_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_run_v2_condition_proto_enumTypes[0].Descriptor()
}

func (Condition_State) Type() protoreflect.EnumType {
	return &file_google_cloud_run_v2_condition_proto_enumTypes[0]
}

func (x Condition_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Condition_State.Descriptor instead.
func (Condition_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_run_v2_condition_proto_rawDescGZIP(), []int{0, 0}
}

// Represents the severity of the condition failures.
type Condition_Severity int32

const (
	// Unspecified severity
	Condition_SEVERITY_UNSPECIFIED Condition_Severity = 0
	// Error severity.
	Condition_ERROR Condition_Severity = 1
	// Warning severity.
	Condition_WARNING Condition_Severity = 2
	// Info severity.
	Condition_INFO Condition_Severity = 3
)

// Enum value maps for Condition_Severity.
var (
	Condition_Severity_name = map[int32]string{
		0: "SEVERITY_UNSPECIFIED",
		1: "ERROR",
		2: "WARNING",
		3: "INFO",
	}
	Condition_Severity_value = map[string]int32{
		"SEVERITY_UNSPECIFIED": 0,
		"ERROR":                1,
		"WARNING":              2,
		"INFO":                 3,
	}
)

func (x Condition_Severity) Enum() *Condition_Severity {
	p := new(Condition_Severity)
	*p = x
	return p
}

func (x Condition_Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Condition_Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_run_v2_condition_proto_enumTypes[1].Descriptor()
}

func (Condition_Severity) Type() protoreflect.EnumType {
	return &file_google_cloud_run_v2_condition_proto_enumTypes[1]
}

func (x Condition_Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Condition_Severity.Descriptor instead.
func (Condition_Severity) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_run_v2_condition_proto_rawDescGZIP(), []int{0, 1}
}

// Reasons common to all types of conditions.
type Condition_CommonReason int32

const (
	// Default value.
	Condition_COMMON_REASON_UNDEFINED Condition_CommonReason = 0
	// Reason unknown. Further details will be in message.
	Condition_UNKNOWN Condition_CommonReason = 1
	// Revision creation process failed.
	Condition_REVISION_FAILED Condition_CommonReason = 3
	// Timed out waiting for completion.
	Condition_PROGRESS_DEADLINE_EXCEEDED Condition_CommonReason = 4
	// The container image path is incorrect.
	Condition_CONTAINER_MISSING Condition_CommonReason = 6
	// Insufficient permissions on the container image.
	Condition_CONTAINER_PERMISSION_DENIED Condition_CommonReason = 7
	// Container image is not authorized by policy.
	Condition_CONTAINER_IMAGE_UNAUTHORIZED Condition_CommonReason = 8
	// Container image policy authorization check failed.
	Condition_CONTAINER_IMAGE_AUTHORIZATION_CHECK_FAILED Condition_CommonReason = 9
	// Insufficient permissions on encryption key.
	Condition_ENCRYPTION_KEY_PERMISSION_DENIED Condition_CommonReason = 10
	// Permission check on encryption key failed.
	Condition_ENCRYPTION_KEY_CHECK_FAILED Condition_CommonReason = 11
	// At least one Access check on secrets failed.
	Condition_SECRETS_ACCESS_CHECK_FAILED Condition_CommonReason = 12
	// Waiting for operation to complete.
	Condition_WAITING_FOR_OPERATION Condition_CommonReason = 13
	// System will retry immediately.
	Condition_IMMEDIATE_RETRY Condition_CommonReason = 14
	// System will retry later; current attempt failed.
	Condition_POSTPONED_RETRY Condition_CommonReason = 15
	// An internal error occurred. Further information may be in the message.
	Condition_INTERNAL Condition_CommonReason = 16
)

// Enum value maps for Condition_CommonReason.
var (
	Condition_CommonReason_name = map[int32]string{
		0:  "COMMON_REASON_UNDEFINED",
		1:  "UNKNOWN",
		3:  "REVISION_FAILED",
		4:  "PROGRESS_DEADLINE_EXCEEDED",
		6:  "CONTAINER_MISSING",
		7:  "CONTAINER_PERMISSION_DENIED",
		8:  "CONTAINER_IMAGE_UNAUTHORIZED",
		9:  "CONTAINER_IMAGE_AUTHORIZATION_CHECK_FAILED",
		10: "ENCRYPTION_KEY_PERMISSION_DENIED",
		11: "ENCRYPTION_KEY_CHECK_FAILED",
		12: "SECRETS_ACCESS_CHECK_FAILED",
		13: "WAITING_FOR_OPERATION",
		14: "IMMEDIATE_RETRY",
		15: "POSTPONED_RETRY",
		16: "INTERNAL",
	}
	Condition_CommonReason_value = map[string]int32{
		"COMMON_REASON_UNDEFINED":                    0,
		"UNKNOWN":                                    1,
		"REVISION_FAILED":                            3,
		"PROGRESS_DEADLINE_EXCEEDED":                 4,
		"CONTAINER_MISSING":                          6,
		"CONTAINER_PERMISSION_DENIED":                7,
		"CONTAINER_IMAGE_UNAUTHORIZED":               8,
		"CONTAINER_IMAGE_AUTHORIZATION_CHECK_FAILED": 9,
		"ENCRYPTION_KEY_PERMISSION_DENIED":           10,
		"ENCRYPTION_KEY_CHECK_FAILED":                11,
		"SECRETS_ACCESS_CHECK_FAILED":                12,
		"WAITING_FOR_OPERATION":                      13,
		"IMMEDIATE_RETRY":                            14,
		"POSTPONED_RETRY":                            15,
		"INTERNAL":                                   16,
	}
)

func (x Condition_CommonReason) Enum() *Condition_CommonReason {
	p := new(Condition_CommonReason)
	*p = x
	return p
}

func (x Condition_CommonReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Condition_CommonReason) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_run_v2_condition_proto_enumTypes[2].Descriptor()
}

func (Condition_CommonReason) Type() protoreflect.EnumType {
	return &file_google_cloud_run_v2_condition_proto_enumTypes[2]
}

func (x Condition_CommonReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Condition_CommonReason.Descriptor instead.
func (Condition_CommonReason) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_run_v2_condition_proto_rawDescGZIP(), []int{0, 2}
}

// Reasons specific to Revision resource.
type Condition_RevisionReason int32

const (
	// Default value.
	Condition_REVISION_REASON_UNDEFINED Condition_RevisionReason = 0
	// Revision in Pending state.
	Condition_PENDING Condition_RevisionReason = 1
	// Revision is in Reserve state.
	Condition_RESERVE Condition_RevisionReason = 2
	// Revision is Retired.
	Condition_RETIRED Condition_RevisionReason = 3
	// Revision is being retired.
	Condition_RETIRING Condition_RevisionReason = 4
	// Revision is being recreated.
	Condition_RECREATING Condition_RevisionReason = 5
	// There was a health check error.
	Condition_HEALTH_CHECK_CONTAINER_ERROR Condition_RevisionReason = 6
	// Health check failed due to user error from customized path of the
	// container. System will retry.
	Condition_CUSTOMIZED_PATH_RESPONSE_PENDING Condition_RevisionReason = 7
	// A revision with min_instance_count > 0 was created and is reserved, but
	// it was not configured to serve traffic, so it's not live. This can also
	// happen momentarily during traffic migration.
	Condition_MIN_INSTANCES_NOT_PROVISIONED Condition_RevisionReason = 8
	// The maximum allowed number of active revisions has been reached.
	Condition_ACTIVE_REVISION_LIMIT_REACHED Condition_RevisionReason = 9
	// There was no deployment defined.
	// This value is no longer used, but Services created in older versions of
	// the API might contain this value.
	Condition_NO_DEPLOYMENT Condition_RevisionReason = 10
	// A revision's container has no port specified since the revision is of a
	// manually scaled service with 0 instance count
	Condition_HEALTH_CHECK_SKIPPED Condition_RevisionReason = 11
)

// Enum value maps for Condition_RevisionReason.
var (
	Condition_RevisionReason_name = map[int32]string{
		0:  "REVISION_REASON_UNDEFINED",
		1:  "PENDING",
		2:  "RESERVE",
		3:  "RETIRED",
		4:  "RETIRING",
		5:  "RECREATING",
		6:  "HEALTH_CHECK_CONTAINER_ERROR",
		7:  "CUSTOMIZED_PATH_RESPONSE_PENDING",
		8:  "MIN_INSTANCES_NOT_PROVISIONED",
		9:  "ACTIVE_REVISION_LIMIT_REACHED",
		10: "NO_DEPLOYMENT",
		11: "HEALTH_CHECK_SKIPPED",
	}
	Condition_RevisionReason_value = map[string]int32{
		"REVISION_REASON_UNDEFINED":        0,
		"PENDING":                          1,
		"RESERVE":                          2,
		"RETIRED":                          3,
		"RETIRING":                         4,
		"RECREATING":                       5,
		"HEALTH_CHECK_CONTAINER_ERROR":     6,
		"CUSTOMIZED_PATH_RESPONSE_PENDING": 7,
		"MIN_INSTANCES_NOT_PROVISIONED":    8,
		"ACTIVE_REVISION_LIMIT_REACHED":    9,
		"NO_DEPLOYMENT":                    10,
		"HEALTH_CHECK_SKIPPED":             11,
	}
)

func (x Condition_RevisionReason) Enum() *Condition_RevisionReason {
	p := new(Condition_RevisionReason)
	*p = x
	return p
}

func (x Condition_RevisionReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Condition_RevisionReason) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_run_v2_condition_proto_enumTypes[3].Descriptor()
}

func (Condition_RevisionReason) Type() protoreflect.EnumType {
	return &file_google_cloud_run_v2_condition_proto_enumTypes[3]
}

func (x Condition_RevisionReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Condition_RevisionReason.Descriptor instead.
func (Condition_RevisionReason) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_run_v2_condition_proto_rawDescGZIP(), []int{0, 3}
}

// Reasons specific to Execution resource.
type Condition_ExecutionReason int32

const (
	// Default value.
	Condition_EXECUTION_REASON_UNDEFINED Condition_ExecutionReason = 0
	// Internal system error getting execution status. System will retry.
	Condition_JOB_STATUS_SERVICE_POLLING_ERROR Condition_ExecutionReason = 1
	// A task reached its retry limit and the last attempt failed due to the
	// user container exiting with a non-zero exit code.
	Condition_NON_ZERO_EXIT_CODE Condition_ExecutionReason = 2
	// The execution was cancelled by users.
	Condition_CANCELLED Condition_ExecutionReason = 3
)

// Enum value maps for Condition_ExecutionReason.
var (
	Condition_ExecutionReason_name = map[int32]string{
		0: "EXECUTION_REASON_UNDEFINED",
		1: "JOB_STATUS_SERVICE_POLLING_ERROR",
		2: "NON_ZERO_EXIT_CODE",
		3: "CANCELLED",
	}
	Condition_ExecutionReason_value = map[string]int32{
		"EXECUTION_REASON_UNDEFINED":       0,
		"JOB_STATUS_SERVICE_POLLING_ERROR": 1,
		"NON_ZERO_EXIT_CODE":               2,
		"CANCELLED":                        3,
	}
)

func (x Condition_ExecutionReason) Enum() *Condition_ExecutionReason {
	p := new(Condition_ExecutionReason)
	*p = x
	return p
}

func (x Condition_ExecutionReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Condition_ExecutionReason) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_run_v2_condition_proto_enumTypes[4].Descriptor()
}

func (Condition_ExecutionReason) Type() protoreflect.EnumType {
	return &file_google_cloud_run_v2_condition_proto_enumTypes[4]
}

func (x Condition_ExecutionReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Condition_ExecutionReason.Descriptor instead.
func (Condition_ExecutionReason) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_run_v2_condition_proto_rawDescGZIP(), []int{0, 4}
}

// Defines a status condition for a resource.
type Condition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type is used to communicate the status of the reconciliation process.
	// See also:
	// https://github.com/knative/serving/blob/main/docs/spec/errors.md#error-conditions-and-reporting
	// Types common to all resources include:
	// * "Ready": True when the Resource is ready.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// State of the condition.
	State Condition_State `protobuf:"varint,2,opt,name=state,proto3,enum=google.cloud.run.v2.Condition_State" json:"state,omitempty"`
	// Human readable message indicating details about the current status.
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=last_transition_time,json=lastTransitionTime,proto3" json:"last_transition_time,omitempty"`
	// How to interpret failures of this condition, one of Error, Warning, Info
	Severity Condition_Severity `protobuf:"varint,5,opt,name=severity,proto3,enum=google.cloud.run.v2.Condition_Severity" json:"severity,omitempty"`
	// The reason for this condition. Depending on the condition type,
	// it will populate one of these fields.
	// Successful conditions cannot have a reason.
	//
	// Types that are assignable to Reasons:
	//
	//	*Condition_Reason
	//	*Condition_RevisionReason_
	//	*Condition_ExecutionReason_
	Reasons isCondition_Reasons `protobuf_oneof:"reasons"`
}

func (x *Condition) Reset() {
	*x = Condition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v2_condition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Condition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Condition) ProtoMessage() {}

func (x *Condition) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v2_condition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Condition.ProtoReflect.Descriptor instead.
func (*Condition) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v2_condition_proto_rawDescGZIP(), []int{0}
}

func (x *Condition) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Condition) GetState() Condition_State {
	if x != nil {
		return x.State
	}
	return Condition_STATE_UNSPECIFIED
}

func (x *Condition) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Condition) GetLastTransitionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastTransitionTime
	}
	return nil
}

func (x *Condition) GetSeverity() Condition_Severity {
	if x != nil {
		return x.Severity
	}
	return Condition_SEVERITY_UNSPECIFIED
}

func (m *Condition) GetReasons() isCondition_Reasons {
	if m != nil {
		return m.Reasons
	}
	return nil
}

func (x *Condition) GetReason() Condition_CommonReason {
	if x, ok := x.GetReasons().(*Condition_Reason); ok {
		return x.Reason
	}
	return Condition_COMMON_REASON_UNDEFINED
}

func (x *Condition) GetRevisionReason() Condition_RevisionReason {
	if x, ok := x.GetReasons().(*Condition_RevisionReason_); ok {
		return x.RevisionReason
	}
	return Condition_REVISION_REASON_UNDEFINED
}

func (x *Condition) GetExecutionReason() Condition_ExecutionReason {
	if x, ok := x.GetReasons().(*Condition_ExecutionReason_); ok {
		return x.ExecutionReason
	}
	return Condition_EXECUTION_REASON_UNDEFINED
}

type isCondition_Reasons interface {
	isCondition_Reasons()
}

type Condition_Reason struct {
	// A common (service-level) reason for this condition.
	Reason Condition_CommonReason `protobuf:"varint,6,opt,name=reason,proto3,enum=google.cloud.run.v2.Condition_CommonReason,oneof"`
}

type Condition_RevisionReason_ struct {
	// A reason for the revision condition.
	RevisionReason Condition_RevisionReason `protobuf:"varint,9,opt,name=revision_reason,json=revisionReason,proto3,enum=google.cloud.run.v2.Condition_RevisionReason,oneof"`
}

type Condition_ExecutionReason_ struct {
	// A reason for the execution condition.
	ExecutionReason Condition_ExecutionReason `protobuf:"varint,11,opt,name=execution_reason,json=executionReason,proto3,enum=google.cloud.run.v2.Condition_ExecutionReason,oneof"`
}

func (*Condition_Reason) isCondition_Reasons() {}

func (*Condition_RevisionReason_) isCondition_Reasons() {}

func (*Condition_ExecutionReason_) isCondition_Reasons() {}

var File_google_cloud_run_v2_condition_proto protoreflect.FileDescriptor

var file_google_cloud_run_v2_condition_proto_rawDesc = []byte{
	0x0a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x75, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x0c, 0x0a, 0x09,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e,
	0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x4c, 0x0a, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x6c,
	0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x43, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x58, 0x0a,
	0x0f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x5b, 0x0a, 0x10, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x22, 0x7f, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a,
	0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x43,
	0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x4e, 0x43, 0x49,
	0x4c, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13,
	0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45,
	0x44, 0x45, 0x44, 0x10, 0x04, 0x22, 0x46, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e,
	0x47, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x03, 0x22, 0xb2, 0x03,
	0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1b,
	0x0a, 0x17, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f,
	0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x56, 0x49,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1e, 0x0a,
	0x1a, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x44, 0x45, 0x41, 0x44, 0x4c, 0x49,
	0x4e, 0x45, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x04, 0x12, 0x15, 0x0a,
	0x11, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x45, 0x52, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4e, 0x47, 0x10, 0x06, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x45,
	0x52, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4e,
	0x49, 0x45, 0x44, 0x10, 0x07, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e,
	0x45, 0x52, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f,
	0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x08, 0x12, 0x2e, 0x0a, 0x2a, 0x43, 0x4f, 0x4e, 0x54, 0x41,
	0x49, 0x4e, 0x45, 0x52, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f,
	0x52, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x09, 0x12, 0x24, 0x0a, 0x20, 0x45, 0x4e, 0x43, 0x52, 0x59,
	0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4b, 0x45, 0x59, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x1f, 0x0a,
	0x1b, 0x45, 0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4b, 0x45, 0x59, 0x5f,
	0x43, 0x48, 0x45, 0x43, 0x4b, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x1f,
	0x0a, 0x1b, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x53, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x0c, 0x12,
	0x19, 0x0a, 0x15, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x4f,
	0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0d, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4d,
	0x4d, 0x45, 0x44, 0x49, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x54, 0x52, 0x59, 0x10, 0x0e, 0x12,
	0x13, 0x0a, 0x0f, 0x50, 0x4f, 0x53, 0x54, 0x50, 0x4f, 0x4e, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x54,
	0x52, 0x59, 0x10, 0x0f, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c,
	0x10, 0x10, 0x22, 0xaf, 0x02, 0x0a, 0x0e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x56, 0x49, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x45, 0x10, 0x02, 0x12, 0x0b,
	0x0a, 0x07, 0x52, 0x45, 0x54, 0x49, 0x52, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x52,
	0x45, 0x54, 0x49, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x45, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x20, 0x0a, 0x1c, 0x48, 0x45, 0x41,
	0x4c, 0x54, 0x48, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49,
	0x4e, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x06, 0x12, 0x24, 0x0a, 0x20, 0x43,
	0x55, 0x53, 0x54, 0x4f, 0x4d, 0x49, 0x5a, 0x45, 0x44, 0x5f, 0x50, 0x41, 0x54, 0x48, 0x5f, 0x52,
	0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10,
	0x07, 0x12, 0x21, 0x0a, 0x1d, 0x4d, 0x49, 0x4e, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43,
	0x45, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e,
	0x45, 0x44, 0x10, 0x08, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x52,
	0x45, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x52, 0x45,
	0x41, 0x43, 0x48, 0x45, 0x44, 0x10, 0x09, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x5f, 0x44, 0x45,
	0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x0a, 0x12, 0x18, 0x0a, 0x14, 0x48, 0x45,
	0x41, 0x4c, 0x54, 0x48, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x5f, 0x53, 0x4b, 0x49, 0x50, 0x50,
	0x45, 0x44, 0x10, 0x0b, 0x22, 0x7e, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x58, 0x45, 0x43, 0x55,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x44, 0x45,
	0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x4a, 0x4f, 0x42, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x50, 0x4f,
	0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x16, 0x0a,
	0x12, 0x4e, 0x4f, 0x4e, 0x5f, 0x5a, 0x45, 0x52, 0x4f, 0x5f, 0x45, 0x58, 0x49, 0x54, 0x5f, 0x43,
	0x4f, 0x44, 0x45, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c,
	0x45, 0x44, 0x10, 0x03, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x42,
	0x63, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x42, 0x0e, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x76, 0x32,
	0x3b, 0x72, 0x75, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_run_v2_condition_proto_rawDescOnce sync.Once
	file_google_cloud_run_v2_condition_proto_rawDescData = file_google_cloud_run_v2_condition_proto_rawDesc
)

func file_google_cloud_run_v2_condition_proto_rawDescGZIP() []byte {
	file_google_cloud_run_v2_condition_proto_rawDescOnce.Do(func() {
		file_google_cloud_run_v2_condition_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_run_v2_condition_proto_rawDescData)
	})
	return file_google_cloud_run_v2_condition_proto_rawDescData
}

var file_google_cloud_run_v2_condition_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_google_cloud_run_v2_condition_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_run_v2_condition_proto_goTypes = []interface{}{
	(Condition_State)(0),           // 0: google.cloud.run.v2.Condition.State
	(Condition_Severity)(0),        // 1: google.cloud.run.v2.Condition.Severity
	(Condition_CommonReason)(0),    // 2: google.cloud.run.v2.Condition.CommonReason
	(Condition_RevisionReason)(0),  // 3: google.cloud.run.v2.Condition.RevisionReason
	(Condition_ExecutionReason)(0), // 4: google.cloud.run.v2.Condition.ExecutionReason
	(*Condition)(nil),              // 5: google.cloud.run.v2.Condition
	(*timestamppb.Timestamp)(nil),  // 6: google.protobuf.Timestamp
}
var file_google_cloud_run_v2_condition_proto_depIdxs = []int32{
	0, // 0: google.cloud.run.v2.Condition.state:type_name -> google.cloud.run.v2.Condition.State
	6, // 1: google.cloud.run.v2.Condition.last_transition_time:type_name -> google.protobuf.Timestamp
	1, // 2: google.cloud.run.v2.Condition.severity:type_name -> google.cloud.run.v2.Condition.Severity
	2, // 3: google.cloud.run.v2.Condition.reason:type_name -> google.cloud.run.v2.Condition.CommonReason
	3, // 4: google.cloud.run.v2.Condition.revision_reason:type_name -> google.cloud.run.v2.Condition.RevisionReason
	4, // 5: google.cloud.run.v2.Condition.execution_reason:type_name -> google.cloud.run.v2.Condition.ExecutionReason
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_run_v2_condition_proto_init() }
func file_google_cloud_run_v2_condition_proto_init() {
	if File_google_cloud_run_v2_condition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_run_v2_condition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition); i {
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
	file_google_cloud_run_v2_condition_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Condition_Reason)(nil),
		(*Condition_RevisionReason_)(nil),
		(*Condition_ExecutionReason_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_run_v2_condition_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_run_v2_condition_proto_goTypes,
		DependencyIndexes: file_google_cloud_run_v2_condition_proto_depIdxs,
		EnumInfos:         file_google_cloud_run_v2_condition_proto_enumTypes,
		MessageInfos:      file_google_cloud_run_v2_condition_proto_msgTypes,
	}.Build()
	File_google_cloud_run_v2_condition_proto = out.File
	file_google_cloud_run_v2_condition_proto_rawDesc = nil
	file_google_cloud_run_v2_condition_proto_goTypes = nil
	file_google_cloud_run_v2_condition_proto_depIdxs = nil
}
