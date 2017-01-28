package controller

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"../viewmodel"
	"../../util"
	"../../data"
)

/** Авторизация пользователя. */
func SighIn(context *gin.Context) {
	// Parameter's
	var email = context.Query("email")
	var password = context.Query("password")

	// Validation
	var errors = []viewmodel.ErrorViewModel{}
	err := util.ValidateLogin(email)
	if err != nil {
		errors = append(errors, viewmodel.ErrorViewModel{Message: err.Error()})
	}
	err = util.ValidatePassword(password)
	if err != nil {
		errors = append(errors, viewmodel.ErrorViewModel{Message: err.Error()})
	}

	if len(errors) == 0 {
		// Валидация пройдена, поиск в БД
		account, err := data.FindUserByLoginPassword(email, password)

		if err != nil {
			errors = append(errors, viewmodel.ErrorViewModel{Message: err.Error()})
		} else if account == nil {
			// Пользователь не найден в БД
			errors = append(errors, viewmodel.ErrorViewModel{Message: "User not found"})
		} else if err == nil && account != nil {
			// Пользователь найден, генерируем токен
			var tokenString string
			tokenString, err = util.CreateToken(account.Id)
			if err == nil {
				context.JSON(http.StatusOK, viewmodel.TokenViewModel{Token: tokenString})
				return
			}
		}
	}
	if len(errors) > 0 {
		context.JSON(http.StatusBadRequest, errors)
	}
}
