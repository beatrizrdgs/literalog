package authors

import (
	"github.com/literalog/library/pkg/models"
)

type Validator struct{}

func (v *Validator) Validate(a *models.Author) error {
	switch {
	case v.validateName(a.Name) != nil:
		return ErrInvalidName
	default:
		return nil
	}
}

func (v *Validator) validateName(name string) error {
	if name == "" {
		return ErrEmptyName
	}
	return nil
}
