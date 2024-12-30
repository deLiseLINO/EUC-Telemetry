package clickhouseconnect

import (
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

func New(address, user, password, database string) (*driver.Conn, error) {
	conn, err := clickhouse.Open(
		&clickhouse.Options{
			Addr: []string{address},
			Auth: clickhouse.Auth{
				Database: database,
				Username: user,
				Password: password,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return &conn, nil
}
