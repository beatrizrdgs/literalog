package genre

import "github.com/literalog/library/pkg/models"

type Validator struct{}

func (v *Validator) Validate(genre *models.Genre) error {
	switch {
	case v.validateName(genre.Name) != nil:
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
