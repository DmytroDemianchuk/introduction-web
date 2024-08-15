package repository

import (
	"context"

	"github.com/dmytrodemianchuk/go-auth-mongo/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BooksRepository interface {
	Create(ctx context.Context, book domain.Book) error
	GetByID(ctx context.Context, id string) (domain.Book, error)
	GetAll(ctx context.Context) ([]domain.Book, error)
	Update(ctx context.Context, id string, book domain.Book) error
	Delete(ctx context.Context, id string) error
}

type mongoBooksRepository struct {
	collection *mongo.Collection
}

func NewBooksRepository(db *mongo.Database) BooksRepository {
	return &mongoBooksRepository{
		collection: db.Collection("books"),
	}
}

func (r *mongoBooksRepository) Create(ctx context.Context, book domain.Book) error {
	_, err := r.collection.InsertOne(ctx, book)
	return err
}

func (r *mongoBooksRepository) GetByID(ctx context.Context, id string) (domain.Book, error) {
	var book domain.Book
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&book)
	return book, err
}

func (r *mongoBooksRepository) GetAll(ctx context.Context) ([]domain.Book, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var books []domain.Book
	for cursor.Next(ctx) {
		var book domain.Book
		if err := cursor.Decode(&book); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *mongoBooksRepository) Update(ctx context.Context, id string, book domain.Book) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": book})
	return err
}

func (r *mongoBooksRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
