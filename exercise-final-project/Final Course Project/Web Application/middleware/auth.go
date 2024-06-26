package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		
		cookie, err := ctx.Cookie("session_token")
		contentType := ctx.GetHeader("Content-Type")

        if err != nil {
            if err == http.ErrNoCookie && contentType == "application/json"{
                ctx.AbortWithStatus(http.StatusUnauthorized)
                return
            }
			ctx.AbortWithStatus(http.StatusSeeOther)
            return
        }

		claims := &model.Claims{}

		token, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})

		if token == nil {
            ctx.AbortWithStatus(http.StatusBadRequest)
            return
        }

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		
		ctx.Set("email", claims.Email)
		ctx.Next()
	})
}
