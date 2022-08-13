package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jose1up/go-gorm-resapi/db"

	"github.com/jose1up/go-gorm-resapi/models"

	"github.com/jose1up/go-gorm-resapi/routes"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})

	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)
	//users routes
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{user_id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PotsUserHandler).Methods("POST")
	r.HandleFunc("/users/{user_id}", routes.DeleteUserHandler).Methods("DELETE")
	//Task routers
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{task_id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{task_id}", routes.DeleteTasksHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)

}
