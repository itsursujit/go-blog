package middleware

import (
	"blog/app/config"
	"github.com/iris-contrib/middleware/jwt"
)

var (
	// JWT JWT Middleware
	JWT *jwt.Middleware
)

func initJWT() {
	JWT = jwt.New(jwt.Config{
		Extractor: jwt.FromAuthHeader,

		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			key := config.NewConfig(nil).GetValue("jwt_signature_key")
			return []byte(key), nil
		},

		SigningMethod: jwt.SigningMethodHS256,
	})
}
