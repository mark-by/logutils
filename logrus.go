package logutils

import "github.com/sirupsen/logrus"

func InitLogrus(filename, level string) {
	Init(&Options{
		File: filename,
		SetFormat: func() {
			logrus.SetFormatter(&logrus.JSONFormatter{
				TimestampFormat: "2006-01-02 15:04:05",
			})
		},
		SetOutput: logrus.SetOutput,
		SetLevel: func() {
			logrus.SetLevel(chooseLevel(level))
		},
	})
}

func chooseLevel(logLevel string) logrus.Level {
	switch logLevel {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}
