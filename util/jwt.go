package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"zzidun.tech/vforum0/controller"
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

func TokenReleaseAccess(userType uint, userId uint, name string) (aToken, rToken string, err error) {
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
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(gSecret)

	// refresh token 不需要存任何自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(), // 过期时间
		Issuer:    "vforum",                                // 签发人
	}).SignedString(gSecret)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return
}

func TokenRelease(userType uint, userId uint, name string) (token string, err error) {
	// 创建一个我们自己的声明
	c := Claims{
		userType,
		userId, // 自定义字段
		name,   // 自定义字段
		jwt.StandardClaims{ // JWT规定的7个官方字段
			ExpiresAt: time.Now().Add(gExpireDuration).Unix(), // 过期时间
			Issuer:    "vforum",                               // 签发人
		},
	}
	// 加密并获得完整的编码后的字符串token
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(gSecret)

	// refresh token 不需要存任何自定义数据
	//rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
	//	ExpiresAt: time.Now().Add(time.Second * 30).Unix(), // 过期时间
	//	Issuer:    "bluebell",                              // 签发人
	//}).SignedString(mySecret)	// 使用指定的secret签名并获得完整的编码后的字符串token
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

func TokenRefresh(aToken, rToken string) (newAToken, newRToken string, err error) {
	// refresh token无效直接返回
	if _, err = jwt.Parse(rToken, JwtKeyGet); err != nil {
		return
	}

	// 从旧access token中解析出claims数据	解析出payload负载信息
	var claims Claims
	_, err = jwt.ParseWithClaims(aToken, &claims, JwtKeyGet)
	v, _ := err.(*jwt.ValidationError)

	// 当access token是过期错误 并且 refresh token没有过期时就创建一个新的access token
	if v.Errors == jwt.ValidationErrorExpired {
		return TokenReleaseAccess(claims.UserType, claims.UserId, claims.Name)
	}
	return
}

func JwtAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseErrorWithMsg(c, controller.CodeInvalidToken, "请求头缺少Auth Token")
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseErrorWithMsg(c, controller.CodeInvalidToken, "Token格式不对")
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := TokenParse(parts[1])
		if err != nil {
			fmt.Println(err)
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set("userType", mc.UserType)
		c.Set("userId", mc.UserId)
		c.Next() // 后续的处理函数可以用过c.Get(ContextUserIDKey)来获取当前请求的用户信息
	}
}

// func Token_Release(user *model.User) (string, error) {
// 	expiration_time := time.Now().Add(24 * time.Hour)
// 	claims := &Claims{
// 		UserId: user.ID,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expiration_time.Unix(),
// 			IssuedAt:  time.Now().Unix(),
// 			Issuer:    "zzidun.tech",
// 			Subject:   "user token",
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	authorization, err := token.SignedString(gJwtKey)

// 	if err != nil {
// 		return "", err
// 	}

// 	return authorization, nil
// }

// func Token_Parse(authorization string) (*jwt.Token, *Claims, error) {
// 	claims := &Claims{}

// 	token, err := jwt.ParseWithClaims(authorization, claims, func(token *jwt.Token) (i interface{}, err error) {
// 		return gJwtKey, nil
// 	})

// 	return token, claims, err
// }
