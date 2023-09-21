package main

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxPoolIface interface {
	Acquire(ctx context.Context) (*pgxpool.Conn, error)
	AcquireAllIdle(ctx context.Context) []*pgxpool.Conn
	AcquireFunc(ctx context.Context, f func(*pgxpool.Conn) error) error
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	Close()
	Config() *pgxpool.Config
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Ping(ctx context.Context) error
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Reset()
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	Stat() *pgxpool.Stat
}

var (
	// return your actual connection (*pgxpool.Conn)
	PgxPool = func() PgxPoolIface {
		pool, _ := pgxpool.New(context.Background(), "connect string")
		return pool
	}
)

func PgxPoolHandler() error {
	pool := PgxPool()

	rows, err := pool.Query(context.Background(), `select * from foo;`)
	if err != nil {
		return err
	}
	defer func() {
		rows.Close()
		pool.Close()
	}()

	var result int
	for rows.Next() {
		if err = rows.Scan(&result); err != nil {
			return err
		}
	}
	return nil
}
