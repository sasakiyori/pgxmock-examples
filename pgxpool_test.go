package main

import (
	"testing"

	"github.com/pashagolub/pgxmock/v3"
)

// embed pgxmock.PgxConnIface
type PgxPoolMock struct {
	pgxmock.PgxPoolIface
}

// add impl by your need
// ...

func TestPgxPoolHandler(t *testing.T) {
	mock, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
	if err != nil {
		t.Error(err)
	}

	PgxPool = func() PgxPoolIface {
		return &PgxPoolMock{mock}
	}

	row := pgxmock.NewRows([]string{"result"}).AddRow(1)
	mock.ExpectQuery("select").WillReturnRows(row)

	if err = PgxPoolHandler(); err != nil {
		t.Error(err)
	}
}
