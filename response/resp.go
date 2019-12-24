package response

import (
	"gw/conf"

	"github.com/gin-gonic/gin"
)

//通用返回
type Resp struct {
	Code int         `json:"code, string"`
	Msg  string      `json:"msg, string"`
	Data interface{} `json:"data"`
}

func Response(c *gin.Context, code int, v interface{}) {
	d := [...]int{}
	out := Resp{
		Code: code,
		Msg:  conf.Code[code],
		Data: d,
	}

	if v != nil {
		out.Data = v
	}
	c.JSON(code, out)
}
