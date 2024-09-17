package edit

import (
	"avito_task/internal/requests"
	"github.com/google/uuid"
)

type EditBidRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (e *EditBidRequest) toDTO(uuid uuid.UUID) requests.EditBidReq {
	return requests.EditBidReq{
		Name:        e.Name,
		Description: e.Description,
		BidId:       uuid,
	}
}
