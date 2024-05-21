```
.
|——.gitignore
|——go.mod
|——go.sum
|——main.go    // 项目入口 main 包
|——LICENSE
|——README.md
|——boot    // 项目初始化目录
|  └──boot.go
|——config    // 这里通常维护一些本地调试用的样例配置文件
|  └──autoload    // 配置文件的结构体定义包
|     └──app.go
|     └──logger.go
|     └──mysql.go
|     └──redis.go
|     └──server.go
|  └──config.example.ini    // .ini 配置示例文件
|  └──config.example.yaml    // .yaml 配置示例文件
|  └──config.go    // 配置初始化文件
|——data    // 数据初始化目录
|  └──data.go
|  └──mysql.go
|  └──redis.go
|——internal    // 该服务所有不对外暴露的代码，通常的业务逻辑都在这下面，使用internal避免错误引用
|  └──controller    // 控制器代码
|     └──v1
|        └──auth.go    // 完整流程演示代码，包含数据库表的操作
|        └──helloword.go    // 基础演示代码
|     └──base.go
|  └──middleware    // 中间件目录
|     └──cors.go
|     └──logger.go
|     └──recovery.go
|     └──requestCost.go
|  └──model    // 业务数据访问
|     └──admin_users.go
|     └──base.go
|  └──pkg    // 内部使用包
|     └──errors    // 错误定义
|        └──code.go
|        └──en-us.go
|        └──zh-cn.go
|     └──logger    // 日志处理
|        └──logger.go
|     └──response    // 统一响应输出
|        └──response.go
|  └──routers    // 路由定义
|     └──apiRouter.go
|     └──router.go
|  └──service    // 业务逻辑
|     └──auth.go
|  └──validator    // 请求参数验证器
|     └──form    // 表单参数定义
|        └──auth.go
|     └──validator.go
|——pkg    // 可以被外部使用的包
|  └──convert    // 数据类型转换
|     └──convert.go
|  └──utils    // 帮助函数
|     └──utils.go
```
# migrate

```
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate -database 'mysql://root:sky@tcp(127.0.0.1:3307)/ginframe?charset=utf8mb4&parseTime=True&loc=Local' -path data/migrations up

rm 表
migrate -database 'mysql://root:sky@tcp(127.0.0.1:3307)/ginframe?charset=utf8mb4&parseTime=True&loc=Local' -path data/migrations down

```

# run
```
1. 先打开config中的mysql，配置mysql创建ginframe数据库
2. 执行数据库migrate
3.  go run main.go server -R true (将api加载到数据库中)
4.  go run main.go server 

```