package routers

import (
	"net/http"
	"test/src/controller"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter(engine *gin.Engine) *gin.Engine {
	// api 接口路由配置

	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	g := engine.Group("/api")
	{
		g.GET("test", controller.Test)
	}
	return engine
}
