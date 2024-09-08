// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: example/bookstore/v1/bookstore.proto

package bookstore

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A Book resource.
type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Field for path.
	Path string `protobuf:"bytes,10000,opt,name=path,proto3" json:"path,omitempty"`
	// Field for id.
	Id string `protobuf:"bytes,10001,opt,name=id,proto3" json:"id,omitempty"`
	// Field for edition.
	Edition int32 `protobuf:"varint,4,opt,name=edition,proto3" json:"edition,omitempty"`
	// Field for isbn.
	Isbn []string `protobuf:"bytes,1,rep,name=isbn,proto3" json:"isbn,omitempty"`
	// Field for price.
	Price float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	// Field for published.
	Published bool `protobuf:"varint,3,opt,name=published,proto3" json:"published,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_bookstore_v1_bookstore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_example_bookstore_v1_bookstore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_example_bookstore_v1_bookstore_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Book) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Book) GetEdition() int32 {
	if x != nil {
		return x.Edition
	}
	return 0
}

func (x *Book) GetIsbn() []string {
	if x != nil {
		return x.Isbn
	}
	return nil
}

func (x *Book) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Book) GetPublished() bool {
	if x != nil {
		return x.Published
	}
	return false
}

// A Create request for a  Book resource.
type CreateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A field for the parent of Book
	Parent string `protobuf:"bytes,10013,opt,name=parent,proto3" json:"parent,omitempty"`
	// An id that uniquely identifies the resource within the collection%!(EXTRA string=Book)
	Id string `protobuf:"bytes,10014,opt,name=id,proto3" json:"id,omitempty"`
	// The resource to perform the operation on.
	Book *Book `protobuf:"bytes,10015,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *CreateBookRequest) Reset() {
	*x = CreateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_bookstore_v1_bookstore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookRequest) ProtoMessage() {}

func (x *CreateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_bookstore_v1_bookstore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookRequest.ProtoReflect.Descriptor instead.
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return file_example_bookstore_v1_bookstore_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBookRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateBookRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateBookRequest) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

// Request message for the GetBook method
type GetBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The globally unique identifier for the resource
	Path string `protobuf:"bytes,10018,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *GetBookRequest) Reset() {
	*x = GetBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_bookstore_v1_bookstore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookRequest) ProtoMessage() {}

func (x *GetBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_bookstore_v1_bookstore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookRequest.ProtoReflect.Descriptor instead.
func (*GetBookRequest) Descriptor() ([]byte, []int) {
	return file_example_bookstore_v1_bookstore_proto_rawDescGZIP(), []int{2}
}

func (x *GetBookRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// Request message for the UpdateBook method
type UpdateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The globally unique identifier for the resource
	Path string `protobuf:"bytes,10018,opt,name=path,proto3" json:"path,omitempty"`
	// The resource to perform the operation on.
	Book *Book `protobuf:"bytes,10015,opt,name=book,proto3" json:"book,omitempty"`
	// The update mask for the resource
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,10012,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateBookRequest) Reset() {
	*x = UpdateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_bookstore_v1_bookstore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookRequest) ProtoMessage() {}

func (x *UpdateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_bookstore_v1_bookstore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookRequest.ProtoReflect.Descriptor instead.
func (*UpdateBookRequest) Descriptor() ([]byte, []int) {
	return file_example_bookstore_v1_bookstore_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateBookRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *UpdateBookRequest) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

func (x *UpdateBookRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// Request message for the DeleteBook method
type DeleteBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The globally unique identifier for the resource
	Path string `protobuf:"bytes,10018,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *DeleteBookRequest) Reset() {
	*x = DeleteBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_bookstore_v1_bookstore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookRequest) ProtoMessage() {}

