package usecases

import (
	"log"
	"techiebutler/assessment/domain"
)

// EmployeeInteractor represents the interactor for employee-related use cases.
type EmployeeInteractor struct {
	EmployeeRepository domain.EmployeeRepository // EmployeeRepository is the repository for employee operations.
}

// NewEmployeeInteractor creates a new instance of EmployeeInteractor with the given repository.
func NewEmployeeInteractor(repository domain.EmployeeRepository) EmployeeInteractor {
	return EmployeeInteractor{repository}
}

// CreateEmployee creates a new employee.
func (interactor *EmployeeInteractor) CreateEmployee(employee domain.Employee) (domain.Employee, error) {
	employee, err := interactor.EmployeeRepository.CreateEmployee(employee)
	if err != nil {
		log.Println(err.Error())
		return domain.Employee{}, err
	}
	return employee, nil
}

// GetEmployeeByID retrieves an employee by ID.
func (interactor *EmployeeInteractor) GetEmployeeByID(employee domain.Employee) (domain.Employee, error) {
	employee, err := interactor.EmployeeRepository.GetEmployeeByID(employee)
	if err != nil {
		log.Println(err.Error())
		return domain.Employee{}, err
	}
	return employee, nil
}

// UpdateEmployee updates an existing employee.
func (interactor *EmployeeInteractor) UpdateEmployee(employee domain.Employee) error {
	err := interactor.EmployeeRepository.UpdateEmployee(employee)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

// DeleteEmployee deletes an employee.
func (interactor *EmployeeInteractor) DeleteEmployee(employee domain.Employee) error {
	err := interactor.EmployeeRepository.DeleteEmployee(employee)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
