# tags通用实现

## 目录结构

imply
    body
        body1
            body1Imp.go //依赖body.go中定义的结构体,body1的具体实现
        body2
            body2Imp.go //依赖body.go中定义的结构体，body2的具体实现
        body.go
    driver
        driver_body1.go //依赖body.go、body1.go、driver.go
        driver_body2.go //依赖body.go、body2.go、driver.go
        driver.go       //依赖body.go

main.go //直接调用driver的接口

## 结构说明

driver_body1.go 中包含 tags 编译选项，来控制是否编译 body1
driver_body1.go 必须在 driver 目录，如果放在 body1 中的话，body1目录不会被依赖，就不会被编译了
而放在 driver 目录的话，driver.go被调用的时候，driver目录被编译，根据是否编译 driver_body1.go 来触发是否编译 body1

## 编译命令

go build -a -v -work -tags="body1"
go build -a -v -work -tags="body2"
go build -a -v -work -tags="body1 body2"
