package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(" hello there ")
	r := gin.Default()
	fmt.Println(r)
	fmt.Println("hello there")
}
