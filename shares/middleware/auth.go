package middleware

import (
	"gin_project_manage_server/shares/global"
	"gin_project_manage_server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AbortResponse(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"message": "Authorization failed",
	})
}

func ValidateAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			AbortResponse(c)
			return
		}

		authorization := c.GetHeader("Authorization")[strings.LastIndex(c.GetHeader("Authorization"), "Bearer ")+7:]
		data, err := utils.ParseAuthorization(authorization)
		if err != nil {
			AbortResponse(c)
			return
		}

		if do := global.GvaRedis.Do(global.GvaCtx, "get", data.Email+"_token"); do.Val() != authorization {
			AbortResponse(c)
			return
		} else {
			c.Set("uid", data.Uid)
			c.Next()
		}
	}
}
