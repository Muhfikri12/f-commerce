package database

import (
	"context"
	"finance/config"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	rdb     *redis.Client
	expired time.Duration
	prefix  string
}

func newRedisClient(url, password string, dbIndex int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       dbIndex,
	})
}

func NewCache(cfg *config.Config, expired int) *Cache {
	return &Cache{
		rdb:     newRedisClient(cfg.Redis.Url, cfg.Redis.Password, 0),
		expired: time.Duration(expired) * time.Second,
		prefix:  cfg.Redis.Prefix,
	}
}

func (c *Cache) SaveToken(name string, value string) error {
	return c.rdb.Set(context.Background(), c.prefix+"_"+name, value, 24*time.Hour).Err()
}

func (c *Cache) Delete(name string) error {
	return c.rdb.Del(context.Background(), c.prefix+"_"+name).Err()
}
