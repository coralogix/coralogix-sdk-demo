// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: com/coralogix/alerts/v2/alert.proto

package v2

import (
	v1 "coralogix-sdk/alerts/v1"
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

type Alert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                         *wrapperspb.StringValue    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                       *wrapperspb.StringValue    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description                *wrapperspb.StringValue    `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	IsActive                   *wrapperspb.BoolValue      `protobuf:"bytes,4,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Severity                   v1.AlertSeverity           `protobuf:"varint,5,opt,name=severity,proto3,enum=com.coralogix.alerts.v1.AlertSeverity" json:"severity,omitempty"`
	Expiration                 *v1.Date                   `protobuf:"bytes,6,opt,name=expiration,proto3" json:"expiration,omitempty"`
	Condition                  *AlertCondition            `protobuf:"bytes,7,opt,name=condition,proto3" json:"condition,omitempty"`
	ShowInInsight              *ShowInInsight             `protobuf:"bytes,8,opt,name=show_in_insight,json=showInInsight,proto3,oneof" json:"show_in_insight,omitempty"`
	NotificationGroups         []*AlertNotificationGroups `protobuf:"bytes,9,rep,name=notification_groups,json=notificationGroups,proto3" json:"notification_groups,omitempty"`
	Filters                    *v1.AlertFilters           `protobuf:"bytes,10,opt,name=filters,proto3" json:"filters,omitempty"`
	ActiveWhen                 *v1.AlertActiveWhen        `protobuf:"bytes,11,opt,name=active_when,json=activeWhen,proto3" json:"active_when,omitempty"`
	NotificationPayloadFilters []*wrapperspb.StringValue  `protobuf:"bytes,12,rep,name=notification_payload_filters,json=notificationPayloadFilters,proto3" json:"notification_payload_filters,omitempty"`
	MetaLabels                 []*v1.MetaLabel            `protobuf:"bytes,13,rep,name=meta_labels,json=metaLabels,proto3" json:"meta_labels,omitempty"`
	MetaLabelsStrings          []*wrapperspb.StringValue  `protobuf:"bytes,14,rep,name=meta_labels_strings,json=metaLabelsStrings,proto3" json:"meta_labels_strings,omitempty"`
	TracingAlert               *v1.TracingAlert           `protobuf:"bytes,15,opt,name=tracing_alert,json=tracingAlert,proto3" json:"tracing_alert,omitempty"`
	UniqueIdentifier           *wrapperspb.StringValue    `protobuf:"bytes,16,opt,name=unique_identifier,json=uniqueIdentifier,proto3" json:"unique_identifier,omitempty"`
}

func (x *Alert) Reset() {
	*x = Alert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_alerts_v2_alert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alert) ProtoMessage() {}

func (x *Alert) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_alerts_v2_alert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alert.ProtoReflect.Descriptor instead.
func (*Alert) Descriptor() ([]byte, []int) {
	return file_com_coralogix_alerts_v2_alert_proto_rawDescGZIP(), []int{0}
}

func (x *Alert) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Alert) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Alert) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Alert) GetIsActive() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsActive
	}
	return nil
}

func (x *Alert) GetSeverity() v1.AlertSeverity {
	if x != nil {
		return x.Severity
	}
	return v1.AlertSeverity(0)
}

func (x *Alert) GetExpiration() *v1.Date {
	if x != nil {
		return x.Expiration
	}
	return nil
}

func (x *Alert) GetCondition() *AlertCondition {
	if x != nil {
		return x.Condition
	}
	return nil
}

func (x *Alert) GetShowInInsight() *ShowInInsight {
	if x != nil {
		return x.ShowInInsight
	}
	return nil
}

func (x *Alert) GetNotificationGroups() []*AlertNotificationGroups {
	if x != nil {
		return x.NotificationGroups
	}
	return nil
}

func (x *Alert) GetFilters() *v1.AlertFilters {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *Alert) GetActiveWhen() *v1.AlertActiveWhen {
	if x != nil {
		return x.ActiveWhen
	}
	return nil
}

func (x *Alert) GetNotificationPayloadFilters() []*wrapperspb.StringValue {
	if x != nil {
		return x.NotificationPayloadFilters
	}
	return nil
}

func (x *Alert) GetMetaLabels() []*v1.MetaLabel {
	if x != nil {
		return x.MetaLabels
	}
	return nil
}

func (x *Alert) GetMetaLabelsStrings() []*wrapperspb.StringValue {
	if x != nil {
		return x.MetaLabelsStrings
	}
	return nil
}

func (x *Alert) GetTracingAlert() *v1.TracingAlert {
	if x != nil {
		return x.TracingAlert
	}
	return nil
}

func (x *Alert) GetUniqueIdentifier() *wrapperspb.StringValue {
	if x != nil {
		return x.UniqueIdentifier
	}
	return nil
}

var File_com_coralogix_alerts_v2_alert_proto protoreflect.FileDescriptor

var file_com_coralogix_alerts_v2_alert_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x77, 0x68, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x09, 0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x37, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x42, 0x0a, 0x08, 0x73, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x65, 0x76, 0x65,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x3d,
	0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74,
	0x65, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a,
	0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x0f, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x69, 0x6e, 0x5f,
	0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x49, 0x6e, 0x49, 0x6e,
	0x73, 0x69, 0x67, 0x68, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x68, 0x6f, 0x77, 0x49, 0x6e, 0x49,
	0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01, 0x01, 0x12, 0x61, 0x0a, 0x13, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x12, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x3f, 0x0a, 0x07,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x49, 0x0a,
	0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x77, 0x68, 0x65, 0x6e, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x57, 0x68, 0x65, 0x6e, 0x52, 0x0a, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x57, 0x68, 0x65, 0x6e, 0x12, 0x5e, 0x0a, 0x1c, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x1a, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x43, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x61,
	0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x4c, 0x0a,
	0x13, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x5f, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x6d, 0x65, 0x74, 0x61, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x4a, 0x0a, 0x0d, 0x74,
	0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x49, 0x0a, 0x11, 0x75, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x10, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x69, 0x6e, 0x5f, 0x69,
	0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x42, 0x1d, 0x5a, 0x1b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x73, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_alerts_v2_alert_proto_rawDescOnce sync.Once
	file_com_coralogix_alerts_v2_alert_proto_rawDescData = file_com_coralogix_alerts_v2_alert_proto_rawDesc
)

