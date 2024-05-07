package usecases

import (
	"techiebutler/assessment/domain"
	"testing"
)

// MockEmployeeRepository is a mock implementation of domain.EmployeeRepository for testing purposes.
type MockEmployeeRepository struct {
	CreateEmployeeFunc  func(employee domain.Employee) (domain.Employee, error)
	GetEmployeeByIDFunc func(employee domain.Employee) (domain.Employee, error)
	UpdateEmployeeFunc  func(employee domain.Employee) error
	DeleteEmployeeFunc  func(employee domain.Employee) error
}

func (m *MockEmployeeRepository) CreateEmployee(employee domain.Employee) (domain.Employee, error) {
	if m.CreateEmployeeFunc != nil {
		return m.CreateEmployeeFunc(employee)
	}
	return domain.Employee{}, nil
}

func (m *MockEmployeeRepository) GetEmployeeByID(employee domain.Employee) (domain.Employee, error) {
	if m.GetEmployeeByIDFunc != nil {
		return m.GetEmployeeByIDFunc(employee)
	}
	return domain.Employee{}, nil
}

func (m *MockEmployeeRepository) UpdateEmployee(employee domain.Employee) error {
	if m.UpdateEmployeeFunc != nil {
		return m.UpdateEmployeeFunc(employee)
	}
	return nil
}

func (m *MockEmployeeRepository) DeleteEmployee(employee domain.Employee) error {
	if m.DeleteEmployeeFunc != nil {
		return m.DeleteEmployeeFunc(employee)
	}
	return nil
}

func TestCreateEmployee(t *testing.T) {
	mockRepo := &MockEmployeeRepository{
		CreateEmployeeFunc: func(employee domain.Employee) (domain.Employee, error) {
			return domain.Employee{ID: 1, Name: "John", Position: "Developer", Salary: 50000}, nil
		},
	}
	interactor := NewEmployeeInteractor(mockRepo)
	employee := domain.Employee{Name: "John", Position: "Developer", Salary: 50000}
	createdEmployee, err := interactor.CreateEmployee(employee)
	if err != nil {
		t.Errorf("CreateEmployee returned an unexpected error: %v", err)
	}
	if createdEmployee.ID != 1 {
		t.Errorf("Expected employee ID to be 1, got %d", createdEmployee.ID)
	}
}

func TestGetEmployeeByID(t *testing.T) {
	mockRepo := &MockEmployeeRepository{
		GetEmployeeByIDFunc: func(employee domain.Employee) (domain.Employee, error) {
			return domain.Employee{ID: 1, Name: "John", Position: "Developer", Salary: 50000}, nil
		},
	}
	interactor := NewEmployeeInteractor(mockRepo)
	employee := domain.Employee{ID: 1}
	retrievedEmployee, err := interactor.GetEmployeeByID(employee)
	if err != nil {
		t.Errorf("GetEmployeeByID returned an unexpected error: %v", err)
	}
	if retrievedEmployee.Name != "John" {
		t.Errorf("Expected employee name to be John, got %s", retrievedEmployee.Name)
	}
}

func TestUpdateEmployee(t *testing.T) {
	mockRepo := &MockEmployeeRepository{
		UpdateEmployeeFunc: func(employee domain.Employee) error {
			return nil
		},
	}
	interactor := NewEmployeeInteractor(mockRepo)
	employee := domain.Employee{ID: 1, Name: "John", Position: "Developer", Salary: 50000}
	err := interactor.UpdateEmployee(employee)
	if err != nil {
		t.Errorf("UpdateEmployee returned an unexpected error: %v", err)
	}
}

func TestDeleteEmployee(t *testing.T) {
	mockRepo := &MockEmployeeRepository{
		DeleteEmployeeFunc: func(employee domain.Employee) error {
			return nil
		},
	}
	interactor := NewEmployeeInteractor(mockRepo)
	employee := domain.Employee{ID: 1}
	err := interactor.DeleteEmployee(employee)
	if err != nil {
		t.Errorf("DeleteEmployee returned an unexpected error: %v", err)
	}
}
