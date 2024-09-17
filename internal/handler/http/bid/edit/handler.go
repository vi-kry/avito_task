package edit

import (
	"context"
	"encoding/json"
	"net/http"

	"avito_task/internal/model"
	"avito_task/internal/requests"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type bidEditor interface {
	EditBid(ctx context.Context, req requests.EditBidReq) (model.Bid, error)
}

type logger interface {
	Error(msg string, args ...any)
}

type Handler struct {
	log       logger
	bidEditor bidEditor
}

func NewHandler(bidEditor bidEditor, log logger) *Handler {
	return &Handler{
		log:       log,
		bidEditor: bidEditor,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	bidId := chi.URLParam(r, "bidId")

	if bidId == "" {
		http.Error(w, "bidId is required", http.StatusBadRequest)
		return
	}

	var request EditBidRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.log.Error("failed to decode request:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	bidUUID, err := uuid.Parse(bidId)
	if err != nil {
		h.log.Error("failed to parse id:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := request.toDTO(bidUUID)
	resp, err := h.bidEditor.EditBid(r.Context(), req)
	if err != nil {
		h.log.Error("failed to edit bid:", err)
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
