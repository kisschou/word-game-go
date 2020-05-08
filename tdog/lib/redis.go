package lib

import (
	RedisModel "github.com/go-redis/redis/v7"
)

type Redis struct {
	Engine *RedisModel.Client
}

func (redis *Redis) NewEngine() {
	conf := new(Config)
	host := conf.Get("cache.master.host").String()
	port := conf.Get("cache.master.port").String()
	pass := conf.Get("cache.master.pass").String()

	redis.Engine = RedisModel.NewClient(&RedisModel.Options{
		Addr:     host + ":" + port,
		Password: pass,
		DB:       0,
	})
}

func (redis *Redis) Ping() bool {
	_, err := redis.Engine.Ping().Result()
	if err != nil {
		logger := Logger{Level: 0, Key: "error"}
		logger.New(err.Error())
		return false
	}
	return true
}

func (redis *Redis) Select(index int) {
	redis.Engine = RedisModel.NewClient(&RedisModel.Options{DB: index})
}
