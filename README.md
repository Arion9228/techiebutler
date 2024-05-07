This project handles the CRUD operations for an Employee. 
I am using Clean Architecture design of Separtion of concerns in this developement, where I have divided the section into 4 major components, namely, domain, infrastructure, interface, usecases.
Domain stores the blueprint of the Object that we want to work upon (Employee details in this case)
Infrastructure layer contains the implementation of DB and Mux-router
Interface layer is subdivided in two sections, controller (handles API request) and repository (handling communication with data layer)
Usecases is used for handling business usecase logic.
