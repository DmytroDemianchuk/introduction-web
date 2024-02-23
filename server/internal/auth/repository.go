package auth

import (
	"context"

	"github.com/dmytrodemianchuk/go-auth/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, username, password string) (*models.User, error)
}
