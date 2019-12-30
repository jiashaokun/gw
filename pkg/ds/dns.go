package ds

import (
	"fmt"
	"strconv"
	"strings"

	"gw/library"
	"gw/util"

	"github.com/gin-gonic/gin"
)

type Dns struct {
	Ds    int
	Pth   string
	To    string
	Query string
	Ctx   *gin.Context
}

func (d *Dns) GetRestUrl() {
	//不需要dns
	if d.Ds != 1 {
		d.To = getUrl(d.Pth, 0)
		d.Query = util.GetRequestUrl(d.To, d.Ctx)
		return
	}

	//需要dns
	key := util.CacheKey(fmt.Sprintf("%s_request_dns_num", d.Pth))
	//获取当前访问次数
	incr := library.GetCache(key)
	num, _ := strconv.Atoi(incr)

	l := getList(d.Pth)
	//计算应该访问第几个
	uk := num % len(l)

	d.To = getUrl(d.Pth, uk)

	d.Query = util.GetRequestUrl(d.To, d.Ctx)

	//cache ++
	library.Incr(key)
}

//url to []string
func getList(u string) []string {
	str := strings.TrimRight(u, ",")
	list := strings.Split(str, ",")

	return list
}

//get url
func getUrl(u string, k int) string {
	ls := getList(u)

	return ls[k]
}
