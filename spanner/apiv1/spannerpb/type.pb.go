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
// 	protoc        v3.21.5
// source: google/spanner/v1/type.proto

package spannerpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// `TypeCode` is used as part of [Type][google.spanner.v1.Type] to
// indicate the type of a Cloud Spanner value.
//
// Each legal value of a type can be encoded to or decoded from a JSON
// value, using the encodings described below. All Cloud Spanner values can
// be `null`, regardless of type; `null`s are always encoded as a JSON
// `null`.
type TypeCode int32

const (
	// Not specified.
	TypeCode_TYPE_CODE_UNSPECIFIED TypeCode = 0
	// Encoded as JSON `true` or `false`.
	TypeCode_BOOL TypeCode = 1
	// Encoded as `string`, in decimal format.
	TypeCode_INT64 TypeCode = 2
	// Encoded as `number`, or the strings `"NaN"`, `"Infinity"`, or
	// `"-Infinity"`.
	TypeCode_FLOAT64 TypeCode = 3
	// Encoded as `string` in RFC 3339 timestamp format. The time zone
	// must be present, and must be `"Z"`.
	//
	// If the schema has the column option
	// `allow_commit_timestamp=true`, the placeholder string
	// `"spanner.commit_timestamp()"` can be used to instruct the system
	// to insert the commit timestamp associated with the transaction
	// commit.
	TypeCode_TIMESTAMP TypeCode = 4
	// Encoded as `string` in RFC 3339 date format.
	TypeCode_DATE TypeCode = 5
	// Encoded as `string`.
	TypeCode_STRING TypeCode = 6
	// Encoded as a base64-encoded `string`, as described in RFC 4648,
	// section 4.
	TypeCode_BYTES TypeCode = 7
	// Encoded as `list`, where the list elements are represented
	// according to
	// [array_element_type][google.spanner.v1.Type.array_element_type].
	TypeCode_ARRAY TypeCode = 8
	// Encoded as `list`, where list element `i` is represented according
	// to [struct_type.fields[i]][google.spanner.v1.StructType.fields].
	TypeCode_STRUCT TypeCode = 9
	// Encoded as `string`, in decimal format or scientific notation format.
	// <br>Decimal format:
	// <br>`[+-]Digits[.[Digits]]` or
	// <br>`[+-][Digits].Digits`
	//
	// Scientific notation:
	// <br>`[+-]Digits[.[Digits]][ExponentIndicator[+-]Digits]` or
	// <br>`[+-][Digits].Digits[ExponentIndicator[+-]Digits]`
	// <br>(ExponentIndicator is `"e"` or `"E"`)
	TypeCode_NUMERIC TypeCode = 10
	// Encoded as a JSON-formatted `string` as described in RFC 7159. The
	// following rules are applied when parsing JSON input:
	//
	//   - Whitespace characters are not preserved.
	//   - If a JSON object has duplicate keys, only the first key is preserved.
	//   - Members of a JSON object are not guaranteed to have their order
	//     preserved.
	//   - JSON array elements will have their order preserved.
	TypeCode_JSON TypeCode = 11
)

// Enum value maps for TypeCode.
var (
	TypeCode_name = map[int32]string{
		0:  "TYPE_CODE_UNSPECIFIED",
		1:  "BOOL",
		2:  "INT64",
		3:  "FLOAT64",
		4:  "TIMESTAMP",
		5:  "DATE",
		6:  "STRING",
		7:  "BYTES",
		8:  "ARRAY",
		9:  "STRUCT",
		10: "NUMERIC",
		11: "JSON",
	}
	TypeCode_value = map[string]int32{
		"TYPE_CODE_UNSPECIFIED": 0,
		"BOOL":                  1,
		"INT64":                 2,
		"FLOAT64":               3,
		"TIMESTAMP":             4,
		"DATE":                  5,
		"STRING":                6,
		"BYTES":                 7,
		"ARRAY":                 8,
		"STRUCT":                9,
		"NUMERIC":               10,
		"JSON":                  11,
	}
)

func (x TypeCode) Enum() *TypeCode {
	p := new(TypeCode)
	*p = x
	return p
}

