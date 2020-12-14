package infras

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"net"
	"net/http"
)

func HttpResponse(code int, msg string, data interface{}) gin.H {
	return gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}

func VerifyIpAddr(addr string) error {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return err
	}
	if ip := net.ParseIP(host); ip == nil {
		return errors.New("addr invalid")
	}
	return nil
}

func DirectCallOption(addr string) client.CallOption {
	return func(co *client.CallOptions) {
		co.Address = []string{addr}
		co.CallWrappers = nil //此处直接根据addr进行服务调用，不走全局的client中间件
	}
}

func DurationExceedsWriteTimeout(r *http.Request, seconds float64) bool {
	srv, ok := r.Context().Value(http.ServerContextKey).(*http.Server)
	return ok && srv.WriteTimeout != 0 && seconds >= srv.WriteTimeout.Seconds()
}

func ServeError(w http.ResponseWriter, status int, txt string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Go-Pprof", "1")
	w.Header().Del("Content-Disposition")
	w.WriteHeader(status)
	fmt.Fprintln(w, txt)
}
