package changeStatus

import (
	"avito_task/internal/requests"
	"context"
	"encoding/json"
	"net/http"
)

type bidStatusChanger interface {
	ChangeStatusBid(ctx context.Context, req requests.ChangeStatusBidReq) error
}

type logger interface {
	Error(msg string, args ...any)
}

type Handler struct {
	log              logger
	bidStatusChanger bidStatusChanger
}

func NewHandler(bidStatusChanger bidStatusChanger, log logger) *Handler {
	return &Handler{
		bidStatusChanger: bidStatusChanger,
		log:              log,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	var request ChangeStatusRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.log.Error("failed to decode request:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	req := request.toDTO()

	err = h.bidStatusChanger.ChangeStatusBid(r.Context(), req)
	if err != nil {
		h.log.Error("failed to change status:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
