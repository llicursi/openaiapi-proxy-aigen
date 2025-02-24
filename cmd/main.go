package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "inspect-proxy/docs" // This line is important! Import the docs
	"inspect-proxy/internal/config"
	"inspect-proxy/internal/handlers"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Inspect Proxy API
// @version         1.0
// @description     A proxy service for inspecting and forwarding API requests.
// @host           localhost:8080
// @BasePath       /v1

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Create handler
	handler := handlers.NewHandler(cfg)
	handlers.SetupRoutes(handler)

	// Add Swagger documentation endpoint
	http.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("swagger/doc.json"), // Remove the leading slash
	))

	// Start server using configured port
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Environment: %s", os.Getenv("ENV"))
	log.Printf("Server starting on %s", addr)
	log.Printf("API Documentation available at http://localhost:%d/swagger/index.html", cfg.Server.Port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
