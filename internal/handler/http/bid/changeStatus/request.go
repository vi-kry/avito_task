package changeStatus

import (
	"avito_task/internal/requests"
	"github.com/google/uuid"
)

type ChangeStatusRequest struct {
	Status string    `json:"status"`
	BidId  uuid.UUID `json:"bid_id"`
}

func (c *ChangeStatusRequest) toDTO() requests.ChangeStatusBidReq {
	return requests.ChangeStatusBidReq{
		Status: c.Status,
		BidId:  c.BidId,
	}
}
