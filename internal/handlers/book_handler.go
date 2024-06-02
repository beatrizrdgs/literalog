package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/beatrizrdgs/literalog/internal/models"
	"github.com/beatrizrdgs/literalog/internal/services"
	"github.com/go-chi/chi/v5"
)

type BookHandler struct {
	svc *services.BookService
}

func NewBookHandler(svc *services.BookService) *BookHandler {
	return &BookHandler{svc: svc}
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req models.BookRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.svc.CreateBook(ctx, models.NewBook(req))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *BookHandler) GetBookById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	fmt.Println(id)
	book, err := h.svc.GetBookById(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	books, err := h.svc.GetAllBooks(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
