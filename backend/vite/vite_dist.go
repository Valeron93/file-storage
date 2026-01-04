//go:build dist

package vite

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed dist/**
var viteFS embed.FS

var Handler http.Handler

func init() {
	sub, err := fs.Sub(viteFS, "dist")
	if err != nil {
		panic(err)
	}
	Handler = http.FileServerFS(sub)
}
