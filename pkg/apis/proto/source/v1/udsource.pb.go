// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: pkg/apis/proto/source/v1/udsource.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// event_time is the time associated with each datum.
	EventTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
}

func (x *EventTime) Reset() {
	*x = EventTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTime) ProtoMessage() {}

func (x *EventTime) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTime.ProtoReflect.Descriptor instead.
func (*EventTime) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_source_v1_udsource_proto_rawDescGZIP(), []int{0}
}

func (x *EventTime) GetEventTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTime
	}
	return nil
}

// *
// DatumResponse represents a datum response element.
type DatumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// payload is the payload of the datum.
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// offset is the offset information of the datum.
	Offset *Offset `protobuf:"bytes,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// keys is an optional list of keys associated with the datum.
	// key is the "key" attribute in (key,value) as in the map-reduce paradigm.
	// we add this optional field to support the use case where the user defined source can provide keys for the datum.
	// e.g. Kafka and Redis Stream message usually include information about the keys.
	Keys []string `protobuf:"bytes,3,rep,name=keys,proto3" json:"keys,omitempty"`
	// event_time is the time associated with each datum.
	// we add this optional field to support the use case where the user defined source can provide event time for the datum.
	// e.g. Kafka and Redis Stream message usually include information about the event time.
	EventTime *EventTime `protobuf:"bytes,4,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
}

func (x *DatumResponse) Reset() {
	*x = DatumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatumResponse) ProtoMessage() {}

func (x *DatumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatumResponse.ProtoReflect.Descriptor instead.
func (*DatumResponse) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_source_v1_udsource_proto_rawDescGZIP(), []int{1}
}

func (x *DatumResponse) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *DatumResponse) GetOffset() *Offset {
	if x != nil {
		return x.Offset
	}
	return nil
}

func (x *DatumResponse) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *DatumResponse) GetEventTime() *EventTime {
	if x != nil {
		return x.EventTime
	}
	return nil
}

// *
// DatumResponseList represents a list of datum response elements.
type DatumResponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Elements []*DatumResponse `protobuf:"bytes,1,rep,name=elements,proto3" json:"elements,omitempty"`
}

func (x *DatumResponseList) Reset() {
	*x = DatumResponseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatumResponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatumResponseList) ProtoMessage() {}

func (x *DatumResponseList) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatumResponseList.ProtoReflect.Descriptor instead.
func (*DatumResponseList) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_source_v1_udsource_proto_rawDescGZIP(), []int{2}
}

func (x *DatumResponseList) GetElements() []*DatumResponse {
	if x != nil {
		return x.Elements
	}
	return nil
}

// AckRequest is the request for acknowledging datum.
// it takes a list of offsets to be acknowledged.
type AckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offsets []*Offset `protobuf:"bytes,1,rep,name=offsets,proto3" json:"offsets,omitempty"`
}

func (x *AckRequest) Reset() {
	*x = AckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckRequest) ProtoMessage() {}

func (x *AckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckRequest.ProtoReflect.Descriptor instead.
func (*AckRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_source_v1_udsource_proto_rawDescGZIP(), []int{3}
}

func (x *AckRequest) GetOffsets() []*Offset {
	if x != nil {
		return x.Offsets
	}
	return nil
}

// AckResponseList represents a list of ack response elements.
type AckResponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offsets []*Offset `protobuf:"bytes,1,rep,name=offsets,proto3" json:"offsets,omitempty"`
}

func (x *AckResponseList) Reset() {
	*x = AckResponseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckResponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckResponseList) ProtoMessage() {}

func (x *AckResponseList) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckResponseList.ProtoReflect.Descriptor instead.
func (*AckResponseList) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_source_v1_udsource_proto_rawDescGZIP(), []int{4}
}

func (x *AckResponseList) GetOffsets() []*Offset {
	if x != nil {
		return x.Offsets
	}
	return nil
}

