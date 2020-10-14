package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/kisschou/tdog"
)

type (
	Feign struct {
		Base tdog.Controller
	}
)

func (feign *Feign) Http() {
	data := feign.Base.Req.Put
	dataJson, _ := json.Marshal(data)
	FeignCore := tdog.NewFeign()
	code, res := FeignCore.Decoder(string(dataJson)).Target()
	if code == 0 {
		feign.Base.Res.JSON(http.StatusInternalServerError, tdog.H{
			"code": res,
		})
		return
	}

	feign.Base.Res.String(code, res)
}

func (feign *Feign) Ping() {
	feign.Base.Res.String(http.StatusOK, "Pong~")
}
