package main

import (
	"log"
	"net/http"

	"github.com/Valeron93/file-storage/backend/api"
	"github.com/Valeron93/file-storage/backend/auth"
	"github.com/Valeron93/file-storage/backend/migrations"
	"github.com/Valeron93/file-storage/cmd/database"
	"github.com/Valeron93/file-storage/frontend"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	db, err := database.OpenSQLite("./tmp/db.sqlite")
	if err != nil {
		log.Panicf("failed to open database: %v", err)
	}

	defer db.Close()

	if err := migrations.RunMigrations(db); err != nil {
		panic(err)
	}

	auth := auth.NewSQLite(db)

	authApi := api.NewAuthAPI(auth)

	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)

	r.Handle("/*", frontend.Handler)

	r.Route("/api", func(r chi.Router) {
		r.Post("/register", authApi.HandleRegister)
		r.Post("/login", authApi.HandleLogin)
	})

	const addr = ":3000"
	log.Printf("listening HTTP on %#v", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Printf("http error: %v", err)
	}
}
