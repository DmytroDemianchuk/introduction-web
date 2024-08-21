package repository

import (
	"context"
	"errors"

	"github.com/dmytrodemianchuk/go-auth-mongo/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepository struct {
	collection *mongo.Collection
}

func NewUsersRepository(db *mongo.Database) *UsersRepository {
	return &UsersRepository{
		collection: db.Collection("users"),
	}
}

func (r *UsersRepository) Create(ctx context.Context, user domain.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *UsersRepository) GetByID(ctx context.Context, id string) (domain.User, error) {
	var user domain.User
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{}, err
	}
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}
	return user, nil
}

func (r *UsersRepository) GetByCredentials(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}
	return user, nil
}

func (r *UsersRepository) GetByName(ctx context.Context, name string) (domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(ctx, bson.M{"name": name}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}
	return user, nil
}
