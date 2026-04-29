package middleware

import (
	"fmt"
	"net/http"
	"urlShortner/auth"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		claims := &auth.Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return auth.JWTKEY, nil
		})
		if err != nil || !token.Valid {
			fmt.Println(err, token)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// access data

		fmt.Println(claims.Username)
		next.ServeHTTP(w, r)
	})
}
