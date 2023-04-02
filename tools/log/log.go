package log

import (
	"github.com/CreFire/rain/tools/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

type Level = zapcore.Level

type Logger struct {
	l *zap.Logger
	// https://pkg.go.dev/go.uber.org/zap#example-AtomicLevel
	al *zap.AtomicLevel
}

func New(out io.Writer, level Level, opts ...Option) *Logger {
	if out == nil {
		out = os.Stderr
	}
	al := zap.NewAtomicLevelAt(level)
	var cfg zapcore.Encoder
	if config.IsDev() {
		cfg = getDevEncoder()
	} else {
		cfg = getProdEncoder()
	}
	core := zapcore.NewCore(cfg, zapcore.AddSync(out), al)
	return &Logger{l: zap.New(core, opts...), al: &al}
}

func NewDefault() *Logger {
	al := zap.NewAtomicLevelAt(getLogLevel(config.Conf.Log.Level))
	var core zapcore.Core
	if config.IsDev() {
		core = zapcore.NewCore(getDevEncoder(), zapcore.AddSync(os.Stdout), al)
	} else {
		core = zapcore.NewCore(getProdEncoder(), getWriter(config.Conf.Log), al)
	}
	// 传入 zap.AddCaller() 显示打日志点的文件名和行数
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.DPanicLevel))
	return &Logger{l: logger, al: &al}
}

// SetLevel 动态更改日志级别
// 对于使用 NewTee 创建的 Logger 无效，因为 NewTee 本意是根据不同日志级别
// 创建的多个 zap.Core，不应该通过 SetLevel 将多个 zap.Core 日志级别统一
func (l *Logger) SetLevel(level Level) {
	if l.al != nil {
		l.al.SetLevel(level)
	}
}

type Field = zap.Field

func (l *Logger) Debug(msg string, fields ...Field) {
	l.l.Debug(msg, fields...)
}
func (l *Logger) Debugf(msg string, fields ...interface{}) {
	l.l.Sugar().Debug(msg, fields)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.l.Info(msg, fields...)
}

func (l *Logger) Infof(msg string, fields ...Field) {
	l.l.Sugar().Info(msg, fields)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.l.Warn(msg, fields...)
}

func (l *Logger) Warnf(msg string, fields ...Field) {
	l.l.Sugar().Warn(msg, fields)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.l.Error(msg, fields...)
}

func (l *Logger) Errorf(msg string, fields ...Field) {
	l.l.Sugar().Error(msg, fields)
}

func (l *Logger) Panic(msg string, fields ...Field) {
	l.l.Panic(msg, fields...)
}

func (l *Logger) Panicf(msg string, fields ...Field) {
	l.l.Sugar().Panic(msg, fields)
}

func (l *Logger) Fatal(msg string, fields ...Field) {
	l.l.Fatal(msg, fields...)
}

func (l *Logger) Fatalf(msg string, fields ...Field) {
	l.l.Sugar().Panic(msg, fields)
}

func (l *Logger) Sync() error {
	return l.l.Sync()
}

var std = NewDefault()

func Default() *Logger         {
	return std
}
func ReplaceDefault(l *Logger) { std = l }

func SetLevel(level Level) { std.SetLevel(level) }

func Debug(msg string, fields ...Field) { std.Debug(msg, fields...) }
func Info(msg string, fields ...Field)  { std.Info(msg, fields...) }
func Warn(msg string, fields ...Field)  { std.Warn(msg, fields...) }
func Error(msg string, fields ...Field) { std.Error(msg, fields...) }
func Panic(msg string, fields ...Field) { std.Panic(msg, fields...) }
func Fatal(msg string, fields ...Field) { std.Fatal(msg, fields...) }

func Sync() error { return std.Sync() }
