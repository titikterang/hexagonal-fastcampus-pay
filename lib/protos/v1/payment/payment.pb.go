// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: lib/protos/v1/payment/payment.proto

package payment

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PaymentStatus int32

const (
	PaymentStatus_UNSPECIFIED PaymentStatus = 0
	PaymentStatus_UNPAID      PaymentStatus = 1
	PaymentStatus_PENDING     PaymentStatus = 2
	PaymentStatus_FAILED      PaymentStatus = 3
	PaymentStatus_PAID        PaymentStatus = 4
)

// Enum value maps for PaymentStatus.
var (
	PaymentStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNPAID",
		2: "PENDING",
		3: "FAILED",
		4: "PAID",
	}
	PaymentStatus_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNPAID":      1,
		"PENDING":     2,
		"FAILED":      3,
		"PAID":        4,
	}
)

func (x PaymentStatus) Enum() *PaymentStatus {
	p := new(PaymentStatus)
	*p = x
	return p
}

func (x PaymentStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_lib_protos_v1_payment_payment_proto_enumTypes[0].Descriptor()
}

func (PaymentStatus) Type() protoreflect.EnumType {
	return &file_lib_protos_v1_payment_payment_proto_enumTypes[0]
}

func (x PaymentStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentStatus.Descriptor instead.
func (PaymentStatus) EnumDescriptor() ([]byte, []int) {
	return file_lib_protos_v1_payment_payment_proto_rawDescGZIP(), []int{0}
}

type PaymentStatusPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvoiceId string `protobuf:"bytes,1,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
}

func (x *PaymentStatusPayload) Reset() {
	*x = PaymentStatusPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_protos_v1_payment_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentStatusPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentStatusPayload) ProtoMessage() {}

func (x *PaymentStatusPayload) ProtoReflect() protoreflect.Message {
	mi := &file_lib_protos_v1_payment_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentStatusPayload.ProtoReflect.Descriptor instead.
func (*PaymentStatusPayload) Descriptor() ([]byte, []int) {
	return file_lib_protos_v1_payment_payment_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentStatusPayload) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

type PaymentPrecheckPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountNo string `protobuf:"bytes,1,opt,name=account_no,json=accountNo,proto3" json:"account_no,omitempty"`
}

func (x *PaymentPrecheckPayload) Reset() {
	*x = PaymentPrecheckPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_protos_v1_payment_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentPrecheckPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentPrecheckPayload) ProtoMessage() {}

func (x *PaymentPrecheckPayload) ProtoReflect() protoreflect.Message {
	mi := &file_lib_protos_v1_payment_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentPrecheckPayload.ProtoReflect.Descriptor instead.
func (*PaymentPrecheckPayload) Descriptor() ([]byte, []int) {
	return file_lib_protos_v1_payment_payment_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentPrecheckPayload) GetAccountNo() string {
	if x != nil {
		return x.AccountNo
	}
	return ""
}

type PaymentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datetime string        `protobuf:"bytes,1,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Status   PaymentStatus `protobuf:"varint,2,opt,name=status,proto3,enum=fastcampus.payment.v1.PaymentStatus" json:"status,omitempty"`
}

func (x *PaymentInfo) Reset() {
	*x = PaymentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_protos_v1_payment_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentInfo) ProtoMessage() {}

func (x *PaymentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_lib_protos_v1_payment_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentInfo.ProtoReflect.Descriptor instead.
func (*PaymentInfo) Descriptor() ([]byte, []int) {
	return file_lib_protos_v1_payment_payment_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentInfo) GetDatetime() string {
	if x != nil {
		return x.Datetime
	}
	return ""
}

func (x *PaymentInfo) GetStatus() PaymentStatus {
	if x != nil {
		return x.Status
	}
	return PaymentStatus_UNSPECIFIED
}

type PaymentStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvoiceId      string         `protobuf:"bytes,1,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	Status         PaymentStatus  `protobuf:"varint,2,opt,name=status,proto3,enum=fastcampus.payment.v1.PaymentStatus" json:"status,omitempty"`
	Amount         int64          `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Datetime       string         `protobuf:"bytes,4,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Message        string         `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	PaymentHistory []*PaymentInfo `protobuf:"bytes,6,rep,name=payment_history,json=paymentHistory,proto3" json:"payment_history,omitempty"`
}

func (x *PaymentStatusResponse) Reset() {
	*x = PaymentStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_protos_v1_payment_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentStatusResponse) ProtoMessage() {}

