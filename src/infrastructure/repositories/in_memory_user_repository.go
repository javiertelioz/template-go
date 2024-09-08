package repositories

import (
	"errors"
	"sync"

	"github.com/google/uuid"

	"github.com/javiertelioz/template-clean-architecture-go/src/domain/entities/user"
	"github.com/javiertelioz/template-clean-architecture-go/src/domain/repositories"
)

type InMemoryUserRepository struct {
	mu    sync.RWMutex
	users map[string]*user.User[string]
}

func NewInMemoryUserRepository() repositories.UserRepository {
	repo := &InMemoryUserRepository{
		users: make(map[string]*user.User[string]),
	}

	setup(repo)

	return repo
}

func (r *InMemoryUserRepository) Create(u *user.User[string]) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[u.GetID()]; exists {
		return errors.New("user already exists")
	}

	id := uuid.New().String()
	newUser := user.NewUser[string](
		user.WithID(id),
		user.WithName(u.GetName()),
		user.WithEmail(u.GetEmail()),
	)

	r.users[id] = newUser

	return nil
}

func (r *InMemoryUserRepository) GetUsers() ([]*user.User[string], error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var users []*user.User[string]
	for _, u := range r.users {
		users = append(users, u)
	}

	return users, nil
}

func (r *InMemoryUserRepository) GetByID(id string) (*user.User[string], error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	u, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	return u, nil
}

func (r *InMemoryUserRepository) GetByEmail(email string) (*user.User[string], error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, u := range r.users {
		if u.GetEmail() == email {
			return u, nil
		}
	}

	return &user.User[string]{}, errors.New("user not found")
}

func (r *InMemoryUserRepository) Update(u *user.User[string]) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[u.GetID()]; !exists {
		return errors.New("user not found")
	}

	r.users[u.GetID()] = u
	return nil
}

func (r *InMemoryUserRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}

	delete(r.users, id)
	return nil
}
