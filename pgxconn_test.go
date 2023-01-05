package main

import (
	"context"
	"testing"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgconn/stmtcache"
	"github.com/jackc/pgx/v4"
	"github.com/pashagolub/pgxmock"
)

// embed pgxmock.PgxConnIface
type PgxConnMock struct {
	pgxmock.PgxConnIface
}

// add impl by your need
func (c *PgxConnMock) Config() *pgx.ConnConfig {
	return nil
}
func (c *PgxConnMock) IsClosed() bool {
	return false
}
func (c *PgxConnMock) StatementCache() stmtcache.Cache {
	return nil
}
func (c *PgxConnMock) WaitForNotification(ctx context.Context) (*pgconn.Notification, error) {
	return nil, nil
}

func TestPgxConnHandler(t *testing.T) {
	mock, err := pgxmock.NewConn(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
	if err != nil {
		t.Error(err)
	}
	PgxConn = func() PgxConnIface {
		return &PgxConnMock{mock}
	}

	row := pgxmock.NewRows([]string{"result"}).AddRow(1)
	mock.ExpectQuery("select").WillReturnRows(row)

	if err = PgxConnHandler(); err != nil {
		t.Error(err)
	}
}
