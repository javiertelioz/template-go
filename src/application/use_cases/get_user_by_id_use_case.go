package use_cases

import (
	"github.com/javiertelioz/template-clean-architecture-go/src/application/dto/user"
	"github.com/javiertelioz/template-clean-architecture-go/src/domain/repositories"
)

type GetUserByIDUseCase struct {
	userRepository repositories.UserRepository
}

func NewGetUserByIDUseCase(userRepository repositories.UserRepository) *GetUserByIDUseCase {
	return &GetUserByIDUseCase{
		userRepository: userRepository,
	}
}

func (uc *GetUserByIDUseCase) Execute(id string) (user.GetUserDTO, error) {
	domainUser, err := uc.userRepository.GetByID(id)

	if err != nil {
		return user.GetUserDTO{}, err
	}

	return user.GetUserDTO{
		ID:    domainUser.GetID(),
		Name:  domainUser.GetName(),
		Email: domainUser.GetEmail(),
	}, nil
}
