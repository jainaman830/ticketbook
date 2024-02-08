// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.26.0--rc1
// source: ticket.proto

package generated

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

// Message for representing a user.
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID     string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	FirstName  string `protobuf:"bytes,2,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName   string `protobuf:"bytes,3,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Email      string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	CreatedOn  string `protobuf:"bytes,5,opt,name=CreatedOn,proto3" json:"CreatedOn,omitempty"`
	ModifiedOn string `protobuf:"bytes,6,opt,name=ModifiedOn,proto3" json:"ModifiedOn,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetCreatedOn() string {
	if x != nil {
		return x.CreatedOn
	}
	return ""
}

func (x *User) GetModifiedOn() string {
	if x != nil {
		return x.ModifiedOn
	}
	return ""
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,2,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateUserRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// Message for representing a train ticket purchase.
type Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From       string  `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To         string  `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	User       *User   `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	PricePaid  float32 `protobuf:"fixed32,4,opt,name=price_paid,json=pricePaid,proto3" json:"price_paid,omitempty"`
	Section    string  `protobuf:"bytes,5,opt,name=section,proto3" json:"section,omitempty"`
	SeatNumber string  `protobuf:"bytes,6,opt,name=seat_number,json=seatNumber,proto3" json:"seat_number,omitempty"`
	CreatedOn  string  `protobuf:"bytes,7,opt,name=CreatedOn,proto3" json:"CreatedOn,omitempty"`
	ModifiedOn string  `protobuf:"bytes,8,opt,name=ModifiedOn,proto3" json:"ModifiedOn,omitempty"`
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

func (x *Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticket.ProtoReflect.Descriptor instead.
func (*Ticket) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{2}
}

func (x *Ticket) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Ticket) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Ticket) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Ticket) GetPricePaid() float32 {
	if x != nil {
		return x.PricePaid
	}
	return 0
}

func (x *Ticket) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

func (x *Ticket) GetSeatNumber() string {
	if x != nil {
		return x.SeatNumber
	}
	return ""
}

func (x *Ticket) GetCreatedOn() string {
	if x != nil {
		return x.CreatedOn
	}
	return ""
}

func (x *Ticket) GetModifiedOn() string {
	if x != nil {
		return x.ModifiedOn
	}
	return ""
}

type Section struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section        string `protobuf:"bytes,1,opt,name=Section,proto3" json:"Section,omitempty"`
	TotalSeats     int32  `protobuf:"varint,2,opt,name=TotalSeats,proto3" json:"TotalSeats,omitempty"`
	AvailableSeats int32  `protobuf:"varint,3,opt,name=AvailableSeats,proto3" json:"AvailableSeats,omitempty"`
	CreatedOn      string `protobuf:"bytes,4,opt,name=CreatedOn,proto3" json:"CreatedOn,omitempty"`
	ModifiedOn     string `protobuf:"bytes,5,opt,name=ModifiedOn,proto3" json:"ModifiedOn,omitempty"`
}

func (x *Section) Reset() {
	*x = Section{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Section) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Section) ProtoMessage() {}

func (x *Section) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Section.ProtoReflect.Descriptor instead.
func (*Section) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{3}
}

func (x *Section) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

func (x *Section) GetTotalSeats() int32 {
	if x != nil {
		return x.TotalSeats
	}
	return 0
}

func (x *Section) GetAvailableSeats() int32 {
	if x != nil {
		return x.AvailableSeats
	}
	return 0
}

func (x *Section) GetCreatedOn() string {
	if x != nil {
		return x.CreatedOn
	}
	return ""
}

func (x *Section) GetModifiedOn() string {
	if x != nil {
		return x.ModifiedOn
	}
	return ""
}

type AllSections struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sections []*Section `protobuf:"bytes,1,rep,name=sections,proto3" json:"sections,omitempty"`
}

func (x *AllSections) Reset() {
	*x = AllSections{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllSections) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllSections) ProtoMessage() {}

func (x *AllSections) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllSections.ProtoReflect.Descriptor instead.
func (*AllSections) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{4}
}

func (x *AllSections) GetSections() []*Section {
	if x != nil {
		return x.Sections
	}
	return nil
}

type AllUsers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *AllUsers) Reset() {
	*x = AllUsers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllUsers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllUsers) ProtoMessage() {}

