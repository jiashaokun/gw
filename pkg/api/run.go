package api

import (
	"encoding/json"

	"gw/backend"
	"gw/conf"
	"gw/library"
	"gw/util"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type G struct {
	//请求时间
	RequestTime string
	//返回时间
	ResponseTime string
	//结果channel
	Rch chan string
	//error channel
	Ech chan error
	//MongoData
	Md backend.MongoInfo
	//com request path
	Path string
	//request url path
	To string
	//Query all url
	Query string
}

func (g *G) SetInfo(c *gin.Context) error {
	g.Path = util.GetRequestUri(c)

	//get cache
	key := util.CacheKey(g.Path)
	minfo := new(backend.MongoInfo)

	if str := library.GetCache(key); str != "" {
		json.Unmarshal([]byte(str), minfo)
		g.Md = *minfo
		return nil
	}

	if err := library.FindOne("wg", bson.M{"path": g.Path}, minfo); err != nil {
		return err
	}
	g.Md = *minfo
	//set cache
	info, err := json.Marshal(g.Md)
	if err != nil {
		return err
	}

	library.SetCache(key, string(info), conf.CacheMongoAllTime)

	return nil
}
