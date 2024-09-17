package fetch

import (
	"avito_task/internal/model"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type repoBid interface {
	FetchBidsByUserId(ctx context.Context, uuid uuid.UUID) ([]model.Bid, error)
}

type repoEmployee interface {
	FetchEmployeeByUsername(ctx context.Context, username string) (model.Employee, error)
}

type UseCase struct {
	bid      repoBid
	employee repoEmployee
}

func NewBidUseCase(bid repoBid, employee repoEmployee) *UseCase {
	return &UseCase{
		bid:      bid,
		employee: employee,
	}
}

func (u *UseCase) FetchBids(ctx context.Context, username string) ([]model.Bid, error) {
	const op = "useCase.bid.FetchBids"
	employee, err := u.FetchEmployeeUseCase(ctx, username)
	if err != nil {
		return []model.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	bids, err := u.bid.FetchBidsByUserId(ctx, employee.ID)
	if err != nil {
		return []model.Bid{}, fmt.Errorf("%s: %w", op, err)
	}
	for i := range bids {
		bids[i].CreatorUsername = username
	}
	return bids, nil
}

func (u *UseCase) FetchEmployeeUseCase(ctx context.Context, username string) (model.Employee, error) {
	const op = "useCase.bid.FetchEmployeeByUsernameUseCase"
	emp, err := u.employee.FetchEmployeeByUsername(ctx, username)
	if err != nil {
		return model.Employee{}, fmt.Errorf("%s: %w", op, err)
	}
	return emp, nil
}
