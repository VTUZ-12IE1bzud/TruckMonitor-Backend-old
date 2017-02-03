package service

import (
	"net/http"
	"gopkg.in/gin-gonic/gin.v1"
	"TruckMonitor-Backend/api"
)

func (resource *Resource) NotFoundError(context *gin.Context) {
	context.JSON(http.StatusNotFound, api.NewError("Not Found"))
}
