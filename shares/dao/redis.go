package dao

import (
	"fmt"
	"gin_project_manage_server/model"
	"gin_project_manage_server/shares/global"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v3"
	"os"
)

func InitRedis() {
	var ds = &model.DATABASES{}
	filePath := "shares/config/databases.yaml"
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic("Failed to read file -> " + err.Error())
	}
	if err = yaml.Unmarshal(file, ds); err != nil {
		panic("Yaml unmarshal failed -> " + err.Error())
	}

	DSN := fmt.Sprintf("%s:%v", ds.REDIS.Host, ds.REDIS.Port)
	global.GvaRedis = redis.NewClient(&redis.Options{
		Addr: DSN,
	})
}
