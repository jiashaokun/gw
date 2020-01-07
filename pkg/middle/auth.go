package middle

import (
	"fmt"
	"net/url"
	"strings"

	"gw/backend"
	"gw/conf"
	"gw/library"
	"gw/response"
	"gw/util"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// 统一鉴权 从 header 中获取 Content-Src 和 Content-Md5 并结合参数和秘钥进行排序签名计算,具体描述可参照 README.md 中签名文案
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if conf.AccessToken != 1 {
			return
		}
		date := util.GetDay()
		src := c.GetHeader("Content-Src")
		sign := c.GetHeader("Content-Md5")

		//获取src配置对应的秘钥
		var info backend.MongoAuth

		library.FindOne("auth", bson.M{"src": src}, &info)

		//获取请求方式
		method := strings.ToUpper(c.Request.Method)

		switch method {
		case "GET":
			query := c.Request.URL.RawQuery
			params, _ := url.ParseQuery(query)
			query = params.Encode()

			signKey := fmt.Sprintf("%s%s%s", query, info.Key, date)
			auth := util.CacheKey(signKey)

			if sign != auth {
				response.Response(c, 503, nil)
				c.Abort()
			}
			return
		case "POST":
			c.Request.ParseForm()
			query := c.Request.PostForm.Encode()

			signKey := fmt.Sprintf("%s%s%s", query, info.Key, date)
			auth := util.CacheKey(signKey)

			if sign != auth {
				response.Response(c, 503, nil)
				c.Abort()
			}
			return
		}
		return
	}
}
