package dao

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"gin-web-template/internal/config"

	"github.com/go-redis/redis/v8"
)

var (
	redisInstance          *RedisCache
	onceRedisInitilization sync.Once
)

// RedisCache 封装 Redis 客户端
type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

// InitRedisDB 初始化 Redis 客户端
func InitRedisDB() error {
	var err error
	var cfg *config.Config
	onceRedisInitilization.Do(func() {
		if cfg, err = config.ServerConfig(); err != nil {
			log.Println("无法获取服务器配置: ", err)
		}
		addr := cfg.GetRedisAddr()
		password := cfg.GetRedisPassword()

		client, initErr := connectToRedis(addr, password, context.Background())
		if initErr != nil {
			err = initErr
			return
		}

		redisInstance = client
		log.Println("Redis连接成功")
	})

	return err
}

// connectToRedis 建立 Redis 连接
func connectToRedis(addr, password string, ctx context.Context) (*RedisCache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0, // 默认使用 0 号数据库
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisCache{
		client: rdb,
		ctx:    ctx,
	}, nil
}

// GetRedisInstance 获取 Redis 实例
func GetRedisInstance() (*RedisCache, error) {
	if redisInstance == nil {
		return nil, errors.New("Redis实例未初始化")
	}
	return redisInstance, nil
}

// Set 设置键值
func (r *RedisCache) Set(key string, value interface{}, expiration ...time.Duration) error {
	ttl := time.Duration(0)
	if len(expiration) > 0 {
		ttl = expiration[0]
	}
	return r.client.Set(r.ctx, key, value, ttl).Err()
}

// Get 获取键值
func (r *RedisCache) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

// Del 删除键
func (r *RedisCache) Del(key string) error {
	return r.client.Del(r.ctx, key).Err()
}

// Close 关闭 Redis 连接
func (r *RedisCache) Close() error {
	return r.client.Close()
}
