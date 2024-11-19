package cmd

import (
	"gin_project_manage_server/api/account"
	"gin_project_manage_server/api/task"
	"gin_project_manage_server/api/user"
	"github.com/gin-gonic/gin"
)

func Run() {
	app := gin.Default()

	api := app.Group("api")
	account.InitAccount(api)
	user.InitUser(api)
	task.InitTask(api)

	if err := app.Run(":8888"); err != nil {
		panic("Service startup failed -> " + err.Error())
	}
}
