package api

import "github.com/golang-jwt/jwt/v5"

type jwtCustomClaims struct {
	Username string `json:"username"`
	ID       uint   `json:"id"`
	jwt.RegisteredClaims
}
