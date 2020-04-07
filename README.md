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
