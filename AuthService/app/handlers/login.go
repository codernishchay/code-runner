package handlers

import (
	"authservice/app/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	fmt.Println(c.Request.Body)
	fmt.Println("login success")
	var credentials models.AuthModel
	fmt.Println(credentials)
	err := json.NewDecoder(c.Request.Body).Decode(&credentials)
	fmt.Println(err)
	json.NewEncoder(c.Writer).Encode("hi guys")
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &models.JwtClaims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
