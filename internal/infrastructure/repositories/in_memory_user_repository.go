package repositories

import (
	"errors"
	"github.com/google/uuid"

	"github.com/javiertelioz/template-clean-architecture-go/internal/domain/entities/user"
	"github.com/javiertelioz/template-clean-architecture-go/internal/domain/repositories"
)

type InMemoryUserRepository struct {
	users map[string]user.User[string]
}

func NewInMemoryUserRepository() repositories.UserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]user.User[string]),
	}
}

func (repo *InMemoryUserRepository) Create(user *user.User[string]) error {
	if _, exists := repo.users[user.GetID()]; exists {
		return errors.New("user already exists")
	}

	id := uuid.New().String()

	repo.users[id] = *user

	return nil
}

func (repo *InMemoryUserRepository) GetByID(id string) (*user.User[string], error) {
	existsUser, exists := repo.users[id]

	if !exists {
		return &user.User[string]{}, nil
	}

	return &existsUser, nil
}

func (repo *InMemoryUserRepository) GetByEmail(email string) (*user.User[string], error) {
	existsUser, exists := repo.users[email]
	if !exists {
		return &user.User[string]{}, nil
	}

	return &existsUser, nil
}
