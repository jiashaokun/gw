package api

import (
	"fmt"

	"gw/library"
	"gw/pkg/api"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {

	list, err := library.FindAll("wg")
	if err != nil {
		panic(fmt.Sprintf("Api Route Was Wrong Err Was %s", err))
	}

	//动态加载路由,根据mongoDB中的path加载
	for _, v := range list {
		pth := v.Path
		r.Any(pth, api.Run)
	}

	//r.Any("/api/*action", api.Run)
}
