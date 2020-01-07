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

	//检查数据库是否存在该接口
	library.FindOne("wg", bson.M{"path": info.Path}, &res)
	if res.To != "" {
		response.Response(c, 500, nil)
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

//新增接口路由组
func AddGroup(c *gin.Context) {
	var res, info backend.MongoGroup

	info.CreateTime = util.GetTime()
	info.UpdateTime = util.GetTime()

	if err := c.ShouldBind(&info); err != nil {
		response.Response(c, 504, nil)
		return
	}

	//检查数据库是否存在该接口
	library.FindOne("group", bson.M{"group": info.Group}, &res)

	if res.Id != "" {
		response.Response(c, 500, nil)
		return
	}

	//add
	info.Id = uuid.NewV4().String()
	if err := library.Add("group", info); err != nil {
		response.Response(c, 500, nil)
		return
	}

	response.Response(c, 200, nil)
}

//签名分配秘钥接口
func AddAuth(c *gin.Context) {
	var res, info backend.MongoAuth

	info.CreateTime = util.GetTime()
	info.UpdateTime = util.GetTime()

	if err := c.ShouldBind(&info); err != nil {
		response.Response(c, 504, nil)
		return
	}

	//检查数据库是否存在该接口
	library.FindOne("auth", bson.M{"src": info.Src}, &res)

	if res.Id != "" {
		response.Response(c, 500, nil)
		return
	}

	//add
	info.Id = uuid.NewV4().String()
	info.Key = util.CacheKey(info.Id)

	if err := library.Add("auth", info); err != nil {
		response.Response(c, 500, nil)
		return
	}

	response.Response(c, 200, nil)
}
