package main

import (
	"context"
	"database/sql"
	"log"
	"log/slog"
	"net/http"

	"github.com/Valeron93/file-storage/backend/migrations"
	"github.com/Valeron93/file-storage/backend/vite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

func main() {

	db := openDB()
	defer db.Close()

	if err := migrations.RunMigrations(db); err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)

	r.Handle("/*", vite.Handler)

	r.Route("/api", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})
	})

	const addr = ":3000"
	slog.Info("listening http", slog.Any("addr", addr))
	if err := http.ListenAndServe(addr, r); err != nil {
		slog.Error("error: http", "err", err)
	}
}

func openDB() *sql.DB {

	db, err := sql.Open("sqlite", "./tmp/db.sqlite3")
	if err != nil {
		log.Panic(err)
	}

	return db
}
