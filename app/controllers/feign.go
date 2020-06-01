package controllers

import (
	"encoding/json"

	"wordgame/tdog/core"
)

type (
	Feign struct {
		Base core.Controller
	}
)

func (feign *Feign) Http() {
	data := feign.Base.Req.Put
	dataJson, _ := json.Marshal(data)
	FeignCore := core.NewFeign()
	FeignCore.Decoder(string(dataJson)).Target()
}
