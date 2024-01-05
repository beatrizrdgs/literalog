package book

import "github.com/literalog/library/pkg/models"

type Validator struct{}

func (v *Validator) Validate(b *models.Book) error {
	switch {
	case v.validateTitle(b.Title) != nil:
		return ErrInvalidTitle
	default:
		return nil
	}
}

func (v *Validator) validateTitle(title string) error {
	if title == "" {
		return ErrEmptyTitle
	}

	if len(title) < 3 || len(title) > 50 {
		return ErrInvalidTitleLength
	}

	return nil
}
