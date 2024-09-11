package tender

import (
	"avito_task/internal/model"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type TenderRepo struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewTenderRepo(db *pgxpool.Pool, log *slog.Logger) *TenderRepo {
	return &TenderRepo{
		db:  db,
		log: log,
	}
}

func (tr *TenderRepo) Create(ctx context.Context, ten model.Tender) (model.Tender, error) {
	const op = "repository.tender.Create"

	query := `INSERT INTO tender (name, description, service_type, status, organization_id, creator_username) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`

	row := tr.db.QueryRow(ctx, query, ten.Name, ten.Description, ten.ServiceType, ten.Status, ten.OrganizationId, ten.CreatorUsername)

	var res TenderRow

	if err := row.Scan(&res); err != nil {
		tr.log.Error(op, slog.String("error", err.Error()))
		return model.Tender{}, fmt.Errorf("%s: scan result: %w", op, err)
	}
	return toModel(res), nil
}
