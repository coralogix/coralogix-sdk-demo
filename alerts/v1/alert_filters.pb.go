// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: com/coralogix/alerts/v1/alert_filters.proto

package v1

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

type AlertFilters_LogSeverity int32

const (
	AlertFilters_LOG_SEVERITY_DEBUG_OR_UNSPECIFIED AlertFilters_LogSeverity = 0
	AlertFilters_LOG_SEVERITY_VERBOSE              AlertFilters_LogSeverity = 1
	AlertFilters_LOG_SEVERITY_INFO                 AlertFilters_LogSeverity = 2
	AlertFilters_LOG_SEVERITY_WARNING              AlertFilters_LogSeverity = 3
	AlertFilters_LOG_SEVERITY_ERROR                AlertFilters_LogSeverity = 4
	AlertFilters_LOG_SEVERITY_CRITICAL             AlertFilters_LogSeverity = 5
)

// Enum value maps for AlertFilters_LogSeverity.
var (
	AlertFilters_LogSeverity_name = map[int32]string{
		0: "LOG_SEVERITY_DEBUG_OR_UNSPECIFIED",
		1: "LOG_SEVERITY_VERBOSE",
		2: "LOG_SEVERITY_INFO",
		3: "LOG_SEVERITY_WARNING",
		4: "LOG_SEVERITY_ERROR",
		5: "LOG_SEVERITY_CRITICAL",
	}
	AlertFilters_LogSeverity_value = map[string]int32{
		"LOG_SEVERITY_DEBUG_OR_UNSPECIFIED": 0,
		"LOG_SEVERITY_VERBOSE":              1,
		"LOG_SEVERITY_INFO":                 2,
		"LOG_SEVERITY_WARNING":              3,
		"LOG_SEVERITY_ERROR":                4,
		"LOG_SEVERITY_CRITICAL":             5,
	}
)

func (x AlertFilters_LogSeverity) Enum() *AlertFilters_LogSeverity {
	p := new(AlertFilters_LogSeverity)
	*p = x
	return p
}

func (x AlertFilters_LogSeverity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlertFilters_LogSeverity) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_alerts_v1_alert_filters_proto_enumTypes[0].Descriptor()
}

func (AlertFilters_LogSeverity) Type() protoreflect.EnumType {
	return &file_com_coralogix_alerts_v1_alert_filters_proto_enumTypes[0]
}

func (x AlertFilters_LogSeverity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlertFilters_LogSeverity.Descriptor instead.
func (AlertFilters_LogSeverity) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_alerts_v1_alert_filters_proto_rawDescGZIP(), []int{0, 0}
}

type AlertFilters_FilterType int32

const (
	AlertFilters_FILTER_TYPE_TEXT_OR_UNSPECIFIED AlertFilters_FilterType = 0
	AlertFilters_FILTER_TYPE_TEMPLATE            AlertFilters_FilterType = 1
	AlertFilters_FILTER_TYPE_RATIO               AlertFilters_FilterType = 2
	AlertFilters_FILTER_TYPE_UNIQUE_COUNT        AlertFilters_FilterType = 3
	AlertFilters_FILTER_TYPE_TIME_RELATIVE       AlertFilters_FilterType = 4
	AlertFilters_FILTER_TYPE_METRIC              AlertFilters_FilterType = 5
	AlertFilters_FILTER_TYPE_TRACING             AlertFilters_FilterType = 6
	AlertFilters_FILTER_TYPE_FLOW                AlertFilters_FilterType = 7
)

// Enum value maps for AlertFilters_FilterType.
var (
	AlertFilters_FilterType_name = map[int32]string{
		0: "FILTER_TYPE_TEXT_OR_UNSPECIFIED",
		1: "FILTER_TYPE_TEMPLATE",
		2: "FILTER_TYPE_RATIO",
		3: "FILTER_TYPE_UNIQUE_COUNT",
		4: "FILTER_TYPE_TIME_RELATIVE",
		5: "FILTER_TYPE_METRIC",
		6: "FILTER_TYPE_TRACING",
		7: "FILTER_TYPE_FLOW",
	}
	AlertFilters_FilterType_value = map[string]int32{
		"FILTER_TYPE_TEXT_OR_UNSPECIFIED": 0,
		"FILTER_TYPE_TEMPLATE":            1,
		"FILTER_TYPE_RATIO":               2,
		"FILTER_TYPE_UNIQUE_COUNT":        3,
		"FILTER_TYPE_TIME_RELATIVE":       4,
		"FILTER_TYPE_METRIC":              5,
		"FILTER_TYPE_TRACING":             6,
		"FILTER_TYPE_FLOW":                7,
	}
)

