package localstorage

import (
	"context"
	"sync"

	"github.com/dmytrodemianchuk/go-auth/internal/auth"
	"github.com/dmytrodemianchuk/go-auth/internal/models"
)

type UserLocalStorage struct {
	users map[string]*models.User
	mutex *sync.Mutex
}

func NewUserLocalStorage() *UserLocalStorage {
	return &UserLocalStorage{
		users: make(map[string]*models.User),
		mutex: new(sync.Mutex),
	}
}

func (s *UserLocalStorage) CreateUser(ctx context.Context, user *models.User) error {
	s.mutex.Lock()
	s.users[user.ID] = user
	s.mutex.Unlock()

	return nil
}

func (s *UserLocalStorage) GetUser(ctx context.Context, email, password string) (*models.User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, user := range s.users {
		if user.Email == email && user.Password == password {
			return user, nil
		}
	}
	return nil, auth.ErrUserNotFound
}
