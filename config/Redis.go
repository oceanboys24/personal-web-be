package config

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)



var (
	ctx         = context.Background()
	RedisClient *redis.Client
)

func ConnectRedis() {
	dsn := os.Getenv("REDIS_URL")
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		log.Fatalf("Fail to parse Redis URL: %v", err)
	}
	RedisClient = redis.NewClient(opt)

	// Test Redis
	if err := RedisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("Can't Connect to Redis: %v", err)
	}
}