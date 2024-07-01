package database

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Host     string
	Port     int
	DB       int
	Password string
}

func NewRedisClient(ctx context.Context, cfg RedisConfig) *redis.Client {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if pong := client.Ping(ctx); pong.String() != "ping: PONG" {
		// TODO: logging and panic
	}

	return client
}
