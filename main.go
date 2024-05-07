package main

import (
	"fmt"
	"log"
	"net/http"
	"techiebutler/assessment/infrastructure/db"
	"techiebutler/assessment/infrastructure/router"
	"techiebutler/assessment/interface/controllers"
	"techiebutler/assessment/interface/repository"
	"techiebutler/assessment/usecases"
)

// httpRouter is the HTTP router used for handling incoming requests.
var (
	httpRouter router.MuxRouter = *router.NewMuxRouter()
	dbHandler  db.DBHandler     // dbHandler is the handler for database operations.
)

// getEmployeeController initializes and returns the employee controller.
func getEmployeeController() controllers.EmployeeController {
	employeeRepo := repository.NewEmployeeRepo(dbHandler)               // Initialize employee repository.
	employeeInteractor := usecases.NewEmployeeInteractor(employeeRepo) // Initialize employee interactor.
	employeeController := controllers.NewEmployeeController(employeeInteractor) // Initialize employee controller.
	return *employeeController
}

func main() {
	// Define a route for checking if the application is up and running.
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "App is up and running..")
	})

	var err error
	// Initialize the database handler.
	dbHandler, err = db.NewDBHandler("postgres://postgres:postgres@123@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Println("Unable to connect to the DataBase")
		return
	}

	// Initialize the employee controller.
	employeeController := getEmployeeController()

	// Define routes for employee operations.
	httpRouter.POST("/employee/add", employeeController.Add)
	httpRouter.GET("/employee/get", employeeController.Get)
	httpRouter.PUT("/employee/update", employeeController.Update)
	httpRouter.DELETE("/employee/delete", employeeController.Delete)

	// Start serving HTTP requests on port 8000.
	httpRouter.SERVE(":8000")
}
