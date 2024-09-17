package getByTenderId

import (
	"context"
	"encoding/json"
	"net/http"

	"avito_task/internal/model"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type bidGetter interface {
	FetchBidsByTenderId(ctx context.Context, tenderId uuid.UUID) ([]model.Bid, error)
}

type logger interface {
	Error(msg string, args ...any)
}

type Handler struct {
	bidGetter bidGetter
	log       logger
}

func NewHandler(bidGetter bidGetter, logger logger) *Handler {
	return &Handler{
		bidGetter: bidGetter,
		log:       logger,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	tenderId := chi.URLParam(r, "tenderId")
	if tenderId == "" {
		http.Error(w, "tenderId is required", http.StatusBadRequest)
		return
	}

	tenderUUID, err := uuid.Parse(tenderId)
	if err != nil {
		h.log.Error("failed to parse id:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.bidGetter.FetchBidsByTenderId(r.Context(), tenderUUID)
	if err != nil {
		h.log.Error("failed to fetch bids:", err)
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
