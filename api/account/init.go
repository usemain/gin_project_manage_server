package account

import (
	"gin_project_manage_server/shares/middleware"
	"github.com/gin-gonic/gin"
)

func InitAccount(api *gin.RouterGroup) {
	accountGroup := api.Group("/account", middleware.ValidateAuthorization())
	{
		accountGroup.GET("login", Login)
	}
}
