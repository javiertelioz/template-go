package use_cases

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/javiertelioz/template-clean-architecture-go/src/application/dto/user"
	"github.com/javiertelioz/template-clean-architecture-go/src/application/use_cases"
	userentity "github.com/javiertelioz/template-clean-architecture-go/src/domain/entities/user"
	"github.com/javiertelioz/template-clean-architecture-go/test/mocks/repository"
)

func benchmarkCreateUserUseCase(b *testing.B, dto user.CreateUserDTO, existingUser *userentity.User[string], getByEmailError, createError error) {
	userRepository := new(repository.MockUserRepository)
	useCase := use_cases.NewCreateUserUseCase(userRepository)

	userRepository.On("GetByEmail", dto.Email).Return(existingUser, getByEmailError)
	userRepository.On("Create", mock.Anything).Return(createError)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = useCase.Execute(dto)
	}
}

func BenchmarkCreateUserUseCaseWithValidUser(b *testing.B) {
	dto := user.CreateUserDTO{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	existingUser := userentity.NewUser[string]()
	benchmarkCreateUserUseCase(b, dto, existingUser, nil, nil)
}

func BenchmarkCreateUserUseCaseWithExistingEmail(b *testing.B) {
	dto := user.CreateUserDTO{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	existingUser := userentity.NewUser(
		userentity.WithID("e17d5135-989e-4977-99ef-495c0ab7cd00"),
		userentity.WithName("Jane Doe"),
		userentity.WithEmail(dto.Email),
	)
	benchmarkCreateUserUseCase(b, dto, existingUser, nil, errors.New("email already exists"))
}

func BenchmarkCreateUserUseCaseWithInvalidUser(b *testing.B) {
	dto := user.CreateUserDTO{
		Name:  "",
		Email: "invalid-email",
	}
	existingUser := userentity.NewUser[string]()
	benchmarkCreateUserUseCase(b, dto, existingUser, nil, nil)
}

func BenchmarkCreateUserUseCaseWithValidUserWhenCallExecuteThenReturnErrorOnGetUserByEmail(b *testing.B) {
	dto := user.CreateUserDTO{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	existingUser := userentity.NewUser[string]()
	benchmarkCreateUserUseCase(b, dto, existingUser, errors.New("provider error"), nil)
}
