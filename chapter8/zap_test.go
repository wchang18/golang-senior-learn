package chapter8

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"testing"
)

func TestExample(t *testing.T) {
	logger := zap.NewExample()
	logger.Info("user info", zap.String("name", "tom"), zap.Int("age", 18))
	logger.Sugar().Infow("user data",
		"url", "www.qq.com",
		"telephone", 12312312311,
		"delete", false,
	)
}

func TestDevelop(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	logger.Error("test log", zap.String("error_info", "param error"))
	logger.Sugar().Infof("user name %s", "tom")

	logger.Sugar().Debugw("map data",
		"url", "www.qq.com",
	)
}

func TestProduct(t *testing.T) {
	logger, _ := zap.NewProduction()
	logger.Info("product data")
	logger.Sugar().Warnw("product warning", "reason", "status problem")
	logger.Error("product error")
}

func TestNew1(t *testing.T) {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(os.Stdout), zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	logger.Info("data-info")
	logger.Debug("data-debug")
}

func TestNew2(t *testing.T) {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	file, _ := os.Create("./info.log")
	fileWriter := zapcore.AddSync(file)

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(fileWriter, zapcore.AddSync(os.Stdout)), zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	logger.Info("new2-info")
	logger.Debug("new2-debug")
}

func TestNew3(t *testing.T) {
	lumberWriteSyncer := &lumberjack.Logger{
		Filename: "./date-log.log",
		MaxSize:  500,   // 最大MB
		MaxAge:   28,    // 最大天数
		Compress: false, //Compress确定是否应该使用gzip压缩已旋转的日志文件。默认值是不执行压缩。
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "dateTime"

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(lumberWriteSyncer), zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller(), zap.Fields(zap.String("appName", "ad-server")))
	logger.Info("new3-info")
	logger.Debug("new3-debug")
}
