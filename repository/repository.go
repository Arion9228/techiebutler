package repository

import "techiebutler/assessment/domain"

type DBHandler interface {
	CreateEmployee(domain.Employee) error
	GetEmployeeByID(domain.Employee) (domain.Employee, error)
	UpdateEmployee(domain.Employee) error
	DeleteEmployee(domain.Employee) error
}
