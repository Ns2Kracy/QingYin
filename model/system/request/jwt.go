package request

import (
	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	UserId   uint
	Username string
}
