package domain

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type ID struct {
	value string
}

func (i ID) Value() string {
	return i.value
}
func NewIDFromString(value string) (ID, error) {
	if value == "" {
		return ID{}, errors.New("id vazio não é válido")
	}
	return ID{value: value}, nil
}
func (i ID) GenerateNew() (ID, error) {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1000)
	newValue := fmt.Sprintf("cv-%03d", randomNum)

	if newValue == "" {
		return ID{}, errors.New("generated ID is empty")
	}

	return ID{value: newValue}, nil
}
