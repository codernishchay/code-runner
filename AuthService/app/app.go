package app

import (
	"authservice/app/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func App(r *gin.Engine) {
	fmt.Println("hello app")
	r.Use(middleware.VerifyToken)
	Router(r)
}
