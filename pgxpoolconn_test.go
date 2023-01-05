package main

import (
	"testing"

	"github.com/pashagolub/pgxmock"
)

// embed pgxmock.PgxConnIface
type PgxPoolConnMock struct {
	pgxmock.PgxConnIface
}

// "Release" is needed by pgxpool.Conn and usually used
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
