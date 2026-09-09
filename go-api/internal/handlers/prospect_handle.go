package handlers

import (
	"encoding/json"
	"net/http"

	"go-api/internal/services"
)

type ProspectHandler struct {
	Service *services.ProspectService
}

func (h *ProspectHandler) GetRowCount(w http.ResponseWriter, r *http.Request) {
	count, err := h.Service.GetRowCount()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(count)
}
