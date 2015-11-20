package routes

import (
	"github.com/exitcodezero/picloud/middleware"
	"github.com/exitcodezero/picloud/socket"
	"github.com/exitcodezero/picloud/info"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Router setups all the API routes and middleware
func Router() *mux.Router {
	common := alice.New(middleware.Authentication, middleware.RecoverHandler)

	infoHandlers := handlers.MethodHandler{
		"GET": common.ThenFunc(info.Handler),
	}

	socketHandlers := handlers.MethodHandler{
		"GET": common.ThenFunc(socket.Handler),
	}

	router := mux.NewRouter()
	router.Handle("/connect", socketHandlers)
	router.Handle("/info", infoHandlers)

	return router
}