func (x AlertFilters_FilterType) Enum() *AlertFilters_FilterType {
	p := new(AlertFilters_FilterType)
	*p = x
	return p
}

func (x AlertFilters_FilterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlertFilters_FilterType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_alerts_v1_alert_filters_proto_enumTypes[1].Descriptor()
}

func (AlertFilters_FilterType) Type() protoreflect.EnumType {
	return &file_com_coralogix_alerts_v1_alert_filters_proto_enumTypes[1]
}

func (x AlertFilters_FilterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlertFilters_FilterType.Descriptor instead.
func (AlertFilters_FilterType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_alerts_v1_alert_filters_proto_rawDescGZIP(), []int{0, 1}
}

type AlertFilters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Severities  []AlertFilters_LogSeverity    `protobuf:"varint,1,rep,packed,name=severities,proto3,enum=com.coralogix.alerts.v1.AlertFilters_LogSeverity" json:"severities,omitempty"`
	Metadata    *AlertFilters_MetadataFilters `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Alias       *wrapperspb.StringValue       `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`
	Text        *wrapperspb.StringValue       `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	RatioAlerts []*AlertFilters_RatioAlert    `protobuf:"bytes,5,rep,name=ratio_alerts,json=ratioAlerts,proto3" json:"ratio_alerts,omitempty"`
	FilterType  AlertFilters_FilterType       `protobuf:"varint,6,opt,name=filter_type,json=filterType,proto3,enum=com.coralogix.alerts.v1.AlertFilters_FilterType" json:"filter_type,omitempty"`
}

func (x *AlertFilters) Reset() {
	*x = AlertFilters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_alerts_v1_alert_filters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertFilters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertFilters) ProtoMessage() {}

func (x *AlertFilters) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_alerts_v1_alert_filters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertFilters.ProtoReflect.Descriptor instead.
func (*AlertFilters) Descriptor() ([]byte, []int) {
	return file_com_coralogix_alerts_v1_alert_filters_proto_rawDescGZIP(), []int{0}
}

func (x *AlertFilters) GetSeverities() []AlertFilters_LogSeverity {
	if x != nil {
		return x.Severities
	}
	return nil
}

func (x *AlertFilters) GetMetadata() *AlertFilters_MetadataFilters {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *AlertFilters) GetAlias() *wrapperspb.StringValue {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *AlertFilters) GetText() *wrapperspb.StringValue {
	if x != nil {
		return x.Text
	}
	return nil
}

func (x *AlertFilters) GetRatioAlerts() []*AlertFilters_RatioAlert {
	if x != nil {
		return x.RatioAlerts
	}
	return nil
}

func (x *AlertFilters) GetFilterType() AlertFilters_FilterType {
	if x != nil {
		return x.FilterType
	}
	return AlertFilters_FILTER_TYPE_TEXT_OR_UNSPECIFIED
}

