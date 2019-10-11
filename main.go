package main

import (
	"kunder-workshop-bk-go/internal/rest"
	"log"
	"net/http"
)

// How to try it: go run main.go
func main() {

	log.Print("Starting the service...")
	router := rest.Router()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8000", router))
}
