package main

import (
	/*"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"*/
	//"test3/database"

	"log"
	"net/http"
	"test3/controllers"
	"test3/database"
	"test3/entity"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	initDB()
	log.Println("Startint the HTTP server on port 8080")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreateStudent).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllStudent).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetStudentByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdateStudentByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeleteStudentByID).Methods("DELETE")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "ey",
			Password:   "code64CRUD",
			DB:         "students",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Student{})
}
