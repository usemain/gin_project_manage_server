package middleware

import (
	"gin_project_manage_server/internal/global"
	"gin_project_manage_server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ValidateAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Authorization failed",
			})
			return
		}

		authorization := c.GetHeader("Authorization")[strings.LastIndex(c.GetHeader("Authorization"), "Bearer ")+7:]
		data, err := utils.ParseAuthorization(authorization)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Authorization failed",
			})
			return
		}

		if do := global.GVA_REDIS.Do(global.GVA_CTX, "get", data.Email+"_token"); do.Val() != authorization {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Authorization failed",
			})
			return
		} else {
			c.Set("uid", data.Uid)
			c.Next()
		}
	}
}
