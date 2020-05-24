# word-game-go
仅仅是一个微笑开始

## 说明
本项目主要是参考gin改造而成，改动太多，难以详尽描述。使用go-mod做包管理。

## 引用包
包名 | 版本 | 用途
:--: | :--: | :--:
[go-redis](https://github.com/go-redis/redis) | v7.2.0 | 处理redis连接和redis方法实现
[xorm](https://github.com/xormplus/xorm) | v0.7.9 | 数据库的orm处理，当前主要针对mysql数据库使用
[httprouter](https://github.com/julienschmidt/httprouter) | v1.3.0 | 据说能提高40倍效率的玩意，用于路由处理
[logrus](https://github.com/sirupsen/logrus) | v1.5.0 | 日志处理
[viper](https://github.com/spf13/viper) | v1.6.3 | 配置文件获取，配置项读取，感觉用起来会有卡顿和偶发性的取key失败的问题，有更好的替代方案时考虑更换
[go-swagger](https://github.com/go-swagger/go-swagger) | v0.23.0 | 因为当前版本只针对RESTFul所以swagger还是一个不错的工具

## 结构
```
├── README-LEARN.md                     // 所谓的学习资料
├── README-REDIS.md                     // Redis的key字典
├── README.md                           // md
├── TODO.md                             // 只是概念下的待做清单
├── app                                 // 业务主目录
│   ├── controllers                     // 业务控制器目录
│   │   ├── auth.go
│   │   ├── member.go
│   │   └── ...
│   ├── models                          // 业务模型目录
│   │   ├── user.go
│   │   ├── ...
│   ├── routers                         // 路由目录
│   │   ├── README.md                   // 路由添加规则说明文档
│   │   ├── router.go
│   │   └── ...
│   └── services                        // 服务目录
│       ├── member.go
│       └── ...
├── config                              // 配置目录
│   ├── app.toml                        // 项目基础配置
│   ├── cache.toml                      // 说是缓存配置其实也就是redis的配置
│   ├── database.toml                   // 说是数据库配置其实也就是mysql配置
│   ├── error.toml                      // 说是错误配置，其实也没有用起来
│   └── ...
├── data                                // 数据目录
│   ├── protobuf                        // protobuf文件目录，rpc铺路
│   │   ├── doc                         // 说明文件,说的啥不是我写的,我也不知道
│   │   │   └── api.md
│   │   └── proto                       // proto文件目录
│   │       ├── admin_api.proto
│   │       ├── api.proto
│   │       ├── app_api.proto
│   │       ├── base.proto
│   │       └── ...
│   └── script                          // 脚本目录
│       ├── protobuf-builder.sh
│       └── ...
├── go.mod                              // mod文件
├── go.sum
├── main.go                             // 入口
└── tdog                                // 核心功能实现
    ├── core                            // 核心包
    │   ├── controller.go               // 基础控制器类，希望每个控制器都引用他
    │   ├── jwt.go                      // 一个垃圾到可能会被嘲讽的东西
    │   ├── model.go                    // 基础模型类，希望每个模型都引用他
    │   ├── request.go                  // 请求数据类，所谓的统一入口
    │   ├── response.go                 // 返回数据类，所谓的统一出口
    │   ├── router.go                   // 路由解析类，httprouter作用于此
    │   └── service.go                  // 基础服务类，希望每个服务都引用他
    └── lib                             // 附加方法
        ├── config.go                   // 配置文件获取类，viper作用于此
        ├── crypt.go                    // 加密方法都放在这里
        ├── logger.go                   // 日志类, logrus作用于此
        ├── mysql.go                    // MySQL操作类, xorm作用于此
        ├── redis.go                    // Redis操作类, go-redis作用于此
        └── util.go                     // 基础方法都在这里
```

## 使用说明
```
## 安装golang环境等问题不再重复叙述

## clone项目
shell> git clone https://github.com/kisschou/word-game-go.git
shell> cd word-game-go

## 运行项目
shell> go run .
```

#### 业务添加说明
* 路由
> 路由添加请参考 [app/routers/README.md](#)，里面有详细描述。

* 控制器
> 不懂怎么给描述，索性直接上demo
```
package controllers

import (
    "net/http"

    "wordgame/tdog/services"
    "wordgame/tdog/core"
)

type Demo struct() {
    Base core.Controller // 这里需引用基础控制器文件
}

func (demo *Demo) Hello() {
    name := ""

    // 因为目前没有做自动验证，所以得自己做参数判断
    if _, ok := demo.Base.Req.Params["name"]; ok {
        if len(demo.Base.Req.Params["name"]) > 0 {
            name = demo.Base.Req.Params["name"][0]
        }
    }

    // 请求services
    DemoService := new(services.Demo)
    name = DemoService.GetName(name)

    // 数据返回
    // 返回json
    member.Base.Res.JSON(http.StatusOK, core.H{
        "success" => "Hello " + name,
    })
    // 返回string
    member.Base.Res.String(http.StatusOK, "Hello " + name)
    // 返回xml
    ...
    // 返回data
    ...
}
```

* 服务
> 不懂怎么给描述，索性直接上demo
```
package services

import (
    "time"

    "wordgame/app/models"
    "wordgame/tdog/core"
)

type (
    DemoInfo string

    Demo struct {
        Base core.Service // 这里需引用基础服务文件
    }
)

func (demo *Demo) GetName(name string) (retStr DemoInfo) {
    demo.Base.Redis.NewEngine() // 初始化redis控件
    if demo.Base.Redis.Engine.Exists(name).Val() > 0 { // 判断key是否存在redis
        retStr = demo.Base.Redis.Engine.Get(name).Val() // 存在取出数据返回
        return
    }

    // 不存在，从数据库中获取
    DemoModel := new(models.DemoModel)
    retStr = DemoModel.GetName(name)

    // 数据存储到redis
    // 更多的redis操作请参考go-redis文档
    demo.Base.Redis.Engine.SetNX(name, retStr, time.Duration(0)*time.Second)

    return
}
```

* 模型
> 模型的生成建议使用xorm通过表结构自动生成模型后编写
> xorm 生成模型参考: https://github.com/go-xorm/cmd/blob/master/README.md
```
package models

import (
    "wordgame/tdog/core"
)

type (
    // 这个是模块定义的结构体
    DemoModel struct {
        Base core.Model // 这里需引用基础模型文件
    }

    // 数据库Column的结构体
    // 可以通过xorm工具生成
    Demo struct {
        Name       string      `xorm:"comment('测试名') unique CHAR(40)"`
        RetName    string      `xorm:"comment('测试返回名') CHAR(40)"`
    }
)

func (demoModel *DemoModel) GetName(name) (retStr string) {
    demoInfo := new(Demo)
    demoModel.Base.Sql.NewEngine() // 初始化数据库驱动
    // 从数据库中读取数据
    // 更多的数据库操作请参考xorm文档
    result, _ := demoModel.Base.Sql.Engine.Where("name=?", name).Get(demoInfo)
    if !result {
        // 数据不存在
        return
    }
    retStr = demoInfo.RetName
    return
}
```

## 其他
Kisschou&copy;2020.All Rights.
