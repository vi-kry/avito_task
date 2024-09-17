package employee

import (
	"context"
	"fmt"

	"avito_task/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepo struct {
	db *pgxpool.Pool
}

func NewEmployeeRepo(db *pgxpool.Pool) *EmployeeRepo {
	return &EmployeeRepo{
		db: db,
	}
}

func (e *EmployeeRepo) FetchEmployeeByUsername(ctx context.Context, username string) (model.Employee, error) {
	const op = "repository.employee.FetchEmployeeByUsername"

	query := `SELECT * FROM employee WHERE username = $1`
	row := e.db.QueryRow(ctx, query, username)

	var res EmployeeRow

	if err := row.Scan(&res.ID, &res.Username, &res.FirstName, &res.LastName, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return model.Employee{}, fmt.Errorf("%s: scan result: %w", op, err)
	}
	return toModel(res), nil
}

func (e *EmployeeRepo) FetchEmployeeById(ctx context.Context, uuid uuid.UUID) (model.Employee, error) {
	const op = "repository.employee.FetchEmployeeById"

	query := `SELECT * FROM employee WHERE id = $1`
	row := e.db.QueryRow(ctx, query, uuid)

	var res EmployeeRow

	if err := row.Scan(&res.ID, &res.Username, &res.FirstName, &res.LastName, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return model.Employee{}, fmt.Errorf("%s: scan result: %w", op, err)
	}
	return toModel(res), nil
}
