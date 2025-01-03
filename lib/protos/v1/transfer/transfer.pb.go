// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.29.1
// source: lib/protos/v1/transfer/transfer.proto

package transfer

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	money "google.golang.org/genproto/googleapis/type/money"
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

type TransferType int32

const (
	TransferType_UNSPECIFIED     TransferType = 0
	TransferType_BETWEEN_ACCOUNT TransferType = 1
	TransferType_BETWEEN_BANK    TransferType = 2
)

// Enum value maps for TransferType.
var (
	TransferType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "BETWEEN_ACCOUNT",
		2: "BETWEEN_BANK",
	}
	TransferType_value = map[string]int32{
		"UNSPECIFIED":     0,
		"BETWEEN_ACCOUNT": 1,
		"BETWEEN_BANK":    2,
	}
)

func (x TransferType) Enum() *TransferType {
	p := new(TransferType)
	*p = x
	return p
}

func (x TransferType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransferType) Descriptor() protoreflect.EnumDescriptor {
	return file_lib_protos_v1_transfer_transfer_proto_enumTypes[0].Descriptor()
}

func (TransferType) Type() protoreflect.EnumType {
	return &file_lib_protos_v1_transfer_transfer_proto_enumTypes[0]
}

func (x TransferType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransferType.Descriptor instead.
func (TransferType) EnumDescriptor() ([]byte, []int) {
	return file_lib_protos_v1_transfer_transfer_proto_rawDescGZIP(), []int{0}
}

type TransferInfoPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrxId string `protobuf:"bytes,1,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
}

func (x *TransferInfoPayload) Reset() {
	*x = TransferInfoPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_protos_v1_transfer_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferInfoPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferInfoPayload) ProtoMessage() {}

func (x *TransferInfoPayload) ProtoReflect() protoreflect.Message {
	mi := &file_lib_protos_v1_transfer_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferInfoPayload.ProtoReflect.Descriptor instead.
func (*TransferInfoPayload) Descriptor() ([]byte, []int) {
	return file_lib_protos_v1_transfer_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *TransferInfoPayload) GetTrxId() string {
	if x != nil {
		return x.TrxId
	}
	return ""
}

type TransferInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrxId     string       `protobuf:"bytes,1,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	AccountNo string       `protobuf:"bytes,2,opt,name=account_no,json=accountNo,proto3" json:"account_no,omitempty"`
	Amount    *money.Money `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransferInfoResponse) Reset() {
	*x = TransferInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_protos_v1_transfer_transfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferInfoResponse) ProtoMessage() {}

func (x *TransferInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lib_protos_v1_transfer_transfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferInfoResponse.ProtoReflect.Descriptor instead.
func (*TransferInfoResponse) Descriptor() ([]byte, []int) {
	return file_lib_protos_v1_transfer_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *TransferInfoResponse) GetTrxId() string {
	if x != nil {
		return x.TrxId
	}
	return ""
}

func (x *TransferInfoResponse) GetAccountNo() string {
	if x != nil {
		return x.AccountNo
	}
	return ""
}

func (x *TransferInfoResponse) GetAmount() *money.Money {
	if x != nil {
		return x.Amount
	}
	return nil
}

type TransferBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceAccountNumber string       `protobuf:"bytes,1,opt,name=source_account_number,json=sourceAccountNumber,proto3" json:"source_account_number,omitempty"`
	Amount              int64        `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	DestinationAccount  string       `protobuf:"bytes,3,opt,name=destination_account,json=destinationAccount,proto3" json:"destination_account,omitempty"`
	AccountCode         int64        `protobuf:"varint,4,opt,name=account_code,json=accountCode,proto3" json:"account_code,omitempty"`
	TransferType        TransferType `protobuf:"varint,5,opt,name=transfer_type,json=transferType,proto3,enum=fastcampus.transfer.v1.TransferType" json:"transfer_type,omitempty"`
	RequestId           string       `protobuf:"bytes,6,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *TransferBalanceRequest) Reset() {
	*x = TransferBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_protos_v1_transfer_transfer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferBalanceRequest) ProtoMessage() {}

func (x *TransferBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lib_protos_v1_transfer_transfer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferBalanceRequest.ProtoReflect.Descriptor instead.
func (*TransferBalanceRequest) Descriptor() ([]byte, []int) {
	return file_lib_protos_v1_transfer_transfer_proto_rawDescGZIP(), []int{2}
}

func (x *TransferBalanceRequest) GetSourceAccountNumber() string {
	if x != nil {
		return x.SourceAccountNumber
	}
	return ""
}

func (x *TransferBalanceRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransferBalanceRequest) GetDestinationAccount() string {
	if x != nil {
		return x.DestinationAccount
	}
	return ""
}

func (x *TransferBalanceRequest) GetAccountCode() int64 {
	if x != nil {
		return x.AccountCode
	}
	return 0
}

func (x *TransferBalanceRequest) GetTransferType() TransferType {
	if x != nil {
		return x.TransferType
	}
	return TransferType_UNSPECIFIED
}

func (x *TransferBalanceRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type TransferBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	TrxId   string `protobuf:"bytes,3,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	Amount  int64  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransferBalanceResponse) Reset() {
	*x = TransferBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_protos_v1_transfer_transfer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferBalanceResponse) ProtoMessage() {}

