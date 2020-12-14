package client

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"monitor/proto/monitor"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"time"
)

var (
	units = []string{" bytes", "KB", "MB", "GB", "TB", "PB"}
)

func NewMonitor() *Monitor {
	return &Monitor{}
}

type Monitor struct{}

func (m *Monitor) GC(ctx context.Context, in *monitor.Empty, out *monitor.Empty) error {
	runtime.GC()
	return nil
}

func (m *Monitor) Stack(ctx context.Context, in *monitor.Empty, out *monitor.GoPsStackResponse) error {
	buff := bytes.NewBuffer([]byte{})

	if err := pprof.Lookup("goroutine").WriteTo(buff, 2); err != nil {
		return err
	}
	*out = monitor.GoPsStackResponse{
		Content: buff.Bytes(),
	}
	return nil
}

func (m *Monitor) Stats(ctx context.Context, in *monitor.Empty, out *monitor.GoPsStatsResponse) error {
	*out = monitor.GoPsStatsResponse{
		Goroutines: int32(runtime.NumGoroutine()),
		OSThreads:  int32(pprof.Lookup("threadcreate").Count()),
		GOMAXPROCS: int32(runtime.GOMAXPROCS(0)),
	}
	return nil
}

func (m *Monitor) MemStats(ctx context.Context, in *monitor.Empty, out *monitor.GoPsMemStatsResponse) error {
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	*out = monitor.GoPsMemStatsResponse{
		Alloc:            formatBytes(s.Alloc),
		TotalAlloc:       formatBytes(s.TotalAlloc),
		Sys:              formatBytes(s.Sys),
		Lookups:          int64(s.Lookups),
		Mallocs:          int64(s.Mallocs),
		Frees:            int64(s.Frees),
		HeapAlloc:        formatBytes(s.HeapAlloc),
		HeapSys:          formatBytes(s.HeapSys),
		HeapIdle:         formatBytes(s.HeapIdle),
		HeapInUse:        formatBytes(s.HeapInuse),
		HeapReleased:     formatBytes(s.HeapReleased),
		HeapObjects:      int64(s.HeapObjects),
		StackInUse:       formatBytes(s.StackInuse),
		StackSys:         formatBytes(s.StackSys),
		StackMspanInuse:  formatBytes(s.MSpanInuse),
		StackMspanSys:    formatBytes(s.MSpanSys),
		StackMcacheInuse: formatBytes(s.MCacheInuse),
		StackMcacheSys:   formatBytes(s.MCacheSys),
		OtherSys:         formatBytes(s.OtherSys),
		GcSys:            formatBytes(s.GCSys),
		NextGc:           "when heap-alloc >=" + formatBytes(s.NextGC),
		GcPauseTotal:     int64(time.Duration(s.PauseTotalNs)),
		GcPause:          int64(s.PauseNs[(s.NumGC+255)%256]),
		GcPauseEnd:       int64(s.PauseEnd[(s.NumGC+255)%256]),
		NumGc:            int64(s.NumGC),
		NumForcedGc:      int64(s.NumForcedGC),
		GcCpuFraction:    float32(s.GCCPUFraction),
		EnableGc:         s.EnableGC,
		DebugGc:          s.DebugGC,
	}
	lastGC := "-"
	if s.LastGC != 0 {
		lastGC = fmt.Sprint(time.Unix(0, int64(s.LastGC)))
	}
	out.LastGc = lastGC
	return nil
}
func (m *Monitor) CpuProfiles(ctx context.Context, in *monitor.GoPsProfilesRequest, out *monitor.GoPsProfilesResponse) error {
	buff := bytes.NewBuffer([]byte{})
	if err := pprof.StartCPUProfile(buff); err != nil {
		return err
	}
	time.Sleep(30 * time.Second)
	pprof.StopCPUProfile()
	*out = monitor.GoPsProfilesResponse{
		Content: buff.Bytes(),
	}
	return nil
}

func (m *Monitor) HeapProfiles(ctx context.Context, in *monitor.Empty, out *monitor.GoPsProfilesResponse) error {
	buff := bytes.NewBuffer([]byte{})
	if err := pprof.WriteHeapProfile(buff); err != nil {
		return err
	}
	*out = monitor.GoPsProfilesResponse{
		Content: buff.Bytes(),
	}
	return nil
}

func (m *Monitor) BinaryDump(ctx context.Context, in *monitor.Empty, out *monitor.GoPsBinaryResponse) error {
	path, err := os.Executable()
	if err != nil {
		return err
	}
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	buff := bytes.NewBuffer([]byte{})

	_, err = bufio.NewReader(f).WriteTo(buff)
	*out = monitor.GoPsBinaryResponse{
		Content: buff.Bytes(),
	}
	return err
}

func (m *Monitor) Trace(ctx context.Context, in *monitor.Empty, out *monitor.GoPsProfilesResponse) error {
	buff := bytes.NewBuffer([]byte{})
	if err := trace.Start(buff); err != nil {
		return err
	}

	time.Sleep(5 * time.Second)
	trace.Stop()
	*out = monitor.GoPsProfilesResponse{
		Content: buff.Bytes(),
	}
	return nil
}

func formatBytes(val uint64) string {
	var i int
	var target uint64
	for i = range units {
		target = 1 << uint(10*(i+1))
		if val < target {
			break
		}
	}
	if i > 0 {
		return fmt.Sprintf("%0.2f%s (%d bytes)", float64(val)/(float64(target)/1024), units[i], val)
	}
	return fmt.Sprintf("%d bytes", val)
}
