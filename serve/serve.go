package serve

import (
	"github.com/gin-gonic/gin"
	"go-study/common/response"
	"go-study/common/token"
)

func Ping(c *gin.Context) {
	c.JSON(200, response.NewResponse(200, "成功", struct {
		Token string `json:"token"`
	}{
		Token: token.GetToken(),
	}))
}
