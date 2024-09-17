package create

import (
	"avito_task/internal/model"
	"github.com/google/uuid"
)

type CreateTenderRequest struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	ServiceType     string    `json:"serviceType"`
	Status          string    `json:"status"`
	OrganizationId  uuid.UUID `json:"organizationId"`
	CreatorUsername string    `json:"creatorUsername"`
}

func (c *CreateTenderRequest) toDTO() model.CreateTenderReq {

	return model.CreateTenderReq{
		Name:            c.Name,
		Description:     c.Description,
		ServiceType:     c.ServiceType,
		Status:          c.Status,
		OrganizationId:  c.OrganizationId,
		CreatorUsername: c.CreatorUsername,
	}
}
