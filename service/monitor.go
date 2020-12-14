package service

import (
	"context"
	"errors"
	"monitor/provider"
	"time"
)

var MonitorSrv = &monitorSrv{p: provider.DefaultProvider()}

type monitorSrv struct {
	p provider.Monitor
}

func (m *monitorSrv) ListProviders() []string {
	pro := []string{}
	for k := range provider.Provider {
		pro = append(pro, k)
	}
	return pro
}

func (m *monitorSrv) ChangeProvider(name string) error {
	if pro, ok := provider.Provider[name]; ok {
		m.p = pro
		return nil
	} else {
		return errors.New("provider not found")
	}
}

func (m *monitorSrv) Stats(ctx context.Context, addr string) (*provider.Stats, error) {
	return m.p.Stats(ctx, addr)
}

func (m *monitorSrv) Stack(ctx context.Context, addr string) ([]byte, error) {
	return m.p.Stack(ctx, addr)
}

func (m *monitorSrv) MemStats(ctx context.Context, addr string) (*provider.MemStats, error) {
	return m.p.MemStats(ctx, addr)
}

func (m *monitorSrv) CpuProfiles(ctx context.Context, addr string, duration time.Duration) ([]byte, error) {
	return m.p.CpuProfiles(ctx, addr, duration)
}

func (m *monitorSrv) HeapProfiles(ctx context.Context, addr string) ([]byte, error) {
	return m.p.HeapProfiles(ctx, addr)
}

func (m *monitorSrv) BinaryDump(ctx context.Context, addr string) ([]byte, error) {
	return m.p.BinaryDump(ctx, addr)
}
