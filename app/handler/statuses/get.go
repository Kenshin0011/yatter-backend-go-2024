package statuses

import (
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
)

// Handle request for `GET /v1/statuses/{id}`
func (h *handler) FindByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	dto, err := h.statusUsecase.FindByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.Status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
