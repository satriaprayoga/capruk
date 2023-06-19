package capruk

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

func setupRedis() *redis.Client {
	now := time.Now()
	conString := fmt.Sprintf("%s:%d", Config.Redis.Host, Config.Redis.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     conString,
		Password: Config.Redis.Password,
		DB:       Config.Redis.DB,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	timeSpent := time.Since(now)
	log.Printf("Config redis is ready in %v", timeSpent)
	return rdb
}
