package use_cases

import (
	"github.com/javiertelioz/template-clean-architecture-go/internal/application/dto/user"
	"github.com/javiertelioz/template-clean-architecture-go/internal/domain/repositories"
)

type GetUserByIDUseCase struct {
	UserRepository repositories.UserRepository
}

func NewGetUserByIDUseCase(userRepository repositories.UserRepository) *GetUserByIDUseCase {
	return &GetUserByIDUseCase{
		UserRepository: userRepository,
	}
}

func (uc *GetUserByIDUseCase) Execute(id string) (user.GetUserDTO, error) {
	domainUser, err := uc.UserRepository.GetByID(id)

	if err != nil {
		return user.GetUserDTO{}, err
	}

	return user.GetUserDTO{
		ID:    domainUser.GetID(),
		Name:  domainUser.GetName(),
		Email: domainUser.GetEmail(),
	}, nil
}
