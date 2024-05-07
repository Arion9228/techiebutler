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

var (
	httpRouter router.MuxRouter = *router.NewMuxRouter()
	dbHandler  db.DBHandler
)

func getEmployeeController() controllers.EmployeeController {
	employeeRepo := repository.NewEmployeeRepo(dbHandler)
	employeeInteractor := usecases.NewEmployeeInteractor(employeeRepo)
	employeeController := controllers.NewEmployeeController(employeeInteractor)
	return *employeeController
}

func main() {

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "App is up and running..")
	})
	var err error
	dbHandler, err = db.NewDBHandler("postgres://postgres:postgres@123@localhost:5432/techiebutler?sslmode=disable")
	if err != nil {
		log.Println("Unable to connect to the DataBase")
		return
	}
	httpRouter.SERVE(":8000")
	employeeController := getEmployeeController()
	httpRouter.POST("/employee/add", employeeController.Add)
	httpRouter.GET("/employee/get", employeeController.Get)
	httpRouter.PUT("/employee/update", employeeController.Update)
	httpRouter.DELETE("/employee/delete", employeeController.Delete)
}
