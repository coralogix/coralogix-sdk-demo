// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: com/coralogix/quota/v1/policies_service.proto

package policies

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_com_coralogix_quota_v1_policies_service_proto protoreflect.FileDescriptor

var file_com_coralogix_quota_v1_policies_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71,
	0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x26, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71,
	0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x30, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71,
	0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74,
	0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x3a, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71,
	0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71,
	0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x6c, 0x6b,
	0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x3c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x6c, 0x6b, 0x5f, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75,
	0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x67, 0x67, 0x6c,
	0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xcb, 0x08, 0x0a, 0x0f, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x72, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0xc2,
	0xb8, 0x02, 0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x20, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x7e, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0xc2, 0xb8, 0x02, 0x0f,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x7e, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0xc2, 0xb8, 0x02, 0x0f,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x97, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0xc2,
	0xb8, 0x02, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x20, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x20, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x7e, 0x0a, 0x0c, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0xc2, 0xb8, 0x02, 0x0f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x20, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x8a, 0x01, 0x0a, 0x0f, 0x52, 0x65,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x2e, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75,
	0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75,
	0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16,
	0xc2, 0xb8, 0x02, 0x12, 0x0a, 0x10, 0x52, 0x65, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x20, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x9c, 0x01, 0x0a, 0x13, 0x42, 0x75, 0x6c, 0x6b, 0x54,
	0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x32,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71,
	0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x54, 0x65, 0x73, 0x74,
	0x4c, 0x6f, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x6c, 0x6b,
	0x54, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0xc2, 0xb8, 0x02, 0x18, 0x0a, 0x16, 0x42,
	0x75, 0x6c, 0x6b, 0x20, 0x54, 0x65, 0x73, 0x74, 0x20, 0x4c, 0x6f, 0x67, 0x20, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x7e, 0x0a, 0x0c, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x6f, 0x67, 0x67, 0x6c, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x67, 0x67,
	0x6c, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x13, 0xc2, 0xb8, 0x02, 0x0f, 0x0a, 0x0d, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x20, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x18, 0x5a, 0x16, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_com_coralogix_quota_v1_policies_service_proto_goTypes = []interface{}{
	(*GetPolicyRequest)(nil),            // 0: com.coralogix.quota.v1.GetPolicyRequest
	(*CreatePolicyRequest)(nil),         // 1: com.coralogix.quota.v1.CreatePolicyRequest
	(*UpdatePolicyRequest)(nil),         // 2: com.coralogix.quota.v1.UpdatePolicyRequest
	(*GetCompanyPoliciesRequest)(nil),   // 3: com.coralogix.quota.v1.GetCompanyPoliciesRequest
	(*DeletePolicyRequest)(nil),         // 4: com.coralogix.quota.v1.DeletePolicyRequest
	(*ReorderPoliciesRequest)(nil),      // 5: com.coralogix.quota.v1.ReorderPoliciesRequest
	(*BulkTestLogPoliciesRequest)(nil),  // 6: com.coralogix.quota.v1.BulkTestLogPoliciesRequest
	(*TogglePolicyRequest)(nil),         // 7: com.coralogix.quota.v1.TogglePolicyRequest
	(*GetPolicyResponse)(nil),           // 8: com.coralogix.quota.v1.GetPolicyResponse
	(*CreatePolicyResponse)(nil),        // 9: com.coralogix.quota.v1.CreatePolicyResponse
	(*UpdatePolicyResponse)(nil),        // 10: com.coralogix.quota.v1.UpdatePolicyResponse
	(*GetCompanyPoliciesResponse)(nil),  // 11: com.coralogix.quota.v1.GetCompanyPoliciesResponse
	(*DeletePolicyResponse)(nil),        // 12: com.coralogix.quota.v1.DeletePolicyResponse
	(*ReorderPoliciesResponse)(nil),     // 13: com.coralogix.quota.v1.ReorderPoliciesResponse
	(*BulkTestLogPoliciesResponse)(nil), // 14: com.coralogix.quota.v1.BulkTestLogPoliciesResponse
	(*TogglePolicyResponse)(nil),        // 15: com.coralogix.quota.v1.TogglePolicyResponse
}
var file_com_coralogix_quota_v1_policies_service_proto_depIdxs = []int32{
	0,  // 0: com.coralogix.quota.v1.PoliciesService.GetPolicy:input_type -> com.coralogix.quota.v1.GetPolicyRequest
	1,  // 1: com.coralogix.quota.v1.PoliciesService.CreatePolicy:input_type -> com.coralogix.quota.v1.CreatePolicyRequest
	2,  // 2: com.coralogix.quota.v1.PoliciesService.UpdatePolicy:input_type -> com.coralogix.quota.v1.UpdatePolicyRequest
	3,  // 3: com.coralogix.quota.v1.PoliciesService.GetCompanyPolicies:input_type -> com.coralogix.quota.v1.GetCompanyPoliciesRequest
	4,  // 4: com.coralogix.quota.v1.PoliciesService.DeletePolicy:input_type -> com.coralogix.quota.v1.DeletePolicyRequest
	5,  // 5: com.coralogix.quota.v1.PoliciesService.ReorderPolicies:input_type -> com.coralogix.quota.v1.ReorderPoliciesRequest
	6,  // 6: com.coralogix.quota.v1.PoliciesService.BulkTestLogPolicies:input_type -> com.coralogix.quota.v1.BulkTestLogPoliciesRequest
	7,  // 7: com.coralogix.quota.v1.PoliciesService.TogglePolicy:input_type -> com.coralogix.quota.v1.TogglePolicyRequest
	8,  // 8: com.coralogix.quota.v1.PoliciesService.GetPolicy:output_type -> com.coralogix.quota.v1.GetPolicyResponse
	9,  // 9: com.coralogix.quota.v1.PoliciesService.CreatePolicy:output_type -> com.coralogix.quota.v1.CreatePolicyResponse
	10, // 10: com.coralogix.quota.v1.PoliciesService.UpdatePolicy:output_type -> com.coralogix.quota.v1.UpdatePolicyResponse
	11, // 11: com.coralogix.quota.v1.PoliciesService.GetCompanyPolicies:output_type -> com.coralogix.quota.v1.GetCompanyPoliciesResponse
	12, // 12: com.coralogix.quota.v1.PoliciesService.DeletePolicy:output_type -> com.coralogix.quota.v1.DeletePolicyResponse
	13, // 13: com.coralogix.quota.v1.PoliciesService.ReorderPolicies:output_type -> com.coralogix.quota.v1.ReorderPoliciesResponse
	14, // 14: com.coralogix.quota.v1.PoliciesService.BulkTestLogPolicies:output_type -> com.coralogix.quota.v1.BulkTestLogPoliciesResponse
	15, // 15: com.coralogix.quota.v1.PoliciesService.TogglePolicy:output_type -> com.coralogix.quota.v1.TogglePolicyResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogix_quota_v1_policies_service_proto_init() }
func file_com_coralogix_quota_v1_policies_service_proto_init() {
	if File_com_coralogix_quota_v1_policies_service_proto != nil {
		return
	}
	file_com_coralogix_quota_v1_audit_log_proto_init()
	file_com_coralogix_quota_v1_get_policy_request_proto_init()
	file_com_coralogix_quota_v1_get_policy_response_proto_init()
	file_com_coralogix_quota_v1_create_policy_request_proto_init()
	file_com_coralogix_quota_v1_create_policy_response_proto_init()
	file_com_coralogix_quota_v1_update_policy_request_proto_init()
	file_com_coralogix_quota_v1_update_policy_response_proto_init()
	file_com_coralogix_quota_v1_get_company_policies_request_proto_init()
	file_com_coralogix_quota_v1_get_company_policies_response_proto_init()
	file_com_coralogix_quota_v1_delete_policy_request_proto_init()
	file_com_coralogix_quota_v1_delete_policy_response_proto_init()
	file_com_coralogix_quota_v1_reorder_policies_request_proto_init()
	file_com_coralogix_quota_v1_reorder_policies_response_proto_init()
	file_com_coralogix_quota_v1_bulk_test_log_policies_request_proto_init()
	file_com_coralogix_quota_v1_bulk_test_log_policies_response_proto_init()
	file_com_coralogix_quota_v1_toggle_policy_request_proto_init()
	file_com_coralogix_quota_v1_toggle_policy_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_quota_v1_policies_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_quota_v1_policies_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_quota_v1_policies_service_proto_depIdxs,
	}.Build()
	File_com_coralogix_quota_v1_policies_service_proto = out.File
	file_com_coralogix_quota_v1_policies_service_proto_rawDesc = nil
	file_com_coralogix_quota_v1_policies_service_proto_goTypes = nil
	file_com_coralogix_quota_v1_policies_service_proto_depIdxs = nil
}
