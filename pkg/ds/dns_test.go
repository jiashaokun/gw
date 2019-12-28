package ds

import (
	"fmt"
	"strconv"
	"testing"

	"gw/library"
	"gw/util"
)

func TestDns_GetRestUrl(t *testing.T) {
	u := "http://guazi.com,http://guazi2.com,http://guazi3.com"
	dns := Dns{Ds: 1, Pth: u}
	dns.GetRestUrl()

	//需要dns
	key := util.CacheKey(fmt.Sprintf("%s_getresurl", u))
	//获取当前访问次数
	incr := library.GetCache(key)
	num, _ := strconv.Atoi(incr)

	l := getList(u)
	//计算应该访问第几个(先访问被测试函数，里面对 num 进行了 ++ 操作，所以这边需要进行 --)
	uk := (num - 1) % len(l)

	to := getUrl(u, uk)

	if dns.To != to {
		t.Fatalf("Dns Get Url Was Wrong Want %s now %s", to, dns.To)
	}
}