type AlertFilters_MetadataFilters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Categories   []*wrapperspb.StringValue `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	Applications []*wrapperspb.StringValue `protobuf:"bytes,2,rep,name=applications,proto3" json:"applications,omitempty"`
	Subsystems   []*wrapperspb.StringValue `protobuf:"bytes,3,rep,name=subsystems,proto3" json:"subsystems,omitempty"`
	Computers    []*wrapperspb.StringValue `protobuf:"bytes,4,rep,name=computers,proto3" json:"computers,omitempty"`
	Classes      []*wrapperspb.StringValue `protobuf:"bytes,5,rep,name=classes,proto3" json:"classes,omitempty"`
	Methods      []*wrapperspb.StringValue `protobuf:"bytes,6,rep,name=methods,proto3" json:"methods,omitempty"`
	IpAddresses  []*wrapperspb.StringValue `protobuf:"bytes,7,rep,name=ip_addresses,json=ipAddresses,proto3" json:"ip_addresses,omitempty"`
}

func (x *AlertFilters_MetadataFilters) Reset() {
	*x = AlertFilters_MetadataFilters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_alerts_v1_alert_filters_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertFilters_MetadataFilters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertFilters_MetadataFilters) ProtoMessage() {}

func (x *AlertFilters_MetadataFilters) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_alerts_v1_alert_filters_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertFilters_MetadataFilters.ProtoReflect.Descriptor instead.
func (*AlertFilters_MetadataFilters) Descriptor() ([]byte, []int) {
	return file_com_coralogix_alerts_v1_alert_filters_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AlertFilters_MetadataFilters) GetCategories() []*wrapperspb.StringValue {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *AlertFilters_MetadataFilters) GetApplications() []*wrapperspb.StringValue {
	if x != nil {
		return x.Applications
	}
	return nil
}

func (x *AlertFilters_MetadataFilters) GetSubsystems() []*wrapperspb.StringValue {
	if x != nil {
		return x.Subsystems
	}
	return nil
}

func (x *AlertFilters_MetadataFilters) GetComputers() []*wrapperspb.StringValue {
	if x != nil {
		return x.Computers
	}
	return nil
}

func (x *AlertFilters_MetadataFilters) GetClasses() []*wrapperspb.StringValue {
	if x != nil {
		return x.Classes
	}
	return nil
}

func (x *AlertFilters_MetadataFilters) GetMethods() []*wrapperspb.StringValue {
	if x != nil {
		return x.Methods
	}
	return nil
}

func (x *AlertFilters_MetadataFilters) GetIpAddresses() []*wrapperspb.StringValue {
	if x != nil {
		return x.IpAddresses
	}
	return nil
}

type AlertFilters_RatioAlert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alias        *wrapperspb.StringValue    `protobuf:"bytes,1,opt,name=alias,proto3" json:"alias,omitempty"`
	Text         *wrapperspb.StringValue    `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Severities   []AlertFilters_LogSeverity `protobuf:"varint,3,rep,packed,name=severities,proto3,enum=com.coralogix.alerts.v1.AlertFilters_LogSeverity" json:"severities,omitempty"`
	Applications []*wrapperspb.StringValue  `protobuf:"bytes,4,rep,name=applications,proto3" json:"applications,omitempty"`
	Subsystems   []*wrapperspb.StringValue  `protobuf:"bytes,5,rep,name=subsystems,proto3" json:"subsystems,omitempty"`
	GroupBy      []*wrapperspb.StringValue  `protobuf:"bytes,6,rep,name=group_by,json=groupBy,proto3" json:"group_by,omitempty"`
}

func (x *AlertFilters_RatioAlert) Reset() {
	*x = AlertFilters_RatioAlert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_alerts_v1_alert_filters_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertFilters_RatioAlert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertFilters_RatioAlert) ProtoMessage() {}

func (x *AlertFilters_RatioAlert) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_alerts_v1_alert_filters_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertFilters_RatioAlert.ProtoReflect.Descriptor instead.
func (*AlertFilters_RatioAlert) Descriptor() ([]byte, []int) {
	return file_com_coralogix_alerts_v1_alert_filters_proto_rawDescGZIP(), []int{0, 1}
}

func (x *AlertFilters_RatioAlert) GetAlias() *wrapperspb.StringValue {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *AlertFilters_RatioAlert) GetText() *wrapperspb.StringValue {
	if x != nil {
		return x.Text
	}
	return nil
}

func (x *AlertFilters_RatioAlert) GetSeverities() []AlertFilters_LogSeverity {
	if x != nil {
		return x.Severities
	}
	return nil
}

func (x *AlertFilters_RatioAlert) GetApplications() []*wrapperspb.StringValue {
	if x != nil {
		return x.Applications
	}
	return nil
}

