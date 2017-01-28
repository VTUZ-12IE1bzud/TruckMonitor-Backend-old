package util

import (
	"github.com/dgrijalva/jwt-go"
	"errors"
)

var tokenEncodeString = []byte("37FjfjU7vka80OU3r520Yy2T7h0p7hAM")

type TokenClaims struct {
	UserId int    `json:"id"`
	jwt.StandardClaims
}

/** Создать токен. */
func CreateToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "TruckMonitor",
		},
	})
	return token.SignedString(tokenEncodeString)
}

/** Разобрать токен. */
func ParseToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return tokenEncodeString, nil
	})
	if err != nil {
		return 0, err
	} else {
		if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
			return claims.UserId, nil
		} else {
			return 0, errors.New("Invalid token")
		}
	}
}
