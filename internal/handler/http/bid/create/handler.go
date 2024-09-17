package create

import (
	"avito_task/internal/model"
	"avito_task/internal/requests"
	"context"
	"encoding/json"
	"net/http"
)

type bidCreator interface {
	CreateBidUseCase(ctx context.Context, req requests.CreateBidReq) (model.Bid, error)
	FetchEmployeeUseCase(ctx context.Context, username string) (model.Employee, error)
}

type logger interface {
	Error(msg string, args ...any)
}

type Handler struct {
	log        logger
	bidCreator bidCreator
}

func NewHandler(bidCreator bidCreator, logger logger) *Handler {
	return &Handler{
		bidCreator: bidCreator,
		log:        logger,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	var request CreateBidRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.log.Error("failed to decode request:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	req := request.toDTO()
	resp, err := h.bidCreator.CreateBidUseCase(r.Context(), req)
	if err != nil {
		h.log.Error("failed to create bid:", err)
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