func (x *PaymentStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lib_protos_v1_payment_payment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentStatusResponse.ProtoReflect.Descriptor instead.
func (*PaymentStatusResponse) Descriptor() ([]byte, []int) {
	return file_lib_protos_v1_payment_payment_proto_rawDescGZIP(), []int{3}
}

func (x *PaymentStatusResponse) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

func (x *PaymentStatusResponse) GetStatus() PaymentStatus {
	if x != nil {
		return x.Status
	}
	return PaymentStatus_UNSPECIFIED
}

func (x *PaymentStatusResponse) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PaymentStatusResponse) GetDatetime() string {
	if x != nil {
		return x.Datetime
	}
	return ""
}

func (x *PaymentStatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PaymentStatusResponse) GetPaymentHistory() []*PaymentInfo {
	if x != nil {
		return x.PaymentHistory
	}
	return nil
}

type SubmitPaymentPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MerchantId string `protobuf:"bytes,1,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	Amount     int64  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	AccountNo  string `protobuf:"bytes,3,opt,name=account_no,json=accountNo,proto3" json:"account_no,omitempty"`
}

func (x *SubmitPaymentPayload) Reset() {
	*x = SubmitPaymentPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_protos_v1_payment_payment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitPaymentPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitPaymentPayload) ProtoMessage() {}

func (x *SubmitPaymentPayload) ProtoReflect() protoreflect.Message {
	mi := &file_lib_protos_v1_payment_payment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitPaymentPayload.ProtoReflect.Descriptor instead.
func (*SubmitPaymentPayload) Descriptor() ([]byte, []int) {
	return file_lib_protos_v1_payment_payment_proto_rawDescGZIP(), []int{4}
}

func (x *SubmitPaymentPayload) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *SubmitPaymentPayload) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *SubmitPaymentPayload) GetAccountNo() string {
	if x != nil {
		return x.AccountNo
	}
	return ""
}

type SubmitPaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvoiceId string        `protobuf:"bytes,1,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	Status    PaymentStatus `protobuf:"varint,2,opt,name=status,proto3,enum=fastcampus.payment.v1.PaymentStatus" json:"status,omitempty"`
	Message   string        `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SubmitPaymentResponse) Reset() {
	*x = SubmitPaymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_protos_v1_payment_payment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitPaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitPaymentResponse) ProtoMessage() {}

func (x *SubmitPaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lib_protos_v1_payment_payment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitPaymentResponse.ProtoReflect.Descriptor instead.
func (*SubmitPaymentResponse) Descriptor() ([]byte, []int) {
	return file_lib_protos_v1_payment_payment_proto_rawDescGZIP(), []int{5}
}

func (x *SubmitPaymentResponse) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

func (x *SubmitPaymentResponse) GetStatus() PaymentStatus {
	if x != nil {
		return x.Status
	}
	return PaymentStatus_UNSPECIFIED
}

func (x *SubmitPaymentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email         string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Fullname      string `protobuf:"bytes,2,opt,name=fullname,proto3" json:"fullname,omitempty"`
	AccountNumber string `protobuf:"bytes,3,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	Status        string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UserInfoResponse) Reset() {
	*x = UserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_protos_v1_payment_payment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResponse) ProtoMessage() {}

func (x *UserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lib_protos_v1_payment_payment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResponse.ProtoReflect.Descriptor instead.
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return file_lib_protos_v1_payment_payment_proto_rawDescGZIP(), []int{6}
}

func (x *UserInfoResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfoResponse) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *UserInfoResponse) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

func (x *UserInfoResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type PaymentPrecheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance int64             `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
	Info    *UserInfoResponse `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *PaymentPrecheckResponse) Reset() {
	*x = PaymentPrecheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_protos_v1_payment_payment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentPrecheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentPrecheckResponse) ProtoMessage() {}

