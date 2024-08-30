package middleware

import (
	"context"
	"log"
	"net/http"
	"v3/config"
)

type contextKey string

const AuthenticatedClientKey contextKey = "authenticatedClient"

func AuthMiddleware(t any) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Retrieve the OAuth token from the session or context
			if r.URL.Path == "/login" || r.URL.Path == "/callback" {
				next.ServeHTTP(w, r)
				return
			}
			session, _ := config.Store.Get(r, "auth-session")
			token, ok := session.Values["token"]
			// Implement this function to retrieve the token
			if !ok || token == nil {
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}

			// Create an authenticated client

			// Store the authenticated client in the context
			log.Println("eggggggg")
			ctx := context.WithValue(r.Context(), AuthenticatedClientKey, "egg")

			f, ok := ctx.Value(AuthenticatedClientKey).(string)
			log.Print("ayo", f, ok)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
