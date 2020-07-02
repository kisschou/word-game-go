急需修改解决问题:
0. 内部对外发送请求
1. 日志模块重构 => 上ELK
2. 新开线程启动WebSocket服务 => 加上RabbitMQ
3. Swagger 或 RESTFul换gRPC
4. 自动验证..
5. 因viper模块不够稳定, 考虑替换viper模块, 感觉可以自己利用sqlite构建一个。有空试试。
6. 加入Oauth2

整理微服务:
1. 用户服务 - feature_kisschou_user
2. 鉴权服务 - feature_kisschou_auth
3. 日志服务 - logger_service
4. 短信服务 - sms_service
5. 验证码服务 - verifycode_service
6. 调试工具 - doctool
7. 网关服务 - feature_kisschou_feign
8. 许可服务 (鉴权子服务) - permission_service
9. 文件存储服务 - feature_kisschou_file

继续做:
so many..
