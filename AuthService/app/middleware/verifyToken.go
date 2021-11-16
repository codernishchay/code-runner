package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func VerifyToken(c *gin.Context) {
	fmt.Println(c.Request)
}
