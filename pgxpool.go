package main

import (
	"context"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	// return your actual pool (*pgxpool.Pool)
	PgxPool = func() pgxpoolmock.PgxPool {
		pool, _ := pgxpool.Connect(context.Background(), "connect string")
		return pool
	}
)

func PgxPoolHandler() error {
	rows, err := PgxPool().Query(context.Background(), `select * from foo;`)
	if err != nil {
		return err
	}
	defer rows.Close()

	var result int
	for rows.Next() {
		if err = rows.Scan(&result); err != nil {
			return err
		}
	}
	return nil
}
