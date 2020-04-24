# word-game-go
仅仅是一个微笑开始

Reference
https://juejin.im/post/5d2a85f5f265da1bd2611b31

待处理
1. Protobuf
2. 登录的Session存储及使用问题:
    a. 用户session:{:token} 没有过期
    b. 用户session:{:openId} 有一天的过期时间
    如果session:{:openId}过期了, 这时候用户重新登录, 这时候用户之前的session:{:token}, 如果没有再次请求不会自动清理, 造成脏数据.


## Other
1. [gojsonq](https://github.com/thedevsaddam/gojsonq) - JSON处理库 [相关文档](https://juejin.im/post/5e53e2efe51d4527196d5352)
2. [copier](https://github.com/jinzhu/copier) - Struct|Map赋值 [相关文档](https://juejin.im/post/5e6b81e551882549281c08c1)
3. [watermill](https://github.com/ThreeDotsLabs/watermill) - 异步消息解决方案, 它支持消息重传、保存消息, 后启动的订阅者也能收到前面发布的消息 [相关文档](https://juejin.im/post/5e5b6e80f265da571d2f2cc4)
4. [email](https://github.com/jordan-wright/email) - 邮件发送 [相关文档](https://juejin.im/post/5e4902b851882549327a41a2)
5. [carbon](https://github.com/uniplaces/carbon) - 时间扩展库 [相关文档](https://juejin.im/post/5e4793dbf265da575d20d488)
6. [cast](https://github.com/spf13/cast) - 类型转换库, 用于将一个类型转为另一个类型 [相关文档](https://juejin.im/post/5e39590d518825490455c385)
7. [gentleman](https://https://gopkg.in/h2non/gentleman.v2) - 功能齐全、插件驱动的 HTTP 客户端 [相关文档](https://juejin.im/post/5e8dee63518825736b749555?utm_source=gold_browser_extension)
8. [logrus](https://github.com/sirupsen/logrus) - 完全兼容标准的log库，还支持文本、JSON 两种日志输出格式 [相关文档](https://juejin.im/post/5e3e768a6fb9a07ccd51793e)
9. [viper](https://github.com/spf13/viper) - 配置解决方案，拥有丰富的特性 [相关文档](https://juejin.im/post/5e24f1bc518825263237edb2)
10. [cobra](https://github.com/spf13/cobra/cobra) - 命令行程序库, 可以用来编写命令行程序 [相关文档](https://juejin.im/post/5e22cfb35188252ca21bb781)
11. [flag](https://golang.org/pkg/flag/) - 用于解析命令行选项 [相关文档](https://juejin.im/post/5e1b33da51882536a627f17b)
12. [go-app](https://github.com/maxence-charriere/go-app/v6) - 是一个使用 Go + WebAssembly 技术编写渐进式 Web 应用的库 [相关文档](https://juejin.im/post/5ea1ac82f265da47b55508b7?utm_source=gold_browser_extension)

## Docs
0. [优秀的开源项目](https://juejin.im/entry/5c00e5fce51d4550c76d9097)
1. [命名规范](https://juejin.im/post/5c6b6c266fb9a04a08223d6c)
2. [高级数据类型](https://juejin.im/post/5e88c68151882573716a8f88)
3. [学习使用 Go 的反射](https://juejin.im/post/5e23c94c5188254db85f01c2)
4. [go中的引用类型](https://juejin.im/post/5e6d8d9e6fb9a07cc97db58c)
5. [深入学习解析HTTP请求](https://juejin.im/post/5e547c2a518825495d69acfc)
6. [http server](https://juejin.im/post/5dd11baff265da0c0c1fe813)
7. [Golang应付百万级请求/分钟](https://juejin.im/post/5db1464b6fb9a0202a261ca9)
8. [Golang的pprof](https://juejin.im/post/5e5b6591518825492f771540)
9. [WebSocket 双端实践（iOS/ Golang）](https://juejin.im/post/5e450b21f265da573f356711)
10. [分布式从ACID、CAP、BASE的理论推进](https://juejin.im/post/5e7d5cb36fb9a03c75752ec0)
11. [Struct超详细讲解](https://juejin.im/post/5ca2f37ce51d4502a27f0539)
12. [微服务统一认证与授权的 Go 语言实现](https://juejin.im/post/5e305be9e51d4531220273e2)
13. [Go 语言内存分配器的实现原理](https://juejin.im/entry/5e5c69af51882549063a8e13)
14. [用Golang构建gRPC服务](https://juejin.im/post/5d994445e51d45782935346b)
15. [go interface详细介绍](https://juejin.im/post/5d8877f1f265da03986c311c)
16. [Golang的sync.WaitGroup 实现逻辑和源码解析](https://juejin.im/post/5e5b62f86fb9a07cb1578fda)
17. [Go 并发模式](https://juejin.im/post/5e554edb6fb9a07ca453439d)
18. [提升科研远程办公效率：代码同步工具sync-go](https://juejin.im/post/5e45203fe51d4526d43f2818)
19. [如何"优雅"地发布 go module 模块](https://juejin.im/post/5e4ccabf6fb9a07ca24f49d4)
20. [Go 实现百万 WebSocket 连接](https://juejin.im/post/5d48f1cd6fb9a06b233ca719)
21. [手把手教你写一个完美的Golang Dockerfile](https://juejin.im/post/5e569409e51d4526ca15ce56)
22. [探究 Go 语言 defer 语句的三种机制](https://juejin.im/post/5e5b4a53f265da570829ed35)
23. [database/sql数据库连接池](https://juejin.im/post/5d624abde51d45621655352c)
24. [理解Sync.Pool的设计思想](https://juejin.im/post/5de4d8a05188256e8d33c7d1)
25. [Golang处理PDF](https://juejin.im/post/5e0c698d5188253aaf656925)
26. [golang核心原理-协程栈](https://juejin.im/post/5da7385ae51d45782a478d2d)
27. [搞懂Go垃圾回收](https://juejin.im/post/5d56b47a5188250541792ede)
28. [走进Golang之运行与Plan9汇编](https://juejin.im/post/5ddde5eef265da05ec6b62ec)
29. [通用抽奖工具之系统设计](https://juejin.im/post/5e0d52906fb9a0481759db76)
30. [通过gin理解中间件](https://juejin.im/post/5e8ed66be51d4546d23c0c12)
31. [Go Web编程--SecureCookie实现客户端Session管理](https://juejin.im/post/5e68b927e51d4526f76eccdf)
32. [Go Web 编程--如何确保Cookie数据的安全传输](https://juejin.im/post/5e64a448518825490966eac5)
