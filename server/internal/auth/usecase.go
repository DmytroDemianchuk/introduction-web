package auth

import (
	"context"

	"github.com/dmytrodemianchuk/go-auth/internal/models"
)

const CtxUserKey = "user"

type UseCase interface {
	SignUp(ctx context.Context, username, password string) error
	SignIn(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accesToken string) (*models.User, error)
}
