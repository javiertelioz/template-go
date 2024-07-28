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

func NewUser[T string | int64](opts ...UserOption[T]) *User[T] {
	user := &User[T]{}
	for _, opt := range opts {
		opt(user)
	}
	return user
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

type UserOption[T string | int64] func(*User[T])

func WithID[T string | int64](id T) UserOption[T] {
	return func(u *User[T]) {
		u.id = id
	}
}

func WithName[T string | int64](name T) UserOption[T] {
	return func(u *User[T]) {
		u.name = string(name)
	}
}

func WithEmail[T string | int64](email T) UserOption[T] {
	return func(u *User[T]) {
		u.email = string(email)
	}
}
