package api

import (
	"fmt"
	flow2 "gw/pkg/flow"
	"net/http"
	"time"

	"gw/library"
	ds "gw/pkg/dns"
	"gw/util"

	"github.com/gin-gonic/gin"
)

// 入口函数
func Run(c *gin.Context) {
	//设置 global
	var glb G
	glb.Rch = make(chan string)
	glb.Ech = make(chan error)

	go func(c *gin.Context, glb *G) {
		glb.RequestTime = util.GetTime()

		//设置请求访问的数据
		if err := glb.SetInfo(c); err != nil {
			glb.Ech <- err
			return
		}

		//获取要访问的url
		dns := ds.Dns{
			Ds:  glb.Md.Dns,
			Pth: glb.Md.To,
			Ctx: c,
		}
		dns.GetRestUrl()
		glb.To = dns.To
		glb.Query = dns.Query
		
		//流量检查
		flow := flow2.Flow{
			Path: glb.To,
			Num:  glb.Md.Flow,
		}
		if err := flow.Check(); err != nil {
			glb.Ech <- err
			return
		}

		//发起请求
		hp := library.HttpRequest{
			Method:    glb.Md.Method,
			To:        glb.To,
			Query:     glb.Query,
			Out:       glb.Md.Timeout,
			CacheTime: glb.Md.CacheTime,
		}

		//发起请求
		body, err := hp.Http()
		if err != nil {
			glb.Ech <- err
			return
		}
		glb.Rch <- body
	}(c, &glb)

	select {
	case rch := <-glb.Rch:
		c.String(http.StatusOK, rch)
	case ech := <-glb.Ech:
		c.String(http.StatusInternalServerError, fmt.Sprintln(ech))
	case <-time.After(10 * time.Second):
		c.String(http.StatusNotFound, "request time out")
	}
}
