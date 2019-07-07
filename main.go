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

// getLists ... indexes lists
func getLists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lists)
}

// getList ... show specific checklist
func getList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, item := range lists {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// createList ... create new checklist
func createList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newList List
	json.NewDecoder(r.Body).Decode(&newList)
	newList.ID = strconv.Itoa(len(lists) + 1)

	lists = append(lists, newList)

	json.NewEncoder(w).Encode(newList)
}

// updateList ... update a list
func updateList(w http.ResponseWriter, r *http.Request) {
	// Get headers and params
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Iterate over lists to find correct list to update
	for i, item := range lists {
		if item.ID == params["id"] {

			// remove element i from lists slice
			// slicing works like python
			// ... tells Go to treat each element as separate argument
			lists = append(lists[:i], lists[i+1:]...)

			// Create new, updated list with same id as the old version
			// of the list
			var newList List
			json.NewDecoder(r.Body).Decode(&newList)
			newList.ID = params["id"]
			lists = append(lists, newList)
			json.NewEncoder(w).Encode(newList)
			return
		}
	}
}

// deleteList ... delete list
func deleteList(w http.ResponseWriter, r *http.Request) {
	// Get headers and params
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range lists {
		if item.ID == params["id"] {
			lists = append(lists[:i], lists[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(lists)
}

func main() {

	// Sample data
	lists = append(lists,
		List{ID: "1", Name: "Katter"},
		List{ID: "2", Name: "Planter"})

	// Initialize router
	router := mux.NewRouter()

	// Endpoints
	router.HandleFunc("/lists", getLists).Methods("GET")
	router.HandleFunc("/lists/{id}", getList).Methods("GET")
	router.HandleFunc("/lists/{id}", createList).Methods("POST")
	router.HandleFunc("/lists/{id}", deleteList).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}
