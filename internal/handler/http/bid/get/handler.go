package get

import (
	"context"
	"encoding/json"
	"net/http"

	"avito_task/internal/model"
)

type bidGetter interface {
	FetchBids(ctx context.Context, username string) ([]model.Bid, error)
	FetchEmployeeUseCase(ctx context.Context, username string) (model.Employee, error)
}

type logger interface {
	Error(msg string, args ...any)
}

type Handler struct {
	log       logger
	bidGetter bidGetter
}

func NewHandler(bidGetter bidGetter, logger logger) *Handler {
	return &Handler{
		bidGetter: bidGetter,
		log:       logger,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	resp, err := h.bidGetter.FetchBids(r.Context(), username)
	if err != nil {
		h.log.Error("failed to get bids:", err)
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
