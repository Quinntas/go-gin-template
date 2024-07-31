package utils

import (
	"github.com/golang-jwt/jwt"
)

func signJWT[T jwt.Claims](claims T, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func parseJWT[T jwt.Claims](tokenString string, claims T, secret string) (T, error) {
	parsed, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return claims, err
	}
	if parsedClaims, ok := parsed.Claims.(T); ok && parsed.Valid {
		return parsedClaims, nil
	}
	return claims, jwt.NewValidationError("invalid token", jwt.ValidationErrorClaimsInvalid)
}
