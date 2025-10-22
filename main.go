package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BerkAkipek/go-two-endpoints/handlers"
	"github.com/BerkAkipek/go-two-endpoints/routes"
)

func main() {
	if err := handlers.InitDB(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	if err := handlers.SetupSchema(); err != nil {
		log.Fatalf("Failed to setup database schema: %v", err)
	}

	router := routes.SetupRouter()
	port := ":8080"

	fmt.Printf("🚀 Server running on http://localhost%s\n", port)

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
