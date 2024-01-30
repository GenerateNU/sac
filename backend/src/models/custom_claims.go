package types

import "github.com/golang-jwt/jwt"

type CustomClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
}
