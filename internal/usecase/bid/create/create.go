package create

import (
	"avito_task/internal/model"
	"avito_task/internal/requests"
	"context"
	"fmt"
)

type repoBid interface {
	CreateBid(ctx context.Context, req requests.CreateBidReq) (model.Bid, error)
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

func (u *UseCase) CreateBidUseCase(ctx context.Context, req requests.CreateBidReq) (model.Bid, error) {
	const op = "useCase.bid.CreateBidUseCase"
	employee, err := u.FetchEmployeeUseCase(ctx, req.CreatorUsername)
	if err != nil {
		return model.Bid{}, fmt.Errorf("%s: %w", op, err)
	}
	req.UserId = employee.ID
	req.Status = model.StatusBidCreated

	bid, err := u.bid.CreateBid(ctx, req)
	if err != nil {
		return model.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	bid.UserId = employee.ID
	bid.CreatorUsername = employee.Username
	return bid, nil
}

func (u *UseCase) FetchEmployeeUseCase(ctx context.Context, username string) (model.Employee, error) {
	const op = "useCase.bid.FetchEmployeeByUsernameUseCase"
	emp, err := u.employee.FetchEmployeeByUsername(ctx, username)
	if err != nil {
		return model.Employee{}, fmt.Errorf("%s: %w", op, err)
	}
	return emp, nil
}
