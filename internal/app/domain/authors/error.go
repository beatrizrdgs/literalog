package authors

import (
	"net/http"

	"github.com/literalog/cerrors"
)

var (
	ErrEmptyID     = cerrors.New("empty author id", http.StatusBadRequest)
	ErrInvalidName = cerrors.New("invalid name", http.StatusBadRequest)
	ErrEmptyName   = cerrors.New("empty name", http.StatusBadRequest)
	ErrNotFound    = cerrors.New("author not found", http.StatusNotFound)
)
