// @Author moruikang
// @Date 2025/3/22 03:58:00
// @Desc Redis操作相关接口

package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"macro-mall/internal/admin/config"
	"macro-mall/internal/pkg/options"
	"sync"
	"time"
)

var (
	Factory   RedisFactory
	redisOnce sync.Once
)

type RedisFactory interface {

	// 查找所有以指定前缀开头的键
	GetAllPrefixKey(priefix string) ([]string, error)
	// 删除所有以指定前缀开头的键
	DeleteAllPrefixKey(priefix string) error
	// 获取指定键的值
	GetKey(key string) (string, error)
	// 设置指定键的值
	SetKey(key string, value interface{}) error

	HGetAllKey(key string) (map[string]string, error)

	// 以下是Redis 原生用法

	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration ...time.Duration) *redis.StatusCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Exists(ctx context.Context, keys ...string) *redis.IntCmd
	Keys(ctx context.Context, pattern string) *redis.StringSliceCmd
	MGet(ctx context.Context, keys ...string) *redis.SliceCmd
	// 设置多个键值对: MSet(ctx,"name", "Alice", "age", 30)
	MSet(ctx context.Context, values ...interface{}) *redis.StatusCmd
	HGet(ctx context.Context, key string, field string) *redis.StringCmd
	HSet(ctx context.Context, key string, field string, value interface{}) *redis.IntCmd
	HGetAll(ctx context.Context, key string) *redis.MapStringStringCmd
	HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd
	HMSet(ctx context.Context, key string, values ...interface{}) *redis.BoolCmd
	HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd
}

type Redis struct {
	Ctx    context.Context
	Client *redis.Client
}

var _ RedisFactory = (*Redis)(nil)

func NewRedisFactory(ctx context.Context, client *redis.Client) RedisFactory {
	return &Redis{
		Ctx:    ctx,
		Client: client,
	}
}

func SetRedisFactory(f RedisFactory) {

	if Factory != nil {
		Factory = f
	}
}

func GetRedisFactoryOr(opts *options.RedisOptions) (RedisFactory, error) {

	if opts == nil && Factory == nil {
		opts = &config.GlobalConfig.Redis
	}

	var err error
	var redisIns *redis.Client

	redisOnce.Do(func() {
		redisIns = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", opts.Host, opts.Port),
			Password: opts.Password,
			DB:       opts.Database,
		})
		ctx := context.Background()
		err = redisIns.Info(ctx).Err()
		if err != nil {
			log.Errorf("Redis Connect Fail, ERR = %s", err.Error())
		}
		Factory = NewRedisFactory(ctx, redisIns)
	})

	if Factory == nil || err != nil {
		log.Errorf("fali to get redis store factory %s", err.Error())
	}
	return Factory, nil
}

func (r Redis) GetAllPrefixKey(prefix string) ([]string, error) {
	/*    keys, err := r.Client.Keys(r.Ctx, prefix+"*").Result()
	      if err != nil {
	          return nil, err
	      }
	      return keys, nil*/
	// 避免 keys * 命令，优化为scan命令
	var keys []string
	var cursor uint64
	var n int
	for {
		var err error
		var batchKeys []string

		batchKeys, cursor, err = r.Client.Scan(r.Ctx, cursor, prefix+"*", 100).Result()
		if err != nil {
			return nil, err
		}
		n += len(keys)
		keys = append(keys, batchKeys...)
		if cursor == 0 {
			break
		}
	}
	return keys, nil

}

func (r Redis) DeleteAllPrefixKey(prefix string) error {
	keys, err := r.GetAllPrefixKey(prefix)
	if err != nil {
		return err
	}
	if len(keys) == 0 {
		return nil
	}
	return r.Client.Del(r.Ctx, keys...).Err()
}

func (r Redis) GetKey(key string) (string, error) {
	val, err := r.Client.Get(r.Ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func (r Redis) HGetAllKey(key string) (map[string]string, error) {
	return r.Client.HGetAll(r.Ctx, key).Result()
}

func (r Redis) SetKey(key string, value interface{}) error {
	return r.Client.Set(r.Ctx, key, value, 0).Err()
}

func (r Redis) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.Client.Get(ctx, key)
}

func (r Redis) Set(ctx context.Context, key string, value interface{}, expiration ...time.Duration) *redis.StatusCmd {
	var exp time.Duration
	if len(expiration) > 0 {
		exp = expiration[0]
	}
	return r.Client.Set(ctx, key, value, exp)
}

func (r Redis) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	return r.Client.Del(ctx, keys...)
}

func (r Redis) Exists(ctx context.Context, keys ...string) *redis.IntCmd {
	return r.Client.Exists(ctx, keys...)
}

func (r Redis) Keys(ctx context.Context, pattern string) *redis.StringSliceCmd {
	return r.Client.Keys(ctx, pattern)
}

func (r Redis) MGet(ctx context.Context, keys ...string) *redis.SliceCmd {
	return r.Client.MGet(ctx, keys...)
}

func (r Redis) MSet(ctx context.Context, values ...interface{}) *redis.StatusCmd {
	return r.Client.MSet(ctx, values...)
}

func (r Redis) HGet(ctx context.Context, key string, field string) *redis.StringCmd {
	return r.Client.HGet(ctx, key, field)
}

func (r Redis) HSet(ctx context.Context, key string, field string, value interface{}) *redis.IntCmd {
	return r.Client.HSet(ctx, key, field, value)
}

func (r Redis) HGetAll(ctx context.Context, key string) *redis.MapStringStringCmd {
	return r.Client.HGetAll(ctx, key)
}

func (r Redis) HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd {
	return r.Client.HMGet(ctx, key, fields...)
}

func (r Redis) HMSet(ctx context.Context, key string, values ...interface{}) *redis.BoolCmd {
	return r.Client.HMSet(ctx, key, values...)
}

func (r Redis) HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd {
	return r.Client.HDel(ctx, key, fields...)
}
