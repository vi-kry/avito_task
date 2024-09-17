package tender

import "avito_task/internal/model"

func toModel(row tenderRow) model.Tender {
	return model.Tender{
		ID:             row.ID,
		Name:           row.Name,
		Description:    row.Description,
		ServiceType:    row.ServiceType,
		Status:         row.Status,
		OrganizationId: row.OrganizationId,
		EmployeeId:     row.UserId,
	}
}

func toModels(rows []tenderRow) []model.Tender {
	tenders := make([]model.Tender, 0, len(rows))

	for _, row := range rows {
		tenders = append(tenders, toModel(row))
	}

	return tenders
}
