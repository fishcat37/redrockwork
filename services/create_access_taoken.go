package services

import (
	"redrockCommerce/config"
	"redrockCommerce/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte(config.JwtSecretKey)

func CreateAccessToken(user model.User) (string, error) {
	claims := model.CustomClaims{
		Username: user.Username,
		Password: user.Password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24 小时后过期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "redrock",
			Subject:   "myjwt",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}
