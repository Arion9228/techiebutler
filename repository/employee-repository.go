package repository

import (
	"techiebutler/assessment/domain"
)

type EmployeeRepo struct {
	handler DBHandler
}

func NewEmployeeRepo(handler DBHandler) EmployeeRepo {
	return EmployeeRepo{handler}
}

func (repo EmployeeRepo) CreateEmployee(employee domain.Employee) error {
	err := repo.handler.CreateEmployee(employee)
	if err != nil {
		return err
	}
	return nil
}

func (repo EmployeeRepo) GetEmployeeByID(employee domain.Employee) (domain.Employee, error) {
	employee, err := repo.handler.GetEmployeeByID(employee)
	if err != nil {
		return domain.Employee{}, err
	}
	return employee, nil
}

func (repo EmployeeRepo) UpdateEmployee(employee domain.Employee) error {
	err := repo.handler.UpdateEmployee(employee)
	if err != nil {
		return err
	}
	return nil
}

func (repo EmployeeRepo) DeleteEmployee(employee domain.Employee) error {
	err := repo.handler.DeleteEmployee(employee)
	if err != nil {
		return err
	}
	return nil
}
