package rest

import (
	"net/http"
	"strings"

	"github.com/dmytrodemianchuk/go-auth-mongo/internal/service"
)

type AuthMiddleware struct {
	UsersService *service.Users
}

func NewAuthMiddleware(usersService *service.Users) *AuthMiddleware {
	return &AuthMiddleware{UsersService: usersService}
}

func (am *AuthMiddleware) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")
		_, err := am.UsersService.ParseToken(r.Context(), token)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
