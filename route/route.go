package route

import (
	"gw/route/admin"

	"github.com/gin-gonic/gin"
)

// 调用设置路由
func Route() *gin.Engine {
	r := gin.Default()
	//外部调用路由
	//api.Route(r)

	//后台接口调用路由
	admin.Route(r)
	return r
}
