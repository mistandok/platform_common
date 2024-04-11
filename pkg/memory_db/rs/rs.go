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

// NewRs новый интерфейс для работы с БД.
func NewRs(pool *redis.Pool) memory_db.DB {
	return &rs{
		pool: pool,
	}
}

// DoContext выполнение команды с учетом контекста.
func (r *rs) DoContext(ctx context.Context, commandName string, args ...interface{}) (reply interface{}, err error) {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении соединения к Redis из пула соединений: %w", err)
	}
	defer func() {
		_ = conn.Close()
	}()

	reply, err = conn.Do(commandName, args...)
	if err != nil {
		return nil, fmt.Errorf("ошибка при попытке выполнения команды в Redis: %w", err)
	}

	return reply, nil
}

// Close закрытие соединения с БД.
func (r *rs) Close() error {
	return r.pool.Close()
}

// String конвертация ответа от бд к стрингу.
func (r *rs) String(reply interface{}, err error) (string, error) {
	return redis.String(reply, err)
}
