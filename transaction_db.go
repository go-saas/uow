package uow

import (
	"context"
	"database/sql"
)

type TransactionalDb interface {
	// Begin a transaction
	Begin(ctx context.Context, opt ...*sql.TxOptions) (db Txn, err error)
}

type Txn interface {
	Commit() error
	Rollback() error
}

type DbFactory func(ctx context.Context, key string) TransactionalDb