package use_cases

import (
	"errors"
	"testing"

	"github.com/javiertelioz/template-clean-architecture-go/src/application/use_cases"
	userentity "github.com/javiertelioz/template-clean-architecture-go/src/domain/entities/user"
	"github.com/javiertelioz/template-clean-architecture-go/test/mocks/repository"
)

func benchmarkDeleteUserByIDUseCase(
	b *testing.B,
	id string,
	existingUser *userentity.User[string],
	getByIDError error,
	deleteError error,
) {
	userRepository := new(repository.MockUserRepository)
	useCase := use_cases.NewDeleteUserByIDUseCase(userRepository)

	userRepository.On("GetByID", id).Return(existingUser, getByIDError)
	userRepository.On("Delete", id).Return(deleteError)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = useCase.Execute(id)
	}
}

func BenchmarkDeleteUserByIDUseCaseWithValidID(b *testing.B) {
	id := "e17d5135-989e-4977-99ef-495c0ab7cd00"
	existingUser := userentity.NewUser(
		userentity.WithID(id),
		userentity.WithName("John Doe"),
		userentity.WithEmail("john.doe@example.com"),
	)
	benchmarkDeleteUserByIDUseCase(b, id, existingUser, nil, nil)
}

func BenchmarkDeleteUserByIDUseCaseWithNonExistentID(b *testing.B) {
	id := "non-existent-id"
	benchmarkDeleteUserByIDUseCase(b, id, nil, errors.New("user not found"), nil)
}

func BenchmarkDeleteUserByIDUseCaseWithRepositoryError(b *testing.B) {
	id := "e17d5135-989e-4977-99ef-495c0ab7cd00"
	existingUser := userentity.NewUser(
		userentity.WithID(id),
		userentity.WithName("John Doe"),
		userentity.WithEmail("john.doe@example.com"),
	)

	benchmarkDeleteUserByIDUseCase(b, id, existingUser, nil, errors.New("repository error"))
}
