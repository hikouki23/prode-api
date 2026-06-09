package main

import (
	"context"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
)

// Middleware para proteger las rutas del prode
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		sessionClaims, ok := clerk.SessionClaimsFromContext(ctx)
		if !ok {
			http.Error(w, "Rajá de acá, no estás logueado", http.StatusUnauthorized)
			return
		}

		userID := sessionClaims.Subject

		ctx = context.WithValue(ctx, "userID", userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
