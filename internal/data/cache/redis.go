//存入的数据一个小时内，可以在缓存里取

package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/WH-5/url-gin-gorm/configs"
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(config configs.RdConfig) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.DB,
	})
	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, fmt.Errorf("connect redis failed: %w", err)
	}
	return &RedisClient{client: client}, nil
}
func (c *RedisClient) close() error {
	return c.client.Close()
}

// SetURL 存入<code,url>
func (c *RedisClient) SetURL(shortcode, url string) error {
	result, err := c.client.Exists(context.Background(), shortcode).Result()
	if err != nil {
		return fmt.Errorf("redis set failed: %w", err)
	}
	if result == 1 {
		return fmt.Errorf("redis set failed: url %s already exists", url)
	}
	_, err = c.client.Set(context.Background(), shortcode, url, 0).Result()
	if err != nil {
		return fmt.Errorf("redis set failed: %w", err)
	}
	return nil
}

func (c *RedisClient) GetURL(shortcode string) (url string, err error) {
	result, err := c.client.Get(context.Background(), shortcode).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("redis get failed: %w", err)
	}
	return result, nil
}
