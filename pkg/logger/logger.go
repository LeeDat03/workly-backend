package logger

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
)

var std *log.Logger

func Init(level, prefix string) {
	std = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.DateTime,
		Prefix:          prefix,
	})
	std.SetLevel(parseLevel(level))
}

func parseLevel(level string) log.Level {
	switch level {
	case "debug":
		return log.DebugLevel
	case "info":
		return log.InfoLevel
	case "warn":
		return log.WarnLevel
	case "error":
		return log.ErrorLevel
	case "fatal":
		return log.FatalLevel
	default:
		return log.InfoLevel
	}
}

func init() {
	Init("info", "")
}

// Simple wrapper functions for common use
func Debug(msg interface{}, keyvals ...interface{}) { std.Debug(msg, keyvals...) }
func Info(msg interface{}, keyvals ...interface{})  { std.Info(msg, keyvals...) }
func Warn(msg interface{}, keyvals ...interface{})  { std.Warn(msg, keyvals...) }
func Error(msg interface{}, keyvals ...interface{}) { std.Error(msg, keyvals...) }
func Fatal(msg interface{}, keyvals ...interface{}) { std.Fatal(msg, keyvals...) }
func Debugf(format string, args ...interface{})     { std.Debugf(format, args...) }
func Infof(format string, args ...interface{})      { std.Infof(format, args...) }
func Warnf(format string, args ...interface{})      { std.Warnf(format, args...) }
func Errorf(format string, args ...interface{})     { std.Errorf(format, args...) }
func Fatalf(format string, args ...interface{})     { std.Fatalf(format, args...) }
func With(keyvals ...interface{}) *log.Logger       { return std.With(keyvals...) }
