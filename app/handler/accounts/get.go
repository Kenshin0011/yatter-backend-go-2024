package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)


// Handle request for `GET /v1/accounts/{username}`
func (h *handler) FindByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	if username == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	dto, err := h.accountUsecase.FindByUsername(ctx, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.Account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}