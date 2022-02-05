package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/viniciuslunardi/api-go-rest.git/database"
	"github.com/viniciuslunardi/api-go-rest.git/routes/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPeople(w http.ResponseWriter, r *http.Request) {
	var people []models.Person
	database.DB.Find(&people)
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var person models.Person

	database.DB.First(&person, id)

	json.NewEncoder(w).Encode(person)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	json.NewDecoder(r.Body).Decode(&person)
	database.DB.Create(&person)
	json.NewEncoder(w).Encode(person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var person models.Person

	database.DB.Delete(&person, id)

	json.NewEncoder(w).Encode(person)
}

func EditPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var person models.Person
	database.DB.First(&person, id)
	json.NewDecoder(r.Body).Decode(&person)
	database.DB.Save(&person)

	json.NewEncoder(w).Encode(person)
}
