package auth

import (
	"context"

	"github.com/Valeron93/file-storage/backend/model"
)

type Auth interface {
	RegisterUser(
		ctx context.Context,
		username, password string,
	) (model.User, error)

	AuthenticateUser(
		ctx context.Context,
		username, password string,
	) (model.Session, error)

	GetSession(
		ctx context.Context,
		token string,
	) (model.Session, error)
}
