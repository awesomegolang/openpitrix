// Code generated by protoc-gen-go. DO NOT EDIT.
// source: repo_indexer.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type IndexRepoRequest struct {
	RepoId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=repo_id,json=repoId" json:"repo_id,omitempty"`
}

func (m *IndexRepoRequest) Reset()                    { *m = IndexRepoRequest{} }
func (m *IndexRepoRequest) String() string            { return proto.CompactTextString(m) }
func (*IndexRepoRequest) ProtoMessage()               {}
func (*IndexRepoRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *IndexRepoRequest) GetRepoId() *google_protobuf.StringValue {
	if m != nil {
		return m.RepoId
	}
	return nil
}

type IndexRepoResponse struct {
	RepoTask *RepoTask `protobuf:"bytes,1,opt,name=repo_task,json=repoTask" json:"repo_task,omitempty"`
}

func (m *IndexRepoResponse) Reset()                    { *m = IndexRepoResponse{} }
func (m *IndexRepoResponse) String() string            { return proto.CompactTextString(m) }
func (*IndexRepoResponse) ProtoMessage()               {}
func (*IndexRepoResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *IndexRepoResponse) GetRepoTask() *RepoTask {
	if m != nil {
		return m.RepoTask
	}
	return nil
}

type RepoTask struct {
	RepoTaskId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=repo_task_id,json=repoTaskId" json:"repo_task_id,omitempty"`
	RepoId     *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=repo_id,json=repoId" json:"repo_id,omitempty"`
	Owner      *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
	Status     *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	Result     *google_protobuf.StringValue `protobuf:"bytes,5,opt,name=result" json:"result,omitempty"`
	CreateTime *google_protobuf1.Timestamp  `protobuf:"bytes,6,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	StatusTime *google_protobuf1.Timestamp  `protobuf:"bytes,7,opt,name=status_time,json=statusTime" json:"status_time,omitempty"`
}

func (m *RepoTask) Reset()                    { *m = RepoTask{} }
func (m *RepoTask) String() string            { return proto.CompactTextString(m) }
func (*RepoTask) ProtoMessage()               {}
func (*RepoTask) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *RepoTask) GetRepoTaskId() *google_protobuf.StringValue {
	if m != nil {
		return m.RepoTaskId
	}
	return nil
}

func (m *RepoTask) GetRepoId() *google_protobuf.StringValue {
	if m != nil {
		return m.RepoId
	}
	return nil
}

func (m *RepoTask) GetOwner() *google_protobuf.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *RepoTask) GetStatus() *google_protobuf.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *RepoTask) GetResult() *google_protobuf.StringValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *RepoTask) GetCreateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *RepoTask) GetStatusTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

type DescribeRepoTasksRequest struct {
	RepoTaskId []string `protobuf:"bytes,1,rep,name=repo_task_id,json=repoTaskId" json:"repo_task_id,omitempty"`
	RepoId     []string `protobuf:"bytes,2,rep,name=repo_id,json=repoId" json:"repo_id,omitempty"`
	Owner      []string `protobuf:"bytes,3,rep,name=owner" json:"owner,omitempty"`
	Status     []string `protobuf:"bytes,4,rep,name=status" json:"status,omitempty"`
	Limit      uint32   `protobuf:"varint,5,opt,name=limit" json:"limit,omitempty"`
	Offset     uint32   `protobuf:"varint,6,opt,name=offset" json:"offset,omitempty"`
}

func (m *DescribeRepoTasksRequest) Reset()                    { *m = DescribeRepoTasksRequest{} }
func (m *DescribeRepoTasksRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeRepoTasksRequest) ProtoMessage()               {}
func (*DescribeRepoTasksRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *DescribeRepoTasksRequest) GetRepoTaskId() []string {
	if m != nil {
		return m.RepoTaskId
	}
	return nil
}

func (m *DescribeRepoTasksRequest) GetRepoId() []string {
	if m != nil {
		return m.RepoId
	}
	return nil
}

func (m *DescribeRepoTasksRequest) GetOwner() []string {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *DescribeRepoTasksRequest) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeRepoTasksRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeRepoTasksRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type DescribeRepoTasksResponse struct {
	TotalCount  uint32      `protobuf:"varint,1,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	RepoTaskSet []*RepoTask `protobuf:"bytes,2,rep,name=repo_task_set,json=repoTaskSet" json:"repo_task_set,omitempty"`
}

