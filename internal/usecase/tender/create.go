package tender

import (
	"avito_task/internal/model"
	"context"
	"fmt"
)

type repoTender interface {
	CreateTender(context.Context, *model.CreateTenderReq) (model.Tender, error)
}

type repoEmployee interface {
	FetchEmployeeByUsername(ctx context.Context, username string) (model.Employee, error)
}

type UseCase struct {
	tender   repoTender
	employee repoEmployee
}

func NewTenderUseCase(tender repoTender, employee repoEmployee) *UseCase {
	return &UseCase{
		tender:   tender,
		employee: employee,
	}
}

func (u *UseCase) CreateTenderUseCase(ctx context.Context, req *model.CreateTenderReq) (model.CreateTenderResp, error) {
	const op = "useCase.tender.CreateTenderUseCase"
	employee, err := u.FetchEmployeeByUsernameUseCase(ctx, req.CreatorUsername)
	if err != nil {
		return model.CreateTenderResp{}, fmt.Errorf("%s: %w", op, err)
	}

	req.UserId = employee.ID
	req.Status = model.StatusTenderCreated

	tender, err := u.tender.CreateTender(ctx, req)
	if err != nil {
		return model.CreateTenderResp{}, fmt.Errorf("%s: %w", op, err)
	}

	return model.CreateTenderResp{
		ID:              tender.ID,
		Name:            tender.Name,
		Description:     tender.Description,
		ServiceType:     tender.ServiceType,
		Status:          tender.Status,
		OrganizationId:  tender.OrganizationId,
		CreatorUsername: employee.Username,
	}, nil
}

func (u *UseCase) FetchEmployeeByUsernameUseCase(ctx context.Context, username string) (model.FetchEmployeeByUsernameResp, error) {
	const op = "useCase.tender.FetchEmployeeByUsernameUseCase"
	emp, err := u.employee.FetchEmployeeByUsername(ctx, username)
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
