package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/viniciuslunardi/api-go-rest.git/controllers"
	"github.com/viniciuslunardi/api-go-rest.git/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)

	// middleware
	r.Use(middleware.ContentTypeMiddeware)

	// custom routes
	r.HandleFunc("/api/people", controllers.AllPeople).Methods("Get")
	r.HandleFunc("/api/people/{id}", controllers.GetPerson).Methods("Get")
	r.HandleFunc("/api/people", controllers.CreatePerson).Methods("Post")
	r.HandleFunc("/api/people/{id}", controllers.DeletePerson).Methods("Delete")
	r.HandleFunc("/api/people/{id}", controllers.EditPerson).Methods("Patch")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
