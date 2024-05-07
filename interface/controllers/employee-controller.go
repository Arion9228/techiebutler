package controllers

import (
	"encoding/json"
	"net/http"
	"techiebutler/assessment/domain"
	"techiebutler/assessment/usecases"
)

type EmployeeController struct {
	employeeInteractor usecases.EmployeeInteractor
}

type ErrorResponse struct {
	Message string
}

const (
	CONTENT_TYPE     = "Content-Type"
	APPLICATION_JSON = "application/json"
)

func NewEmployeeController(employeeInteractor usecases.EmployeeInteractor) *EmployeeController {
	return &EmployeeController{employeeInteractor}
}

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
