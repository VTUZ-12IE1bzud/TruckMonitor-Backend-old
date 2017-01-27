package router

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func Start() {
	restRouter := gin.Default()

	restRouter.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})

	restRouter.Run(":8080")
}