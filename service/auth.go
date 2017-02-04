package service

import (
	"net/http"
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/go-playground/validator.v9"
	"TruckMonitor-Backend/api"
	"TruckMonitor-Backend/token"
)

type SignInValid struct {
	Email    string 	`validate:"required,email"`
	Password string     `validate:"required"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (resource *Resource) AuthRequiredMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Validate Token
		tokenString := context.Request.Header.Get("X-Auth-Token")
		tokenData, err := token.Parse(tokenString)
		if err != nil {
			context.AbortWithStatus(http.StatusUnauthorized)
		} else {
			context.Set("user", tokenData.UserId)
			context.Next()
		}
	}
}

func (resource *Resource) SighIn(context *gin.Context) {
	// Parameter's
	var email = context.Query("email")
	var password = context.Query("password")

	// Validation
	valid := SignInValid{Email: email, Password: password}
	if err := validate.Struct(valid); err != nil {
		context.JSON(http.StatusBadRequest, api.NewError(err.Error()))
		return
	}

	var account Account = resource.db

	accountData, err := account.AccountFindByEmailPassword(email, password)
	if err != nil {
		context.JSON(http.StatusBadRequest, api.NewError(err.Error()))
		return
	}

	accountToken := token.UserToken{Account: accountData}
	tokenString, err := accountToken.Create()
	if err != nil {
		context.JSON(http.StatusBadRequest, api.NewError(err.Error()))
		return
	}

	context.JSON(http.StatusOK, api.NewToken(tokenString))
}
