package task

import (
	"gin_project_manage_server/shares/middleware"
	"github.com/gin-gonic/gin"
)

func InitTask(api *gin.RouterGroup) {
	taskGroup := api.Group("/task", middleware.ValidateAuthorization())
	{
		taskGroup.GET("list", List)
	}
}
