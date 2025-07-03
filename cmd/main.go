package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to FIVEN, An open-source exam system.")
	fmt.Println("If you got any trouble, you can send it to https://github.com/Lparksi/fiven-front/issues")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}
