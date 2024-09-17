package employee

import "avito_task/internal/model"

func toModel(row EmployeeRow) model.Employee {
	return model.Employee{
		ID:        row.ID,
		Username:  row.Username,
		FirstName: row.FirstName,
		LastName:  row.LastName,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}

}
