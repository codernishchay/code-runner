package middleware

import (
	"authservice/app/models"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func VerifyToken(c *gin.Context) {
	fmt.Println("verify token just called")
	tokenstr, err := c.Request.Cookie("token")
	if err != nil {
		c.Writer.WriteHeader(http.StatusUnauthorized)
	}
	claims := &models.JwtClaims{}

	token, err := jwt.ParseWithClaims(tokenstr.Value, claims, func(token *jwt.Token) (interface{}, error) {
		jwtKey := []byte(os.Getenv("JWT_KEY"))
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.Writer.WriteHeader(http.StatusUnauthorized)
		}
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !token.Valid {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	c.Next()
	c.Writer.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}
