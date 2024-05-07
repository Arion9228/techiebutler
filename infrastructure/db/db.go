package db

import (
	"database/sql"
	"log"
	"techiebutler/assessment/domain"
	"sync"

	_ "github.com/lib/pq"
)

const (
	STATUS_ACTIVE  = 1 // STATUS_ACTIVE represents the active status of an employee
	STATUS_DELETED = 0 // STATUS_DELETED represents the deleted status of an employee
)

// DBHandler represents a handler for database operations
type DBHandler struct {
	DB     *sql.DB    // DB is the database connection
	dbLock sync.Mutex // dbLock is a Mutex for protecting concurrent access to the database
}

// NewDBHandler creates a new DBHandler instance with the given connection string
func NewDBHandler(connectString string) (DBHandler, error) {
	dbHandler := DBHandler{}
	db, err := sql.Open("postgres", connectString)
	if err != nil {
		log.Fatal(err)
		return dbHandler, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return dbHandler, err
	}
	dbHandler.DB = db
	return dbHandler, nil
}

// CreateEmployee inserts a new employee into the database
func (dbHandler DBHandler) CreateEmployee(employee domain.Employee) (domain.Employee, error) {
	dbHandler.dbLock.Lock()
	defer dbHandler.dbLock.Unlock()

	err := dbHandler.DB.QueryRow("INSERT INTO employees (name, position, salary, status) VALUES ($1, $2, $3, $4) RETURNING id",
		employee.Name, employee.Position, employee.Salary, STATUS_ACTIVE).Scan(&employee.ID)
	if err != nil {
		return domain.Employee{}, err
	}
	return employee, nil
}

// UpdateEmployee updates an existing employee in the database
func (dbHandler DBHandler) UpdateEmployee(employee domain.Employee) error {
	dbHandler.dbLock.Lock()
	defer dbHandler.dbLock.Unlock()

	_, err := dbHandler.DB.Exec("UPDATE employees SET name = $1, position = $2, salary = $3 WHERE id = $4",
		employee.Name, employee.Position, employee.Salary, employee.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetEmployeeByID retrieves an employee from the database by ID
func (dbHandler DBHandler) GetEmployeeByID(employee domain.Employee) (domain.Employee, error) {
	dbHandler.dbLock.Lock()
	defer dbHandler.dbLock.Unlock()

	err := dbHandler.DB.QueryRow("SELECT name, position, salary FROM employees WHERE id = $1", employee.ID).
		Scan(&employee.Name, &employee.Position, &employee.Salary)
	if err != nil {
		return domain.Employee{}, err
	}

	return employee, nil
}

// DeleteEmployee marks an employee as deleted in the database
func (dbHandler DBHandler) DeleteEmployee(employee domain.Employee) error {
	dbHandler.dbLock.Lock()
	defer dbHandler.dbLock.Unlock()

	_, err := dbHandler.DB.Exec("UPDATE employees SET status = $1 WHERE id = $2",
		STATUS_DELETED, employee.ID)
	if err != nil {
		return err
	}
	return nil
}
