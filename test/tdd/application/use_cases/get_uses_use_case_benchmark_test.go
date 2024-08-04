package use_cases

import (
	"errors"
	"testing"

	"github.com/javiertelioz/template-clean-architecture-go/internal/application/use_cases"
	userentity "github.com/javiertelioz/template-clean-architecture-go/internal/domain/entities/user"
	"github.com/javiertelioz/template-clean-architecture-go/test/mocks/repository"
)

func benchmarkGetUsesUseCase(b *testing.B, domainUsers []*userentity.User[string], getUsersError error) {
	userRepository := new(repository.MockUserRepository)
	useCase := use_cases.NewGetUsesUseCase(userRepository)

	userRepository.On("GetUsers").Return(domainUsers, getUsersError)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = useCase.Execute()
	}
}

func BenchmarkGetUsesUseCaseWithUsers(b *testing.B) {
	domainUsers := []*userentity.User[string]{
		userentity.NewUser(
			userentity.WithID("1"),
			userentity.WithName("John Doe"),
			userentity.WithEmail("john.doe@example.com"),
		),
		userentity.NewUser(
			userentity.WithID("2"),
			userentity.WithName("Jane Doe"),
			userentity.WithEmail("jane.doe@example.com"),
		),
	}
	benchmarkGetUsesUseCase(b, domainUsers, nil)
}

func BenchmarkGetUsesUseCaseWithNoUsers(b *testing.B) {
	domainUsers := []*userentity.User[string]{}
	benchmarkGetUsesUseCase(b, domainUsers, nil)
}

func BenchmarkGetUsesUseCaseWithRepositoryError(b *testing.B) {
	benchmarkGetUsesUseCase(b, nil, errors.New("repository error"))
}
