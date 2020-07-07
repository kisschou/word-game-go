package controllers

import (
	"net/http"

	"wordgame/app/services"
	"wordgame/tdog/core"
)

type VerificationCode struct {
	Base core.Controller
}

// swagger:operation POST /verification_code/get verification_code get
// ---
// summary: 获取验证码授权码
// description: 获取验证码授权码
// parameters:
// - name: Authorization
//   in: header
//   description: 授权信息
//   type: string
//   required: false
// - name: service
//   in: body
//   description: 服务标识
//   type: string
//   required: true
// - name: type
//   in: body
//   description: 验证码类型
//   type: string
//   required: true
// - name: extras
//   in: body
//   description: 附加信息
//   type: string
//   required: true
// - name: ip_addr
//   in: body
//   description: 请求发起者ip地址
//   type: string
//   required: true
// responses:
//   200: repoResp
//   401: badReq
func (verificationCode *VerificationCode) Get() {
	service := ""
	verificationCodeType := ""
	extras := ""
	ipAddr := verificationCode.Base.Req.IpAddr

	if _, ok := verificationCode.Base.Req.Params["service"]; ok {
		if len(verificationCode.Base.Req.Params["service"]) > 0 {
			service = verificationCode.Base.Req.Params["service"][0]
		}
	}

	if _, ok := verificationCode.Base.Req.Params["type"]; ok {
		if len(verificationCode.Base.Req.Params["type"]) > 0 {
			verificationCodeType = verificationCode.Base.Req.Params["type"][0]
		}
	}

	if _, ok := verificationCode.Base.Req.Params["extras"]; ok {
		if len(verificationCode.Base.Req.Params["extras"]) > 0 {
			extras = verificationCode.Base.Req.Params["extras"][0]
		}
	}

	if _, ok := verificationCode.Base.Req.Params["ip_addr"]; ok {
		if len(verificationCode.Base.Req.Params["ip_addr"]) > 0 {
			ipAddr = verificationCode.Base.Req.Params["ip_addr"][0]
		}
	}

	if len(service) < 1 || len(verificationCodeType) < 1 || len(extras) < 1 {
		verificationCode.Base.Res.JSON(http.StatusInternalServerError, core.H{
			"code": "ERROR_REQUEST_PARAMS",
		})
		return
	}

	VerificationCodeService := new(services.VerificationCode)
	VerificationCodeService.IpAddr = verificationCode.Base.Req.IpAddr
	captchaAccess := VerificationCodeService.Build(service, verificationCodeType, extras, ipAddr)
	verificationCode.Base.Res.JSON(http.StatusOK, core.H{
		"message":        "success",
		"captcha_access": captchaAccess,
	})
}

// swagger:operation POST /verification_code/getPic verification_code getPic
// ---
// summary: 获取验证码图片
// description: 获取验证码图片
// parameters:
// - name: Authorization
//   in: header
//   description: 授权信息
//   type: string
//   required: false
// - name: captcha_access
//   in: body
//   description: 验证码授权码
//   type: string
//   required: true
// responses:
//   200: repoResp
//   401: badReq
func (verificationCode *VerificationCode) GetPic() {
	captchaAccess := ""
	if _, ok := verificationCode.Base.Req.Params["captcha_access"]; ok {
		if len(verificationCode.Base.Req.Params["captcha_access"]) > 0 {
			captchaAccess = verificationCode.Base.Req.Params["captcha_access"][0]
		}
	}

	if len(captchaAccess) < 1 {
		verificationCode.Base.Res.JSON(http.StatusInternalServerError, core.H{
			"code": "ERROR_REQUEST_PARAMS",
		})
		return
	}

	VerificationCodeService := new(services.VerificationCode)
	code, err := VerificationCodeService.GetCode(captchaAccess)
	if err != nil {
		verificationCode.Base.Res.JSON(http.StatusInternalServerError, core.H{
			"code": err.Error(),
		})
		return
	}
	verificationCode.Base.Res.Captcha(code)
}

// swagger:operation POST /verification_code/verify verification_code verify
// ---
// summary: 验证码校验
// description: 验证码校验
// parameters:
// - name: Authorization
//   in: header
//   description: 授权信息
//   type: string
//   required: false
// - name: captcha_access
//   in: body
//   description: 验证码授权码
//   type: string
//   required: true
// - name: code
//   in: body
//   description: 验证码
//   type: string
//   required: true
// responses:
//   200: repoResp
//   401: badReq
func (verificationCode *VerificationCode) Verify() {
	code := ""
	captchaAccess := ""

	if _, ok := verificationCode.Base.Req.Params["captcha_access"]; ok {
		if len(verificationCode.Base.Req.Params["captcha_access"]) > 0 {
			captchaAccess = verificationCode.Base.Req.Params["captcha_access"][0]
		}
	}

	if _, ok := verificationCode.Base.Req.Params["code"]; ok {
		if len(verificationCode.Base.Req.Params["code"]) > 0 {
			code = verificationCode.Base.Req.Params["code"][0]
		}
	}

	if len(captchaAccess) < 1 || len(code) < 4 {
		verificationCode.Base.Res.JSON(http.StatusInternalServerError, core.H{
			"code": "ERROR_REQUEST_PARAMS",
		})
		return
	}

	VerificationCodeService := new(services.VerificationCode)
	err := VerificationCodeService.Verify(captchaAccess, code)
	if err != nil {
		verificationCode.Base.Res.JSON(http.StatusInternalServerError, core.H{
			"code": err.Error(),
		})
		return
	}
	verificationCode.Base.Res.JSON(http.StatusOK, core.H{
		"message": "success",
	})
}

func (verificationCode *VerificationCode) Ping() {
	verificationCode.Base.Res.String(http.StatusOK, "Pong")
}
