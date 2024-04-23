package main

import (
	"flag"
	"github.com/gin-gonic/gin"
)

var staticRoot string //静态文件根目录
var port string       //端口

func main() {
	flag.StringVar(&staticRoot, "s", "", "静态文件根目录")
	flag.StringVar(&port, "p", "", "端口")
	flag.Parse()
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Static("/static", staticRoot)
	r.Run(port)
}
