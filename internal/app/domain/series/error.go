package series

import (
	"net/http"

	"github.com/literalog/cerrors"
)

var (
	ErrEmptyId = cerrors.New("empty id", http.StatusBadRequest)
)
