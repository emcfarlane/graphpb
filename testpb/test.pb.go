// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: testpb/test.proto

package testpb

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Text      string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"` // The resource content
	UserId    string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testpb_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_testpb_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_testpb_test_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Message) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetMessageRequestOne struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // Mapped to URL path
}

func (x *GetMessageRequestOne) Reset() {
	*x = GetMessageRequestOne{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testpb_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageRequestOne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageRequestOne) ProtoMessage() {}

func (x *GetMessageRequestOne) ProtoReflect() protoreflect.Message {
	mi := &file_testpb_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageRequestOne.ProtoReflect.Descriptor instead.
func (*GetMessageRequestOne) Descriptor() ([]byte, []int) {
	return file_testpb_test_proto_rawDescGZIP(), []int{1}
}

func (x *GetMessageRequestOne) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetMessageRequestTwo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string                           `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"` // Mapped to URL path
	Revision  int64                            `protobuf:"varint,2,opt,name=revision,proto3" json:"revision,omitempty"`                   // Mapped to URL query parameter `revision`
	Sub       *GetMessageRequestTwo_SubMessage `protobuf:"bytes,3,opt,name=sub,proto3" json:"sub,omitempty"`                              // Mapped to URL query parameter `sub.subfield`
	UserId    string                           `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`          // Additional binding
}

func (x *GetMessageRequestTwo) Reset() {
	*x = GetMessageRequestTwo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testpb_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageRequestTwo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageRequestTwo) ProtoMessage() {}

func (x *GetMessageRequestTwo) ProtoReflect() protoreflect.Message {
	mi := &file_testpb_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageRequestTwo.ProtoReflect.Descriptor instead.
func (*GetMessageRequestTwo) Descriptor() ([]byte, []int) {
	return file_testpb_test_proto_rawDescGZIP(), []int{2}
}

func (x *GetMessageRequestTwo) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *GetMessageRequestTwo) GetRevision() int64 {
	if x != nil {
		return x.Revision
	}
	return 0
}

func (x *GetMessageRequestTwo) GetSub() *GetMessageRequestTwo_SubMessage {
	if x != nil {
		return x.Sub
	}
	return nil
}

func (x *GetMessageRequestTwo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UpdateMessageRequestOne struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string   `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"` // Mapped to the URL
	Message   *Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`                      // Mapped to the body
}

func (x *UpdateMessageRequestOne) Reset() {
	*x = UpdateMessageRequestOne{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testpb_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMessageRequestOne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMessageRequestOne) ProtoMessage() {}

func (x *UpdateMessageRequestOne) ProtoReflect() protoreflect.Message {
	mi := &file_testpb_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMessageRequestOne.ProtoReflect.Descriptor instead.
func (*UpdateMessageRequestOne) Descriptor() ([]byte, []int) {
	return file_testpb_test_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateMessageRequestOne) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *UpdateMessageRequestOne) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type UploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string             `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	File     *httpbody.HttpBody `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *UploadFileRequest) Reset() {
	*x = UploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testpb_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequest) ProtoMessage() {}

func (x *UploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testpb_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequest.ProtoReflect.Descriptor instead.
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return file_testpb_test_proto_rawDescGZIP(), []int{4}
}

func (x *UploadFileRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *UploadFileRequest) GetFile() *httpbody.HttpBody {
	if x != nil {
		return x.File
	}
	return nil
}

type Scalars struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp   *timestamppb.Timestamp  `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Duration    *durationpb.Duration    `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	BoolValue   *wrapperspb.BoolValue   `protobuf:"bytes,3,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	Int32Value  *wrapperspb.Int32Value  `protobuf:"bytes,4,opt,name=int32_value,json=int32Value,proto3" json:"int32_value,omitempty"`
	Int64Value  *wrapperspb.Int64Value  `protobuf:"bytes,5,opt,name=int64_value,json=int64Value,proto3" json:"int64_value,omitempty"`
	Uint32Value *wrapperspb.UInt32Value `protobuf:"bytes,6,opt,name=uint32_value,json=uint32Value,proto3" json:"uint32_value,omitempty"`
	Uint64Value *wrapperspb.UInt64Value `protobuf:"bytes,7,opt,name=uint64_value,json=uint64Value,proto3" json:"uint64_value,omitempty"`
	FloatValue  *wrapperspb.FloatValue  `protobuf:"bytes,8,opt,name=float_value,json=floatValue,proto3" json:"float_value,omitempty"`
	DoubleValue *wrapperspb.DoubleValue `protobuf:"bytes,9,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	BytesValue  *wrapperspb.BytesValue  `protobuf:"bytes,10,opt,name=bytes_value,json=bytesValue,proto3" json:"bytes_value,omitempty"`
	StringValue *wrapperspb.StringValue `protobuf:"bytes,11,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"` // TODO google.protobuf.FieldMask mask = 12;
}

func (x *Scalars) Reset() {
	*x = Scalars{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testpb_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scalars) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scalars) ProtoMessage() {}

func (x *Scalars) ProtoReflect() protoreflect.Message {
	mi := &file_testpb_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scalars.ProtoReflect.Descriptor instead.
func (*Scalars) Descriptor() ([]byte, []int) {
	return file_testpb_test_proto_rawDescGZIP(), []int{5}
}

func (x *Scalars) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Scalars) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *Scalars) GetBoolValue() *wrapperspb.BoolValue {
	if x != nil {
		return x.BoolValue
	}
	return nil
}

func (x *Scalars) GetInt32Value() *wrapperspb.Int32Value {
	if x != nil {
		return x.Int32Value
	}
	return nil
}

func (x *Scalars) GetInt64Value() *wrapperspb.Int64Value {
	if x != nil {
		return x.Int64Value
	}
	return nil
}

func (x *Scalars) GetUint32Value() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Uint32Value
	}
	return nil
}

func (x *Scalars) GetUint64Value() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Uint64Value
	}
	return nil
}

func (x *Scalars) GetFloatValue() *wrapperspb.FloatValue {
	if x != nil {
		return x.FloatValue
	}
	return nil
}

func (x *Scalars) GetDoubleValue() *wrapperspb.DoubleValue {
	if x != nil {
		return x.DoubleValue
	}
	return nil
}

func (x *Scalars) GetBytesValue() *wrapperspb.BytesValue {
	if x != nil {
		return x.BytesValue
	}
	return nil
}

func (x *Scalars) GetStringValue() *wrapperspb.StringValue {
	if x != nil {
		return x.StringValue
	}
	return nil
}

type GetMessageRequestTwo_SubMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subfield string `protobuf:"bytes,1,opt,name=subfield,proto3" json:"subfield,omitempty"`
}

func (x *GetMessageRequestTwo_SubMessage) Reset() {
	*x = GetMessageRequestTwo_SubMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testpb_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageRequestTwo_SubMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageRequestTwo_SubMessage) ProtoMessage() {}

func (x *GetMessageRequestTwo_SubMessage) ProtoReflect() protoreflect.Message {
	mi := &file_testpb_test_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageRequestTwo_SubMessage.ProtoReflect.Descriptor instead.
func (*GetMessageRequestTwo_SubMessage) Descriptor() ([]byte, []int) {
	return file_testpb_test_proto_rawDescGZIP(), []int{2, 0}
}

func (x *GetMessageRequestTwo_SubMessage) GetSubfield() string {
	if x != nil {
		return x.Subfield
	}
	return ""
}

var File_testpb_test_proto protoreflect.FileDescriptor

var file_testpb_test_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x70, 0x62, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74,
	0x74, 0x70, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x2a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xd7, 0x01,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x54, 0x77, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x41, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x70, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x54, 0x77, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x03, 0x73, 0x75, 0x62, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x28, 0x0a,
	0x0a, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x75, 0x62, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x75, 0x62, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x6b, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f,
	0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x12, 0x31, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x70, 0x62, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x59, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22,
	0xb1, 0x05, 0x0a, 0x07, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a,
	0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x62, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f,
	0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x32, 0xbb, 0x06, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x12, 0x72, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x6e, 0x65, 0x12, 0x24, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x70, 0x62, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x6e, 0x65, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x70, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x6e, 0x61,
	0x6d, 0x65, 0x2f, 0x2a, 0x7d, 0x12, 0xbe, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x77, 0x6f, 0x12, 0x24, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x70,
	0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x77, 0x6f, 0x1a, 0x17, 0x2e,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x70, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x68, 0x12, 0x19,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x5a, 0x1e, 0x12, 0x1c, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5a, 0x2b, 0x12, 0x29, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x7d, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x70,
	0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x6e, 0x65,
	0x1a, 0x17, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x70, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x24, 0x32, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f,
	0x7b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x70, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x70, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x70, 0x62, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x23, 0x32, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f,
	0x62, 0x6f, 0x64, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x5e, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x70, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f,
	0x7b, 0x74, 0x65, 0x78, 0x74, 0x3d, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x3a, 0x63, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x53, 0x0a, 0x0b, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x4f, 0x6e, 0x65, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x70, 0x62,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12,
	0x0b, 0x2f, 0x7b, 0x74, 0x65, 0x78, 0x74, 0x7d, 0x2f, 0x6f, 0x6e, 0x65, 0x12, 0x53, 0x0a, 0x0b,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x77, 0x6f, 0x12, 0x17, 0x2e, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x70, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x13, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x7b, 0x74, 0x65, 0x78, 0x74, 0x7d, 0x2f, 0x74, 0x77,
	0x6f, 0x32, 0xee, 0x01, 0x0a, 0x05, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x6a, 0x0a, 0x0e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x21, 0x2e,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x70, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74,
	0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x11,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x7d, 0x3a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x79, 0x0a, 0x13, 0x4c, 0x61, 0x72, 0x67, 0x65,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x21,
	0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x70, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48,
	0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22,
	0x17, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x2f, 0x7b, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x3a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x28, 0x01,
	0x30, 0x01, 0x32, 0x5c, 0x0a, 0x09, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x12,
	0x4f, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x70, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72,
	0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6d, 0x63, 0x66, 0x61, 0x72, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x70,
	0x62, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testpb_test_proto_rawDescOnce sync.Once
	file_testpb_test_proto_rawDescData = file_testpb_test_proto_rawDesc
)

func file_testpb_test_proto_rawDescGZIP() []byte {
	file_testpb_test_proto_rawDescOnce.Do(func() {
		file_testpb_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_testpb_test_proto_rawDescData)
	})
	return file_testpb_test_proto_rawDescData
}

var file_testpb_test_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_testpb_test_proto_goTypes = []interface{}{
	(*Message)(nil),                         // 0: graphpb.testpb.Message
	(*GetMessageRequestOne)(nil),            // 1: graphpb.testpb.GetMessageRequestOne
	(*GetMessageRequestTwo)(nil),            // 2: graphpb.testpb.GetMessageRequestTwo
	(*UpdateMessageRequestOne)(nil),         // 3: graphpb.testpb.UpdateMessageRequestOne
	(*UploadFileRequest)(nil),               // 4: graphpb.testpb.UploadFileRequest
	(*Scalars)(nil),                         // 5: graphpb.testpb.Scalars
	(*GetMessageRequestTwo_SubMessage)(nil), // 6: graphpb.testpb.GetMessageRequestTwo.SubMessage
	(*httpbody.HttpBody)(nil),               // 7: google.api.HttpBody
	(*timestamppb.Timestamp)(nil),           // 8: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),             // 9: google.protobuf.Duration
	(*wrapperspb.BoolValue)(nil),            // 10: google.protobuf.BoolValue
	(*wrapperspb.Int32Value)(nil),           // 11: google.protobuf.Int32Value
	(*wrapperspb.Int64Value)(nil),           // 12: google.protobuf.Int64Value
	(*wrapperspb.UInt32Value)(nil),          // 13: google.protobuf.UInt32Value
	(*wrapperspb.UInt64Value)(nil),          // 14: google.protobuf.UInt64Value
	(*wrapperspb.FloatValue)(nil),           // 15: google.protobuf.FloatValue
	(*wrapperspb.DoubleValue)(nil),          // 16: google.protobuf.DoubleValue
	(*wrapperspb.BytesValue)(nil),           // 17: google.protobuf.BytesValue
	(*wrapperspb.StringValue)(nil),          // 18: google.protobuf.StringValue
	(*emptypb.Empty)(nil),                   // 19: google.protobuf.Empty
}
var file_testpb_test_proto_depIdxs = []int32{
	6,  // 0: graphpb.testpb.GetMessageRequestTwo.sub:type_name -> graphpb.testpb.GetMessageRequestTwo.SubMessage
	0,  // 1: graphpb.testpb.UpdateMessageRequestOne.message:type_name -> graphpb.testpb.Message
	7,  // 2: graphpb.testpb.UploadFileRequest.file:type_name -> google.api.HttpBody
	8,  // 3: graphpb.testpb.Scalars.timestamp:type_name -> google.protobuf.Timestamp
	9,  // 4: graphpb.testpb.Scalars.duration:type_name -> google.protobuf.Duration
	10, // 5: graphpb.testpb.Scalars.bool_value:type_name -> google.protobuf.BoolValue
	11, // 6: graphpb.testpb.Scalars.int32_value:type_name -> google.protobuf.Int32Value
	12, // 7: graphpb.testpb.Scalars.int64_value:type_name -> google.protobuf.Int64Value
	13, // 8: graphpb.testpb.Scalars.uint32_value:type_name -> google.protobuf.UInt32Value
	14, // 9: graphpb.testpb.Scalars.uint64_value:type_name -> google.protobuf.UInt64Value
	15, // 10: graphpb.testpb.Scalars.float_value:type_name -> google.protobuf.FloatValue
	16, // 11: graphpb.testpb.Scalars.double_value:type_name -> google.protobuf.DoubleValue
	17, // 12: graphpb.testpb.Scalars.bytes_value:type_name -> google.protobuf.BytesValue
	18, // 13: graphpb.testpb.Scalars.string_value:type_name -> google.protobuf.StringValue
	1,  // 14: graphpb.testpb.Messaging.GetMessageOne:input_type -> graphpb.testpb.GetMessageRequestOne
	2,  // 15: graphpb.testpb.Messaging.GetMessageTwo:input_type -> graphpb.testpb.GetMessageRequestTwo
	3,  // 16: graphpb.testpb.Messaging.UpdateMessage:input_type -> graphpb.testpb.UpdateMessageRequestOne
	0,  // 17: graphpb.testpb.Messaging.UpdateMessageBody:input_type -> graphpb.testpb.Message
	0,  // 18: graphpb.testpb.Messaging.Action:input_type -> graphpb.testpb.Message
	0,  // 19: graphpb.testpb.Messaging.VariableOne:input_type -> graphpb.testpb.Message
	0,  // 20: graphpb.testpb.Messaging.VariableTwo:input_type -> graphpb.testpb.Message
	4,  // 21: graphpb.testpb.Files.UploadDownload:input_type -> graphpb.testpb.UploadFileRequest
	4,  // 22: graphpb.testpb.Files.LargeUploadDownload:input_type -> graphpb.testpb.UploadFileRequest
	5,  // 23: graphpb.testpb.WellKnown.Check:input_type -> graphpb.testpb.Scalars
	0,  // 24: graphpb.testpb.Messaging.GetMessageOne:output_type -> graphpb.testpb.Message
	0,  // 25: graphpb.testpb.Messaging.GetMessageTwo:output_type -> graphpb.testpb.Message
	0,  // 26: graphpb.testpb.Messaging.UpdateMessage:output_type -> graphpb.testpb.Message
	0,  // 27: graphpb.testpb.Messaging.UpdateMessageBody:output_type -> graphpb.testpb.Message
	19, // 28: graphpb.testpb.Messaging.Action:output_type -> google.protobuf.Empty
	19, // 29: graphpb.testpb.Messaging.VariableOne:output_type -> google.protobuf.Empty
	19, // 30: graphpb.testpb.Messaging.VariableTwo:output_type -> google.protobuf.Empty
	7,  // 31: graphpb.testpb.Files.UploadDownload:output_type -> google.api.HttpBody
	7,  // 32: graphpb.testpb.Files.LargeUploadDownload:output_type -> google.api.HttpBody
	19, // 33: graphpb.testpb.WellKnown.Check:output_type -> google.protobuf.Empty
	24, // [24:34] is the sub-list for method output_type
	14, // [14:24] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_testpb_test_proto_init() }
func file_testpb_test_proto_init() {
	if File_testpb_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testpb_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_testpb_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageRequestOne); i {
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
		file_testpb_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageRequestTwo); i {
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
		file_testpb_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMessageRequestOne); i {
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
		file_testpb_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileRequest); i {
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
		file_testpb_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scalars); i {
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
		file_testpb_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageRequestTwo_SubMessage); i {
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
			RawDescriptor: file_testpb_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_testpb_test_proto_goTypes,
		DependencyIndexes: file_testpb_test_proto_depIdxs,
		MessageInfos:      file_testpb_test_proto_msgTypes,
	}.Build()
	File_testpb_test_proto = out.File
	file_testpb_test_proto_rawDesc = nil
	file_testpb_test_proto_goTypes = nil
	file_testpb_test_proto_depIdxs = nil
}
