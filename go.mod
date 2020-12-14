module monitor

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace coco-server/util => lemco.dev/coco-server/util v0.0.0-20201203040948-69e41a48425c

require (
	coco-server/util v0.0.0-00010101000000-000000000000
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v7 v7.4.0
	github.com/golang/protobuf v1.4.0
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.9.1
	github.com/micro/go-plugins/store/redis/v2 v2.9.1
)
