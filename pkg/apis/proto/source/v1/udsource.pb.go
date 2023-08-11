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

// ReadRequest is the request for reading datum stream from user defined source.
type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required field indicating the request.
	Request *ReadRequest_Request `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_source_v1_udsource_proto_rawDescGZIP(), []int{0}
}

func (x *ReadRequest) GetRequest() *ReadRequest_Request {
	if x != nil {
		return x.Request
	}
	return nil
}

// DatumResponse represents a datum response element.
type DatumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required field holding the result.
	Result *DatumResponse_Result `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
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

func (x *DatumResponse) GetResult() *DatumResponse_Result {
	if x != nil {
		return x.Result
	}
	return nil
}

// AckRequest is the request for acknowledging datum.
// It takes a list of offsets to be acknowledged.
type AckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required field holding the request. The list will be ordered and will have the same order as the original Read response.
	Request *AckRequest_Request `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *AckRequest) Reset() {
	*x = AckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckRequest) ProtoMessage() {}

func (x *AckRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AckRequest.ProtoReflect.Descriptor instead.
func (*AckRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_source_v1_udsource_proto_rawDescGZIP(), []int{2}
}

func (x *AckRequest) GetRequest() *AckRequest_Request {
	if x != nil {
		return x.Request
	}
	return nil
}

// AckResponse is the response for acknowledging datum. It contains one empty field confirming
// the batch of offsets that have been successfully acknowledged. The contract between client and server
// is that the server will only return the AckResponse if the ack request is successful.
// If the server hangs during the ack request, the client can decide to timeout and error out the data forwarder.
// The reason why we define such contract is that we always expect the server to be able to process the ack request.
// Client is expected to send the AckRequest to the server with offsets that are strictly
// corresponding to the previously read batch. If the client sends the AckRequest with offsets that are not,
// it is considered as a client error and the server will not return the AckResponse.
type AckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required field holding the result.
	Result *AckResponse_Result `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AckResponse) Reset() {
	*x = AckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckResponse) ProtoMessage() {}

func (x *AckResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AckResponse.ProtoReflect.Descriptor instead.
func (*AckResponse) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_source_v1_udsource_proto_rawDescGZIP(), []int{3}
}

func (x *AckResponse) GetResult() *AckResponse_Result {
	if x != nil {
		return x.Result
	}
	return nil
}

// ReadyResponse is the health check result for user defined source.
type ReadyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required field holding the health check result.
	Ready bool `protobuf:"varint,1,opt,name=ready,proto3" json:"ready,omitempty"`
}

func (x *ReadyResponse) Reset() {
	*x = ReadyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyResponse) ProtoMessage() {}

func (x *ReadyResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ReadyResponse.ProtoReflect.Descriptor instead.
func (*ReadyResponse) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_source_v1_udsource_proto_rawDescGZIP(), []int{4}
}

func (x *ReadyResponse) GetReady() bool {
	if x != nil {
		return x.Ready
	}
	return false
}

// PendingResponse is the response for the pending request.
type PendingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required field holding the result.
	Result *PendingResponse_Result `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PendingResponse) Reset() {
	*x = PendingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PendingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PendingResponse) ProtoMessage() {}

func (x *PendingResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PendingResponse.ProtoReflect.Descriptor instead.
func (*PendingResponse) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_source_v1_udsource_proto_rawDescGZIP(), []int{5}
}

func (x *PendingResponse) GetResult() *PendingResponse_Result {
	if x != nil {
		return x.Result
	}
	return nil
}

// Offset is the offset of the datum.
type Offset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// offset is the offset of the datum. This field is required.
	// We define Offset as a byte array because different input data sources can have different representations for Offset.
	// The only way to generalize it is to define it as a byte array,
	// Such that we can let the UDSource to de-serialize the offset using its own interpretation logics.
	Offset []byte `protobuf:"bytes,1,opt,name=offset,proto3" json:"offset,omitempty"`
	// Optional partition_id indicates which partition of the source the datum belongs to.
	// It is useful for sources that have multiple partitions. e.g. Kafka.
	// If the partition_id is not specified, it is assumed that the source has a single partition.
	PartitionId string `protobuf:"bytes,2,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
}

func (x *Offset) Reset() {
	*x = Offset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Offset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Offset) ProtoMessage() {}

func (x *Offset) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Offset.ProtoReflect.Descriptor instead.
func (*Offset) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_source_v1_udsource_proto_rawDescGZIP(), []int{6}
}

func (x *Offset) GetOffset() []byte {
	if x != nil {
		return x.Offset
	}
	return nil
}

