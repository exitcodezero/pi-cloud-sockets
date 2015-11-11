package middleware

import (
	"log"
	"net/http"
)

// RecoverHandler keeps the app from crashing and returns 500 errors
func RecoverHandler(n http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		n.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
