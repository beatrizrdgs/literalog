package author

import (
	"net/http"

	"github.com/literalog/cerrors"
)

var (
	ErrEmptyID     = cerrors.New("empty id", http.StatusBadRequest)
	ErrInvalidName = cerrors.New("invalid name", http.StatusBadRequest)
	ErrEmptyName   = cerrors.New("empty name", http.StatusBadRequest)
)
