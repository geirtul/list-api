package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// ========== Structs ==========

// Checklist implements ...
type Checklist struct {
	ID    string `json:"id,omitempty"`
	Items string `json:"items,omitempty"`
}

// Functions

// GetChecklists ...
func GetChecklists(w http.ResponseWriter, r *http.Request) {
}

// GetChecklist ...
func GetChecklist(w http.ResponseWriter, r *http.Request) {
}

// CreateChecklist ...
func CreateChecklist(w http.ResponseWriter, r *http.Request) {
}

// DeleteChecklist ...
func DeleteChecklist(w http.ResponseWriter, r *http.Request) {
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/checklists", GetChecklists).Methods("GET")
	router.HandleFunc("/checklists/{id}", GetChecklist).Methods("GET")
	router.HandleFunc("/checklists/{id}", CreateChecklist).Methods("POST")
	router.HandleFunc("/checklists/{id}", DeleteChecklist).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
