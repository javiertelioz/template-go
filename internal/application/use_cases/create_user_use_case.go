package use_cases

import (
	"errors"
	"fmt"

	"github.com/javiertelioz/template-clean-architecture-go/internal/application/dto/user"
	userentity "github.com/javiertelioz/template-clean-architecture-go/internal/domain/entities/user"
	"github.com/javiertelioz/template-clean-architecture-go/internal/domain/repositories"
)

type CreateUserUseCase struct {
	UserRepository repositories.UserRepository
}

func NewCreateUserUseCase(
	userRepository repositories.UserRepository,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: userRepository,
	}
}

func (uc *CreateUserUseCase) Execute(user user.CreateUserDTO) error {

	exists, _ := uc.UserRepository.GetByEmail(user.Email)

	if exists != nil {
		return errors.New("email already exists")
	}

	domainUser := userentity.NewUser(
		// userentity.WithID(user.ID),
		userentity.WithName(user.Name),
		userentity.WithEmail(user.Email),
	)

	validationErrors := domainUser.Validate()
	if !validationErrors.IsEmpty() {
		fmt.Printf("Errores de validación: %+v \n", validationErrors)
		return validationErrors
	}

	return uc.UserRepository.Create(domainUser)
}
