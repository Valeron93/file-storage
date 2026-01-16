package api

import (
	"net/http"

	"github.com/Valeron93/file-storage/backend/auth"
)

type AuthAPI struct {
	a auth.Auth
}

func NewAuthAPI(a auth.Auth) *AuthAPI {
	return &AuthAPI{
		a: a,
	}
}

func (a *AuthAPI) HandleRegister(w http.ResponseWriter, r *http.Request) {

}

func (a *AuthAPI) HandleLogin(w http.ResponseWriter, r *http.Request) {

}
