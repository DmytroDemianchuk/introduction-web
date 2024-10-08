package repository

import (
	"context"

	"github.com/dmytrodemianchuk/go-auth-mongo/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepository struct {
	db *mongo.Database
}

func NewUsersRepository(db *mongo.Database) *UsersRepository {
	return &UsersRepository{db: db}
}

func (r *UsersRepository) Create(ctx context.Context, user domain.User) error {
	_, err := r.db.Collection("users").InsertOne(ctx, user)
	return err
}

func (r *UsersRepository) GetByCredentials(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	err := r.db.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return user, domain.ErrUserNotFound
	}
	if err != nil {
		return user, err
	}
	return user, nil
}
