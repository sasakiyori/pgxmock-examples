package main

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type PgxConnIface interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	Close(ctx context.Context) error
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	Deallocate(ctx context.Context, name string) error
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	PgConn() *pgconn.PgConn
	Ping(ctx context.Context) error
	Prepare(ctx context.Context, name string, sql string) (sd *pgconn.StatementDescription, err error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	// functions pgxmock doesn't implement, add them by your need
	Config() *pgx.ConnConfig
	IsClosed() bool
	DeallocateAll(ctx context.Context) error
	LoadType(ctx context.Context, typeName string) (*pgtype.Type, error)
	TypeMap() *pgtype.Map
	WaitForNotification(ctx context.Context) (*pgconn.Notification, error)
}

var (
	// return your actual connection (*pgx.Conn)
	PgxConn = func() PgxConnIface {
		conn, _ := pgx.Connect(context.Background(), "connect string")
		return conn
	}
)

func PgxConnHandler() error {
	conn := PgxConn()
	rows, err := conn.Query(context.Background(), `select * from foo;`)
	if err != nil {
		return err
	}
	defer func() {
		rows.Close()
		conn.Close(context.Background())
	}()

	var result int
	for rows.Next() {
		if err = rows.Scan(&result); err != nil {
			return err
		}
	}
	return nil
}
