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
// source: google/cloud/batch/v1/volume.proto

package batchpb

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

// Volume and mount parameters to be associated with a TaskSpec. A TaskSpec
// might describe zero, one, or multiple volumes to be mounted as part of the
// task.
type Volume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The source for the volume.
	//
	// Types that are assignable to Source:
	//
	//	*Volume_Nfs
	//	*Volume_Gcs
	//	*Volume_DeviceName
	Source isVolume_Source `protobuf_oneof:"source"`
	// Mount path for the volume, e.g. /mnt/share
	MountPath string `protobuf:"bytes,4,opt,name=mount_path,json=mountPath,proto3" json:"mount_path,omitempty"`
	// Mount options
	// For Google Cloud Storage, mount options are the global options supported by
	// gcsfuse tool. Batch will use them to mount the volume with the following
	// command:
	// "gcsfuse [global options] bucket mountpoint".
	// For PD, NFS, mount options are these supported by /etc/fstab. Batch will
	// use Fstab to mount such volumes.
	// https://help.ubuntu.com/community/Fstab
	MountOptions []string `protobuf:"bytes,5,rep,name=mount_options,json=mountOptions,proto3" json:"mount_options,omitempty"`
}

func (x *Volume) Reset() {
	*x = Volume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_batch_v1_volume_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Volume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Volume) ProtoMessage() {}

func (x *Volume) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_batch_v1_volume_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Volume.ProtoReflect.Descriptor instead.
func (*Volume) Descriptor() ([]byte, []int) {
	return file_google_cloud_batch_v1_volume_proto_rawDescGZIP(), []int{0}
}

func (m *Volume) GetSource() isVolume_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *Volume) GetNfs() *NFS {
	if x, ok := x.GetSource().(*Volume_Nfs); ok {
		return x.Nfs
	}
	return nil
}

func (x *Volume) GetGcs() *GCS {
	if x, ok := x.GetSource().(*Volume_Gcs); ok {
		return x.Gcs
	}
	return nil
}

func (x *Volume) GetDeviceName() string {
	if x, ok := x.GetSource().(*Volume_DeviceName); ok {
		return x.DeviceName
	}
	return ""
}

func (x *Volume) GetMountPath() string {
	if x != nil {
		return x.MountPath
	}
	return ""
}

func (x *Volume) GetMountOptions() []string {
	if x != nil {
		return x.MountOptions
	}
	return nil
}

type isVolume_Source interface {
	isVolume_Source()
}

type Volume_Nfs struct {
	// An NFS source for the volume (could be a Filestore, for example).
	Nfs *NFS `protobuf:"bytes,1,opt,name=nfs,proto3,oneof"`
}

type Volume_Gcs struct {
	// A Google Cloud Storage source for the volume.
	Gcs *GCS `protobuf:"bytes,3,opt,name=gcs,proto3,oneof"`
}

type Volume_DeviceName struct {
	// Device name of an attached disk
	DeviceName string `protobuf:"bytes,6,opt,name=device_name,json=deviceName,proto3,oneof"`
}

func (*Volume_Nfs) isVolume_Source() {}

func (*Volume_Gcs) isVolume_Source() {}

func (*Volume_DeviceName) isVolume_Source() {}

