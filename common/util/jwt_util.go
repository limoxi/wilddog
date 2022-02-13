package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/limoxi/ghost"
	"time"
)

const salt = "kL9jasA2ksd1aAkd6ak3s"

func EncodeJwtToken(data ghost.Map, exp time.Duration) string{
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": float64(time.Now().Add(exp).Unix()),
		"data": data,
	})
	ts, err := token.SignedString([]byte(salt))
	if err != nil{
		panic(err)
	}
	return ts
}

func DecodeJwtToken(tokenStr string) ghost.Map{
	token, err := jwt.Parse(tokenStr, func (token *jwt.Token) (interface{}, error){
		return []byte(salt), nil
	})
	if token.Valid{
		return token.Claims.(jwt.MapClaims)["data"].(ghost.Map)
	}else{
		if jve, ok := err.(*jwt.ValidationError); ok{
			if jve.Errors & jwt.ValidationErrorExpired == 1{
				panic(ghost.NewBusinessError("token已过期"))
			}
		}
		panic(err)
	}
}