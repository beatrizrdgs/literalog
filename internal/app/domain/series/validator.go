package series

import "github.com/literalog/library/pkg/models"

type Validator struct{}

func (v *Validator) Validate(series *models.Series) error {
	switch {
	case v.validateName(series.Name) != nil:
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