func (x *PaymentPrecheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lib_protos_v1_payment_payment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentPrecheckResponse.ProtoReflect.Descriptor instead.
func (*PaymentPrecheckResponse) Descriptor() ([]byte, []int) {
	return file_lib_protos_v1_payment_payment_proto_rawDescGZIP(), []int{7}
}

func (x *PaymentPrecheckResponse) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *PaymentPrecheckResponse) GetInfo() *UserInfoResponse {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_lib_protos_v1_payment_payment_proto protoreflect.FileDescriptor

var file_lib_protos_v1_payment_payment_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x66, 0x61, 0x73, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x75,
	0x73, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x27, 0x6c, 0x69,
	0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x35, 0x0a, 0x14, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x16, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x65, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x6f,
	0x22, 0x67, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x66, 0x61,
	0x73, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x8f, 0x02, 0x0a, 0x15, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x24, 0x2e, 0x66, 0x61, 0x73, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4b,
	0x0a, 0x0f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x66, 0x61, 0x73, 0x74, 0x63, 0x61,
	0x6d, 0x70, 0x75, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x6e, 0x0a, 0x14, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x15,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x66, 0x61, 0x73, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x75,
	0x73, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x83, 0x01, 0x0a,
	0x10, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x70, 0x0a, 0x17, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x65,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x66, 0x61, 0x73, 0x74, 0x63, 0x61, 0x6d, 0x70,
	0x75, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x2a, 0x4f, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x50, 0x41, 0x49, 0x44,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12,
	0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x50,
	0x41, 0x49, 0x44, 0x10, 0x04, 0x32, 0x81, 0x05, 0x0a, 0x0c, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc8, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x2e, 0x66, 0x61,
	0x73, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x2c, 0x2e, 0x66, 0x61, 0x73, 0x74, 0x63,
	0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x59, 0x92, 0x41, 0x3c, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x47, 0x65, 0x74, 0x20, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x20, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4a, 0x0b, 0x0a, 0x03, 0x32, 0x30,
	0x30, 0x12, 0x04, 0x0a, 0x02, 0x4f, 0x4b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0xe3, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x65, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2d, 0x2e, 0x66,
	0x61, 0x73, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x65, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x2e, 0x2e, 0x66, 0x61,
	0x73, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x65, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6a, 0x92, 0x41, 0x4b,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x47, 0x65, 0x74, 0x20, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x50, 0x72, 0x65, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x20,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x65, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x4a, 0x0b,
	0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x04, 0x0a, 0x02, 0x4f, 0x4b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70,
	0x72, 0x65, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12, 0xbf, 0x01, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x66, 0x61, 0x73, 0x74,
	0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x2c, 0x2e, 0x66, 0x61, 0x73, 0x74, 0x63, 0x61, 0x6d,
	0x70, 0x75, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0x92, 0x41, 0x33, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79,
	0x12, 0x0e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x20, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2a, 0x0d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4a,
	0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x04, 0x0a, 0x02, 0x4f, 0x4b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x42, 0xf6, 0x01, 0x92, 0x41, 0xa3, 0x01,
	0x12, 0x7a, 0x0a, 0x26, 0x46, 0x61, 0x73, 0x74, 0x20, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x20,
	0x50, 0x61, 0x79, 0x20, 0x28, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x41, 0x50, 0x49,
	0x20, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x29, 0x22, 0x4b, 0x0a, 0x13, 0x46, 0x61,
	0x73, 0x74, 0x20, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x20, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x1e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x74, 0x69, 0x6b, 0x74, 0x65, 0x72, 0x61, 0x6e,
	0x67, 0x1a, 0x14, 0x61, 0x7a, 0x77, 0x61, 0x72, 0x2e, 0x6e, 0x72, 0x73, 0x74, 0x40, 0x67, 0x6d,
	0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01, 0x01, 0x32,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x69, 0x74, 0x69, 0x6b, 0x74, 0x65, 0x72, 0x61, 0x6e, 0x67, 0x2f, 0x68, 0x65, 0x78, 0x61,
	0x67, 0x6f, 0x6e, 0x61, 0x6c, 0x2d, 0x66, 0x61, 0x73, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73,
	0x2d, 0x70, 0x61, 0x79, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x3b, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lib_protos_v1_payment_payment_proto_rawDescOnce sync.Once
	file_lib_protos_v1_payment_payment_proto_rawDescData = file_lib_protos_v1_payment_payment_proto_rawDesc
)

func file_lib_protos_v1_payment_payment_proto_rawDescGZIP() []byte {
	file_lib_protos_v1_payment_payment_proto_rawDescOnce.Do(func() {
		file_lib_protos_v1_payment_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_lib_protos_v1_payment_payment_proto_rawDescData)
	})
	return file_lib_protos_v1_payment_payment_proto_rawDescData
}

