package main

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// PgxPoolConnIface mock pgxpool.Conn
type PgxPoolConnIface interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
	Conn() *pgx.Conn
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Ping(ctx context.Context) error
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Release()
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
}

var (
	// return your actual connection (*pgxpool.Conn)
	PgxPoolConn = func() PgxPoolConnIface {
		pool, _ := pgxpool.Connect(context.Background(), "connect string")
		conn, _ := pool.Acquire(context.Background())
		return conn
	}
)

func PgxPoolConnHandler() error {
	conn := PgxPoolConn()
	rows, err := conn.Query(context.Background(), `select * from foo;`)
	if err != nil {
		return err
	}
	defer func() {
		rows.Close()
		conn.Release()
	}()

	var result int
	for rows.Next() {
		if err = rows.Scan(&result); err != nil {
			return err
		}
	}
	return nil
}