func (x *AllUsers) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllUsers.ProtoReflect.Descriptor instead.
func (*AllUsers) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{5}
}

func (x *AllUsers) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

// Message for representing seat allocation.
type SeatAllocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tickets []*Ticket `protobuf:"bytes,1,rep,name=tickets,proto3" json:"tickets,omitempty"`
}

func (x *SeatAllocation) Reset() {
	*x = SeatAllocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeatAllocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeatAllocation) ProtoMessage() {}

func (x *SeatAllocation) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeatAllocation.ProtoReflect.Descriptor instead.
func (*SeatAllocation) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{6}
}

func (x *SeatAllocation) GetTickets() []*Ticket {
	if x != nil {
		return x.Tickets
	}
	return nil
}

// Boolean message type
type Bool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Bool) Reset() {
	*x = Bool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bool) ProtoMessage() {}

func (x *Bool) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bool.ProtoReflect.Descriptor instead.
func (*Bool) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{7}
}

func (x *Bool) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type UseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *UseRequest) Reset() {
	*x = UseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UseRequest) ProtoMessage() {}

func (x *UseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UseRequest.ProtoReflect.Descriptor instead.
func (*UseRequest) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{8}
}

func (x *UseRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type SectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section string `protobuf:"bytes,1,opt,name=Section,proto3" json:"Section,omitempty"`
}

func (x *SectionRequest) Reset() {
	*x = SectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SectionRequest) ProtoMessage() {}

func (x *SectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SectionRequest.ProtoReflect.Descriptor instead.
func (*SectionRequest) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{9}
}

func (x *SectionRequest) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

// Empty response message.
type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{10}
}

var File_ticket_proto protoreflect.FileDescriptor

var file_ticket_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x22,
	0xac, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4f, 0x6e, 0x22, 0x63,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0xef, 0x01, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x29, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x61, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x74, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61,
	0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x4f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x4f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x4f, 0x6e, 0x22, 0xa9, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x61, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x61, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65,
	0x61, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f,
	0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4f,
	0x6e, 0x22, 0x43, 0x0a, 0x0b, 0x41, 0x6c, 0x6c, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x34, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x37, 0x0a, 0x08, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x2b, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22,
	0x43, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x31, 0x0a, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x07, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x22, 0x1c, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x24, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x97, 0x07, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x43, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a,
	0x0c, 0x56, 0x69, 0x65, 0x77, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x41, 0x6c, 0x6c, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x48, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x44, 0x0a, 0x0e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x15,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x1e, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0e, 0x50, 0x75,
	0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x3d,
	0x0a, 0x0b, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x15, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x1a, 0x17, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x4f, 0x0a,
	0x12, 0x56, 0x69, 0x65, 0x77, 0x53, 0x65, 0x61, 0x74, 0x73, 0x42, 0x79, 0x53, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1f, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x53, 0x65, 0x61, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f,
	0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12,
	0x17, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x12,
	0x3e, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x12, 0x17, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x42,
	0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ticket_proto_rawDescOnce sync.Once
	file_ticket_proto_rawDescData = file_ticket_proto_rawDesc
)

func file_ticket_proto_rawDescGZIP() []byte {
	file_ticket_proto_rawDescOnce.Do(func() {
		file_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_ticket_proto_rawDescData)
	})
	return file_ticket_proto_rawDescData
}

