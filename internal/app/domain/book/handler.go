package book

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

	Routes() *mux.Router
}

type handler struct {
	service Service
	router  *mux.Router
}

func NewHandler(s Service) Handler {
	h := &handler{
		service: s,
		router:  mux.NewRouter(),
	}

	h.setupRoutes()

	return h
}

func (h *handler) setupRoutes() {
	h.router.HandleFunc("/", h.Create).Methods(http.MethodPost)
	h.router.HandleFunc("/isbn/{isbn}", h.CreateByISBN).Methods(http.MethodPost)
	h.router.HandleFunc("/", h.Update).Methods(http.MethodPut)
	h.router.HandleFunc("/{id}", h.Delete).Methods(http.MethodDelete)
	h.router.HandleFunc("/{id}", h.GetByID).Methods(http.MethodGet)
	h.router.HandleFunc("/", h.GetAll).Methods(http.MethodGet)
}

func (h handler) Routes() *mux.Router {
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
	isbn := mux.Vars(r)["isbn"]

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
	id := mux.Vars(r)["id"]

	if err := h.service.Delete(ctx, id); err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
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

func (h *handler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]

	book, err := h.service.GetByID(ctx, id)
	if err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
