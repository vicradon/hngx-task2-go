package main

import (
	"fmt"
	"net/http"
	personcrud "webapis/personcrudapi/person"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api", personcrud.GetAllPersons).Methods("GET")
	router.HandleFunc("/api", personcrud.CreateNewPerson).Methods("POST")
	router.HandleFunc("/api/{id:[0-9]+}", personcrud.GetPersonDetails).Methods("GET")
	router.HandleFunc("/api/{id:[0-9]+}", personcrud.UpdatePersonDetails).Methods("PATCH")
	router.HandleFunc("/api/{id:[0-9]+}", personcrud.DeletePerson).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
