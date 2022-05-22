package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var gRedis *redis.Client

func Init() (err error) {
	host := viper.Get("redis.host")
	port := viper.Get("redis.port")
	password := viper.Get("redis.password")
	database := viper.Get("redis.database")
	poolsize := viper.Get("redis.poolsize")

	gRedis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, int(port.(float64))),
		Password: password.(string),       // no password set
		DB:       int(database.(float64)), // use default DB
		PoolSize: int(poolsize.(float64)),
	})

	_, err = gRedis.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
