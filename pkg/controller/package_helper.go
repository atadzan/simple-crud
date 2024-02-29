package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/atadzan/simple-crud/pkg/models"
	"github.com/gofiber/fiber/v2"
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
func getAuthorIDFromCtx(c *fiber.Ctx) (authorId uint32) {
	val := c.Locals("authorId")
	if val != nil {
		authorId = val.(uint32)
	}
	return
}

func getPaginationParams(c *fiber.Ctx) (resp models.PaginationParams) {
	limit, err := strconv.ParseUint(c.Query("limit"), 10, 64)
	if err != nil || limit == 0 {
		resp.Limit = 15
	} else {
		resp.Limit = limit
	}

	page, err := strconv.ParseUint(c.Query("page"), 10, 64)
	if err != nil || page == 0 {
		resp.Offset = 0
	} else {
		resp.Offset = page * resp.Limit
	}
	return
}

func getFilterParams(c *fiber.Ctx) (resp models.BookFilter) {
	authorId, err := strconv.ParseUint(c.Query("authorId"), 10, 64)
	if err != nil {
		authorId = 0
	}
	resp.AuthorId = uint32(authorId)
	genreId, err := strconv.ParseUint(c.Query("genreId"), 10, 64)
	if err != nil {
		authorId = 0
	}
	resp.GenreId = uint32(genreId)
	return
}
