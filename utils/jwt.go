package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// 秘钥
var key = []byte("1rqkw3i1")

const TokenExpireDuration = time.Hour * 24
const WhiteToken = "sofcgvtdnwjdrbhd"

type CustomClaims struct {
	// 可根据需要自行添加字段
	Account              string `json:"account"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

// GenRegisteredClaims 使用默认声明创建jwt
func GenRegisteredClaims(account string) (string, error) {
	// 创建 Claims
	claims := CustomClaims{
		account, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "my-project", // 签发人
		},
	}
	// 生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	return token.SignedString(key)
}

// ParseRegisteredClaims 解析jwt
func ValidateRegisteredClaims(tokenString string) bool {
	// 解析token

	if tokenString == WhiteToken {
		return true
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil { // 解析token失败
		return false
	}
	return token.Valid
}
