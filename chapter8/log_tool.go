package chapter8

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type LogTool struct {
	conf   LogConfig
	Logger *zap.Logger
}

type LogConfig struct {
	ServiceName string `json:"serviceName"` // 服务名称
	Mode        string `json:"mode"`        // 日志模式，console、file、double
	Encoding    string `json:"encoding"`    // 日志格式,json、console
	Level       string `json:"level"`       // 日志级别
	Path        string `json:"path"`        // 日志文件路径
	MaxSize     int    `json:"maxSize"`     // 单位M
	MaxAge      int    `json:"maxAge"`      // 单位天
	AppName     string `json:"appName"`     //APP Name
}

func NewLogTool(c LogConfig) *LogTool {
	tool := &LogTool{
		conf: c,
	}
	tool.Init()
	return tool
}

func (z *LogTool) Init() {
	core := zapcore.NewCore(z.getEncoder(), z.getLogWriter(), z.getLogLevel())
	z.Logger = zap.New(core, zap.AddCaller(), zap.Fields(zap.String("appName", z.conf.AppName)))
}

func (z *LogTool) getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if z.conf.Encoding == "json" {
		return zapcore.NewJSONEncoder(encoderConfig)
	} else {
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
}

func (z *LogTool) getLogWriter() zapcore.WriteSyncer {
	fileName := fmt.Sprintf("%s/%s.log", z.conf.Path, z.conf.ServiceName)
	lumberWriteSyncer := &lumberjack.Logger{
		Filename: fileName,
		MaxSize:  z.conf.MaxSize,
		MaxAge:   z.conf.MaxAge,
		Compress: false,
	}
	switch z.conf.Mode {
	case "console":
		return zapcore.AddSync(os.Stdout)
	case "file":
		return zapcore.AddSync(lumberWriteSyncer)
	case "double":
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberWriteSyncer))
	}
	return zapcore.AddSync(os.Stdout)
}

func (z *LogTool) getLogLevel() zap.AtomicLevel {
	atomicLevel := zap.NewAtomicLevel()
	switch z.conf.Level {
	case "debug":
		atomicLevel.SetLevel(zapcore.DebugLevel)
	case "info":
		atomicLevel.SetLevel(zapcore.InfoLevel)
	case "error":
		atomicLevel.SetLevel(zapcore.ErrorLevel)
	case "fatal":
		atomicLevel.SetLevel(zapcore.PanicLevel)
	default:
		atomicLevel.SetLevel(zapcore.InfoLevel)
	}
	return atomicLevel
}

func (z *LogTool) Debug(msg string, kv ...any) {
	z.Logger.Sugar().Debugw(msg, z.getField(kv...)...)
}

func (z *LogTool) Info(msg string, kv ...any) {
	z.Logger.Sugar().Infow(msg, z.getField(kv...)...)
}

func (z *LogTool) Error(msg string, kv ...any) {
	z.Logger.Sugar().Errorw(msg, z.getField(kv...)...)
}

func (z *LogTool) Fatal(msg string, kv ...any) {
	z.Logger.Sugar().Fatalw(msg, z.getField(kv...)...)
}

func (z *LogTool) getField(kv ...any) []any {
	return kv
}

func (z *LogTool) DebugWithContext(ctx context.Context, msg string, kv ...any) {
	z.Logger.Sugar().Debugw(msg, z.getFieldWithContext(ctx, kv...)...)
}

func (z *LogTool) InfoWithContext(ctx context.Context, msg string, kv ...any) {
	z.Logger.Sugar().Infow(msg, z.getFieldWithContext(ctx, kv...)...)
}

func (z *LogTool) ErrorWithContext(ctx context.Context, msg string, kv ...any) {
	z.Logger.Sugar().Errorw(msg, z.getFieldWithContext(ctx, kv...)...)
}

func (z *LogTool) FatalWithContext(ctx context.Context, msg string, kv ...any) {
	z.Logger.Sugar().Fatalw(msg, z.getFieldWithContext(ctx, kv...)...)
}

func (z *LogTool) getFieldWithContext(ctx context.Context, kv ...any) []any {
	list := append([]any{}, TraceId, GetTraceIdFromContext(ctx))
	return append(list, kv...)
}

const TraceId = "traceId"

func NewTraceContext(ctx context.Context, traceId string) context.Context {
	return context.WithValue(ctx, TraceId, traceId)
}

func GetTraceIdFromContext(ctx context.Context) string {
	traceId := ctx.Value(TraceId)
	if traceId == nil {
		return ""
	}
	return traceId.(string)
}
