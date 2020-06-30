# word-game-go
仅仅是一个微笑开始

## 说明
文件服务

## 服务配置说明
配置文件位于 `config/file.toml` 中。
字段 | 说明
-- | --
file_request_url | 文件请求地址, 即获取文件的地址
file_save_path | 文件存储位置(基础目录), 该目录会与相对地址进行拼接

## 服务启动说明
默认端口为 `8006`, 如需修改在 `config/app.toml` 中进行修改，修改后需同时修改 `config/api_url.toml` 和 网关服务的 `config/api_url.toml` 中的相关配置。
```
## 安装golang环境等问题不再重复叙述

## clone项目
shell> git clone -b feature_kisschou_file https://github.com/kisschou/word-game-go.git <目录名>
shell> cd word-game-go ( <目录名> )

## 运行项目
shell> go run .
## 后台运行
shell> go run . -d=true

## 项目打包
## ...Linux or Mac OS:
shell> go build .
## 会在目录下生成一个名为wordgame的二进制文件
## 运行:
shell> ./wordgame
## 后台运行
shell> ./wordgame -d=true
## ...Windows:
shell> CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build .
## 会生成一个名为 wordgame.exe 的可执行文件
## 在Windows下双击即可运行
```

## 其他
Kisschou&copy;2020.All Rights.
