// main.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/random/monster/human", getHuman).Methods("GET")
	router.HandleFunc("/api/random/monster/beast", getBeast).Methods("GET")

	log.Println("Server starting on :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
