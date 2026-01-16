package auth

import (
	"context"
	"database/sql"

	"github.com/Valeron93/file-storage/backend/model"
)

type authSQLite struct {
	db *sql.DB
}

func NewSQLite(db *sql.DB) Auth {
	return &authSQLite{
		db: db,
	}
}

// AuthenticateUser implements [Auth].
func (a *authSQLite) AuthenticateUser(ctx context.Context, username string, password string) (model.Session, error) {
	panic("unimplemented")
}

// RegisterUser implements [Auth].
func (a *authSQLite) RegisterUser(ctx context.Context, username string, password string) (model.User, error) {
	panic("unimplemented")
}

// GetSession implements [Auth].
func (a *authSQLite) GetSession(ctx context.Context, token string) (model.Session, error) {
	panic("unimplemented")
}
