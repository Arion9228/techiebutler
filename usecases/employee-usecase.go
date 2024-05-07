package usecases

import (
	"log"
	"techiebutler/assessment/domain"
)

type EmployeeInteractor struct {
	EmployeeRepository domain.EmployeeRepository
}

func NewEmployeeInteractor(repository domain.EmployeeRepository) EmployeeInteractor {
	return EmployeeInteractor{repository}
}

func (interactor *EmployeeInteractor) CreateEmployee(employee domain.Employee) (domain.Employee, error) {
	employee, err := interactor.EmployeeRepository.CreateEmployee(employee)
	if err != nil {
		log.Println(err.Error())
		return domain.Employee{}, err
	}
	return employee, nil
}

func (interactor *EmployeeInteractor) GetEmployeeByID(employee domain.Employee) (domain.Employee, error) {
	employee, err := interactor.EmployeeRepository.GetEmployeeByID(employee)
	if err != nil {
		log.Println(err.Error())
		return domain.Employee{}, err
	}
	return employee, nil
}

func (interactor *EmployeeInteractor) UpdateEmployee(employee domain.Employee) error {
	err := interactor.EmployeeRepository.UpdateEmployee(employee)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (interactor *EmployeeInteractor) DeleteEmployee(employee domain.Employee) error {
	err := interactor.EmployeeRepository.DeleteEmployee(employee)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
