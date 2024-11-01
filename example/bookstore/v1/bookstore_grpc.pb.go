// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: example/bookstore/v1/bookstore.proto

package bookstore

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Bookstore_CreateBook_FullMethodName        = "/example.bookstore.v1.Bookstore/CreateBook"
	Bookstore_GetBook_FullMethodName           = "/example.bookstore.v1.Bookstore/GetBook"
	Bookstore_UpdateBook_FullMethodName        = "/example.bookstore.v1.Bookstore/UpdateBook"
	Bookstore_DeleteBook_FullMethodName        = "/example.bookstore.v1.Bookstore/DeleteBook"
	Bookstore_ListBooks_FullMethodName         = "/example.bookstore.v1.Bookstore/ListBooks"
	Bookstore_ApplyBook_FullMethodName         = "/example.bookstore.v1.Bookstore/ApplyBook"
	Bookstore_CreateBookEdition_FullMethodName = "/example.bookstore.v1.Bookstore/CreateBookEdition"
	Bookstore_GetBookEdition_FullMethodName    = "/example.bookstore.v1.Bookstore/GetBookEdition"
	Bookstore_DeleteBookEdition_FullMethodName = "/example.bookstore.v1.Bookstore/DeleteBookEdition"
	Bookstore_ListBookEditions_FullMethodName  = "/example.bookstore.v1.Bookstore/ListBookEditions"
	Bookstore_CreateIsbn_FullMethodName        = "/example.bookstore.v1.Bookstore/CreateIsbn"
	Bookstore_GetIsbn_FullMethodName           = "/example.bookstore.v1.Bookstore/GetIsbn"
	Bookstore_CreatePublisher_FullMethodName   = "/example.bookstore.v1.Bookstore/CreatePublisher"
	Bookstore_GetPublisher_FullMethodName      = "/example.bookstore.v1.Bookstore/GetPublisher"
	Bookstore_UpdatePublisher_FullMethodName   = "/example.bookstore.v1.Bookstore/UpdatePublisher"
	Bookstore_DeletePublisher_FullMethodName   = "/example.bookstore.v1.Bookstore/DeletePublisher"
	Bookstore_ListPublishers_FullMethodName    = "/example.bookstore.v1.Bookstore/ListPublishers"
	Bookstore_ApplyPublisher_FullMethodName    = "/example.bookstore.v1.Bookstore/ApplyPublisher"
)

