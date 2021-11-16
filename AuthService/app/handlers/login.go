package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	fmt.Println(c.Request.Body)
	fmt.Println("login success")
}
