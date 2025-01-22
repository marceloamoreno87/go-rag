package handlers

import (
	"encoding/json"
	"net/http"

	"go-rag/internal/services"
	"go-rag/pkg/models"
)

type Handler struct {
	weaviateService services.WeaviateService
}

func NewHandler(
	weaviateService *services.WeaviateService,
) *Handler {
	return &Handler{
		weaviateService: *weaviateService,
	}
}

func (h *Handler) GenerateHandler(w http.ResponseWriter, r *http.Request) {
	var req models.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vectorResponse, err := h.weaviateService.Search(r.Context(), req.Query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var res models.Response
	vectorResponseBytes, err := json.Marshal(vectorResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(vectorResponseBytes, &res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res.Data.Get.QuestionTest[0].Additional.Generate)
}

func (h *Handler) AddDataHandler(w http.ResponseWriter, r *http.Request) {
	if err := h.weaviateService.AddDefaultData(r.Context()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
