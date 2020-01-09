package admin

import (
	"gw/backend"
	"gw/library"
	"gw/response"
	"gw/util"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// 新增接口，增加配置接口
func Add(c *gin.Context) {
	var res, info backend.MongoInfo

	info.CreateTime = util.GetTime()
	info.UpdateTime = util.GetTime()

	if err := c.ShouldBind(&info); err != nil {
		response.Response(c, 504, nil)
		return
	}

	//检查是否group存在
	var group backend.MongoGroup
	library.FindOne("group", bson.M{"id": info.GroupId}, &group)

	if group.Id == "" {
		response.Response(c, 501, nil)
		return
	}

	//检查数据库是否存在该接口
	library.FindOne("wg", bson.M{"path": info.Path}, &res)
	if res.To != "" {
		response.Response(c, 555, nil)
		return
	}

	//add
	info.Id = uuid.NewV4().String()
	if err := library.Add("wg", info); err != nil {
		response.Response(c, 500, nil)
		return
	}

	response.Response(c, 200, nil)
}
