package controller

import (
	"net/http"
	"test/util"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Data interface{}
}

func Test(c *gin.Context) {
	data := Data{Data: 123}
	c.JSON(http.StatusOK, util.NewApiJsonResult("200", "success").Simple(data))
}
