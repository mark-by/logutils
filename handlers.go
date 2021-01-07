package logutils

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"io"
	"io/ioutil"
	"os"
)

// /logs
func GetLogs(ctx *fasthttp.RequestCtx) {
	logrus.Info("Hiiii")
	if logFile == nil {
		ctx.Response.SetStatusCode(fasthttp.StatusNoContent)
		return
	}

	logs, err := ioutil.ReadFile(logFile.Name())
	ctx.SetContentType("text/plain")
	if err != nil {
		serveError(ctx, err)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(logs)
}

// logs/reset
func ResetLogs(ctx *fasthttp.RequestCtx) {
	tmpFile, err := os.Create(logFile.Name() + ".tmp")

	if err != nil {
		serveError(ctx, err)
		return
	}

	currFile, err := os.Open(logFile.Name())
	if err != nil {
		serveError(ctx, err)
		return
	}

	buffer := bytes.Buffer{}

	_, err = io.Copy(&buffer, currFile)
	if err != nil {
		serveError(ctx, err)
		return
	}

	ctx.Response.SetBody(buffer.Bytes())

	go func() {
		_, _ = io.Copy(tmpFile, &buffer)
		_ = currFile.Close()
		_ = tmpFile.Close()

		_ = logFile.Truncate(0)
		_, _ = logFile.Seek(0, 0)
	}()
}

func serveError(ctx *fasthttp.RequestCtx, err error) {
	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	ctx.SetBody([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
}
