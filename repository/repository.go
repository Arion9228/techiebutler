package repository

import "techiebutler/assessment/domain"

type DBHandler interface {
	CreateEmployee(employee domain.Employee) error
	GetEmployeeByID(id int) (domain.Employee, error)
	UpdateEmployee(employee domain.Employee) error
	DeleteEmployee(id int) error
}
