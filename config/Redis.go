package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)



var (
	ctx         = context.Background()
	RedisClient *redis.Client
)

func ConnectRedis() {
	opt, err := redis.ParseURL("rediss://default:AYcdAAIjcDEwNmQ2MDFjMjQ4NTE0YzU4OGZjNTA3NmI5ZWI5NDM0YnAxMA@literate-badger-34589.upstash.io:6379")
	if err != nil {
		log.Fatalf("Fail to parse Redis URL: %v", err)
	}
	RedisClient = redis.NewClient(opt)

	// Test Redis
	if err := RedisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("Can't Connect to Redis: %v", err)
	}
}