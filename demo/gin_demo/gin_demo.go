package main

// 安装库
// go get github.com/gin-gonic/gin
// go get github.com/gin-contrib/sse
// go get gopkg.in/yaml.v2

import "github.com/gin-gonic/gin"

func main() {
    // Default 返回一个默认的路由引擎
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        // 输出json结果给调用方
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.GET("/test", testHandle)

    //r.Run() // listen and serve on 0.0.0.0:8080
    r.Run(":9090")
}

func testHandle(c *gin.Context) {
   c.JSON(200, gin.H{
       "message": "test",
   })
}