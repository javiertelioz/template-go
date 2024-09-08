package use_cases

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/javiertelioz/template-clean-architecture-go/src/application/dto/user"
	"github.com/javiertelioz/template-clean-architecture-go/src/application/use_cases"
	userentity "github.com/javiertelioz/template-clean-architecture-go/src/domain/entities/user"
	"github.com/javiertelioz/template-clean-architecture-go/test/mocks/repository"
)

func benchmarkUpdateUserByIDUseCase(b *testing.B, dto user.UpdateUserDTO, existingUser *userentity.User[string], getByIDError, updateError error) {
	userRepository := new(repository.MockUserRepository)
	useCase := use_cases.NewUpdateUserByIDUseCase(userRepository)

	userRepository.On("GetByID", dto.ID).Return(existingUser, getByIDError)
	userRepository.On("Update", mock.Anything).Return(updateError)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = useCase.Execute(dto)
	}
}

func BenchmarkUpdateUserByIDUseCaseWithValidUser(b *testing.B) {
	dto := user.UpdateUserDTO{
		ID:    "1",
		Name:  "John Doe Updated",
		Email: "john.doe.updated@example.com",
		Dob: user.DateOfBirth{
			Time: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	existingUser := userentity.NewUser(
		userentity.WithID(dto.ID),
		userentity.WithName("John Doe"),
		userentity.WithEmail("john.doe@example.com"),
		userentity.WithDob[string](time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)),
	)
	benchmarkUpdateUserByIDUseCase(b, dto, existingUser, nil, nil)
}

func BenchmarkUpdateUserByIDUseCaseWithNonExistentID(b *testing.B) {
	dto := user.UpdateUserDTO{
		ID:    "non-existent-id",
		Name:  "John Doe Updated",
		Email: "john.doe.updated@example.com",
		Dob: user.DateOfBirth{
			Time: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	benchmarkUpdateUserByIDUseCase(b, dto, nil, errors.New("user not found"), nil)
}

func BenchmarkUpdateUserByIDUseCaseWithInvalidUser(b *testing.B) {
	dto := user.UpdateUserDTO{
		ID:    "1",
		Name:  "",
		Email: "invalid-email",
		Dob: user.DateOfBirth{
			Time: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	existingUser := userentity.NewUser(
		userentity.WithID(dto.ID),
		userentity.WithName("John Doe"),
		userentity.WithEmail("john.doe@example.com"),
		userentity.WithDob[string](time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)),
	)
	benchmarkUpdateUserByIDUseCase(b, dto, existingUser, nil, nil)
}

func BenchmarkUpdateUserByIDUseCaseWithRepositoryError(b *testing.B) {
	dto := user.UpdateUserDTO{
		ID:    "1",
		Name:  "John Doe Updated",
		Email: "john.doe.updated@example.com",
		Dob: user.DateOfBirth{
			Time: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	existingUser := userentity.NewUser(
		userentity.WithID(dto.ID),
		userentity.WithName("John Doe"),
		userentity.WithEmail("john.doe@example.com"),
		userentity.WithDob[string](time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)),
	)
	benchmarkUpdateUserByIDUseCase(b, dto, existingUser, nil, errors.New("repository error"))
}
