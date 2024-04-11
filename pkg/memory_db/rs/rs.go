package rs

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/mistandok/platform_common/pkg/memory_db"
)

var _ memory_db.DB = (*rs)(nil)

type rs struct {
	pool *redis.Pool
}

func NewRs(pool *redis.Pool) memory_db.DB {
	return &rs{
		pool: pool,
	}
}

func (r *rs) DoContext(ctx context.Context, commandName string, args ...interface{}) (reply interface{}, err error) {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении соединения к Redis из пула соединений: %w", err)
	}
	defer func() {
		_ = conn.Close()
	}()

	reply, err = redis.String(conn.Do(commandName, args...))
	if err != nil {
		return nil, fmt.Errorf("ошибка при попытке выполнения команды в Redis: %w", err)
	}

	return reply, nil
}

func (r *rs) Close() error {
	return r.pool.Close()
}

func (r *rs) String(reply interface{}, err error) (string, error) {
	return redis.String(reply, err)
}