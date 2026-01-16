package database

import (
	"context"
	"database/sql"

	"modernc.org/sqlite"
)

func init() {
	// SQLite settings
	sqlite.RegisterConnectionHook(func(conn sqlite.ExecQuerierContext, dsn string) error {
		query := `
		PRAGMA foreign_keys = ON;
		PRAGMA journal_mode = WAL;
		PRAGMA synchronous = NORMAL;
		PRAGMA temp_store = MEMORY;
		PRAGMA busy_timeout = 5000;
		PRAGMA mmap_size = 268435456;
		`
		_, err := conn.ExecContext(context.Background(), query, nil)
		return err
	})
}

func OpenSQLite(filename string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./tmp/db.sqlite3")
	return db, err
}
