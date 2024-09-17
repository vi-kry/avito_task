package model

import "github.com/google/uuid"

const (
	StatusBidCreated = "CREATED"
)

type Bid struct {
	ID              uuid.UUID
	Name            string
	Description     string
	Status          string
	TenderId        uuid.UUID
	OrganizationId  uuid.UUID
	UserId          uuid.UUID
	CreatorUsername string
}
