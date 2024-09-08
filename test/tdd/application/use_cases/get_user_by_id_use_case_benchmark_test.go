package use_cases

import (
	"errors"
	"testing"

	"github.com/javiertelioz/template-clean-architecture-go/src/application/use_cases"
	userentity "github.com/javiertelioz/template-clean-architecture-go/src/domain/entities/user"
	"github.com/javiertelioz/template-clean-architecture-go/test/mocks/repository"
)

func benchmarkGetUserByIDUseCase(b *testing.B, id string, existingUser *userentity.User[string], getByIDError error) {
	userRepository := new(repository.MockUserRepository)
	useCase := use_cases.NewGetUserByIDUseCase(userRepository)

	userRepository.On("GetByID", id).Return(existingUser, getByIDError)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = useCase.Execute(id)
	}
}

func BenchmarkGetUserByIDUseCaseWithValidID(b *testing.B) {
	id := "e17d5135-989e-4977-99ef-495c0ab7cd00"
	existingUser := userentity.NewUser(
		userentity.WithID(id),
		userentity.WithName("John Doe"),
		userentity.WithEmail("john.doe@example.com"),
	)
	benchmarkGetUserByIDUseCase(b, id, existingUser, nil)
}

func BenchmarkGetUserByIDUseCaseWithNonExistentID(b *testing.B) {
	id := "non-existent-id"
	benchmarkGetUserByIDUseCase(b, id, nil, errors.New("user not found"))
}

func BenchmarkGetUserByIDUseCaseWithRepositoryError(b *testing.B) {
	id := "e17d5135-989e-4977-99ef-495c0ab7cd00"
	existingUser := userentity.NewUser(
		userentity.WithID(id),
		userentity.WithName("John Doe"),
		userentity.WithEmail("john.doe@example.com"),
	)
	benchmarkGetUserByIDUseCase(b, id, existingUser, errors.New("repository error"))
}