// BookstoreClient is the client API for Bookstore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookstoreClient interface {
	// An aep-compliant Create method for book.
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error)
	// An aep-compliant Get method for book.
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error)
	// An aep-compliant Update method for book.
	UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error)
	// An aep-compliant Delete method for book.
	DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// An aep-compliant List method for books.
	ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error)
	// An aep-compliant Apply method for books.
	ApplyBook(ctx context.Context, in *ApplyBookRequest, opts ...grpc.CallOption) (*Book, error)
	// An aep-compliant Create method for book-edition.
	CreateBookEdition(ctx context.Context, in *CreateBookEditionRequest, opts ...grpc.CallOption) (*BookEdition, error)
	// An aep-compliant Get method for book-edition.
	GetBookEdition(ctx context.Context, in *GetBookEditionRequest, opts ...grpc.CallOption) (*BookEdition, error)
	// An aep-compliant Delete method for book-edition.
	DeleteBookEdition(ctx context.Context, in *DeleteBookEditionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// An aep-compliant List method for book-editions.
	ListBookEditions(ctx context.Context, in *ListBookEditionsRequest, opts ...grpc.CallOption) (*ListBookEditionsResponse, error)
	// An aep-compliant Create method for isbn.
	CreateIsbn(ctx context.Context, in *CreateIsbnRequest, opts ...grpc.CallOption) (*Isbn, error)
	// An aep-compliant Get method for isbn.
	GetIsbn(ctx context.Context, in *GetIsbnRequest, opts ...grpc.CallOption) (*Isbn, error)
	// An aep-compliant Create method for publisher.
	CreatePublisher(ctx context.Context, in *CreatePublisherRequest, opts ...grpc.CallOption) (*Publisher, error)
	// An aep-compliant Get method for publisher.
	GetPublisher(ctx context.Context, in *GetPublisherRequest, opts ...grpc.CallOption) (*Publisher, error)
	// An aep-compliant Update method for publisher.
	UpdatePublisher(ctx context.Context, in *UpdatePublisherRequest, opts ...grpc.CallOption) (*Publisher, error)
	// An aep-compliant Delete method for publisher.
	DeletePublisher(ctx context.Context, in *DeletePublisherRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// An aep-compliant List method for publishers.
	ListPublishers(ctx context.Context, in *ListPublishersRequest, opts ...grpc.CallOption) (*ListPublishersResponse, error)
	// An aep-compliant Apply method for publishers.
	ApplyPublisher(ctx context.Context, in *ApplyPublisherRequest, opts ...grpc.CallOption) (*Publisher, error)
}

type bookstoreClient struct {
	cc grpc.ClientConnInterface
}

func NewBookstoreClient(cc grpc.ClientConnInterface) BookstoreClient {
	return &bookstoreClient{cc}
}

func (c *bookstoreClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, Bookstore_CreateBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, Bookstore_GetBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, Bookstore_UpdateBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Bookstore_DeleteBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error) {
	out := new(ListBooksResponse)
	err := c.cc.Invoke(ctx, Bookstore_ListBooks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) ApplyBook(ctx context.Context, in *ApplyBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, Bookstore_ApplyBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) CreateBookEdition(ctx context.Context, in *CreateBookEditionRequest, opts ...grpc.CallOption) (*BookEdition, error) {
	out := new(BookEdition)
	err := c.cc.Invoke(ctx, Bookstore_CreateBookEdition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) GetBookEdition(ctx context.Context, in *GetBookEditionRequest, opts ...grpc.CallOption) (*BookEdition, error) {
	out := new(BookEdition)
	err := c.cc.Invoke(ctx, Bookstore_GetBookEdition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) DeleteBookEdition(ctx context.Context, in *DeleteBookEditionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Bookstore_DeleteBookEdition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) ListBookEditions(ctx context.Context, in *ListBookEditionsRequest, opts ...grpc.CallOption) (*ListBookEditionsResponse, error) {
	out := new(ListBookEditionsResponse)
	err := c.cc.Invoke(ctx, Bookstore_ListBookEditions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) CreateIsbn(ctx context.Context, in *CreateIsbnRequest, opts ...grpc.CallOption) (*Isbn, error) {
	out := new(Isbn)
	err := c.cc.Invoke(ctx, Bookstore_CreateIsbn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) GetIsbn(ctx context.Context, in *GetIsbnRequest, opts ...grpc.CallOption) (*Isbn, error) {
	out := new(Isbn)
	err := c.cc.Invoke(ctx, Bookstore_GetIsbn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) CreatePublisher(ctx context.Context, in *CreatePublisherRequest, opts ...grpc.CallOption) (*Publisher, error) {
	out := new(Publisher)
	err := c.cc.Invoke(ctx, Bookstore_CreatePublisher_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) GetPublisher(ctx context.Context, in *GetPublisherRequest, opts ...grpc.CallOption) (*Publisher, error) {
	out := new(Publisher)
	err := c.cc.Invoke(ctx, Bookstore_GetPublisher_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) UpdatePublisher(ctx context.Context, in *UpdatePublisherRequest, opts ...grpc.CallOption) (*Publisher, error) {
	out := new(Publisher)
	err := c.cc.Invoke(ctx, Bookstore_UpdatePublisher_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) DeletePublisher(ctx context.Context, in *DeletePublisherRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Bookstore_DeletePublisher_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) ListPublishers(ctx context.Context, in *ListPublishersRequest, opts ...grpc.CallOption) (*ListPublishersResponse, error) {
	out := new(ListPublishersResponse)
	err := c.cc.Invoke(ctx, Bookstore_ListPublishers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) ApplyPublisher(ctx context.Context, in *ApplyPublisherRequest, opts ...grpc.CallOption) (*Publisher, error) {
	out := new(Publisher)
	err := c.cc.Invoke(ctx, Bookstore_ApplyPublisher_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookstoreServer is the server API for Bookstore service.
// All implementations must embed UnimplementedBookstoreServer
// for forward compatibility
type BookstoreServer interface {
	// An aep-compliant Create method for book.
	CreateBook(context.Context, *CreateBookRequest) (*Book, error)
	// An aep-compliant Get method for book.
	GetBook(context.Context, *GetBookRequest) (*Book, error)
	// An aep-compliant Update method for book.
	UpdateBook(context.Context, *UpdateBookRequest) (*Book, error)
	// An aep-compliant Delete method for book.
	DeleteBook(context.Context, *DeleteBookRequest) (*emptypb.Empty, error)
	// An aep-compliant List method for books.
	ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error)
	// An aep-compliant Apply method for books.
	ApplyBook(context.Context, *ApplyBookRequest) (*Book, error)
	// An aep-compliant Create method for book-edition.
	CreateBookEdition(context.Context, *CreateBookEditionRequest) (*BookEdition, error)
	// An aep-compliant Get method for book-edition.
	GetBookEdition(context.Context, *GetBookEditionRequest) (*BookEdition, error)
	// An aep-compliant Delete method for book-edition.
	DeleteBookEdition(context.Context, *DeleteBookEditionRequest) (*emptypb.Empty, error)
	// An aep-compliant List method for book-editions.
	ListBookEditions(context.Context, *ListBookEditionsRequest) (*ListBookEditionsResponse, error)
	// An aep-compliant Create method for isbn.
	CreateIsbn(context.Context, *CreateIsbnRequest) (*Isbn, error)
	// An aep-compliant Get method for isbn.
	GetIsbn(context.Context, *GetIsbnRequest) (*Isbn, error)
	// An aep-compliant Create method for publisher.
	CreatePublisher(context.Context, *CreatePublisherRequest) (*Publisher, error)
	// An aep-compliant Get method for publisher.
	GetPublisher(context.Context, *GetPublisherRequest) (*Publisher, error)
	// An aep-compliant Update method for publisher.
	UpdatePublisher(context.Context, *UpdatePublisherRequest) (*Publisher, error)
	// An aep-compliant Delete method for publisher.
	DeletePublisher(context.Context, *DeletePublisherRequest) (*emptypb.Empty, error)
	// An aep-compliant List method for publishers.
	ListPublishers(context.Context, *ListPublishersRequest) (*ListPublishersResponse, error)
	// An aep-compliant Apply method for publishers.
	ApplyPublisher(context.Context, *ApplyPublisherRequest) (*Publisher, error)
	mustEmbedUnimplementedBookstoreServer()
}

// UnimplementedBookstoreServer must be embedded to have forward compatible implementations.
type UnimplementedBookstoreServer struct {
}

func (UnimplementedBookstoreServer) CreateBook(context.Context, *CreateBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBookstoreServer) GetBook(context.Context, *GetBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBookstoreServer) UpdateBook(context.Context, *UpdateBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedBookstoreServer) DeleteBook(context.Context, *DeleteBookRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedBookstoreServer) ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBooks not implemented")
}
func (UnimplementedBookstoreServer) ApplyBook(context.Context, *ApplyBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyBook not implemented")
}
func (UnimplementedBookstoreServer) CreateBookEdition(context.Context, *CreateBookEditionRequest) (*BookEdition, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBookEdition not implemented")
}
func (UnimplementedBookstoreServer) GetBookEdition(context.Context, *GetBookEditionRequest) (*BookEdition, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookEdition not implemented")
}
func (UnimplementedBookstoreServer) DeleteBookEdition(context.Context, *DeleteBookEditionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBookEdition not implemented")
}
func (UnimplementedBookstoreServer) ListBookEditions(context.Context, *ListBookEditionsRequest) (*ListBookEditionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBookEditions not implemented")
}
func (UnimplementedBookstoreServer) CreateIsbn(context.Context, *CreateIsbnRequest) (*Isbn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIsbn not implemented")
}
func (UnimplementedBookstoreServer) GetIsbn(context.Context, *GetIsbnRequest) (*Isbn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIsbn not implemented")
}
func (UnimplementedBookstoreServer) CreatePublisher(context.Context, *CreatePublisherRequest) (*Publisher, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePublisher not implemented")
}
func (UnimplementedBookstoreServer) GetPublisher(context.Context, *GetPublisherRequest) (*Publisher, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublisher not implemented")
}
func (UnimplementedBookstoreServer) UpdatePublisher(context.Context, *UpdatePublisherRequest) (*Publisher, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePublisher not implemented")
}
func (UnimplementedBookstoreServer) DeletePublisher(context.Context, *DeletePublisherRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePublisher not implemented")
}
func (UnimplementedBookstoreServer) ListPublishers(context.Context, *ListPublishersRequest) (*ListPublishersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPublishers not implemented")
}
func (UnimplementedBookstoreServer) ApplyPublisher(context.Context, *ApplyPublisherRequest) (*Publisher, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyPublisher not implemented")
}
func (UnimplementedBookstoreServer) mustEmbedUnimplementedBookstoreServer() {}

// UnsafeBookstoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookstoreServer will
// result in compilation errors.
type UnsafeBookstoreServer interface {
	mustEmbedUnimplementedBookstoreServer()
}

func RegisterBookstoreServer(s grpc.ServiceRegistrar, srv BookstoreServer) {
	s.RegisterService(&Bookstore_ServiceDesc, srv)
}

func _Bookstore_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_CreateBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_GetBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_UpdateBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).UpdateBook(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_DeleteBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).DeleteBook(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_ListBooks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).ListBooks(ctx, req.(*ListBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_ApplyBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).ApplyBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_ApplyBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).ApplyBook(ctx, req.(*ApplyBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_CreateBookEdition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookEditionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).CreateBookEdition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_CreateBookEdition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).CreateBookEdition(ctx, req.(*CreateBookEditionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_GetBookEdition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookEditionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).GetBookEdition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_GetBookEdition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).GetBookEdition(ctx, req.(*GetBookEditionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_DeleteBookEdition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookEditionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).DeleteBookEdition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_DeleteBookEdition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).DeleteBookEdition(ctx, req.(*DeleteBookEditionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_ListBookEditions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBookEditionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).ListBookEditions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_ListBookEditions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).ListBookEditions(ctx, req.(*ListBookEditionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_CreateIsbn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIsbnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).CreateIsbn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_CreateIsbn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).CreateIsbn(ctx, req.(*CreateIsbnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_GetIsbn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIsbnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).GetIsbn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_GetIsbn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).GetIsbn(ctx, req.(*GetIsbnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_CreatePublisher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePublisherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).CreatePublisher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_CreatePublisher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).CreatePublisher(ctx, req.(*CreatePublisherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_GetPublisher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublisherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).GetPublisher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_GetPublisher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).GetPublisher(ctx, req.(*GetPublisherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_UpdatePublisher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePublisherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).UpdatePublisher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_UpdatePublisher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).UpdatePublisher(ctx, req.(*UpdatePublisherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_DeletePublisher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePublisherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).DeletePublisher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_DeletePublisher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).DeletePublisher(ctx, req.(*DeletePublisherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_ListPublishers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPublishersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).ListPublishers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_ListPublishers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).ListPublishers(ctx, req.(*ListPublishersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_ApplyPublisher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyPublisherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).ApplyPublisher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_ApplyPublisher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).ApplyPublisher(ctx, req.(*ApplyPublisherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Bookstore_ServiceDesc is the grpc.ServiceDesc for Bookstore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bookstore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.bookstore.v1.Bookstore",
	HandlerType: (*BookstoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBook",
			Handler:    _Bookstore_CreateBook_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _Bookstore_GetBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _Bookstore_UpdateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _Bookstore_DeleteBook_Handler,
		},
		{
			MethodName: "ListBooks",
			Handler:    _Bookstore_ListBooks_Handler,
		},
		{
			MethodName: "ApplyBook",
			Handler:    _Bookstore_ApplyBook_Handler,
		},
		{
			MethodName: "CreateBookEdition",
			Handler:    _Bookstore_CreateBookEdition_Handler,
		},
		{
			MethodName: "GetBookEdition",
			Handler:    _Bookstore_GetBookEdition_Handler,
		},
		{
			MethodName: "DeleteBookEdition",
			Handler:    _Bookstore_DeleteBookEdition_Handler,
		},
		{
			MethodName: "ListBookEditions",
			Handler:    _Bookstore_ListBookEditions_Handler,
		},
		{
			MethodName: "CreateIsbn",
			Handler:    _Bookstore_CreateIsbn_Handler,
		},
		{
			MethodName: "GetIsbn",
			Handler:    _Bookstore_GetIsbn_Handler,
		},
		{
			MethodName: "CreatePublisher",
			Handler:    _Bookstore_CreatePublisher_Handler,
		},
		{
			MethodName: "GetPublisher",
			Handler:    _Bookstore_GetPublisher_Handler,
		},
		{
			MethodName: "UpdatePublisher",
			Handler:    _Bookstore_UpdatePublisher_Handler,
		},
		{
			MethodName: "DeletePublisher",
			Handler:    _Bookstore_DeletePublisher_Handler,
		},
		{
			MethodName: "ListPublishers",
			Handler:    _Bookstore_ListPublishers_Handler,
		},
		{
			MethodName: "ApplyPublisher",
			Handler:    _Bookstore_ApplyPublisher_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/bookstore/v1/bookstore.proto",
}
