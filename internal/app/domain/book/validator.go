package book

import "github.com/literalog/library/pkg/models"

type Validator struct{}

func (v *Validator) Validate(book *models.Book) error {
	switch {
	case v.validateTitle(book.Title) != nil:
		return ErrInvalidTitle
	default:
		return nil
	}
}

func (v *Validator) validateTitle(title string) error {
	if title == "" {
		return ErrEmptyTitle
	}

	if len(title) < 2 || len(title) > 100 {
		return ErrInvalidTitleLength
	}

	return nil
}
