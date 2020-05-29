package controllers

import (
	"wordgame/tdog/core"
)

type (
	Feign struct {
		Base core.Controller
	}
)

func (feign *Feign) Http() {
}
