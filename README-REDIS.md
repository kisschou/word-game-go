# Redis字段字典

​	基于后期存在Redis使用频繁导致键值过多出现用时查不到, 或者命名各种不统一、存储位置怪异造成管理麻烦, 整理繁琐, 故而一开始构建此文件约束。

------

#### DB(0) 用户数据

> 主要用于存储用户相关的数据

* session数据

| key  | type | desc |
| ---- | ---- | ---- |
| session:{:token} | string | session在用 |
| session:user:{:openId} | hash | 在线用户信息列表 |

#### DB(1)