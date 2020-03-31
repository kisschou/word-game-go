package redis

import (
	"fmt"
	RedisModel "github.com/go-redis/redis/v7"
	"github.com/rs/xid"
	"time"
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

func SessionId() string {
	// 切换到db(0)
	RedisSelectDb(0)

	key := "session:"
	id := xid.New()
	sessionId := id.String()

	for {
		sessionId = id.String()
		key += sessionId
		if Redis.Exists(key).Val() == 0 {
			Redis.SetNX(key, "", time.Duration(86400)*time.Second)
			break
		}
	}

	return sessionId
}
