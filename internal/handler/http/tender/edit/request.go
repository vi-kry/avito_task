package edit

import (
	"avito_task/internal/model"
	"github.com/google/uuid"
)

type EditTenderRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (e *EditTenderRequest) toDTO(uuid uuid.UUID) model.EditTenderReq {
	return model.EditTenderReq{
		Name:        e.Name,
		Description: e.Description,
		TenderId:    uuid,
	}
}
