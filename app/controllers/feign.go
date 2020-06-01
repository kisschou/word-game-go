package controllers

import (
	"encoding/json"
	"net/http"

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
	code, res := FeignCore.Decoder(string(dataJson)).Target()
	if code == 0 {
		feign.Base.Res.JSON(http.StatusInternalServerError, core.H{
			"code": res,
		})
		return
	}

	feign.Base.Res.String(code, res)
}
