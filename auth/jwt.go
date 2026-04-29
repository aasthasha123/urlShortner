package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTKEY = []byte("secret_key")

type Claims struct {
	Username string `json:"username"`

	jwt.RegisteredClaims
}

func GenerateToken(userName string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := Claims{Username: userName, RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(expirationTime)}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTKEY)
}
