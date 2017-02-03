package api

import "TruckMonitor-Backend/model"

/** Модель представления. Аккаунт пользователя. */
type Account struct {
	Id          int `json:"id"`
	Role        *Role `json:"role"`
	Email       string `json:"email"`
	Surname     string `json:"surname"`
	Name        string `json:"name"`
	Patronymic  string `json:"patronymic"`
	DateOfBirth string `json:"dateOfBirth"`
	Photo       string `json:"photo"`
	Phone       string `json:"phone"`
}

func NewAccount(account *model.Account, role *model.Role) *Account {
	roleVM := NewRole(role)
	return &Account{
		Id:          account.Id,
		Role:        roleVM,
		Email:       account.Email,
		Surname:     account.Surname,
		Name:        account.Name,
		Patronymic:  account.Patronymic,
		DateOfBirth: account.DateOfBirth,
		Photo:       account.Photo,
		Phone:       account.Phone}
}
