package user_token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

// 一些常量
var (
	keyString string = "lxc"
)

type JWT struct {
	SigningKey []byte
}

func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(keyString), nil // 这是我的secret
	}
}

type CustomClaims struct {
	ID       uint   `json:"user_id"`
	NickName string `json:"nick_name"`
	jwt.RegisteredClaims
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	res, err := token.SignedString(j.SigningKey)
	fmt.Println("err:", err)
	return res, err
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, Secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenHandleError("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpiredError()
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenHandleError("token not active yet")
			} else {
				return nil, TokenHandleError("couldn't handle this token")
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenHandleError("couldn't handle this token")
}

func GenerateToken(claims CustomClaims) (string, error) {
	j := &JWT{[]byte(keyString)}
	token, err := j.CreateToken(claims)
	if err != nil {
		return "", TokenCreateError()
	}
	return token, nil
}

func CheckToken(tokenString string) (*CustomClaims, error) {
	j := &JWT{[]byte(keyString)}
	return j.ParseToken(tokenString)

}
