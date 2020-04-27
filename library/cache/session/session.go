package session

import (
	"fmt"

	"github.com/rs/xid"

	. "wordgame/library/cache/redis"
	"wordgame/library/logger"
)

type Session interface {
}

type Config struct {
}

func (conf *Config) New() (sid string) {
	// RedisSelectDb(0)

	id := xid.New()
	sessionId := id.String()

	return sessionId
}