func (x *TransferBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lib_protos_v1_transfer_transfer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferBalanceResponse.ProtoReflect.Descriptor instead.
func (*TransferBalanceResponse) Descriptor() ([]byte, []int) {
	return file_lib_protos_v1_transfer_transfer_proto_rawDescGZIP(), []int{3}
}

func (x *TransferBalanceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TransferBalanceResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TransferBalanceResponse) GetTrxId() string {
	if x != nil {
		return x.TrxId
	}
	return ""
}

func (x *TransferBalanceResponse) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransferHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *TransferHistoryRequest) Reset() {
	*x = TransferHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_protos_v1_transfer_transfer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferHistoryRequest) ProtoMessage() {}

func (x *TransferHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lib_protos_v1_transfer_transfer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferHistoryRequest.ProtoReflect.Descriptor instead.
func (*TransferHistoryRequest) Descriptor() ([]byte, []int) {
	return file_lib_protos_v1_transfer_transfer_proto_rawDescGZIP(), []int{4}
}

func (x *TransferHistoryRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type TransferHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransferHistory []*TransferBalanceResponse `protobuf:"bytes,1,rep,name=transfer_history,json=transferHistory,proto3" json:"transfer_history,omitempty"`
}

func (x *TransferHistoryResponse) Reset() {
	*x = TransferHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_protos_v1_transfer_transfer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferHistoryResponse) ProtoMessage() {}

func (x *TransferHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lib_protos_v1_transfer_transfer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferHistoryResponse.ProtoReflect.Descriptor instead.
func (*TransferHistoryResponse) Descriptor() ([]byte, []int) {
	return file_lib_protos_v1_transfer_transfer_proto_rawDescGZIP(), []int{5}
}

func (x *TransferHistoryResponse) GetTransferHistory() []*TransferBalanceResponse {
	if x != nil {
		return x.TransferHistory
	}
	return nil
}

var File_lib_protos_v1_transfer_transfer_proto protoreflect.FileDescriptor

var file_lib_protos_v1_transfer_transfer_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x66, 0x61, 0x73, 0x74, 0x63, 0x61, 0x6d,
	0x70, 0x75, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x27, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x6c, 0x69,
	0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x74, 0x72, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x72, 0x78, 0x49, 0x64, 0x22, 0x78, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a,
	0x06, 0x74, 0x72, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x72, 0x78, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4e, 0x6f, 0x12, 0x2a, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0xa2, 0x02, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x49, 0x0a, 0x0d, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x24, 0x2e, 0x66, 0x61, 0x73, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x22, 0x7a, 0x0a, 0x17, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x72, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x72, 0x78, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x2c, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x75,
	0x0a, 0x17, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x10, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x66, 0x61, 0x73, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x2a, 0x46, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x45, 0x54, 0x57, 0x45, 0x45,
	0x4e, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x42,
	0x45, 0x54, 0x57, 0x45, 0x45, 0x4e, 0x5f, 0x42, 0x41, 0x4e, 0x4b, 0x10, 0x02, 0x32, 0xbc, 0x04,
	0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x72, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x12, 0x2b, 0x2e, 0x66, 0x61, 0x73, 0x74, 0x63,
	0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x2c, 0x2e, 0x66, 0x61, 0x73, 0x74, 0x63, 0x61, 0x6d, 0x70,
	0x75, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xda, 0x01, 0x0a, 0x15, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x2e, 0x2e, 0x66, 0x61, 0x73, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2f, 0x2e, 0x66, 0x61, 0x73, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x60, 0x92, 0x41, 0x3e, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12,
	0x15, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x55, 0x73, 0x65, 0x72, 0x20, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2a, 0x0e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4a, 0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x04, 0x0a,
	0x02, 0x4f, 0x4b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0xd7, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2e, 0x2e, 0x66, 0x61, 0x73, 0x74,
	0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x66, 0x61, 0x73, 0x74,
	0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60, 0x92, 0x41, 0x41, 0x0a,
	0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x14, 0x47, 0x65, 0x74, 0x20, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2a,
	0x12, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x4a, 0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x04, 0x0a, 0x02, 0x4f, 0x4b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x42, 0xf9, 0x01, 0x92,
	0x41, 0xa4, 0x01, 0x12, 0x7b, 0x0a, 0x27, 0x46, 0x61, 0x73, 0x74, 0x20, 0x43, 0x61, 0x6d, 0x70,
	0x75, 0x73, 0x20, 0x50, 0x61, 0x79, 0x20, 0x28, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x20, 0x41, 0x50, 0x49, 0x20, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x29, 0x22, 0x4b,
	0x0a, 0x13, 0x46, 0x61, 0x73, 0x74, 0x20, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x20, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x74, 0x69, 0x6b, 0x74,
	0x65, 0x72, 0x61, 0x6e, 0x67, 0x1a, 0x14, 0x61, 0x7a, 0x77, 0x61, 0x72, 0x2e, 0x6e, 0x72, 0x73,
	0x74, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31, 0x2e, 0x30,
	0x2a, 0x01, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x74, 0x69, 0x6b, 0x74, 0x65, 0x72, 0x61, 0x6e, 0x67, 0x2f,
	0x68, 0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e, 0x61, 0x6c, 0x2d, 0x66, 0x61, 0x73, 0x74, 0x63, 0x61,
	0x6d, 0x70, 0x75, 0x73, 0x2d, 0x70, 0x61, 0x79, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x3b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lib_protos_v1_transfer_transfer_proto_rawDescOnce sync.Once
	file_lib_protos_v1_transfer_transfer_proto_rawDescData = file_lib_protos_v1_transfer_transfer_proto_rawDesc
)

func file_lib_protos_v1_transfer_transfer_proto_rawDescGZIP() []byte {
	file_lib_protos_v1_transfer_transfer_proto_rawDescOnce.Do(func() {
		file_lib_protos_v1_transfer_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_lib_protos_v1_transfer_transfer_proto_rawDescData)
	})
	return file_lib_protos_v1_transfer_transfer_proto_rawDescData
}

