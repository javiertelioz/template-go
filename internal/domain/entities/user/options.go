package user

import (
	"fmt"
	"time"
)

type UserOption[T string | int64] func(*User[T])

func NewUser[T string | int64](opts ...UserOption[T]) *User[T] {
	user := &User[T]{}

	for _, opt := range opts {
		opt(user)
	}

	return user
}

func WithID[T string | int64](id T) UserOption[T] {
	return func(u *User[T]) {
		u.id = id
	}
}

func WithName[T string | int64](name T) UserOption[T] {
	return func(u *User[T]) {
		u.name = fmt.Sprint(name)
	}
}

func WithEmail[T string | int64](email T) UserOption[T] {
	return func(u *User[T]) {
		u.email = fmt.Sprint(email)
	}
}

func WithDob[T string | int64](dob time.Time) UserOption[T] {
	return func(u *User[T]) {
		u.dob = dob
	}
}
