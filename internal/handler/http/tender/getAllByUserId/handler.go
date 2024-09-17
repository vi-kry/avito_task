package getAllByUserId

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"avito_task/internal/model"
)

type tendersByUserIdGetter interface {
	GetTendersByUserId(context.Context, *model.GetTendersByUserIdReq) ([]model.Tender, error)
	GetEmployeeByUsername(ctx context.Context, username string) (model.FetchEmployeeByUsernameResp, error)
}

type Handler struct {
	log                   *slog.Logger
	tendersByUserIdGetter tendersByUserIdGetter
}

func NewHandler(
	tendersByUserIdGetter tendersByUserIdGetter,
	log *slog.Logger,
) *Handler {
	return &Handler{
		tendersByUserIdGetter: tendersByUserIdGetter,
		log:                   log,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {

	username := r.URL.Query().Get("username")
	resp, err := h.tendersByUserIdGetter.GetTendersByUserId(r.Context(), &model.GetTendersByUserIdReq{Username: username})
	if err != nil {
		h.log.Error("failed to get tenders:", err)
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
