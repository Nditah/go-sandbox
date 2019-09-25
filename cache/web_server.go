package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

var items = make(map[string]string)

func main() {
	items["name"] = "Roman"
	items["age"] = "35"
	items["weight"] = "82.5kg"
	items["car"] = "Renault"

	router := mux.NewRouter()
	router.HandleFunc("/items/{key}", CreateString).Methods("POST")
	router.HandleFunc("/items/{key}", ReadString).Methods("GET")
	router.HandleFunc("/items/{key}", DeleteString).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}

func CreateString(writer http.ResponseWriter, request *http.Request) {
	// TODO: add implementation
}

func ReadString(writer http.ResponseWriter, request *http.Request) {
	// TODO: add implementation
}

func DeleteString(writer http.ResponseWriter, request *http.Request) {
	// TODO: add implementation
}