func (x *DeleteBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_bookstore_v1_bookstore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookRequest.ProtoReflect.Descriptor instead.
func (*DeleteBookRequest) Descriptor() ([]byte, []int) {
	return file_example_bookstore_v1_bookstore_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteBookRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// Request message for the ListBook method
type ListBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A field for the parent of Book
	Parent string `protobuf:"bytes,10013,opt,name=parent,proto3" json:"parent,omitempty"`
	// The page token indicating the starting point of the page
	PageToken string `protobuf:"bytes,10010,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// The maximum number of resources to return in a single page.
	MaxPageSize int32 `protobuf:"varint,10017,opt,name=max_page_size,json=maxPageSize,proto3" json:"max_page_size,omitempty"`
}

func (x *ListBookRequest) Reset() {
	*x = ListBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_bookstore_v1_bookstore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBookRequest) ProtoMessage() {}

func (x *ListBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_bookstore_v1_bookstore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBookRequest.ProtoReflect.Descriptor instead.
func (*ListBookRequest) Descriptor() ([]byte, []int) {
	return file_example_bookstore_v1_bookstore_proto_rawDescGZIP(), []int{5}
}

func (x *ListBookRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListBookRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListBookRequest) GetMaxPageSize() int32 {
	if x != nil {
		return x.MaxPageSize
	}
	return 0
}

// Response message for the ListBook method
type ListBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of books
	Results []*Book `protobuf:"bytes,10016,rep,name=results,proto3" json:"results,omitempty"`
	// The page token indicating the ending point of this response.
	NextPageToken string `protobuf:"bytes,10011,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListBookResponse) Reset() {
	*x = ListBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_bookstore_v1_bookstore_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBookResponse) ProtoMessage() {}

func (x *ListBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_bookstore_v1_bookstore_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBookResponse.ProtoReflect.Descriptor instead.
func (*ListBookResponse) Descriptor() ([]byte, []int) {
	return file_example_bookstore_v1_bookstore_proto_rawDescGZIP(), []int{6}
}

func (x *ListBookResponse) GetResults() []*Book {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ListBookResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Request message for the ApplyBook method
type ApplyBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The globally unique identifier for the resource
	Path string `protobuf:"bytes,10018,opt,name=path,proto3" json:"path,omitempty"`
	// The resource to perform the operation on.
	Book *Book `protobuf:"bytes,10015,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *ApplyBookRequest) Reset() {
	*x = ApplyBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_bookstore_v1_bookstore_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyBookRequest) ProtoMessage() {}

func (x *ApplyBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_bookstore_v1_bookstore_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyBookRequest.ProtoReflect.Descriptor instead.
func (*ApplyBookRequest) Descriptor() ([]byte, []int) {
	return file_example_bookstore_v1_bookstore_proto_rawDescGZIP(), []int{7}
}

func (x *ApplyBookRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ApplyBookRequest) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

var File_example_bookstore_v1_bookstore_proto protoreflect.FileDescriptor

var file_example_bookstore_v1_bookstore_proto_rawDesc = []byte{
	0x0a, 0x24, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d,
	0x01, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x13, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x90, 0x4e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x0f, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x91, 0x4e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x69, 0x73, 0x62, 0x6e,
	0x12, 0x19, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x22, 0x7b,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x9d, 0x4e,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x00, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x9e, 0x4e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x9f, 0x4e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x49, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0xa2, 0x4e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xe0, 0x41,
	0x02, 0xfa, 0x41, 0x1c, 0x0a, 0x1a, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0xc0, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0xa2, 0x4e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xe0, 0x41, 0x02,
	0xfa, 0x41, 0x1c, 0x0a, 0x1a, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x34, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x9f, 0x4e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x3c, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x9c, 0x4e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x4c, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0xa2, 0x4e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xe0,
	0x41, 0x02, 0xfa, 0x41, 0x1c, 0x0a, 0x1a, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x77, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x9d, 0x4e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x02,
	0xfa, 0x41, 0x00, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x9a, 0x4e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6d,
	0x61, 0x78, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0xa1, 0x4e, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x22, 0x72, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0xa0, 0x4e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x9b,
	0x4e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x81, 0x01, 0x0a, 0x10, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0xa2, 0x4e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x1c,
	0x0a, 0x1a, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x34, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x9f, 0x4e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x32, 0xdf, 0x05, 0x0a, 0x09, 0x42, 0x6f, 0x6f,
	0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x7e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x27, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x2b, 0xda, 0x41, 0x0b, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x2c, 0x62, 0x6f, 0x6f, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a,
	0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x0f, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x7d, 0x12, 0x6b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x12, 0x24, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x22, 0x1e, 0xda, 0x41, 0x04, 0x70, 0x61, 0x74, 0x68, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x12, 0x0f, 0x2f, 0x7b, 0x70, 0x61, 0x74, 0x68, 0x3d, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x2f, 0x2a, 0x7d, 0x12, 0x88, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x12, 0x27, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x35, 0xda, 0x41, 0x10, 0x62, 0x6f, 0x6f, 0x6b,
	0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1c, 0x3a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x32, 0x14, 0x2f, 0x7b, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x70, 0x61, 0x74, 0x68, 0x3d, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x6d,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x27, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1e, 0xda,
	0x41, 0x04, 0x70, 0x61, 0x74, 0x68, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x2a, 0x0f, 0x2f, 0x7b,
	0x70, 0x61, 0x74, 0x68, 0x3d, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x7b, 0x0a,
	0x08, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x25, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x7b, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x3d, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x7d, 0x12, 0x6e, 0x0a, 0x09, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x26, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x1d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x17, 0x3a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x1a, 0x0f, 0x2f, 0x7b, 0x70, 0x61, 0x74,
	0x68, 0x3d, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x2a, 0x7d, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_bookstore_v1_bookstore_proto_rawDescOnce sync.Once
	file_example_bookstore_v1_bookstore_proto_rawDescData = file_example_bookstore_v1_bookstore_proto_rawDesc
)

func file_example_bookstore_v1_bookstore_proto_rawDescGZIP() []byte {
	file_example_bookstore_v1_bookstore_proto_rawDescOnce.Do(func() {
		file_example_bookstore_v1_bookstore_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_bookstore_v1_bookstore_proto_rawDescData)
	})
	return file_example_bookstore_v1_bookstore_proto_rawDescData
}

var file_example_bookstore_v1_bookstore_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_example_bookstore_v1_bookstore_proto_goTypes = []interface{}{
	(*Book)(nil),                  // 0: example.bookstore.v1.Book
	(*CreateBookRequest)(nil),     // 1: example.bookstore.v1.CreateBookRequest
	(*GetBookRequest)(nil),        // 2: example.bookstore.v1.GetBookRequest
	(*UpdateBookRequest)(nil),     // 3: example.bookstore.v1.UpdateBookRequest
	(*DeleteBookRequest)(nil),     // 4: example.bookstore.v1.DeleteBookRequest
	(*ListBookRequest)(nil),       // 5: example.bookstore.v1.ListBookRequest
	(*ListBookResponse)(nil),      // 6: example.bookstore.v1.ListBookResponse
	(*ApplyBookRequest)(nil),      // 7: example.bookstore.v1.ApplyBookRequest
	(*fieldmaskpb.FieldMask)(nil), // 8: google.protobuf.FieldMask
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_example_bookstore_v1_bookstore_proto_depIdxs = []int32{
	0,  // 0: example.bookstore.v1.CreateBookRequest.book:type_name -> example.bookstore.v1.Book
	0,  // 1: example.bookstore.v1.UpdateBookRequest.book:type_name -> example.bookstore.v1.Book
	8,  // 2: example.bookstore.v1.UpdateBookRequest.update_mask:type_name -> google.protobuf.FieldMask
	0,  // 3: example.bookstore.v1.ListBookResponse.results:type_name -> example.bookstore.v1.Book
	0,  // 4: example.bookstore.v1.ApplyBookRequest.book:type_name -> example.bookstore.v1.Book
	1,  // 5: example.bookstore.v1.Bookstore.CreateBook:input_type -> example.bookstore.v1.CreateBookRequest
	2,  // 6: example.bookstore.v1.Bookstore.GetBook:input_type -> example.bookstore.v1.GetBookRequest
	3,  // 7: example.bookstore.v1.Bookstore.UpdateBook:input_type -> example.bookstore.v1.UpdateBookRequest
	4,  // 8: example.bookstore.v1.Bookstore.DeleteBook:input_type -> example.bookstore.v1.DeleteBookRequest
	5,  // 9: example.bookstore.v1.Bookstore.ListBook:input_type -> example.bookstore.v1.ListBookRequest
	7,  // 10: example.bookstore.v1.Bookstore.ApplyBook:input_type -> example.bookstore.v1.ApplyBookRequest
	0,  // 11: example.bookstore.v1.Bookstore.CreateBook:output_type -> example.bookstore.v1.Book
	0,  // 12: example.bookstore.v1.Bookstore.GetBook:output_type -> example.bookstore.v1.Book
	0,  // 13: example.bookstore.v1.Bookstore.UpdateBook:output_type -> example.bookstore.v1.Book
	9,  // 14: example.bookstore.v1.Bookstore.DeleteBook:output_type -> google.protobuf.Empty
	6,  // 15: example.bookstore.v1.Bookstore.ListBook:output_type -> example.bookstore.v1.ListBookResponse
	0,  // 16: example.bookstore.v1.Bookstore.ApplyBook:output_type -> example.bookstore.v1.Book
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_example_bookstore_v1_bookstore_proto_init() }
func file_example_bookstore_v1_bookstore_proto_init() {
	if File_example_bookstore_v1_bookstore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_bookstore_v1_bookstore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_example_bookstore_v1_bookstore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookRequest); i {
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
		file_example_bookstore_v1_bookstore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookRequest); i {
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
		file_example_bookstore_v1_bookstore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBookRequest); i {
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
		file_example_bookstore_v1_bookstore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBookRequest); i {
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
		file_example_bookstore_v1_bookstore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBookRequest); i {
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
		file_example_bookstore_v1_bookstore_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBookResponse); i {
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
		file_example_bookstore_v1_bookstore_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyBookRequest); i {
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
			RawDescriptor: file_example_bookstore_v1_bookstore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_bookstore_v1_bookstore_proto_goTypes,
		DependencyIndexes: file_example_bookstore_v1_bookstore_proto_depIdxs,
		MessageInfos:      file_example_bookstore_v1_bookstore_proto_msgTypes,
	}.Build()
	File_example_bookstore_v1_bookstore_proto = out.File
	file_example_bookstore_v1_bookstore_proto_rawDesc = nil
	file_example_bookstore_v1_bookstore_proto_goTypes = nil
	file_example_bookstore_v1_bookstore_proto_depIdxs = nil
}
