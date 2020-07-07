# Redis字段字典

​	基于后期存在Redis使用频繁导致键值过多出现用时查不到, 或者命名各种不统一、存储位置怪异造成管理麻烦, 整理繁琐, 故而一开始构建此文件约束。

------

#### DB(0) 用户数据

> 主要用于存储用户相关的数据

* 用户数据

| key  | type | desc |
| ---- | ---- | ---- |
| user:openid:{:USER_OPEN_ID} | string | openId和用户id关联 |
| user:info:{:USER_ID} | hash | 用户信息 |

* 验证码数据
| key  | type | desc |
| ---- | ---- | ---- |
| captcha:{:CAPTCHA_ACCESS} | string | 验证码授权码关联验证码, 存活期默认3分钟 |

#### DB(1)
