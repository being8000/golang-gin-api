package pkg

import (
	"fmt"

	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func IsDev() bool {
	return viper.GetString("env") == "dev"
}

func InitConfig() {
	env := pflag.String("env", "dev", "system environment")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	log.Println("Configuration initializing, environment: ", *env)
	if !IsDev() {
		gin.SetMode(gin.ReleaseMode)
	}
	file := fmt.Sprintf("./app/conf/%s.yaml", *env)
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Fatal error config file: ", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("Configuration changed!")
	})
	log.Println("Configuration initialized successfully")
}