// *
// ReadyResponse is the health check result.
type ReadyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ready bool `protobuf:"varint,1,opt,name=ready,proto3" json:"ready,omitempty"`
}

func (x *ReadyResponse) Reset() {
	*x = ReadyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyResponse) ProtoMessage() {}

func (x *ReadyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyResponse.ProtoReflect.Descriptor instead.
func (*ReadyResponse) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_source_v1_udsource_proto_rawDescGZIP(), []int{5}
}

func (x *ReadyResponse) GetReady() bool {
	if x != nil {
		return x.Ready
	}
	return false
}

// *
// ReadRequest is the request for reading datum stream from user defined source.
type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// num_records is the number of records to read.
	NumRecords uint64 `protobuf:"varint,1,opt,name=num_records,json=numRecords,proto3" json:"num_records,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_source_v1_udsource_proto_rawDescGZIP(), []int{6}
}

func (x *ReadRequest) GetNumRecords() uint64 {
	if x != nil {
		return x.NumRecords
	}
	return 0
}

// *
// Offset is the offset of the datum.
type Offset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// offset is the offset of the datum.
	// we define Offset as a byte array because different input data sources can have different representations for Offset.
	// the only way to generalize it is to define it as a byte array,
	// such that we can let the UDSource to de-serialize the offset using its own interpretation logics.
	Offset []byte `protobuf:"bytes,1,opt,name=offset,proto3" json:"offset,omitempty"`
	// is_acked indicates whether the offset is successfully acknowledged by the user defined source.
	IsAcked bool `protobuf:"varint,2,opt,name=is_acked,json=isAcked,proto3" json:"is_acked,omitempty"`
}

func (x *Offset) Reset() {
	*x = Offset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Offset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Offset) ProtoMessage() {}

func (x *Offset) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Offset.ProtoReflect.Descriptor instead.
func (*Offset) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_source_v1_udsource_proto_rawDescGZIP(), []int{7}
}

func (x *Offset) GetOffset() []byte {
	if x != nil {
		return x.Offset
	}
	return nil
}

func (x *Offset) GetIsAcked() bool {
	if x != nil {
		return x.IsAcked
	}
	return false
}

// *
// PendingResponse is the response for the pending request.
type PendingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// count is the number of pending records at the user defined source.
	Count uint64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *PendingResponse) Reset() {
	*x = PendingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PendingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PendingResponse) ProtoMessage() {}

func (x *PendingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PendingResponse.ProtoReflect.Descriptor instead.
func (*PendingResponse) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_source_v1_udsource_proto_rawDescGZIP(), []int{8}
}

func (x *PendingResponse) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_pkg_apis_proto_source_v1_udsource_proto protoreflect.FileDescriptor

var file_pkg_apis_proto_source_v1_udsource_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x64, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x46, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x9d, 0x01, 0x0a, 0x0d, 0x44,
	0x61, 0x74, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x29, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x33, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x11, 0x44, 0x61,
	0x74, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x34, 0x0a, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61,
	0x74, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x39, 0x0a, 0x0a, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x52, 0x07, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x73,
	0x22, 0x3e, 0x0a, 0x0f, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x52, 0x07, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x73,
	0x22, 0x25, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x22, 0x2e, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6e, 0x75, 0x6d,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x3b, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f,
	0x61, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41,
	0x63, 0x6b, 0x65, 0x64, 0x22, 0x27, 0x0a, 0x0f, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x89, 0x02,
	0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x16, 0x2e, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x61, 0x74, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x30, 0x01, 0x12, 0x38, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x15, 0x2e, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3d, 0x0a,
	0x07, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1a, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x07,
	0x49, 0x73, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x18, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x75, 0x6d, 0x61, 0x70, 0x72, 0x6f, 0x6a,
	0x2f, 0x6e, 0x75, 0x6d, 0x61, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_apis_proto_source_v1_udsource_proto_rawDescOnce sync.Once
	file_pkg_apis_proto_source_v1_udsource_proto_rawDescData = file_pkg_apis_proto_source_v1_udsource_proto_rawDesc
)

func file_pkg_apis_proto_source_v1_udsource_proto_rawDescGZIP() []byte {
	file_pkg_apis_proto_source_v1_udsource_proto_rawDescOnce.Do(func() {
		file_pkg_apis_proto_source_v1_udsource_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_apis_proto_source_v1_udsource_proto_rawDescData)
	})
	return file_pkg_apis_proto_source_v1_udsource_proto_rawDescData
}

var file_pkg_apis_proto_source_v1_udsource_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_apis_proto_source_v1_udsource_proto_goTypes = []interface{}{
	(*EventTime)(nil),             // 0: source.v1.EventTime
	(*DatumResponse)(nil),         // 1: source.v1.DatumResponse
	(*DatumResponseList)(nil),     // 2: source.v1.DatumResponseList
	(*AckRequest)(nil),            // 3: source.v1.AckRequest
	(*AckResponseList)(nil),       // 4: source.v1.AckResponseList
	(*ReadyResponse)(nil),         // 5: source.v1.ReadyResponse
	(*ReadRequest)(nil),           // 6: source.v1.ReadRequest
	(*Offset)(nil),                // 7: source.v1.Offset
	(*PendingResponse)(nil),       // 8: source.v1.PendingResponse
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 10: google.protobuf.Empty
}
var file_pkg_apis_proto_source_v1_udsource_proto_depIdxs = []int32{
	9,  // 0: source.v1.EventTime.event_time:type_name -> google.protobuf.Timestamp
	7,  // 1: source.v1.DatumResponse.offset:type_name -> source.v1.Offset
	0,  // 2: source.v1.DatumResponse.event_time:type_name -> source.v1.EventTime
	1,  // 3: source.v1.DatumResponseList.elements:type_name -> source.v1.DatumResponse
	7,  // 4: source.v1.AckRequest.offsets:type_name -> source.v1.Offset
	7,  // 5: source.v1.AckResponseList.offsets:type_name -> source.v1.Offset
	6,  // 6: source.v1.UserDefinedSource.Read:input_type -> source.v1.ReadRequest
	3,  // 7: source.v1.UserDefinedSource.Ack:input_type -> source.v1.AckRequest
	10, // 8: source.v1.UserDefinedSource.Pending:input_type -> google.protobuf.Empty
	10, // 9: source.v1.UserDefinedSource.IsReady:input_type -> google.protobuf.Empty
	2,  // 10: source.v1.UserDefinedSource.Read:output_type -> source.v1.DatumResponseList
	4,  // 11: source.v1.UserDefinedSource.Ack:output_type -> source.v1.AckResponseList
	8,  // 12: source.v1.UserDefinedSource.Pending:output_type -> source.v1.PendingResponse
	5,  // 13: source.v1.UserDefinedSource.IsReady:output_type -> source.v1.ReadyResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_apis_proto_source_v1_udsource_proto_init() }
func file_pkg_apis_proto_source_v1_udsource_proto_init() {
	if File_pkg_apis_proto_source_v1_udsource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventTime); i {
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
		file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatumResponse); i {
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
		file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatumResponseList); i {
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
		file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckRequest); i {
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
		file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckResponseList); i {
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
		file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadyResponse); i {
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
		file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Offset); i {
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
		file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PendingResponse); i {
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
			RawDescriptor: file_pkg_apis_proto_source_v1_udsource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_apis_proto_source_v1_udsource_proto_goTypes,
		DependencyIndexes: file_pkg_apis_proto_source_v1_udsource_proto_depIdxs,
		MessageInfos:      file_pkg_apis_proto_source_v1_udsource_proto_msgTypes,
	}.Build()
	File_pkg_apis_proto_source_v1_udsource_proto = out.File
	file_pkg_apis_proto_source_v1_udsource_proto_rawDesc = nil
	file_pkg_apis_proto_source_v1_udsource_proto_goTypes = nil
	file_pkg_apis_proto_source_v1_udsource_proto_depIdxs = nil
}