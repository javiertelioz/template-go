package use_cases

import (
	"fmt"
	"github.com/javiertelioz/template-clean-architecture-go/src/application/dto/user"
	userentity "github.com/javiertelioz/template-clean-architecture-go/src/domain/entities/user"
	"github.com/javiertelioz/template-clean-architecture-go/src/domain/repositories"
)

type UpdateUserByIDUseCase struct {
	userRepository repositories.UserRepository
}

func NewUpdateUserByIDUseCase(userRepository repositories.UserRepository) *UpdateUserByIDUseCase {
	return &UpdateUserByIDUseCase{
		userRepository: userRepository,
	}
}

func (uc *UpdateUserByIDUseCase) Execute(user user.UpdateUserDTO) error {
	_, err := uc.userRepository.GetByID(user.ID)
	if err != nil {
		return err
	}

	domainUser := userentity.NewUser[string](
		userentity.WithID(user.ID),
		userentity.WithName(user.Name),
		userentity.WithEmail(user.Email),
		userentity.WithDob[string](user.Dob.Time),
	)

	validationErrors := domainUser.Validate()
	if !validationErrors.IsEmpty() {
		fmt.Printf("Validation errors: %+v\n", validationErrors)
		return validationErrors
	}

	err = uc.userRepository.Update(domainUser)
	if err != nil {
		return err
	}

	return nil
}
