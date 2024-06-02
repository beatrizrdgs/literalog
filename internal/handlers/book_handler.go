package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/beatrizrdgs/literalog/internal/services"
	"github.com/go-chi/chi/v5"
)

type BookHandler struct {
	svc *services.BookService
}

func NewBookHandler(svc *services.BookService) *BookHandler {
	return &BookHandler{svc: svc}
}

// func (h *BookHandler) Add(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	req := new(models.BookRequest)
// 	json.NewDecoder(r.Body).Decode(req)
// 	book := models.NewBook(*req)
// 	err := h.svc.Add(ctx, book)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(book)
// }

func (h *BookHandler) GetById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	fmt.Println(id)
	book, err := h.svc.GetById(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
