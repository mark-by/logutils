package logutils

import (
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"net"
	"testing"
)

func TestServe(t *testing.T) {
	InitLogrus("logs.txt", "debug")

	go func() {
		listener, err := net.Listen("tcp", ":8990")
		if err != nil {
			logrus.Fatal("Fail to create listener: ", err)
		}
		err = fasthttp.Serve(listener, ResetLogs)
		if err != nil {
			logrus.Fatal("Fail to serve: ", err)
		}
	}()

	listener, err := net.Listen("tcp", ":8989")
	if err != nil {
		logrus.Fatal("Fail to create listener: ", err)
	}
	err = fasthttp.Serve(listener, GetLogs)
	if err != nil {
		logrus.Fatal("Fail to serve: ", err)
	}
}
