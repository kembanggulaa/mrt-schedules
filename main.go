package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")

	var router = gin.Default()

	router.Run(":8080")
}
