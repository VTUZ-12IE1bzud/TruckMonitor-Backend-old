package service

import (
	"net/http"
	"gopkg.in/gin-gonic/gin.v1"
	"TruckMonitor-Backend/model"
	"TruckMonitor-Backend/api"
)

type Account interface {
	AccountFindById(id int) (*model.Account, error)
	AccountFindByEmailPassword(email string, password string) (*model.Account, error)
}

func (resource *Resource) GetMe(context *gin.Context) {
	userId := context.MustGet("user")
	if userId == nil {
		context.JSON(http.StatusBadRequest, api.NewError("User not found"))
	} else {
		var account Account = resource.db
		var role Role = resource.db

		dataAccount, err := account.AccountFindById(userId.(int))
		if err != nil {
			context.JSON(http.StatusBadRequest, api.NewError(err.Error()))
			return
		}
		dataRole, err := role.RoleFindById(dataAccount.RoleId)
		if err != nil {
			context.JSON(http.StatusBadRequest, api.NewError(err.Error()))
			return
		}

		context.JSON(http.StatusOK, api.NewAccount(dataAccount, dataRole))
	}
}
