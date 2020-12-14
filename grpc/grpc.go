package grpc

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	monitor "monitor/client"
	"monitor/config"
	"time"
)

var grpcErr error

func Run(ch chan<- error) {
	defer func() {
		ch <- grpcErr
	}()

	config.Micro.Init(
		micro.Name(config.Conf.ServiceName),
		micro.Version(config.Conf.Version),

		micro.RegisterInterval(5*time.Second),
		micro.RegisterTTL(10*time.Second),
	)

	grpcErr = config.Micro.Options().Client.Init(
		client.Retries(3),
		client.RequestTimeout(1*time.Minute),
		client.PoolTTL(30*time.Minute),
	)

	grpcErr = micro.RegisterHandler(config.Micro.Server(), monitor.NewMonitor())
	grpcErr = config.Micro.Run()
}
