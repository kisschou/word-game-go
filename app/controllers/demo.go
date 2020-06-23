package controllers

import (
	"wordgame/tdog/core"
)

type Demo struct {
	Base core.Controller
}

func (demo *Demo) Hello() {
	data := []string{"bgbiao", "biaoge"}
	demo.Base.Res.Html(data)
}
