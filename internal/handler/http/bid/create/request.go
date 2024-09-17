package create

import (
	"avito_task/internal/requests"
	"github.com/google/uuid"
)

type CreateBidRequest struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Status          string    `json:"status"`
	TenderId        uuid.UUID `json:"tenderId"`
	OrganisationId  uuid.UUID `json:"organizationId"`
	CreatorUsername string    `json:"creatorUsername"`
}

func (c *CreateBidRequest) toDTO() requests.CreateBidReq {
	return requests.CreateBidReq{
		Name:            c.Name,
		Description:     c.Description,
		Status:          c.Status,
		TenderId:        c.TenderId,
		OrganizationId:  c.OrganisationId,
		CreatorUsername: c.CreatorUsername,
	}
}
