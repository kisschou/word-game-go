package controllers

import (
	"net/http"

	"wordgame/tdog/core"
)

type Logger struct {
	Base core.Controller
}

// swagger:operation POST /logger/get logger get
// ---
// summary: 获取日志数据
// description: 获取日志数据
// parameters:
// - name: Authorization
//   in: header
//   description: 授权信息
//   type: string
//   required: true
// - name: user_id
//   in: body
//   description: 用户id
//   type: string
//   required: false
// responses:
//   200: repoResp
//   401: errMsg
func (logger *Logger) Get() {
	member.Base.Res.JSON(http.StatusOK, core.H{
		"message": "Connection success",
	})
}

// swagger:operation POST /logger/set logger set
// ---
// summary: 写入日志数据
// description: 写入日志数据
// parameters:
// - name: Authorization
//   in: header
//   description: 授权信息
//   type: string
//   required: true
// - name: user_id
//   in: body
//   description: 用户id
//   type: string
//   required: false
// responses:
//   200: repoResp
//   401: errMsg
func (logger *Logger) Set() {
	member.Base.Res.JSON(http.StatusOK, core.H{
		"message": "Connection success",
	})
}

func (logger *Logger) Ping() {
	logger.Base.Res.String(http.StatusOK, "Pong")
}
