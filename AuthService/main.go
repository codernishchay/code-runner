package main

import (
	"authservice/app"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(" hello there ")
	r := gin.Default()
	fmt.Println(r)
	app.App(r)
	fmt.Println("hello there")
	r.Run()
}
