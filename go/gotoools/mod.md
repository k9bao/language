# 1. go mod

- [1. go mod](#1-go-mod)
  - [1.1. 简介](#11-简介)
  - [1.2. 命令](#12-命令)
  - [1.3. 实践](#13-实践)
  - [1.4. go.mod,go.sum](#14-gomodgosum)
  - [1.5. 参考资料](#15-参考资料)

## 1.1. 简介

&emsp;&emsp;golang 包管理工具

配置命令( `go env` 可以查看)：

- GO111MODULE
  - GO111MODULE=off，go命令行将不会支持module功能，寻找依赖包的方式将会沿用旧版本那种通过vendor目录或者GOPATH模式来查找。
  - GO111MODULE=on，go命令行会使用modules，而一点也不会去GOPATH目录下查找。
  - GO111MODULE=auto，默认值，go命令行将会根据当前目录来决定是否启用module功能。这种情况下可以分为两种情形：
    - 当前目录在GOPATH/src之外且该目录包含go.mod文件
    - 当前文件在包含go.mod文件的目录下面。

## 1.2. 命令

```bash
$ go help mod
Go mod provides access to operations on modules.

Note that support for modules is built into all the go commands,
not just 'go mod'. For example, day-to-day adding, removing, upgrading,
and downgrading of dependencies should be done using 'go get'.
See 'go help modules' for an overview of module functionality.

Usage:

        go mod <command> [arguments]

The commands are:

        download    download modules to local cache
        edit        edit go.mod from tools or scripts
        graph       print module requirement graph
        init        initialize new module in current directory
        tidy        add missing and remove unused modules
        vendor      make vendored copy of dependencies
        verify      verify dependencies have expected content
        why         explain why packages or modules are needed

Use "go help mod <command>" for more information about a command.
```

| 命令     | 说明                                                                  |
| -------- | --------------------------------------------------------------------- |
| download | download modules to local cache(下载依赖包)                           |
| edit     | edit go.mod from tools or scripts（编辑go.mod)                        |
| graph    | print module requirement graph (打印模块依赖图)                       |
| verify   | initialize new module in current directory（在当前目录初始化mod）     |
| tidy     | add missing and remove unused modules(拉取缺少的模块，移除不用的模块) |
| vendor   | make vendored copy of dependencies(将依赖复制到vendor下)              |
| verify   | verify dependencies have expected content (验证依赖是否正确）         |
| why      | explain why packages or modules are needed(解释为什么需要依赖)        |

## 1.3. 实践

```cmd
mkdir Gone
cd Gone
go mod init Gone
```

go.mod文件一旦创建后，它的内容将会被go toolchain全面掌控。
go toolchain会在各类命令执行时，比如go get、go build、go mod等修改和维护go.mod文件。

创建 main.go文件

```golang
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

执行 go run main.go 运行代码会发现 go mod 会自动查找依赖自动下载

## 1.4. go.mod,go.sum

go.mod 提供了module, require、replace和exclude 四个命令

module 语句指定包的名字（路径）
require 语句指定的依赖项模块
replace 语句可以替换依赖项模块,
exclude 语句可以忽略依赖项模块

```bash
module gorm

go 1.12

require (
    gorm.io/driver/sqlite v1.1.4
    gorm.io/gorm v1.20.7
)

replace (
    github.com/eclipse/paho.mqtt.golang v1.2.0 => ./mqtt/paho.mqtt.golang
)

```

go module 安装 package 的原則是先拉最新的 release tag，若无tag则拉最新的commit
go 会自动生成一个 go.sum 文件来记录 dependency tree

`go list -m -u all` 来检查可以升级的package
`go get -u need-upgrade-package` 升级后会将新的依赖版本更新到go.mod *
`go get -u` 升级所有依赖

默认会下载依赖包到 `$GOPATH/pkg/mod` 目录

## 1.5. 参考资料

1. [go mod使用](https://www.jianshu.com/p/760c97ff644c)
