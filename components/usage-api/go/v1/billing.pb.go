// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: usage/v1/billing.proto

package v1

import (
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

type ReconcileInvoicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReconcileInvoicesRequest) Reset() {
	*x = ReconcileInvoicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usage_v1_billing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReconcileInvoicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconcileInvoicesRequest) ProtoMessage() {}

func (x *ReconcileInvoicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_usage_v1_billing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconcileInvoicesRequest.ProtoReflect.Descriptor instead.
func (*ReconcileInvoicesRequest) Descriptor() ([]byte, []int) {
	return file_usage_v1_billing_proto_rawDescGZIP(), []int{0}
}

type ReconcileInvoicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReconcileInvoicesResponse) Reset() {
	*x = ReconcileInvoicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usage_v1_billing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReconcileInvoicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconcileInvoicesResponse) ProtoMessage() {}

func (x *ReconcileInvoicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_usage_v1_billing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconcileInvoicesResponse.ProtoReflect.Descriptor instead.
func (*ReconcileInvoicesResponse) Descriptor() ([]byte, []int) {
	return file_usage_v1_billing_proto_rawDescGZIP(), []int{1}
}

type FinalizeInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvoiceId string `protobuf:"bytes,1,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
}

func (x *FinalizeInvoiceRequest) Reset() {
	*x = FinalizeInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usage_v1_billing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalizeInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalizeInvoiceRequest) ProtoMessage() {}

func (x *FinalizeInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_usage_v1_billing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalizeInvoiceRequest.ProtoReflect.Descriptor instead.
func (*FinalizeInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_usage_v1_billing_proto_rawDescGZIP(), []int{2}
}

func (x *FinalizeInvoiceRequest) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

type FinalizeInvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FinalizeInvoiceResponse) Reset() {
	*x = FinalizeInvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usage_v1_billing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalizeInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalizeInvoiceResponse) ProtoMessage() {}

func (x *FinalizeInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_usage_v1_billing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalizeInvoiceResponse.ProtoReflect.Descriptor instead.
func (*FinalizeInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_usage_v1_billing_proto_rawDescGZIP(), []int{3}
}

type CancelSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriptionId string `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
}

func (x *CancelSubscriptionRequest) Reset() {
	*x = CancelSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usage_v1_billing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelSubscriptionRequest) ProtoMessage() {}

func (x *CancelSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_usage_v1_billing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*CancelSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_usage_v1_billing_proto_rawDescGZIP(), []int{4}
}

func (x *CancelSubscriptionRequest) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

type CancelSubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelSubscriptionResponse) Reset() {
	*x = CancelSubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usage_v1_billing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelSubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelSubscriptionResponse) ProtoMessage() {}

func (x *CancelSubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_usage_v1_billing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelSubscriptionResponse.ProtoReflect.Descriptor instead.
func (*CancelSubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_usage_v1_billing_proto_rawDescGZIP(), []int{5}
}

type GetStripeCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Identifier:
	//
	//	*GetStripeCustomerRequest_AttributionId
	//	*GetStripeCustomerRequest_StripeCustomerId
	Identifier isGetStripeCustomerRequest_Identifier `protobuf_oneof:"identifier"`
}

func (x *GetStripeCustomerRequest) Reset() {
	*x = GetStripeCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usage_v1_billing_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStripeCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStripeCustomerRequest) ProtoMessage() {}

func (x *GetStripeCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_usage_v1_billing_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStripeCustomerRequest.ProtoReflect.Descriptor instead.
func (*GetStripeCustomerRequest) Descriptor() ([]byte, []int) {
	return file_usage_v1_billing_proto_rawDescGZIP(), []int{6}
}

func (m *GetStripeCustomerRequest) GetIdentifier() isGetStripeCustomerRequest_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (x *GetStripeCustomerRequest) GetAttributionId() string {
	if x, ok := x.GetIdentifier().(*GetStripeCustomerRequest_AttributionId); ok {
		return x.AttributionId
	}
	return ""
}

func (x *GetStripeCustomerRequest) GetStripeCustomerId() string {
	if x, ok := x.GetIdentifier().(*GetStripeCustomerRequest_StripeCustomerId); ok {
		return x.StripeCustomerId
	}
	return ""
}

type isGetStripeCustomerRequest_Identifier interface {
	isGetStripeCustomerRequest_Identifier()
}

type GetStripeCustomerRequest_AttributionId struct {
	AttributionId string `protobuf:"bytes,1,opt,name=attribution_id,json=attributionId,proto3,oneof"`
}

type GetStripeCustomerRequest_StripeCustomerId struct {
	StripeCustomerId string `protobuf:"bytes,2,opt,name=stripe_customer_id,json=stripeCustomerId,proto3,oneof"`
}

func (*GetStripeCustomerRequest_AttributionId) isGetStripeCustomerRequest_Identifier() {}

func (*GetStripeCustomerRequest_StripeCustomerId) isGetStripeCustomerRequest_Identifier() {}

type GetStripeCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer      *StripeCustomer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	AttributionId string          `protobuf:"bytes,2,opt,name=attribution_id,json=attributionId,proto3" json:"attribution_id,omitempty"`
}

func (x *GetStripeCustomerResponse) Reset() {
	*x = GetStripeCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usage_v1_billing_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStripeCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStripeCustomerResponse) ProtoMessage() {}

func (x *GetStripeCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_usage_v1_billing_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStripeCustomerResponse.ProtoReflect.Descriptor instead.
func (*GetStripeCustomerResponse) Descriptor() ([]byte, []int) {
	return file_usage_v1_billing_proto_rawDescGZIP(), []int{7}
}

func (x *GetStripeCustomerResponse) GetCustomer() *StripeCustomer {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *GetStripeCustomerResponse) GetAttributionId() string {
	if x != nil {
		return x.AttributionId
	}
	return ""
}

type StripeCustomer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StripeCustomer) Reset() {
	*x = StripeCustomer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usage_v1_billing_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StripeCustomer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StripeCustomer) ProtoMessage() {}

