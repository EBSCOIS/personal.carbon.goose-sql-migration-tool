package goose

import (
	"database/sql"
	"fmt"
)

// OpenDBWithDriver creates a connection to a database, and modifies goose
// internals to be compatible with the supplied driver by calling SetDialect.
func OpenDBWithDriver(driver string, dbstring string) (*sql.DB, error) {
	if err := SetDialect(driver); err != nil {
		return nil, err
	}

	switch driver {
	case "mssql":
		driver = "sqlserver"
	case "redshift":
		driver = "pgx"
	case "tidb":
		driver = "mysql"
	}

	switch driver {
	case "pgx", "sqlite3", "sqlite", "mysql", "sqlserver", "clickhouse", "vertica", "azuresql":
		return sql.Open(driver, dbstring)
	default:
		return nil, fmt.Errorf("unsupported driver %s", driver)
	}
}
