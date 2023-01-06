// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: services/calculate/v1/calculate.proto

package calculatev1

import (
	v1 "github.com/pitakill/proto.test/go/calculator/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AdditionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation      v1.Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=calculator.v1.Operation" json:"operation,omitempty"`
	OperatorFirst  float64      `protobuf:"fixed64,2,opt,name=operator_first,json=operatorFirst,proto3" json:"operator_first,omitempty"`
	OperatorSecond float64      `protobuf:"fixed64,3,opt,name=operator_second,json=operatorSecond,proto3" json:"operator_second,omitempty"`
}

func (x *AdditionRequest) Reset() {
	*x = AdditionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_calculate_v1_calculate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdditionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdditionRequest) ProtoMessage() {}

func (x *AdditionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_calculate_v1_calculate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdditionRequest.ProtoReflect.Descriptor instead.
func (*AdditionRequest) Descriptor() ([]byte, []int) {
	return file_services_calculate_v1_calculate_proto_rawDescGZIP(), []int{0}
}

func (x *AdditionRequest) GetOperation() v1.Operation {
	if x != nil {
		return x.Operation
	}
	return v1.Operation(0)
}

func (x *AdditionRequest) GetOperatorFirst() float64 {
	if x != nil {
		return x.OperatorFirst
	}
	return 0
}

func (x *AdditionRequest) GetOperatorSecond() float64 {
	if x != nil {
		return x.OperatorSecond
	}
	return 0
}

type AdditionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float64 `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AdditionResponse) Reset() {
	*x = AdditionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_calculate_v1_calculate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdditionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdditionResponse) ProtoMessage() {}

func (x *AdditionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_calculate_v1_calculate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdditionResponse.ProtoReflect.Descriptor instead.
func (*AdditionResponse) Descriptor() ([]byte, []int) {
	return file_services_calculate_v1_calculate_proto_rawDescGZIP(), []int{1}
}

func (x *AdditionResponse) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

type AdditionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation      v1.Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=calculator.v1.Operation" json:"operation,omitempty"`
	OperatorFirst  float64      `protobuf:"fixed64,2,opt,name=operator_first,json=operatorFirst,proto3" json:"operator_first,omitempty"`
	OperatorSecond float64      `protobuf:"fixed64,3,opt,name=operator_second,json=operatorSecond,proto3" json:"operator_second,omitempty"`
}

func (x *AdditionsRequest) Reset() {
	*x = AdditionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_calculate_v1_calculate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdditionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdditionsRequest) ProtoMessage() {}

func (x *AdditionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_calculate_v1_calculate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdditionsRequest.ProtoReflect.Descriptor instead.
func (*AdditionsRequest) Descriptor() ([]byte, []int) {
	return file_services_calculate_v1_calculate_proto_rawDescGZIP(), []int{2}
}

func (x *AdditionsRequest) GetOperation() v1.Operation {
	if x != nil {
		return x.Operation
	}
	return v1.Operation(0)
}

func (x *AdditionsRequest) GetOperatorFirst() float64 {
	if x != nil {
		return x.OperatorFirst
	}
	return 0
}

func (x *AdditionsRequest) GetOperatorSecond() float64 {
	if x != nil {
		return x.OperatorSecond
	}
	return 0
}

type AdditionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float64             `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	Errors map[string]v1.Error `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=calculator.v1.Error"`
}

func (x *AdditionsResponse) Reset() {
	*x = AdditionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_calculate_v1_calculate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdditionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdditionsResponse) ProtoMessage() {}

func (x *AdditionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_calculate_v1_calculate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdditionsResponse.ProtoReflect.Descriptor instead.
func (*AdditionsResponse) Descriptor() ([]byte, []int) {
	return file_services_calculate_v1_calculate_proto_rawDescGZIP(), []int{3}
}

func (x *AdditionsResponse) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *AdditionsResponse) GetErrors() map[string]v1.Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

type DivisionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation      v1.Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=calculator.v1.Operation" json:"operation,omitempty"`
	OperatorFirst  float64      `protobuf:"fixed64,2,opt,name=operator_first,json=operatorFirst,proto3" json:"operator_first,omitempty"`
	OperatorSecond float64      `protobuf:"fixed64,3,opt,name=operator_second,json=operatorSecond,proto3" json:"operator_second,omitempty"`
}

func (x *DivisionRequest) Reset() {
	*x = DivisionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_calculate_v1_calculate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DivisionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DivisionRequest) ProtoMessage() {}

