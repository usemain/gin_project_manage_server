package utils

import (
	"errors"
	"gin_project_manage_server/shares/config"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type MyClaims struct {
	Uid   string `json:"uid"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func MakeAuthorization(Uid, email string) (tokenString string, err error) {
	claim := MyClaims{
		Uid:   Uid,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(720 * time.Hour * time.Duration(1))), // 过期时间30天
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                         // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                         // 生效时间
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err = token.SignedString([]byte(config.SECRET))
	return tokenString, err
}

func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET), nil // secret
	}
}

func ParseAuthorization(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, Secret())
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("未知令牌")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("令牌已过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("令牌尚未激活")
			} else {
				return nil, errors.New("无法处理此令牌")
			}
		}
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无法处理此令牌")
}
