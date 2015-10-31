package middleware

import (
    "net/http"
    "app/config"
)

// Authentication validates the API Key passed in via the "X-API-Key" header or a query param "apiKey"
func Authentication(n http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        queryParams := r.URL.Query()

        apiKeyFromHeader := r.Header.Get("X-API-Key")
        apiKeyFromQuery := ""
        if len(queryParams["apiKey"]) > 0 {
            apiKeyFromQuery = queryParams["apiKey"][0]
        }

        if apiKeyFromHeader == config.APIKey || apiKeyFromQuery == config.APIKey {
            n.ServeHTTP(w, r)
            return
        }
        http.Error(w, "Invalid API Key", http.StatusUnauthorized)
    }
    return http.HandlerFunc(fn)
}
