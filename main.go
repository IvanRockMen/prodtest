package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	base_path := route.PathPrefix("/api").Subrouter()

	base_path.HandleFunc("/createUser", createUser).Methods("POST")
	base_path.HandleFunc("/getAllUsers", getAllUsers).Methods("GET")
	base_path.HandleFunc("/getUser", getUser).Methods("POST")
	base_path.HandleFunc("/update", updateUser).Methods("PUT")
	base_path.HandleFunc("/delete/{id}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":80", base_path))
}
