// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: com/coralogix/rules/v1/rule_matcher.proto

package parsingrules

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SeverityConstraint_Value int32

const (
	SeverityConstraint_VALUE_DEBUG_OR_UNSPECIFIED SeverityConstraint_Value = 0
	SeverityConstraint_VALUE_VERBOSE              SeverityConstraint_Value = 1
	SeverityConstraint_VALUE_INFO                 SeverityConstraint_Value = 2
	SeverityConstraint_VALUE_WARNING              SeverityConstraint_Value = 3
	SeverityConstraint_VALUE_ERROR                SeverityConstraint_Value = 4
	SeverityConstraint_VALUE_CRITICAL             SeverityConstraint_Value = 5
)

// Enum value maps for SeverityConstraint_Value.
var (
	SeverityConstraint_Value_name = map[int32]string{
		0: "VALUE_DEBUG_OR_UNSPECIFIED",
		1: "VALUE_VERBOSE",
		2: "VALUE_INFO",
		3: "VALUE_WARNING",
		4: "VALUE_ERROR",
		5: "VALUE_CRITICAL",
	}
	SeverityConstraint_Value_value = map[string]int32{
		"VALUE_DEBUG_OR_UNSPECIFIED": 0,
		"VALUE_VERBOSE":              1,
		"VALUE_INFO":                 2,
		"VALUE_WARNING":              3,
		"VALUE_ERROR":                4,
		"VALUE_CRITICAL":             5,
	}
)

func (x SeverityConstraint_Value) Enum() *SeverityConstraint_Value {
	p := new(SeverityConstraint_Value)
	*p = x
	return p
}

func (x SeverityConstraint_Value) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SeverityConstraint_Value) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_rules_v1_rule_matcher_proto_enumTypes[0].Descriptor()
}

func (SeverityConstraint_Value) Type() protoreflect.EnumType {
	return &file_com_coralogix_rules_v1_rule_matcher_proto_enumTypes[0]
}

func (x SeverityConstraint_Value) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SeverityConstraint_Value.Descriptor instead.
func (SeverityConstraint_Value) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_rules_v1_rule_matcher_proto_rawDescGZIP(), []int{3, 0}
}

type RuleMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Constraint:
	//	*RuleMatcher_ApplicationName
	//	*RuleMatcher_SubsystemName
	//	*RuleMatcher_Severity
	Constraint isRuleMatcher_Constraint `protobuf_oneof:"constraint"`
}

func (x *RuleMatcher) Reset() {
	*x = RuleMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuleMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleMatcher) ProtoMessage() {}

func (x *RuleMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleMatcher.ProtoReflect.Descriptor instead.
func (*RuleMatcher) Descriptor() ([]byte, []int) {
	return file_com_coralogix_rules_v1_rule_matcher_proto_rawDescGZIP(), []int{0}
}

func (m *RuleMatcher) GetConstraint() isRuleMatcher_Constraint {
	if m != nil {
		return m.Constraint
	}
	return nil
}

func (x *RuleMatcher) GetApplicationName() *ApplicationNameConstraint {
	if x, ok := x.GetConstraint().(*RuleMatcher_ApplicationName); ok {
		return x.ApplicationName
	}
	return nil
}

func (x *RuleMatcher) GetSubsystemName() *SubsystemNameConstraint {
	if x, ok := x.GetConstraint().(*RuleMatcher_SubsystemName); ok {
		return x.SubsystemName
	}
	return nil
}

func (x *RuleMatcher) GetSeverity() *SeverityConstraint {
	if x, ok := x.GetConstraint().(*RuleMatcher_Severity); ok {
		return x.Severity
	}
	return nil
}

type isRuleMatcher_Constraint interface {
	isRuleMatcher_Constraint()
}

type RuleMatcher_ApplicationName struct {
	ApplicationName *ApplicationNameConstraint `protobuf:"bytes,1,opt,name=application_name,json=applicationName,proto3,oneof"`
}

type RuleMatcher_SubsystemName struct {
	SubsystemName *SubsystemNameConstraint `protobuf:"bytes,2,opt,name=subsystem_name,json=subsystemName,proto3,oneof"`
}

type RuleMatcher_Severity struct {
	Severity *SeverityConstraint `protobuf:"bytes,9,opt,name=severity,proto3,oneof"`
}

func (*RuleMatcher_ApplicationName) isRuleMatcher_Constraint() {}

func (*RuleMatcher_SubsystemName) isRuleMatcher_Constraint() {}

func (*RuleMatcher_Severity) isRuleMatcher_Constraint() {}

type ApplicationNameConstraint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ApplicationNameConstraint) Reset() {
	*x = ApplicationNameConstraint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationNameConstraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationNameConstraint) ProtoMessage() {}

func (x *ApplicationNameConstraint) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationNameConstraint.ProtoReflect.Descriptor instead.
func (*ApplicationNameConstraint) Descriptor() ([]byte, []int) {
	return file_com_coralogix_rules_v1_rule_matcher_proto_rawDescGZIP(), []int{1}
}

func (x *ApplicationNameConstraint) GetValue() *wrapperspb.StringValue {
	if x != nil {
		return x.Value
	}
	return nil
}

type SubsystemNameConstraint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SubsystemNameConstraint) Reset() {
	*x = SubsystemNameConstraint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubsystemNameConstraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubsystemNameConstraint) ProtoMessage() {}

func (x *SubsystemNameConstraint) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubsystemNameConstraint.ProtoReflect.Descriptor instead.
func (*SubsystemNameConstraint) Descriptor() ([]byte, []int) {
	return file_com_coralogix_rules_v1_rule_matcher_proto_rawDescGZIP(), []int{2}
}

