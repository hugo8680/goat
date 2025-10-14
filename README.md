# GOAT SCAFFOLD
> 基于Go + Gin + Gorm的快速开发脚手架
> 
> v2025-10-15
## 基础信息
* Golang ^1.25
* MySQL 8.0.x
* Redis 7.x.x
* Gin ^1.11.0
* Gorm ^1.31.0

## 关于
* 项目基于 [Ruoyi-Vue项目](https://gitee.com/y_project/RuoYi-Vue)设计扩展
* 项目基于 [Ruoyi-Go项目](https://github.com/mengxiangyu996/ruoyi-go)重构
* 项目前端基于 [Ruoyi-Vue3](https://github.com/yangzongzhuan/RuoYi-Vue3)重构
* 项目追求简单易用，可用来快速开发小型项目，未设计较复杂的功能及第三方API接入

## 模块
* `framework` 定义项目基础配置，基础响应解构，路由，gin.Engine初始化配置及数据源接入配置
* `common` 定义若干基础工具包和通用常量
* `middleware` 定义通用中间件
* `model` 定义通用domain及dto模型
* `route` 定义基础路由
* `service` 定义通用service
* `api/controller` 定义通用接口
* `api/validator` 定义接口参数校验方法
* `application.yml` 配置文件

## 启动方法
普通运行
``` shell
    go run main.go
```
编译打包
```shell
    go build main.go
```

## 编写规则

1. 路由可以定义新文件在RunServer方法传入可变参数
2. 若使用其他数据库可在framework下添加新的connector文件
3. controller中尽量只涉及参数校验，实际业务逻辑交由service处理