package tender

import (
	"avito_task/internal/model"
	"context"
	"log/slog"
)

type Repository interface {
	Create(context.Context, model.Tender) (model.Tender, error)
}

type TenderUseCase struct {
	repo Repository
	log  *slog.Logger
}

func NewTenderUseCase(repo Repository, log *slog.Logger) *TenderUseCase {
	return &TenderUseCase{
		repo: repo,
		log:  log,
	}
}

func (t *TenderUseCase) CreateTender(ctx context.Context, tender *model.Tender) (model.Tender, error) {
	const op = "usecase.tender.Create"
	//
}
