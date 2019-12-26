package library

import (
	"errors"
	"strings"
	"time"

	"gw/conf"

	"github.com/valyala/fasthttp"
)

type HttpRequest struct {
	Method string
	To     string
	Out    int
}

func (h *HttpRequest) Http() (string, error) {
	var body string
	var err error
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