func (x *StripeCustomer) ProtoReflect() protoreflect.Message {
	mi := &file_usage_v1_billing_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StripeCustomer.ProtoReflect.Descriptor instead.
func (*StripeCustomer) Descriptor() ([]byte, []int) {
	return file_usage_v1_billing_proto_rawDescGZIP(), []int{8}
}

func (x *StripeCustomer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_usage_v1_billing_proto protoreflect.FileDescriptor

var file_usage_v1_billing_proto_rawDesc = []byte{
	0x0a, 0x16, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x22, 0x1a, 0x0a, 0x18, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1b,
	0x0a, 0x19, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x37, 0x0a, 0x16, 0x46,
	0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x44, 0x0a, 0x19, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x1c, 0x0a, 0x1a, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x70,
	0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x27, 0x0a, 0x0e, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x73, 0x74, 0x72,
	0x69, 0x70, 0x65, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x10, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x78, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x72, 0x69, 0x70, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x20, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x32, 0x8d, 0x03, 0x0a, 0x0e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x11, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63,
	0x69, 0x6c, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e,
	0x63, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x20, 0x2e, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x61, 0x0a, 0x12, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x70, 0x65,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x70,
	0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x69, 0x74, 0x70,
	0x6f, 0x64, 0x2f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_usage_v1_billing_proto_rawDescOnce sync.Once
	file_usage_v1_billing_proto_rawDescData = file_usage_v1_billing_proto_rawDesc
)

func file_usage_v1_billing_proto_rawDescGZIP() []byte {
	file_usage_v1_billing_proto_rawDescOnce.Do(func() {
		file_usage_v1_billing_proto_rawDescData = protoimpl.X.CompressGZIP(file_usage_v1_billing_proto_rawDescData)
	})
	return file_usage_v1_billing_proto_rawDescData
}

var file_usage_v1_billing_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_usage_v1_billing_proto_goTypes = []interface{}{
	(*ReconcileInvoicesRequest)(nil),   // 0: usage.v1.ReconcileInvoicesRequest
	(*ReconcileInvoicesResponse)(nil),  // 1: usage.v1.ReconcileInvoicesResponse
	(*FinalizeInvoiceRequest)(nil),     // 2: usage.v1.FinalizeInvoiceRequest
	(*FinalizeInvoiceResponse)(nil),    // 3: usage.v1.FinalizeInvoiceResponse
	(*CancelSubscriptionRequest)(nil),  // 4: usage.v1.CancelSubscriptionRequest
	(*CancelSubscriptionResponse)(nil), // 5: usage.v1.CancelSubscriptionResponse
	(*GetStripeCustomerRequest)(nil),   // 6: usage.v1.GetStripeCustomerRequest
	(*GetStripeCustomerResponse)(nil),  // 7: usage.v1.GetStripeCustomerResponse
	(*StripeCustomer)(nil),             // 8: usage.v1.StripeCustomer
}
var file_usage_v1_billing_proto_depIdxs = []int32{
	8, // 0: usage.v1.GetStripeCustomerResponse.customer:type_name -> usage.v1.StripeCustomer
	0, // 1: usage.v1.BillingService.ReconcileInvoices:input_type -> usage.v1.ReconcileInvoicesRequest
	2, // 2: usage.v1.BillingService.FinalizeInvoice:input_type -> usage.v1.FinalizeInvoiceRequest
	4, // 3: usage.v1.BillingService.CancelSubscription:input_type -> usage.v1.CancelSubscriptionRequest
	6, // 4: usage.v1.BillingService.GetStripeCustomer:input_type -> usage.v1.GetStripeCustomerRequest
	1, // 5: usage.v1.BillingService.ReconcileInvoices:output_type -> usage.v1.ReconcileInvoicesResponse
	3, // 6: usage.v1.BillingService.FinalizeInvoice:output_type -> usage.v1.FinalizeInvoiceResponse
	5, // 7: usage.v1.BillingService.CancelSubscription:output_type -> usage.v1.CancelSubscriptionResponse
	7, // 8: usage.v1.BillingService.GetStripeCustomer:output_type -> usage.v1.GetStripeCustomerResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_usage_v1_billing_proto_init() }
func file_usage_v1_billing_proto_init() {
	if File_usage_v1_billing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_usage_v1_billing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReconcileInvoicesRequest); i {
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
		file_usage_v1_billing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReconcileInvoicesResponse); i {
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
		file_usage_v1_billing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalizeInvoiceRequest); i {
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
		file_usage_v1_billing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalizeInvoiceResponse); i {
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
		file_usage_v1_billing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelSubscriptionRequest); i {
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
		file_usage_v1_billing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelSubscriptionResponse); i {
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
		file_usage_v1_billing_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStripeCustomerRequest); i {
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
		file_usage_v1_billing_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStripeCustomerResponse); i {
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
		file_usage_v1_billing_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StripeCustomer); i {
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
	file_usage_v1_billing_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*GetStripeCustomerRequest_AttributionId)(nil),
		(*GetStripeCustomerRequest_StripeCustomerId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_usage_v1_billing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_usage_v1_billing_proto_goTypes,
		DependencyIndexes: file_usage_v1_billing_proto_depIdxs,
		MessageInfos:      file_usage_v1_billing_proto_msgTypes,
	}.Build()
	File_usage_v1_billing_proto = out.File
	file_usage_v1_billing_proto_rawDesc = nil
	file_usage_v1_billing_proto_goTypes = nil
	file_usage_v1_billing_proto_depIdxs = nil
}
