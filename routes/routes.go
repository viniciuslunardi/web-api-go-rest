package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/viniciuslunardi/api-go-rest.git/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)

	// custom routes
	r.HandleFunc("/api/people", controllers.AllPeople).Methods("Get")
	r.HandleFunc("/api/people/{id}", controllers.GetPerson).Methods("Get")

	// error handler
	log.Fatal(http.ListenAndServe(":8000", r))
}
