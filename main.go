package main

/*
Following this guide for RESTful API in go:
https://medium.com/@johnteckert/building-a-restful-api-with-go-part-1-9e234774b14d
*/

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// ========== Structs ==========

// List implements ...
type List struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var lists []List

// Functions

// GetLists ... indexes lists
func GetLists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lists)
}

// GetList ... show specific checklist
func GetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range lists {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// CreateList ... create new checklist
func CreateList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newList List
	json.NewDecoder(r.Body).Decode(&newList)
	newList.ID = strconv.Itoa(len(lists) + 1)
	lists = append(lists, newList)
	json.NewEncoder(w).Encode(newList)
}

// UpdateList ... update a checklist
func UpdateList(w http.ResponseWriter, r *http.Request) {
}

// DeleteList ... delete checklist
func DeleteList(w http.ResponseWriter, r *http.Request) {
}

func main() {

	// Sample data
	lists = append(lists,
		List{ID: "1", Name: "Katter"},
		List{ID: "2", Name: "Planter"})

	// Initialize router
	router := mux.NewRouter()

	// Endpoints
	router.HandleFunc("/lists", GetLists).Methods("GET")
	router.HandleFunc("/lists/{id}", GetList).Methods("GET")
	router.HandleFunc("/lists/{id}", CreateList).Methods("POST")
	router.HandleFunc("/lists/{id}", DeleteList).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}