func file_com_coralogix_alerts_v2_alert_proto_rawDescGZIP() []byte {
	file_com_coralogix_alerts_v2_alert_proto_rawDescOnce.Do(func() {
		file_com_coralogix_alerts_v2_alert_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_alerts_v2_alert_proto_rawDescData)
	})
	return file_com_coralogix_alerts_v2_alert_proto_rawDescData
}

var file_com_coralogix_alerts_v2_alert_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_alerts_v2_alert_proto_goTypes = []interface{}{
	(*Alert)(nil),                   // 0: com.coralogix.alerts.v2.Alert
	(*wrapperspb.StringValue)(nil),  // 1: google.protobuf.StringValue
	(*wrapperspb.BoolValue)(nil),    // 2: google.protobuf.BoolValue
	(v1.AlertSeverity)(0),           // 3: com.coralogix.alerts.v1.AlertSeverity
	(*v1.Date)(nil),                 // 4: com.coralogix.alerts.v1.Date
	(*AlertCondition)(nil),          // 5: com.coralogix.alerts.v2.AlertCondition
	(*ShowInInsight)(nil),           // 6: com.coralogix.alerts.v2.ShowInInsight
	(*AlertNotificationGroups)(nil), // 7: com.coralogix.alerts.v2.AlertNotificationGroups
	(*v1.AlertFilters)(nil),         // 8: com.coralogix.alerts.v1.AlertFilters
	(*v1.AlertActiveWhen)(nil),      // 9: com.coralogix.alerts.v1.AlertActiveWhen
	(*v1.MetaLabel)(nil),            // 10: com.coralogix.alerts.v1.MetaLabel
	(*v1.TracingAlert)(nil),         // 11: com.coralogix.alerts.v1.TracingAlert
}
var file_com_coralogix_alerts_v2_alert_proto_depIdxs = []int32{
	1,  // 0: com.coralogix.alerts.v2.Alert.id:type_name -> google.protobuf.StringValue
	1,  // 1: com.coralogix.alerts.v2.Alert.name:type_name -> google.protobuf.StringValue
	1,  // 2: com.coralogix.alerts.v2.Alert.description:type_name -> google.protobuf.StringValue
	2,  // 3: com.coralogix.alerts.v2.Alert.is_active:type_name -> google.protobuf.BoolValue
	3,  // 4: com.coralogix.alerts.v2.Alert.severity:type_name -> com.coralogix.alerts.v1.AlertSeverity
	4,  // 5: com.coralogix.alerts.v2.Alert.expiration:type_name -> com.coralogix.alerts.v1.Date
	5,  // 6: com.coralogix.alerts.v2.Alert.condition:type_name -> com.coralogix.alerts.v2.AlertCondition
	6,  // 7: com.coralogix.alerts.v2.Alert.show_in_insight:type_name -> com.coralogix.alerts.v2.ShowInInsight
	7,  // 8: com.coralogix.alerts.v2.Alert.notification_groups:type_name -> com.coralogix.alerts.v2.AlertNotificationGroups
	8,  // 9: com.coralogix.alerts.v2.Alert.filters:type_name -> com.coralogix.alerts.v1.AlertFilters
	9,  // 10: com.coralogix.alerts.v2.Alert.active_when:type_name -> com.coralogix.alerts.v1.AlertActiveWhen
	1,  // 11: com.coralogix.alerts.v2.Alert.notification_payload_filters:type_name -> google.protobuf.StringValue
	10, // 12: com.coralogix.alerts.v2.Alert.meta_labels:type_name -> com.coralogix.alerts.v1.MetaLabel
	1,  // 13: com.coralogix.alerts.v2.Alert.meta_labels_strings:type_name -> google.protobuf.StringValue
	11, // 14: com.coralogix.alerts.v2.Alert.tracing_alert:type_name -> com.coralogix.alerts.v1.TracingAlert
	1,  // 15: com.coralogix.alerts.v2.Alert.unique_identifier:type_name -> google.protobuf.StringValue
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_com_coralogix_alerts_v2_alert_proto_init() }
func file_com_coralogix_alerts_v2_alert_proto_init() {
	if File_com_coralogix_alerts_v2_alert_proto != nil {
		return
	}
	file_com_coralogix_alerts_v2_alert_condition_proto_init()
	file_com_coralogix_alerts_v2_alert_notification_group_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_alerts_v2_alert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alert); i {
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
	file_com_coralogix_alerts_v2_alert_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_alerts_v2_alert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_alerts_v2_alert_proto_goTypes,
		DependencyIndexes: file_com_coralogix_alerts_v2_alert_proto_depIdxs,
		MessageInfos:      file_com_coralogix_alerts_v2_alert_proto_msgTypes,
	}.Build()
	File_com_coralogix_alerts_v2_alert_proto = out.File
	file_com_coralogix_alerts_v2_alert_proto_rawDesc = nil
	file_com_coralogix_alerts_v2_alert_proto_goTypes = nil
	file_com_coralogix_alerts_v2_alert_proto_depIdxs = nil
}
