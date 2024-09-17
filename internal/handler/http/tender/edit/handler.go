package edit

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"avito_task/internal/model"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type tenderEditor interface {
	EditTender(context.Context, *model.EditTenderReq) (model.Tender, error)
}

type Handler struct {
	log          *slog.Logger
	tenderEditor tenderEditor
}

func NewHandler(tender tenderEditor, log *slog.Logger) *Handler {
	return &Handler{
		log:          log,
		tenderEditor: tender,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {

	tenderId := chi.URLParam(r, "tenderId")

	if tenderId == "" {
		http.Error(w, "tenderId is required", http.StatusBadRequest)
		return
	}

	var request EditTenderRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.log.Error("failed to decode request:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	tenderUUID, err := uuid.Parse(tenderId)
	if err != nil {
		h.log.Error("failed to parse id:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := request.toDTO(tenderUUID)
	resp, err := h.tenderEditor.EditTender(r.Context(), &req)
	if err != nil {
		h.log.Error("failed to edit tender:", err)
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
