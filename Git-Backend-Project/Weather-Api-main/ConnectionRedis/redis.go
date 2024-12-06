package connectionredis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func RedisConnection(addr string, ctx context.Context) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	log.Printf("Redis is running on: %s\n", addr)

	return rdb
}
