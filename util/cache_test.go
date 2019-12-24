package util

import "testing"

func TestCacheKey(t *testing.T) {
	md := CacheKey("test")
	want := "098f6bcd4621d373cade4e832627b4f6"
	if md != "098f6bcd4621d373cade4e832627b4f6" {
		t.Fatalf("Util Md5 Func Was Wrong Wang %s Now %s", want, md)
	}
}
