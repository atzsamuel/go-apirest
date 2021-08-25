package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"test3/database"
	"test3/entity"

	"github.com/gorilla/mux"
)

// get all students
func GetAllStudent(w http.ResponseWriter, r *http.Request) {
	var students []entity.Student
	database.Connector.Find(&students)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(students)
}

// getperson by id
func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var student entity.Student
	database.Connector.First(&student, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

// create student
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var student entity.Student
	json.Unmarshal(requestBody, &student)

	database.Connector.Create(student)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

// udpate student by id
func UpdateStudentByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var student entity.Student
	json.Unmarshal(requestBody, &student)
	database.Connector.Save(student)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

//delet student by id
func DeleteStudentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var student entity.Student
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&student)
	w.WriteHeader(http.StatusNoContent)
}
