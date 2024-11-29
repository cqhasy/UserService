package data

import (
	"UserService/internal/conf"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Mock configuration for database connection
func getMysqlConfig() *conf.Data {
	return &conf.Data{
		Database: &conf.Data_Database{
			Dsn: "root:ccnu@tcp(127.0.0.1:3307)/users?charset=utf8mb4&parseTime=True&loc=Local",
		},
	}
}

func getRedisConfig() *conf.Redis {
	return &conf.Redis{
		Database: &conf.Redis_Database{
			Addr:     "127.0.0.1:6379",
			Password: "ccnu",
			Db:       0,
		},
	}
}

// Test NewDB function
func TestNewDB(t *testing.T) {
	c := getMysqlConfig()

	db := NewDB(c)
	assert.NotNil(t, db, "Database connection should not be nil")

	sqlDB, err := db.DB()
	assert.NoError(t, err, "Failed to get SQL DB")
	defer sqlDB.Close()

	err = sqlDB.Ping()
	assert.NoError(t, err, "Failed to ping database")
}

func TestNewRedisClient(t *testing.T) {
	c := getRedisConfig()
	client := NewRedisClient(c)

	// 验证 Redis 客户端是否非 nil
	assert.NotNil(t, client, "Redis client should not be nil")

	// 测试 Redis 客户端是否能成功连接
	// 使用 Ping 命令来测试连接
	_, err := client.Ping(context.Background()).Result()

	// 验证连接是否成功
	assert.NoError(t, err, "Failed to connect to Redis")

	// 关闭 Redis 客户端连接
	defer client.Close()
}
