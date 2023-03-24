package infrastructures

import (
	"app/config"
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func NewRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",config.Get().Redis.Host,config.Get().Redis.Port),
		Username: config.Get().Redis.Username,
		Password: config.Get().Redis.Password,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("failed to ping redis: %s\n", err.Error())
	}

	return client
}
