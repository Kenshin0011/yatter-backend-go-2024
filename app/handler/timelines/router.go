package timelines

import (
	"net/http"
	"yatter-backend-go/app/usecase"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	statusUsecase usecase.Status
}

// Create Handler for `/v1/timelines/`
func NewRouter(su usecase.Status) http.Handler {
	r := chi.NewRouter()

	h := &handler{
		statusUsecase: su,
	}

	r.Get("/public", h.FindPublicTimeline)

	return r
}
