package use_cases

import (
	"errors"
	"fmt"

	"github.com/javiertelioz/template-clean-architecture-go/internal/application/dto/user"
	userentity "github.com/javiertelioz/template-clean-architecture-go/internal/domain/entities/user"
	"github.com/javiertelioz/template-clean-architecture-go/internal/domain/repositories"
)

type CreateUserUseCase struct {
	userRepository repositories.UserRepository
}

func NewCreateUserUseCase(
	userRepository repositories.UserRepository,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *CreateUserUseCase) Execute(user user.CreateUserDTO) error {
	domainUser := userentity.NewUser(
		userentity.WithName(user.Name),
		userentity.WithEmail(user.Email),
	)

	validationErrors := domainUser.Validate()

	if !validationErrors.IsEmpty() {
		fmt.Printf("Errores de validación: %+v", validationErrors)
		return validationErrors
	}

	exists, err := uc.userRepository.GetByEmail(user.Email)

	if err != nil {
		return err
	}

	if exists.GetID() != "" {
		return errors.New("email already exists")
	}

	return uc.userRepository.Create(domainUser)
}
