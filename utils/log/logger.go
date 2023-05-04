package log

import (
	"errors"
	"github.com/CreFire/rain/utils/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

type Logger struct {
	logger *zap.Logger
}

func init() {
	exLogger, _ = New(config.Conf.Log)
}
func New(cfg *config.Log) (*Logger, error) {
	var (
		level    zapcore.Level
		writer   zapcore.WriteSyncer
		encoder  zapcore.Encoder
		core     zapcore.Core
		fileCore zapcore.Core
	)

	// 解析日志级别
	if err := level.UnmarshalText([]byte(cfg.Level)); err != nil {
		return nil, err
	}

	// 设置日志编码器
	switch cfg.Encoding {
	case "json":
		encoder = zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	case "console":
		encoder = zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	case "dev":
		encoder = getDevEncoder()
	case "prod":
		encoder = getProdEncoder()
	default:
		return nil, errors.New("invalid encoding")
	}
	// 设置日志输出
	if cfg.Stdout {
		writer = zapcore.Lock(os.Stdout)
	}
	writer = zapcore.AddSync(io.Discard)
	if cfg.Filename != "" {
		fileWriter := zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.Filename,
			MaxSize:    cfg.Maxsize,
			MaxAge:     cfg.MaxAge,
			MaxBackups: cfg.FileMaxBackups,
			LocalTime:  true,
			Compress:   cfg.Compress,
		})
		fileCore = zapcore.NewCore(
			encoder,
			fileWriter,
			level,
		)
		if writer != nil {
			writer = zapcore.NewMultiWriteSyncer(writer, fileWriter)
		} else {
			writer = fileWriter
		}
	}
	zap.NewProduction(zap.WithCaller(true), zap.AddStacktrace(zapcore.DPanicLevel), zap.IncreaseLevel(level))
	// 组合日志核心
	core = zapcore.NewCore(encoder, writer, level)

	// 添加 Caller 和 StackTrace
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.DPanicLevel))

	if fileCore != nil {
		fileLogger := zap.New(fileCore, zap.AddCaller(), zap.AddStacktrace(zapcore.DPanicLevel))
		defer fileLogger.Sync()
		fileLogger.Info("Logger initialization successful")
	}
	return &Logger{logger: logger}, nil
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *Logger) DPanic(msg string, fields ...zap.Field) {
	l.logger.DPanic(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...zap.Field) {
	l.logger.Panic(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.logger.Fatal(msg, fields...)
}
