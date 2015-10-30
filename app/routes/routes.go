package routes


import (
	"github.com/justinas/alice"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
    "app/middleware"
    "app/socket"
)

// Router setups all the API routes and middleware
func Router() *mux.Router  {
    common := alice.New(middleware.Authentication, middleware.RecoverHandler)

	socketHandlers := handlers.MethodHandler{
		"GET": common.ThenFunc(socket.Handler),
	}

    router := mux.NewRouter()
	router.Handle("/connect", socketHandlers)

    return router
}
