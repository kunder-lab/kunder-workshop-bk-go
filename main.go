package main

import (
	"koath/internal/rest"
	"log"
	"net/http"

	"cloud.google.com/go/profiler"
)

// How to try it: go run main.go
func main() {
	log.Print("Starting Profiler")
	if err := profiler.Start(profiler.Config{
		Service:        "koath_service",
		ServiceVersion: "1.0.0",
	}); err != nil {
		// TODO: Handle error.
		log.Print("The profiler has an error")
		log.Print(err)
	}
	log.Print("Starting the service...")
	router := rest.Router()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8000", router))
}
