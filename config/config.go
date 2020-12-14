package config

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	configs "github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/cmd"
	"github.com/micro/go-micro/v2/config/source/etcd"
	"os"
	"time"
)

var (
	EtcdAddr = ""
	Conf     = &Config{}

	StartTime = time.Now()
	ExitC     = make(chan bool)
	Micro     = micro.NewService(micro.AfterStop(afterExit))
)

func Init() {
	err := extractEtcdAddr()
	if err != nil {
		panic(nil)
	}
	etcdSrc := etcd.NewSource(
		etcd.WithAddress(EtcdAddr),
		etcd.WithPrefix("/config/monitor"),
		etcd.StripPrefix(true),
	)
	conf, err := configs.NewConfig()
	if err != nil {
		panic(err)
	}
	if err := conf.Load(etcdSrc); err != nil {
		panic(err)
	}
	if err = conf.Scan(Conf); err != nil {
		panic(err)
	}

	Micro.Init()
	initRedis()
}

func afterExit() error {
	close(ExitC)
	return nil
}

type Config struct {
	ServiceName string `json:"service_name"`
	Version     string `json:"version"`
	Env         string `json:"env"`

	Port int `json:"port"`

	Redis struct {
		Addr string `json:"addr"`
		Pwd  string `json:"pwd"`
	} `json:"redis"`
}

func extractEtcdAddr() error {
	app := cli.NewApp()
	app.Name = "coco-mgr"
	app.Flags = cmd.DefaultCmd.App().Flags
	app.Action = func(ctx *cli.Context) error {
		EtcdAddr = ctx.String("registry_address")
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		return err
	}
	return nil
}
