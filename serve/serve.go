package serve

import (
	"github.com/gin-gonic/gin"
	"go-study/common/status"
)

func Ping(c *gin.Context) {
	response := status.Response{
		Code:    status.OK,
		Message: "成功",
		Data:    []any{1, 2, 3, 4},
	}
	c.JSON(200, response)
}
