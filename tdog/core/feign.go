package core

import (
	_ "wordgame/tdog/lib"
)

type (
	Feign struct {
		Method      string
		BaseUrlType string
		Url         string
		Header      string
		Body        string
	}
)

func (feign *Feign) NewClient() {
}
