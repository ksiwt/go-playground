package main

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

const expiration = 10*time.Second

type redisClient struct {
	client redis.Client
}

func NewRedisClient (client *redis.Client) *redisClient {
	return &redisClient{
		client: *client,
	}
}

func (r *redisClient) Set(ctx context.Context, key, value string) error {
	if err := r.client.Set(ctx, key, value, expiration).Err(); err != nil {
		return err
	}
	return nil
}

func (r *redisClient) Get(ctx context.Context, key string) (string, error) {
	res, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", nil
	}

	return res, nil
}