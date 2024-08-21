package service

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dmytrodemianchuk/go-auth-mongo/internal/domain"
	"github.com/dmytrodemianchuk/go-auth-mongo/internal/repository"
	"github.com/dmytrodemianchuk/go-auth-mongo/pkg/hash"
)

type Users struct {
	repo         *repository.UsersRepository
	hasher       *hash.Hasher
	jwtSecret    []byte
	tokenExpires time.Duration
}

func NewUsers(repo *repository.UsersRepository, hasher *hash.Hasher, jwtSecret []byte, tokenExpires time.Duration) *Users {
	return &Users{repo: repo, hasher: hasher, jwtSecret: jwtSecret, tokenExpires: tokenExpires}
}

type SignUpInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Users) SignUp(ctx context.Context, inp SignUpInput) error {
	hashedPassword, err := s.hasher.Hash(inp.Password)
	if err != nil {
		return err
	}

	user := domain.User{
		Name:         inp.Name,
		Email:        inp.Email,
		Password:     hashedPassword,
		RegisteredAt: time.Now(),
	}

	return s.repo.Create(ctx, user)
}

func (s *Users) SignIn(ctx context.Context, inp SignInInput) (string, error) {
	user, err := s.repo.GetByCredentials(ctx, inp.Email)
	if err != nil {
		return "", err
	}

	if err := s.hasher.Compare(user.Password, inp.Password); err != nil {
		return "", err
	}

	token, err := s.createToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Users) ParseToken(ctx context.Context, tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return s.jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, ok := claims["id"].(string)
		if !ok {
			return "", jwt.ErrSignatureInvalid
		}
		return id, nil
	}

	return "", jwt.ErrSignatureInvalid
}

func (s *Users) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Users) GetUserByName(ctx context.Context, name string) (domain.User, error) {
	return s.repo.GetByName(ctx, name)
}

func (s *Users) createToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(s.tokenExpires).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}
