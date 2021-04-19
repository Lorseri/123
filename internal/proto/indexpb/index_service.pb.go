// Code generated by protoc-gen-go. DO NOT EDIT.
// source: index_service.proto

package indexpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	commonpb "github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	internalpb2 "github.com/zilliztech/milvus-distributed/internal/proto/internalpb2"
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

type RegisterNodeRequest struct {
	Base                 *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Address              *commonpb.Address `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RegisterNodeRequest) Reset()         { *m = RegisterNodeRequest{} }
func (m *RegisterNodeRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterNodeRequest) ProtoMessage()    {}
func (*RegisterNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{0}
}

func (m *RegisterNodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterNodeRequest.Unmarshal(m, b)
}
func (m *RegisterNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterNodeRequest.Marshal(b, m, deterministic)
}
func (m *RegisterNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterNodeRequest.Merge(m, src)
}
func (m *RegisterNodeRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterNodeRequest.Size(m)
}
func (m *RegisterNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterNodeRequest proto.InternalMessageInfo

func (m *RegisterNodeRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *RegisterNodeRequest) GetAddress() *commonpb.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

type RegisterNodeResponse struct {
	InitParams           *internalpb2.InitParams `protobuf:"bytes,1,opt,name=init_params,json=initParams,proto3" json:"init_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RegisterNodeResponse) Reset()         { *m = RegisterNodeResponse{} }
func (m *RegisterNodeResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterNodeResponse) ProtoMessage()    {}
func (*RegisterNodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{1}
}

func (m *RegisterNodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterNodeResponse.Unmarshal(m, b)
}
func (m *RegisterNodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterNodeResponse.Marshal(b, m, deterministic)
}
func (m *RegisterNodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterNodeResponse.Merge(m, src)
}
func (m *RegisterNodeResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterNodeResponse.Size(m)
}
func (m *RegisterNodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterNodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterNodeResponse proto.InternalMessageInfo

func (m *RegisterNodeResponse) GetInitParams() *internalpb2.InitParams {
	if m != nil {
		return m.InitParams
	}
	return nil
}

type IndexStatesRequest struct {
	IndexID              int64    `protobuf:"varint,1,opt,name=indexID,proto3" json:"indexID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexStatesRequest) Reset()         { *m = IndexStatesRequest{} }
func (m *IndexStatesRequest) String() string { return proto.CompactTextString(m) }
func (*IndexStatesRequest) ProtoMessage()    {}
func (*IndexStatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{2}
}

func (m *IndexStatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexStatesRequest.Unmarshal(m, b)
}
func (m *IndexStatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexStatesRequest.Marshal(b, m, deterministic)
}
func (m *IndexStatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexStatesRequest.Merge(m, src)
}
func (m *IndexStatesRequest) XXX_Size() int {
	return xxx_messageInfo_IndexStatesRequest.Size(m)
}
func (m *IndexStatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexStatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IndexStatesRequest proto.InternalMessageInfo

func (m *IndexStatesRequest) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

type IndexStatesResponse struct {
	Status               *commonpb.Status    `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	State                commonpb.IndexState `protobuf:"varint,2,opt,name=state,proto3,enum=milvus.proto.common.IndexState" json:"state,omitempty"`
	IndexID              int64               `protobuf:"varint,3,opt,name=indexID,proto3" json:"indexID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IndexStatesResponse) Reset()         { *m = IndexStatesResponse{} }
func (m *IndexStatesResponse) String() string { return proto.CompactTextString(m) }
func (*IndexStatesResponse) ProtoMessage()    {}
func (*IndexStatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{3}
}

func (m *IndexStatesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexStatesResponse.Unmarshal(m, b)
}
func (m *IndexStatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexStatesResponse.Marshal(b, m, deterministic)
}
func (m *IndexStatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexStatesResponse.Merge(m, src)
}
func (m *IndexStatesResponse) XXX_Size() int {
	return xxx_messageInfo_IndexStatesResponse.Size(m)
}
func (m *IndexStatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexStatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IndexStatesResponse proto.InternalMessageInfo

func (m *IndexStatesResponse) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *IndexStatesResponse) GetState() commonpb.IndexState {
	if m != nil {
		return m.State
	}
	return commonpb.IndexState_NONE
}

func (m *IndexStatesResponse) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

type BuildIndexRequest struct {
	DataPaths            []string                 `protobuf:"bytes,2,rep,name=data_paths,json=dataPaths,proto3" json:"data_paths,omitempty"`
	TypeParams           []*commonpb.KeyValuePair `protobuf:"bytes,3,rep,name=type_params,json=typeParams,proto3" json:"type_params,omitempty"`
	IndexParams          []*commonpb.KeyValuePair `protobuf:"bytes,4,rep,name=index_params,json=indexParams,proto3" json:"index_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *BuildIndexRequest) Reset()         { *m = BuildIndexRequest{} }
func (m *BuildIndexRequest) String() string { return proto.CompactTextString(m) }
func (*BuildIndexRequest) ProtoMessage()    {}
func (*BuildIndexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{4}
}

func (m *BuildIndexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildIndexRequest.Unmarshal(m, b)
}
func (m *BuildIndexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildIndexRequest.Marshal(b, m, deterministic)
}
func (m *BuildIndexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildIndexRequest.Merge(m, src)
}
func (m *BuildIndexRequest) XXX_Size() int {
	return xxx_messageInfo_BuildIndexRequest.Size(m)
}
func (m *BuildIndexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildIndexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuildIndexRequest proto.InternalMessageInfo

func (m *BuildIndexRequest) GetDataPaths() []string {
	if m != nil {
		return m.DataPaths
	}
	return nil
}

func (m *BuildIndexRequest) GetTypeParams() []*commonpb.KeyValuePair {
	if m != nil {
		return m.TypeParams
	}
	return nil
}

func (m *BuildIndexRequest) GetIndexParams() []*commonpb.KeyValuePair {
	if m != nil {
		return m.IndexParams
	}
	return nil
}

type BuildIndexResponse struct {
	Status               *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IndexID              int64            `protobuf:"varint,2,opt,name=indexID,proto3" json:"indexID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BuildIndexResponse) Reset()         { *m = BuildIndexResponse{} }
func (m *BuildIndexResponse) String() string { return proto.CompactTextString(m) }
func (*BuildIndexResponse) ProtoMessage()    {}
func (*BuildIndexResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{5}
}

func (m *BuildIndexResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildIndexResponse.Unmarshal(m, b)
}
func (m *BuildIndexResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildIndexResponse.Marshal(b, m, deterministic)
}
func (m *BuildIndexResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildIndexResponse.Merge(m, src)
}
func (m *BuildIndexResponse) XXX_Size() int {
	return xxx_messageInfo_BuildIndexResponse.Size(m)
}
func (m *BuildIndexResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildIndexResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BuildIndexResponse proto.InternalMessageInfo

func (m *BuildIndexResponse) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *BuildIndexResponse) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

type BuildIndexCmd struct {
	IndexID              int64              `protobuf:"varint,1,opt,name=indexID,proto3" json:"indexID,omitempty"`
	Req                  *BuildIndexRequest `protobuf:"bytes,2,opt,name=req,proto3" json:"req,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BuildIndexCmd) Reset()         { *m = BuildIndexCmd{} }
func (m *BuildIndexCmd) String() string { return proto.CompactTextString(m) }
func (*BuildIndexCmd) ProtoMessage()    {}
func (*BuildIndexCmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{6}
}

func (m *BuildIndexCmd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildIndexCmd.Unmarshal(m, b)
}
func (m *BuildIndexCmd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildIndexCmd.Marshal(b, m, deterministic)
}
func (m *BuildIndexCmd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildIndexCmd.Merge(m, src)
}
func (m *BuildIndexCmd) XXX_Size() int {
	return xxx_messageInfo_BuildIndexCmd.Size(m)
}
func (m *BuildIndexCmd) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildIndexCmd.DiscardUnknown(m)
}

var xxx_messageInfo_BuildIndexCmd proto.InternalMessageInfo

func (m *BuildIndexCmd) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

func (m *BuildIndexCmd) GetReq() *BuildIndexRequest {
	if m != nil {
		return m.Req
	}
	return nil
}

type BuildIndexNotification struct {
	Status               *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IndexID              int64            `protobuf:"varint,2,opt,name=indexID,proto3" json:"indexID,omitempty"`
	IndexFilePaths       []string         `protobuf:"bytes,3,rep,name=index_file_paths,json=indexFilePaths,proto3" json:"index_file_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BuildIndexNotification) Reset()         { *m = BuildIndexNotification{} }
func (m *BuildIndexNotification) String() string { return proto.CompactTextString(m) }
func (*BuildIndexNotification) ProtoMessage()    {}
func (*BuildIndexNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{7}
}

func (m *BuildIndexNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildIndexNotification.Unmarshal(m, b)
}
func (m *BuildIndexNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildIndexNotification.Marshal(b, m, deterministic)
}
func (m *BuildIndexNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildIndexNotification.Merge(m, src)
}
func (m *BuildIndexNotification) XXX_Size() int {
	return xxx_messageInfo_BuildIndexNotification.Size(m)
}
func (m *BuildIndexNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildIndexNotification.DiscardUnknown(m)
}

var xxx_messageInfo_BuildIndexNotification proto.InternalMessageInfo

func (m *BuildIndexNotification) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *BuildIndexNotification) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

func (m *BuildIndexNotification) GetIndexFilePaths() []string {
	if m != nil {
		return m.IndexFilePaths
	}
	return nil
}

type IndexFilePathRequest struct {
	IndexID              int64    `protobuf:"varint,1,opt,name=indexID,proto3" json:"indexID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexFilePathRequest) Reset()         { *m = IndexFilePathRequest{} }
func (m *IndexFilePathRequest) String() string { return proto.CompactTextString(m) }
func (*IndexFilePathRequest) ProtoMessage()    {}
func (*IndexFilePathRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{8}
}

func (m *IndexFilePathRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexFilePathRequest.Unmarshal(m, b)
}
func (m *IndexFilePathRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexFilePathRequest.Marshal(b, m, deterministic)
}
func (m *IndexFilePathRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexFilePathRequest.Merge(m, src)
}
func (m *IndexFilePathRequest) XXX_Size() int {
	return xxx_messageInfo_IndexFilePathRequest.Size(m)
}
func (m *IndexFilePathRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexFilePathRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IndexFilePathRequest proto.InternalMessageInfo

func (m *IndexFilePathRequest) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

type IndexFilePathsResponse struct {
	Status               *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IndexID              int64            `protobuf:"varint,2,opt,name=indexID,proto3" json:"indexID,omitempty"`
	IndexFilePaths       []string         `protobuf:"bytes,3,rep,name=index_file_paths,json=indexFilePaths,proto3" json:"index_file_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *IndexFilePathsResponse) Reset()         { *m = IndexFilePathsResponse{} }
func (m *IndexFilePathsResponse) String() string { return proto.CompactTextString(m) }
func (*IndexFilePathsResponse) ProtoMessage()    {}
func (*IndexFilePathsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{9}
}

func (m *IndexFilePathsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexFilePathsResponse.Unmarshal(m, b)
}
func (m *IndexFilePathsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexFilePathsResponse.Marshal(b, m, deterministic)
}
func (m *IndexFilePathsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexFilePathsResponse.Merge(m, src)
}
func (m *IndexFilePathsResponse) XXX_Size() int {
	return xxx_messageInfo_IndexFilePathsResponse.Size(m)
}
func (m *IndexFilePathsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexFilePathsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IndexFilePathsResponse proto.InternalMessageInfo

func (m *IndexFilePathsResponse) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *IndexFilePathsResponse) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

func (m *IndexFilePathsResponse) GetIndexFilePaths() []string {
	if m != nil {
		return m.IndexFilePaths
	}
	return nil
}

type IndexMeta struct {
	State                commonpb.IndexState `protobuf:"varint,1,opt,name=state,proto3,enum=milvus.proto.common.IndexState" json:"state,omitempty"`
	IndexID              int64               `protobuf:"varint,2,opt,name=indexID,proto3" json:"indexID,omitempty"`
	EnqueTime            int64               `protobuf:"varint,3,opt,name=enque_time,json=enqueTime,proto3" json:"enque_time,omitempty"`
	ScheduleTime         int64               `protobuf:"varint,4,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	BuildCompleteTime    int64               `protobuf:"varint,5,opt,name=build_complete_time,json=buildCompleteTime,proto3" json:"build_complete_time,omitempty"`
	Req                  *BuildIndexRequest  `protobuf:"bytes,6,opt,name=req,proto3" json:"req,omitempty"`
	IndexFilePaths       []string            `protobuf:"bytes,7,rep,name=index_file_paths,json=indexFilePaths,proto3" json:"index_file_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IndexMeta) Reset()         { *m = IndexMeta{} }
func (m *IndexMeta) String() string { return proto.CompactTextString(m) }
func (*IndexMeta) ProtoMessage()    {}
func (*IndexMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{10}
}

func (m *IndexMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexMeta.Unmarshal(m, b)
}
func (m *IndexMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexMeta.Marshal(b, m, deterministic)
}
func (m *IndexMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexMeta.Merge(m, src)
}
func (m *IndexMeta) XXX_Size() int {
	return xxx_messageInfo_IndexMeta.Size(m)
}
func (m *IndexMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexMeta.DiscardUnknown(m)
}

var xxx_messageInfo_IndexMeta proto.InternalMessageInfo

func (m *IndexMeta) GetState() commonpb.IndexState {
	if m != nil {
		return m.State
	}
	return commonpb.IndexState_NONE
}

func (m *IndexMeta) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

func (m *IndexMeta) GetEnqueTime() int64 {
	if m != nil {
		return m.EnqueTime
	}
	return 0
}

func (m *IndexMeta) GetScheduleTime() int64 {
	if m != nil {
		return m.ScheduleTime
	}
	return 0
}

func (m *IndexMeta) GetBuildCompleteTime() int64 {
	if m != nil {
		return m.BuildCompleteTime
	}
	return 0
}

func (m *IndexMeta) GetReq() *BuildIndexRequest {
	if m != nil {
		return m.Req
	}
	return nil
}

func (m *IndexMeta) GetIndexFilePaths() []string {
	if m != nil {
		return m.IndexFilePaths
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterNodeRequest)(nil), "milvus.proto.index.RegisterNodeRequest")
	proto.RegisterType((*RegisterNodeResponse)(nil), "milvus.proto.index.RegisterNodeResponse")
	proto.RegisterType((*IndexStatesRequest)(nil), "milvus.proto.index.IndexStatesRequest")
	proto.RegisterType((*IndexStatesResponse)(nil), "milvus.proto.index.IndexStatesResponse")
	proto.RegisterType((*BuildIndexRequest)(nil), "milvus.proto.index.BuildIndexRequest")
	proto.RegisterType((*BuildIndexResponse)(nil), "milvus.proto.index.BuildIndexResponse")
	proto.RegisterType((*BuildIndexCmd)(nil), "milvus.proto.index.BuildIndexCmd")
	proto.RegisterType((*BuildIndexNotification)(nil), "milvus.proto.index.BuildIndexNotification")
	proto.RegisterType((*IndexFilePathRequest)(nil), "milvus.proto.index.IndexFilePathRequest")
	proto.RegisterType((*IndexFilePathsResponse)(nil), "milvus.proto.index.IndexFilePathsResponse")
	proto.RegisterType((*IndexMeta)(nil), "milvus.proto.index.IndexMeta")
}

func init() { proto.RegisterFile("index_service.proto", fileDescriptor_a5d2036b4df73e0a) }

var fileDescriptor_a5d2036b4df73e0a = []byte{
	// 714 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcd, 0x4e, 0xdb, 0x40,
	0x10, 0x26, 0x24, 0x80, 0x32, 0x84, 0x08, 0x36, 0xa8, 0x42, 0x69, 0x51, 0xc1, 0x55, 0x4b, 0x84,
	0x54, 0x07, 0x05, 0xb5, 0x3d, 0x56, 0x04, 0xd4, 0x2a, 0xaa, 0x40, 0xc8, 0x54, 0x3d, 0x50, 0x55,
	0xd1, 0xda, 0x1e, 0xc8, 0xaa, 0xfe, 0xc3, 0xbb, 0x46, 0x85, 0x4b, 0x0f, 0x7d, 0x82, 0x9e, 0xfb,
	0x18, 0xbd, 0xf6, 0xe1, 0x2a, 0xaf, 0xd7, 0x89, 0x4d, 0x4c, 0x02, 0xfd, 0xb9, 0x79, 0x67, 0xbf,
	0x99, 0xf9, 0xe6, 0x9b, 0x9d, 0x31, 0x34, 0x98, 0x67, 0xe3, 0x97, 0x3e, 0xc7, 0xf0, 0x92, 0x59,
	0xa8, 0x07, 0xa1, 0x2f, 0x7c, 0x42, 0x5c, 0xe6, 0x5c, 0x46, 0x3c, 0x39, 0xe9, 0x12, 0xd1, 0xac,
	0x59, 0xbe, 0xeb, 0xfa, 0x5e, 0x62, 0x6b, 0xd6, 0x99, 0x27, 0x30, 0xf4, 0xa8, 0x93, 0x9c, 0xb5,
	0xaf, 0xd0, 0x30, 0xf0, 0x9c, 0x71, 0x81, 0xe1, 0x91, 0x6f, 0xa3, 0x81, 0x17, 0x11, 0x72, 0x41,
	0x76, 0xa0, 0x62, 0x52, 0x8e, 0x6b, 0xa5, 0x8d, 0x52, 0x6b, 0xb1, 0xf3, 0x48, 0xcf, 0xc5, 0x55,
	0x01, 0x0f, 0xf9, 0x79, 0x97, 0x72, 0x34, 0x24, 0x92, 0xbc, 0x84, 0x05, 0x6a, 0xdb, 0x21, 0x72,
	0xbe, 0x36, 0x3b, 0xc1, 0x69, 0x2f, 0xc1, 0x18, 0x29, 0x58, 0x3b, 0x85, 0xd5, 0x3c, 0x01, 0x1e,
	0xf8, 0x1e, 0x47, 0xd2, 0x85, 0x45, 0xe6, 0x31, 0xd1, 0x0f, 0x68, 0x48, 0x5d, 0xae, 0x88, 0x6c,
	0xea, 0x37, 0x0a, 0x54, 0xb5, 0xf4, 0x3c, 0x26, 0x8e, 0x25, 0xd0, 0x00, 0x36, 0xfc, 0xd6, 0x74,
	0x20, 0xbd, 0x58, 0x83, 0x13, 0x41, 0x05, 0xf2, 0xb4, 0xb6, 0x35, 0x58, 0x90, 0xca, 0xf4, 0x0e,
	0x64, 0xd4, 0xb2, 0x91, 0x1e, 0xb5, 0x1f, 0x25, 0x68, 0xe4, 0x1c, 0x14, 0x97, 0x5d, 0x98, 0xe7,
	0x82, 0x8a, 0x28, 0xa5, 0xf1, 0xb0, 0xb0, 0xb4, 0x13, 0x09, 0x31, 0x14, 0x94, 0xbc, 0x80, 0xb9,
	0xf8, 0x0b, 0xa5, 0x1c, 0xf5, 0xce, 0xe3, 0x42, 0x9f, 0x51, 0x36, 0x23, 0x41, 0x67, 0xd9, 0x95,
	0xf3, 0xec, 0x7e, 0x95, 0x60, 0xa5, 0x1b, 0x31, 0xc7, 0x96, 0x4e, 0x69, 0x35, 0xeb, 0x00, 0x36,
	0x15, 0xb4, 0x1f, 0x50, 0x31, 0x88, 0xa5, 0x2f, 0xb7, 0xaa, 0x46, 0x35, 0xb6, 0x1c, 0xc7, 0x86,
	0x58, 0x46, 0x71, 0x15, 0x60, 0x2a, 0x63, 0x79, 0xa3, 0x3c, 0x2e, 0xa3, 0xe2, 0xf2, 0x0e, 0xaf,
	0x3e, 0x50, 0x27, 0xc2, 0x63, 0xca, 0x42, 0x03, 0x62, 0xaf, 0x44, 0x46, 0x72, 0x00, 0xb5, 0xe4,
	0xb1, 0xa9, 0x20, 0x95, 0xbb, 0x06, 0x59, 0x94, 0x6e, 0xaa, 0x19, 0x16, 0x90, 0x2c, 0xfb, 0xbf,
	0x91, 0x36, 0xa3, 0xd1, 0x6c, 0x5e, 0x23, 0x13, 0x96, 0x46, 0x49, 0xf6, 0x5d, 0xfb, 0xf6, 0x66,
	0x93, 0x57, 0x50, 0x0e, 0xf1, 0x42, 0x3d, 0xd6, 0xa7, 0xfa, 0xf8, 0xe4, 0xe8, 0x63, 0x62, 0x1b,
	0xb1, 0x87, 0xf6, 0xbd, 0x04, 0x0f, 0x46, 0x57, 0x47, 0xbe, 0x60, 0x67, 0xcc, 0xa2, 0x82, 0xf9,
	0xde, 0x3f, 0xae, 0x86, 0xb4, 0x60, 0x39, 0x11, 0xfe, 0x8c, 0x39, 0xa8, 0x3a, 0x5c, 0x96, 0x1d,
	0xae, 0x4b, 0xfb, 0x1b, 0xe6, 0xa0, 0x6c, 0xb3, 0xb6, 0x03, 0xab, 0xbd, 0xac, 0x65, 0xfa, 0x5b,
	0x8f, 0xab, 0xc8, 0xb9, 0xf0, 0xff, 0xd4, 0x93, 0x7b, 0x54, 0xf1, 0x73, 0x16, 0xaa, 0x92, 0xd3,
	0x21, 0x0a, 0x3a, 0x1a, 0xa0, 0xd2, 0x9f, 0x0e, 0xd0, 0x0d, 0x22, 0xeb, 0x00, 0xe8, 0x5d, 0x44,
	0xd8, 0x17, 0xcc, 0x45, 0x35, 0x5d, 0x55, 0x69, 0x79, 0xcf, 0x5c, 0x24, 0x4f, 0x60, 0x89, 0x5b,
	0x03, 0xb4, 0x23, 0x47, 0x21, 0x2a, 0x12, 0x51, 0x4b, 0x8d, 0x12, 0xa4, 0x43, 0xc3, 0x8c, 0x7b,
	0xdf, 0xb7, 0x7c, 0x37, 0x70, 0x50, 0x28, 0xe8, 0x9c, 0x84, 0xae, 0xc8, 0xab, 0x7d, 0x75, 0x23,
	0xf1, 0xea, 0x95, 0xcd, 0xdf, 0xf7, 0x95, 0x15, 0xaa, 0xb6, 0x50, 0xa4, 0x5a, 0xe7, 0x5b, 0x05,
	0x6a, 0x89, 0x0c, 0xc9, 0xbf, 0x80, 0x58, 0x50, 0xcb, 0xae, 0x54, 0xb2, 0x55, 0x94, 0xb6, 0x60,
	0xeb, 0x37, 0x5b, 0xd3, 0x81, 0xc9, 0x13, 0xd1, 0x66, 0xc8, 0x27, 0x80, 0x11, 0x73, 0x72, 0xb7,
	0xca, 0x9a, 0xcf, 0xa6, 0xc1, 0x86, 0xe1, 0x2d, 0xa8, 0xbf, 0x45, 0x91, 0x59, 0xc6, 0xa4, 0xd0,
	0x77, 0x7c, 0xbd, 0x37, 0xb7, 0xa6, 0xe2, 0x86, 0x49, 0x3e, 0xc3, 0x4a, 0x9a, 0x64, 0x28, 0x27,
	0x69, 0xdd, 0xea, 0x7f, 0x63, 0xb8, 0x9a, 0xdb, 0x53, 0x91, 0x3c, 0x27, 0xd8, 0xb2, 0xdc, 0x15,
	0x57, 0x19, 0xd9, 0xb6, 0x27, 0xeb, 0x91, 0xdd, 0x2d, 0xcd, 0x49, 0x53, 0xa8, 0xcd, 0x74, 0x3e,
	0xaa, 0xd1, 0x91, 0x1d, 0x3f, 0xca, 0x35, 0x67, 0x73, 0x72, 0x96, 0x7d, 0xd7, 0x9e, 0x12, 0xbc,
	0xbb, 0x77, 0xfa, 0xfa, 0x9c, 0x89, 0x41, 0x64, 0xc6, 0x37, 0xed, 0x6b, 0xe6, 0x38, 0xec, 0x5a,
	0xa0, 0x35, 0x68, 0x27, 0x5e, 0xcf, 0x6d, 0xc6, 0x45, 0xc8, 0xcc, 0x48, 0xa0, 0xdd, 0x4e, 0x7f,
	0xca, 0x6d, 0x19, 0xaa, 0x2d, 0xb3, 0x05, 0xa6, 0x39, 0x2f, 0x8f, 0xbb, 0xbf, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x89, 0xd5, 0x9c, 0x9e, 0xb8, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IndexServiceClient is the client API for IndexService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IndexServiceClient interface {
	//*
	// @brief This method is used to create collection
	//
	// @param CollectionSchema, use to provide collection information to be created.
	//
	// @return Status
	RegisterNode(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*RegisterNodeResponse, error)
	BuildIndex(ctx context.Context, in *BuildIndexRequest, opts ...grpc.CallOption) (*BuildIndexResponse, error)
	GetIndexStates(ctx context.Context, in *IndexStatesRequest, opts ...grpc.CallOption) (*IndexStatesResponse, error)
	GetIndexFilePaths(ctx context.Context, in *IndexFilePathRequest, opts ...grpc.CallOption) (*IndexFilePathsResponse, error)
	NotifyBuildIndex(ctx context.Context, in *BuildIndexNotification, opts ...grpc.CallOption) (*commonpb.Status, error)
}

type indexServiceClient struct {
	cc *grpc.ClientConn
}

func NewIndexServiceClient(cc *grpc.ClientConn) IndexServiceClient {
	return &indexServiceClient{cc}
}

func (c *indexServiceClient) RegisterNode(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*RegisterNodeResponse, error) {
	out := new(RegisterNodeResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.index.IndexService/RegisterNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) BuildIndex(ctx context.Context, in *BuildIndexRequest, opts ...grpc.CallOption) (*BuildIndexResponse, error) {
	out := new(BuildIndexResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.index.IndexService/BuildIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) GetIndexStates(ctx context.Context, in *IndexStatesRequest, opts ...grpc.CallOption) (*IndexStatesResponse, error) {
	out := new(IndexStatesResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.index.IndexService/GetIndexStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) GetIndexFilePaths(ctx context.Context, in *IndexFilePathRequest, opts ...grpc.CallOption) (*IndexFilePathsResponse, error) {
	out := new(IndexFilePathsResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.index.IndexService/GetIndexFilePaths", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) NotifyBuildIndex(ctx context.Context, in *BuildIndexNotification, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.index.IndexService/NotifyBuildIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexServiceServer is the server API for IndexService service.
type IndexServiceServer interface {
	//*
	// @brief This method is used to create collection
	//
	// @param CollectionSchema, use to provide collection information to be created.
	//
	// @return Status
	RegisterNode(context.Context, *RegisterNodeRequest) (*RegisterNodeResponse, error)
	BuildIndex(context.Context, *BuildIndexRequest) (*BuildIndexResponse, error)
	GetIndexStates(context.Context, *IndexStatesRequest) (*IndexStatesResponse, error)
	GetIndexFilePaths(context.Context, *IndexFilePathRequest) (*IndexFilePathsResponse, error)
	NotifyBuildIndex(context.Context, *BuildIndexNotification) (*commonpb.Status, error)
}

// UnimplementedIndexServiceServer can be embedded to have forward compatible implementations.
type UnimplementedIndexServiceServer struct {
}

func (*UnimplementedIndexServiceServer) RegisterNode(ctx context.Context, req *RegisterNodeRequest) (*RegisterNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNode not implemented")
}
func (*UnimplementedIndexServiceServer) BuildIndex(ctx context.Context, req *BuildIndexRequest) (*BuildIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildIndex not implemented")
}
func (*UnimplementedIndexServiceServer) GetIndexStates(ctx context.Context, req *IndexStatesRequest) (*IndexStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndexStates not implemented")
}
func (*UnimplementedIndexServiceServer) GetIndexFilePaths(ctx context.Context, req *IndexFilePathRequest) (*IndexFilePathsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndexFilePaths not implemented")
}
func (*UnimplementedIndexServiceServer) NotifyBuildIndex(ctx context.Context, req *BuildIndexNotification) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyBuildIndex not implemented")
}

func RegisterIndexServiceServer(s *grpc.Server, srv IndexServiceServer) {
	s.RegisterService(&_IndexService_serviceDesc, srv)
}

func _IndexService_RegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).RegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.index.IndexService/RegisterNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).RegisterNode(ctx, req.(*RegisterNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_BuildIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).BuildIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.index.IndexService/BuildIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).BuildIndex(ctx, req.(*BuildIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_GetIndexStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).GetIndexStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.index.IndexService/GetIndexStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).GetIndexStates(ctx, req.(*IndexStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_GetIndexFilePaths_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexFilePathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).GetIndexFilePaths(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.index.IndexService/GetIndexFilePaths",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).GetIndexFilePaths(ctx, req.(*IndexFilePathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_NotifyBuildIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildIndexNotification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).NotifyBuildIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.index.IndexService/NotifyBuildIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).NotifyBuildIndex(ctx, req.(*BuildIndexNotification))
	}
	return interceptor(ctx, in, info, handler)
}

var _IndexService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "milvus.proto.index.IndexService",
	HandlerType: (*IndexServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNode",
			Handler:    _IndexService_RegisterNode_Handler,
		},
		{
			MethodName: "BuildIndex",
			Handler:    _IndexService_BuildIndex_Handler,
		},
		{
			MethodName: "GetIndexStates",
			Handler:    _IndexService_GetIndexStates_Handler,
		},
		{
			MethodName: "GetIndexFilePaths",
			Handler:    _IndexService_GetIndexFilePaths_Handler,
		},
		{
			MethodName: "NotifyBuildIndex",
			Handler:    _IndexService_NotifyBuildIndex_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "index_service.proto",
}

// IndexNodeClient is the client API for IndexNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IndexNodeClient interface {
	//*
	// @brief This method is used to create collection
	//
	// @param CollectionSchema, use to provide collection information to be created.
	//
	// @return Status
	BuildIndex(ctx context.Context, in *BuildIndexCmd, opts ...grpc.CallOption) (*commonpb.Status, error)
}

type indexNodeClient struct {
	cc *grpc.ClientConn
}

func NewIndexNodeClient(cc *grpc.ClientConn) IndexNodeClient {
	return &indexNodeClient{cc}
}

func (c *indexNodeClient) BuildIndex(ctx context.Context, in *BuildIndexCmd, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.index.IndexNode/BuildIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexNodeServer is the server API for IndexNode service.
type IndexNodeServer interface {
	//*
	// @brief This method is used to create collection
	//
	// @param CollectionSchema, use to provide collection information to be created.
	//
	// @return Status
	BuildIndex(context.Context, *BuildIndexCmd) (*commonpb.Status, error)
}

// UnimplementedIndexNodeServer can be embedded to have forward compatible implementations.
type UnimplementedIndexNodeServer struct {
}

func (*UnimplementedIndexNodeServer) BuildIndex(ctx context.Context, req *BuildIndexCmd) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildIndex not implemented")
}

func RegisterIndexNodeServer(s *grpc.Server, srv IndexNodeServer) {
	s.RegisterService(&_IndexNode_serviceDesc, srv)
}

func _IndexNode_BuildIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildIndexCmd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexNodeServer).BuildIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.index.IndexNode/BuildIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexNodeServer).BuildIndex(ctx, req.(*BuildIndexCmd))
	}
	return interceptor(ctx, in, info, handler)
}

var _IndexNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "milvus.proto.index.IndexNode",
	HandlerType: (*IndexNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuildIndex",
			Handler:    _IndexNode_BuildIndex_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "index_service.proto",
}
