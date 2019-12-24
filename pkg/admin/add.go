package admin

import (
	"fmt"
	"gw/library"
	"gw/util"

	"gw/response"
	"gw/backend/admin"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// 新增接口，增加配置接口
func Add(c *gin.Context) {
	var res, info admin.Add

	info.CreateTime = util.GetTime()
	info.UpdateTime = util.GetTime()

	if err := c.ShouldBind(&info); err != nil {
		response.Response(c, 504, nil)
	}

	fmt.Println(info.Path)
	//检查数据库是否存在该接口
	library.FindOne("wg", bson.M{"path": info.Path}, &res)
	if res.To != "" {
		response.Response(c, 500, nil)
	}

	//add
	if err := library.Add("wg", info); err != nil {
		response.Response(c, 500, nil)
	}

	response.Response(c, 200, nil)
}
