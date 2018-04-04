package gost

import (
	"fmt"
	"time"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     50,
		PoolTimeout:  30 * time.Second,
		Password: "redis",
		DB: 0,
	})
	pong,err := RedisClient.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(pong,"redis connect SUCCESS")
	}
}


func testRedis() {
	err := RedisClient.Set("key", "value", 0).Err()
	if err == nil {
		fmt.Println("set ok")
	}
}
