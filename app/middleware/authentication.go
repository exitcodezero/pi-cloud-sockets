package middleware

import (
    "net/http"
    "app/config"
)

// Authentication validates the API Key passed in via the X-API-Key header
func Authentication(n http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        apiKey := r.Header.Get("X-API-Key")
        if apiKey != config.APIKey {
            message := "Invalid API Key"
            http.Error(w, message, http.StatusUnauthorized)
        }
        n.ServeHTTP(w, r)
    }
    return http.HandlerFunc(fn)
}
