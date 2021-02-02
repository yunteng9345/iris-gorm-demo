package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"log"
	"print-chn/models"
)

const JwtKey = "percy"

func GetJWT() *jwtmiddleware.Middleware {
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		// 验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtKey), nil
		},
		// 加密方式
		SigningMethod: jwt.SigningMethodHS256,
		ErrorHandler: func(context iris.Context, err error) {
			fmt.Println("错误:", err)
			result := models.Result{Code: -1, Msg: "认证失败，请重新登录认证"}
			i, err := context.JSON(result)
			if err != nil {
				log.Println(i, err)
			}
		},
	})
	return jwtHandler
}
