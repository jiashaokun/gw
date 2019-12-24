package admin

import (
	"fmt"

	"gw/response"
	"gw/backend/admin"

	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	var info admin.Add

	if err := c.ShouldBind(&info); err != nil {
		fmt.Println(err)
		response.Response(c, 504, [...]int{}, "参数验证失败")
	}

	fmt.Println(info)
}
