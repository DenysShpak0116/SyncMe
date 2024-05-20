package main

import (
	"server/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitializeDBConnection()
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	router.Run(":8080")
}
