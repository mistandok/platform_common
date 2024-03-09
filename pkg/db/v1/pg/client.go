package pg

import (
	"context"
	"fmt"
	"github.com/mistandok/platform_common/pkg/db/v1"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type pgClient struct {
	masterDBC v1.DB
	logger    *zerolog.Logger
}

// New новый клиент для работы с Postgres
func New(ctx context.Context, dsn string, logger *zerolog.Logger) (v1.Client, error) {
	pgxConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка при формировании конфига для pgxpool: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("ошибка при подключении к БД: %v", err)
	}

	return &pgClient{
		masterDBC: NewDB(pool, logger),
		logger:    logger,
	}, nil
}

// DB доступ к интерфейсу базы данных
func (c *pgClient) DB() v1.DB {
	return c.masterDBC
}

// Close закрытие соединений
func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
