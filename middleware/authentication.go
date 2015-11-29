package middleware

import (
	"github.com/exitcodezero/picloud/config"
	"net/http"
)

// Authentication validates the API Key passed in via the "X-API-Key" header or a query param "apiKey"
func Authentication(n http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			apiKey = r.URL.Query().Get("apiKey")
		}

		if apiKey != config.APIKey {
			http.Error(w, "Invalid API Key", http.StatusUnauthorized)
			return
		}

		n.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
