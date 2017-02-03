package psql

import (
	"TruckMonitor-Backend/model"
)

func (db *DB) AccountFindById(id int) (*model.Account, error) {
	var account model.Account
	row := db.DataBase.QueryRow("SELECT * FROM public.account WHERE id=$1", id)
	err := row.Scan(&account.Id, &account.RoleId, &account.Email, &account.Password, &account.Surname, &account.Name,
		&account.Patronymic, &account.DateOfBirth, &account.Photo, &account.Phone)
	if err != nil {
		return &account, err
	}
	return &account, nil
}
func (db *DB) AccountFindByEmailPassword(email string, password string) (*model.Account, error) {
	var account model.Account
	row := db.DataBase.QueryRow("SELECT * FROM public.account WHERE email=$1 AND password=$2", email, password)
	err := row.Scan(&account.Id, &account.RoleId, &account.Email, &account.Password, &account.Surname, &account.Name,
		&account.Patronymic, &account.DateOfBirth, &account.Photo, &account.Phone)
	if err != nil {
		return &account, err
	}
	return &account, nil
}
