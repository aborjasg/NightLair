package handlers

import (
	"encoding/json"
	"net/http"

	"go-api/internal/models"
	"go-api/internal/services"
)

type ProspectHandler struct {
	Service *services.ProspectService
}

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
}

func (h *ProspectHandler) HelloAPI(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	response := models.Message{
		Status:  "success",
		Message: "Hello from your first Go API!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *ProspectHandler) GetRowCount(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	count, err := h.Service.GetRowCount()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(count)
}
