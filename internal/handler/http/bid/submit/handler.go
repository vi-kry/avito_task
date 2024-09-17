package submit

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type bidSubmitter interface {
	SubmitBid(ctx context.Context, bidId uuid.UUID) error
}

type logger interface {
	Error(msg string, args ...any)
}

type Handler struct {
	log          logger
	bidSubmitter bidSubmitter
}

func NewHandler(bidSubmitter bidSubmitter, log logger) *Handler {
	return &Handler{
		bidSubmitter: bidSubmitter,
		log:          log,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	var request SubmitBidRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.log.Error("failed to decode request:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = h.bidSubmitter.SubmitBid(r.Context(), request.BidId)
	if err != nil {
		h.log.Error("failed to submit bid:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
