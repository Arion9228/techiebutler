package repository

import (
	"techiebutler/assessment/domain"
	"testing"
)

// MockDBHandler is a mock implementation of DBHandler for testing purposes.
type MockDBHandler struct {
	CreateEmployeeFunc  func(domain.Employee) (domain.Employee, error)
	GetEmployeeByIDFunc func(domain.Employee) (domain.Employee, error)
	UpdateEmployeeFunc  func(domain.Employee) error
	DeleteEmployeeFunc  func(domain.Employee) error
}

func (m *MockDBHandler) CreateEmployee(employee domain.Employee) (domain.Employee, error) {
	if m.CreateEmployeeFunc != nil {
		return m.CreateEmployeeFunc(employee)
	}
	return domain.Employee{}, nil
}

func (m *MockDBHandler) GetEmployeeByID(employee domain.Employee) (domain.Employee, error) {
	if m.GetEmployeeByIDFunc != nil {
		return m.GetEmployeeByIDFunc(employee)
	}
	return domain.Employee{}, nil
}

func (m *MockDBHandler) UpdateEmployee(employee domain.Employee) error {
	if m.UpdateEmployeeFunc != nil {
		return m.UpdateEmployeeFunc(employee)
	}
	return nil
}

func (m *MockDBHandler) DeleteEmployee(employee domain.Employee) error {
	if m.DeleteEmployeeFunc != nil {
		return m.DeleteEmployeeFunc(employee)
	}
	return nil
}

func TestCreateEmployee(t *testing.T) {
	mockHandler := &MockDBHandler{
		CreateEmployeeFunc: func(employee domain.Employee) (domain.Employee, error) {
			return domain.Employee{ID: 1, Name: "John", Position: "Developer", Salary: 50000}, nil
		},
	}
	repo := NewEmployeeRepo(mockHandler)
	employee := domain.Employee{Name: "John", Position: "Developer", Salary: 50000}
	createdEmployee, err := repo.CreateEmployee(employee)
	if err != nil {
		t.Errorf("CreateEmployee returned an unexpected error: %v", err)
	}
	if createdEmployee.ID != 1 {
		t.Errorf("Expected employee ID to be 1, got %d", createdEmployee.ID)
	}
}

func TestGetEmployeeByID(t *testing.T) {
	mockHandler := &MockDBHandler{
		GetEmployeeByIDFunc: func(employee domain.Employee) (domain.Employee, error) {
			return domain.Employee{ID: 1, Name: "John", Position: "Developer", Salary: 50000}, nil
		},
	}
	repo := NewEmployeeRepo(mockHandler)
	employee := domain.Employee{ID: 1}
	retrievedEmployee, err := repo.GetEmployeeByID(employee)
	if err != nil {
		t.Errorf("GetEmployeeByID returned an unexpected error: %v", err)
	}
	if retrievedEmployee.Name != "John" {
		t.Errorf("Expected employee name to be John, got %s", retrievedEmployee.Name)
	}
}

func TestUpdateEmployee(t *testing.T) {
	mockHandler := &MockDBHandler{
		UpdateEmployeeFunc: func(employee domain.Employee) error {
			return nil
		},
	}
	repo := NewEmployeeRepo(mockHandler)
	employee := domain.Employee{ID: 1, Name: "John", Position: "Developer", Salary: 50000}
	err := repo.UpdateEmployee(employee)
	if err != nil {
		t.Errorf("UpdateEmployee returned an unexpected error: %v", err)
	}
}

func TestDeleteEmployee(t *testing.T) {
	mockHandler := &MockDBHandler{
		DeleteEmployeeFunc: func(employee domain.Employee) error {
			return nil
		},
	}
	repo := NewEmployeeRepo(mockHandler)
	employee := domain.Employee{ID: 1}
	err := repo.DeleteEmployee(employee)
	if err != nil {
		t.Errorf("DeleteEmployee returned an unexpected error: %v", err)
	}
}
