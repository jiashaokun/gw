package util

import (
	"crypto/md5"
	"encoding/hex"
)

// 获取key
func CacheKey(k string) string {
	ctx := md5.New()
	ctx.Write([]byte(k))

	key := hex.EncodeToString(ctx.Sum(nil))

	return key
}
