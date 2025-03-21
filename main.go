package main

import (
	"fmt"
	"log"
	"net/http"
)

// Configurable global variable for the healthcheck endpoint
var (
	HealthCheckPath = "/healthcheck"
	Port            = 8006
)

func healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("success"))
	if err != nil {
		fmt.Println("error writing response", err)
	}
}

func main() {
	// Register the healthcheck handler with the configurable path
	http.HandleFunc(HealthCheckPath, healthCheckHandler)

	// Start the server
	fmt.Printf("Server starting on port %d with healthcheck endpoint at %s\n", Port, HealthCheckPath)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", Port), nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
