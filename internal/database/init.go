package database

import (
	"database/sql"
	"database/sql/driver"

	"modernc.org/sqlite"
)

// SQLiteDriverWithForeignKeys wraps the sqlite driver to enable foreign keys
type SQLiteDriverWithForeignKeys struct {
	*sqlite.Driver
}

func (d *SQLiteDriverWithForeignKeys) Open(name string) (driver.Conn, error) {
	conn, err := d.Driver.Open(name)
	if err != nil {
		return nil, err
	}

	// Enable foreign keys on connection
	if execer, ok := conn.(driver.Execer); ok {
		_, err = execer.Exec("PRAGMA foreign_keys = ON", nil)
		if err != nil {
			conn.Close()
			return nil, err
		}
	}

	return conn, nil
}

func init() {
	sql.Register("sqlite3", &SQLiteDriverWithForeignKeys{Driver: &sqlite.Driver{}})
}
