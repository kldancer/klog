package klog

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var baseLogger = logrus.New()

const (
	InfoImage  = "📝"
	DebugImage = "⚙️"
	WarnImage  = "⚠️"
	ErrorImage = "❌"
	FatalImage = "️☠️"
)

func init() {
	baseLogger.SetOutput(os.Stdout)
	formatter := &logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	}
	baseLogger.SetFormatter(formatter)
	baseLogger.SetLevel(logrus.InfoLevel)
}

func GetLogger() *logrus.Logger {
	return baseLogger
}

func SetLogLevel(logLevel string) {
	if logLevel == "debug" {
		baseLogger.SetLevel(logrus.DebugLevel)
	} else {
		baseLogger.SetLevel(logrus.InfoLevel)
	}
}

func Info(args ...interface{}) {
	baseLogger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	baseLogger.Infof("%s %s", InfoImage, message)
}

func Debug(args ...interface{}) {
	baseLogger.Debug(args...)
}
func Debugf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	baseLogger.Debugf("%s %s", DebugImage, message)
}

func Warn(args ...interface{}) {
	baseLogger.Warn(args)
}

func Warnf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	baseLogger.Warnf("%s %s", WarnImage, message)
}

func Error(args ...interface{}) {
	baseLogger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	baseLogger.Errorf("%s %s", ErrorImage, message)
}

func Fatalf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	baseLogger.Fatalf("%s %s", FatalImage, message)
	baseLogger.Exit(1)
}

func Fatal(args ...interface{}) {
	baseLogger.Fatal(args...)
	baseLogger.Exit(1)
}
