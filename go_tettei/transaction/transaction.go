package transaction

import (
	"context"
	"database/sql"
	"fmt"
)

type Transaction interface {
	DoInTx(context.Context, func(ctx context.Context) (any, error)) (any, error)
}

type tx struct {
	db *sql.DB
}

func NewTransaction(db *sql.DB) Transaction {
	return &tx{db}
}
func (t *tx) DoInTx(ctx context.Context, f func(ctx context.Context) (any, error)) (any, error) {
	tx, err := t.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		fmt.Println("begin error")
		return nil, err
	}

	ctx = context.WithValue(ctx, "tx", tx)
	v, err := f(ctx)
	if err != nil {
		fmt.Println("f error")
		_ = tx.Rollback()
		return nil, err
	}

	fmt.Println("commit.")
	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
	}
	return v, nil
}
