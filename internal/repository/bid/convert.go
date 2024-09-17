package bid

import "avito_task/internal/model"

func toModel(row bidRow) model.Bid {
	return model.Bid{
		ID:             row.ID,
		Name:           row.Name,
		Description:    row.Description,
		Status:         row.Status,
		TenderId:       row.TenderId,
		OrganizationId: row.OrganizationId,
		UserId:         row.UserId,
	}
}

func toModels(rows []bidRow) []model.Bid {
	bids := make([]model.Bid, 0, len(rows))

	for _, row := range rows {
		bids = append(bids, toModel(row))
	}
	return bids
}
