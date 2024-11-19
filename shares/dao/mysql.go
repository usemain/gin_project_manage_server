package dao

import (
	"fmt"
	"gin_project_manage_server/model"
	"gin_project_manage_server/shares/global"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const (
	formatDns = "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local"
)

func InitMySQL() {
	var ds = &model.DATABASES{}
	filePath := "shares/config/databases.yaml"
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic("Failed to read file -> " + err.Error())
	}
	if err = yaml.Unmarshal(file, ds); err != nil {
		panic("Yaml unmarshal failed -> " + err.Error())
	}

	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Info,
		Colorful:      true,
	})

	DSN := fmt.Sprintf(
		formatDns,
		ds.MySQL.User,
		ds.MySQL.Password,
		ds.MySQL.Host,
		ds.MySQL.Port,
		ds.MySQL.Database,
		ds.MySQL.Charset,
	)
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("MySQL connect failed -> " + err.Error())
	}

	global.GvaDB = db
}
