package use_cases

import (
	"github.com/javiertelioz/template-clean-architecture-go/src/domain/repositories"
)

type DeleteUserByIDUseCase struct {
	userRepository repositories.UserRepository
}

func NewDeleteUserByIDUseCase(userRepository repositories.UserRepository) *DeleteUserByIDUseCase {
	return &DeleteUserByIDUseCase{
		userRepository: userRepository,
	}
}

func (uc *DeleteUserByIDUseCase) Execute(id string) error {
	_, err := uc.userRepository.GetByID(id)
	if err != nil {
		return err
	}

	err = uc.userRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
