// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateTaskRequest struct {
	JobId                *wrappers.StringValue `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	NodeId               *wrappers.StringValue `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Target               *wrappers.StringValue `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	TaskAction           *wrappers.StringValue `protobuf:"bytes,4,opt,name=task_action,json=taskAction,proto3" json:"task_action,omitempty"`
	Directive            *wrappers.StringValue `protobuf:"bytes,5,opt,name=directive,proto3" json:"directive,omitempty"`
	FailureAllowed       *wrappers.BoolValue   `protobuf:"bytes,6,opt,name=failure_allowed,json=failureAllowed,proto3" json:"failure_allowed,omitempty"`
	Status               *wrappers.StringValue `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateTaskRequest) Reset()         { *m = CreateTaskRequest{} }
func (m *CreateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTaskRequest) ProtoMessage()    {}
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_0656cacb78bf7829, []int{0}
}
func (m *CreateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskRequest.Unmarshal(m, b)
}
func (m *CreateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskRequest.Marshal(b, m, deterministic)
}
func (dst *CreateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskRequest.Merge(dst, src)
}
func (m *CreateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTaskRequest.Size(m)
}
func (m *CreateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskRequest proto.InternalMessageInfo

func (m *CreateTaskRequest) GetJobId() *wrappers.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *CreateTaskRequest) GetNodeId() *wrappers.StringValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *CreateTaskRequest) GetTarget() *wrappers.StringValue {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *CreateTaskRequest) GetTaskAction() *wrappers.StringValue {
	if m != nil {
		return m.TaskAction
	}
	return nil
}

func (m *CreateTaskRequest) GetDirective() *wrappers.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

func (m *CreateTaskRequest) GetFailureAllowed() *wrappers.BoolValue {
	if m != nil {
		return m.FailureAllowed
	}
	return nil
}

func (m *CreateTaskRequest) GetStatus() *wrappers.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

