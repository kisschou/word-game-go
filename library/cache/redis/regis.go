package redis

import (
	RedisModel "github.com/go-redis/redis/v7"

	"wordgame/library/config"
	"wordgame/library/logger"
)

type RedisInfo struct{}

var Redis *RedisModel.Client

func init() {
	c := config.Config{}
	redisHost := c.Get("redis.master.host").(string)
	redisPort := c.Get("redis.master.port").(string)
	redisPass := c.Get("redis.master.pass").(string)

	Redis = RedisModel.NewClient(&RedisModel.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPass,
		DB:       0,
	})

	_, err := Redis.Ping().Result()
	if err != nil {
		logger.LoggerToFile("error", err.Error(), 0)
	}
}

func RedisSelectDb(index int) {
	Redis = RedisModel.NewClient(&RedisModel.Options{DB: index})
}
