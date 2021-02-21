package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func init() {}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	r := NewRedisClient(client)
	defer r.client.Close()

	ctx := context.Background()

	err := r.Set(ctx, "hoge", "huga")
	if err != nil {
		panic(err)
	}

	val, err := r.Get(ctx, "hoge")
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
