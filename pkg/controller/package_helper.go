package controller

import (
	"fmt"
	"time"

	"github.com/atadzan/simple-crud/pkg/models"
	"github.com/golang-jwt/jwt/v5"
)

func generateAccessToken(authorId uint32, jwtSigningKey string) string {
	claims := &models.JWTClaims{
		AuthorId: authorId,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now().UTC()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	resultToken, err := accessToken.SignedString([]byte(jwtSigningKey))
	if err != nil {
		return ""
	}
	return resultToken
}

func parseAccessToken(accessToken, jwtSigningKey string) (uint32, error) {
	var claims models.JWTClaims
	resultToken, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSigningKey), nil
	})
	switch {
	case !resultToken.Valid:
		return 0, fmt.Errorf("invalid token")
	case err != nil:
		return 0, err
	}
	return claims.AuthorId, err
}
