package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"zzidun.tech/vforum0/model"
)

var g_jwt_key = []byte("deep_dark_fantastic")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func Token_Release(user *model.User) (string, error) {
	expiration_time := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration_time.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "zzidun.tech",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	authorization, err := token.SignedString(g_jwt_key)

	if err != nil {
		return "", err
	}

	return authorization, nil
}

func Token_Parse(authorization string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(authorization, claims, func(token *jwt.Token) (i interface{}, err error) {
		return g_jwt_key, nil
	})

	return token, claims, err
}
