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

// 获取wg列表
func ListWg(c *gin.Context) {
	var info backend.MongoWgListApi

	var det, res backend.MongoInfo
	var detArray []*backend.MongoInfo
	var resp []backend.MongoInfo

	if err := c.ShouldBind(&info); err != nil {
		response.Response(c, 504, nil)
		return
	}

	w := bson.M{}
	if info.Name != "" {
		w["name"] = info.Name
	}

	if info.Id != "" {
		w["id"] = info.Id
	}

	if info.GroupId != "" {
		w["group_id"] = info.GroupId
	}

	detArray, err := library.FindAllWg("wg", w, &det)
	if err != nil {
		response.Response(c, 555, nil)
		return
	}

	if len(detArray) == 0 {
		response.Response(c, 200, resp)
		return
	}

	for _, v := range detArray {
		res.Id = v.Id
		res.To = v.To
		res.Dns = v.Dns
		res.Path = v.Path
		res.Name = v.Name
		res.Flow = v.Flow
		res.Decay = v.Decay
		res.Method = v.Method
		res.GroupId = v.GroupId
		res.Timeout = v.Timeout
		res.DecayTime = v.DecayTime
		res.CacheTime = v.CacheTime
		res.CreateTime = v.CreateTime
		res.UpdateTime = v.UpdateTime

		resp = append(resp, res)
	}
	response.Response(c, 200, resp)
}
