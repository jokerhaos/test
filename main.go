package main

import (
	"test/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由
	r := gin.Default()
	routers.InitRouter(r)
	// 2.监听端口，默认在8080
	r.Run(":8000")
}
