package book

import (
	"net/http"

	"github.com/literalog/cerrors"
)

var (
	ErrEmptyId            = cerrors.New("empty id", http.StatusBadRequest)
	ErrInvalidTitle       = cerrors.New("invalid title", http.StatusBadRequest)
	ErrEmptyTitle         = cerrors.New("empty title", http.StatusBadRequest)
	ErrInvalidTitleLength = cerrors.New("title must be between x and y", http.StatusBadRequest)
	ErrAuthorNotFound     = cerrors.New("", http.StatusBadRequest)
)