func (x *AlertFilters_RatioAlert) GetSubsystems() []*wrapperspb.StringValue {
	if x != nil {
		return x.Subsystems
	}
	return nil
}

func (x *AlertFilters_RatioAlert) GetGroupBy() []*wrapperspb.StringValue {
	if x != nil {
		return x.GroupBy
	}
	return nil
}

var File_com_coralogix_alerts_v1_alert_filters_proto protoreflect.FileDescriptor

var file_com_coralogix_alerts_v1_alert_filters_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x0d, 0x0a, 0x0c, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x51, 0x0a, 0x0a, 0x73, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x0a,
	0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x51, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a,
	0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61,
	0x73, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x53, 0x0a, 0x0c, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x5f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x0b, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x51, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x1a, 0xbc, 0x03, 0x0a, 0x0f,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x3c, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x40, 0x0a,
	0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x3c, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x3a, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x73, 0x12, 0x36, 0x0a, 0x07, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65,
	0x73, 0x12, 0x36, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x3f, 0x0a, 0x0c, 0x69, 0x70, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x69,
	0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x1a, 0xfe, 0x02, 0x0a, 0x0a, 0x52,
	0x61, 0x74, 0x69, 0x6f, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x30, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x51, 0x0a, 0x0a, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x40, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3c, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x37, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x62, 0x79, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x22, 0xb2, 0x01, 0x0a, 0x0b,
	0x4c, 0x6f, 0x67, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x21, 0x4c,
	0x4f, 0x47, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x44, 0x45, 0x42, 0x55,
	0x47, 0x5f, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49,
	0x54, 0x59, 0x5f, 0x56, 0x45, 0x52, 0x42, 0x4f, 0x53, 0x45, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11,
	0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x46,
	0x4f, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52,
	0x49, 0x54, 0x59, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x16, 0x0a,
	0x12, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x45, 0x56,
	0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x05,
	0x22, 0xe6, 0x01, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x23, 0x0a, 0x1f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54,
	0x45, 0x58, 0x54, 0x5f, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x54, 0x45, 0x4d, 0x50, 0x4c, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x15,
	0x0a, 0x11, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x49, 0x51, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x55, 0x4e,
	0x54, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x56, 0x45,
	0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49,
	0x4c, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x49, 0x4e,
	0x47, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x46, 0x4c, 0x4f, 0x57, 0x10, 0x07, 0x42, 0x1d, 0x5a, 0x1b, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_alerts_v1_alert_filters_proto_rawDescOnce sync.Once
	file_com_coralogix_alerts_v1_alert_filters_proto_rawDescData = file_com_coralogix_alerts_v1_alert_filters_proto_rawDesc
)

func file_com_coralogix_alerts_v1_alert_filters_proto_rawDescGZIP() []byte {
	file_com_coralogix_alerts_v1_alert_filters_proto_rawDescOnce.Do(func() {
		file_com_coralogix_alerts_v1_alert_filters_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_alerts_v1_alert_filters_proto_rawDescData)
	})
	return file_com_coralogix_alerts_v1_alert_filters_proto_rawDescData
}

