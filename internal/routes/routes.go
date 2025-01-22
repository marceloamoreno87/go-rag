package routes

import (
	"fmt"
	"net/http"

	"go-rag/internal/handlers"
	"go-rag/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
)

func SetupRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	cfg := weaviate.Config{
		Host:    "weaviate:8080",
		Scheme:  "http",
		Headers: nil,
	}

	client, err := weaviate.NewClient(cfg)
	if err != nil {
		fmt.Println(err)
	}

	weaviateService := services.NewWeaviateService(client)
	h := handlers.NewHandler(weaviateService)

	r.Post("/generate", h.GenerateHandler)
	r.Post("/add-data", h.AddDataHandler)

	return r
}
