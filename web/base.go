package web

import (
	"github.com/gin-gonic/gin"
	"monitor/infras"
	"monitor/service"
	"net/http"
)

var base = &baseCtl{}

type baseCtl struct{}

func (b *baseCtl) init() {
	router.GET("/", b.index)
	router.GET("/service", b.serverList)
	router.GET("/service/nodes", b.serverNode)
	router.GET("/service/providers", b.providers)
}

func (b *baseCtl) index(c *gin.Context) {
	index := service.BaseSrv.Index()
	c.Writer.WriteString(index)
}

func (b *baseCtl) serverList(c *gin.Context) {
	list := service.BaseSrv.ServerList()
	c.JSON(http.StatusOK, infras.HttpResponse(0, "success", list))
}

func (b *baseCtl) serverNode(c *gin.Context) {
	query := c.Query("name")
	list := service.BaseSrv.NodeList(query)
	c.JSON(http.StatusOK, infras.HttpResponse(0, "success", list))
}

func (b *baseCtl) providers(c *gin.Context) {
	providers := service.MonitorSrv.ListProviders()
	c.JSON(http.StatusOK, infras.HttpResponse(http.StatusOK, "success", providers))
}
