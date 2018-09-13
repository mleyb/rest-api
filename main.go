package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", getPeople).Methods("GET")
	router.HandleFunc("/people/{id}", getPerson).Methods("GET")
	router.HandleFunc("/people/{id}", createPerson).Methods("POST")
	router.HandleFunc("/people/{id}", deletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	log.Println("getPeople")
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	log.Println("getPerson")
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	log.Println("createPerson")
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("deletePerson")
}
