package controller

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"../viewmodel"
)

func GetMe(context *gin.Context) {
	user := context.MustGet("user")
	if user == nil {
		context.JSON(http.StatusBadRequest, viewmodel.ErrorViewModel{Message: "User not found"})
	} else {
		context.JSON(http.StatusOK, user)
	}
}
