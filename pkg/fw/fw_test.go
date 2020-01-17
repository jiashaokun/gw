package fw

import (
	"fmt"
	"testing"

	"gw/library"
	"gw/util"
)

func TestFlow_Check(t *testing.T) {
	path := "http://guozi.com"
	fw := Flow{
		Path: path,
		Num:  10,
	}

	key := util.CacheKey(fmt.Sprintf("flow_%s", path, util.GetDay()))
	n := library.GetCache(key)

	err := fw.Check()

	if n == "10" && err != nil {
		t.Fatalf("Flow Num Was error limit 10 now %s err was nil", n)
	}
}
