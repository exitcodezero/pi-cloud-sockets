package middleware

import (
	"github.com/gorilla/context"
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

		clientName := r.Header.Get("X-API-Client-Name")
		if clientName == "" {
			clientName = r.URL.Query().Get("clientName")
		}

		if clientName == "" {
			http.Error(w, "Invalid client name", http.StatusUnauthorized)
			return
		}
		context.Set(r, "ClientName", clientName)

		n.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
