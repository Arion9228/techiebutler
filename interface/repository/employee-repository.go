package repository

import (
	"techiebutler/assessment/domain"
)

// DBHandler defines the interface for database operations.
type DBHandler interface {
	CreateEmployee(domain.Employee) (domain.Employee, error) // CreateEmployee creates a new employee.
	GetEmployeeByID(domain.Employee) (domain.Employee, error) // GetEmployeeByID retrieves an employee by ID.
	UpdateEmployee(domain.Employee) error // UpdateEmployee updates an existing employee.
	DeleteEmployee(domain.Employee) error // DeleteEmployee deletes an employee.
}

// EmployeeRepo represents a repository for employee-related operations.
type EmployeeRepo struct {
	handler DBHandler // handler is the instance of DBHandler interface.
}

// NewEmployeeRepo creates a new instance of EmployeeRepo with the given handler.
func NewEmployeeRepo(handler DBHandler) EmployeeRepo {
	return EmployeeRepo{handler}
}

// CreateEmployee creates a new employee.
func (repo EmployeeRepo) CreateEmployee(employee domain.Employee) (domain.Employee, error) {
	employee, err := repo.handler.CreateEmployee(employee)
	if err != nil {
		return domain.Employee{}, err
	}
	return employee, nil
}

// GetEmployeeByID retrieves an employee by ID.
func (repo EmployeeRepo) GetEmployeeByID(employee domain.Employee) (domain.Employee, error) {
	employee, err := repo.handler.GetEmployeeByID(employee)
	if err != nil {
		return domain.Employee{}, err
	}
	return employee, nil
}

// UpdateEmployee updates an existing employee.
func (repo EmployeeRepo) UpdateEmployee(employee domain.Employee) error {
	err := repo.handler.UpdateEmployee(employee)
	if err != nil {
		return err
	}
	return nil
}

// DeleteEmployee deletes an employee.
func (repo EmployeeRepo) DeleteEmployee(employee domain.Employee) error {
	err := repo.handler.DeleteEmployee(employee)
	if err != nil {
		return err
	}
	return nil
}
