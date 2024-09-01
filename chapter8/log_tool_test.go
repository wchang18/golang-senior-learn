package chapter8

import (
	"context"
	"github.com/gogf/gf/v2/util/grand"
	"sync"
	"testing"
)

var (
	Logger *LogTool
	config = LogConfig{
		AppName:     "app-test",
		Encoding:    "json",
		Level:       "debug",
		MaxAge:      7,
		MaxSize:     1,
		Mode:        "double",
		Path:        "./log",
		ServiceName: "test-server",
	}
	once = sync.Once{}
)

func GetLogger() *LogTool {
	once.Do(func() {
		Logger = NewLogTool(config)
	})
	return Logger
}

func TestLogTool(t *testing.T) {
	logger := GetLogger()
	logger.Logger.Info("test")
	logger.Logger.Sugar().Infow("test2", "k", "val")
}

func TestLogTool2(t *testing.T) {
	logger := GetLogger()
	logger.Debug("data-debug", "k-d", "v-d")
	logger.Info("data-info")
}

func TestLogTool3(t *testing.T) {
	logger := GetLogger()
	traceId := grand.S(32)
	ctx := NewTraceContext(context.Background(), traceId)
	logger.DebugWithContext(ctx, "start")
	logger.DebugWithContext(ctx, "step1")
	logger.DebugWithContext(ctx, "step2")
	logger.DebugWithContext(ctx, "step3")
	logger.DebugWithContext(ctx, "end")
}

func TestLogTool4(t *testing.T) {
	logger := GetLogger()
	for i := 0; i < 1000; i++ {
		traceId := grand.S(32)
		ctx := NewTraceContext(context.Background(), traceId)
		logger.DebugWithContext(ctx, "start")
		logger.DebugWithContext(ctx, "step1")
		logger.DebugWithContext(ctx, "step2")
		logger.DebugWithContext(ctx, "step3")
		logger.DebugWithContext(ctx, "end")
	}

}
