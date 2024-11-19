package user

import (
	"github.com/gin-gonic/gin"
)

func InitUser(api *gin.RouterGroup) {
	userGroup := api.Group("/user")
	{
		userGroup.GET("userinfo", Userinfo)
	}
}
