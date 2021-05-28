package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/welcome", handler).Methods("GET")

	http.ListenAndServe(":8080", router)
}

func handler(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "Welcome to Birdpedia!")
}
