package pkg

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type Redis struct {
	DB *redis.Client
}

func (r *Redis) NewRedisClient() *redis.Client {
	config := viper.GetStringMap("redis")
	r.DB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config["host"], config["port"]),
		Password: viper.GetString("redis.password"), // 没有密码，默认值
		DB:       viper.GetInt("redis.database"),    // 默认DB 0
	})
	return r.DB
}
