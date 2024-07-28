package user

import "strings"

type Name struct {
	value string
}

func NewName(name string) Name {
	return Name{
		strings.ToLower(name),
	}
}

func (n Name) GetValue() string {
	return n.value
}
