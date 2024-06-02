package initializer

import (
	"github.com/beatrizrdgs/literalog/internal/handlers"
	"github.com/beatrizrdgs/literalog/internal/models"
	"github.com/beatrizrdgs/literalog/internal/repository/mongodb"
	"github.com/beatrizrdgs/literalog/internal/services"
	"github.com/literalog/go-wise/wise"
)

var logbookSvc *services.LogbookService

func initLogbookHandler() *handlers.LogbookHandler {
	return handlers.NewLogbookHandler(logbookSvc)
}

func initLogbookRepo() (wise.MongoRepository[models.Logbook], error) {
	coll, err := mongodb.GetCollection("logbooks")
	if err != nil {
		return nil, err
	}

	repo, err := wise.NewMongoSimpleRepository[models.Logbook](coll)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func initLogbookSvc() (*services.LogbookService, error) {
	repo, err := initLogbookRepo()
	if err != nil {
		return nil, err
	}

	logbookSvc, err = services.NewLogbookService(repo)
	if err != nil {
		return nil, err
	}

	return logbookSvc, nil
}
