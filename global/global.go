package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/lsy88/jsonwizard/config"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

type jwDb struct {
	DB    *gorm.DB
	Mongo *mongo.Client
}

var (
	JW_DB                  jwDb
	JW_CONFIG              config.Server
	JW_VP                  *viper.Viper
	JW_LOG                 *zap.Logger
	JW_REDIS               *redis.Client
	JW_Concurrency_Control = &singleflight.Group{}
)

const (
	ConfigEnv         = "JW_CONFIG"
	ConfigDefaultFile = "config.yml"
)
