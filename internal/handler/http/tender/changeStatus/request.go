package changeStatus

import (
	"avito_task/internal/model"
	"github.com/google/uuid"
)

type ChangeStatusRequest struct {
	Status   string    `json:"status"`
	TenderId uuid.UUID `json:"tender_id"`
}

func (c *ChangeStatusRequest) toDTO() model.ChangeStatusTenderReq {
	return model.ChangeStatusTenderReq{
		Status:   c.Status,
		TenderId: c.TenderId,
	}
}
