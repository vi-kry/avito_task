package create

import (
	"avito_task/internal/model"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type tenderCreator interface {
	CreateTenderUseCase(context.Context, *model.CreateTenderReq) (model.CreateTenderResp, error)
	FetchEmployeeByUsernameUseCase(ctx context.Context, username string) (model.FetchEmployeeByUsernameResp, error)
}

type Handler struct {
	log           *slog.Logger
	tenderCreator tenderCreator
}

func NewHandler(tenderCreator tenderCreator, log *slog.Logger) *Handler {
	return &Handler{
		tenderCreator: tenderCreator,
		log:           log,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	var request CreateTenderRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.log.Error("failed to decode request:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// todo: добавить err = request.validate()

	req := request.toDTO()
	resp, err := h.tenderCreator.CreateTenderUseCase(r.Context(), &req)
	if err != nil {
		h.log.Error("failed to create tender:", err)
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
