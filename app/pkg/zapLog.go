package pkg

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func testMain() {

	for i := 1; i < 8000; i++ {
		simpleHttpGet()
		simpleHttpGet()
	}
}

func InitZapLogger() {
	var coreArr []zapcore.Core
	encoder := getEncoder()
	writeSyncer := zapcore.AddSync(GetLogWriter())
	// writeSyncerErr := getLogWriter(viper.GetString("log.error"))
	// err := zapcore.NewCore(encoder, writeSyncerErr, zapcore.ErrorLevel)
	// coreArr = append(coreArr, err)
	info := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)
	coreArr = append(coreArr, info)
	core := zapcore.NewTee(coreArr...)
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(filename string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    5,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func simpleHttpGet() {
	zap.L().Info("Trying to hit GET request for %s url")
	zap.L().Error("Error fetching URL %s : Error = %s url err")
	zap.L().Warn("Success! statusCode = %s for URL %s resp.Status url")
}