var file_lib_protos_v1_payment_payment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_lib_protos_v1_payment_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_lib_protos_v1_payment_payment_proto_goTypes = []any{
	(PaymentStatus)(0),              // 0: fastcampus.payment.v1.PaymentStatus
	(*PaymentStatusPayload)(nil),    // 1: fastcampus.payment.v1.PaymentStatusPayload
	(*PaymentPrecheckPayload)(nil),  // 2: fastcampus.payment.v1.PaymentPrecheckPayload
	(*PaymentInfo)(nil),             // 3: fastcampus.payment.v1.PaymentInfo
	(*PaymentStatusResponse)(nil),   // 4: fastcampus.payment.v1.PaymentStatusResponse
	(*SubmitPaymentPayload)(nil),    // 5: fastcampus.payment.v1.SubmitPaymentPayload
	(*SubmitPaymentResponse)(nil),   // 6: fastcampus.payment.v1.SubmitPaymentResponse
	(*UserInfoResponse)(nil),        // 7: fastcampus.payment.v1.UserInfoResponse
	(*PaymentPrecheckResponse)(nil), // 8: fastcampus.payment.v1.PaymentPrecheckResponse
}
var file_lib_protos_v1_payment_payment_proto_depIdxs = []int32{
	0, // 0: fastcampus.payment.v1.PaymentInfo.status:type_name -> fastcampus.payment.v1.PaymentStatus
	0, // 1: fastcampus.payment.v1.PaymentStatusResponse.status:type_name -> fastcampus.payment.v1.PaymentStatus
	3, // 2: fastcampus.payment.v1.PaymentStatusResponse.payment_history:type_name -> fastcampus.payment.v1.PaymentInfo
	0, // 3: fastcampus.payment.v1.SubmitPaymentResponse.status:type_name -> fastcampus.payment.v1.PaymentStatus
	7, // 4: fastcampus.payment.v1.PaymentPrecheckResponse.info:type_name -> fastcampus.payment.v1.UserInfoResponse
	1, // 5: fastcampus.payment.v1.MoneyService.GetPaymentStatus:input_type -> fastcampus.payment.v1.PaymentStatusPayload
	2, // 6: fastcampus.payment.v1.MoneyService.GetPaymentPrecheckInfo:input_type -> fastcampus.payment.v1.PaymentPrecheckPayload
	5, // 7: fastcampus.payment.v1.MoneyService.SubmitPayment:input_type -> fastcampus.payment.v1.SubmitPaymentPayload
	4, // 8: fastcampus.payment.v1.MoneyService.GetPaymentStatus:output_type -> fastcampus.payment.v1.PaymentStatusResponse
	8, // 9: fastcampus.payment.v1.MoneyService.GetPaymentPrecheckInfo:output_type -> fastcampus.payment.v1.PaymentPrecheckResponse
	6, // 10: fastcampus.payment.v1.MoneyService.SubmitPayment:output_type -> fastcampus.payment.v1.SubmitPaymentResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_lib_protos_v1_payment_payment_proto_init() }
func file_lib_protos_v1_payment_payment_proto_init() {
	if File_lib_protos_v1_payment_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lib_protos_v1_payment_payment_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PaymentStatusPayload); i {
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
		file_lib_protos_v1_payment_payment_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PaymentPrecheckPayload); i {
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
		file_lib_protos_v1_payment_payment_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PaymentInfo); i {
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
		file_lib_protos_v1_payment_payment_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PaymentStatusResponse); i {
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
		file_lib_protos_v1_payment_payment_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SubmitPaymentPayload); i {
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
		file_lib_protos_v1_payment_payment_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SubmitPaymentResponse); i {
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
		file_lib_protos_v1_payment_payment_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UserInfoResponse); i {
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
		file_lib_protos_v1_payment_payment_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*PaymentPrecheckResponse); i {
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
			RawDescriptor: file_lib_protos_v1_payment_payment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lib_protos_v1_payment_payment_proto_goTypes,
		DependencyIndexes: file_lib_protos_v1_payment_payment_proto_depIdxs,
		EnumInfos:         file_lib_protos_v1_payment_payment_proto_enumTypes,
		MessageInfos:      file_lib_protos_v1_payment_payment_proto_msgTypes,
	}.Build()
	File_lib_protos_v1_payment_payment_proto = out.File
	file_lib_protos_v1_payment_payment_proto_rawDesc = nil
	file_lib_protos_v1_payment_payment_proto_goTypes = nil
	file_lib_protos_v1_payment_payment_proto_depIdxs = nil
}