func (x *SubsystemNameConstraint) GetValue() *wrapperspb.StringValue {
	if x != nil {
		return x.Value
	}
	return nil
}

type SeverityConstraint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value SeverityConstraint_Value `protobuf:"varint,1,opt,name=value,proto3,enum=com.coralogix.rules.v1.SeverityConstraint_Value" json:"value,omitempty"`
}

func (x *SeverityConstraint) Reset() {
	*x = SeverityConstraint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeverityConstraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeverityConstraint) ProtoMessage() {}

func (x *SeverityConstraint) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeverityConstraint.ProtoReflect.Descriptor instead.
func (*SeverityConstraint) Descriptor() ([]byte, []int) {
	return file_com_coralogix_rules_v1_rule_matcher_proto_rawDescGZIP(), []int{3}
}

func (x *SeverityConstraint) GetValue() SeverityConstraint_Value {
	if x != nil {
		return x.Value
	}
	return SeverityConstraint_VALUE_DEBUG_OR_UNSPECIFIED
}

var File_com_coralogix_rules_v1_rule_matcher_proto protoreflect.FileDescriptor

var file_com_coralogix_rules_v1_rule_matcher_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x02, 0x0a, 0x0b, 0x52, 0x75, 0x6c, 0x65, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x12, 0x5e, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x72, 0x75,
	0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x58, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x61,
	0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0d,
	0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a,
	0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74,
	0x79, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x08, 0x73,
	0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x42, 0x0c, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x74, 0x22, 0x4f, 0x0a, 0x19, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4d, 0x0a, 0x17, 0x53, 0x75, 0x62, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x74, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xe1, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x46, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e,
	0x0a, 0x1a, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x44, 0x45, 0x42, 0x55, 0x47, 0x5f, 0x4f, 0x52,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11,
	0x0a, 0x0d, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x56, 0x45, 0x52, 0x42, 0x4f, 0x53, 0x45, 0x10,
	0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10,
	0x02, 0x12, 0x11, 0x0a, 0x0d, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x49,
	0x4e, 0x47, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x43,
	0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x05, 0x42, 0x1c, 0x5a, 0x1a, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x69,
	0x6e, 0x67, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_rules_v1_rule_matcher_proto_rawDescOnce sync.Once
	file_com_coralogix_rules_v1_rule_matcher_proto_rawDescData = file_com_coralogix_rules_v1_rule_matcher_proto_rawDesc
)

func file_com_coralogix_rules_v1_rule_matcher_proto_rawDescGZIP() []byte {
	file_com_coralogix_rules_v1_rule_matcher_proto_rawDescOnce.Do(func() {
		file_com_coralogix_rules_v1_rule_matcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_rules_v1_rule_matcher_proto_rawDescData)
	})
	return file_com_coralogix_rules_v1_rule_matcher_proto_rawDescData
}

var file_com_coralogix_rules_v1_rule_matcher_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogix_rules_v1_rule_matcher_proto_goTypes = []interface{}{
	(SeverityConstraint_Value)(0),     // 0: com.coralogix.rules.v1.SeverityConstraint.Value
	(*RuleMatcher)(nil),               // 1: com.coralogix.rules.v1.RuleMatcher
	(*ApplicationNameConstraint)(nil), // 2: com.coralogix.rules.v1.ApplicationNameConstraint
	(*SubsystemNameConstraint)(nil),   // 3: com.coralogix.rules.v1.SubsystemNameConstraint
	(*SeverityConstraint)(nil),        // 4: com.coralogix.rules.v1.SeverityConstraint
	(*wrapperspb.StringValue)(nil),    // 5: google.protobuf.StringValue
}
var file_com_coralogix_rules_v1_rule_matcher_proto_depIdxs = []int32{
	2, // 0: com.coralogix.rules.v1.RuleMatcher.application_name:type_name -> com.coralogix.rules.v1.ApplicationNameConstraint
	3, // 1: com.coralogix.rules.v1.RuleMatcher.subsystem_name:type_name -> com.coralogix.rules.v1.SubsystemNameConstraint
	4, // 2: com.coralogix.rules.v1.RuleMatcher.severity:type_name -> com.coralogix.rules.v1.SeverityConstraint
	5, // 3: com.coralogix.rules.v1.ApplicationNameConstraint.value:type_name -> google.protobuf.StringValue
	5, // 4: com.coralogix.rules.v1.SubsystemNameConstraint.value:type_name -> google.protobuf.StringValue
	0, // 5: com.coralogix.rules.v1.SeverityConstraint.value:type_name -> com.coralogix.rules.v1.SeverityConstraint.Value
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_com_coralogix_rules_v1_rule_matcher_proto_init() }
func file_com_coralogix_rules_v1_rule_matcher_proto_init() {
	if File_com_coralogix_rules_v1_rule_matcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuleMatcher); i {
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
		file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationNameConstraint); i {
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
		file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubsystemNameConstraint); i {
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
		file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeverityConstraint); i {
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
	file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RuleMatcher_ApplicationName)(nil),
		(*RuleMatcher_SubsystemName)(nil),
		(*RuleMatcher_Severity)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_rules_v1_rule_matcher_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_rules_v1_rule_matcher_proto_goTypes,
		DependencyIndexes: file_com_coralogix_rules_v1_rule_matcher_proto_depIdxs,
		EnumInfos:         file_com_coralogix_rules_v1_rule_matcher_proto_enumTypes,
		MessageInfos:      file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes,
	}.Build()
	File_com_coralogix_rules_v1_rule_matcher_proto = out.File
	file_com_coralogix_rules_v1_rule_matcher_proto_rawDesc = nil
	file_com_coralogix_rules_v1_rule_matcher_proto_goTypes = nil
	file_com_coralogix_rules_v1_rule_matcher_proto_depIdxs = nil
}