package get

import "net/http"

type UseCase interface {
	Do()
}

type Handler struct {
	uc UseCase
}

func NewHandler(uc UseCase) *Handler {
	return &Handler{uc}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	h.uc.Do()
}
