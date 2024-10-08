package bid

import (
	"context"
	"fmt"

	"avito_task/internal/model"
	"avito_task/internal/requests"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BidRepo struct {
	db *pgxpool.Pool
}

func NewBidRepo(db *pgxpool.Pool) *BidRepo {
	return &BidRepo{
		db: db,
	}
}

func (br *BidRepo) CreateBid(ctx context.Context, req requests.CreateBidReq) (model.Bid, error) {
	const op = "repository.bid.CreateBid"
	query := `INSERT INTO bid (name, description, status, tender_id, organization_id, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, current_timestamp, current_timestamp) RETURNING *`

	row := br.db.QueryRow(ctx, query, req.Name, req.Description, req.Status, req.TenderId, req.OrganizationId, req.UserId)

	var res bidRow

	if err := row.Scan(&res.ID, &res.Name, &res.Description, &res.Status, &res.TenderId, &res.OrganizationId, &res.UserId, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return model.Bid{}, fmt.Errorf("%s: scan result: %w", op, err)
	}

	return toModel(res), nil
}

func (br *BidRepo) FetchBidsByUserId(ctx context.Context, uuid uuid.UUID) ([]model.Bid, error) {
	const op = "repository.bid.FetchBidsByUserId"

	query := `SELECT * FROM bid WHERE user_id = $1`
	rows, err := br.db.Query(ctx, query, uuid)
	if err != nil {
		return nil, fmt.Errorf("%s: execute statement: %w", op, err)
	}

	var res []bidRow

	for rows.Next() {
		var r bidRow

		if err = rows.Scan(&r.ID, &r.Name, &r.Description, &r.Status, &r.TenderId, &r.OrganizationId, &r.UserId, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, fmt.Errorf("%s: scan result: %w", op, err)
		}
		res = append(res, r)
	}

	return toModels(res), nil
}

func (br *BidRepo) FetchBidsByTenderId(ctx context.Context, tenderId uuid.UUID) ([]model.Bid, error) {
	const op = "repository.bid.FetchBidsByTenderId"

	query := `SELECT * FROM bid WHERE tender_id = $1`
	rows, err := br.db.Query(ctx, query, tenderId)
	if err != nil {
		return nil, fmt.Errorf("%s: execute statement: %w", op, err)
	}

	var res []bidRow

	for rows.Next() {
		var r bidRow

		if err = rows.Scan(&r.ID, &r.Name, &r.Description, &r.Status, &r.TenderId, &r.OrganizationId, &r.UserId, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, fmt.Errorf("%s: scan result: %w", op, err)
		}
		res = append(res, r)
	}

	return toModels(res), nil
}

func (br *BidRepo) EditBid(ctx context.Context, req requests.EditBidReq) (model.Bid, error) {
	const op = "repository.bid.EditBid"
	query := `UPDATE bid SET name = $1, description = $2, updated_at = current_timestamp WHERE id = $3 RETURNING *`

	row := br.db.QueryRow(ctx, query, req.Name, req.Description, req.BidId)

	var res bidRow

	if err := row.Scan(&res.ID, &res.Name, &res.Description, &res.Status, &res.TenderId, &res.OrganizationId, &res.UserId, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return model.Bid{}, fmt.Errorf("%s: scan result: %w", op, err)
	}
	return toModel(res), nil
}

func (br *BidRepo) ChangeStatusBid(ctx context.Context, req requests.ChangeStatusBidReq) error {
	const op = "repository.bid.ChangeStatusBid"
	fmt.Println(req)
	query := `UPDATE bid SET status = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2`

	_, err := br.db.Exec(ctx, query, req.Status, req.BidId)
	if err != nil {
		return fmt.Errorf("%s: execute statement: %w", op, err)
	}
	return nil
}

func (br *BidRepo) FetchBidByBidId(ctx context.Context, bidId uuid.UUID) (model.Bid, error) {
	const op = "repository.bid.FetchBidByBidId"

	fmt.Println(bidId)
	query := `SELECT * FROM bid WHERE id = $1`

	row := br.db.QueryRow(ctx, query, bidId)

	var res bidRow

	if err := row.Scan(&res.ID, &res.Name, &res.Description, &res.Status, &res.TenderId, &res.OrganizationId, &res.UserId, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return model.Bid{}, fmt.Errorf("%s: scan result: %w", op, err)
	}
	fmt.Println(res)

	return toModel(res), nil
}
