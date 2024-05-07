package domain

type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

type EmployeeRepository interface {
	CreateEmployee(employee Employee) error
	GetEmployeeByID(id int) (Employee, error)
	UpdateEmployee(employee Employee) error
	DeleteEmployee(id int) error
}
