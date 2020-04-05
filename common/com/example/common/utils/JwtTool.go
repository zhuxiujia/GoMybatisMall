package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func Sign(data jwt.MapClaims, secret string) string { //生成token,算法hs256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	access_token, _ := token.SignedString([]byte(secret))
	fmt.Println("access_token=" + access_token)
	return access_token
}

func VerifyToken(input string, secret string) (jwt.Claims, error) {
	token, err := jwt.Parse(input, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if jwt.SigningMethodHS256.Alg() != token.Header["alg"] {
		return nil, errors.New("header err!")
	}
	fmt.Println("verify pass!!!")
	return token.Claims, nil
}
