package use_cases

import (
	"github.com/javiertelioz/template-clean-architecture-go/internal/application/dto/user"
	"github.com/javiertelioz/template-clean-architecture-go/internal/domain/repositories"
)

type GetUsesUseCase struct {
	userRepository repositories.UserRepository
}

func NewGetUsesUseCase(userRepository repositories.UserRepository) *GetUsesUseCase {
	return &GetUsesUseCase{
		userRepository: userRepository,
	}
}

func (uc *GetUsesUseCase) Execute() ([]user.GetUserDTO, error) {
	domainUsers, err := uc.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}

	var users []user.GetUserDTO

	for _, u := range domainUsers {
		users = append(users, user.GetUserDTO{
			ID:    u.GetID(),
			Name:  u.GetName(),
			Email: u.GetEmail(),
		})
	}

	return users, nil
}
