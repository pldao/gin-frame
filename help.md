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



## 详解



1. 从main.go进入项目，到init选择启动的服务，启动服务为 go run main.go server -R
2. 如果设置-R true就会将api加载到数据库中
3. 在internal/routers/router.go 中的SetAdminApiRoute方法中正式进入路由主体
4. 此时外层已结束需要对SetAdminApiRoute()进行修改路由为自己的路由，开始正式开发
5. 在SetAdminApiRoute方法中统一以api/v1为路由前缀
6. 在SetAdminApiRoute方法中的admin/login方法是通过username和password生成用户端请求的token
7. 提前在数据库中a_admin_user中存入账号和密码:

```sql
-- 创建管理员表
CREATE TABLE IF NOT EXISTS `a_admin_user` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `nickname` varchar(30) NOT NULL DEFAULT '' COMMENT '昵称',
    `username` varchar(30) NOT NULL DEFAULT '' COMMENT '用户名',
    `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
    `mobile` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '手机号',
    `email` varchar(120) NOT NULL DEFAULT '' COMMENT '邮箱',
    `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
    `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1启用 2禁用',
    `is_admin` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否管理员 1是 2不是',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` int NOT NULL DEFAULT '0' COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `a_u_u_d` (`username`,`deleted_at`) USING BTREE,
    UNIQUE KEY `a_u_p_d` (`mobile`,`deleted_at`) USING BTREE,
    UNIQUE KEY `a_u_e_d` (`email`,`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='后台管理用户表';

-- 初始密码 123456
INSERT INTO `a_admin_user` (`id`, `nickname`, `username`, `password`, `mobile`, `email`, `avatar`, `status`, `is_admin`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, '超级管理员', 'super_admin', '$2a$10$OuKQoJGH7xkCgwFISmDve.euBDbOCnYEJX6R22QMeLxCLwdoJ4iyi', '18888888888', 'admin@go-layout.com', '', 1, 1, '2023-05-01 00:00:00', '2023-05-01 00:00:00', 0);

```

这样就可以请求这个接口生成access token

```sh
req:
curl -X POST http://127.0.0.1:9001/api/v1/admin/login\?username\=super_admin\&password\=123456

res:
{"code":0,"msg":"OK","data":{"access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJtb2JpbGUiOiIxODg4ODg4ODg4OCIsIm5pY2tuYW1lIjoi6LaF57qn566h55CG5ZGYIiwiaXNzIjoiZ28tbGF5b3V0Iiwic3ViIjoicGMtYWRtaW4iLCJleHAiOjE3MTYyOTEzMjR9.lnze0vBLbhG8dc14L8UwXcwZoZqe_oUuMY8T1z1pZQs","token_type":"Bearer","expires_at":1716291324},"cost":"117.186042ms"}
```

8.之后的方法中需要带Authorization为Bearer,输入上述生成的token

```sh
curl --location --request GET 'http://127.0.0.1:9001/api/v1/admin-user/get' --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJtb2JpbGUiOiIxODg4ODg4ODg4OCIsIm5pY2tuYW1lIjoi6LaF57qn566h55CG5ZGYIiwiaXNzIjoiZ28tbGF5b3V0Iiwic3ViIjoicGMtYWRtaW4iLCJleHAiOjE3MTYyOTEzMjR9.lnze0vBLbhG8dc14L8UwXcwZoZqe_oUuMY8T1z1pZQs' 
```

这个token除了主动生成，在有效时间内有效，如果刷新时间大于0那么就会在返回体中返回新的token，如果超时没有请求就得重新调用接口生成
