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

func Response(c *gin.Context, code int, v interface{}, msg string) {
	out := Resp{
		Code: code,
		Msg:  conf.Code[code],
	}

	if msg != "" {
		out.Msg = msg
	}

	out.Data = v
	c.JSON(code, out)
}