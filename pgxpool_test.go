package main

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/golang/mock/gomock"
)

type RegexpMatcher struct {
	pattern *regexp.Regexp
}

func Regexp(pattern string) *RegexpMatcher {
	return &RegexpMatcher{
		pattern: regexp.MustCompile(pattern),
	}
}

func (m *RegexpMatcher) String() string {
	return fmt.Sprintf("matches pattern /%v/", m.pattern)
}

func (m *RegexpMatcher) Matches(x interface{}) bool {
	s, ok := x.(string)
	if !ok {
		return false
	}

	return m.pattern.MatchString(s)
}

func TestPgxPoolHandler(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	PgxPool = func() pgxpoolmock.PgxPool {
		return mockPool
	}

	rows := pgxpoolmock.NewRows([]string{"result"}).AddRow(1).ToPgxRows()
	mockPool.EXPECT().Query(gomock.Any(), Regexp("select")).Return(rows, nil)

	if err := PgxPoolHandler(); err != nil {
		t.Error(err)
	}
}
