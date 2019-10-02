package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-sandbox/cache/datastore"
	"net/http"
)

// TODO: switch to usage of in-memory storage
var storage = datastore.New()

func main() {
	// TODO: populate items storage externally, not in code
	storage.Set("name", "Roman")
	storage.Set("age", "35")
	storage.Set("weight", "82.5kg")
	storage.Set("car", "Renault")

	router := mux.NewRouter()
	router.HandleFunc("/items/{key}", CreateString).Methods("POST")
	router.HandleFunc("/items/keys", ReadStringKeys).Methods("GET")
	router.HandleFunc("/items/{key}", ReadString).Methods("GET")
	router.HandleFunc("/items/{key}", DeleteString).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}

func CreateString(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["key"]

	var value string
	err := json.NewDecoder(request.Body).Decode(&value)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(500)
		return
	}

	storage.Set(key, value)
	resultJson, err := json.Marshal(value)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(500)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write(resultJson)
}

func ReadString(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["key"]
	value, ok := storage.Get(key)
	if !ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(404)
		return
	}

	resultJson, err := json.Marshal(value)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(500)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(resultJson)
}

func ReadStringKeys(writer http.ResponseWriter, request *http.Request) {
	keys := storage.GetKeys()
	resultJson, err := json.Marshal(keys)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(500)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(resultJson)
}

func DeleteString(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["key"]
	if ok := storage.Delete(key); !ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(404)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}