var file_lib_protos_v1_transfer_transfer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_lib_protos_v1_transfer_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_lib_protos_v1_transfer_transfer_proto_goTypes = []any{
	(TransferType)(0),               // 0: fastcampus.transfer.v1.TransferType
	(*TransferInfoPayload)(nil),     // 1: fastcampus.transfer.v1.TransferInfoPayload
	(*TransferInfoResponse)(nil),    // 2: fastcampus.transfer.v1.TransferInfoResponse
	(*TransferBalanceRequest)(nil),  // 3: fastcampus.transfer.v1.TransferBalanceRequest
	(*TransferBalanceResponse)(nil), // 4: fastcampus.transfer.v1.TransferBalanceResponse
	(*TransferHistoryRequest)(nil),  // 5: fastcampus.transfer.v1.TransferHistoryRequest
	(*TransferHistoryResponse)(nil), // 6: fastcampus.transfer.v1.TransferHistoryResponse
	(*money.Money)(nil),             // 7: google.type.Money
}
var file_lib_protos_v1_transfer_transfer_proto_depIdxs = []int32{
	7, // 0: fastcampus.transfer.v1.TransferInfoResponse.amount:type_name -> google.type.Money
	0, // 1: fastcampus.transfer.v1.TransferBalanceRequest.transfer_type:type_name -> fastcampus.transfer.v1.TransferType
	4, // 2: fastcampus.transfer.v1.TransferHistoryResponse.transfer_history:type_name -> fastcampus.transfer.v1.TransferBalanceResponse
	1, // 3: fastcampus.transfer.v1.TransferService.GetTransferInfoByID:input_type -> fastcampus.transfer.v1.TransferInfoPayload
	3, // 4: fastcampus.transfer.v1.TransferService.SubmitTransferBalance:input_type -> fastcampus.transfer.v1.TransferBalanceRequest
	5, // 5: fastcampus.transfer.v1.TransferService.GetTransferHistory:input_type -> fastcampus.transfer.v1.TransferHistoryRequest
	2, // 6: fastcampus.transfer.v1.TransferService.GetTransferInfoByID:output_type -> fastcampus.transfer.v1.TransferInfoResponse
	4, // 7: fastcampus.transfer.v1.TransferService.SubmitTransferBalance:output_type -> fastcampus.transfer.v1.TransferBalanceResponse
	6, // 8: fastcampus.transfer.v1.TransferService.GetTransferHistory:output_type -> fastcampus.transfer.v1.TransferHistoryResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_lib_protos_v1_transfer_transfer_proto_init() }
func file_lib_protos_v1_transfer_transfer_proto_init() {
	if File_lib_protos_v1_transfer_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lib_protos_v1_transfer_transfer_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TransferInfoPayload); i {
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
		file_lib_protos_v1_transfer_transfer_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TransferInfoResponse); i {
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
		file_lib_protos_v1_transfer_transfer_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TransferBalanceRequest); i {
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
		file_lib_protos_v1_transfer_transfer_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TransferBalanceResponse); i {
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
		file_lib_protos_v1_transfer_transfer_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*TransferHistoryRequest); i {
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
		file_lib_protos_v1_transfer_transfer_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*TransferHistoryResponse); i {
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
			RawDescriptor: file_lib_protos_v1_transfer_transfer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lib_protos_v1_transfer_transfer_proto_goTypes,
		DependencyIndexes: file_lib_protos_v1_transfer_transfer_proto_depIdxs,
		EnumInfos:         file_lib_protos_v1_transfer_transfer_proto_enumTypes,
		MessageInfos:      file_lib_protos_v1_transfer_transfer_proto_msgTypes,
	}.Build()
	File_lib_protos_v1_transfer_transfer_proto = out.File
	file_lib_protos_v1_transfer_transfer_proto_rawDesc = nil
	file_lib_protos_v1_transfer_transfer_proto_goTypes = nil
	file_lib_protos_v1_transfer_transfer_proto_depIdxs = nil
}
