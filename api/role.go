package api

import "TruckMonitor-Backend/model"

/** Модель. Роль пользователя. */
type Role struct {
	Id   int `json:"id"`
	Name string `json:"name"`
}

func NewRole(role *model.Role) *Role {
	return &Role{Id: role.Id, Name: role.Name}
}
