package models

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type AuthParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (p *AuthParams) Validate() (err error) {
	if len(p.Username) == 0 {
		return fmt.Errorf("empty username")
	}
	if len(p.Password) > 8 {
		return fmt.Errorf("password should be longer than 8 chars")
	}
	return
}

type JWTClaims struct {
	AuthorId uint32
	jwt.RegisteredClaims
}
