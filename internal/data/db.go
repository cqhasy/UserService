package data

import (
	"UserService/internal/conf"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Data struct {
	db *gorm.DB
	re *redis.Client
}

func NewData(c *conf.Data, r *conf.Redis, logger log.Logger) (*Data, func(), error) {
	db := NewDB(c)
	re := NewRedisClient(r)
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db: db,
		re: re,
	}, cleanup, nil
}

func NewDB(c *conf.Data) *gorm.DB { //注意：需要多返回一个err
	if c == nil || c.Database == nil || c.Database.Dsn == "" {
		panic("invalid database configuration")
	}
	db, err := gorm.Open(mysql.Open(c.Database.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(
		&User{},
	); err != nil {
		panic(err)
	}
	return db
}

func NewRedisClient(c *conf.Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Database.Addr,
		Password: c.Database.Password,
		DB:       int(c.Database.Db),
	})

	// 测试连接
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		panic("failed to connect to Redis: " + err.Error())
	}

	return rdb
}
