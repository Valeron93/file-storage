package auth

import (
	"context"
	"net/http"

	"github.com/Valeron93/file-storage/backend/model"
)

const SessionCookieName = "session_id"

type MiddlewareFunc func(next http.Handler) http.Handler
type SessionCtxKey struct{}

// will error with 401 Unauthorized if no valid
// cookie was provided ([SessionCookieName]), otherwise will inject
// [model.Session] into request's [context.Context]
// using [SessionCtxKey]
func MustBeAuthorized(auth Auth) MiddlewareFunc {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(SessionCookieName)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			token := cookie.Value
			session, err := auth.GetSession(r.Context(), token)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			newCtx := context.WithValue(r.Context(), SessionCtxKey{}, session)
			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}

// Get the session from [context.Context].
// Panics if session was not found in ctx,
// so should be used carefully
func SessionFromCtx(ctx context.Context) model.Session {
	session, ok := ctx.Value(SessionCtxKey{}).(model.Session)
	if !ok {
		panic("Failed to get session from context")
	}

	return session
}
