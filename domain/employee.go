package domain

// Employee represents the structure of an employee
type Employee struct {
	ID       int     `json:"id"`               // ID of the employee
	Name     string  `json:"name,omitempty"`  // Name of the employee
	Position string  `json:"position,omitempty"`  // Position of the employee
	Salary   float64 `json:"salary,omitempty"`    // Salary of the employee
}

// EmployeeRepository is an interface defining methods for CRUD operations on Employee entities
type EmployeeRepository interface {
	// CreateEmployee creates a new employee
	CreateEmployee(Employee) (Employee, error)

	// GetEmployeeByID retrieves an employee by ID
	GetEmployeeByID(Employee) (Employee, error)

	// UpdateEmployee updates an existing employee
	UpdateEmployee(Employee) error

	// DeleteEmployee deletes an employee
	DeleteEmployee(Employee) error
}
