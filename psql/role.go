package psql

import "TruckMonitor-Backend/model"

func (db *DB) RoleFindById(id int) (*model.Role, error) {
	var role model.Role
	row := db.DataBase.QueryRow("SELECT * FROM public.role WHERE id=$1", id)
	err := row.Scan(&role.Id, &role.Name)
	if err != nil {
		return &role, err
	}
	return &role, err
}
