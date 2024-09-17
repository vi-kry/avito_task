package tender

import (
	"avito_task/internal/model"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type repoTenderGetTendersByUserId interface {
	FetchTendersByUserId(context.Context, uuid.UUID) ([]model.Tender, error)
}

type repoEmployeeGetTendersByUserId interface {
	FetchEmployeeByUsername(ctx context.Context, username string) (model.Employee, error)
}

type GetTendersByUserIdUseCase struct {
	tender   repoTenderGetTendersByUserId
	employee repoEmployeeGetTendersByUserId
}

func NewGetTendersByUserIdUseCase(
	tender repoTenderGetTendersByUserId,
	employee repoEmployeeGetTendersByUserId,
) *GetTendersByUserIdUseCase {
	return &GetTendersByUserIdUseCase{
		tender:   tender,
		employee: employee,
	}
}

func (g *GetTendersByUserIdUseCase) GetTendersByUserId(ctx context.Context, req *model.GetTendersByUserIdReq) ([]model.Tender, error) {
	const op = "useCase.tender.GetTendersByUserId"
	employee, err := g.GetEmployeeByUsername(ctx, req.Username)
	if err != nil {
		return []model.Tender{}, fmt.Errorf("%s: %w", op, err)
	}

	tenders, err := g.tender.FetchTendersByUserId(ctx, employee.ID)
	if err != nil {
		return []model.Tender{}, fmt.Errorf("%s: %w", op, err)
	}
	return tenders, nil
}

// POVTOR
func (g *GetTendersByUserIdUseCase) GetEmployeeByUsername(ctx context.Context, username string) (model.FetchEmployeeByUsernameResp, error) {
	const op = "useCase.tender.FetchEmployeeByUsername"
	emp, err := g.employee.FetchEmployeeByUsername(ctx, username)
	if err != nil {
		return model.FetchEmployeeByUsernameResp{}, fmt.Errorf("%s: %w", op, err)
	}
	return model.FetchEmployeeByUsernameResp{
		ID:        emp.ID,
		Username:  emp.Username,
		FirstName: emp.FirstName,
		LastName:  emp.LastName,
		CreatedAt: emp.CreatedAt,
		UpdatedAt: emp.UpdatedAt,
	}, nil
}
