package controllers

import (
	"encoding/json"
	"net/http"
	"techiebutler/assessment/domain"
	"techiebutler/assessment/usecases"
)

// EmployeeController handles HTTP requests related to employee operations.
type EmployeeController struct {
	employeeInteractor usecases.EmployeeInteractor // employeeInteractor is an instance of the EmployeeInteractor.
}

// ErrorResponse represents the structure of error responses.
type ErrorResponse struct {
	Message string // Message contains the error message.
}

const (
	CONTENT_TYPE     = "Content-Type"        // CONTENT_TYPE is the header key for content type.
	APPLICATION_JSON = "application/json"    // APPLICATION_JSON is the value for JSON content type.
)

// NewEmployeeController creates a new instance of EmployeeController.
func NewEmployeeController(employeeInteractor usecases.EmployeeInteractor) *EmployeeController {
	return &EmployeeController{employeeInteractor}
}

// Add handles the HTTP POST request to create a new employee.
func (controller *EmployeeController) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	var employee domain.Employee
	err := json.NewDecoder(req.Body).Decode(&employee)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}
	employee, err2 := controller.employeeInteractor.CreateEmployee(employee)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(employee)
}

// Get handles the HTTP GET request to retrieve an employee.
func (controller *EmployeeController) Get(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	var employee domain.Employee
	err := json.NewDecoder(req.Body).Decode(&employee)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}
	data, err2 := controller.employeeInteractor.GetEmployeeByID(employee)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(data)
}

// Update handles the HTTP PUT request to update an employee.
func (controller *EmployeeController) Update(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	var employee domain.Employee
	err := json.NewDecoder(req.Body).Decode(&employee)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}
	err2 := controller.employeeInteractor.UpdateEmployee(employee)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
}

// Delete handles the HTTP DELETE request to delete an employee.
func (controller *EmployeeController) Delete(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	var employee domain.Employee
	err := json.NewDecoder(req.Body).Decode(&employee)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}
	err2 := controller.employeeInteractor.DeleteEmployee(employee)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
}
