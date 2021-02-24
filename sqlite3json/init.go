package sqlite3json

import (
	"database/sql"

	sqlite "github.com/mattn/go-sqlite3"
	"github.com/pauloavelar/go-sqlite3-json/sqlite3json/functions"
)

func init() {
	sql.Register(DialectName, &sqlite.SQLiteDriver{
		ConnectHook: func(conn *sqlite.SQLiteConn) error {
			for name, fn := range functions.Enabled {
				if err := conn.RegisterFunc(name, fn, true); err != nil {
					return err
				}
			}
			return nil
		},
	})
}
