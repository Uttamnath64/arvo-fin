package storage

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
	ctx    context.Context
}

func NewRedisClient(ctx context.Context, addr, password string, db int) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Ping to ensure Redis connection is successful
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	return &RedisClient{Client: rdb, ctx: ctx}
}

// SetValue stores a key-value pair with TTL
func (r *RedisClient) SetValue(key, value string, ttl int) error {
	return r.Client.Set(r.ctx, key, value, time.Duration(ttl)*time.Second).Err()
}

// GetValue retrieves the value for a given key
func (r *RedisClient) GetValue(key string) (string, error) {
	return r.Client.Get(r.ctx, key).Result()
}

// DeleteKey removes a key from Redis
func (r *RedisClient) DeleteKey(key string) error {
	return r.Client.Del(r.ctx, key).Err()
}
