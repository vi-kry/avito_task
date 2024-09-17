package changeStatus

import (
	"context"
	"encoding/json"
	"net/http"

	"avito_task/internal/model"
)

type tenderStatusChanger interface {
	ChangeStatusTender(context.Context, *model.ChangeStatusTenderReq) error
}

type logger interface {
	Error(msg string, args ...any)
}

type Handler struct {
	log                 logger
	tenderStatusChanger tenderStatusChanger
}

func NewHandler(tenderStatusChanger tenderStatusChanger, log logger) *Handler {
	return &Handler{
		tenderStatusChanger: tenderStatusChanger,
		log:                 log,
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

	err = h.tenderStatusChanger.ChangeStatusTender(r.Context(), &req)
	if err != nil {
		h.log.Error("failed to change status:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
