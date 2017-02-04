package service

import (
	"gopkg.in/gin-gonic/gin.v1"
	"TruckMonitor-Backend/psql"
	"TruckMonitor-Backend/env"
)

type (
	AppService struct{}
	Resource struct {
		db *psql.DB
	}
)

func (app *AppService) Run() {
	// DataBase
	db := psql.Connection()
	defer db.Close()
	// TODO: [Debug]
	db.SchemeInitialization()

	resource := Resource{db: db}
	// Route
	api := gin.Default()
	api.NoRoute(resource.NotFoundError)
	v1 := api.Group("/api/v1")
	{
		v1.GET("/login", resource.SighIn)

		authorized := v1
		authorized.Use(resource.AuthRequiredMiddleware())
		{
			authorized.GET("/me", resource.GetMe)
		}
	}
	api.Run(env.ServiceHost)
}
