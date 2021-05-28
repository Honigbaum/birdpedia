package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/welcome", handler).Methods("GET")
	return router
}

func main() {
	router := newRouter()
	http.ListenAndServe(":8080", router)
}

func handler(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "Welcome to Birdpedia!")
}
