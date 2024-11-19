package global

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	GvaDB    *gorm.DB
	GvaRedis *redis.Client
	GvaCtx   = context.Background()
)
