package provider

import (
	"context"
	"time"
)

var Provider = make(map[string]Monitor)

func Register(name string, m Monitor) {
	Provider[name] = m
}

type ProtocolType string

const (
	ProtocolHttp     ProtocolType = "http"
	ProtocolHttpGrpc ProtocolType = "grpc"
)

// 监控服务数据提供者
type Monitor interface {
	// 提供者提供数据的协议类型
	Protocol() ProtocolType
	// 基本状态接口
	Stats(ctx context.Context, addr string) (*Stats, error)
	// Stack 栈跟踪
	Stack(ctx context.Context, addr string) ([]byte, error)
	// MemStats 内存详细状态
	MemStats(ctx context.Context, addr string) (*MemStats, error)
	// CpuProfiles Cpu 数据
	CpuProfiles(ctx context.Context, addr string, duration time.Duration) ([]byte, error)
	// HeapProfiles 堆内存数据
	HeapProfiles(ctx context.Context, addr string) ([]byte, error)
	// BinaryDump 二进制文件数据
	BinaryDump(ctx context.Context, addr string) ([]byte, error)
}

type Stats struct {
	Goroutines int32 `json:"goroutines,omitempty"`
	OSThreads  int32 `json:"OSThreads,omitempty"`
	GOMAXPROCS int32 `json:"GOMAXPROCS,omitempty"`
	NumCPU     int32 `json:"NumCPU,omitempty"`
}

// GoPsMemStatsResponse 内存状态返回数据
type MemStats struct {
	Alloc            string  `protobuf:"bytes,1,opt,name=Alloc,proto3" json:"Alloc,omitempty"`
	TotalAlloc       string  `protobuf:"bytes,2,opt,name=TotalAlloc,proto3" json:"TotalAlloc,omitempty"`
	Sys              string  `protobuf:"bytes,3,opt,name=Sys,proto3" json:"Sys,omitempty"`
	Lookups          int64   `protobuf:"varint,4,opt,name=Lookups,proto3" json:"Lookups,omitempty"`
	Mallocs          int64   `protobuf:"varint,5,opt,name=Mallocs,proto3" json:"Mallocs,omitempty"`
	Frees            int64   `protobuf:"varint,6,opt,name=Frees,proto3" json:"Frees,omitempty"`
	HeapAlloc        string  `protobuf:"bytes,7,opt,name=HeapAlloc,proto3" json:"HeapAlloc,omitempty"`
	HeapSys          string  `protobuf:"bytes,8,opt,name=HeapSys,proto3" json:"HeapSys,omitempty"`
	HeapIdle         string  `protobuf:"bytes,9,opt,name=HeapIdle,proto3" json:"HeapIdle,omitempty"`
	HeapInUse        string  `protobuf:"bytes,10,opt,name=HeapInUse,proto3" json:"HeapInUse,omitempty"`
	HeapReleased     string  `protobuf:"bytes,11,opt,name=HeapReleased,proto3" json:"HeapReleased,omitempty"`
	HeapObjects      int64   `protobuf:"varint,12,opt,name=HeapObjects,proto3" json:"HeapObjects,omitempty"`
	StackInUse       string  `protobuf:"bytes,13,opt,name=StackInUse,proto3" json:"StackInUse,omitempty"`
	StackSys         string  `protobuf:"bytes,14,opt,name=StackSys,proto3" json:"StackSys,omitempty"`
	StackMspanInuse  string  `protobuf:"bytes,15,opt,name=StackMspanInuse,proto3" json:"StackMspanInuse,omitempty"`
	StackMspanSys    string  `protobuf:"bytes,16,opt,name=StackMspanSys,proto3" json:"StackMspanSys,omitempty"`
	StackMcacheInuse string  `protobuf:"bytes,17,opt,name=StackMcacheInuse,proto3" json:"StackMcacheInuse,omitempty"`
	StackMcacheSys   string  `protobuf:"bytes,18,opt,name=StackMcacheSys,proto3" json:"StackMcacheSys,omitempty"`
	OtherSys         string  `protobuf:"bytes,19,opt,name=OtherSys,proto3" json:"OtherSys,omitempty"`
	GcSys            string  `protobuf:"bytes,20,opt,name=GcSys,proto3" json:"GcSys,omitempty"`
	NextGc           string  `protobuf:"bytes,21,opt,name=NextGc,proto3" json:"NextGc,omitempty"`
	LastGc           string  `protobuf:"bytes,22,opt,name=LastGc,proto3" json:"LastGc,omitempty"`
	GcPauseTotal     int64   `protobuf:"varint,23,opt,name=gcPauseTotal,proto3" json:"gcPauseTotal,omitempty"`
	GcPause          int64   `protobuf:"varint,24,opt,name=GcPause,proto3" json:"GcPause,omitempty"`
	GcPauseEnd       int64   `protobuf:"varint,25,opt,name=GcPauseEnd,proto3" json:"GcPauseEnd,omitempty"`
	NumGc            int64   `protobuf:"varint,26,opt,name=NumGc,proto3" json:"NumGc,omitempty"`
	NumForcedGc      int64   `protobuf:"varint,27,opt,name=NumForcedGc,proto3" json:"NumForcedGc,omitempty"`
	GcCpuFraction    float32 `protobuf:"fixed32,28,opt,name=GcCpuFraction,proto3" json:"GcCpuFraction,omitempty"`
	EnableGc         bool    `protobuf:"varint,29,opt,name=EnableGc,proto3" json:"EnableGc,omitempty"`
	DebugGc          bool    `protobuf:"varint,30,opt,name=DebugGc,proto3" json:"DebugGc,omitempty"`
}
