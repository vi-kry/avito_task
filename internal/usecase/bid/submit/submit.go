package submit

import (
	"context"
	"fmt"

	"avito_task/internal/model"
	"github.com/google/uuid"
)

type repoBid interface {
	FetchBidByBidId(ctx context.Context, bidId uuid.UUID) (model.Bid, error)
}

type repoTender interface {
	ChangeStatusTender(ctx context.Context, req *model.ChangeStatusTenderReq) error
}

type UseCase struct {
	bid    repoBid
	tender repoTender
}

func NewSubmitUseCase(bid repoBid, tender repoTender) *UseCase {
	return &UseCase{
		bid:    bid,
		tender: tender,
	}
}

func (u *UseCase) SubmitBid(ctx context.Context, bidId uuid.UUID) error {
	const op = "useCase.bid.submitBid"
	fmt.Println(bidId, "USECASE")
	bid, err := u.bid.FetchBidByBidId(ctx, bidId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	req := &model.ChangeStatusTenderReq{
		Status:   model.StatusTenderClosed,
		TenderId: bid.TenderId,
	}

	err = u.tender.ChangeStatusTender(ctx, req)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
