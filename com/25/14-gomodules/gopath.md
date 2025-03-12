## 一、什么是Go Modules?

Go modules是Go语言的依赖解决方案，发布于Go1.11，成长于Go1.12，丰富于Go1.13，自Go1.14起正式推荐在生产中使用。

Go modules目前已集成在Go的工具链中，只要安装了Go环境，即可使用。它的出现解决了Go 1.11之前几个常见的争议问题：

1. 长久以来Go语言的依赖管理难题。
2. 改变并逐渐替代原有的GOPATH使用模式。
3. 统一社区中其他的依赖管理工具，并提供迁移功能。

## 二、GOPATH的工作模式

Go Modules的目的之一就是淘汰GOPATH，那么GOPATH是什么？为什么在Go1.11前使用GOPATH，而Go1.11后就开始逐步建议使用Go
modules，不再推荐GOPATH的模式了呢？

### (1) Wait is GOPATH?

通过`$ go env`命令查看，示例输出：

```
GOPATH="/home/itheima/go"
...
```

## 三、Go Modules模式

接下来用Go Modules的方式创建一个项目，建议与GOPATH分开，不要将项目创建在GOPATH/src下。

### (1) go mod命令

| 命令              | 作用                 |
|-----------------|--------------------|
| go mod init     | 生成go.mod文件         |
| go mod download | 下载go.mod文件中指明的所有依赖 |
| go mod tidy     | 整理现有的依赖            |
| go mod graph    | 查看现有的依赖结构          |
| go mod edit     | 编辑go.mod文件         |
| go mod vendor   | 导出项目所有的依赖到vendor目录 |
| go mod verify   | 校验一个模块是否被篡改过       |
| go mod why      | 查看为什么需要依赖某模块       |

### (2) go mod环境变量

可通过`go env`命令查看，示例输出：

```
GO111MODULE="auto"
GOPROXY="https://proxy.golang.org,direct"
GONOPROXY=""
GOSUMDB="sum.golang.org"
GONOSUMDB=""
GOPRIVATE=""
...
```