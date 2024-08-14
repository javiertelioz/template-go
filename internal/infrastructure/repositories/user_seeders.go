package repositories

import (
	"time"

	"github.com/javiertelioz/template-clean-architecture-go/internal/domain/entities/user"
)

func setup(repo *InMemoryUserRepository) {
	user1 := user.NewUser[string](
		user.WithID("e17d5135-989e-4977-99ef-495c0ab7cd00"),
		user.WithName("Joe"),
		user.WithEmail("joe@example.com"),
	)

	user2 := user.NewUser[string](
		user.WithID("a0ff5e6e-908c-4534-b83f-cd85a66bda0a"),
		user.WithName("Jane"),
		user.WithEmail("jane@example.com"),
		user.WithDob[string](time.Now().AddDate(-30, 0, 0)),
	)

	repo.users[user1.GetID()] = user1
	repo.users[user2.GetID()] = user2
}
