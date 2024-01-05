package author

import (
	"errors"

	"github.com/literalog/library/pkg/models"
)

type Validator struct{}

func (v *Validator) Validate(a *models.Author) error {
	switch {
	case v.validateName(a.Name) != nil:
		return errors.New("invalid author")
	default:
		return nil
	}
}

func (v *Validator) validateName(name string) error {
	if name == "" {
		return errors.New("empty name")
	}
	return nil
}
