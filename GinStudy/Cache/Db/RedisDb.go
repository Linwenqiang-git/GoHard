package db

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"

	glocConfig "github.linwenqiang.com/GinStudy/AppSetting"
	tool "github.linwenqiang.com/GinStudy/Tool"
)

type RedisStore struct {
	redisClient *redis.Client
	ctx         context.Context
}

//设置缓存
func (cs *RedisStore) Set(key string, value string) {
	err := cs.redisClient.Set(cs.ctx, key, value, time.Minute*10).Err()
	if err != nil {
		log.Println(err.Error())
	}
}

//获取缓存
func (cs *RedisStore) Get(key string, clear bool) string {
	val, err := cs.redisClient.Get(cs.ctx, key).Result()
	tool.PanicError(err)
	if clear {
		err := cs.redisClient.Del(cs.ctx, key).Err()
		tool.PanicError(err)
	}
	return val
}

var Redis *redis.Client

func InitRediStore() *RedisStore {
	config := glocConfig.GetConfig().RedisConfig

	Redis = redis.NewClient(&redis.Options{
		Addr:     config.Addr + ":" + string(config.Port),
		Password: "",
		DB:       config.Db,
	})
	var ctx = context.Background()
	customeStore := &RedisStore{
		redisClient: Redis,
		ctx:         ctx}
	//base64Captcha.SetCustomStore(customeStore)

	return customeStore
}
