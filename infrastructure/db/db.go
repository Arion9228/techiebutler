package db

import (
	"database/sql"
	"log"
	"techiebutler/assessment/domain"
)

const (
	STATUS_ACTIVE  = 1
	STATUS_DELETED = 0
)

type DBHandler struct {
	DB *sql.DB
}

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

func (dbHandler DBHandler) CreateEmployee(employee domain.Employee) (int, error) {
	err := dbHandler.DB.QueryRow("INSERT INTO employee (name, position, salary, status) VALUES ($1, $2, $3, $4) RETURNING id",
		employee.Name, employee.Position, employee.Salary, STATUS_ACTIVE).Scan(&employee.ID)
	if err != nil {
		return -1, err
	}
	return employee.ID, nil
}

func (dbHandler *DBHandler) UpdateEmployee(employee domain.Employee) error {
	_, err := dbHandler.DB.Exec("UPDATE employee SET name = $1, position = $2, salary = $3 WHERE id = $5",
		employee.Name, employee.Position, employee.Salary, employee.ID)
	if err != nil {
		return err
	}
	return nil
}

func (dbHandler *DBHandler) GetEmployeeByID(employee domain.Employee) (domain.Employee, error) {
	err := dbHandler.DB.QueryRow("SELECT id, name, position, salary FROM employee WHERE id = $1", employee.ID).
		Scan(&employee.Name, &employee.Position, &employee.Salary)
	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (dbHandler *DBHandler) DeleteEmployee(employee domain.Employee) error {
	_, err := dbHandler.DB.Exec("UPDATE employee SET salary = $1 WHERE id = $2",
		STATUS_DELETED, employee.ID)
	if err != nil {
		return err
	}
	return nil
}
