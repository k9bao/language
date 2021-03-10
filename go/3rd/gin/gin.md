# 1. gin

- [1. gin](#1-gin)
  - [1.1. 简介](#11-简介)
  - [1.2. header](#12-header)
    - [1.2.1. 设置cookie](#121-设置cookie)
  - [1.3. 路由规则](#13-路由规则)
    - [1.3.1. 简单路由](#131-简单路由)
    - [1.3.2. 路由分组](#132-路由分组)
    - [1.3.3. 静态文件路径](#133-静态文件路径)
    - [1.3.4. 重定向](#134-重定向)
  - [1.4. 获取query和form表单数据](#14-获取query和form表单数据)
  - [1.5. 接收文件](#15-接收文件)
  - [1.6. 中间件](#16-中间件)
    - [1.6.1. 自定义中间件](#161-自定义中间件)
    - [1.6.2. 中间件中使用Goroutines](#162-中间件中使用goroutines)
  - [1.7. 日志](#17-日志)
    - [1.7.1. 中间件日志格式](#171-中间件日志格式)
    - [1.7.2. 自定义路由日志](#172-自定义路由日志)
  - [1.8. 模型绑定和验证](#18-模型绑定和验证)
    - [1.8.1. 自定义验证器](#181-自定义验证器)
    - [1.8.2. 绑定Get参数或者Post参数](#182-绑定get参数或者post参数)
    - [1.8.3. 绑定uri](#183-绑定uri)
    - [1.8.4. 绑定HTML复选框](#184-绑定html复选框)
    - [1.8.5. 绑定Post参数](#185-绑定post参数)
    - [1.8.6. XML、JSON、YAML和ProtoBuf 渲染](#186-xmljsonyaml和protobuf-渲染)
  - [1.9. 支持Let's Encrypt证书](#19-支持lets-encrypt证书)
  - [1.10. 测试](#110-测试)
  - [1.11. 参考资料](#111-参考资料)

## 1.1. 简介

```golang
func main() {
    router := gin.Default()

    s := &http.Server{
        Addr:           ":8080",
        Handler:        router,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe()
}
func main() {
    r := gin.Default()
    // Listen and serve on 0.0.0.0:8080
    r.Run(":8080")
}
```

## 1.2. header

### 1.2.1. 设置cookie

```golang
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/cookie", func(c *gin.Context) {
        cookie, err := c.Cookie("gin_cookie")
        if err != nil {
            cookie = "NotSet"
            c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
        }

        fmt.Printf("Cookie value: %s \n", cookie)
    })

    router.Run()
}
```

## 1.3. 路由规则

### 1.3.1. 简单路由

```golang
// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
router.GET("/user/:name", func(c *gin.Context){})

// 这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
// 如果没有其他路由器匹配/user/john，它将重定向到/user/john/
router.GET("/user/:name/*action", func(c *gin.Context) {})
```

### 1.3.2. 路由分组

```golang
func main() {
    router := gin.Default()

    // Simple group: v1
    v1 := router.Group("/v1")
    {
        v1.POST("/login", loginEndpoint)
        v1.POST("/submit", submitEndpoint)
        v1.POST("/read", readEndpoint)
    }
    // Simple group: v2
    v2 := router.Group("/v2")
    {
        v2.POST("/login", loginEndpoint)
        v2.POST("/submit", submitEndpoint)
        v2.POST("/read", readEndpoint)
    }

    router.Run(":8080")
}
```

### 1.3.3. 静态文件路径

```golang
func main() {
    router := gin.Default()
    router.Static("/assets", "./assets")
    router.StaticFS("/more_static", http.Dir("my_file_system"))
    router.StaticFile("/favicon.ico", "./resources/favicon.ico")

    // Listen and serve on 0.0.0.0:8080
    router.Run(":8080")
}
```

### 1.3.4. 重定向

```golang
r.GET("/test", func(c *gin.Context) {
    c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
})
r.GET("/test", func(c *gin.Context) {
    c.Request.URL.Path = "/test2"
    r.HandleContext(c)
})
r.GET("/test2", func(c *gin.Context) {
    c.JSON(200, gin.H{"hello": "world"})
})
```

## 1.4. 获取query和form表单数据

```golang
// 匹配的url格式:  /welcome?firstname=Jane&lastname=Doe
router.GET("/welcome", func(c *gin.Context) {
    firstname := c.DefaultQuery("firstname", "Guest")
    lastname := c.Query("lastname") // 是 c.Request.URL.Query().Get("lastname") 的简写
    c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
})

// POST /post?id=1234&page=1 HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
// name=manu&message=this_is_great
router.POST("/post", func(c *gin.Context) {
    id := c.Query("id")
    page := c.DefaultQuery("page", "0")
    name := c.PostForm("name")
    message := c.PostForm("message")

    fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
})
```

## 1.5. 接收文件

```golang
// 给表单限制上传大小 (默认 32 MiB)
// router.MaxMultipartMemory = 8 << 20  // 8 MiB
router.POST("/upload", func(c *gin.Context) {
    // 单文件
    file, _ := c.FormFile("file")
    log.Println(file.Filename)

    // 上传文件到指定的路径
    // c.SaveUploadedFile(file, dst)
    c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
})

//多文件上传
// curl -X POST http://localhost:8080/upload \
//   -F "upload[]=@/Users/appleboy/test1.zip" \
//   -F "upload[]=@/Users/appleboy/test2.zip" \
//   -H "Content-Type: multipart/form-data"
router.POST("/upload", func(c *gin.Context) {
    // 多文件
    form, _ := c.MultipartForm()
    files := form.File["upload[]"]

    for _, file := range files {
        log.Println(file.Filename)

        // 上传文件到指定的路径
        // c.SaveUploadedFile(file, dst)
    }
    c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
})
```

## 1.6. 中间件

```golang
// 默认启动方式，包含 Logger、Recovery 中间件
r := gin.Default()

//不包含中间件
r := gin.New()
// 全局中间件
// 使用 Logger 中间件
r.Use(gin.Logger())

// 使用 Recovery 中间件
r.Use(gin.Recovery())

// 路由添加中间件，可以添加任意多个
r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

// 路由组中添加中间件
// authorized := r.Group("/", AuthRequired())
```

### 1.6.1. 自定义中间件

```golang
func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now()
        // Set example variable
        c.Set("example", "12345")
        // before request
        c.Next()
        // after request
        latency := time.Since(t)
        log.Print(latency)

        // access the status we are sending
        status := c.Writer.Status()
        log.Println(status)
    }
}

func main() {
    r := gin.New()
    r.Use(Logger())

    r.GET("/test", func(c *gin.Context) {
        example := c.MustGet("example").(string)
        // it would print: "12345"
        log.Println(example)
    })
    // Listen and serve on 0.0.0.0:8080
    r.Run(":8080")
}
```

### 1.6.2. 中间件中使用Goroutines

```golang
func main() {
    r := gin.Default()

    r.GET("/long_async", func(c *gin.Context) {
        // 创建要在goroutine中使用的副本
        cCp := c.Copy()
        go func() {
            // simulate a long task with time.Sleep(). 5 seconds
            time.Sleep(5 * time.Second)
            // 这里使用你创建的副本
            log.Println("Done! in path " + cCp.Request.URL.Path)
        }()
    })

    r.GET("/long_sync", func(c *gin.Context) {
        // simulate a long task with time.Sleep(). 5 seconds
        time.Sleep(5 * time.Second)

        // 这里没有使用goroutine，所以不用使用副本
        log.Println("Done! in path " + c.Request.URL.Path)
    })
    // Listen and serve on 0.0.0.0:8080
    r.Run(":8080")
}
```

## 1.7. 日志

### 1.7.1. 中间件日志格式

```golang
func main() {
    // 禁用控制台颜色
    gin.DisableConsoleColor()

    // 创建记录日志的文件
    f, _ := os.Create("gin.log")
    gin.DefaultWriter = io.MultiWriter(f)

    // 如果需要将日志同时写入文件和控制台，请使用以下代码
    // gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

    router := gin.New()
    // LoggerWithFormatter 中间件会将日志写入 gin.DefaultWriter
    // By default gin.DefaultWriter = os.Stdout
    router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

        // 你的自定义格式
        return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
                param.ClientIP,
                param.TimeStamp.Format(time.RFC1123),
                param.Method,
                param.Path,
                param.Request.Proto,
                param.StatusCode,
                param.Latency,
                param.Request.UserAgent(),
                param.ErrorMessage,
        )
    }))
    router.Use(gin.Recovery())

    router.Run(":8080")
}
```

### 1.7.2. 自定义路由日志

```golang
import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
        log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
    }

    r.POST("/foo", func(c *gin.Context) {
        c.JSON(http.StatusOK, "foo")
    })

    r.GET("/bar", func(c *gin.Context) {
        c.JSON(http.StatusOK, "bar")
    })

    r.GET("/status", func(c *gin.Context) {
        c.JSON(http.StatusOK, "ok")
    })

    // Listen and Server in http://0.0.0.0:8080
    r.Run()
}
```

## 1.8. 模型绑定和验证

```golang
// 绑定为json
type Login struct {
    User     string `form:"user" json:"user" xml:"user"  binding:"required"`
    Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
    router := gin.Default()

    // Example for binding JSON ({"user": "manu", "password": "123"})
    router.POST("/loginJSON", func(c *gin.Context) {
        var json Login
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if json.User != "manu" || json.Password != "123" {
            c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
    })

    // Example for binding XML (
    //  <?xml version="1.0" encoding="UTF-8"?>
    //  <root>
    //      <user>user</user>
    //      <password>123</password>
    //  </root>)
    router.POST("/loginXML", func(c *gin.Context) {
        var xml Login
        if err := c.ShouldBindXML(&xml); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if xml.User != "manu" || xml.Password != "123" {
            c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
    })

    // Example for binding a HTML form (user=manu&password=123)
    router.POST("/loginForm", func(c *gin.Context) {
        var form Login
        // This will infer what binder to use depending on the content-type header.
        if err := c.ShouldBind(&form); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if form.User != "manu" || form.Password != "123" {
            c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
    })

    // Listen and serve on 0.0.0.0:8080
    router.Run(":8080")
}
```

### 1.8.1. 自定义验证器

### 1.8.2. 绑定Get参数或者Post参数

### 1.8.3. 绑定uri

### 1.8.4. 绑定HTML复选框

### 1.8.5. 绑定Post参数

### 1.8.6. XML、JSON、YAML和ProtoBuf 渲染

## 1.9. 支持Let's Encrypt证书

```golang
package main

import (
    "log"

    "github.com/gin-gonic/autotls"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/acme/autocert"
)

func main() {
    r := gin.Default()

    // Ping handler
    r.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

    log.Fatal(autotls.Run(r, "example1.com", "example2.com"))
}

func main() {
    r := gin.Default()

    // Ping handler
    r.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

    m := autocert.Manager{
        Prompt:     autocert.AcceptTOS,
        HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
        Cache:      autocert.DirCache("/var/www/.cache"),
    }

    log.Fatal(autotls.RunWithManager(r, &m))
}
```

## 1.10. 测试

```golang
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/ping", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    assert.Equal(t, "pong", w.Body.String())
}
```

## 1.11. 参考资料

1. [官网](https://github.com/gin-gonic/gin)
2. [中文doc参考](https://github.com/asong2020/Golang_Dream/blob/master/Gin/Doc)
3. [Gin框架中文文档](https://www.jianshu.com/p/98965b3ff638)
