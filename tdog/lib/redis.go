package lib

import (
	"os"

	RedisModel "github.com/go-redis/redis/v7"
)

type Redis struct {
}

var RedisEngine *RedisModel.Client

func init() {
	var conf Config
	host := conf.Get("cache.master.host").(string)
	port := conf.Get("cache.master.port").(string)
	pass := conf.Get("cache.master.pass").(string)

	RedisEngine = RedisModel.NewClient(&RedisModel.Options{
		Addr:     host + ":" + port,
		Password: pass,
		DB:       0,
	})

	_, err := RedisEngine.Ping().Result()
	if err != nil {
		logger := Logger{Level: 0, Key: "error"}
		logger.New(err.Error())
		os.Exit(0)
	}
}

func (redis *Redis) Select(index int) {
	RedisEngine = RedisModel.NewClient(&RedisModel.Options{DB: index})
}
