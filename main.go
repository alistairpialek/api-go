package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alistairpialek/api-go/v1/routes"
	"github.com/alistairpialek/api-go/v1/utils"

	"github.com/gorilla/mux"
)

func handleRequests() {
	// Strictslash will redirect URL routes with a trailing / to the non-slash route. E.g path/ -> path
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc(routes.CalculateEndpoint, routes.PostCalculate).Methods("POST")
	router.HandleFunc(routes.HealthEndpoint, routes.GetHealth).Methods("GET")
	router.HandleFunc(routes.MetadataEndpoint, routes.GetMetadata).Methods("GET")
	router.Use(utils.MetricsMiddleware)

	log.Printf("Git commit: %s", utils.GitCommit)
	log.Printf("Listening on port: %s", os.Getenv("LISTEN_PORT"))

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%s", os.Getenv("LISTEN_PORT")),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func main() {
	handleRequests()
}
