package core

import (
	"wordgame/tdog/lib"
)

type (
	Error struct {
	}
)

func (e *Error) GetError(errCode string) (errMsg string) {
	ConfigLib := new(lib.Config)
	errMsg = errCode
	if ConfigLib.Get("error." + errCode).IsExists() {
		errMsg = ConfigLib.Get("error." + errCode).String()
	}
	return
}