var file_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_ticket_proto_goTypes = []interface{}{
	(*User)(nil),              // 0: train_ticketing.User
	(*CreateUserRequest)(nil), // 1: train_ticketing.CreateUserRequest
	(*Ticket)(nil),            // 2: train_ticketing.Ticket
	(*Section)(nil),           // 3: train_ticketing.Section
	(*AllSections)(nil),       // 4: train_ticketing.AllSections
	(*AllUsers)(nil),          // 5: train_ticketing.AllUsers
	(*SeatAllocation)(nil),    // 6: train_ticketing.SeatAllocation
	(*Bool)(nil),              // 7: train_ticketing.Bool
	(*UseRequest)(nil),        // 8: train_ticketing.UseRequest
	(*SectionRequest)(nil),    // 9: train_ticketing.SectionRequest
	(*EmptyResponse)(nil),     // 10: train_ticketing.EmptyResponse
}
var file_ticket_proto_depIdxs = []int32{
	0,  // 0: train_ticketing.Ticket.user:type_name -> train_ticketing.User
	3,  // 1: train_ticketing.AllSections.sections:type_name -> train_ticketing.Section
	0,  // 2: train_ticketing.AllUsers.users:type_name -> train_ticketing.User
	2,  // 3: train_ticketing.SeatAllocation.tickets:type_name -> train_ticketing.Ticket
	3,  // 4: train_ticketing.TrainTicketing.CreateSection:input_type -> train_ticketing.Section
	9,  // 5: train_ticketing.TrainTicketing.ViewSections:input_type -> train_ticketing.SectionRequest
	9,  // 6: train_ticketing.TrainTicketing.DeleteSections:input_type -> train_ticketing.SectionRequest
	3,  // 7: train_ticketing.TrainTicketing.ModifySections:input_type -> train_ticketing.Section
	1,  // 8: train_ticketing.TrainTicketing.CreateUser:input_type -> train_ticketing.CreateUserRequest
	8,  // 9: train_ticketing.TrainTicketing.GetUsers:input_type -> train_ticketing.UseRequest
	0,  // 10: train_ticketing.TrainTicketing.ModifyUser:input_type -> train_ticketing.User
	0,  // 11: train_ticketing.TrainTicketing.RemoveUser:input_type -> train_ticketing.User
	2,  // 12: train_ticketing.TrainTicketing.PurchaseTicket:input_type -> train_ticketing.Ticket
	0,  // 13: train_ticketing.TrainTicketing.ViewReceipt:input_type -> train_ticketing.User
	3,  // 14: train_ticketing.TrainTicketing.ViewSeatsBySection:input_type -> train_ticketing.Section
	2,  // 15: train_ticketing.TrainTicketing.CancelReceipt:input_type -> train_ticketing.Ticket
	2,  // 16: train_ticketing.TrainTicketing.ModifySeat:input_type -> train_ticketing.Ticket
	3,  // 17: train_ticketing.TrainTicketing.CreateSection:output_type -> train_ticketing.Section
	4,  // 18: train_ticketing.TrainTicketing.ViewSections:output_type -> train_ticketing.AllSections
	7,  // 19: train_ticketing.TrainTicketing.DeleteSections:output_type -> train_ticketing.Bool
	3,  // 20: train_ticketing.TrainTicketing.ModifySections:output_type -> train_ticketing.Section
	0,  // 21: train_ticketing.TrainTicketing.CreateUser:output_type -> train_ticketing.User
	5,  // 22: train_ticketing.TrainTicketing.GetUsers:output_type -> train_ticketing.AllUsers
	0,  // 23: train_ticketing.TrainTicketing.ModifyUser:output_type -> train_ticketing.User
	10, // 24: train_ticketing.TrainTicketing.RemoveUser:output_type -> train_ticketing.EmptyResponse
	2,  // 25: train_ticketing.TrainTicketing.PurchaseTicket:output_type -> train_ticketing.Ticket
	2,  // 26: train_ticketing.TrainTicketing.ViewReceipt:output_type -> train_ticketing.Ticket
	6,  // 27: train_ticketing.TrainTicketing.ViewSeatsBySection:output_type -> train_ticketing.SeatAllocation
	7,  // 28: train_ticketing.TrainTicketing.CancelReceipt:output_type -> train_ticketing.Bool
	2,  // 29: train_ticketing.TrainTicketing.ModifySeat:output_type -> train_ticketing.Ticket
	17, // [17:30] is the sub-list for method output_type
	4,  // [4:17] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_ticket_proto_init() }
func file_ticket_proto_init() {
	if File_ticket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ticket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_ticket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_ticket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ticket); i {
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
		file_ticket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Section); i {
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
		file_ticket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllSections); i {
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
		file_ticket_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllUsers); i {
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
		file_ticket_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeatAllocation); i {
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
		file_ticket_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bool); i {
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
		file_ticket_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UseRequest); i {
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
		file_ticket_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SectionRequest); i {
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
		file_ticket_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
			RawDescriptor: file_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ticket_proto_goTypes,
		DependencyIndexes: file_ticket_proto_depIdxs,
		MessageInfos:      file_ticket_proto_msgTypes,
	}.Build()
	File_ticket_proto = out.File
	file_ticket_proto_rawDesc = nil
	file_ticket_proto_goTypes = nil
	file_ticket_proto_depIdxs = nil
}