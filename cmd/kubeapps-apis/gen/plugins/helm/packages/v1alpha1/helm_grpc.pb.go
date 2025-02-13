// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: kubeappsapis/plugins/helm/packages/v1alpha1/helm.proto

package v1alpha1

import (
	context "context"
	v1alpha1 "github.com/vmware-tanzu/kubeapps/cmd/kubeapps-apis/gen/core/packages/v1alpha1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HelmPackagesServiceClient is the client API for HelmPackagesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelmPackagesServiceClient interface {
	// GetAvailablePackageSummaries returns the available packages managed by the 'helm' plugin
	GetAvailablePackageSummaries(ctx context.Context, in *v1alpha1.GetAvailablePackageSummariesRequest, opts ...grpc.CallOption) (*v1alpha1.GetAvailablePackageSummariesResponse, error)
	// GetAvailablePackageDetail returns the package details managed by the 'helm' plugin
	GetAvailablePackageDetail(ctx context.Context, in *v1alpha1.GetAvailablePackageDetailRequest, opts ...grpc.CallOption) (*v1alpha1.GetAvailablePackageDetailResponse, error)
	// GetAvailablePackageVersions returns the package versions managed by the 'helm' plugin
	GetAvailablePackageVersions(ctx context.Context, in *v1alpha1.GetAvailablePackageVersionsRequest, opts ...grpc.CallOption) (*v1alpha1.GetAvailablePackageVersionsResponse, error)
	// GetInstalledPackageSummaries returns the installed packages managed by the 'helm' plugin
	GetInstalledPackageSummaries(ctx context.Context, in *v1alpha1.GetInstalledPackageSummariesRequest, opts ...grpc.CallOption) (*v1alpha1.GetInstalledPackageSummariesResponse, error)
	// GetInstalledPackageDetail returns the requested installed package managed by the 'helm' plugin
	GetInstalledPackageDetail(ctx context.Context, in *v1alpha1.GetInstalledPackageDetailRequest, opts ...grpc.CallOption) (*v1alpha1.GetInstalledPackageDetailResponse, error)
	// CreateInstalledPackage creates an installed package based on the request.
	CreateInstalledPackage(ctx context.Context, in *v1alpha1.CreateInstalledPackageRequest, opts ...grpc.CallOption) (*v1alpha1.CreateInstalledPackageResponse, error)
	// UpdateInstalledPackage updates an installed package based on the request.
	UpdateInstalledPackage(ctx context.Context, in *v1alpha1.UpdateInstalledPackageRequest, opts ...grpc.CallOption) (*v1alpha1.UpdateInstalledPackageResponse, error)
	// DeleteInstalledPackage deletes an installed package based on the request.
	DeleteInstalledPackage(ctx context.Context, in *v1alpha1.DeleteInstalledPackageRequest, opts ...grpc.CallOption) (*v1alpha1.DeleteInstalledPackageResponse, error)
	// RollbackInstalledPackage updates an installed package based on the request.
	RollbackInstalledPackage(ctx context.Context, in *RollbackInstalledPackageRequest, opts ...grpc.CallOption) (*RollbackInstalledPackageResponse, error)
	// GetInstalledPackageResourceRefs returns the references for the Kubernetes resources created by
	// an installed package.
	GetInstalledPackageResourceRefs(ctx context.Context, in *v1alpha1.GetInstalledPackageResourceRefsRequest, opts ...grpc.CallOption) (*v1alpha1.GetInstalledPackageResourceRefsResponse, error)
}

type helmPackagesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelmPackagesServiceClient(cc grpc.ClientConnInterface) HelmPackagesServiceClient {
	return &helmPackagesServiceClient{cc}
}

