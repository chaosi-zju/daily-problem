package model

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var (
	TokenExpired     = errors.New("Token is expired. ")
	TokenNotValidYet = errors.New("Token not active yet. ")
	TokenMalformed   = errors.New("That's not even a token. ")
	TokenInvalid     = errors.New("Couldn't handle this token. ")
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	UserID uint   `json:"userId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// NewJWT: 新建一个jwt实例
func NewJWT() *JWT {
	signKey := viper.GetString("JWT-SignKey")
	return &JWT{
		[]byte(signKey),
	}
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if e, ok := err.(*jwt.ValidationError); ok {
			if e.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if e.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if e.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			}
		}
		return nil, TokenInvalid
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}
