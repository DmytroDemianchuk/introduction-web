package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dmytrodemianchuk/go-auth-mongo/internal/domain"
	"github.com/dmytrodemianchuk/go-auth-mongo/pkg/hash" // Import the hash package
	"github.com/golang-jwt/jwt"
)

type UsersRepository interface {
	Create(ctx context.Context, user domain.User) error
	GetByCredentials(ctx context.Context, email string) (domain.User, error) // Fetch by email only
}

type Users struct {
	repo       UsersRepository
	hasher     *hash.Hasher
	hmacSecret []byte
	tokenTtl   time.Duration
}

func NewUsers(repo UsersRepository, hasher *hash.Hasher, secret []byte, ttl time.Duration) *Users {
	return &Users{
		repo:       repo,
		hasher:     hasher,
		hmacSecret: secret,
		tokenTtl:   ttl,
	}
}

func (s *Users) SignUp(ctx context.Context, inp domain.SignUpInput) error {
	password, err := s.hasher.Hash(inp.Password)
	if err != nil {
		return err
	}

	user := domain.User{
		Name:         inp.Name,
		Email:        inp.Email,
		Password:     password,
		RegisteredAt: time.Now(),
	}

	return s.repo.Create(ctx, user)
}

func (s *Users) SignIn(ctx context.Context, inp domain.SignInInput) (string, error) {
	user, err := s.repo.GetByCredentials(ctx, inp.Email)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return "", domain.ErrUserNotFound
		}
		return "", err
	}

	if err := s.hasher.Compare(user.Password, inp.Password); err != nil {
		return "", domain.ErrUserNotFound
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   user.ID,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(s.tokenTtl).Unix(),
	})

	return token.SignedString(s.hmacSecret)
}

func (s *Users) ParseToken(ctx context.Context, token string) (string, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return s.hmacSecret, nil
	})
	if err != nil {
		return "", err
	}

	if !t.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	subject, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("invalid subject")
	}

	return subject, nil
}
