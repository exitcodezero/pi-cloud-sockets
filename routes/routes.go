package routes

import (
	"github.com/exitcodezero/picloud/config"
	"github.com/exitcodezero/picloud/info"
	"github.com/exitcodezero/picloud/middleware"
	"github.com/exitcodezero/picloud/socket"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Router setups all the API routes and middleware
func Router() *mux.Router {
	common := alice.New(middleware.Authentication, middleware.RecoverHandler)
	noAuth := alice.New(middleware.RecoverHandler)

	infoSocket := handlers.MethodHandler{
		"GET": noAuth.ThenFunc(info.SocketHandler),
	}

	pubSubSocket := handlers.MethodHandler{
		"GET": common.ThenFunc(socket.Handler),
	}

	router := mux.NewRouter()

	router.Handle("/connect", pubSubSocket)

	if config.EnableInfoSocket != "" {
		router.Handle("/socket/info", infoSocket)
	}

	return router
}
