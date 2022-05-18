package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type Claims struct {
	UserType uint   `json:"user_type"`
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

var gSecret = []byte("deep_dark_fantastic")

const gExpireDuration = time.Hour * 2

func JwtInit() (err error) {
	tmp := viper.Get("secret")

	err = json.Unmarshal([]byte(tmp.(string)), &gSecret)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Print(gSecret)

	return
}

func JwtKeyGet(_ *jwt.Token) (i interface{}, err error) {
	return gSecret, nil
}

func TokenRelease(userType uint, userId uint, name string) (token string, err error) {
	// 创建一个我们自己的声明
	c := Claims{
		userType,
		userId, // 自定义字段
		name,   // 自定义字段
		jwt.StandardClaims{ // JWT规定的7个官方字段
			ExpiresAt: time.Now().Add(
				time.Duration(viper.GetInt("secret.expire")) * time.Hour).Unix(), // 过期时间
			Issuer: "vforum", // 签发人
		},
	}
	// 加密并获得完整的编码后的字符串token
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(gSecret)

	// 使用指定的secret签名并获得完整的编码后的字符串token
	return
}

func TokenParse(tokenString string) (claims *Claims, err error) {
	// 解析token
	var token *jwt.Token
	claims = new(Claims)
	token, err = jwt.ParseWithClaims(tokenString, claims, JwtKeyGet)
	if err != nil {
		return
	}
	if !token.Valid { // 校验token
		err = errors.New("invalid token")
	}
	return
}

func TokenRefresh(token string) (newToken string, err error) {

	// 从旧access token中解析出claims数据	解析出payload负载信息
	var claims Claims
	_, err = jwt.ParseWithClaims(token, &claims, JwtKeyGet)
	v, _ := err.(*jwt.ValidationError)

	// 当access token是过期错误 并且 refresh token没有过期时就创建一个新的access token
	if v.Errors == jwt.ValidationErrorExpired {
		return TokenRelease(claims.UserType, claims.UserId, claims.Name)
	}
	return
}
