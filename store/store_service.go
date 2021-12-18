package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// Define the struct wrapper around raw Redis client
type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
    ctx = context.Background()
)

// In real life, we should use some rdbms with lru policy instead using redis only, redis should be used only for high frequently used urls
const CacheDuration = 6 * time.Hour

// Initializing the store service and return a store pointer 
func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

/* Save mapping between the longUrl and the generated shortUrl url */

func SaveUrlMapping(shortUrl string, longUrl string, userId string){ 
	err := storeService.redisClient.Set(ctx, shortUrl, longUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - longUrl: %s\n", err, shortUrl, longUrl))
	}
}

func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}