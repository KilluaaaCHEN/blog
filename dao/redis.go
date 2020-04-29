package dao

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var redisInst *redis.Client
var redisOnce sync.Once

type Redis struct {
	Port     int
	Host     string
	Password string
	Db       int
}

func InstRedis() *redis.Client {

	redisOnce.Do(func() {

		var config Redis
		if err := viper.UnmarshalKey("database.redis", &config); err != nil {
			log.Fatal(err)
		}

		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
			Password: config.Password,
			DB:       config.Db,
		})
		_, err := client.Ping().Result()
		if err != nil {
			panic(err)
		}
		redisInst = client
		fmt.Println("redis connected...")
	})
	return redisInst
}
