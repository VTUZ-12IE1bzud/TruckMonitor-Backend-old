package presentation

import (
	"gopkg.in/gin-gonic/gin.v1"
	"./controller"
	"../util"
	"../data"
	"net/http"
)

const UserContextKey  = "user"

func StartRoute() {
	restRouter := gin.Default()

	restRouter.NoRoute(controller.NotFoundError)

	v1 := restRouter.Group("/api/v1")
	{
		v1.GET("/login", controller.SighIn)

		authorized := v1
		authorized.Use(AuthRequiredMiddleware())
		{
			authorized.GET("/me", controller.GetMe)
		}
	}

	restRouter.Run(":8080")
}

func AuthRequiredMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Validate Token
		token := context.Request.Header.Get("X-Auth-Token")
		userId, err := util.ParseToken(token)
		if err != nil {
			context.AbortWithStatus(http.StatusUnauthorized)
		} else {
			user, err := data.FindUserById(userId)
			if err != nil || user == nil {
				// Пользователь не найден.
				context.AbortWithStatus(http.StatusUnauthorized)
			} else {
				context.Set(UserContextKey, user)
				context.Next()
			}
		}
	}
}