package main

import (
	"golang/internal/config"
	"golang/internal/handler"
	"golang/internal/repository"
	"golang/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg := config.NewConfig()

	// Initialize repository
	citiesFilePath, err := repository.GetCitiesFilePath()
	if err != nil {
		log.Fatalf("Failed to get cities file path: %v", err)
	}
	citiesRepo := repository.NewCitiesRepository(citiesFilePath)

	// Initialize service
	distributorService := service.NewDistributorService(citiesRepo)

	// Initialize handler
	distributorHandler := handler.NewDistributorHandler(distributorService)

	// Create router and register routes
	router := mux.NewRouter()
	router.HandleFunc("/distributors/permissions", distributorHandler.GetPermissions).Methods(http.MethodPost)

	// Start server
	listenAddr := cfg.Port
	log.Printf("About to listen on port %s.", listenAddr)
	log.Fatal(http.ListenAndServe(":"+listenAddr, router))
}
