package main

import (
	"fmt"
	"net/http"

	"github.com/Lazer430/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReportCaller(true)        // include calling method in log output
	var r *chi.Mux = chi.NewRouter() // create a new router of type chi
	handlers.RegisterHandlers(r)     // register the handlers

	// start the server
	fmt.Println("Starting server on port 8080")
	err := http.ListenAndServe("localhost:8080", r)

	if err != nil {
		log.Error(err)
		fmt.Println("Error starting server")
	}
}
