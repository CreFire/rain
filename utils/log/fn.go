package log

import "os"

func Debug(msg string, fields ...any) {
	exLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...any) {
	exLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...any) {
	exLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...any) {
	exLogger.Error(msg, fields...)
}

func Err(err error) {
	exLogger.Error("error", "err", err.Error())
}

func Fatal(msg string, fields ...any) {
	exLogger.Error(msg, fields...)
	os.Exit(1)
}
