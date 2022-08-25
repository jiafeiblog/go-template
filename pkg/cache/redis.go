package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisClient 全局 redis 客户端
var RedisClient *redis.Client

// InitRedis 初始化 redis 客户端
func InitRedis(conn, username, password string) {
	client := redis.NewClient(&redis.Options{
		Addr:     conn,
		Username: username,
		Password: password,
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}

// GetCacheBytes 从 redis 中获取 Bytes 类型的缓存
func GetCacheBytes(key string) (resp []byte, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	resp, err = RedisClient.Get(ctx, key).Bytes()
	return
}

// SetCacheJson 将 interface 类型的值放入 redis 缓存中
func SetCacheJson(key string, value interface{}, cacheTime time.Duration) (err error) {
	bytes, _ := json.Marshal(value)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	err = RedisClient.Set(ctx, key, bytes, cacheTime).Err()
	return err
}

// GetCacheString 从 redis 中获取 string 类型的缓存
func GetCacheString(key string) (val string, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	resp := RedisClient.Get(ctx, key)
	val, err = resp.Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

// SetCacheString 将 string 类型的值放入 redis 缓存中
func SetCacheString(key, value string, cacheTime time.Duration) (err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	err = RedisClient.Set(ctx, key, value, cacheTime).Err()
	return err
}

// DelCacheString 从 redis 中删除 string 类型的缓存
func DelCacheString(key string) (err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	err = RedisClient.Del(ctx, key).Err()
	return err
}

// SetCacheSet 添加集合
func SetCacheSet(key, value string) (err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	err = RedisClient.SAdd(ctx, key, value).Err()
	return err
}

// GetCacheSetSize 获取集合大小
func GetCacheSetSize(key string) (size int64, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	card, err := RedisClient.SCard(ctx, key).Result()
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return card, nil
}
