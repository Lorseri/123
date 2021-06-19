// Code generated by protoc-gen-go. DO NOT EDIT.
// source: segcore.proto

package segcorepb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	schemapb "github.com/milvus-io/milvus/internal/proto/schemapb"
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

type RetrieveRequest struct {
	Ids                  *schemapb.IDs `protobuf:"bytes,1,opt,name=ids,proto3" json:"ids,omitempty"`
	OutputFields         []string      `protobuf:"bytes,2,rep,name=output_fields,json=outputFields,proto3" json:"output_fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RetrieveRequest) Reset()         { *m = RetrieveRequest{} }
func (m *RetrieveRequest) String() string { return proto.CompactTextString(m) }
func (*RetrieveRequest) ProtoMessage()    {}
func (*RetrieveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d79fce784797357, []int{0}
}

func (m *RetrieveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveRequest.Unmarshal(m, b)
}
func (m *RetrieveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveRequest.Marshal(b, m, deterministic)
}
func (m *RetrieveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveRequest.Merge(m, src)
}
func (m *RetrieveRequest) XXX_Size() int {
	return xxx_messageInfo_RetrieveRequest.Size(m)
}
func (m *RetrieveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveRequest proto.InternalMessageInfo

func (m *RetrieveRequest) GetIds() *schemapb.IDs {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *RetrieveRequest) GetOutputFields() []string {
	if m != nil {
		return m.OutputFields
	}
	return nil
}

type RetrieveResults struct {
	Ids                  *schemapb.IDs         `protobuf:"bytes,1,opt,name=ids,proto3" json:"ids,omitempty"`
	FieldsData           []*schemapb.FieldData `protobuf:"bytes,2,rep,name=fields_data,json=fieldsData,proto3" json:"fields_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RetrieveResults) Reset()         { *m = RetrieveResults{} }
func (m *RetrieveResults) String() string { return proto.CompactTextString(m) }
func (*RetrieveResults) ProtoMessage()    {}
func (*RetrieveResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d79fce784797357, []int{1}
}

func (m *RetrieveResults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveResults.Unmarshal(m, b)
}
func (m *RetrieveResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveResults.Marshal(b, m, deterministic)
}
func (m *RetrieveResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveResults.Merge(m, src)
}
func (m *RetrieveResults) XXX_Size() int {
	return xxx_messageInfo_RetrieveResults.Size(m)
}
func (m *RetrieveResults) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveResults.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveResults proto.InternalMessageInfo

func (m *RetrieveResults) GetIds() *schemapb.IDs {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *RetrieveResults) GetFieldsData() []*schemapb.FieldData {
	if m != nil {
		return m.FieldsData
	}
	return nil
}

type LoadFieldMeta struct {
	MinTimestamp         int64    `protobuf:"varint,1,opt,name=min_timestamp,json=minTimestamp,proto3" json:"min_timestamp,omitempty"`
	MaxTimestamp         int64    `protobuf:"varint,2,opt,name=max_timestamp,json=maxTimestamp,proto3" json:"max_timestamp,omitempty"`
	RowCount             int64    `protobuf:"varint,3,opt,name=row_count,json=rowCount,proto3" json:"row_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadFieldMeta) Reset()         { *m = LoadFieldMeta{} }
func (m *LoadFieldMeta) String() string { return proto.CompactTextString(m) }
func (*LoadFieldMeta) ProtoMessage()    {}
func (*LoadFieldMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d79fce784797357, []int{2}
}

func (m *LoadFieldMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadFieldMeta.Unmarshal(m, b)
}
func (m *LoadFieldMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadFieldMeta.Marshal(b, m, deterministic)
}
func (m *LoadFieldMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadFieldMeta.Merge(m, src)
}
func (m *LoadFieldMeta) XXX_Size() int {
	return xxx_messageInfo_LoadFieldMeta.Size(m)
}
func (m *LoadFieldMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadFieldMeta.DiscardUnknown(m)
}

var xxx_messageInfo_LoadFieldMeta proto.InternalMessageInfo

func (m *LoadFieldMeta) GetMinTimestamp() int64 {
	if m != nil {
		return m.MinTimestamp
	}
	return 0
}

func (m *LoadFieldMeta) GetMaxTimestamp() int64 {
	if m != nil {
		return m.MaxTimestamp
	}
	return 0
}

func (m *LoadFieldMeta) GetRowCount() int64 {
	if m != nil {
		return m.RowCount
	}
	return 0
}

type LoadSegmentMeta struct {
	// TODOs
	Metas                []*LoadFieldMeta `protobuf:"bytes,1,rep,name=metas,proto3" json:"metas,omitempty"`
	TotalSize            int64            `protobuf:"varint,2,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *LoadSegmentMeta) Reset()         { *m = LoadSegmentMeta{} }
func (m *LoadSegmentMeta) String() string { return proto.CompactTextString(m) }
func (*LoadSegmentMeta) ProtoMessage()    {}
func (*LoadSegmentMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d79fce784797357, []int{3}
}

func (m *LoadSegmentMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadSegmentMeta.Unmarshal(m, b)
}
func (m *LoadSegmentMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadSegmentMeta.Marshal(b, m, deterministic)
}
func (m *LoadSegmentMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadSegmentMeta.Merge(m, src)
}
func (m *LoadSegmentMeta) XXX_Size() int {
	return xxx_messageInfo_LoadSegmentMeta.Size(m)
}
func (m *LoadSegmentMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadSegmentMeta.DiscardUnknown(m)
}

var xxx_messageInfo_LoadSegmentMeta proto.InternalMessageInfo

func (m *LoadSegmentMeta) GetMetas() []*LoadFieldMeta {
	if m != nil {
		return m.Metas
	}
	return nil
}

func (m *LoadSegmentMeta) GetTotalSize() int64 {
	if m != nil {
		return m.TotalSize
	}
	return 0
}

func init() {
	proto.RegisterType((*RetrieveRequest)(nil), "milvus.proto.segcore.RetrieveRequest")
	proto.RegisterType((*RetrieveResults)(nil), "milvus.proto.segcore.RetrieveResults")
	proto.RegisterType((*LoadFieldMeta)(nil), "milvus.proto.segcore.LoadFieldMeta")
	proto.RegisterType((*LoadSegmentMeta)(nil), "milvus.proto.segcore.LoadSegmentMeta")
}

func init() { proto.RegisterFile("segcore.proto", fileDescriptor_1d79fce784797357) }

var fileDescriptor_1d79fce784797357 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4b, 0xf3, 0x30,
	0x18, 0xc6, 0xd9, 0xca, 0xf7, 0xe1, 0xb2, 0x8d, 0x41, 0xf1, 0x50, 0x14, 0x65, 0x74, 0x97, 0x21,
	0xd8, 0xc2, 0x14, 0xc1, 0x93, 0xa0, 0x43, 0x10, 0xf4, 0x92, 0x79, 0xf2, 0x52, 0xd2, 0xf6, 0x75,
	0x0b, 0x36, 0x4d, 0x6d, 0xde, 0x74, 0x63, 0x07, 0xff, 0x76, 0x49, 0x52, 0x71, 0x83, 0x5d, 0xbc,
	0x25, 0x4f, 0x9e, 0xe7, 0xfd, 0x3d, 0x2f, 0x21, 0x43, 0x05, 0xcb, 0x4c, 0xd6, 0x10, 0x55, 0xb5,
	0x44, 0xe9, 0x1f, 0x0b, 0x5e, 0x34, 0x5a, 0xb9, 0x5b, 0xd4, 0xbe, 0x9d, 0x0c, 0x54, 0xb6, 0x02,
	0xc1, 0x9c, 0x1a, 0xa6, 0x64, 0x44, 0x01, 0x6b, 0x0e, 0x0d, 0x50, 0xf8, 0xd4, 0xa0, 0xd0, 0xbf,
	0x20, 0x1e, 0xcf, 0x55, 0xd0, 0x19, 0x77, 0xa6, 0xfd, 0x59, 0x10, 0xed, 0x0f, 0x71, 0xd9, 0xa7,
	0xb9, 0xa2, 0xc6, 0xe4, 0x4f, 0xc8, 0x50, 0x6a, 0xac, 0x34, 0x26, 0xef, 0x1c, 0x8a, 0x5c, 0x05,
	0xdd, 0xb1, 0x37, 0xed, 0xd1, 0x81, 0x13, 0x1f, 0xad, 0x16, 0x7e, 0xed, 0x32, 0x94, 0x2e, 0x50,
	0xfd, 0x89, 0x71, 0x47, 0xfa, 0x6e, 0x78, 0x92, 0x33, 0x64, 0x96, 0xd0, 0x9f, 0x9d, 0x1f, 0xcc,
	0x58, 0xe0, 0x9c, 0x21, 0xa3, 0xc4, 0x45, 0xcc, 0x39, 0x6c, 0xc8, 0xf0, 0x59, 0xb2, 0xdc, 0x3e,
	0xbe, 0x00, 0x32, 0xd3, 0x5a, 0xf0, 0x32, 0x41, 0x2e, 0x40, 0x21, 0x13, 0x95, 0xed, 0xe1, 0xd1,
	0x81, 0xe0, 0xe5, 0xeb, 0x8f, 0x66, 0x4d, 0x6c, 0xb3, 0x63, 0xea, 0xb6, 0x26, 0xb6, 0xf9, 0x35,
	0x9d, 0x92, 0x5e, 0x2d, 0xd7, 0x49, 0x26, 0x75, 0x89, 0x81, 0x67, 0x0d, 0x47, 0xb5, 0x5c, 0x3f,
	0x98, 0x7b, 0xf8, 0x41, 0x46, 0x86, 0xbb, 0x80, 0xa5, 0x80, 0x12, 0x2d, 0xf9, 0x96, 0xfc, 0x13,
	0x80, 0xcc, 0x6c, 0x6e, 0xb6, 0x98, 0x44, 0x87, 0xbe, 0x28, 0xda, 0x6b, 0x4b, 0x5d, 0xc2, 0x3f,
	0x23, 0x04, 0x25, 0xb2, 0x22, 0x51, 0x7c, 0x0b, 0x6d, 0x99, 0x9e, 0x55, 0x16, 0x7c, 0x0b, 0xf7,
	0x37, 0x6f, 0xd7, 0x4b, 0x8e, 0x2b, 0x9d, 0x46, 0x99, 0x14, 0xb1, 0x1b, 0x7b, 0xc9, 0x65, 0x7b,
	0x8a, 0x79, 0x89, 0x50, 0x97, 0xac, 0x88, 0x2d, 0x29, 0x6e, 0x49, 0x55, 0x9a, 0xfe, 0xb7, 0xc2,
	0xd5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac, 0x78, 0x16, 0xc2, 0x3c, 0x02, 0x00, 0x00,
}
