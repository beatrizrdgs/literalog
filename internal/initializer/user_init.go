package initializer

import (
	"github.com/beatrizrdgs/literalog/internal/handlers"
	"github.com/beatrizrdgs/literalog/internal/models"
	"github.com/beatrizrdgs/literalog/internal/repository/mongodb"
	"github.com/beatrizrdgs/literalog/internal/services"
	"github.com/literalog/go-wise/wise"
)

var userSvc *services.UserService

func initUserHandler() *handlers.UserHandler {
	return handlers.NewUserHandler(userSvc)
}

func initUserRepo() (wise.MongoRepository[models.User], error) {
	coll, err := mongodb.GetCollection("Users")
	if err != nil {
		return nil, err
	}

	repo, err := wise.NewMongoSimpleRepository[models.User](coll)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func initUserSvc() (*services.UserService, error) {
	repo, err := initUserRepo()
	if err != nil {
		return nil, err
	}

	userSvc, err = services.NewUserService(repo)
	if err != nil {
		return nil, err
	}

	return userSvc, nil
}
