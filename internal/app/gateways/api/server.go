package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/literalog/library/internal/app/domain/authors"
	"github.com/literalog/library/internal/app/domain/book"
	"github.com/literalog/library/internal/app/domain/genre"
	"github.com/literalog/library/internal/app/domain/series"
	"github.com/literalog/library/internal/app/gateways/apis"
	"github.com/literalog/library/internal/app/gateways/database/mongodb"
)

type Server struct {
	port     string
	logLevel int
	router   *chi.Mux
}

func NewServer(port string) Server {
	storage, err := mongodb.NewMongoStorage()
	if err != nil {
		log.Fatal(err)
	}

	db := storage.Client.Database("library")

	authorRepository := mongodb.NewAuthorRepository(db.Collection("authors"))
	authorService := authors.NewService(authorRepository)
	authorHandler := authors.NewHandler(authorService)

	seriesRepository := mongodb.NewSeriesRepository(db.Collection("series"))
	seriesService := series.NewService(seriesRepository)
	seriesHandler := series.NewHandler(seriesService)

	genreRepository := mongodb.NewGenreRepository(db.Collection("genre"))
	genreService := genre.NewService(genreRepository)
	genreHandler := genre.NewHandler(genreService)

	isbnRepository, err := apis.NewGBooksAPI("", "https://www.googleapis.com/books/v1")
	if err != nil {
		log.Fatal(err)
	}

	bookRepository := mongodb.NewBookRepository(db.Collection("books"))
	bookService := book.NewService(bookRepository, isbnRepository, authorService, seriesService, genreService)
	bookHandler := book.NewHandler(bookService)

	router := chi.NewRouter()

	router.Mount("/authors", authorHandler.Routes())
	router.Mount("/series", seriesHandler.Routes())
	router.Mount("/genre", genreHandler.Routes())
	router.Mount("/books", bookHandler.Routes())

	return Server{
		port:     port,
		logLevel: 1,
		router:   router,
	}
}

func (s *Server) ServeHttp() error {
	log.Println("Server listening on", s.port)
	return http.ListenAndServe(s.port, s.router)
}