// Represents an NFS server and remote path: <server>:<remote_path>
type NFS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URI of the NFS server, e.g. an IP address.
	Server string `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	// Remote source path exported from NFS, e.g., "/share".
	RemotePath string `protobuf:"bytes,2,opt,name=remote_path,json=remotePath,proto3" json:"remote_path,omitempty"`
}

func (x *NFS) Reset() {
	*x = NFS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_batch_v1_volume_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NFS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NFS) ProtoMessage() {}

func (x *NFS) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_batch_v1_volume_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NFS.ProtoReflect.Descriptor instead.
func (*NFS) Descriptor() ([]byte, []int) {
	return file_google_cloud_batch_v1_volume_proto_rawDescGZIP(), []int{1}
}

func (x *NFS) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *NFS) GetRemotePath() string {
	if x != nil {
		return x.RemotePath
	}
	return ""
}

// Represents a Google Cloud Storage volume source config.
type GCS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Remote path, either a bucket name or a subdirectory of a bucket, e.g.:
	// bucket_name, bucket_name/subdirectory/
	RemotePath string `protobuf:"bytes,1,opt,name=remote_path,json=remotePath,proto3" json:"remote_path,omitempty"`
}

func (x *GCS) Reset() {
	*x = GCS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_batch_v1_volume_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCS) ProtoMessage() {}

func (x *GCS) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_batch_v1_volume_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCS.ProtoReflect.Descriptor instead.
func (*GCS) Descriptor() ([]byte, []int) {
	return file_google_cloud_batch_v1_volume_proto_rawDescGZIP(), []int{2}
}

func (x *GCS) GetRemotePath() string {
	if x != nil {
		return x.RemotePath
	}
	return ""
}

var File_google_cloud_batch_v1_volume_proto protoreflect.FileDescriptor

var file_google_cloud_batch_v1_volume_proto_rawDesc = []byte{
	0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x22, 0xd9, 0x01, 0x0a, 0x06,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x03, 0x6e, 0x66, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x46, 0x53, 0x48,
	0x00, 0x52, 0x03, 0x6e, 0x66, 0x73, 0x12, 0x2e, 0x0a, 0x03, 0x67, 0x63, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x43, 0x53, 0x48,
	0x00, 0x52, 0x03, 0x67, 0x63, 0x73, 0x12, 0x21, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0c, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x08, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x3e, 0x0a, 0x03, 0x4e, 0x46, 0x53, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x26, 0x0a, 0x03, 0x47, 0x43, 0x53, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x42,
	0xb7, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2f,
	0x76, 0x31, 0x3b, 0x62, 0x61, 0x74, 0x63, 0x68, 0xa2, 0x02, 0x03, 0x47, 0x43, 0x42, 0xaa, 0x02,
	0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x5c, 0x56, 0x31, 0xea, 0x02,
	0x18, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_batch_v1_volume_proto_rawDescOnce sync.Once
	file_google_cloud_batch_v1_volume_proto_rawDescData = file_google_cloud_batch_v1_volume_proto_rawDesc
)

func file_google_cloud_batch_v1_volume_proto_rawDescGZIP() []byte {
	file_google_cloud_batch_v1_volume_proto_rawDescOnce.Do(func() {
		file_google_cloud_batch_v1_volume_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_batch_v1_volume_proto_rawDescData)
	})
	return file_google_cloud_batch_v1_volume_proto_rawDescData
}

var file_google_cloud_batch_v1_volume_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_batch_v1_volume_proto_goTypes = []interface{}{
	(*Volume)(nil), // 0: google.cloud.batch.v1.Volume
	(*NFS)(nil),    // 1: google.cloud.batch.v1.NFS
	(*GCS)(nil),    // 2: google.cloud.batch.v1.GCS
}
var file_google_cloud_batch_v1_volume_proto_depIdxs = []int32{
	1, // 0: google.cloud.batch.v1.Volume.nfs:type_name -> google.cloud.batch.v1.NFS
	2, // 1: google.cloud.batch.v1.Volume.gcs:type_name -> google.cloud.batch.v1.GCS
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_batch_v1_volume_proto_init() }
func file_google_cloud_batch_v1_volume_proto_init() {
	if File_google_cloud_batch_v1_volume_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_batch_v1_volume_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Volume); i {
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
		file_google_cloud_batch_v1_volume_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NFS); i {
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
		file_google_cloud_batch_v1_volume_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCS); i {
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
	file_google_cloud_batch_v1_volume_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Volume_Nfs)(nil),
		(*Volume_Gcs)(nil),
		(*Volume_DeviceName)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_batch_v1_volume_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_batch_v1_volume_proto_goTypes,
		DependencyIndexes: file_google_cloud_batch_v1_volume_proto_depIdxs,
		MessageInfos:      file_google_cloud_batch_v1_volume_proto_msgTypes,
	}.Build()
	File_google_cloud_batch_v1_volume_proto = out.File
	file_google_cloud_batch_v1_volume_proto_rawDesc = nil
	file_google_cloud_batch_v1_volume_proto_goTypes = nil
	file_google_cloud_batch_v1_volume_proto_depIdxs = nil
}
