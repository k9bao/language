package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 匹配的url格式:  /someGet?firstname=Jane&lastname=Doe
func getMethod(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // 是 c.Request.URL.Query().Get("lastname") 的简写

	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

//Content-Type: application/x-www-form-urlencoded
func postMethod(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}

func defaultMethod(c *gin.Context) {
	c.String(http.StatusOK, "Hello %v %v %v", c.Request.RemoteAddr, c.Request.Method, c.Request.URL.Path)
}

func main_method() {
	// Disable Console Color
	// gin.DisableConsoleColor()

	// 使用默认中间件创建一个gin路由器
	// logger and recovery (crash-free) 中间件
	router := gin.Default()

	router.GET("/someGet", getMethod)
	router.POST("/somePost", postMethod)
	router.PUT("/somePut", defaultMethod)
	router.DELETE("/someDelete", defaultMethod)
	router.PATCH("/somePatch", defaultMethod)
	router.HEAD("/someHead", defaultMethod)
	router.OPTIONS("/someOptions", defaultMethod)

	// 默认启动的是 8080端口，也可以自己定义启动端口
	router.Run()
	// router.Run(":3000") for a hard coded port
}
