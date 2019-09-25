package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// TODO: switch to usage of in-memory storage
var items = make(map[string]string)

func main() {
	// TODO: populate items storage externally, not in code
	items["name"] = "Roman"
	items["age"] = "35"
	items["weight"] = "82.5kg"
	items["car"] = "Renault"

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

	items[key] = value
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
	if _, ok := items[key]; !ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(404)
		return
	}

	value := items[key]
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
	keys := extractMapKeys(items)

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
	if _, ok := items[key]; !ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(404)
		return
	}
	delete(items, key)
	writer.WriteHeader(http.StatusNoContent)
}

func extractMapKeys(mapToExtract map[string]string) []string {
	keys := make([]string, 0, len(mapToExtract))
	for k := range mapToExtract {
		keys = append(keys, k)
	}
	return keys
}
