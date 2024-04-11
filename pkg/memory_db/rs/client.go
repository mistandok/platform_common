package rs

import (
	"github.com/gomodule/redigo/redis"
	"github.com/mistandok/platform_common/pkg/memory_db"
)

var _ memory_db.Client = (*redisClient)(nil)

type redisClient struct {
	masterDB memory_db.DB
}

func New(pool *redis.Pool) memory_db.Client {
	return &redisClient{masterDB: NewRs(pool)}
}

func (c *redisClient) DB() memory_db.DB {
	return c.masterDB
}

func (c *redisClient) Close() error {
	if c.masterDB != nil {
		return c.masterDB.Close()
	}

	return nil
}
