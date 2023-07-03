package log

import "go.uber.org/zap"

var (
	exLogger *Logger
)

func GetLog() *Logger {
	return exLogger
}

func Debug(msg string, fields ...zap.Field) {
	exLogger.logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	exLogger.logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	exLogger.logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	exLogger.logger.Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	exLogger.logger.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	exLogger.logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	exLogger.logger.Fatal(msg, fields...)
}
