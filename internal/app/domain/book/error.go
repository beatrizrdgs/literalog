package book

import (
	"net/http"

	"github.com/literalog/cerrors"
)

var (
	ErrEmptyID            = cerrors.New("empty book id", http.StatusBadRequest)
	ErrEmptyISBN          = cerrors.New("empty isbn", http.StatusBadRequest)
	ErrEmptyTitle         = cerrors.New("empty title", http.StatusBadRequest)
	ErrEmptyAuthors       = cerrors.New("empty authors", http.StatusBadRequest)
	ErrInvalidTitle       = cerrors.New("invalid title", http.StatusBadRequest)
	ErrInvalidTitleLength = cerrors.New("title must be between x and y", http.StatusBadRequest)
)
