package repositories

import (
	"app/internal/core/ports"
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type cacheRepository struct {
	rc *redis.Client
}

func NewCacheRepository(rc *redis.Client) ports.CacheRepository {
	return &cacheRepository{rc: rc}
}

func (c *cacheRepository)Get(ctx context.Context,key string) (string, error) {
	return c.rc.Get(ctx,key).Result()
}

func (c *cacheRepository)Set(ctx context.Context, key string, val string,expire time.Duration)error{
	return c.rc.Set(ctx,key,val,expire).Err()
}

func (c *cacheRepository)Delete(ctx context.Context,key string)error{
	return c.rc.Del(ctx,key).Err();
}