package getAll

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"avito_task/internal/model"
)

type tendersGetter interface {
	FetchAllTenders(context.Context) ([]model.Tender, error)
}

type Handler struct {
	log           *slog.Logger
	tendersGetter tendersGetter
}

func NewHandler(tendersGetter tendersGetter, log *slog.Logger) *Handler {
	return &Handler{
		tendersGetter: tendersGetter,
		log:           log,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	resp, err := h.tendersGetter.FetchAllTenders(r.Context())
	if err != nil {
		h.log.Error("failed to fetch all tenders:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		h.log.Error("failed to marshal response:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
	w.WriteHeader(http.StatusOK)
}
