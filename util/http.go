package util

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
)

//获取访问过来的uri
func GetRequestUri(c *gin.Context) string {
	c.Request.ParseForm()
	u, _ := url.Parse(c.Request.RequestURI)

	return u.Path
}

//获取最终请求的请求串儿
func GetRequestUrl(to string, c *gin.Context) string {
	query, method := "", c.Request.Method
	switch method {
	case "GET":
		query = c.Request.URL.RawQuery
		break
	case "POST":
		c.Request.ParseForm()
		param := c.Request.PostForm
		if len(param) > 0 {
			query = param.Encode()
		}
		break
	default:
		//todo any other
		break
	}

	queryStr := to
	if query != "" {
		queryStr = fmt.Sprintf("%s?%s", to, query)
	}
	return queryStr
}
