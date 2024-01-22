package book

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/literalog/cerrors"
	"github.com/literalog/library/pkg/models"
)

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request)
	CreateByISBN(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)

	Routes() *chi.Mux
}

type handler struct {
	service Service
	router  *chi.Mux
}

func NewHandler(s Service) Handler {
	h := &handler{
		service: s,
		router:  chi.NewRouter(),
	}

	h.setupRoutes()

	return h
}

func (h *handler) setupRoutes() {
	h.router.Post("/", h.Create)
	h.router.Post("/isbn/{isbn}", h.CreateByISBN)
	h.router.Put("/", h.Update)
	h.router.Delete("/{id}", h.Delete)
	h.router.Get("/{id}", h.GetByID)
	h.router.Get("/", h.GetAll)
}

func (h handler) Routes() *chi.Mux {
	return h.router
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := new(models.BookRequest)
	json.NewDecoder(r.Body).Decode(&req)

	book := models.NewBook(*req)
	if err := h.service.Create(ctx, book); err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (h *handler) CreateByISBN(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	isbn := chi.URLParam(r, "isbn")

	if err := h.service.CreateByISBN(ctx, isbn); err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := new(models.BookRequest)
	json.NewDecoder(r.Body).Decode(&req)

	book := models.NewBook(*req)
	if err := h.service.Update(ctx, book); err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	if err := h.service.Delete(ctx, id); err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	book, err := h.service.GetByID(ctx, id)
	if err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	books, err := h.service.GetAll(ctx)
	if err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
