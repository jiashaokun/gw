package library

import (
	"fmt"
	"testing"
)

func TestHttpRequest_Http(t *testing.T) {
	hr := HttpRequest{
		Method: "get",
		To:     "http://api.guazi.com/car/info",
		Out:    5,
	}
	body, err := hr.Http()
	if err != nil {
		t.Fatalf(fmt.Sprintf("Http Request Get Was Wrong Err Was %s", err))
	}
	if body == "" {
		t.Fatalf(fmt.Sprintln("Http Request Get Was Wrong Value Was nil"))
	}
}
