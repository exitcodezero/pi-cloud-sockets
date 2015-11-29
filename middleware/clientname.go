package middleware

import (
	"github.com/gorilla/context"
	"net/http"
)

// ClientName parses the "X-API-Client-Name" header or "clientName" query parameter
func ClientName(n http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
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
