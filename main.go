package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// A person
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"surname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// A person's address
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func main() {
	setupPeople()

	router := setupRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}

func setupPeople() {
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Y", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday", Address: &Address{City: "City Z", State: "State Z"}})
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/people", getPeople).Methods("GET")
	router.HandleFunc("/people/{id}", getPerson).Methods("GET")
	router.HandleFunc("/people/{id}", createPerson).Methods("POST")
	router.HandleFunc("/people/{id}", deletePerson).Methods("DELETE")
	return router
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	log.Println("getPeople")
	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		log.Println(err)
	}
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	log.Println("getPerson")
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404"))
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	log.Println("createPerson")
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}
