package core

import (
	"context"
	"fmt"
	"github.com/lsy88/jsonwizard/global"
	"github.com/lsy88/jsonwizard/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

func Gorm() {
	switch global.JW_CONFIG.System.DbType {
	case "mysql":
		global.JW_DB.DB = GormMysql()
	case "mongodb":
		global.JW_DB.Mongo = InitMongoDB()
	default:
		global.JW_DB.DB = GormMysql()
	}
	return
}

//初始化mysql连接
func GormMysql() *gorm.DB {
	m := global.JW_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(),
		DefaultStringSize:         191,   //string类型字段的默认长度
		SkipInitializeWithVersion: false, //根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		fmt.Println("gorm.Open failed：", err)
		return nil
	} else {
		sqlDB, _ := db.DB()
		//fmt.Println("mysql connect success!")
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

//初始化mongodb连接
func InitMongoDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("newclient mongodb error: ", err)
		return nil
	}
	err = client.Connect(context.Background())
	if err != nil {
		fmt.Println("connect mongodb error: ", err)
		return nil
	}
	//连接到MongoDB服务
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("connect mongodb error: ", err)
		return nil
	}
	return client
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.Application{},
	)
	if err != nil {
		global.JW_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.JW_LOG.Info("register table success")
}
