package data

import (
	"UserService/internal/conf"
	"context"
	"errors"
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
	db, err := NewDB(c)
	if err != nil || db == nil {
		return nil, nil, err
	}
	re, err := NewRedisClient(r)
	if err != nil || re == nil {
		sqlDB, _ := db.DB()
		sqlDB.Close() // 防止资源泄漏
		return nil, nil, err
	}
	cleanup := func() {
		helper := log.NewHelper(logger)
		helper.Info("closing the data resources")

		// 关闭数据库连接
		if sqlDB, err := db.DB(); err == nil {
			sqlDB.Close()
		}
		// 关闭 Redis 连接
		if re != nil {
			re.Close()
		}
	}
	return &Data{
		db: db,
		re: re,
	}, cleanup, nil
}

func NewDB(c *conf.Data) (*gorm.DB, error) {
	if c == nil || c.Database == nil || c.Database.Dsn == "" {
		log.Fatal("invalid database configuration")
		return nil, errors.New("invalid database configuration")
	}
	db, err := gorm.Open(mysql.Open(c.Database.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("invalid database configuration")
		return nil, err
	}
	if err := db.AutoMigrate(
		&User{},
	); err != nil {
		log.Fatalf("failed to migrate models: %v", err)
		return nil, err
	}
	return db, nil
}

func NewRedisClient(c *conf.Redis) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Database.Addr,
		Password: c.Database.Password,
		DB:       int(c.Database.Db),
	})

	// 测试连接
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
		return nil, err
	}

	return rdb, nil
}