var file_com_coralogix_alerts_v1_alert_filters_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_com_coralogix_alerts_v1_alert_filters_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogix_alerts_v1_alert_filters_proto_goTypes = []interface{}{
	(AlertFilters_LogSeverity)(0),        // 0: com.coralogix.alerts.v1.AlertFilters.LogSeverity
	(AlertFilters_FilterType)(0),         // 1: com.coralogix.alerts.v1.AlertFilters.FilterType
	(*AlertFilters)(nil),                 // 2: com.coralogix.alerts.v1.AlertFilters
	(*AlertFilters_MetadataFilters)(nil), // 3: com.coralogix.alerts.v1.AlertFilters.MetadataFilters
	(*AlertFilters_RatioAlert)(nil),      // 4: com.coralogix.alerts.v1.AlertFilters.RatioAlert
	(*wrapperspb.StringValue)(nil),       // 5: google.protobuf.StringValue
}
var file_com_coralogix_alerts_v1_alert_filters_proto_depIdxs = []int32{
	0,  // 0: com.coralogix.alerts.v1.AlertFilters.severities:type_name -> com.coralogix.alerts.v1.AlertFilters.LogSeverity
	3,  // 1: com.coralogix.alerts.v1.AlertFilters.metadata:type_name -> com.coralogix.alerts.v1.AlertFilters.MetadataFilters
	5,  // 2: com.coralogix.alerts.v1.AlertFilters.alias:type_name -> google.protobuf.StringValue
	5,  // 3: com.coralogix.alerts.v1.AlertFilters.text:type_name -> google.protobuf.StringValue
	4,  // 4: com.coralogix.alerts.v1.AlertFilters.ratio_alerts:type_name -> com.coralogix.alerts.v1.AlertFilters.RatioAlert
	1,  // 5: com.coralogix.alerts.v1.AlertFilters.filter_type:type_name -> com.coralogix.alerts.v1.AlertFilters.FilterType
	5,  // 6: com.coralogix.alerts.v1.AlertFilters.MetadataFilters.categories:type_name -> google.protobuf.StringValue
	5,  // 7: com.coralogix.alerts.v1.AlertFilters.MetadataFilters.applications:type_name -> google.protobuf.StringValue
	5,  // 8: com.coralogix.alerts.v1.AlertFilters.MetadataFilters.subsystems:type_name -> google.protobuf.StringValue
	5,  // 9: com.coralogix.alerts.v1.AlertFilters.MetadataFilters.computers:type_name -> google.protobuf.StringValue
	5,  // 10: com.coralogix.alerts.v1.AlertFilters.MetadataFilters.classes:type_name -> google.protobuf.StringValue
	5,  // 11: com.coralogix.alerts.v1.AlertFilters.MetadataFilters.methods:type_name -> google.protobuf.StringValue
	5,  // 12: com.coralogix.alerts.v1.AlertFilters.MetadataFilters.ip_addresses:type_name -> google.protobuf.StringValue
	5,  // 13: com.coralogix.alerts.v1.AlertFilters.RatioAlert.alias:type_name -> google.protobuf.StringValue
	5,  // 14: com.coralogix.alerts.v1.AlertFilters.RatioAlert.text:type_name -> google.protobuf.StringValue
	0,  // 15: com.coralogix.alerts.v1.AlertFilters.RatioAlert.severities:type_name -> com.coralogix.alerts.v1.AlertFilters.LogSeverity
	5,  // 16: com.coralogix.alerts.v1.AlertFilters.RatioAlert.applications:type_name -> google.protobuf.StringValue
	5,  // 17: com.coralogix.alerts.v1.AlertFilters.RatioAlert.subsystems:type_name -> google.protobuf.StringValue
	5,  // 18: com.coralogix.alerts.v1.AlertFilters.RatioAlert.group_by:type_name -> google.protobuf.StringValue
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_com_coralogix_alerts_v1_alert_filters_proto_init() }
func file_com_coralogix_alerts_v1_alert_filters_proto_init() {
	if File_com_coralogix_alerts_v1_alert_filters_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_alerts_v1_alert_filters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertFilters); i {
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
		file_com_coralogix_alerts_v1_alert_filters_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertFilters_MetadataFilters); i {
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
		file_com_coralogix_alerts_v1_alert_filters_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertFilters_RatioAlert); i {
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
			RawDescriptor: file_com_coralogix_alerts_v1_alert_filters_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_alerts_v1_alert_filters_proto_goTypes,
		DependencyIndexes: file_com_coralogix_alerts_v1_alert_filters_proto_depIdxs,
		EnumInfos:         file_com_coralogix_alerts_v1_alert_filters_proto_enumTypes,
		MessageInfos:      file_com_coralogix_alerts_v1_alert_filters_proto_msgTypes,
	}.Build()
	File_com_coralogix_alerts_v1_alert_filters_proto = out.File
	file_com_coralogix_alerts_v1_alert_filters_proto_rawDesc = nil
	file_com_coralogix_alerts_v1_alert_filters_proto_goTypes = nil
	file_com_coralogix_alerts_v1_alert_filters_proto_depIdxs = nil
}