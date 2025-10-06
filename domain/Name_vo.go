package domain

import (
	"errors"
	"strings"
)

type Name struct {
	value string
}

func NewName(value string) (Name, error) {
	if value := strings.TrimSpace(value); value == "" {
		return Name{}, errors.New("nome não pode estar vazio ou ser simplesmente espaço")
	}

	return Name{value: value}, nil
}

func (n *Name) Value() string {
	return n.value
}
