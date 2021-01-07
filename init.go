package logutils

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var logFile *os.File

type Options struct {
	File      string
	SetFormat func()
	SetOutput func(io.Writer)
	SetLevel  func()
}

func Init(options *Options) {
	options.SetLevel()
	setOutput(options)
}

func setOutput(options *Options) {
	if options.File == "-" || options.File == "" {
		return
	}

	options.SetFormat()
	file, err := os.Create(options.File)
	logFile = file
	if err != nil {
		logrus.Fatal("Fail to create file for logs: ", err)
	}
	options.SetOutput(file)
}
