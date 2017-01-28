package controller

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"../viewmodel"
)

func NotFoundError(context *gin.Context) {
	context.JSON(http.StatusNotFound, viewmodel.ErrorViewModel{"Not Found"})
}
