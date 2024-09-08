package repository

import (
	"github.com/stretchr/testify/mock"

	userentity "github.com/javiertelioz/template-clean-architecture-go/src/domain/entities/user"
)

// MockUserRepository is a mock implementation of the UserRepository interface
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *userentity.User[string]) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUsers() ([]*userentity.User[string], error) {
	args := m.Called()
	return args.Get(0).([]*userentity.User[string]), args.Error(1)
}

func (m *MockUserRepository) GetByID(id string) (*userentity.User[string], error) {
	args := m.Called(id)
	return args.Get(0).(*userentity.User[string]), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(email string) (*userentity.User[string], error) {
	args := m.Called(email)
	return args.Get(0).(*userentity.User[string]), args.Error(1)
}

func (m *MockUserRepository) Update(user *userentity.User[string]) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
