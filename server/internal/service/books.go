package service

import (
	"context"

	"github.com/dmytrodemianchuk/go-auth-mongo/internal/domain"
	"github.com/dmytrodemianchuk/go-auth-mongo/internal/repository"
)

type Books struct {
	repo repository.BooksRepository
}

func NewBooks(repo repository.BooksRepository) *Books {
	return &Books{repo: repo}
}

func (s *Books) Create(ctx context.Context, book domain.Book) error {
	return s.repo.Create(ctx, book)
}

func (s *Books) GetByID(ctx context.Context, id string) (domain.Book, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Books) GetAll(ctx context.Context) ([]domain.Book, error) {
	return s.repo.GetAll(ctx)
}

func (s *Books) Update(ctx context.Context, id string, book domain.Book) error {
	return s.repo.Update(ctx, id, book)
}

func (s *Books) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
