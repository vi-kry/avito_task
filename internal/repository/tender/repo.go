package tender

import (
	"context"
	"fmt"

	"avito_task/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TenderRepo struct {
	db *pgxpool.Pool
}

func NewTenderRepo(db *pgxpool.Pool) *TenderRepo {
	return &TenderRepo{
		db: db,
	}
}

func (tr *TenderRepo) CreateTender(ctx context.Context, req *model.CreateTenderReq) (model.Tender, error) {
	const op = "repository.tender.Create"
	query := `INSERT INTO tender (name, description, service_type, status, organization_id, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, current_timestamp, current_timestamp) RETURNING *`

	row := tr.db.QueryRow(ctx, query, req.Name, req.Description, req.ServiceType, req.Status, req.OrganizationId, req.UserId)

	var res tenderRow

	if err := row.Scan(&res.ID, &res.Name, &res.Description, &res.ServiceType, &res.Status, &res.OrganizationId, &res.UserId, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return model.Tender{}, fmt.Errorf("%s: scan result: %w", op, err)
	}
	return toModel(res), nil
}

func (tr *TenderRepo) FetchAllTenders(ctx context.Context) ([]model.Tender, error) {
	const op = "repository.tender.FetchAllTenders"
	query := `SELECT * FROM tender WHERE status = $1`
	rows, err := tr.db.Query(ctx, query, model.StatusTenderPublished)
	if err != nil {
		return nil, fmt.Errorf("%s: execute statement: %w", op, err)
	}

	var res []tenderRow

	for rows.Next() {
		var r tenderRow
		if err = rows.Scan(&r.ID, &r.Name, &r.Description, &r.ServiceType, &r.Status, &r.OrganizationId, &r.UserId, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, fmt.Errorf("%s: scan result: %w", op, err)
		}
		res = append(res, r)
	}
	return toModels(res), nil
}

func (tr *TenderRepo) ChangeStatusTender(ctx context.Context, req *model.ChangeStatusTenderReq) error {
	const op = "repository.tender.ChangeStatusTender"
	query := `UPDATE tender SET status = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2`

	_, err := tr.db.Exec(ctx, query, req.Status, req.TenderId)
	if err != nil {
		return fmt.Errorf("%s: execute statement: %w", op, err)
	}
	return nil
}

func (tr *TenderRepo) FetchTendersByUserId(ctx context.Context, uuid uuid.UUID) ([]model.Tender, error) {
	const op = "repository.tender.FetchTendersByUserId"

	query := `SELECT * FROM tender WHERE user_id = $1`
	rows, err := tr.db.Query(ctx, query, uuid)
	if err != nil {
		return nil, fmt.Errorf("%s: execute statement: %w", op, err)
	}

	var res []tenderRow

	for rows.Next() {
		var r tenderRow
		if err = rows.Scan(&r.ID, &r.Name, &r.Description, &r.ServiceType, &r.Status, &r.OrganizationId, &r.UserId, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, fmt.Errorf("%s: scan result: %w", op, err)
		}
		res = append(res, r)
	}
	return toModels(res), nil
}

func (tr *TenderRepo) EditTender(ctx context.Context, req *model.EditTenderReq) (model.Tender, error) {
	const op = "repository.tender.EditTender"
	query := `UPDATE tender SET name = $1, description = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3 RETURNING *`

	row := tr.db.QueryRow(ctx, query, req.Name, req.Description, req.TenderId)

	var res tenderRow

	if err := row.Scan(&res.ID, &res.Name, &res.Description, &res.ServiceType, &res.Status, &res.OrganizationId, &res.UserId, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return model.Tender{}, fmt.Errorf("%s: scan result: %w", op, err)
	}
	return toModel(res), nil
}
