package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Structs

type Checklist struct {
	ID    string `json:"id,omitempty"`
	Items string `json:"firstname,omitempty"`
}

// Functions

func GetChecklists(w http.ResponsWriter, r *http.Request)    {}
func GetChecklist(w http.ResponsWriter, r *http.Request)     {}
func CreateChecklists(w http.ResponsWriter, r *http.Request) {}
func DeleteChecklists(w http.ResponsWriter, r *http.Request) {}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/checklists", GetChecklists).Methods("GET")
	router.HandleFunc("/checklists/{id}", GetChecklist).Methods("GET")
	router.HandleFunc("/checklists/{id}", CreateChecklist).Methods("POST")
	router.HandleFunc("/checklists/{id}", DeleteChecklist).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
