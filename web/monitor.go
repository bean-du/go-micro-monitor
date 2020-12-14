package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"monitor/infras"
	"monitor/service"
	"net/http"
	"strconv"
	"time"
)

var monitor = &monitorCtl{}

type monitorCtl struct{}

func (ctl *monitorCtl) init() {
	router.GET("/stats", ctl.Stats)
	router.GET("/stack", ctl.Stack)
	router.GET("memStats", ctl.MemStats)
	router.POST("/cpuProfiles", ctl.CpuProfiles)
	router.POST("/heapProfiles", ctl.HeapProfiles)
	router.POST("/binaryDump", ctl.BinaryDump)
}

func (ctl *monitorCtl) Stats(c *gin.Context) {
	addr := c.Query("addr")
	stats, err := service.MonitorSrv.Stats(c.Request.Context(), addr)
	if err != nil {
		c.JSON(http.StatusOK, infras.HttpResponse(1, "internal error", nil))
		return
	}
	c.JSON(http.StatusOK, infras.HttpResponse(0, "success", stats))
}

func (ctl *monitorCtl) Stack(c *gin.Context) {
	stack, err := service.MonitorSrv.Stack(c.Request.Context(), c.Query("addr"))
	if err != nil {
		c.JSON(http.StatusOK, infras.HttpResponse(1, "internal error", nil))
	}
	c.JSON(http.StatusOK, infras.HttpResponse(http.StatusOK, "success", string(stack)))
}

func (ctl *monitorCtl) MemStats(c *gin.Context) {
	stats, err := service.MonitorSrv.MemStats(c.Request.Context(), c.Query("addr"))
	if err != nil {
		c.JSON(http.StatusOK, infras.HttpResponse(1, "internal error", nil))
	}
	c.JSON(http.StatusOK, infras.HttpResponse(http.StatusOK, "success", stats))
}

func (ctl *monitorCtl) CpuProfiles(c *gin.Context) {
	serviceName := c.Query("service")

	c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
	sec, err := strconv.ParseInt(c.Query("seconds"), 10, 64)
	if sec <= 0 || err != nil {
		sec = 30
	}
	if infras.DurationExceedsWriteTimeout(c.Request, float64(sec)) {
		infras.ServeError(c.Writer, http.StatusBadRequest, "profile duration exceeds server's WriteTimeout")
		return
	}

	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	des := fmt.Sprintf("attachment; filename=\"%s\"", serviceName+"_cpu_profile")

	c.Writer.Header().Set("Content-Disposition", des)
	profiles, err := service.MonitorSrv.CpuProfiles(c.Request.Context(), c.Query("addr"), time.Duration(sec))
	if err != nil {
		c.JSON(http.StatusOK, infras.HttpResponse(1, "internal error", nil))
	}
	c.Writer.Write(profiles)
}

func (ctl *monitorCtl) HeapProfiles(c *gin.Context) {
	serviceName := c.Query("service")

	c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	des := fmt.Sprintf("attachment; filename=\"%s\"", serviceName+"_heap_profile")
	c.Writer.Header().Set("Content-Disposition", des)
	profiles, err := service.MonitorSrv.HeapProfiles(c.Request.Context(), c.Query("addr"))
	if err != nil {
		c.JSON(http.StatusOK, infras.HttpResponse(1, "internal error", nil))
	}
	c.Writer.Write(profiles)
}

func (ctl *monitorCtl) BinaryDump(c *gin.Context) {
	serviceName := c.Query("service")
	c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	des := fmt.Sprintf("attachment; filename=\"%s\"", serviceName+"_binary")
	c.Writer.Header().Set("Content-Disposition", des)
	dump, err := service.MonitorSrv.BinaryDump(c.Request.Context(), c.Query("addr"))
	if err != nil {
		c.JSON(http.StatusOK, infras.HttpResponse(1, "internal error", nil))
	}
	c.Writer.Write(dump)
}
