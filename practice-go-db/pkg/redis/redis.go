package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type redisClient struct {
	client *redis.Client
}

func NewRedisClient(host, port, password string, db, poolSize int) *redisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
		PoolSize: poolSize,
	})

	return &redisClient{client: client}
}

func (c redisClient) GetConnection() *redis.Client {
	return c.client
}
