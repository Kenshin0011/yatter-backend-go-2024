package timelines

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Params struct {
	Limit int `json:"limit"`
}

// Handle request for `GET /v1/timelines/public`
func (h *handler) FindPublicTimeline(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params, err := parseParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dto, err := h.statusUsecase.FindPublicTimeline(ctx, params.Limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func parseParams(r *http.Request) (Params, error) {
	var p Params
	queryParams := r.URL.Query()
	limitStr := queryParams.Get("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit <= 0 {
			return p, fmt.Errorf("invalid limit: %s", limitStr)
		}
		p.Limit = limit
	} else {
		return p, fmt.Errorf("limit is required")
	}

	return p, nil
}
