package requests

import "github.com/google/uuid"

type ChangeStatusTenderReq struct {
	Status   string
	TenderId uuid.UUID
}

type CreateTenderReq struct {
	Name            string
	Description     string
	ServiceType     string
	Status          string
	OrganizationId  uuid.UUID
	CreatorUsername string
	UserId          uuid.UUID
}

type GetTendersByUserIdReq struct {
	Username string
}

type FetchTendersForEmployeeReq struct {
	UserId uuid.UUID
}

type EditTenderReq struct {
	Name        string
	Description string
	TenderId    uuid.UUID
}
