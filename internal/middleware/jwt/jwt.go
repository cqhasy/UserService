package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaims struct {
	// 可根据需要自行添加字段
	Email                string `json:"email"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

const TokenExpireDuration = time.Hour * 24

// CustomSecret 用于签名的字符串
var CustomSecret = []byte("CCNU-EDU-LLM")

// GenToken 生成JWT
func GenToken(email string) (string, error) {
	// 创建一个我们自己的声明
	claims := CustomClaims{
		email, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "CCNU", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(CustomSecret)
}
