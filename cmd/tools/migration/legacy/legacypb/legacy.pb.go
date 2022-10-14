// Code generated by protoc-gen-go. DO NOT EDIT.
// source: legacy.proto

package legacypb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	commonpb "github.com/milvus-io/milvus-proto/go-api/commonpb"
	schemapb "github.com/milvus-io/milvus-proto/go-api/schemapb"
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

type BuildIndexRequest struct {
	IndexBuildID         int64                    `protobuf:"varint,1,opt,name=indexBuildID,proto3" json:"indexBuildID,omitempty"`
	IndexName            string                   `protobuf:"bytes,2,opt,name=index_name,json=indexName,proto3" json:"index_name,omitempty"`
	IndexID              int64                    `protobuf:"varint,3,opt,name=indexID,proto3" json:"indexID,omitempty"`
	DataPaths            []string                 `protobuf:"bytes,5,rep,name=data_paths,json=dataPaths,proto3" json:"data_paths,omitempty"`
	TypeParams           []*commonpb.KeyValuePair `protobuf:"bytes,6,rep,name=type_params,json=typeParams,proto3" json:"type_params,omitempty"`
	IndexParams          []*commonpb.KeyValuePair `protobuf:"bytes,7,rep,name=index_params,json=indexParams,proto3" json:"index_params,omitempty"`
	NumRows              int64                    `protobuf:"varint,8,opt,name=num_rows,json=numRows,proto3" json:"num_rows,omitempty"`
	FieldSchema          *schemapb.FieldSchema    `protobuf:"bytes,9,opt,name=field_schema,json=fieldSchema,proto3" json:"field_schema,omitempty"`
	SegmentID            int64                    `protobuf:"varint,10,opt,name=segmentID,proto3" json:"segmentID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *BuildIndexRequest) Reset()         { *m = BuildIndexRequest{} }
func (m *BuildIndexRequest) String() string { return proto.CompactTextString(m) }
func (*BuildIndexRequest) ProtoMessage()    {}
func (*BuildIndexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b5c555c498591f0, []int{0}
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

func (m *BuildIndexRequest) GetIndexBuildID() int64 {
	if m != nil {
		return m.IndexBuildID
	}
	return 0
}

func (m *BuildIndexRequest) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

func (m *BuildIndexRequest) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

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

func (m *BuildIndexRequest) GetNumRows() int64 {
	if m != nil {
		return m.NumRows
	}
	return 0
}

func (m *BuildIndexRequest) GetFieldSchema() *schemapb.FieldSchema {
	if m != nil {
		return m.FieldSchema
	}
	return nil
}

func (m *BuildIndexRequest) GetSegmentID() int64 {
	if m != nil {
		return m.SegmentID
	}
	return 0
}

type IndexMeta struct {
	IndexBuildID         int64               `protobuf:"varint,1,opt,name=indexBuildID,proto3" json:"indexBuildID,omitempty"`
	State                commonpb.IndexState `protobuf:"varint,2,opt,name=state,proto3,enum=milvus.proto.common.IndexState" json:"state,omitempty"`
	FailReason           string              `protobuf:"bytes,3,opt,name=fail_reason,json=failReason,proto3" json:"fail_reason,omitempty"`
	Req                  *BuildIndexRequest  `protobuf:"bytes,4,opt,name=req,proto3" json:"req,omitempty"`
	IndexFilePaths       []string            `protobuf:"bytes,5,rep,name=index_file_paths,json=indexFilePaths,proto3" json:"index_file_paths,omitempty"`
	MarkDeleted          bool                `protobuf:"varint,6,opt,name=mark_deleted,json=markDeleted,proto3" json:"mark_deleted,omitempty"`
	NodeID               int64               `protobuf:"varint,7,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	IndexVersion         int64               `protobuf:"varint,8,opt,name=index_version,json=indexVersion,proto3" json:"index_version,omitempty"`
	Recycled             bool                `protobuf:"varint,9,opt,name=recycled,proto3" json:"recycled,omitempty"`
	SerializeSize        uint64              `protobuf:"varint,10,opt,name=serialize_size,json=serializeSize,proto3" json:"serialize_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IndexMeta) Reset()         { *m = IndexMeta{} }
func (m *IndexMeta) String() string { return proto.CompactTextString(m) }
func (*IndexMeta) ProtoMessage()    {}
func (*IndexMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b5c555c498591f0, []int{1}
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

func (m *IndexMeta) GetIndexBuildID() int64 {
	if m != nil {
		return m.IndexBuildID
	}
	return 0
}

func (m *IndexMeta) GetState() commonpb.IndexState {
	if m != nil {
		return m.State
	}
	return commonpb.IndexState_IndexStateNone
}

func (m *IndexMeta) GetFailReason() string {
	if m != nil {
		return m.FailReason
	}
	return ""
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

func (m *IndexMeta) GetMarkDeleted() bool {
	if m != nil {
		return m.MarkDeleted
	}
	return false
}

func (m *IndexMeta) GetNodeID() int64 {
	if m != nil {
		return m.NodeID
	}
	return 0
}

func (m *IndexMeta) GetIndexVersion() int64 {
	if m != nil {
		return m.IndexVersion
	}
	return 0
}

func (m *IndexMeta) GetRecycled() bool {
	if m != nil {
		return m.Recycled
	}
	return false
}

func (m *IndexMeta) GetSerializeSize() uint64 {
	if m != nil {
		return m.SerializeSize
	}
	return 0
}

func init() {
	proto.RegisterType((*BuildIndexRequest)(nil), "milvus.proto.legacy.BuildIndexRequest")
	proto.RegisterType((*IndexMeta)(nil), "milvus.proto.legacy.IndexMeta")
}

func init() { proto.RegisterFile("legacy.proto", fileDescriptor_4b5c555c498591f0) }

var fileDescriptor_4b5c555c498591f0 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6b, 0xdb, 0x30,
	0x14, 0xc6, 0xc9, 0xdc, 0xa6, 0xf1, 0xb3, 0x5b, 0x36, 0x0d, 0x86, 0x56, 0x36, 0xea, 0x66, 0x6c,
	0xf8, 0xb2, 0x04, 0x52, 0x0a, 0x3b, 0x67, 0xa6, 0x10, 0xc6, 0x46, 0x50, 0xa0, 0x87, 0x5d, 0x8c,
	0x12, 0xbf, 0x24, 0x62, 0xb2, 0x9d, 0x4a, 0x72, 0xbb, 0xe4, 0x8f, 0xd8, 0x75, 0xff, 0xee, 0x90,
	0xe4, 0xb6, 0x84, 0xf5, 0xd0, 0x9b, 0xbf, 0x9f, 0xf5, 0x3d, 0xf9, 0x7d, 0x9f, 0x21, 0x96, 0xb8,
	0xe2, 0x8b, 0xed, 0x60, 0xa3, 0x6a, 0x53, 0x93, 0xd7, 0xa5, 0x90, 0xb7, 0x8d, 0xf6, 0x6a, 0xe0,
	0x5f, 0x9d, 0xc6, 0x8b, 0xba, 0x2c, 0xeb, 0xca, 0xc3, 0xd3, 0x58, 0x2f, 0xd6, 0x58, 0x72, 0xaf,
	0xfa, 0x7f, 0x03, 0x78, 0x35, 0x6e, 0x84, 0x2c, 0x26, 0x55, 0x81, 0xbf, 0x19, 0xde, 0x34, 0xa8,
	0x0d, 0xe9, 0x43, 0x2c, 0xac, 0xf6, 0x6f, 0x32, 0xda, 0x49, 0x3a, 0x69, 0xc0, 0xf6, 0x18, 0x79,
	0x0f, 0xe0, 0x74, 0x5e, 0xf1, 0x12, 0xe9, 0x8b, 0xa4, 0x93, 0x86, 0x2c, 0x74, 0xe4, 0x07, 0x2f,
	0x91, 0x50, 0x38, 0x72, 0x62, 0x92, 0xd1, 0xc0, 0xb9, 0xef, 0xa5, 0x35, 0x16, 0xdc, 0xf0, 0x7c,
	0xc3, 0xcd, 0x5a, 0xd3, 0xc3, 0x24, 0xb0, 0x46, 0x4b, 0xa6, 0x16, 0x90, 0x31, 0x44, 0x66, 0xbb,
	0xc1, 0x7c, 0xc3, 0x15, 0x2f, 0x35, 0xed, 0x26, 0x41, 0x1a, 0x8d, 0xce, 0x07, 0x7b, 0x8b, 0xb5,
	0x0b, 0x7d, 0xc3, 0xed, 0x35, 0x97, 0x0d, 0x4e, 0xb9, 0x50, 0x0c, 0xac, 0x6b, 0xea, 0x4c, 0x24,
	0x6b, 0xbf, 0xff, 0x7e, 0xc8, 0xd1, 0x73, 0x87, 0x44, 0xce, 0xd6, 0x4e, 0x79, 0x0b, 0xbd, 0xaa,
	0x29, 0x73, 0x55, 0xdf, 0x69, 0xda, 0xf3, 0x3b, 0x54, 0x4d, 0xc9, 0xea, 0x3b, 0x4d, 0xbe, 0x42,
	0xbc, 0x14, 0x28, 0x8b, 0xdc, 0x87, 0x49, 0xc3, 0xa4, 0x93, 0x46, 0xa3, 0x64, 0xff, 0x82, 0x36,
	0xe8, 0x2b, 0x7b, 0x70, 0xe6, 0x9e, 0x59, 0xb4, 0x7c, 0x14, 0xe4, 0x1d, 0x84, 0x1a, 0x57, 0x25,
	0x56, 0x66, 0x92, 0x51, 0x70, 0x17, 0x3c, 0x82, 0xfe, 0x9f, 0x00, 0x42, 0x57, 0xca, 0x77, 0x34,
	0xfc, 0x59, 0x8d, 0x5c, 0xc2, 0xa1, 0x36, 0xdc, 0xf8, 0x32, 0x4e, 0x46, 0x67, 0x4f, 0xae, 0xeb,
	0x46, 0xce, 0xec, 0x31, 0xe6, 0x4f, 0x93, 0x33, 0x88, 0x96, 0x5c, 0xc8, 0x5c, 0x21, 0xd7, 0x75,
	0xe5, 0xda, 0x0a, 0x19, 0x58, 0xc4, 0x1c, 0x21, 0x5f, 0x20, 0x50, 0x78, 0x43, 0x0f, 0xdc, 0x8e,
	0x9f, 0x06, 0x4f, 0xfc, 0x62, 0x83, 0xff, 0x7e, 0x21, 0x66, 0x2d, 0x24, 0x85, 0x97, 0xbe, 0x87,
	0xa5, 0x90, 0xb8, 0x57, 0xf8, 0x89, 0xe3, 0x57, 0x42, 0xa2, 0x6f, 0xfd, 0x1c, 0xe2, 0x92, 0xab,
	0x5f, 0x79, 0x81, 0x12, 0x0d, 0x16, 0xb4, 0x9b, 0x74, 0xd2, 0x1e, 0x8b, 0x2c, 0xcb, 0x3c, 0x22,
	0x6f, 0xa0, 0x5b, 0xd5, 0x05, 0x4e, 0x32, 0x7a, 0xe4, 0x96, 0x6f, 0x15, 0xf9, 0x00, 0xc7, 0xfe,
	0x92, 0x5b, 0x54, 0x5a, 0xd4, 0x55, 0xdb, 0x95, 0xcf, 0xe6, 0xda, 0x33, 0x72, 0x0a, 0x3d, 0x85,
	0x8b, 0xed, 0x42, 0x62, 0xe1, 0xca, 0xea, 0xb1, 0x07, 0x4d, 0x3e, 0xc2, 0x89, 0x46, 0x25, 0xb8,
	0x14, 0x3b, 0xcc, 0xb5, 0xd8, 0xa1, 0x2b, 0xe3, 0x80, 0x1d, 0x3f, 0xd0, 0x99, 0xd8, 0xe1, 0xf8,
	0xf2, 0xe7, 0xc5, 0x4a, 0x98, 0x75, 0x33, 0xb7, 0x51, 0x0e, 0x7d, 0x0a, 0x9f, 0x45, 0xdd, 0x3e,
	0x0d, 0x45, 0x65, 0x50, 0x55, 0x5c, 0x0e, 0x5d, 0x30, 0x43, 0x1f, 0xcc, 0x66, 0x3e, 0xef, 0x3a,
	0x7d, 0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x46, 0x37, 0x11, 0xcb, 0xa9, 0x03, 0x00, 0x00,
}
