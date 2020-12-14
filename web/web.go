package web

import (
	"coco-server/util/infras"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	maddr "github.com/micro/go-micro/v2/util/addr"
	"github.com/micro/go-micro/v2/web"
	"monitor/config"
	"monitor/web/middle"
	"net"
	"strings"
	"time"
)

var (
	webErr  error
	webName string
	webAddr string
	webSrv  web.Service
	router   = gin.Default()
)

func Run(ch chan<- error)  {
	defer func() {
		ch <- webErr
	}()

	webName = config.Conf.ServiceName + "-web"
	webSrv = web.NewService(
		web.MicroService(config.Micro),
		web.Name(webName),
		web.Version(config.Conf.Version),

		web.RegisterInterval(5*time.Second),
		web.RegisterTTL(10*time.Second),
		web.Address(":9099"),

		web.AfterStart(extractWebAddr),
	)

	webErr = webSrv.Init()
	infras.Throw(webErr)

	router.Use(middle.Log)
	router.Use(cors.Default())
	webSrv.Handle("/", router)

	base.init()
	monitor.init()
	webErr = webSrv.Run()
}


func extractWebAddr() error {
	host, port, err := net.SplitHostPort(webSrv.Options().Address)
	if err != nil {
		return err
	}
	addr, err := maddr.Extract(host)
	if err != nil {
		return err
	}
	if strings.Count(addr, ":") > 0 {
		addr = "[" + addr + "]"
	}

	webAddr = fmt.Sprintf("%v:%v", addr, port)
	return nil
}
