package limit_alg

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisTokenBucket struct {
	redisClient *redis.Client
	bucketSize  int
	refillRate  int
	interval    time.Duration
	key         string
}

// NewRedisTokenBucket creates a new Redis-backed token bucket.
func NewRedisTokenBucket(redisClient *redis.Client, key string, bucketSize, refillRate int, interval time.Duration) *RedisTokenBucket {
	return &RedisTokenBucket{
		redisClient: redisClient,
		bucketSize:  bucketSize,
		refillRate:  refillRate,
		interval:    interval,
		key:         key,
	}
}

// Allow checks if a token is available for the request.
func (t *RedisTokenBucket) Allow(ctx context.Context) (bool, error) {
	// Lua script for atomic token decrement and bucket refill in Redis
	luaScript := `
		local tokens = redis.call("GET", KEYS[1])
		if tokens == false then
			tokens = ARGV[1]
			redis.call("SET", KEYS[1], tokens)
			redis.call("EXPIRE", KEYS[1], ARGV[2])
		end
		if tonumber(tokens) > 0 then
			redis.call("DECR", KEYS[1])
			return 1
		else
			return 0
		end
	`
	bucketTTL := int(t.interval.Seconds()) // Expire time in seconds

	// Execute the Lua script
	result, err := t.redisClient.Eval(ctx, luaScript, []string{t.key}, t.bucketSize, bucketTTL).Int()
	if err != nil {
		return false, err
	}

	return result == 1, nil
}

// refill refills tokens into the Redis bucket at the specified rate.
func (t *RedisTokenBucket) refill(ctx context.Context) {
	ticker := time.NewTicker(t.interval)
	for range ticker.C {
		// Refill tokens up to the bucketSize, but no more than the allowed size
		luaScript := `
			local tokens = redis.call("GET", KEYS[1])
			if tokens == false then
				tokens = 0
			end
			local newTokens = math.min(ARGV[1] + tonumber(tokens), ARGV[2])
			redis.call("SET", KEYS[1], newTokens)
		`

		_, err := t.redisClient.Eval(ctx, luaScript, []string{t.key}, t.refillRate, t.bucketSize).Result()
		if err != nil {
			fmt.Printf("Error refilling tokens: %v\n", err)
		}
	}
}
