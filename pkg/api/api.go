package api

import (
	"fmt"
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
		}
		dns.GetRestUrl()
		glb.To = dns.To

		//发起请求
		hp := library.HttpRequest{
			Method:    glb.Md.Method,
			To:        glb.To,
			Out:       glb.Md.Timeout,
			CacheTime: glb.Md.CacheTime,
		}
		//设置正确的url
		hp.ParserUrl(c)

		//发起请求
		body, err := hp.Http()
		if err != nil {
			glb.Ech <- err
			return
		}
		glb.Rch <- body
		//容错 todo
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
