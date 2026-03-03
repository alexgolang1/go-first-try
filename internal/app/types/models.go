package types

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	ID      int
	Name    string
	Surname string
}

type RedisClient struct {
	client *redis.Client
}
