package app

import (
	"authservice/app/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	fmt.Println("Router")
	r.POST("/login", handlers.Login)
}
