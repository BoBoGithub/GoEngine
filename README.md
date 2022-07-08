# Go的Web项目 - Gin框架实例

## 1. 启动方式:

    go run ./cmd/main.go --config ./config/engine-dev.yml

## 2. 目录结构:
    ├── README.md
    ├── go.mod
    ├── go.sum
    ├── cmd
    │   └── main.go
    ├── routers
    │   └── router.go
    ├── controller
    │   └── Index.go
    ├── middleware
    │   ├── CheckLogin.go
    │   └── ExceptionMiddleWare.go
    ├── service
    │   └── IndexService.go
    └── model
    │   └── UserGameMode.go
    ├── config
    │   ├── engine-dev.yml
    │   └── engine-gray.yml
    ├── libs
    │   ├── db
    │   │   └── DBPool.go
    │   ├── e
    │   │   └── constants.go
    │   ├── helper
    │   │   ├── config.go
    │   │   └── util.go
    │   └── response
    │       └── response.go
    ├── resources
    │   ├── statics
    │   │   └── play.jpg
    │   └── templates
    │       ├── index
    │       │   └── index.html
    │       ├── public
    │       │   ├── footer.html
    │       │   └── header.html
    │       └── user
    │           └── index.html
    └── entity
        └── EngineStruct.go


## 更新记录:

#### 2022/07/08 新提交项目代码至GitHub
