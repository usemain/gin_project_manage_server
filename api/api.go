package api

import (
	"gin_project_manage_server/api/account"
	"gin_project_manage_server/api/task"
	"gin_project_manage_server/api/user"
	"gin_project_manage_server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitAccountRouter(apiGroup *gin.RouterGroup) {
	accountGroup := apiGroup.Group("/account", middleware.ValidateAuthorization())
	{
		accountGroup.GET("login", account.Login)
	}
}

func InitUserRouter(apiGroup *gin.RouterGroup) {
	userGroup := apiGroup.Group("/user")
	{
		userGroup.GET("userinfo", user.Userinfo)
	}
}

func InitTaskRouter(apiGroup *gin.RouterGroup) {
	taskGroup := apiGroup.Group("/task", middleware.ValidateAuthorization())
	{
		taskGroup.GET("list", task.List)
	}
}
