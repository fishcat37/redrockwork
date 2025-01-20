package services

import (
	"redrockCommerce/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateRefreshToken(user model.User) (string, error) {
	claims := model.CustomClaims{
		Username: user.Username,
		Password: user.Password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "redrock",
			Subject:   "myjwt",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}
