package rs

import (
	"github.com/gomodule/redigo/redis"
	"github.com/mistandok/platform_common/pkg/memory_db"
)

var _ memory_db.Client = (*redisClient)(nil)

type redisClient struct {
	masterDB memory_db.DB
}

// New новый клиент для работы с редисом.
func New(pool *redis.Pool) memory_db.Client {
	return &redisClient{masterDB: NewRs(pool)}
}

// DB интерфейс для общений с БД.
func (c *redisClient) DB() memory_db.DB {
	return c.masterDB
}

// Закрытия соединения с БД.
func (c *redisClient) Close() error {
	if c.masterDB != nil {
		return c.masterDB.Close()
	}

	return nil
}
