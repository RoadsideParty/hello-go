package router

import (
	"go-study/serve"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", serve.Ping)
	return r
}
