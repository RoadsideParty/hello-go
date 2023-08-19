package router

import (
	"github.com/gin-gonic/gin"
	"go-study/serve"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", serve.Ping)
	return r
}
