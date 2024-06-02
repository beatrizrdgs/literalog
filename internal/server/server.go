package server

import (
	"log"
	"net/http"
	"text/template"

	"github.com/beatrizrdgs/literalog/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	router *chi.Mux
	port   string
	Handlers
}

type Handlers struct {
	BookHandler    *handlers.BookHandler
	LogbookHandler *handlers.LogbookHandler
	UserHandler    *handlers.UserHandler
}

func NewServer(port string, handlers *Handlers) *Server {
	r := chi.NewRouter()
	r.Use(
		middleware.RealIP,
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Heartbeat("/ping"),
	)

	s := &Server{
		router:   r,
		port:     port,
		Handlers: *handlers,
	}
	s.routes()

	return s
}

func (s *Server) routes() {
	s.router.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			templ := template.Must(template.ParseFiles("internal/assets/static/index.html"))
			templ.Execute(w, nil)
		})
	})

	s.router.Route("/books", func(r chi.Router) {
		// r.Post("/", s.BookHandler.Add)
		r.Get("/{id}", s.BookHandler.GetById)
	})

	// s.router.Route("/logbooks", func(r chi.Router) {
	// 	r.Post("/", s.LogbookHandler.Add)
	// })

	s.router.Route("/users", func(r chi.Router) {
		r.Post("/register", s.UserHandler.Register)
		r.Post("/login", s.UserHandler.Login)
	})
}

func (s *Server) Start() {
	log.Println("Server running on port", s.port)
	http.ListenAndServe(":"+s.port, s.router)
}
