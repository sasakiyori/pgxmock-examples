package main

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

type PgxConnIface interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) (err error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) (err error)
	Close(ctx context.Context) error
	ConnInfo() *pgtype.ConnInfo
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	Deallocate(ctx context.Context, name string) error
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	PgConn() *pgconn.PgConn
	Ping(ctx context.Context) error
	Prepare(ctx context.Context, name string, sql string) (sd *pgconn.StatementDescription, err error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
}

var (
	// return your actual connection (*pgx.Conn)
	PgxConn = func() PgxConnIface {
		conn, _ := pgx.Connect(context.Background(), "connect string")
		return conn
	}
)

func PgxConnHandler() error {
	rows, err := PgxConn().Query(context.Background(), `select * from foo;`)
	if err != nil {
		return err
	}
	defer func() {
		rows.Close()
		PgxConn().Close(context.Background())
	}()

	var result int
	for rows.Next() {
		if err = rows.Scan(&result); err != nil {
			return err
		}
	}
	return nil
}
