package main

import (
	"go-rag/internal/routes"
	"log"
	"net/http"
)

func main() {

	r := routes.SetupRoutes()

	log.Println("Starting server on :3333")
	if err := http.ListenAndServe(":3333", r); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
