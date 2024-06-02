package initializer

import (
	"github.com/beatrizrdgs/literalog/internal/handlers"
	"github.com/beatrizrdgs/literalog/internal/models"
	"github.com/beatrizrdgs/literalog/internal/repository/mongodb"
	"github.com/beatrizrdgs/literalog/internal/services"
	"github.com/literalog/go-wise/wise"
)

var bookSvc *services.BookService

func initBookHandler() *handlers.BookHandler {
	return handlers.NewBookHandler(bookSvc)
}

func initBookRepo() (wise.MongoRepository[models.Book], error) {
	coll, err := mongodb.GetCollection("books")
	if err != nil {
		return nil, err
	}

	repo, err := wise.NewMongoSimpleRepository[models.Book](coll)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func initBookSvc() (*services.BookService, error) {
	repo, err := initBookRepo()
	if err != nil {
		return nil, err
	}

	bookSvc, err = services.NewBookService(repo)
	if err != nil {
		return nil, err
	}

	return bookSvc, nil
}
