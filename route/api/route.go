package api

import (
	"gw/pkg/api"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	r.Any("/user/*action", api.Run)
}
