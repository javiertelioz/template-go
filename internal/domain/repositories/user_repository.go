package repositories

import (
	"github.com/javiertelioz/template-clean-architecture-go/internal/domain/entities/user"
)

type UserRepository interface {
	Create(user *user.User[string]) error
	GetUsers() ([]*user.User[string], error)
	GetByID(id string) (*user.User[string], error)
	GetByEmail(email string) (*user.User[string], error)
	Update(user *user.User[string]) error
	Delete(id string) error
}
