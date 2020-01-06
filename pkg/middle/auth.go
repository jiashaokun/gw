package middle

import (
	"gw/conf"

	"github.com/gin-gonic/gin"
)


// 统一鉴权 从 header 中获取 Content-Src 和 Content-Md5 并结合参数和秘钥进行排序签名计算,具体描述可参照 README.md 中签名文案
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if conf.AccessToken != 1 {
			return
		}
		//src := c.GetHeader("Content-Src")
		//sign := c.GetHeader("Content-Md5")

		//获取 src 配置的签名key
	}
}