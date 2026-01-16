//go:build !dist

package frontend

import (
	"net/http"
)

var Handler = http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}))
