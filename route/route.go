package route

import (
	"gw/pkg/middle"
	"gw/route/admin"
	"gw/route/api"

	"github.com/gin-gonic/gin"
)

// 调用设置路由
func Route() *gin.Engine {
	r := gin.Default()

	//统一鉴权
	r.Use(middle.Auth())

	//后台接口调用路由
	admin.Route(r)

	//外部调用路由
	api.Route(r)

	return r
}
