package log

import (
	"github.com/CreFire/rain/tools/config"
	"github.com/sirupsen/logrus"
	"io"
	"testing"

	"go.uber.org/zap"
)

func BenchmarkLog(b *testing.B) {
	logger, _ := New(config.Conf.Log)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Debug("A walrus appears",
			zap.String("animal", "walrus"),
			zap.Int("number", 1),
			zap.Float64("size", 10.1),
		)
	}
}

func BenchmarkZap(b *testing.B) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logger.Debug("A walrus appears",
			zap.String("animal", "walrus"),
			zap.Int("number", 1),
			zap.Float64("size", 10.1),
		)
	}
}

func BenchmarkLogrus(b *testing.B) {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(io.Discard)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logger.WithFields(logrus.Fields{
			"animal": "walrus",
			"number": 1,
			"size":   10.1,
		}).Debug("A walrus appears")
	}
}
