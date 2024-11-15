package cmd

import (
	"gin_project_manage_server/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run() {
	app := gin.Default()

	app.StaticFS("static", http.Dir("static"))

	apiGroup := app.Group("api")
	api.InitAccountRouter(apiGroup)
	api.InitUserRouter(apiGroup)
	api.InitTaskRouter(apiGroup)

	if err := app.Run(":8888"); err != nil {
		panic("Server run error -> " + err.Error())
	}
}