func (x *Offset) GetPartitionId() string {
	if x != nil {
		return x.PartitionId
	}
	return ""
}

type ReadRequest_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required field indicating the number of records to read.
	NumRecords uint64 `protobuf:"varint,1,opt,name=num_records,json=numRecords,proto3" json:"num_records,omitempty"`
	// Required field indicating the request timeout in milliseconds.
	TimeoutInMs uint64 `protobuf:"varint,2,opt,name=timeout_in_ms,json=timeoutInMs,proto3" json:"timeout_in_ms,omitempty"`
}

func (x *ReadRequest_Request) Reset() {
	*x = ReadRequest_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest_Request) ProtoMessage() {}

func (x *ReadRequest_Request) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ReadRequest_Request.ProtoReflect.Descriptor instead.
func (*ReadRequest_Request) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_source_v1_udsource_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ReadRequest_Request) GetNumRecords() uint64 {
	if x != nil {
		return x.NumRecords
	}
	return 0
}

func (x *ReadRequest_Request) GetTimeoutInMs() uint64 {
	if x != nil {
		return x.TimeoutInMs
	}
	return 0
}

type DatumResponse_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required field holding the payload of the datum.
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// Required field indicating the offset information of the datum.
	Offset *Offset `protobuf:"bytes,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Required field representing the time associated with each datum. It is used for watermarking.
	EventTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// Optional list of keys associated with the datum.
	// Key is the "key" attribute in (key,value) as in the map-reduce paradigm.
	// We add this optional field to support the use case where the user defined source can provide keys for the datum.
	// e.g. Kafka and Redis Stream message usually include information about the keys.
	Keys []string `protobuf:"bytes,4,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *DatumResponse_Result) Reset() {
	*x = DatumResponse_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatumResponse_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatumResponse_Result) ProtoMessage() {}

func (x *DatumResponse_Result) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DatumResponse_Result.ProtoReflect.Descriptor instead.
func (*DatumResponse_Result) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_source_v1_udsource_proto_rawDescGZIP(), []int{1, 0}
}

func (x *DatumResponse_Result) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *DatumResponse_Result) GetOffset() *Offset {
	if x != nil {
		return x.Offset
	}
	return nil
}

func (x *DatumResponse_Result) GetEventTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTime
	}
	return nil
}

func (x *DatumResponse_Result) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type AckRequest_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required field holding a list of offsets to be acknowledged.
	// The offsets must be strictly corresponding to the previously read batch,
	// meaning the offsets must be in the same order as the datum responses in the ReadResponse.
	// By enforcing ordering, we can save deserialization effort on the server side, assuming the server keeps a local copy of the raw/un-serialized offsets.
	Offset []*Offset `protobuf:"bytes,1,rep,name=offset,proto3" json:"offset,omitempty"`
}

func (x *AckRequest_Request) Reset() {
	*x = AckRequest_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckRequest_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckRequest_Request) ProtoMessage() {}

func (x *AckRequest_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckRequest_Request.ProtoReflect.Descriptor instead.
func (*AckRequest_Request) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_source_v1_udsource_proto_rawDescGZIP(), []int{2, 0}
}

func (x *AckRequest_Request) GetOffset() []*Offset {
	if x != nil {
		return x.Offset
	}
	return nil
}

