package main

import (
	"gin_project_manage_server/api/account"
	"gin_project_manage_server/api/task"
	"gin_project_manage_server/api/user"
	"gin_project_manage_server/shares/dao"
	"github.com/gin-gonic/gin"
)

func init() {
	dao.InitMySQL()
	dao.InitRedis()
}

func main() {
	app := gin.Default()

	api := app.Group("api")
	account.InitAccount(api)
	user.InitUser(api)
	task.InitTask(api)

	if err := app.Run(":8888"); err != nil {
		panic("Service startup failed -> " + err.Error())
	}
}
