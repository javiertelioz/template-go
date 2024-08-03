package user

import (
	"time"

	"github.com/javiertelioz/template-clean-architecture-go/internal/domain/entities"
)

type User[T string | int64] struct {
	id    T
	name  string
	email string
	dob   time.Time
}

func (u *User[T]) GetID() T {
	return u.id
}

func (u *User[T]) GetName() string {
	return u.name
}

func (u *User[T]) GetEmail() string {
	return u.email
}

func (u *User[T]) GetDob() time.Time {
	return u.dob
}

func (u *User[T]) Validate() *entities.ValidationErrors {
	errs := &entities.ValidationErrors{}

	if err := ValidateName(u.name); err != nil {
		errs.Add(err)
	}

	if err := ValidateEmail(u.email); err != nil {
		errs.Add(err)
	}

	return errs
}
