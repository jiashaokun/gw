package middle

import (
	"fmt"
	"time"

	"gw/library"
	"gw/util"

	"github.com/gin-gonic/gin"
)

/*
Redis struct
status : 0,1 (0:请求失败 1:请求成功)
request_time : time.Now().Format("2006-01-02 15:04:05")
body : request_body
*/

//获取最后一次写入的body，存入cache
func Body() gin.HandlerFunc {
	return func(context *gin.Context) {
		//获取请求的url /user/info?a=1&b=2
		path := util.GetRequestUri(context)
		url := util.GetRequestUrl(path, context)

		//body有正确的值，重写写入cache
		key := util.CacheKey(fmt.Sprintf("decay_body_key_%s", url))
		//先设置最后一次请求的时间
		library.HSet(key, "request_time", time.Now().Format("2006-01-02 15:04:05"))

		body, err := context.Get("RequestBody")
		if body == nil || err != true {
			library.HSet(key, "status", 0)
			return
		}
		//判断body是否是null字符串
		if body == "null" {
			library.HSet(key, "status", 0)
			return
		}

		//设置 rds 中的结构数据
		library.HSet(key, "body", fmt.Sprintf("%s", body))
		library.HSet(key, "status", 1)
	}
}
