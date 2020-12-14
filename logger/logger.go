package logger

import (
	"coco-server/util/infras"
	"coco-server/util/log"
)

var (
	Log    *log.Logger
	RpcLog *log.Logger
)

func Init() {
	var err error

	Log, err = log.CreateLog("")
	infras.Throw(err)

	RpcLog, err = log.CreateLog("rpc")
	infras.Throw(err)
}