func (m *DescribeRepoTasksResponse) Reset()                    { *m = DescribeRepoTasksResponse{} }
func (m *DescribeRepoTasksResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeRepoTasksResponse) ProtoMessage()               {}
func (*DescribeRepoTasksResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *DescribeRepoTasksResponse) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *DescribeRepoTasksResponse) GetRepoTaskSet() []*RepoTask {
	if m != nil {
		return m.RepoTaskSet
	}
	return nil
}

func init() {
	proto.RegisterType((*IndexRepoRequest)(nil), "openpitrix.IndexRepoRequest")
	proto.RegisterType((*IndexRepoResponse)(nil), "openpitrix.IndexRepoResponse")
	proto.RegisterType((*RepoTask)(nil), "openpitrix.RepoTask")
	proto.RegisterType((*DescribeRepoTasksRequest)(nil), "openpitrix.DescribeRepoTasksRequest")
	proto.RegisterType((*DescribeRepoTasksResponse)(nil), "openpitrix.DescribeRepoTasksResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RepoIndexer service

type RepoIndexerClient interface {
	IndexRepo(ctx context.Context, in *IndexRepoRequest, opts ...grpc.CallOption) (*IndexRepoResponse, error)
	DescribeRepoTasks(ctx context.Context, in *DescribeRepoTasksRequest, opts ...grpc.CallOption) (*DescribeRepoTasksResponse, error)
}

type repoIndexerClient struct {
	cc *grpc.ClientConn
}

func NewRepoIndexerClient(cc *grpc.ClientConn) RepoIndexerClient {
	return &repoIndexerClient{cc}
}

func (c *repoIndexerClient) IndexRepo(ctx context.Context, in *IndexRepoRequest, opts ...grpc.CallOption) (*IndexRepoResponse, error) {
	out := new(IndexRepoResponse)
	err := grpc.Invoke(ctx, "/openpitrix.RepoIndexer/IndexRepo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoIndexerClient) DescribeRepoTasks(ctx context.Context, in *DescribeRepoTasksRequest, opts ...grpc.CallOption) (*DescribeRepoTasksResponse, error) {
	out := new(DescribeRepoTasksResponse)
	err := grpc.Invoke(ctx, "/openpitrix.RepoIndexer/DescribeRepoTasks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RepoIndexer service

type RepoIndexerServer interface {
	IndexRepo(context.Context, *IndexRepoRequest) (*IndexRepoResponse, error)
	DescribeRepoTasks(context.Context, *DescribeRepoTasksRequest) (*DescribeRepoTasksResponse, error)
}

func RegisterRepoIndexerServer(s *grpc.Server, srv RepoIndexerServer) {
	s.RegisterService(&_RepoIndexer_serviceDesc, srv)
}

func _RepoIndexer_IndexRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoIndexerServer).IndexRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.RepoIndexer/IndexRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoIndexerServer).IndexRepo(ctx, req.(*IndexRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoIndexer_DescribeRepoTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRepoTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoIndexerServer).DescribeRepoTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.RepoIndexer/DescribeRepoTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoIndexerServer).DescribeRepoTasks(ctx, req.(*DescribeRepoTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RepoIndexer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.RepoIndexer",
	HandlerType: (*RepoIndexerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IndexRepo",
			Handler:    _RepoIndexer_IndexRepo_Handler,
		},
		{
			MethodName: "DescribeRepoTasks",
			Handler:    _RepoIndexer_DescribeRepoTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "repo_indexer.proto",
}

func init() { proto.RegisterFile("repo_indexer.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0x96, 0x93, 0x26, 0x6d, 0x26, 0xbf, 0xfc, 0x68, 0x97, 0x40, 0xdd, 0x28, 0x50, 0xcb, 0x02,
	0xa9, 0x42, 0xd4, 0x56, 0xc3, 0x1f, 0x21, 0x90, 0x90, 0x0a, 0x08, 0x29, 0x57, 0xb7, 0xe2, 0xc0,
	0x25, 0xda, 0x24, 0x13, 0xcb, 0x6a, 0xe2, 0x5d, 0x76, 0xd7, 0x4d, 0x8f, 0x88, 0x27, 0x40, 0xe5,
	0xc8, 0x91, 0x27, 0xe0, 0x59, 0x78, 0x05, 0x1e, 0x04, 0xed, 0xae, 0xdd, 0x98, 0x86, 0xa8, 0x39,
	0x25, 0x9e, 0xf9, 0xbe, 0x99, 0x6f, 0xbe, 0x19, 0x1b, 0x88, 0x40, 0xce, 0x06, 0x49, 0x3a, 0xc6,
	0x0b, 0x14, 0x01, 0x17, 0x4c, 0x31, 0x02, 0x8c, 0x63, 0xca, 0x13, 0x25, 0x92, 0x8b, 0xce, 0xfd,
	0x98, 0xb1, 0x78, 0x8a, 0xa1, 0xc9, 0x0c, 0xb3, 0x49, 0x38, 0x17, 0x94, 0x73, 0x14, 0xd2, 0x62,
	0x3b, 0xfb, 0xd7, 0xf3, 0x2a, 0x99, 0xa1, 0x54, 0x74, 0xc6, 0x73, 0x40, 0x37, 0x07, 0x50, 0x9e,
	0x84, 0x34, 0x4d, 0x99, 0xa2, 0x2a, 0x61, 0x69, 0x41, 0x7f, 0x6c, 0x7e, 0x46, 0x87, 0x31, 0xa6,
	0x87, 0x72, 0x4e, 0xe3, 0x18, 0x45, 0xc8, 0xb8, 0x41, 0x2c, 0xa3, 0xfd, 0x3e, 0x6c, 0xf7, 0xb5,
	0xd2, 0x08, 0x39, 0x8b, 0xf0, 0x53, 0x86, 0x52, 0x91, 0x67, 0xb0, 0x69, 0x47, 0x18, 0xbb, 0x8e,
	0xe7, 0x1c, 0x34, 0x7b, 0xdd, 0xc0, 0x76, 0x0c, 0x0a, 0x49, 0xc1, 0x89, 0x12, 0x49, 0x1a, 0x7f,
	0xa0, 0xd3, 0x0c, 0xa3, 0xba, 0x06, 0xf7, 0xc7, 0xfe, 0x7b, 0xd8, 0x29, 0x95, 0x92, 0x9c, 0xa5,
	0x12, 0xc9, 0x11, 0x34, 0x4c, 0x2d, 0x45, 0xe5, 0x59, 0x5e, 0xad, 0x1d, 0x2c, 0xcc, 0x08, 0x34,
	0xf8, 0x94, 0xca, 0xb3, 0x68, 0x4b, 0xe4, 0xff, 0xfc, 0xef, 0x55, 0xd8, 0x2a, 0xc2, 0xe4, 0x35,
	0xfc, 0x77, 0xc5, 0x5f, 0x57, 0x10, 0x14, 0xa5, 0xfa, 0xe3, 0xf2, 0x2c, 0x95, 0xf5, 0x67, 0x21,
	0x3d, 0xa8, 0xb1, 0x79, 0x8a, 0xc2, 0xad, 0xae, 0x41, 0xb2, 0x50, 0xf2, 0x14, 0xea, 0x52, 0x51,
	0x95, 0x49, 0x77, 0x63, 0x9d, 0x4e, 0x16, 0xab, 0x59, 0x02, 0x65, 0x36, 0x55, 0x6e, 0x6d, 0x3d,
	0x7d, 0x1a, 0x4b, 0x5e, 0x41, 0x73, 0x24, 0x90, 0x2a, 0x1c, 0xe8, 0xe3, 0x70, 0xeb, 0x86, 0xda,
	0x59, 0xa2, 0x9e, 0x16, 0x97, 0x13, 0x81, 0x85, 0xeb, 0x80, 0x26, 0xdb, 0xe6, 0x96, 0xbc, 0x79,
	0x33, 0xd9, 0xc2, 0x75, 0xc0, 0xff, 0xe9, 0x80, 0xfb, 0x0e, 0xe5, 0x48, 0x24, 0x43, 0x2c, 0xb6,
	0x24, 0x8b, 0xcb, 0xf1, 0x96, 0xb6, 0x55, 0x3d, 0x68, 0xfc, 0xb5, 0x8f, 0xdd, 0xf2, 0x3e, 0x74,
	0xb2, 0x70, 0xbc, 0xbd, 0x70, 0x5c, 0x87, 0x73, 0x4f, 0xef, 0x96, 0x3c, 0x35, 0xe8, 0xdc, 0xb5,
	0x36, 0xd4, 0xa6, 0xc9, 0x2c, 0xb1, 0xa6, 0xb5, 0x22, 0xfb, 0xa0, 0xd1, 0x6c, 0x32, 0x91, 0xa8,
	0x8c, 0x21, 0xad, 0x28, 0x7f, 0xf2, 0xcf, 0x61, 0xef, 0x1f, 0x92, 0xf3, 0x0b, 0xdd, 0x87, 0xa6,
	0x62, 0x8a, 0x4e, 0x07, 0x23, 0x96, 0xa5, 0xca, 0x1c, 0x58, 0x2b, 0x02, 0x13, 0x7a, 0xab, 0x23,
	0xe4, 0x05, 0xb4, 0x16, 0x43, 0xe9, 0xe2, 0x5a, 0xf8, 0xaa, 0x33, 0x6e, 0x16, 0xb3, 0x9e, 0xa0,
	0xea, 0xfd, 0xa8, 0x40, 0x53, 0x67, 0xfa, 0xf6, 0x5b, 0x40, 0x3e, 0x3b, 0xd0, 0xb8, 0x7a, 0x45,
	0x48, 0xb7, 0x5c, 0xe0, 0xfa, 0x4b, 0xd8, 0xb9, 0xb7, 0x22, 0x6b, 0x55, 0xfb, 0xcf, 0x2f, 0x8f,
	0xf7, 0xc8, 0xae, 0x54, 0x54, 0x28, 0x8f, 0x7a, 0xe6, 0x63, 0xe3, 0xe9, 0xd6, 0x9e, 0x56, 0xf9,
	0xe5, 0xd7, 0xef, 0x6f, 0x95, 0xb6, 0x7f, 0x2b, 0x3c, 0x3f, 0x0a, 0x75, 0x50, 0x86, 0x06, 0xf0,
	0xd2, 0x79, 0x44, 0xbe, 0x3a, 0xb0, 0xb3, 0xe4, 0x05, 0x79, 0x50, 0x6e, 0xb6, 0x6a, 0xbb, 0x9d,
	0x87, 0x37, 0xa0, 0x72, 0x69, 0xc1, 0xe5, 0xf1, 0x1d, 0x72, 0x7b, 0x9c, 0xe7, 0x17, 0xaa, 0xa4,
	0x91, 0xb5, 0x4d, 0xfe, 0x2f, 0x64, 0x19, 0x47, 0xe5, 0x9b, 0x8d, 0x8f, 0x15, 0x3e, 0x1c, 0xd6,
	0xcd, 0xdd, 0x3d, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x55, 0x4b, 0x45, 0x3e, 0x05, 0x00,
	0x00,
}
