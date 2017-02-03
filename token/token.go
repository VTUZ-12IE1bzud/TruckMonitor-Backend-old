package token

import (
	"github.com/dgrijalva/jwt-go"
	"TruckMonitor-Backend/config"
	"TruckMonitor-Backend/model"
	"errors"
)

type (
	/** Пользователь. */
	UserToken struct {
		Account *model.Account
	}

	/** Раскодированный токен. */
	DataToken struct {
		UserId    int
		UserEmail string
	}

	/** Структура токена. */
	claims struct {
		UserId    int    `json:"userId"`
		UserEmail string    `json:"userEmail"`
		jwt.StandardClaims
	}
)

/** Ключ шифрования токена. */
var key []byte

func init() {
	key = []byte(config.GetConfiguration().TokenKey)
}

/** Создать токен. */
func (user *UserToken) Create() (string, error) {
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		UserId:         user.Account.Id,
		UserEmail:      user.Account.Email,
		StandardClaims: jwt.StandardClaims{Issuer: "TruckMonitor"},
	})
	newTokenString, err := newToken.SignedString(key)
	return newTokenString, err
}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJUcnVja01vbml0b3IifQ.nCuOoKkNVhQjsHjktCf4GNpNiqjw-ThPGu1iG_OwZyc
/** Разобрать токен. */
func Parse(tokenString string) (*DataToken, error) {
	userToken, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return &DataToken{}, err
	} else {
		if claims, ok := userToken.Claims.(*claims); ok && userToken.Valid {
			return &DataToken{UserId: claims.UserId, UserEmail: claims.UserEmail}, nil
		} else {
			return &DataToken{}, errors.New("Invalid token")
		}
	}
}
