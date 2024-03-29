package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"Message": "pong"})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