func (x *DivisionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_calculate_v1_calculate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DivisionRequest.ProtoReflect.Descriptor instead.
func (*DivisionRequest) Descriptor() ([]byte, []int) {
	return file_services_calculate_v1_calculate_proto_rawDescGZIP(), []int{4}
}

func (x *DivisionRequest) GetOperation() v1.Operation {
	if x != nil {
		return x.Operation
	}
	return v1.Operation(0)
}

func (x *DivisionRequest) GetOperatorFirst() float64 {
	if x != nil {
		return x.OperatorFirst
	}
	return 0
}

func (x *DivisionRequest) GetOperatorSecond() float64 {
	if x != nil {
		return x.OperatorSecond
	}
	return 0
}

type DivisionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float64 `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DivisionResponse) Reset() {
	*x = DivisionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_calculate_v1_calculate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DivisionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DivisionResponse) ProtoMessage() {}

func (x *DivisionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_calculate_v1_calculate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DivisionResponse.ProtoReflect.Descriptor instead.
func (*DivisionResponse) Descriptor() ([]byte, []int) {
	return file_services_calculate_v1_calculate_proto_rawDescGZIP(), []int{5}
}

func (x *DivisionResponse) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_services_calculate_v1_calculate_proto protoreflect.FileDescriptor

var file_services_calculate_v1_calculate_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1e,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99,
	0x01, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x46, 0x69, 0x72, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x22, 0x2a, 0x0a, 0x10, 0x41, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x9a, 0x01, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x46, 0x69, 0x72, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x22, 0xca, 0x01, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x4c, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a,
	0x4f, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x99, 0x01, 0x0a, 0x0f, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x46, 0x69,
	0x72, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x22, 0x2a, 0x0a, 0x10,
	0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xb6, 0x02, 0x0a, 0x10, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a,
	0x08, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x08,
	0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x09, 0x41,
	0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x85, 0x02, 0x0a, 0x38, 0x6e, 0x65, 0x74, 0x2e, 0x70, 0x69, 0x74, 0x61, 0x6b, 0x69,
	0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0e,
	0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x74,
	0x61, 0x6b, 0x69, 0x6c, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x43, 0x58, 0xaa, 0x02, 0x15, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x21, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x17, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_services_calculate_v1_calculate_proto_rawDescOnce sync.Once
	file_services_calculate_v1_calculate_proto_rawDescData = file_services_calculate_v1_calculate_proto_rawDesc
)

func file_services_calculate_v1_calculate_proto_rawDescGZIP() []byte {
	file_services_calculate_v1_calculate_proto_rawDescOnce.Do(func() {
		file_services_calculate_v1_calculate_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_calculate_v1_calculate_proto_rawDescData)
	})
	return file_services_calculate_v1_calculate_proto_rawDescData
}

var file_services_calculate_v1_calculate_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_services_calculate_v1_calculate_proto_goTypes = []interface{}{
	(*AdditionRequest)(nil),   // 0: services.calculate.v1.AdditionRequest
	(*AdditionResponse)(nil),  // 1: services.calculate.v1.AdditionResponse
	(*AdditionsRequest)(nil),  // 2: services.calculate.v1.AdditionsRequest
	(*AdditionsResponse)(nil), // 3: services.calculate.v1.AdditionsResponse
	(*DivisionRequest)(nil),   // 4: services.calculate.v1.DivisionRequest
	(*DivisionResponse)(nil),  // 5: services.calculate.v1.DivisionResponse
	nil,                       // 6: services.calculate.v1.AdditionsResponse.ErrorsEntry
	(v1.Operation)(0),         // 7: calculator.v1.Operation
	(v1.Error)(0),             // 8: calculator.v1.Error
}
var file_services_calculate_v1_calculate_proto_depIdxs = []int32{
	7, // 0: services.calculate.v1.AdditionRequest.operation:type_name -> calculator.v1.Operation
	7, // 1: services.calculate.v1.AdditionsRequest.operation:type_name -> calculator.v1.Operation
	6, // 2: services.calculate.v1.AdditionsResponse.errors:type_name -> services.calculate.v1.AdditionsResponse.ErrorsEntry
	7, // 3: services.calculate.v1.DivisionRequest.operation:type_name -> calculator.v1.Operation
	8, // 4: services.calculate.v1.AdditionsResponse.ErrorsEntry.value:type_name -> calculator.v1.Error
	0, // 5: services.calculate.v1.CalculateService.Addition:input_type -> services.calculate.v1.AdditionRequest
	4, // 6: services.calculate.v1.CalculateService.Division:input_type -> services.calculate.v1.DivisionRequest
	2, // 7: services.calculate.v1.CalculateService.Additions:input_type -> services.calculate.v1.AdditionsRequest
	1, // 8: services.calculate.v1.CalculateService.Addition:output_type -> services.calculate.v1.AdditionResponse
	5, // 9: services.calculate.v1.CalculateService.Division:output_type -> services.calculate.v1.DivisionResponse
	3, // 10: services.calculate.v1.CalculateService.Additions:output_type -> services.calculate.v1.AdditionsResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_services_calculate_v1_calculate_proto_init() }
func file_services_calculate_v1_calculate_proto_init() {
	if File_services_calculate_v1_calculate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_calculate_v1_calculate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdditionRequest); i {
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
		file_services_calculate_v1_calculate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdditionResponse); i {
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
		file_services_calculate_v1_calculate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdditionsRequest); i {
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
		file_services_calculate_v1_calculate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdditionsResponse); i {
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
		file_services_calculate_v1_calculate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DivisionRequest); i {
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
		file_services_calculate_v1_calculate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DivisionResponse); i {
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
			RawDescriptor: file_services_calculate_v1_calculate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_calculate_v1_calculate_proto_goTypes,
		DependencyIndexes: file_services_calculate_v1_calculate_proto_depIdxs,
		MessageInfos:      file_services_calculate_v1_calculate_proto_msgTypes,
	}.Build()
	File_services_calculate_v1_calculate_proto = out.File
	file_services_calculate_v1_calculate_proto_rawDesc = nil
	file_services_calculate_v1_calculate_proto_goTypes = nil
	file_services_calculate_v1_calculate_proto_depIdxs = nil
}