type CreateTaskResponse struct {
	TaskId               *wrappers.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	JobId                *wrappers.StringValue `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateTaskResponse) Reset()         { *m = CreateTaskResponse{} }
func (m *CreateTaskResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTaskResponse) ProtoMessage()    {}
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_0656cacb78bf7829, []int{1}
}
func (m *CreateTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskResponse.Unmarshal(m, b)
}
func (m *CreateTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskResponse.Marshal(b, m, deterministic)
}
func (dst *CreateTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskResponse.Merge(dst, src)
}
func (m *CreateTaskResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTaskResponse.Size(m)
}
func (m *CreateTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskResponse proto.InternalMessageInfo

func (m *CreateTaskResponse) GetTaskId() *wrappers.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *CreateTaskResponse) GetJobId() *wrappers.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

type RetryTasksRequest struct {
	TaskId               []string `protobuf:"bytes,1,rep,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetryTasksRequest) Reset()         { *m = RetryTasksRequest{} }
func (m *RetryTasksRequest) String() string { return proto.CompactTextString(m) }
func (*RetryTasksRequest) ProtoMessage()    {}
func (*RetryTasksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_0656cacb78bf7829, []int{2}
}
func (m *RetryTasksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetryTasksRequest.Unmarshal(m, b)
}
func (m *RetryTasksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetryTasksRequest.Marshal(b, m, deterministic)
}
func (dst *RetryTasksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryTasksRequest.Merge(dst, src)
}
func (m *RetryTasksRequest) XXX_Size() int {
	return xxx_messageInfo_RetryTasksRequest.Size(m)
}
func (m *RetryTasksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryTasksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RetryTasksRequest proto.InternalMessageInfo

func (m *RetryTasksRequest) GetTaskId() []string {
	if m != nil {
		return m.TaskId
	}
	return nil
}

type RetryTasksResponse struct {
	TaskSet              []*Task  `protobuf:"bytes,1,rep,name=task_set,json=taskSet,proto3" json:"task_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetryTasksResponse) Reset()         { *m = RetryTasksResponse{} }
func (m *RetryTasksResponse) String() string { return proto.CompactTextString(m) }
func (*RetryTasksResponse) ProtoMessage()    {}
func (*RetryTasksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_0656cacb78bf7829, []int{3}
}
func (m *RetryTasksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetryTasksResponse.Unmarshal(m, b)
}
func (m *RetryTasksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetryTasksResponse.Marshal(b, m, deterministic)
}
func (dst *RetryTasksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryTasksResponse.Merge(dst, src)
}
func (m *RetryTasksResponse) XXX_Size() int {
	return xxx_messageInfo_RetryTasksResponse.Size(m)
}
func (m *RetryTasksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryTasksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RetryTasksResponse proto.InternalMessageInfo

func (m *RetryTasksResponse) GetTaskSet() []*Task {
	if m != nil {
		return m.TaskSet
	}
	return nil
}

type Task struct {
	TaskId               *wrappers.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	JobId                *wrappers.StringValue `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	TaskAction           *wrappers.StringValue `protobuf:"bytes,3,opt,name=task_action,json=taskAction,proto3" json:"task_action,omitempty"`
	Status               *wrappers.StringValue `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	ErrorCode            *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	Directive            *wrappers.StringValue `protobuf:"bytes,6,opt,name=directive,proto3" json:"directive,omitempty"`
	Executor             *wrappers.StringValue `protobuf:"bytes,7,opt,name=executor,proto3" json:"executor,omitempty"`
	Owner                *wrappers.StringValue `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
	Target               *wrappers.StringValue `protobuf:"bytes,9,opt,name=target,proto3" json:"target,omitempty"`
	NodeId               *wrappers.StringValue `protobuf:"bytes,10,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	CreateTime           *timestamp.Timestamp  `protobuf:"bytes,11,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	StatusTime           *timestamp.Timestamp  `protobuf:"bytes,12,opt,name=status_time,json=statusTime,proto3" json:"status_time,omitempty"`
	FailureAllowed       *wrappers.BoolValue   `protobuf:"bytes,13,opt,name=failure_allowed,json=failureAllowed,proto3" json:"failure_allowed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_0656cacb78bf7829, []int{4}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (dst *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(dst, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetTaskId() *wrappers.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *Task) GetJobId() *wrappers.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *Task) GetTaskAction() *wrappers.StringValue {
	if m != nil {
		return m.TaskAction
	}
	return nil
}

func (m *Task) GetStatus() *wrappers.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Task) GetErrorCode() *wrappers.UInt32Value {
	if m != nil {
		return m.ErrorCode
	}
	return nil
}

func (m *Task) GetDirective() *wrappers.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

func (m *Task) GetExecutor() *wrappers.StringValue {
	if m != nil {
		return m.Executor
	}
	return nil
}

func (m *Task) GetOwner() *wrappers.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Task) GetTarget() *wrappers.StringValue {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Task) GetNodeId() *wrappers.StringValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *Task) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Task) GetStatusTime() *timestamp.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

func (m *Task) GetFailureAllowed() *wrappers.BoolValue {
	if m != nil {
		return m.FailureAllowed
	}
	return nil
}

type DescribeTasksRequest struct {
	TaskId   []string              `protobuf:"bytes,1,rep,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	JobId    []string              `protobuf:"bytes,2,rep,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Executor *wrappers.StringValue `protobuf:"bytes,3,opt,name=executor,proto3" json:"executor,omitempty"`
	Target   *wrappers.StringValue `protobuf:"bytes,4,opt,name=target,proto3" json:"target,omitempty"`
	Status   []string              `protobuf:"bytes,5,rep,name=status,proto3" json:"status,omitempty"`
	// default is 20, max value is 200
	Limit uint32 `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	// default is 0
	Offset               uint32   `protobuf:"varint,7,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescribeTasksRequest) Reset()         { *m = DescribeTasksRequest{} }
func (m *DescribeTasksRequest) String() string { return proto.CompactTextString(m) }
func (*DescribeTasksRequest) ProtoMessage()    {}
func (*DescribeTasksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_0656cacb78bf7829, []int{5}
}
func (m *DescribeTasksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeTasksRequest.Unmarshal(m, b)
}
func (m *DescribeTasksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeTasksRequest.Marshal(b, m, deterministic)
}
func (dst *DescribeTasksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeTasksRequest.Merge(dst, src)
}
func (m *DescribeTasksRequest) XXX_Size() int {
	return xxx_messageInfo_DescribeTasksRequest.Size(m)
}
func (m *DescribeTasksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeTasksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeTasksRequest proto.InternalMessageInfo

func (m *DescribeTasksRequest) GetTaskId() []string {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *DescribeTasksRequest) GetJobId() []string {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *DescribeTasksRequest) GetExecutor() *wrappers.StringValue {
	if m != nil {
		return m.Executor
	}
	return nil
}

func (m *DescribeTasksRequest) GetTarget() *wrappers.StringValue {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *DescribeTasksRequest) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeTasksRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeTasksRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type DescribeTasksResponse struct {
	TotalCount           uint32   `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	TaskSet              []*Task  `protobuf:"bytes,2,rep,name=task_set,json=taskSet,proto3" json:"task_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescribeTasksResponse) Reset()         { *m = DescribeTasksResponse{} }
func (m *DescribeTasksResponse) String() string { return proto.CompactTextString(m) }
func (*DescribeTasksResponse) ProtoMessage()    {}
func (*DescribeTasksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_0656cacb78bf7829, []int{6}
}
func (m *DescribeTasksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeTasksResponse.Unmarshal(m, b)
}
func (m *DescribeTasksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeTasksResponse.Marshal(b, m, deterministic)
}
func (dst *DescribeTasksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeTasksResponse.Merge(dst, src)
}
func (m *DescribeTasksResponse) XXX_Size() int {
	return xxx_messageInfo_DescribeTasksResponse.Size(m)
}
func (m *DescribeTasksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeTasksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeTasksResponse proto.InternalMessageInfo

func (m *DescribeTasksResponse) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *DescribeTasksResponse) GetTaskSet() []*Task {
	if m != nil {
		return m.TaskSet
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTaskRequest)(nil), "openpitrix.CreateTaskRequest")
	proto.RegisterType((*CreateTaskResponse)(nil), "openpitrix.CreateTaskResponse")
	proto.RegisterType((*RetryTasksRequest)(nil), "openpitrix.RetryTasksRequest")
	proto.RegisterType((*RetryTasksResponse)(nil), "openpitrix.RetryTasksResponse")
	proto.RegisterType((*Task)(nil), "openpitrix.Task")
	proto.RegisterType((*DescribeTasksRequest)(nil), "openpitrix.DescribeTasksRequest")
	proto.RegisterType((*DescribeTasksResponse)(nil), "openpitrix.DescribeTasksResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TaskManagerClient is the client API for TaskManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskManagerClient interface {
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	DescribeTasks(ctx context.Context, in *DescribeTasksRequest, opts ...grpc.CallOption) (*DescribeTasksResponse, error)
	RetryTasks(ctx context.Context, in *RetryTasksRequest, opts ...grpc.CallOption) (*RetryTasksResponse, error)
}

type taskManagerClient struct {
	cc *grpc.ClientConn
}

func NewTaskManagerClient(cc *grpc.ClientConn) TaskManagerClient {
	return &taskManagerClient{cc}
}

func (c *taskManagerClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.TaskManager/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerClient) DescribeTasks(ctx context.Context, in *DescribeTasksRequest, opts ...grpc.CallOption) (*DescribeTasksResponse, error) {
	out := new(DescribeTasksResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.TaskManager/DescribeTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerClient) RetryTasks(ctx context.Context, in *RetryTasksRequest, opts ...grpc.CallOption) (*RetryTasksResponse, error) {
	out := new(RetryTasksResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.TaskManager/RetryTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskManagerServer is the server API for TaskManager service.
type TaskManagerServer interface {
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	DescribeTasks(context.Context, *DescribeTasksRequest) (*DescribeTasksResponse, error)
	RetryTasks(context.Context, *RetryTasksRequest) (*RetryTasksResponse, error)
}

func RegisterTaskManagerServer(s *grpc.Server, srv TaskManagerServer) {
	s.RegisterService(&_TaskManager_serviceDesc, srv)
}

func _TaskManager_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.TaskManager/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManager_DescribeTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).DescribeTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.TaskManager/DescribeTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).DescribeTasks(ctx, req.(*DescribeTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManager_RetryTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetryTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).RetryTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.TaskManager/RetryTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).RetryTasks(ctx, req.(*RetryTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.TaskManager",
	HandlerType: (*TaskManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskManager_CreateTask_Handler,
		},
		{
			MethodName: "DescribeTasks",
			Handler:    _TaskManager_DescribeTasks_Handler,
		},
		{
			MethodName: "RetryTasks",
			Handler:    _TaskManager_RetryTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task.proto",
}

func init() { proto.RegisterFile("task.proto", fileDescriptor_task_0656cacb78bf7829) }

var fileDescriptor_task_0656cacb78bf7829 = []byte{
	// 751 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0x3e, 0xdb, 0x8c, 0x09, 0xa5, 0xab, 0x16, 0xac, 0xa8, 0xb4, 0xc1, 0xa7, 0xaa, 0xb4,
	0x09, 0xa4, 0x20, 0xa1, 0x56, 0x1c, 0xd2, 0x70, 0x89, 0x10, 0x97, 0xb4, 0x70, 0xe0, 0x12, 0x6d,
	0xec, 0x89, 0xd9, 0xd6, 0xf5, 0x9a, 0xf5, 0xa6, 0x29, 0x37, 0xc4, 0x81, 0x1f, 0x50, 0x7e, 0x06,
	0x3f, 0x87, 0x2b, 0x47, 0x7e, 0x05, 0xe2, 0x80, 0x76, 0xed, 0x34, 0x4e, 0xd3, 0x52, 0xfb, 0xc2,
	0x29, 0xf2, 0xce, 0x7b, 0xbb, 0x3b, 0xfb, 0xde, 0xbc, 0x00, 0x48, 0x1a, 0x9e, 0x34, 0x02, 0xc1,
	0x25, 0x27, 0xc0, 0x03, 0xf4, 0x03, 0x26, 0x05, 0x3b, 0xaf, 0xad, 0xbb, 0x9c, 0xbb, 0x1e, 0x36,
	0x75, 0x65, 0x30, 0x1a, 0x36, 0xc7, 0x82, 0x06, 0x01, 0x8a, 0x30, 0xc2, 0xd6, 0x36, 0xae, 0xd6,
	0x25, 0x3b, 0xc5, 0x50, 0xd2, 0xd3, 0x20, 0x06, 0xac, 0xc5, 0x00, 0x1a, 0xb0, 0x26, 0xf5, 0x7d,
	0x2e, 0xa9, 0x64, 0xdc, 0x9f, 0xd0, 0xb7, 0xf5, 0x8f, 0xbd, 0xe3, 0xa2, 0xbf, 0x13, 0x8e, 0xa9,
	0xeb, 0xa2, 0x68, 0xf2, 0x40, 0x23, 0xe6, 0xd1, 0xd6, 0xf7, 0x02, 0x2c, 0x77, 0x04, 0x52, 0x89,
	0x47, 0x34, 0x3c, 0xe9, 0xe1, 0xc7, 0x11, 0x86, 0x92, 0xec, 0x42, 0xf9, 0x98, 0x0f, 0xfa, 0xcc,
	0x31, 0x73, 0xf5, 0xdc, 0xa6, 0xd1, 0x5a, 0x6b, 0x44, 0x47, 0x36, 0x26, 0x77, 0x6a, 0x1c, 0x4a,
	0xc1, 0x7c, 0xf7, 0x1d, 0xf5, 0x46, 0xd8, 0x2b, 0x1d, 0xf3, 0x41, 0xd7, 0x21, 0xcf, 0x61, 0xc1,
	0xe7, 0x0e, 0x2a, 0x56, 0x3e, 0x05, 0xab, 0xac, 0xc0, 0x5d, 0x87, 0x3c, 0x83, 0xb2, 0xa4, 0xc2,
	0x45, 0x69, 0x16, 0xd2, 0xb0, 0x22, 0x2c, 0x79, 0x09, 0x86, 0x7a, 0xde, 0x3e, 0xb5, 0x55, 0x37,
	0x66, 0x31, 0x05, 0x55, 0xeb, 0xd1, 0xd6, 0x78, 0xb2, 0x07, 0x15, 0x87, 0x09, 0xb4, 0x25, 0x3b,
	0x43, 0xb3, 0x94, 0x82, 0x3c, 0x85, 0x93, 0x0e, 0x2c, 0x0d, 0x29, 0xf3, 0x46, 0x02, 0xfb, 0xd4,
	0xf3, 0xf8, 0x18, 0x1d, 0xb3, 0xac, 0x77, 0xa8, 0xcd, 0xed, 0x70, 0xc0, 0xb9, 0x17, 0xf1, 0xef,
	0xc6, 0x94, 0x76, 0xc4, 0x50, 0x5d, 0x87, 0x92, 0xca, 0x51, 0x68, 0x2e, 0xa4, 0xe9, 0x3a, 0xc2,
	0x5a, 0x9f, 0x73, 0x40, 0x92, 0x6a, 0x85, 0x01, 0xf7, 0x43, 0x54, 0x2f, 0xaf, 0x1f, 0x23, 0xa5,
	0x5e, 0x65, 0x05, 0xee, 0x3a, 0x09, 0x95, 0xf3, 0xa9, 0x55, 0xb6, 0xb6, 0x61, 0xb9, 0x87, 0x52,
	0x7c, 0x52, 0x17, 0x08, 0x27, 0x7e, 0x79, 0x90, 0xbc, 0x40, 0x61, 0xb3, 0x32, 0x39, 0xc2, 0x6a,
	0x03, 0x49, 0xa2, 0xe3, 0xfb, 0x3e, 0x86, 0x45, 0x0d, 0x0f, 0x51, 0x6a, 0xbc, 0xd1, 0xba, 0xd7,
	0x98, 0x0e, 0x48, 0x43, 0xf7, 0xa6, 0x37, 0x3c, 0x44, 0x69, 0xfd, 0x2e, 0x41, 0x51, 0xad, 0xfc,
	0xcf, 0x2e, 0xaf, 0xda, 0xab, 0x90, 0xd1, 0x5e, 0x53, 0x75, 0x8b, 0xe9, 0xd5, 0x25, 0xfb, 0x00,
	0x28, 0x04, 0x17, 0x7d, 0x9b, 0x3b, 0x37, 0xbb, 0xf2, 0x6d, 0xd7, 0x97, 0xbb, 0xad, 0xd8, 0x95,
	0x1a, 0xdf, 0xe1, 0x0e, 0xce, 0x3a, 0xba, 0x9c, 0xcd, 0xd1, 0x2f, 0x60, 0x11, 0xcf, 0xd1, 0x1e,
	0x49, 0x2e, 0x52, 0xd9, 0xf1, 0x12, 0x4d, 0x5a, 0x50, 0xe2, 0x63, 0x1f, 0x85, 0xb9, 0x98, 0xe6,
	0x6d, 0x35, 0x34, 0x31, 0xf0, 0x95, 0x0c, 0x03, 0x9f, 0x48, 0x17, 0xc8, 0x90, 0x2e, 0xfb, 0x60,
	0xd8, 0x7a, 0x60, 0xfa, 0x2a, 0x45, 0x4d, 0xe3, 0x86, 0x41, 0x3d, 0x9a, 0x44, 0x6c, 0x0f, 0x22,
	0xb8, 0x5a, 0x50, 0xe4, 0x48, 0x9a, 0x88, 0x7c, 0xe7, 0x76, 0x72, 0x04, 0xd7, 0xe4, 0x6b, 0x62,
	0xa2, 0x9a, 0x35, 0x26, 0xac, 0x3f, 0x39, 0x58, 0x79, 0x85, 0xa1, 0x2d, 0xd8, 0x00, 0x53, 0x4d,
	0x1c, 0x59, 0x4d, 0xd8, 0x5d, 0xad, 0xc7, 0x86, 0x4e, 0x4a, 0x5c, 0xc8, 0x24, 0xf1, 0x54, 0xae,
	0x62, 0x06, 0xb9, 0xee, 0x5f, 0x4e, 0x40, 0x29, 0xba, 0x5e, 0xec, 0xf1, 0x15, 0x28, 0x79, 0xec,
	0x94, 0x49, 0x6d, 0xd1, 0x6a, 0x2f, 0xfa, 0x50, 0x68, 0x3e, 0x1c, 0xaa, 0x38, 0x58, 0xd0, 0xcb,
	0xf1, 0x97, 0x85, 0xb0, 0x7a, 0xa5, 0xfb, 0x38, 0x41, 0x36, 0xc0, 0x90, 0x5c, 0x52, 0xaf, 0x6f,
	0xf3, 0x91, 0x2f, 0x75, 0x1e, 0x54, 0x7b, 0xa0, 0x97, 0x3a, 0x6a, 0x65, 0x26, 0x62, 0xf2, 0xb7,
	0x44, 0x4c, 0xeb, 0x67, 0x1e, 0x0c, 0xb5, 0xf2, 0x86, 0xfa, 0xd4, 0x45, 0x41, 0x5e, 0x03, 0x4c,
	0x53, 0x96, 0x3c, 0x4c, 0x12, 0xe7, 0xfe, 0x2b, 0x6b, 0xeb, 0x37, 0x95, 0xe3, 0xab, 0x7e, 0xcd,
	0x41, 0x75, 0xa6, 0x09, 0x52, 0x4f, 0x32, 0xae, 0x53, 0xb7, 0xf6, 0xe8, 0x1f, 0x88, 0x68, 0x5b,
	0xeb, 0xc9, 0x45, 0x7b, 0x8d, 0xd4, 0x9c, 0xb8, 0x56, 0x57, 0xad, 0x84, 0xf5, 0x31, 0x93, 0x1f,
	0xea, 0x43, 0xe6, 0x49, 0x14, 0x5f, 0x7e, 0xfc, 0xfa, 0x96, 0x37, 0x48, 0xa5, 0x79, 0xf6, 0xb4,
	0xa9, 0x8b, 0x64, 0x0c, 0x30, 0xcd, 0xe2, 0xd9, 0xae, 0xe6, 0x12, 0x7d, 0xb6, 0xab, 0xf9, 0x08,
	0xb7, 0xb6, 0x2e, 0xda, 0x55, 0x62, 0x08, 0x55, 0x88, 0xce, 0xd6, 0xe7, 0xad, 0x58, 0x4b, 0x97,
	0xe7, 0x35, 0x75, 0x71, 0x2f, 0xb7, 0x75, 0x50, 0x7c, 0x9f, 0x0f, 0x06, 0x83, 0xb2, 0xf6, 0xcb,
	0xee, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x46, 0x5e, 0x23, 0xe6, 0x17, 0x09, 0x00, 0x00,
}
