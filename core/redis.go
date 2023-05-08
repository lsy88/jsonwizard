package core

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/lsy88/jsonwizard/global"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.JW_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.JW_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.JW_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.JW_REDIS = client
	}
}
