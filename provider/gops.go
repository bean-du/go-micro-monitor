package provider

import (
	"context"
	"encoding/json"
	"errors"
	"monitor/config"
	"monitor/infras"
	"monitor/proto/monitor"
	"time"
)

func init() {
	Register("gops", &GoPs{})
}

func DefaultProvider() Monitor {
	return &GoPs{}
}

type GoPs struct {
}

var _ Monitor = (*GoPs)(nil)

func (g *GoPs) Protocol() ProtocolType {
	return ProtocolHttpGrpc
}

func (g *GoPs) Stats(ctx context.Context, addr string) (*Stats, error) {
	if len(addr) == 0 {
		return nil, errors.New("request Addr is empty")
	}
	if e := infras.VerifyIpAddr(addr); e != nil {
		return nil, e
	}

	in := &monitor.Empty{}
	request := config.MonitorClient.NewRequest("", "Monitor.Stats", in)
	res := &monitor.GoPsStatsResponse{}
	if err := config.MonitorClient.Call(ctx, request, res, infras.DirectCallOption(addr)); err != nil {
		return nil, err
	}
	status := &Stats{
		GOMAXPROCS: res.GOMAXPROCS,
		Goroutines: res.Goroutines,
		OSThreads:  res.OSThreads,
		NumCPU:     res.NumCPU,
	}
	return status, nil
}

func (g *GoPs) Stack(ctx context.Context, addr string) ([]byte, error) {
	if len(addr) == 0 {
		return nil, errors.New("request Addr is empty")
	}
	if e := infras.VerifyIpAddr(addr); e != nil {
		return nil, e
	}
	in := &monitor.Empty{}
	request := config.MonitorClient.NewRequest("", "Monitor.Stack", in)
	res := &monitor.GoPsStackResponse{}

	if err := config.MonitorClient.Call(ctx, request, res, infras.DirectCallOption(addr)); err != nil {
		return nil, err
	}
	return res.Content, nil
}

func (g *GoPs) MemStats(ctx context.Context, addr string) (*MemStats, error) {
	if len(addr) == 0 {
		return nil, errors.New("request Addr is empty")
	}
	if e := infras.VerifyIpAddr(addr); e != nil {
		return nil, e
	}
	in := &monitor.Empty{}
	request := config.MonitorClient.NewRequest("", "Monitor.MemStats", in)
	res := &monitor.GoPsMemStatsResponse{}

	if err := config.MonitorClient.Call(ctx, request, res, infras.DirectCallOption(addr)); err != nil {
		return nil, err
	}
	marshal, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	stats := &MemStats{}
	if err := json.Unmarshal(marshal, stats); err != nil {
		return nil, err
	}
	return stats, nil
}

func (g *GoPs) CpuProfiles(ctx context.Context, addr string, duration time.Duration) ([]byte, error) {
	if len(addr) == 0 {
		return nil, errors.New("request Addr is empty")
	}
	if e := infras.VerifyIpAddr(addr); e != nil {
		return nil, e
	}
	in := &monitor.GoPsProfilesRequest{
		Duration: int64(duration),
	}
	request := config.MonitorClient.NewRequest("", "Monitor.CpuProfiles", in)
	res := &monitor.GoPsProfilesResponse{}

	if err := config.MonitorClient.Call(ctx, request, res, infras.DirectCallOption(addr)); err != nil {
		return nil, err
	}
	return res.Content, nil
}

func (g *GoPs) HeapProfiles(ctx context.Context, addr string) ([]byte, error) {
	if len(addr) == 0 {
		return nil, errors.New("request Addr is empty")
	}
	if e := infras.VerifyIpAddr(addr); e != nil {
		return nil, e
	}
	in := &monitor.Empty{}

	request := config.MonitorClient.NewRequest("", "Monitor.HeapProfiles", in)
	res := &monitor.GoPsProfilesResponse{}

	if err := config.MonitorClient.Call(ctx, request, res, infras.DirectCallOption(addr)); err != nil {
		return nil, err
	}
	return res.Content, nil
}

func (g *GoPs) BinaryDump(ctx context.Context, addr string) ([]byte, error) {
	if len(addr) == 0 {
		return nil, errors.New("request Addr is empty")
	}
	if e := infras.VerifyIpAddr(addr); e != nil {
		return nil, e
	}
	in := &monitor.Empty{}

	request := config.MonitorClient.NewRequest("", "Monitor.BinaryDump", in)
	res := &monitor.GoPsBinaryResponse{}

	if err := config.MonitorClient.Call(ctx, request, res, infras.DirectCallOption(addr)); err != nil {
		return nil, err
	}
	return res.Content, nil
}
