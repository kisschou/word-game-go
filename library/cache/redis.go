package redis

import (
	"fmt"
	RedisModel "github.com/go-redis/redis/v7"
)

type RedisInfo struct{}

var Redis *RedisModel.Client

func init() {
	Redis = RedisModel.NewClient(&RedisModel.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := Redis.Ping().Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pong)
	}
}

func RedisSelectDb(index int) {
	Redis = RedisModel.NewClient(&RedisModel.Options{DB: index})
}
