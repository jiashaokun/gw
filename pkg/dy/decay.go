// 容错
package dy

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

type Decay struct {
	//是否开启容错：0：不开启，1：开启
	Open int
	//发送请求的url /user/info?a=b&c=d
	Url string
	//容错递减间隔时间, 秒
	DecayTime int
	//最后一次请求成功的返回值
	Body string
	//框架结构体,最后一次返回的结果将set到上下文中，用于middleware中更新缓存
	Ctx *gin.Context
}

func (d *Decay) Start() string {
	if d.Open == 0 {
		return ""
	}

	//获取请求的url /user/info?a=1&b=2
	path := util.GetRequestUri(d.Ctx)
	d.Url = util.GetRequestUrl(path, d.Ctx)
	//从cache中获取数据
	body, st := getBody(d.Url, d.DecayTime)
	if st == true {
		return ""
	}

	d.Body = body

	return body
}

//根据时间判断返回body
func getBody(u string, dtm int) (string, bool) {
	key := util.CacheKey(fmt.Sprintf("decay_body_key_%s", u))

	//获取上一次请求的状态，如果是正常则返回空
	if status := library.HGet(key, "status"); status == "" || status == "1" {
		return "", true
	}

	prevTime := library.HGet(key, "request_time")
	if prevTime == "" {
		return "", true
	}

	//计算时间，如果时间在设定时间范围内则直接返回
	to, err := time.ParseInLocation("2006-01-02 15:04:05", prevTime, time.Local)
	if err != nil {
		return "", true
	}
	//上一次请求的时间
	toUx := to.Unix()

	//获取当前时间
	nextTime := util.GetTime()
	ts, err := time.ParseInLocation("2006-01-02 15:04:05", nextTime, time.Local)
	if err != nil {
		return "", true
	}
	tsUx := ts.Unix()
	//间隔时间超过设定的秒数，再次发起请求
	if (tsUx - toUx) > int64(dtm) {
		return "", true
	}

	body := library.HGet(key, "body")
	if body == "" || body == "null" {
		return "", true
	}

	return body, false
}
