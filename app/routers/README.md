## 路由配置文件

#### 说明
    因路由中一定要携带方法，又不存在字符串转方法，考虑过用反射将请求和方法构建对应关系，但是实际实际上会更加麻烦，所以取消了。所以路由没有办法写成配置文件，只能是一个go脚本的存在。

#### 请求方式
    本次版本暂时只提供`RESTFul`方式请求，所以也支持RESTFul的基本请求方式: `GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`, `PATCH`。

#### 配置规则
    路由是按照Gin框架中的路由设计的，也是采用httprouter包进行解析。因为对路由指向的方法都要设置一个`Context`接收参数的方式感觉到不够优雅，所以改变了规则，将参数直接Set到了基础控制器中，所有的控制器都继承基础控制器，从而实现参数的获取。

##### 单路由
```
    // 初始化路由引擎
	r := core.NewEngine()

    // 初始化指向包
	var member controller.Member

    // 设置POST路由
    r.POST("/hello", member.Login, &member.Base)

    说明
    r.POST(访问路径, 触发方法, 基础控制器)
```
|| 注: 基础控制因为不确定方法内有没有改名所以没有定死，在单路由配置中必须填写。

##### 路由组
```
    // 初始化路由引擎
	r := core.NewEngine()

    // 初始化指向包
	var member controller.Member

    // 设置路由组
	memberRouter := r.Group("/member", &member.Base)
	{
		memberRouter.GET("/login", member.Login)
		memberRouter.POST("/login", member.Login)
		memberRouter.GET("/hello", member.Hello)
	}

    // 说明
	memberRouter := r.Group(组访问路径, 基础控制器)
	{
		memberRouter.GET(子访问路径, 触发方法 ...{, 基础控制器})
	}
```
|| 注意:
|| 1. 在路由组配置中，组配置中的基础控制器必填，而子路由中的选填写。
|| 2. 发起请求时，请求地址为: `http(s)://请求地址:请求端口/组访问路径/子访问路径`
