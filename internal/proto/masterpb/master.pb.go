// Code generated by protoc-gen-go. DO NOT EDIT.
// source: master.proto

package masterpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	commonpb "github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	internalpb "github.com/zilliztech/milvus-distributed/internal/proto/internalpb"
	servicepb "github.com/zilliztech/milvus-distributed/internal/proto/servicepb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("master.proto", fileDescriptor_f9c348dec43a6705) }

var fileDescriptor_f9c348dec43a6705 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0xe5, 0x44, 0xcc, 0x86, 0x0f, 0x59, 0x6f, 0x78, 0xeb, 0xc9, 0x80, 0xb4, 0x46, 0xff, 0x80,
	0x01, 0x0e, 0x1c, 0x34, 0x31, 0x70, 0xd3, 0x18, 0xdc, 0x96, 0x0d, 0x4c, 0x6c, 0xbb, 0x75, 0x67,
	0x8a, 0x09, 0xff, 0xc2, 0x7f, 0x6c, 0xda, 0x52, 0xda, 0x15, 0x8a, 0xe8, 0xad, 0x3b, 0xf3, 0xf6,
	0xbd, 0xce, 0xbc, 0x97, 0x65, 0x8d, 0x40, 0x20, 0x49, 0x6d, 0x47, 0x5a, 0x91, 0xe2, 0x17, 0x01,
	0xf8, 0xeb, 0x18, 0xb3, 0x93, 0x9d, 0xb5, 0xba, 0x0d, 0x4f, 0x05, 0x81, 0x0a, 0xb3, 0x62, 0x97,
	0x43, 0x48, 0x52, 0x87, 0xc2, 0x9f, 0x07, 0xb8, 0xdc, 0xd6, 0x3a, 0x28, 0xf5, 0x1a, 0x3c, 0x59,
	0x94, 0x6e, 0xbf, 0xce, 0x58, 0xfd, 0x31, 0xbd, 0xcf, 0x05, 0x3b, 0x1f, 0x69, 0x29, 0x48, 0x8e,
	0x94, 0xef, 0x4b, 0x8f, 0x40, 0x85, 0xdc, 0xb6, 0x0d, 0xa5, 0x9c, 0xd3, 0xfe, 0x09, 0x9c, 0xca,
	0x8f, 0x58, 0x22, 0x75, 0x2f, 0x4d, 0xfc, 0xf6, 0x8f, 0x66, 0x24, 0x28, 0x46, 0xab, 0xc6, 0x5f,
	0x59, 0x6b, 0xac, 0x55, 0x54, 0x12, 0xb8, 0xae, 0x10, 0x30, 0x61, 0x27, 0xd2, 0xbb, 0xac, 0x39,
	0x11, 0x58, 0x62, 0xef, 0x57, 0xb0, 0x1b, 0xa8, 0x9c, 0xdc, 0x32, 0xc1, 0xdb, 0x5d, 0xd9, 0x43,
	0xa5, 0xfc, 0xa9, 0xc4, 0x48, 0x85, 0x28, 0xad, 0x1a, 0x8f, 0x19, 0x1f, 0x4b, 0xf4, 0x34, 0xb8,
	0xe5, 0x3d, 0xdd, 0x54, 0x8d, 0xb1, 0x07, 0xcd, 0xd5, 0xfa, 0x87, 0xd5, 0x0a, 0x60, 0x76, 0x35,
	0x4a, 0x3e, 0xad, 0x1a, 0x7f, 0x67, 0xed, 0xd9, 0x4a, 0x7d, 0x16, 0x6d, 0xac, 0x5c, 0x9d, 0x89,
	0xcb, 0xf5, 0xae, 0x0e, 0xeb, 0xcd, 0x48, 0x43, 0xb8, 0x7c, 0x00, 0xa4, 0xd2, 0x8c, 0x73, 0xd6,
	0xce, 0x0c, 0x7e, 0x12, 0x9a, 0x20, 0x1d, 0x70, 0x70, 0x34, 0x08, 0x3b, 0xdc, 0x89, 0x46, 0xbd,
	0xb0, 0x66, 0x62, 0x70, 0x41, 0xdf, 0x3f, 0x12, 0x83, 0xbf, 0x92, 0xbf, 0xb1, 0xc6, 0x44, 0x60,
	0xc1, 0xdd, 0xab, 0x0e, 0xc1, 0x1e, 0xf5, 0x69, 0x19, 0xd0, 0xac, 0x93, 0x1b, 0x5b, 0xc8, 0x38,
	0xbf, 0x44, 0x60, 0x4f, 0xab, 0x77, 0x58, 0x6b, 0x87, 0x33, 0x03, 0x00, 0xac, 0x95, 0x18, 0xbb,
	0xeb, 0x62, 0xe5, 0xce, 0x0c, 0xd8, 0x3f, 0xec, 0x1f, 0x0e, 0x9f, 0xef, 0x97, 0x40, 0xab, 0xd8,
	0x4d, 0x56, 0xeb, 0x6c, 0xc0, 0xf7, 0x61, 0x43, 0xd2, 0x5b, 0x39, 0x19, 0xc5, 0x60, 0x01, 0x48,
	0x1a, 0xdc, 0x98, 0xe4, 0xc2, 0xc9, 0x55, 0x9d, 0x94, 0xd7, 0xc9, 0x9e, 0xa2, 0xc8, 0x75, 0xeb,
	0xe9, 0xf9, 0xee, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x96, 0xb6, 0xf6, 0x5a, 0xb8, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MasterClient is the client API for Master service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MasterClient interface {
	//*
	// @brief This method is used to create collection
	//
	// @param CreateCollectionRequest, use to provide collection information to be created.
	//
	// @return Status
	CreateCollection(ctx context.Context, in *internalpb.CreateCollectionRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	//*
	// @brief This method is used to delete collection.
	//
	// @param DropCollectionRequest, collection name is going to be deleted.
	//
	// @return Status
	DropCollection(ctx context.Context, in *internalpb.DropCollectionRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	//*
	// @brief This method is used to test collection existence.
	//
	// @param HasCollectionRequest, collection name is going to be tested.
	//
	// @return BoolResponse
	HasCollection(ctx context.Context, in *internalpb.HasCollectionRequest, opts ...grpc.CallOption) (*servicepb.BoolResponse, error)
	//*
	// @brief This method is used to get collection schema.
	//
	// @param DescribeCollectionRequest, target collection name.
	//
	// @return CollectionSchema
	DescribeCollection(ctx context.Context, in *internalpb.DescribeCollectionRequest, opts ...grpc.CallOption) (*servicepb.CollectionDescription, error)
	//*
	// @brief This method is used to list all collections.
	//
	// @return StringListResponse, collection name list
	ShowCollections(ctx context.Context, in *internalpb.ShowCollectionRequest, opts ...grpc.CallOption) (*servicepb.StringListResponse, error)
	//*
	// @brief This method is used to create partition
	//
	// @return Status
	CreatePartition(ctx context.Context, in *internalpb.CreatePartitionRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	//*
	// @brief This method is used to drop partition
	//
	// @return Status
	DropPartition(ctx context.Context, in *internalpb.DropPartitionRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	//*
	// @brief This method is used to test partition existence.
	//
	// @return BoolResponse
	HasPartition(ctx context.Context, in *internalpb.HasPartitionRequest, opts ...grpc.CallOption) (*servicepb.BoolResponse, error)
	//*
	// @brief This method is used to get basic partition infomation.
	//
	// @return PartitionDescription
	DescribePartition(ctx context.Context, in *internalpb.DescribePartitionRequest, opts ...grpc.CallOption) (*servicepb.PartitionDescription, error)
	//*
	// @brief This method is used to show partition information
	//
	// @param ShowPartitionRequest, target collection name.
	//
	// @return StringListResponse
	ShowPartitions(ctx context.Context, in *internalpb.ShowPartitionRequest, opts ...grpc.CallOption) (*servicepb.StringListResponse, error)
}

type masterClient struct {
	cc *grpc.ClientConn
}

func NewMasterClient(cc *grpc.ClientConn) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) CreateCollection(ctx context.Context, in *internalpb.CreateCollectionRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/CreateCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) DropCollection(ctx context.Context, in *internalpb.DropCollectionRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/DropCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) HasCollection(ctx context.Context, in *internalpb.HasCollectionRequest, opts ...grpc.CallOption) (*servicepb.BoolResponse, error) {
	out := new(servicepb.BoolResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/HasCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) DescribeCollection(ctx context.Context, in *internalpb.DescribeCollectionRequest, opts ...grpc.CallOption) (*servicepb.CollectionDescription, error) {
	out := new(servicepb.CollectionDescription)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/DescribeCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) ShowCollections(ctx context.Context, in *internalpb.ShowCollectionRequest, opts ...grpc.CallOption) (*servicepb.StringListResponse, error) {
	out := new(servicepb.StringListResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/ShowCollections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) CreatePartition(ctx context.Context, in *internalpb.CreatePartitionRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/CreatePartition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) DropPartition(ctx context.Context, in *internalpb.DropPartitionRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/DropPartition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) HasPartition(ctx context.Context, in *internalpb.HasPartitionRequest, opts ...grpc.CallOption) (*servicepb.BoolResponse, error) {
	out := new(servicepb.BoolResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/HasPartition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) DescribePartition(ctx context.Context, in *internalpb.DescribePartitionRequest, opts ...grpc.CallOption) (*servicepb.PartitionDescription, error) {
	out := new(servicepb.PartitionDescription)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/DescribePartition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) ShowPartitions(ctx context.Context, in *internalpb.ShowPartitionRequest, opts ...grpc.CallOption) (*servicepb.StringListResponse, error) {
	out := new(servicepb.StringListResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/ShowPartitions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServer is the server API for Master service.
type MasterServer interface {
	//*
	// @brief This method is used to create collection
	//
	// @param CreateCollectionRequest, use to provide collection information to be created.
	//
	// @return Status
	CreateCollection(context.Context, *internalpb.CreateCollectionRequest) (*commonpb.Status, error)
	//*
	// @brief This method is used to delete collection.
	//
	// @param DropCollectionRequest, collection name is going to be deleted.
	//
	// @return Status
	DropCollection(context.Context, *internalpb.DropCollectionRequest) (*commonpb.Status, error)
	//*
	// @brief This method is used to test collection existence.
	//
	// @param HasCollectionRequest, collection name is going to be tested.
	//
	// @return BoolResponse
	HasCollection(context.Context, *internalpb.HasCollectionRequest) (*servicepb.BoolResponse, error)
	//*
	// @brief This method is used to get collection schema.
	//
	// @param DescribeCollectionRequest, target collection name.
	//
	// @return CollectionSchema
	DescribeCollection(context.Context, *internalpb.DescribeCollectionRequest) (*servicepb.CollectionDescription, error)
	//*
	// @brief This method is used to list all collections.
	//
	// @return StringListResponse, collection name list
	ShowCollections(context.Context, *internalpb.ShowCollectionRequest) (*servicepb.StringListResponse, error)
	//*
	// @brief This method is used to create partition
	//
	// @return Status
	CreatePartition(context.Context, *internalpb.CreatePartitionRequest) (*commonpb.Status, error)
	//*
	// @brief This method is used to drop partition
	//
	// @return Status
	DropPartition(context.Context, *internalpb.DropPartitionRequest) (*commonpb.Status, error)
	//*
	// @brief This method is used to test partition existence.
	//
	// @return BoolResponse
	HasPartition(context.Context, *internalpb.HasPartitionRequest) (*servicepb.BoolResponse, error)
	//*
	// @brief This method is used to get basic partition infomation.
	//
	// @return PartitionDescription
	DescribePartition(context.Context, *internalpb.DescribePartitionRequest) (*servicepb.PartitionDescription, error)
	//*
	// @brief This method is used to show partition information
	//
	// @param ShowPartitionRequest, target collection name.
	//
	// @return StringListResponse
	ShowPartitions(context.Context, *internalpb.ShowPartitionRequest) (*servicepb.StringListResponse, error)
}

// UnimplementedMasterServer can be embedded to have forward compatible implementations.
type UnimplementedMasterServer struct {
}

func (*UnimplementedMasterServer) CreateCollection(ctx context.Context, req *internalpb.CreateCollectionRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCollection not implemented")
}
func (*UnimplementedMasterServer) DropCollection(ctx context.Context, req *internalpb.DropCollectionRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropCollection not implemented")
}
func (*UnimplementedMasterServer) HasCollection(ctx context.Context, req *internalpb.HasCollectionRequest) (*servicepb.BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasCollection not implemented")
}
func (*UnimplementedMasterServer) DescribeCollection(ctx context.Context, req *internalpb.DescribeCollectionRequest) (*servicepb.CollectionDescription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeCollection not implemented")
}
func (*UnimplementedMasterServer) ShowCollections(ctx context.Context, req *internalpb.ShowCollectionRequest) (*servicepb.StringListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowCollections not implemented")
}
func (*UnimplementedMasterServer) CreatePartition(ctx context.Context, req *internalpb.CreatePartitionRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePartition not implemented")
}
func (*UnimplementedMasterServer) DropPartition(ctx context.Context, req *internalpb.DropPartitionRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropPartition not implemented")
}
func (*UnimplementedMasterServer) HasPartition(ctx context.Context, req *internalpb.HasPartitionRequest) (*servicepb.BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasPartition not implemented")
}
func (*UnimplementedMasterServer) DescribePartition(ctx context.Context, req *internalpb.DescribePartitionRequest) (*servicepb.PartitionDescription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePartition not implemented")
}
func (*UnimplementedMasterServer) ShowPartitions(ctx context.Context, req *internalpb.ShowPartitionRequest) (*servicepb.StringListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowPartitions not implemented")
}

func RegisterMasterServer(s *grpc.Server, srv MasterServer) {
	s.RegisterService(&_Master_serviceDesc, srv)
}

func _Master_CreateCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.CreateCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).CreateCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/CreateCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).CreateCollection(ctx, req.(*internalpb.CreateCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_DropCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.DropCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).DropCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/DropCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).DropCollection(ctx, req.(*internalpb.DropCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_HasCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.HasCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).HasCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/HasCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).HasCollection(ctx, req.(*internalpb.HasCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_DescribeCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.DescribeCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).DescribeCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/DescribeCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).DescribeCollection(ctx, req.(*internalpb.DescribeCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_ShowCollections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.ShowCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ShowCollections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/ShowCollections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ShowCollections(ctx, req.(*internalpb.ShowCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_CreatePartition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.CreatePartitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).CreatePartition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/CreatePartition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).CreatePartition(ctx, req.(*internalpb.CreatePartitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_DropPartition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.DropPartitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).DropPartition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/DropPartition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).DropPartition(ctx, req.(*internalpb.DropPartitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_HasPartition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.HasPartitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).HasPartition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/HasPartition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).HasPartition(ctx, req.(*internalpb.HasPartitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_DescribePartition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.DescribePartitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).DescribePartition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/DescribePartition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).DescribePartition(ctx, req.(*internalpb.DescribePartitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_ShowPartitions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.ShowPartitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ShowPartitions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/ShowPartitions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ShowPartitions(ctx, req.(*internalpb.ShowPartitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Master_serviceDesc = grpc.ServiceDesc{
	ServiceName: "milvus.proto.master.Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCollection",
			Handler:    _Master_CreateCollection_Handler,
		},
		{
			MethodName: "DropCollection",
			Handler:    _Master_DropCollection_Handler,
		},
		{
			MethodName: "HasCollection",
			Handler:    _Master_HasCollection_Handler,
		},
		{
			MethodName: "DescribeCollection",
			Handler:    _Master_DescribeCollection_Handler,
		},
		{
			MethodName: "ShowCollections",
			Handler:    _Master_ShowCollections_Handler,
		},
		{
			MethodName: "CreatePartition",
			Handler:    _Master_CreatePartition_Handler,
		},
		{
			MethodName: "DropPartition",
			Handler:    _Master_DropPartition_Handler,
		},
		{
			MethodName: "HasPartition",
			Handler:    _Master_HasPartition_Handler,
		},
		{
			MethodName: "DescribePartition",
			Handler:    _Master_DescribePartition_Handler,
		},
		{
			MethodName: "ShowPartitions",
			Handler:    _Master_ShowPartitions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master.proto",
}
