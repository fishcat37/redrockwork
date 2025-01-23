package services

import (
	"fmt"
	"redrockCommerce/config"
	"redrockCommerce/model"

	"github.com/golang-jwt/jwt/v5"
)

func ParseToken(tokenString string) (*model.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("解析token失败:%v", err)
	} else if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
