package data

import (
	"../domain/model"
)

func FindUserByLoginPassword(email string, password string) (*model.User, error) {
	var user model.User
	err := connection.QueryRow("SELECT * FROM public.account WHERE email=$1 AND password=$2", email, password).
		Scan(&user.Id, &user.RoleId, &user.Email, &user.Password, &user.Surname, &user.Name, &user.Patronymic,
		&user.DateOfBirth, &user.Photo, &user.Phone)
	if err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

func FindUserById(id int) (*model.User, error) {
	var user model.User
	err := connection.QueryRow("SELECT * FROM public.account WHERE id=$1", id).
		Scan(&user.Id, &user.RoleId, &user.Email, &user.Password, &user.Surname, &user.Name, &user.Patronymic,
		&user.DateOfBirth, &user.Photo, &user.Phone)
	if err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}
