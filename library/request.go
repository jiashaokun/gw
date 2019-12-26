package library

import (
	"errors"
	"fmt"
	"gw/util"
	"strings"
	"time"

	"gw/conf"

	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
)

type HttpRequest struct {
	Method    string
	To        string
	Out       int
	CacheTime int
}

func (h *HttpRequest) Http() (string, error) {
	var body string
	var err error

	//get cache
	key := util.CacheKey(h.To)
	body = GetCache(key)
	if body != "" {
		return body, nil
	}

	method := strings.ToUpper(h.Method)

	switch method {
	case "GET":
		body, err = get(h.To, h.Out)
		break
	case "POST":
		body, err = post(h.To, h.Out)
		break
	default:
		body, err = "", errors.New("Http Request Any Method")
	}

	SetCache(key, body, h.CacheTime)

	return body, err
}

func get(u string, out int) (string, error) {
	timeout := time.Duration(out) * time.Second

	cli := fasthttp.Client{
		MaxConnsPerHost: conf.ReqMaxConnsPerHost, //最大链接数
		ReadTimeout:     timeout,
	}

	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()

	req.Header.SetContentType("application/json")
	req.Header.SetMethod("GET")
	req.SetRequestURI(u)

	if err := cli.DoTimeout(req, res, timeout); err != nil {
		return "", err
	}

	body := res.Body()
	bodyStr := string(body)

	return bodyStr, nil
}

func post(u string, out int) (string, error) {
	timeout := time.Duration(out) * time.Second

	cli := fasthttp.Client{
		MaxConnsPerHost: conf.ReqMaxConnsPerHost, //最大链接数
		ReadTimeout:     timeout,                 //主动断开时间
	}

	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()

	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.SetRequestURI(u)

	if err := cli.DoTimeout(req, res, timeout); err != nil {
		return "", err
	}

	body := res.Body()
	bodyStr := string(body)

	return bodyStr, nil
}

func (h *HttpRequest) ParserUrl(c *gin.Context) {
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
		//todo add other
		break
	}

	if query != "" {
		h.To = fmt.Sprintf("%s?%s", h.To, query)
	}
}
