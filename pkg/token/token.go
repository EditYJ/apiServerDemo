package token

import (
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

var (
	ErrMissingHeader = errors.New("头部Key==>>`Authorization`的长度是0")
)

type Context struct {
	ID       uint64
	Username string
}

// 验证密钥格式
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// 确保排除了`alg`
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}

// 使用指定的密钥验证令牌，如果令牌有效，则返回上下文。
func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}

	// 解析Token
	token, err := jwt.Parse(tokenString, secretFunc(secret))

	// 处理错误
	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID = uint64(claims["id"].(float64))
		ctx.Username = claims["username"].(string)
		return ctx, nil
	} else {
		return ctx, err
	}
}

// 从请求头获取token并且调用[parse]函数解析他
func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	// 从配置文件中加载jwt
	secret := viper.GetString("jwt_secret")

	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}

	var t string
	// 解析header得到token的部分
	fmt.Sscanf(header, "Bearer %s", &t)
	return Parse(t, secret)
}

// 使用指定的密钥进行上下文签名
func Sign(ctx *gin.Context, c Context, secret string) (tokenString string, err error) {
	// 如果没有密钥规定从Gin的配置文件中加载jwt密钥
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})
	tokenString, err = token.SignedString([]byte(secret))
	return
}
