package middleware

import (
	"fmt"
	"net/http"

	"github.com/matidev200/guardaloya-backend/internal/login"
)



func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		tokenString := r.Header.Get("Authorization")
	
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Missin authorization header")
			return
		}
	
		tokenString = tokenString[len("Bearer "):]
	
		err := login.VerifyToken(tokenString)
	
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
			return
		}

		next.ServeHTTP(w,r)
	}
}