type AckResponse_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required field indicating the ack request is successful.
	Success *emptypb.Empty `protobuf:"bytes,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AckResponse_Result) Reset() {
	*x = AckResponse_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckResponse_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckResponse_Result) ProtoMessage() {}

func (x *AckResponse_Result) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckResponse_Result.ProtoReflect.Descriptor instead.
func (*AckResponse_Result) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_source_v1_udsource_proto_rawDescGZIP(), []int{3, 0}
}

func (x *AckResponse_Result) GetSuccess() *emptypb.Empty {
	if x != nil {
		return x.Success
	}
	return nil
}

type PendingResponse_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required field holding the number of pending records at the user defined source.
	Count uint64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *PendingResponse_Result) Reset() {
	*x = PendingResponse_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PendingResponse_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PendingResponse_Result) ProtoMessage() {}

func (x *PendingResponse_Result) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PendingResponse_Result.ProtoReflect.Descriptor instead.
func (*PendingResponse_Result) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_source_v1_udsource_proto_rawDescGZIP(), []int{5, 0}
}

func (x *PendingResponse_Result) GetCount() uint64 {
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
	0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x38, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4e, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6e, 0x75,
	0x6d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x49, 0x6e, 0x4d, 0x73, 0x22, 0xe7, 0x01, 0x0a,
	0x0d, 0x44, 0x61, 0x74, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x75, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x9c, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x29, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x7b, 0x0a, 0x0a, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x22, 0x80, 0x01, 0x0a, 0x0b, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x3a, 0x0a, 0x06, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x30, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x25, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x22, 0x6c, 0x0a,
	0x0f, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x39, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x1e, 0x0a, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x43, 0x0a, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x32, 0xfc, 0x01, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x52,
	0x65, 0x61, 0x64, 0x46, 0x6e, 0x12, 0x16, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x75, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x36, 0x0a, 0x05, 0x41, 0x63, 0x6b,
	0x46, 0x6e, 0x12, 0x15, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3f, 0x0a, 0x09, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x46, 0x6e, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x49, 0x73, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x75,
	0x6d, 0x61, 0x70, 0x72, 0x6f, 0x6a, 0x2f, 0x6e, 0x75, 0x6d, 0x61, 0x66, 0x6c, 0x6f, 0x77, 0x2d,
	0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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

var file_pkg_apis_proto_source_v1_udsource_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_pkg_apis_proto_source_v1_udsource_proto_goTypes = []interface{}{
	(*ReadRequest)(nil),            // 0: source.v1.ReadRequest
	(*DatumResponse)(nil),          // 1: source.v1.DatumResponse
	(*AckRequest)(nil),             // 2: source.v1.AckRequest
	(*AckResponse)(nil),            // 3: source.v1.AckResponse
	(*ReadyResponse)(nil),          // 4: source.v1.ReadyResponse
	(*PendingResponse)(nil),        // 5: source.v1.PendingResponse
	(*Offset)(nil),                 // 6: source.v1.Offset
	(*ReadRequest_Request)(nil),    // 7: source.v1.ReadRequest.Request
	(*DatumResponse_Result)(nil),   // 8: source.v1.DatumResponse.Result
	(*AckRequest_Request)(nil),     // 9: source.v1.AckRequest.Request
	(*AckResponse_Result)(nil),     // 10: source.v1.AckResponse.Result
	(*PendingResponse_Result)(nil), // 11: source.v1.PendingResponse.Result
	(*timestamppb.Timestamp)(nil),  // 12: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),          // 13: google.protobuf.Empty
}
var file_pkg_apis_proto_source_v1_udsource_proto_depIdxs = []int32{
	7,  // 0: source.v1.ReadRequest.request:type_name -> source.v1.ReadRequest.Request
	8,  // 1: source.v1.DatumResponse.result:type_name -> source.v1.DatumResponse.Result
	9,  // 2: source.v1.AckRequest.request:type_name -> source.v1.AckRequest.Request
	10, // 3: source.v1.AckResponse.result:type_name -> source.v1.AckResponse.Result
	11, // 4: source.v1.PendingResponse.result:type_name -> source.v1.PendingResponse.Result
	6,  // 5: source.v1.DatumResponse.Result.offset:type_name -> source.v1.Offset
	12, // 6: source.v1.DatumResponse.Result.event_time:type_name -> google.protobuf.Timestamp
	6,  // 7: source.v1.AckRequest.Request.offset:type_name -> source.v1.Offset
	13, // 8: source.v1.AckResponse.Result.success:type_name -> google.protobuf.Empty
	0,  // 9: source.v1.Source.ReadFn:input_type -> source.v1.ReadRequest
	2,  // 10: source.v1.Source.AckFn:input_type -> source.v1.AckRequest
	13, // 11: source.v1.Source.PendingFn:input_type -> google.protobuf.Empty
	13, // 12: source.v1.Source.IsReady:input_type -> google.protobuf.Empty
	1,  // 13: source.v1.Source.ReadFn:output_type -> source.v1.DatumResponse
	3,  // 14: source.v1.Source.AckFn:output_type -> source.v1.AckResponse
	5,  // 15: source.v1.Source.PendingFn:output_type -> source.v1.PendingResponse
	4,  // 16: source.v1.Source.IsReady:output_type -> source.v1.ReadyResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_pkg_apis_proto_source_v1_udsource_proto_init() }
func file_pkg_apis_proto_source_v1_udsource_proto_init() {
	if File_pkg_apis_proto_source_v1_udsource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckResponse); i {
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
		file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest_Request); i {
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
			switch v := v.(*DatumResponse_Result); i {
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
		file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckRequest_Request); i {
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
		file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckResponse_Result); i {
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
		file_pkg_apis_proto_source_v1_udsource_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PendingResponse_Result); i {
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
			NumMessages:   12,
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
