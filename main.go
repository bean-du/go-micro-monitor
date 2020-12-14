package main

import (
	"coco-server/util/infras"
	_ "github.com/micro/go-plugins/registry/etcdv3/v2"
	"monitor/config"
	"monitor/grpc"
	"monitor/logger"
	"monitor/web"
	"os"
)

func main() {
	config.Init()
	logger.Init()

	logger.Log.Infow("system start.", "service-name", config.Conf.ServiceName)
	if errs := infras.Await(web.Run, grpc.Run); len(errs) > 0 {
		for _, err := range errs {
			logger.Log.Errw("system error", "err", err)
		}
		os.Exit(-1)
	}
}
