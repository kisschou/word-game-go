package session

import (
	"github.com/rs/xid"
)

func New() string {
	id := xid.New()
	sessionId := id.String()

	return sessionId
}
