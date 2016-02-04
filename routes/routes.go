package routes

import (
	"github.com/exitcodezero/picloud/config"
	"github.com/exitcodezero/picloud/info"
	"github.com/exitcodezero/picloud/middleware"
	"github.com/exitcodezero/picloud/publish"
	"github.com/exitcodezero/picloud/subscribe"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Router setups all the API routes and middleware
func Router() *mux.Router {
	common := alice.New(middleware.Authentication, middleware.ClientName, middleware.RecoverHandler)
	authOnly := alice.New(middleware.Authentication, middleware.RecoverHandler)

	infoSocket := handlers.MethodHandler{
		"GET": authOnly.ThenFunc(info.SocketHandler),
	}

	subSocket := handlers.MethodHandler{
		"GET": common.ThenFunc(subscribe.Handler),
	}

	pubHTTP := handlers.MethodHandler{
		"GET":  common.ThenFunc(publish.HandlerSocket),
		"POST": common.ThenFunc(publish.HandlerHTTP),
	}

	router := mux.NewRouter()

	router.Handle("/publish", pubHTTP)
	router.Handle("/subscribe", subSocket)

	if config.EnableInfoSocket != "" {
		router.Handle("/info", infoSocket)
	}

	return router
}
