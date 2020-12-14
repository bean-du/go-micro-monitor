// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: coco-server/proto/monitor/monitor.proto

package monitor

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Monitor service

func NewMonitorEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Monitor service

type MonitorService interface {
	// GC 一次 gc 操作
	GC(ctx context.Context, in *Empty, opts ...client.CallOption) (*Empty, error)
	// StackStats
	Stack(ctx context.Context, in *Empty, opts ...client.CallOption) (*GoPsStackResponse, error)
	// Stats 基本状态
	Stats(ctx context.Context, in *Empty, opts ...client.CallOption) (*GoPsStatsResponse, error)
	// MemStats 查询内存状态
	MemStats(ctx context.Context, in *Empty, opts ...client.CallOption) (*GoPsMemStatsResponse, error)
	// CpuProfiles CPU监控信息
	CpuProfiles(ctx context.Context, in *GoPsProfilesRequest, opts ...client.CallOption) (*GoPsProfilesResponse, error)
	// HeapProfiles 内存监控信息
	HeapProfiles(ctx context.Context, in *Empty, opts ...client.CallOption) (*GoPsProfilesResponse, error)
	// BinaryDump 二进制文件下载
	BinaryDump(ctx context.Context, in *Empty, opts ...client.CallOption) (*GoPsBinaryResponse, error)
	// Trace 跟踪
	Trace(ctx context.Context, in *Empty, opts ...client.CallOption) (*GoPsProfilesResponse, error)
}

type monitorService struct {
	c    client.Client
	name string
}

func NewMonitorService(name string, c client.Client) MonitorService {
	return &monitorService{
		c:    c,
		name: name,
	}
}

func (c *monitorService) GC(ctx context.Context, in *Empty, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "Monitor.GC", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorService) Stack(ctx context.Context, in *Empty, opts ...client.CallOption) (*GoPsStackResponse, error) {
	req := c.c.NewRequest(c.name, "Monitor.Stack", in)
	out := new(GoPsStackResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorService) Stats(ctx context.Context, in *Empty, opts ...client.CallOption) (*GoPsStatsResponse, error) {
	req := c.c.NewRequest(c.name, "Monitor.Stats", in)
	out := new(GoPsStatsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorService) MemStats(ctx context.Context, in *Empty, opts ...client.CallOption) (*GoPsMemStatsResponse, error) {
	req := c.c.NewRequest(c.name, "Monitor.MemStats", in)
	out := new(GoPsMemStatsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorService) CpuProfiles(ctx context.Context, in *GoPsProfilesRequest, opts ...client.CallOption) (*GoPsProfilesResponse, error) {
	req := c.c.NewRequest(c.name, "Monitor.CpuProfiles", in)
	out := new(GoPsProfilesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorService) HeapProfiles(ctx context.Context, in *Empty, opts ...client.CallOption) (*GoPsProfilesResponse, error) {
	req := c.c.NewRequest(c.name, "Monitor.HeapProfiles", in)
	out := new(GoPsProfilesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorService) BinaryDump(ctx context.Context, in *Empty, opts ...client.CallOption) (*GoPsBinaryResponse, error) {
	req := c.c.NewRequest(c.name, "Monitor.BinaryDump", in)
	out := new(GoPsBinaryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorService) Trace(ctx context.Context, in *Empty, opts ...client.CallOption) (*GoPsProfilesResponse, error) {
	req := c.c.NewRequest(c.name, "Monitor.Trace", in)
	out := new(GoPsProfilesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Monitor service

type MonitorHandler interface {
	// GC 一次 gc 操作
	GC(context.Context, *Empty, *Empty) error
	// StackStats
	Stack(context.Context, *Empty, *GoPsStackResponse) error
	// Stats 基本状态
	Stats(context.Context, *Empty, *GoPsStatsResponse) error
	// MemStats 查询内存状态
	MemStats(context.Context, *Empty, *GoPsMemStatsResponse) error
	// CpuProfiles CPU监控信息
	CpuProfiles(context.Context, *GoPsProfilesRequest, *GoPsProfilesResponse) error
	// HeapProfiles 内存监控信息
	HeapProfiles(context.Context, *Empty, *GoPsProfilesResponse) error
	// BinaryDump 二进制文件下载
	BinaryDump(context.Context, *Empty, *GoPsBinaryResponse) error
	// Trace 跟踪
	Trace(context.Context, *Empty, *GoPsProfilesResponse) error
}

func RegisterMonitorHandler(s server.Server, hdlr MonitorHandler, opts ...server.HandlerOption) error {
	type monitor interface {
		GC(ctx context.Context, in *Empty, out *Empty) error
		Stack(ctx context.Context, in *Empty, out *GoPsStackResponse) error
		Stats(ctx context.Context, in *Empty, out *GoPsStatsResponse) error
		MemStats(ctx context.Context, in *Empty, out *GoPsMemStatsResponse) error
		CpuProfiles(ctx context.Context, in *GoPsProfilesRequest, out *GoPsProfilesResponse) error
		HeapProfiles(ctx context.Context, in *Empty, out *GoPsProfilesResponse) error
		BinaryDump(ctx context.Context, in *Empty, out *GoPsBinaryResponse) error
		Trace(ctx context.Context, in *Empty, out *GoPsProfilesResponse) error
	}
	type Monitor struct {
		monitor
	}
	h := &monitorHandler{hdlr}
	return s.Handle(s.NewHandler(&Monitor{h}, opts...))
}

type monitorHandler struct {
	MonitorHandler
}

func (h *monitorHandler) GC(ctx context.Context, in *Empty, out *Empty) error {
	return h.MonitorHandler.GC(ctx, in, out)
}

func (h *monitorHandler) Stack(ctx context.Context, in *Empty, out *GoPsStackResponse) error {
	return h.MonitorHandler.Stack(ctx, in, out)
}

func (h *monitorHandler) Stats(ctx context.Context, in *Empty, out *GoPsStatsResponse) error {
	return h.MonitorHandler.Stats(ctx, in, out)
}

func (h *monitorHandler) MemStats(ctx context.Context, in *Empty, out *GoPsMemStatsResponse) error {
	return h.MonitorHandler.MemStats(ctx, in, out)
}

func (h *monitorHandler) CpuProfiles(ctx context.Context, in *GoPsProfilesRequest, out *GoPsProfilesResponse) error {
	return h.MonitorHandler.CpuProfiles(ctx, in, out)
}

func (h *monitorHandler) HeapProfiles(ctx context.Context, in *Empty, out *GoPsProfilesResponse) error {
	return h.MonitorHandler.HeapProfiles(ctx, in, out)
}

func (h *monitorHandler) BinaryDump(ctx context.Context, in *Empty, out *GoPsBinaryResponse) error {
	return h.MonitorHandler.BinaryDump(ctx, in, out)
}

func (h *monitorHandler) Trace(ctx context.Context, in *Empty, out *GoPsProfilesResponse) error {
	return h.MonitorHandler.Trace(ctx, in, out)
}