package utils

import (
	"blog/pkg/jwt"
)

type Token interface {
	GenToken(uid int64) (string, error)
	ParseToken(token string) (any, error)
}

type JWTToken struct {
}

func NewJWTToken() *JWTToken {
	return &JWTToken{}
}

// GenToken 生成JWT
func (j *JWTToken) GenToken(uid int64) (string, error) {
	return jwt.GenToken(uid)
}

// ParseToken 解析JWT
func (j *JWTToken) ParseToken(token string) (any, error) {
	return jwt.ParseToken(token)
}
