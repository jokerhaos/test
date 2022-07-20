package main

import (
	"fmt"
	"test/models"
	"test/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 加载env
	godotenv.Load(".env")
	// 创建路由
	r := gin.Default()
	routers.InitRouter(r)
	// 创建数据库链接
	fmt.Println("创建数据库连接")
	models.ConnectDB()

	// 2.监听端口，默认在8080
	r.Run(":8000")
}
