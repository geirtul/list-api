package main

/*
Following this guide for RESTful API in go:
https://medium.com/@johnteckert/building-a-restful-api-with-go-part-1-9e234774b14d
*/

// TODO: Implement necessary functions for Item struct
// TODO: How to add item at list creation?
// TODO: Add Items to sample data

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ============================================================================
// STRUCTS
// ============================================================================

type Item struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type List struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Items []Item `json:"items"`
}

// Collection of lists
var lists []List

// ============================================================================
// FUNCTIONS
// ============================================================================

// ========== Item-functions

// ========== List-functions

// getLists ... indexes lists
func getLists(w http.ResponseWriter, r *http.Request) {
	// Get headers
	w.Header().Set("Content-Type", "application/json")

	// return json encoded list of existing lists
	json.NewEncoder(w).Encode(lists)
}

// getList ... show specific list
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

// createList ... create new list
func createList(w http.ResponseWriter, r *http.Request) {
	// Get headers
	w.Header().Set("Content-Type", "application/json")

	// Init new list with decoded json data and append to lists.
	var newList List
	json.NewDecoder(r.Body).Decode(&newList)
	newList.ID = strconv.Itoa(len(lists) + 1)

	lists = append(lists, newList)

	// return json encoded new list
	json.NewEncoder(w).Encode(newList)
}

// updateList ... update a list
// TODO: Should actually use http PATCH instead, but this works
func updateList(w http.ResponseWriter, r *http.Request) {
	// Get headers and params
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Iterate over existing lists to find item with matching ID.
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

// deleteList ... delete a list
func deleteList(w http.ResponseWriter, r *http.Request) {
	// Get headers and params
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Iterate over existing lists to find matching ID.
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
		List{
			ID:   "1",
			Name: "Katter",
			Items: []Item{
				Item{Name: "Asmo", Count: 1},
				Item{Name: "Luci", Count: 1}},
		},
		List{
			ID:   "2",
			Name: "Planter",
			Items: []Item{
				Item{Name: "Grønne", Count: 2},
				Item{Name: "Gule", Count: 3}},
		},
		List{
			ID:   "3",
			Name: "Handleliste",
			Items: []Item{
				Item{Name: "Ost", Count: 1},
				Item{Name: "Melk", Count: 1}},
		})

	// Initialize router
	router := mux.NewRouter()

	// Endpoints
	router.HandleFunc("/lists", getLists).Methods("GET")
	router.HandleFunc("/lists/{id}", getList).Methods("GET")
	router.HandleFunc("/lists/{id}", createList).Methods("POST")
	router.HandleFunc("/lists/{id}", updateList).Methods("POST")
	router.HandleFunc("/lists/{id}", deleteList).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}
