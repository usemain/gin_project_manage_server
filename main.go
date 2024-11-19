package main

import (
	"gin_project_manage_server/cmd"
	"gin_project_manage_server/shares/dao"
)

func init() {
	dao.InitMySQL()
	dao.InitRedis()
}

func main() {
	cmd.Run()
}
