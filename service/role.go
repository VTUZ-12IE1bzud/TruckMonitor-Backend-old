package service

import "TruckMonitor-Backend/model"

type Role interface {
	RoleFindById(id int) (*model.Role, error)
}
