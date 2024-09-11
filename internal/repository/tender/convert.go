package tender

import "avito_task/internal/model"

func toModel(row TenderRow) model.Tender {
	return model.Tender{
		ID:              row.ID,
		Name:            row.Name,
		Description:     row.Description,
		ServiceType:     row.ServiceType,
		Status:          row.Status,
		OrganizationId:  row.OrganizationId,
		CreatorUsername: row.CreatorUsername,
	}
}
