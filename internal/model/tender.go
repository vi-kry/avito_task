package model

import (
	"github.com/google/uuid"
)

const (
	StatusTenderClosed    = "CLOSED"
	StatusTenderPublished = "PUBLISHED"
	StatusTenderCreated   = "CREATED"
)

type Tender struct {
	ID             uuid.UUID
	Name           string
	Description    string
	ServiceType    string
	Status         string
	OrganizationId uuid.UUID
	EmployeeId     uuid.UUID
}

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

type CreateTenderResp struct {
	ID              uuid.UUID
	Name            string
	Description     string
	ServiceType     string
	Status          string
	OrganizationId  uuid.UUID
	CreatorUsername string
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
