package cache

import (
	"regexp"
	"fmt"
	"context"
	"time"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/harisaginting/tech-test-kredivo/pkg/log"	
	"github.com/harisaginting/tech-test-kredivo/pkg/utils/helper"
)

var (
	ctx = context.Background()
	rdb *redis.Client
	disabled bool
)

func init(){
	cfgredis 	 := helper.ForceInt(helper.MustGetEnv("REDIS"))
	if cfgredis == 0 {
		disabled = false
	}else{
		NewRedisClient()
	}
}

func CreateCacheKey(value string) (cacheKey string) {
	if disabled == true { return }
	prefix := helper.GetEnvOrDefault("APP_NAME", "DEVELOPMENT") + ":" + helper.GetEnvOrDefault("MODE", "LOCAL") + ":"
	re := regexp.MustCompile("=|&")
	cacheKey = prefix + re.ReplaceAllString(value, "_")
	return cacheKey
}

// NewRedisClient
func NewRedisClient() (err error) {
	if disabled == true { return }
	redisHost := helper.MustGetEnv("REDIS_HOST")
	redisPort := helper.GetEnvOrDefault("REDIS_PORT", "6379")
	if redisHost == "" || redisPort == ""{
		log.Fatal(ctx, nil, "Redis Configuration Error")
	}


	redisAddr 		:= redisHost+":"+redisPort
	redisPassword   := helper.GetEnvOrDefault("REDIS_PASSWORD","eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81")
	dbNumber, err 	:= strconv.Atoi(helper.GetEnvOrDefault("REDIS_DB", "0"))
	if err != nil {
		log.Warn(ctx, fmt.Sprintf("Failed to convert string to int : %s ", err))
	}

	// redis client
	rdb = redis.NewClient(&redis.Options{
		Addr:       redisAddr,
		Password:   redisPassword,
		DB:         dbNumber,
		PoolSize:   1000,
		MaxRetries: 2,
	})
	ping, err := rdb.Ping(ctx).Result()
	if err == nil && len(ping) > 0 {
		log.Warn(ctx, fmt.Sprintf("Connected to Redis: %s", redisAddr))
	} else {
		log.Fatal(ctx, err, "Redis Error Connection:")
	}
	return
}

// SetKey set key value redis
func SetKey(key string, value interface{}) error {
	if disabled == true { return nil}
	log.Warn(ctx, fmt.Sprintf("Redis: Set key:", key))
	err := rdb.Set(ctx, key, value, 0).Err()
	if err == redis.Nil {
		log.Warn(ctx, fmt.Sprintf("Redis: Set key Nil: %s", key))
	}
	if err != nil {
		log.Warn(ctx, fmt.Sprintf("Redis: Error Set key: %s", err))
	}
	return nil
}

// GetKey get key of redis
func GetKey(key string) (string, error) {
	if disabled == true { return "",nil }
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		log.Warn(ctx, fmt.Sprintf("Redis: key does not exist %s", err))
		return "", err
	}
	if key != "" {
		log.Warn(ctx, fmt.Sprintf("Redis: getKey %s", key))
	}
	return val, err
}

func SetKeyWithExpired(key string, value interface{}, expiredStr string) error {
	if disabled == true { return nil }
	log.Warn(ctx, fmt.Sprintf("Redis: Set key:", key))
	// sample expiredStr "24h"
	// duration, _ := time.ParseDuration(24)
	duration, _ := time.ParseDuration(expiredStr)
	err := rdb.Set(ctx, key, value, duration).Err()

	if err == redis.Nil {
		log.Warn(ctx, fmt.Sprintf("Redis: Set key Nil: %s", key))
	}
	if err != nil {
		log.Warn(ctx, fmt.Sprintf("Redis: Error Set key: %s", err))
	}

	return nil
}

func GetTTL(key string) (td time.Duration) {
	if disabled == true { return }
	ttl := rdb.TTL(ctx, key)
	if ttl.Err() != nil {
		log.Warn(ctx, fmt.Sprintf("Redis: TTL Error %s : %s ", key, ttl.Err().Error()))
	}
	if ttl.Val() <= -1 {
		log.Warn(ctx, fmt.Sprintf("Redis: No have TTL: %s", key))
	}
	td = ttl.Val()
	return
}

func FlushDB() {
	if disabled == true { return }
	rdb.FlushDB(ctx)
}

func GetAllKeys() (v []string) {
	if disabled == true { return }
	v, _ = rdb.Keys(ctx, "*").Result()
	log.Warn(ctx, fmt.Sprintf("KEYS: %s", v))
	return 
}

func GetListKeys(key string) (v []string) {
	if disabled == true { return }
	v, _ = rdb.Keys(ctx, "*"+key+"*").Result()
	log.Warn(ctx, fmt.Sprintf("KEYS:", v))
	return
}

func DelKey(key string) (v int64) {
	if disabled == true { return }
	v, _ = rdb.Del(ctx, key).Result()
	log.Warn(ctx, fmt.Sprintf("KEYS: %s", v))
	return 
}

func ExpireAt(key string, expire time.Time) (ex bool) {
	if disabled == true { return }
	ex, _ = rdb.ExpireAt(ctx, key, expire).Result()
	log.Warn(ctx, fmt.Sprintf("KEYS:", ex))
	return
}
