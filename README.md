# word-game-go
仅仅是一个微笑开始

## 说明
这是一个`Feign`的服务，就是一个`Gateway`服务。项目依托在`word-game-go`项目的基础上构建，实现的方法已经并入核心包。

## 使用说明
服务核心脚本位于`wordgame/tdog/core/feign.go`, 实现了请求请求体拆解和转发(通过`wordgame/tdog/lib/http_request.go`)的简单功能。

#### 服务启动
脚本 | 说明
-- | --
`wordgame/tdog/core/feign.go` | 核心方法实现
`wordgame/config/api_url` | 转发链接设置

```
# 具体启动服务方式:
# 1. 在controllers中构建一个服务
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
	data := feign.Base.Req.Put // 接收请求数据, 目前只支持application/json
	dataJson, _ := json.Marshal(data)
	FeignCore := core.NewFeign() // 继承核心包中的Feign方法
	code, res := FeignCore.Decoder(string(dataJson)).Target() // 封装数据并分发请求
    // 返回数据说明:
    // code  -  状态码, 为0是脚本错误，其他都是请求服务返回状态码
    // res   -  返回数据, code为0是脚本错误信息，其他都是请求服务返回的信息

	if code == 0 { // 脚本错误，设定code值自动进行错误处理
		feign.Base.Res.JSON(http.StatusInternalServerError, core.H{
			"code": res,
		})
		return
	}

    // 非脚本错误直接返回结果
	feign.Base.Res.String(code, res)
}

# 2. 在`wordgame/app/routers/router.go`中添加路由
var FeignController controllers.Feign
feignRouter := r.Group("/feign", &FeignController.Base)
{
    feignRouter.POST("/http", FeignController.Http)
}

# 3. Test & Run Service.
shell> go run .
## Run in background.
shell> go run -d=true .

## Build & Run Service.
shell> go build -o gateway-service .
shell> ./gateway-service # Run
shell> ./gateway-service -d=true # Run in background
```
> 本项目在windows下运行会报config找不到的问题，这是因为Linux和Windows下路径分隔符"/"和"\\"的问题。

#### 服务请求规则
请求走的也是restful请求，验签之类的问题跟框架内的方式一致不详述。请求体为 `json` 格式, 共包含 `api_key`、`header`、`body` 三个部分。

参数 | 说明 | 类型
-- | -- | --
api_key | 请求接口编号 | string
header | 请求头部 | json
body | 请求体 | json

```
## 示例:
## 请求登陆接口
POST {:IP}:{:PORT}/feign/http
BODY {
    "api_key": "api_2001",
    "header": {
        "Content-Type": "application/json",
        "Authorization": "eyJhbGciOiJIUzI1NiIsInR5cGUiOiJKV1QifQ==.eyJkYXRhIjp7Im9wZW5faWQiOiJzZUd4TGNodXJRcUJzWENHIn0sImV4cCI6NzIwMCwiaXRhIjoxNTkxOTUwMjc5fQ==.756dffe3d786cfc9df042377accaf3b7ebe45fadcd1c9262226b944d71c9e69c"
    },
    "body": {
        "username": "admin",
        "password": "111111"
    }
}
```

## 其他
Kisschou&copy;2020.All Rights.
