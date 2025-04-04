package cache

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func NewRedisClient() *redis.Client {
    return redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        DB:   0,
    })
}

func Create(shortCode string, longURL string) {
    client := NewRedisClient()
    err := client.Set(ctx, shortCode, longURL, 24*time.Hour).Err()
    if err != nil {
        log.Fatal(err)
    }
}

func Get(shortCode string) (string, bool) {
    client := NewRedisClient()
    longURL, err := client.Get(ctx, shortCode).Result()
    if err == redis.Nil {
        return "", false
    }
    return longURL, true
}
func IncrementURLAccess(shortCode string) {
    client := NewRedisClient()
    client.ZIncrBy(ctx, "url_frequency", 1, shortCode) // Increment count by 1
}

func GetAccessCount(shortCode string) (int, error) {
	client := NewRedisClient()
	accessCount, err := client.ZScore(ctx, "url_frequency", shortCode).Result()
	if err == redis.Nil {
		return 0, nil // No access count found
	} else if err != nil {
		return 0, err // Return error if something goes wrong
	}
	return int(accessCount), nil
}

