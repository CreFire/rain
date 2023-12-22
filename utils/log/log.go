package log

import (
	"context"
	"github.com/CreFire/rain/utils/config"
	"github.com/pkg/errors"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log/slog"
	"os"
)

var exLogger *slog.Logger

func init() {
	var err error
	exLogger, err = NewLogger(config.Conf.Log)
	if err != nil {
		exLogger.Error("", "err", err.Error())
	}
}

func NewLogger(cfg *config.Log) (*slog.Logger, error) {

	// 日志输出目标
	var Writers []io.Writer

	// 如果指定了文件名，则输出到文件
	if cfg.Filename != "" {
		logWriter := &lumberjack.Logger{
			Filename:   cfg.Filename,
			MaxSize:    cfg.Maxsize,        // 文件最大大小（MB）
			MaxAge:     cfg.MaxAge,         // 文件最大存活时间（天）
			MaxBackups: cfg.FileMaxBackups, // 最大备份文件个数
			LocalTime:  true,
			Compress:   cfg.Compress, // 是否压缩/归档旧文件
		}
		Writers = append(Writers, logWriter)
	} else {
		// 否则输出到标准输出
		Writers = append(Writers, os.Stdout)
	}

	// 设置日志编码器
	var encoder slog.Handler
	switch cfg.Encoding {
	case "json":
		encoder = slog.NewJSONHandler(io.MultiWriter(Writers...), nil)
	case "text":
		encoder = slog.NewTextHandler(io.MultiWriter(Writers...), nil)
	default:
		return nil, errors.New("invalid encoding")
	}

	// 创建日志记录器
	logger := slog.New(encoder)
	encoder.Enabled(context.Background(), slog.LevelInfo)

	return logger, nil
}