func (x TypeCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TypeCode) Descriptor() protoreflect.EnumDescriptor {
	return file_google_spanner_v1_type_proto_enumTypes[0].Descriptor()
}

func (TypeCode) Type() protoreflect.EnumType {
	return &file_google_spanner_v1_type_proto_enumTypes[0]
}

func (x TypeCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TypeCode.Descriptor instead.
func (TypeCode) EnumDescriptor() ([]byte, []int) {
	return file_google_spanner_v1_type_proto_rawDescGZIP(), []int{0}
}

// `TypeAnnotationCode` is used as a part of [Type][google.spanner.v1.Type] to
// disambiguate SQL types that should be used for a given Cloud Spanner value.
// Disambiguation is needed because the same Cloud Spanner type can be mapped to
// different SQL types depending on SQL dialect. TypeAnnotationCode doesn't
// affect the way value is serialized.
type TypeAnnotationCode int32

const (
	// Not specified.
	TypeAnnotationCode_TYPE_ANNOTATION_CODE_UNSPECIFIED TypeAnnotationCode = 0
	// PostgreSQL compatible NUMERIC type. This annotation needs to be applied to
	// [Type][google.spanner.v1.Type] instances having [NUMERIC][google.spanner.v1.TypeCode.NUMERIC]
	// type code to specify that values of this type should be treated as
	// PostgreSQL NUMERIC values. Currently this annotation is always needed for
	// [NUMERIC][google.spanner.v1.TypeCode.NUMERIC] when a client interacts with PostgreSQL-enabled
	// Spanner databases.
	TypeAnnotationCode_PG_NUMERIC TypeAnnotationCode = 2
	// PostgreSQL compatible JSONB type. This annotation needs to be applied to
	// [Type][google.spanner.v1.Type] instances having [JSON][google.spanner.v1.TypeCode.JSON]
	// type code to specify that values of this type should be treated as
	// PostgreSQL JSONB values. Currently this annotation is always needed for
	// [JSON][google.spanner.v1.TypeCode.JSON] when a client interacts with PostgreSQL-enabled
	// Spanner databases.
	TypeAnnotationCode_PG_JSONB TypeAnnotationCode = 3
)

// Enum value maps for TypeAnnotationCode.
var (
	TypeAnnotationCode_name = map[int32]string{
		0: "TYPE_ANNOTATION_CODE_UNSPECIFIED",
		2: "PG_NUMERIC",
		3: "PG_JSONB",
	}
	TypeAnnotationCode_value = map[string]int32{
		"TYPE_ANNOTATION_CODE_UNSPECIFIED": 0,
		"PG_NUMERIC":                       2,
		"PG_JSONB":                         3,
	}
)

func (x TypeAnnotationCode) Enum() *TypeAnnotationCode {
	p := new(TypeAnnotationCode)
	*p = x
	return p
}

func (x TypeAnnotationCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TypeAnnotationCode) Descriptor() protoreflect.EnumDescriptor {
	return file_google_spanner_v1_type_proto_enumTypes[1].Descriptor()
}

func (TypeAnnotationCode) Type() protoreflect.EnumType {
	return &file_google_spanner_v1_type_proto_enumTypes[1]
}

func (x TypeAnnotationCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TypeAnnotationCode.Descriptor instead.
func (TypeAnnotationCode) EnumDescriptor() ([]byte, []int) {
	return file_google_spanner_v1_type_proto_rawDescGZIP(), []int{1}
}

// `Type` indicates the type of a Cloud Spanner value, as might be stored in a
// table cell or returned from an SQL query.
type Type struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The [TypeCode][google.spanner.v1.TypeCode] for this type.
	Code TypeCode `protobuf:"varint,1,opt,name=code,proto3,enum=google.spanner.v1.TypeCode" json:"code,omitempty"`
	// If [code][google.spanner.v1.Type.code] == [ARRAY][google.spanner.v1.TypeCode.ARRAY], then `array_element_type`
	// is the type of the array elements.
	ArrayElementType *Type `protobuf:"bytes,2,opt,name=array_element_type,json=arrayElementType,proto3" json:"array_element_type,omitempty"`
	// If [code][google.spanner.v1.Type.code] == [STRUCT][google.spanner.v1.TypeCode.STRUCT], then `struct_type`
	// provides type information for the struct's fields.
	StructType *StructType `protobuf:"bytes,3,opt,name=struct_type,json=structType,proto3" json:"struct_type,omitempty"`
	// The [TypeAnnotationCode][google.spanner.v1.TypeAnnotationCode] that disambiguates SQL type that Spanner will
	// use to represent values of this type during query processing. This is
	// necessary for some type codes because a single [TypeCode][google.spanner.v1.TypeCode] can be mapped
	// to different SQL types depending on the SQL dialect. [type_annotation][google.spanner.v1.Type.type_annotation]
	// typically is not needed to process the content of a value (it doesn't
	// affect serialization) and clients can ignore it on the read path.
	TypeAnnotation TypeAnnotationCode `protobuf:"varint,4,opt,name=type_annotation,json=typeAnnotation,proto3,enum=google.spanner.v1.TypeAnnotationCode" json:"type_annotation,omitempty"`
}

func (x *Type) Reset() {
	*x = Type{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_spanner_v1_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Type) ProtoMessage() {}

func (x *Type) ProtoReflect() protoreflect.Message {
	mi := &file_google_spanner_v1_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Type.ProtoReflect.Descriptor instead.
func (*Type) Descriptor() ([]byte, []int) {
	return file_google_spanner_v1_type_proto_rawDescGZIP(), []int{0}
}

func (x *Type) GetCode() TypeCode {
	if x != nil {
		return x.Code
	}
	return TypeCode_TYPE_CODE_UNSPECIFIED
}

func (x *Type) GetArrayElementType() *Type {
	if x != nil {
		return x.ArrayElementType
	}
	return nil
}

func (x *Type) GetStructType() *StructType {
	if x != nil {
		return x.StructType
	}
	return nil
}

func (x *Type) GetTypeAnnotation() TypeAnnotationCode {
	if x != nil {
		return x.TypeAnnotation
	}
	return TypeAnnotationCode_TYPE_ANNOTATION_CODE_UNSPECIFIED
}

// `StructType` defines the fields of a [STRUCT][google.spanner.v1.TypeCode.STRUCT] type.
type StructType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of fields that make up this struct. Order is
	// significant, because values of this struct type are represented as
	// lists, where the order of field values matches the order of
	// fields in the [StructType][google.spanner.v1.StructType]. In turn, the order of fields
	// matches the order of columns in a read request, or the order of
	// fields in the `SELECT` clause of a query.
	Fields []*StructType_Field `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *StructType) Reset() {
	*x = StructType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_spanner_v1_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StructType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StructType) ProtoMessage() {}

func (x *StructType) ProtoReflect() protoreflect.Message {
	mi := &file_google_spanner_v1_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StructType.ProtoReflect.Descriptor instead.
func (*StructType) Descriptor() ([]byte, []int) {
	return file_google_spanner_v1_type_proto_rawDescGZIP(), []int{1}
}

func (x *StructType) GetFields() []*StructType_Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

// Message representing a single field of a struct.
type StructType_Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the field. For reads, this is the column name. For
	// SQL queries, it is the column alias (e.g., `"Word"` in the
	// query `"SELECT 'hello' AS Word"`), or the column name (e.g.,
	// `"ColName"` in the query `"SELECT ColName FROM Table"`). Some
	// columns might have an empty name (e.g., `"SELECT
	// UPPER(ColName)"`). Note that a query result can contain
	// multiple fields with the same name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The type of the field.
	Type *Type `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *StructType_Field) Reset() {
	*x = StructType_Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_spanner_v1_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StructType_Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StructType_Field) ProtoMessage() {}

func (x *StructType_Field) ProtoReflect() protoreflect.Message {
	mi := &file_google_spanner_v1_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StructType_Field.ProtoReflect.Descriptor instead.
func (*StructType_Field) Descriptor() ([]byte, []int) {
	return file_google_spanner_v1_type_proto_rawDescGZIP(), []int{1, 0}
}

func (x *StructType_Field) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StructType_Field) GetType() *Type {
	if x != nil {
		return x.Type
	}
	return nil
}

var File_google_spanner_v1_type_proto protoreflect.FileDescriptor

var file_google_spanner_v1_type_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x93, 0x02, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x45, 0x0a, 0x12, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x10, 0x61, 0x72, 0x72, 0x61, 0x79, 0x45, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0e, 0x74, 0x79, 0x70, 0x65, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x93, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x1a, 0x48, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a, 0xa5,
	0x01, 0x0a, 0x08, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x46,
	0x4c, 0x4f, 0x41, 0x54, 0x36, 0x34, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x49, 0x4d, 0x45,
	0x53, 0x54, 0x41, 0x4d, 0x50, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x41, 0x54, 0x45, 0x10,
	0x05, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x12, 0x09, 0x0a,
	0x05, 0x42, 0x59, 0x54, 0x45, 0x53, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x52, 0x52, 0x41,
	0x59, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x55, 0x43, 0x54, 0x10, 0x09, 0x12,
	0x0b, 0x0a, 0x07, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x49, 0x43, 0x10, 0x0a, 0x12, 0x08, 0x0a, 0x04,
	0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x0b, 0x2a, 0x58, 0x0a, 0x12, 0x54, 0x79, 0x70, 0x65, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x20,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x47, 0x5f, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x49, 0x43,
	0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x47, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x42, 0x10, 0x03,
	0x42, 0xaf, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x54, 0x79, 0x70, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73,
	0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0xaa, 0x02, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x53, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x17, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x70, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_spanner_v1_type_proto_rawDescOnce sync.Once
	file_google_spanner_v1_type_proto_rawDescData = file_google_spanner_v1_type_proto_rawDesc
)

func file_google_spanner_v1_type_proto_rawDescGZIP() []byte {
	file_google_spanner_v1_type_proto_rawDescOnce.Do(func() {
		file_google_spanner_v1_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_spanner_v1_type_proto_rawDescData)
	})
	return file_google_spanner_v1_type_proto_rawDescData
}

var file_google_spanner_v1_type_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_spanner_v1_type_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_spanner_v1_type_proto_goTypes = []interface{}{
	(TypeCode)(0),            // 0: google.spanner.v1.TypeCode
	(TypeAnnotationCode)(0),  // 1: google.spanner.v1.TypeAnnotationCode
	(*Type)(nil),             // 2: google.spanner.v1.Type
	(*StructType)(nil),       // 3: google.spanner.v1.StructType
	(*StructType_Field)(nil), // 4: google.spanner.v1.StructType.Field
}
var file_google_spanner_v1_type_proto_depIdxs = []int32{
	0, // 0: google.spanner.v1.Type.code:type_name -> google.spanner.v1.TypeCode
	2, // 1: google.spanner.v1.Type.array_element_type:type_name -> google.spanner.v1.Type
	3, // 2: google.spanner.v1.Type.struct_type:type_name -> google.spanner.v1.StructType
	1, // 3: google.spanner.v1.Type.type_annotation:type_name -> google.spanner.v1.TypeAnnotationCode
	4, // 4: google.spanner.v1.StructType.fields:type_name -> google.spanner.v1.StructType.Field
	2, // 5: google.spanner.v1.StructType.Field.type:type_name -> google.spanner.v1.Type
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_spanner_v1_type_proto_init() }
func file_google_spanner_v1_type_proto_init() {
	if File_google_spanner_v1_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_spanner_v1_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Type); i {
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
		file_google_spanner_v1_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StructType); i {
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
		file_google_spanner_v1_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StructType_Field); i {
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
			RawDescriptor: file_google_spanner_v1_type_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_spanner_v1_type_proto_goTypes,
		DependencyIndexes: file_google_spanner_v1_type_proto_depIdxs,
		EnumInfos:         file_google_spanner_v1_type_proto_enumTypes,
		MessageInfos:      file_google_spanner_v1_type_proto_msgTypes,
	}.Build()
	File_google_spanner_v1_type_proto = out.File
	file_google_spanner_v1_type_proto_rawDesc = nil
	file_google_spanner_v1_type_proto_goTypes = nil
	file_google_spanner_v1_type_proto_depIdxs = nil
}
