package fw

import (
	"errors"
	"fmt"
	"strconv"

	"gw/util"
	"gw/library"
)

type Flow struct {
	//请求path对应的redis中的访问url
	Path string
	//限制访问次数
	Num int
}

func (f *Flow) Check() error {
	if f.Num == 0 {
		return nil
	}

	key := util.CacheKey(fmt.Sprintf("flow_%s", f.Path))

	//cache 中 path 的请求数 +1
	library.Incr(key)

	//获取 cache 中的请求数
	n := library.GetCache(key)
	num, _ := strconv.Atoi(n)

	if num > f.Num {
		return errors.New("Exceeding the demand limit")
	}

	return nil
}
