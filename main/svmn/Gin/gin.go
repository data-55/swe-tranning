package Gin

import (
	"forbizbe/main/page/home"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/post/home", home.HomeUserList)

	return r
}
