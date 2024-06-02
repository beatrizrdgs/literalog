package initializer

import "github.com/beatrizrdgs/literalog/internal/server"

func InitHandlers() *server.Handlers {
	return &server.Handlers{
		BookHandler:    initBookHandler(),
		LogbookHandler: initLogbookHandler(),
		UserHandler:    initUserHandler(),
	}
}
