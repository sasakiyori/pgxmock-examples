package main

import (
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v3"
)

// embed pgxmock.PgxConnIface
type PgxPoolConnMock struct {
	pgxmock.PgxConnIface
}

// add impl by your need
func (c *PgxPoolConnMock) Hijack() *pgx.Conn {
	return nil
}

func (c *PgxPoolConnMock) Release() {

}

func TestPgxPoolConnHandler(t *testing.T) {
	mock, err := pgxmock.NewConn(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
	if err != nil {
		t.Error(err)
	}

	PgxPoolConn = func() PgxPoolConnIface {
		return &PgxPoolConnMock{mock}
	}

	row := pgxmock.NewRows([]string{"result"}).AddRow(1)
	mock.ExpectQuery("select").WillReturnRows(row)

	if err = PgxPoolConnHandler(); err != nil {
		t.Error(err)
	}
}