func (c *helmPackagesServiceClient) GetAvailablePackageSummaries(ctx context.Context, in *v1alpha1.GetAvailablePackageSummariesRequest, opts ...grpc.CallOption) (*v1alpha1.GetAvailablePackageSummariesResponse, error) {
	out := new(v1alpha1.GetAvailablePackageSummariesResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmPackagesService/GetAvailablePackageSummaries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helmPackagesServiceClient) GetAvailablePackageDetail(ctx context.Context, in *v1alpha1.GetAvailablePackageDetailRequest, opts ...grpc.CallOption) (*v1alpha1.GetAvailablePackageDetailResponse, error) {
	out := new(v1alpha1.GetAvailablePackageDetailResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmPackagesService/GetAvailablePackageDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helmPackagesServiceClient) GetAvailablePackageVersions(ctx context.Context, in *v1alpha1.GetAvailablePackageVersionsRequest, opts ...grpc.CallOption) (*v1alpha1.GetAvailablePackageVersionsResponse, error) {
	out := new(v1alpha1.GetAvailablePackageVersionsResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmPackagesService/GetAvailablePackageVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helmPackagesServiceClient) GetInstalledPackageSummaries(ctx context.Context, in *v1alpha1.GetInstalledPackageSummariesRequest, opts ...grpc.CallOption) (*v1alpha1.GetInstalledPackageSummariesResponse, error) {
	out := new(v1alpha1.GetInstalledPackageSummariesResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmPackagesService/GetInstalledPackageSummaries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helmPackagesServiceClient) GetInstalledPackageDetail(ctx context.Context, in *v1alpha1.GetInstalledPackageDetailRequest, opts ...grpc.CallOption) (*v1alpha1.GetInstalledPackageDetailResponse, error) {
	out := new(v1alpha1.GetInstalledPackageDetailResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmPackagesService/GetInstalledPackageDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helmPackagesServiceClient) CreateInstalledPackage(ctx context.Context, in *v1alpha1.CreateInstalledPackageRequest, opts ...grpc.CallOption) (*v1alpha1.CreateInstalledPackageResponse, error) {
	out := new(v1alpha1.CreateInstalledPackageResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmPackagesService/CreateInstalledPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helmPackagesServiceClient) UpdateInstalledPackage(ctx context.Context, in *v1alpha1.UpdateInstalledPackageRequest, opts ...grpc.CallOption) (*v1alpha1.UpdateInstalledPackageResponse, error) {
	out := new(v1alpha1.UpdateInstalledPackageResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmPackagesService/UpdateInstalledPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helmPackagesServiceClient) DeleteInstalledPackage(ctx context.Context, in *v1alpha1.DeleteInstalledPackageRequest, opts ...grpc.CallOption) (*v1alpha1.DeleteInstalledPackageResponse, error) {
	out := new(v1alpha1.DeleteInstalledPackageResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmPackagesService/DeleteInstalledPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helmPackagesServiceClient) RollbackInstalledPackage(ctx context.Context, in *RollbackInstalledPackageRequest, opts ...grpc.CallOption) (*RollbackInstalledPackageResponse, error) {
	out := new(RollbackInstalledPackageResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmPackagesService/RollbackInstalledPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helmPackagesServiceClient) GetInstalledPackageResourceRefs(ctx context.Context, in *v1alpha1.GetInstalledPackageResourceRefsRequest, opts ...grpc.CallOption) (*v1alpha1.GetInstalledPackageResourceRefsResponse, error) {
	out := new(v1alpha1.GetInstalledPackageResourceRefsResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmPackagesService/GetInstalledPackageResourceRefs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelmPackagesServiceServer is the server API for HelmPackagesService service.
// All implementations should embed UnimplementedHelmPackagesServiceServer
// for forward compatibility
type HelmPackagesServiceServer interface {
	// GetAvailablePackageSummaries returns the available packages managed by the 'helm' plugin
	GetAvailablePackageSummaries(context.Context, *v1alpha1.GetAvailablePackageSummariesRequest) (*v1alpha1.GetAvailablePackageSummariesResponse, error)
	// GetAvailablePackageDetail returns the package details managed by the 'helm' plugin
	GetAvailablePackageDetail(context.Context, *v1alpha1.GetAvailablePackageDetailRequest) (*v1alpha1.GetAvailablePackageDetailResponse, error)
	// GetAvailablePackageVersions returns the package versions managed by the 'helm' plugin
	GetAvailablePackageVersions(context.Context, *v1alpha1.GetAvailablePackageVersionsRequest) (*v1alpha1.GetAvailablePackageVersionsResponse, error)
	// GetInstalledPackageSummaries returns the installed packages managed by the 'helm' plugin
	GetInstalledPackageSummaries(context.Context, *v1alpha1.GetInstalledPackageSummariesRequest) (*v1alpha1.GetInstalledPackageSummariesResponse, error)
	// GetInstalledPackageDetail returns the requested installed package managed by the 'helm' plugin
	GetInstalledPackageDetail(context.Context, *v1alpha1.GetInstalledPackageDetailRequest) (*v1alpha1.GetInstalledPackageDetailResponse, error)
	// CreateInstalledPackage creates an installed package based on the request.
	CreateInstalledPackage(context.Context, *v1alpha1.CreateInstalledPackageRequest) (*v1alpha1.CreateInstalledPackageResponse, error)
	// UpdateInstalledPackage updates an installed package based on the request.
	UpdateInstalledPackage(context.Context, *v1alpha1.UpdateInstalledPackageRequest) (*v1alpha1.UpdateInstalledPackageResponse, error)
	// DeleteInstalledPackage deletes an installed package based on the request.
	DeleteInstalledPackage(context.Context, *v1alpha1.DeleteInstalledPackageRequest) (*v1alpha1.DeleteInstalledPackageResponse, error)
	// RollbackInstalledPackage updates an installed package based on the request.
	RollbackInstalledPackage(context.Context, *RollbackInstalledPackageRequest) (*RollbackInstalledPackageResponse, error)
	// GetInstalledPackageResourceRefs returns the references for the Kubernetes resources created by
	// an installed package.
	GetInstalledPackageResourceRefs(context.Context, *v1alpha1.GetInstalledPackageResourceRefsRequest) (*v1alpha1.GetInstalledPackageResourceRefsResponse, error)
}

// UnimplementedHelmPackagesServiceServer should be embedded to have forward compatible implementations.
type UnimplementedHelmPackagesServiceServer struct {
}

func (UnimplementedHelmPackagesServiceServer) GetAvailablePackageSummaries(context.Context, *v1alpha1.GetAvailablePackageSummariesRequest) (*v1alpha1.GetAvailablePackageSummariesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailablePackageSummaries not implemented")
}
func (UnimplementedHelmPackagesServiceServer) GetAvailablePackageDetail(context.Context, *v1alpha1.GetAvailablePackageDetailRequest) (*v1alpha1.GetAvailablePackageDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailablePackageDetail not implemented")
}
func (UnimplementedHelmPackagesServiceServer) GetAvailablePackageVersions(context.Context, *v1alpha1.GetAvailablePackageVersionsRequest) (*v1alpha1.GetAvailablePackageVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailablePackageVersions not implemented")
}
func (UnimplementedHelmPackagesServiceServer) GetInstalledPackageSummaries(context.Context, *v1alpha1.GetInstalledPackageSummariesRequest) (*v1alpha1.GetInstalledPackageSummariesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstalledPackageSummaries not implemented")
}
func (UnimplementedHelmPackagesServiceServer) GetInstalledPackageDetail(context.Context, *v1alpha1.GetInstalledPackageDetailRequest) (*v1alpha1.GetInstalledPackageDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstalledPackageDetail not implemented")
}
func (UnimplementedHelmPackagesServiceServer) CreateInstalledPackage(context.Context, *v1alpha1.CreateInstalledPackageRequest) (*v1alpha1.CreateInstalledPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInstalledPackage not implemented")
}
func (UnimplementedHelmPackagesServiceServer) UpdateInstalledPackage(context.Context, *v1alpha1.UpdateInstalledPackageRequest) (*v1alpha1.UpdateInstalledPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInstalledPackage not implemented")
}
func (UnimplementedHelmPackagesServiceServer) DeleteInstalledPackage(context.Context, *v1alpha1.DeleteInstalledPackageRequest) (*v1alpha1.DeleteInstalledPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInstalledPackage not implemented")
}
func (UnimplementedHelmPackagesServiceServer) RollbackInstalledPackage(context.Context, *RollbackInstalledPackageRequest) (*RollbackInstalledPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackInstalledPackage not implemented")
}
func (UnimplementedHelmPackagesServiceServer) GetInstalledPackageResourceRefs(context.Context, *v1alpha1.GetInstalledPackageResourceRefsRequest) (*v1alpha1.GetInstalledPackageResourceRefsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstalledPackageResourceRefs not implemented")
}

// UnsafeHelmPackagesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelmPackagesServiceServer will
// result in compilation errors.
type UnsafeHelmPackagesServiceServer interface {
	mustEmbedUnimplementedHelmPackagesServiceServer()
}

func RegisterHelmPackagesServiceServer(s grpc.ServiceRegistrar, srv HelmPackagesServiceServer) {
	s.RegisterService(&HelmPackagesService_ServiceDesc, srv)
}

func _HelmPackagesService_GetAvailablePackageSummaries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.GetAvailablePackageSummariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmPackagesServiceServer).GetAvailablePackageSummaries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmPackagesService/GetAvailablePackageSummaries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmPackagesServiceServer).GetAvailablePackageSummaries(ctx, req.(*v1alpha1.GetAvailablePackageSummariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelmPackagesService_GetAvailablePackageDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.GetAvailablePackageDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmPackagesServiceServer).GetAvailablePackageDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmPackagesService/GetAvailablePackageDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmPackagesServiceServer).GetAvailablePackageDetail(ctx, req.(*v1alpha1.GetAvailablePackageDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelmPackagesService_GetAvailablePackageVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.GetAvailablePackageVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmPackagesServiceServer).GetAvailablePackageVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmPackagesService/GetAvailablePackageVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmPackagesServiceServer).GetAvailablePackageVersions(ctx, req.(*v1alpha1.GetAvailablePackageVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelmPackagesService_GetInstalledPackageSummaries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.GetInstalledPackageSummariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmPackagesServiceServer).GetInstalledPackageSummaries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmPackagesService/GetInstalledPackageSummaries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmPackagesServiceServer).GetInstalledPackageSummaries(ctx, req.(*v1alpha1.GetInstalledPackageSummariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelmPackagesService_GetInstalledPackageDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.GetInstalledPackageDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmPackagesServiceServer).GetInstalledPackageDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmPackagesService/GetInstalledPackageDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmPackagesServiceServer).GetInstalledPackageDetail(ctx, req.(*v1alpha1.GetInstalledPackageDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelmPackagesService_CreateInstalledPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.CreateInstalledPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmPackagesServiceServer).CreateInstalledPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmPackagesService/CreateInstalledPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmPackagesServiceServer).CreateInstalledPackage(ctx, req.(*v1alpha1.CreateInstalledPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelmPackagesService_UpdateInstalledPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.UpdateInstalledPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmPackagesServiceServer).UpdateInstalledPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmPackagesService/UpdateInstalledPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmPackagesServiceServer).UpdateInstalledPackage(ctx, req.(*v1alpha1.UpdateInstalledPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelmPackagesService_DeleteInstalledPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.DeleteInstalledPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmPackagesServiceServer).DeleteInstalledPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmPackagesService/DeleteInstalledPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmPackagesServiceServer).DeleteInstalledPackage(ctx, req.(*v1alpha1.DeleteInstalledPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelmPackagesService_RollbackInstalledPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackInstalledPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmPackagesServiceServer).RollbackInstalledPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmPackagesService/RollbackInstalledPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmPackagesServiceServer).RollbackInstalledPackage(ctx, req.(*RollbackInstalledPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelmPackagesService_GetInstalledPackageResourceRefs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.GetInstalledPackageResourceRefsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmPackagesServiceServer).GetInstalledPackageResourceRefs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmPackagesService/GetInstalledPackageResourceRefs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmPackagesServiceServer).GetInstalledPackageResourceRefs(ctx, req.(*v1alpha1.GetInstalledPackageResourceRefsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HelmPackagesService_ServiceDesc is the grpc.ServiceDesc for HelmPackagesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelmPackagesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kubeappsapis.plugins.helm.packages.v1alpha1.HelmPackagesService",
	HandlerType: (*HelmPackagesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAvailablePackageSummaries",
			Handler:    _HelmPackagesService_GetAvailablePackageSummaries_Handler,
		},
		{
			MethodName: "GetAvailablePackageDetail",
			Handler:    _HelmPackagesService_GetAvailablePackageDetail_Handler,
		},
		{
			MethodName: "GetAvailablePackageVersions",
			Handler:    _HelmPackagesService_GetAvailablePackageVersions_Handler,
		},
		{
			MethodName: "GetInstalledPackageSummaries",
			Handler:    _HelmPackagesService_GetInstalledPackageSummaries_Handler,
		},
		{
			MethodName: "GetInstalledPackageDetail",
			Handler:    _HelmPackagesService_GetInstalledPackageDetail_Handler,
		},
		{
			MethodName: "CreateInstalledPackage",
			Handler:    _HelmPackagesService_CreateInstalledPackage_Handler,
		},
		{
			MethodName: "UpdateInstalledPackage",
			Handler:    _HelmPackagesService_UpdateInstalledPackage_Handler,
		},
		{
			MethodName: "DeleteInstalledPackage",
			Handler:    _HelmPackagesService_DeleteInstalledPackage_Handler,
		},
		{
			MethodName: "RollbackInstalledPackage",
			Handler:    _HelmPackagesService_RollbackInstalledPackage_Handler,
		},
		{
			MethodName: "GetInstalledPackageResourceRefs",
			Handler:    _HelmPackagesService_GetInstalledPackageResourceRefs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kubeappsapis/plugins/helm/packages/v1alpha1/helm.proto",
}

// HelmRepositoriesServiceClient is the client API for HelmRepositoriesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelmRepositoriesServiceClient interface {
	// AddPackageRepository add an existing package repository to the set of ones already managed by the Helm plugin
	AddPackageRepository(ctx context.Context, in *v1alpha1.AddPackageRepositoryRequest, opts ...grpc.CallOption) (*v1alpha1.AddPackageRepositoryResponse, error)
	GetPackageRepositoryDetail(ctx context.Context, in *v1alpha1.GetPackageRepositoryDetailRequest, opts ...grpc.CallOption) (*v1alpha1.GetPackageRepositoryDetailResponse, error)
	GetPackageRepositorySummaries(ctx context.Context, in *v1alpha1.GetPackageRepositorySummariesRequest, opts ...grpc.CallOption) (*v1alpha1.GetPackageRepositorySummariesResponse, error)
	UpdatePackageRepository(ctx context.Context, in *v1alpha1.UpdatePackageRepositoryRequest, opts ...grpc.CallOption) (*v1alpha1.UpdatePackageRepositoryResponse, error)
	DeletePackageRepository(ctx context.Context, in *v1alpha1.DeletePackageRepositoryRequest, opts ...grpc.CallOption) (*v1alpha1.DeletePackageRepositoryResponse, error)
	GetPackageRepositoryPermissions(ctx context.Context, in *v1alpha1.GetPackageRepositoryPermissionsRequest, opts ...grpc.CallOption) (*v1alpha1.GetPackageRepositoryPermissionsResponse, error)
}

type helmRepositoriesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelmRepositoriesServiceClient(cc grpc.ClientConnInterface) HelmRepositoriesServiceClient {
	return &helmRepositoriesServiceClient{cc}
}

func (c *helmRepositoriesServiceClient) AddPackageRepository(ctx context.Context, in *v1alpha1.AddPackageRepositoryRequest, opts ...grpc.CallOption) (*v1alpha1.AddPackageRepositoryResponse, error) {
	out := new(v1alpha1.AddPackageRepositoryResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmRepositoriesService/AddPackageRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helmRepositoriesServiceClient) GetPackageRepositoryDetail(ctx context.Context, in *v1alpha1.GetPackageRepositoryDetailRequest, opts ...grpc.CallOption) (*v1alpha1.GetPackageRepositoryDetailResponse, error) {
	out := new(v1alpha1.GetPackageRepositoryDetailResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmRepositoriesService/GetPackageRepositoryDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helmRepositoriesServiceClient) GetPackageRepositorySummaries(ctx context.Context, in *v1alpha1.GetPackageRepositorySummariesRequest, opts ...grpc.CallOption) (*v1alpha1.GetPackageRepositorySummariesResponse, error) {
	out := new(v1alpha1.GetPackageRepositorySummariesResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmRepositoriesService/GetPackageRepositorySummaries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helmRepositoriesServiceClient) UpdatePackageRepository(ctx context.Context, in *v1alpha1.UpdatePackageRepositoryRequest, opts ...grpc.CallOption) (*v1alpha1.UpdatePackageRepositoryResponse, error) {
	out := new(v1alpha1.UpdatePackageRepositoryResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmRepositoriesService/UpdatePackageRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helmRepositoriesServiceClient) DeletePackageRepository(ctx context.Context, in *v1alpha1.DeletePackageRepositoryRequest, opts ...grpc.CallOption) (*v1alpha1.DeletePackageRepositoryResponse, error) {
	out := new(v1alpha1.DeletePackageRepositoryResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmRepositoriesService/DeletePackageRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helmRepositoriesServiceClient) GetPackageRepositoryPermissions(ctx context.Context, in *v1alpha1.GetPackageRepositoryPermissionsRequest, opts ...grpc.CallOption) (*v1alpha1.GetPackageRepositoryPermissionsResponse, error) {
	out := new(v1alpha1.GetPackageRepositoryPermissionsResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmRepositoriesService/GetPackageRepositoryPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelmRepositoriesServiceServer is the server API for HelmRepositoriesService service.
// All implementations should embed UnimplementedHelmRepositoriesServiceServer
// for forward compatibility
type HelmRepositoriesServiceServer interface {
	// AddPackageRepository add an existing package repository to the set of ones already managed by the Helm plugin
	AddPackageRepository(context.Context, *v1alpha1.AddPackageRepositoryRequest) (*v1alpha1.AddPackageRepositoryResponse, error)
	GetPackageRepositoryDetail(context.Context, *v1alpha1.GetPackageRepositoryDetailRequest) (*v1alpha1.GetPackageRepositoryDetailResponse, error)
	GetPackageRepositorySummaries(context.Context, *v1alpha1.GetPackageRepositorySummariesRequest) (*v1alpha1.GetPackageRepositorySummariesResponse, error)
	UpdatePackageRepository(context.Context, *v1alpha1.UpdatePackageRepositoryRequest) (*v1alpha1.UpdatePackageRepositoryResponse, error)
	DeletePackageRepository(context.Context, *v1alpha1.DeletePackageRepositoryRequest) (*v1alpha1.DeletePackageRepositoryResponse, error)
	GetPackageRepositoryPermissions(context.Context, *v1alpha1.GetPackageRepositoryPermissionsRequest) (*v1alpha1.GetPackageRepositoryPermissionsResponse, error)
}

// UnimplementedHelmRepositoriesServiceServer should be embedded to have forward compatible implementations.
type UnimplementedHelmRepositoriesServiceServer struct {
}

func (UnimplementedHelmRepositoriesServiceServer) AddPackageRepository(context.Context, *v1alpha1.AddPackageRepositoryRequest) (*v1alpha1.AddPackageRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPackageRepository not implemented")
}
func (UnimplementedHelmRepositoriesServiceServer) GetPackageRepositoryDetail(context.Context, *v1alpha1.GetPackageRepositoryDetailRequest) (*v1alpha1.GetPackageRepositoryDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPackageRepositoryDetail not implemented")
}
func (UnimplementedHelmRepositoriesServiceServer) GetPackageRepositorySummaries(context.Context, *v1alpha1.GetPackageRepositorySummariesRequest) (*v1alpha1.GetPackageRepositorySummariesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPackageRepositorySummaries not implemented")
}
func (UnimplementedHelmRepositoriesServiceServer) UpdatePackageRepository(context.Context, *v1alpha1.UpdatePackageRepositoryRequest) (*v1alpha1.UpdatePackageRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePackageRepository not implemented")
}
func (UnimplementedHelmRepositoriesServiceServer) DeletePackageRepository(context.Context, *v1alpha1.DeletePackageRepositoryRequest) (*v1alpha1.DeletePackageRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePackageRepository not implemented")
}
func (UnimplementedHelmRepositoriesServiceServer) GetPackageRepositoryPermissions(context.Context, *v1alpha1.GetPackageRepositoryPermissionsRequest) (*v1alpha1.GetPackageRepositoryPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPackageRepositoryPermissions not implemented")
}

// UnsafeHelmRepositoriesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelmRepositoriesServiceServer will
// result in compilation errors.
type UnsafeHelmRepositoriesServiceServer interface {
	mustEmbedUnimplementedHelmRepositoriesServiceServer()
}

func RegisterHelmRepositoriesServiceServer(s grpc.ServiceRegistrar, srv HelmRepositoriesServiceServer) {
	s.RegisterService(&HelmRepositoriesService_ServiceDesc, srv)
}

func _HelmRepositoriesService_AddPackageRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.AddPackageRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmRepositoriesServiceServer).AddPackageRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmRepositoriesService/AddPackageRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmRepositoriesServiceServer).AddPackageRepository(ctx, req.(*v1alpha1.AddPackageRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelmRepositoriesService_GetPackageRepositoryDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.GetPackageRepositoryDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmRepositoriesServiceServer).GetPackageRepositoryDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmRepositoriesService/GetPackageRepositoryDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmRepositoriesServiceServer).GetPackageRepositoryDetail(ctx, req.(*v1alpha1.GetPackageRepositoryDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelmRepositoriesService_GetPackageRepositorySummaries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.GetPackageRepositorySummariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmRepositoriesServiceServer).GetPackageRepositorySummaries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmRepositoriesService/GetPackageRepositorySummaries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmRepositoriesServiceServer).GetPackageRepositorySummaries(ctx, req.(*v1alpha1.GetPackageRepositorySummariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelmRepositoriesService_UpdatePackageRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.UpdatePackageRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmRepositoriesServiceServer).UpdatePackageRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmRepositoriesService/UpdatePackageRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmRepositoriesServiceServer).UpdatePackageRepository(ctx, req.(*v1alpha1.UpdatePackageRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelmRepositoriesService_DeletePackageRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.DeletePackageRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmRepositoriesServiceServer).DeletePackageRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmRepositoriesService/DeletePackageRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmRepositoriesServiceServer).DeletePackageRepository(ctx, req.(*v1alpha1.DeletePackageRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelmRepositoriesService_GetPackageRepositoryPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.GetPackageRepositoryPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmRepositoriesServiceServer).GetPackageRepositoryPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.helm.packages.v1alpha1.HelmRepositoriesService/GetPackageRepositoryPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmRepositoriesServiceServer).GetPackageRepositoryPermissions(ctx, req.(*v1alpha1.GetPackageRepositoryPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HelmRepositoriesService_ServiceDesc is the grpc.ServiceDesc for HelmRepositoriesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelmRepositoriesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kubeappsapis.plugins.helm.packages.v1alpha1.HelmRepositoriesService",
	HandlerType: (*HelmRepositoriesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPackageRepository",
			Handler:    _HelmRepositoriesService_AddPackageRepository_Handler,
		},
		{
			MethodName: "GetPackageRepositoryDetail",
			Handler:    _HelmRepositoriesService_GetPackageRepositoryDetail_Handler,
		},
		{
			MethodName: "GetPackageRepositorySummaries",
			Handler:    _HelmRepositoriesService_GetPackageRepositorySummaries_Handler,
		},
		{
			MethodName: "UpdatePackageRepository",
			Handler:    _HelmRepositoriesService_UpdatePackageRepository_Handler,
		},
		{
			MethodName: "DeletePackageRepository",
			Handler:    _HelmRepositoriesService_DeletePackageRepository_Handler,
		},
		{
			MethodName: "GetPackageRepositoryPermissions",
			Handler:    _HelmRepositoriesService_GetPackageRepositoryPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kubeappsapis/plugins/helm/packages/v1alpha1/helm.proto",
}
