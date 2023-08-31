package pkg

import (
	"io"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

func GetLogWriter() io.Writer {
	if IsDev() {
		return os.Stdout
	}
	return &lumberjack.Logger{
		Filename:   viper.GetString("log.info"),
		MaxSize:    5,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
}

func GetLogErrorWriter() io.Writer {
	if IsDev() {
		return os.Stdout
	}
	return &lumberjack.Logger{
		Filename:   viper.GetString("log.error"),
		MaxSize:    5,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
}
