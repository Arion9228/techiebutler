package domain

type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name,omitempty"`
	Position string  `json:"position,omitempty"`
	Salary   float64 `json:"salary,omitempty"`
}

type EmployeeRepository interface {
	CreateEmployee(Employee) error
	GetEmployeeByID(Employee) (Employee, error)
	UpdateEmployee(Employee) error
	DeleteEmployee(Employee) error
}
