package memory_db

import "context"

//go:generate ../../bin/mockery --output ./mocks  --inpackage-suffix --all --case snake

// Client интерфейс клиента
type Client interface {
	DB() DB
	Close() error
}

// DB интерфейс для работы с БД
type DB interface {
	QueryExecutor
	ReplyConverter
	Close() error
}

// QueryExecutor интерфейс для выполнения запросов
type QueryExecutor interface {
	DoContext(ctx context.Context, commandName string, args ...interface{}) (reply interface{}, err error)
}

// ReplyConverter интерфейс для конвертирования ответов
type ReplyConverter interface {
	String(reply interface{}, err error) (string, error)
}
