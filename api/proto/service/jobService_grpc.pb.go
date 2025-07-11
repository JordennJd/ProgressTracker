// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: jobService.proto

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	models "progress-tracker/api/proto/models"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	JobService_CreateJob_FullMethodName      = "/jobService.JobService/CreateJob"
	JobService_StartJob_FullMethodName       = "/jobService.JobService/StartJob"
	JobService_CompleteJob_FullMethodName    = "/jobService.JobService/CompleteJob"
	JobService_GetJob_FullMethodName         = "/jobService.JobService/GetJob"
	JobService_GetNextJob_FullMethodName     = "/jobService.JobService/GetNextJob"
	JobService_GetWorkingJobs_FullMethodName = "/jobService.JobService/GetWorkingJobs"
)

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobServiceClient interface {
	CreateJob(ctx context.Context, in *models.CreateJobQuery, opts ...grpc.CallOption) (*models.Result, error)
	StartJob(ctx context.Context, in *models.StartJobQuery, opts ...grpc.CallOption) (*models.Result, error)
	CompleteJob(ctx context.Context, in *models.CompleteJobQuery, opts ...grpc.CallOption) (*models.Result, error)
	GetJob(ctx context.Context, in *models.GetJobQuery, opts ...grpc.CallOption) (*models.GetJobResult, error)
	GetNextJob(ctx context.Context, in *models.GetNextJobRequest, opts ...grpc.CallOption) (*models.GetJobResult, error)
	GetWorkingJobs(ctx context.Context, in *models.GetWorkingJobRequest, opts ...grpc.CallOption) (*models.GetJobsResult, error)
}

type jobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobServiceClient(cc grpc.ClientConnInterface) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) CreateJob(ctx context.Context, in *models.CreateJobQuery, opts ...grpc.CallOption) (*models.Result, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(models.Result)
	err := c.cc.Invoke(ctx, JobService_CreateJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) StartJob(ctx context.Context, in *models.StartJobQuery, opts ...grpc.CallOption) (*models.Result, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(models.Result)
	err := c.cc.Invoke(ctx, JobService_StartJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) CompleteJob(ctx context.Context, in *models.CompleteJobQuery, opts ...grpc.CallOption) (*models.Result, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(models.Result)
	err := c.cc.Invoke(ctx, JobService_CompleteJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetJob(ctx context.Context, in *models.GetJobQuery, opts ...grpc.CallOption) (*models.GetJobResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(models.GetJobResult)
	err := c.cc.Invoke(ctx, JobService_GetJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetNextJob(ctx context.Context, in *models.GetNextJobRequest, opts ...grpc.CallOption) (*models.GetJobResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(models.GetJobResult)
	err := c.cc.Invoke(ctx, JobService_GetNextJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetWorkingJobs(ctx context.Context, in *models.GetWorkingJobRequest, opts ...grpc.CallOption) (*models.GetJobsResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(models.GetJobsResult)
	err := c.cc.Invoke(ctx, JobService_GetWorkingJobs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServiceServer is the server API for JobService service.
// All implementations must embed UnimplementedJobServiceServer
// for forward compatibility.
type JobServiceServer interface {
	CreateJob(context.Context, *models.CreateJobQuery) (*models.Result, error)
	StartJob(context.Context, *models.StartJobQuery) (*models.Result, error)
	CompleteJob(context.Context, *models.CompleteJobQuery) (*models.Result, error)
	GetJob(context.Context, *models.GetJobQuery) (*models.GetJobResult, error)
	GetNextJob(context.Context, *models.GetNextJobRequest) (*models.GetJobResult, error)
	GetWorkingJobs(context.Context, *models.GetWorkingJobRequest) (*models.GetJobsResult, error)
	mustEmbedUnimplementedJobServiceServer()
}

// UnimplementedJobServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedJobServiceServer struct{}

func (UnimplementedJobServiceServer) CreateJob(context.Context, *models.CreateJobQuery) (*models.Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJob not implemented")
}
func (UnimplementedJobServiceServer) StartJob(context.Context, *models.StartJobQuery) (*models.Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartJob not implemented")
}
func (UnimplementedJobServiceServer) CompleteJob(context.Context, *models.CompleteJobQuery) (*models.Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteJob not implemented")
}
func (UnimplementedJobServiceServer) GetJob(context.Context, *models.GetJobQuery) (*models.GetJobResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJob not implemented")
}
func (UnimplementedJobServiceServer) GetNextJob(context.Context, *models.GetNextJobRequest) (*models.GetJobResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNextJob not implemented")
}
func (UnimplementedJobServiceServer) GetWorkingJobs(context.Context, *models.GetWorkingJobRequest) (*models.GetJobsResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkingJobs not implemented")
}
func (UnimplementedJobServiceServer) mustEmbedUnimplementedJobServiceServer() {}
func (UnimplementedJobServiceServer) testEmbeddedByValue()                    {}

// UnsafeJobServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobServiceServer will
// result in compilation errors.
type UnsafeJobServiceServer interface {
	mustEmbedUnimplementedJobServiceServer()
}

func RegisterJobServiceServer(s grpc.ServiceRegistrar, srv JobServiceServer) {
	// If the following call pancis, it indicates UnimplementedJobServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&JobService_ServiceDesc, srv)
}

func _JobService_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.CreateJobQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_CreateJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).CreateJob(ctx, req.(*models.CreateJobQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_StartJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.StartJobQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).StartJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_StartJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).StartJob(ctx, req.(*models.StartJobQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_CompleteJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.CompleteJobQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).CompleteJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_CompleteJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).CompleteJob(ctx, req.(*models.CompleteJobQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.GetJobQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_GetJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetJob(ctx, req.(*models.GetJobQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetNextJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.GetNextJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetNextJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_GetNextJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetNextJob(ctx, req.(*models.GetNextJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetWorkingJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.GetWorkingJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetWorkingJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_GetWorkingJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetWorkingJobs(ctx, req.(*models.GetWorkingJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JobService_ServiceDesc is the grpc.ServiceDesc for JobService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jobService.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _JobService_CreateJob_Handler,
		},
		{
			MethodName: "StartJob",
			Handler:    _JobService_StartJob_Handler,
		},
		{
			MethodName: "CompleteJob",
			Handler:    _JobService_CompleteJob_Handler,
		},
		{
			MethodName: "GetJob",
			Handler:    _JobService_GetJob_Handler,
		},
		{
			MethodName: "GetNextJob",
			Handler:    _JobService_GetNextJob_Handler,
		},
		{
			MethodName: "GetWorkingJobs",
			Handler:    _JobService_GetWorkingJobs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jobService.proto",
}
