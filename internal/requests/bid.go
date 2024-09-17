package requests

import "github.com/google/uuid"

type CreateBidReq struct {
	Name            string
	Description     string
	Status          string
	TenderId        uuid.UUID
	OrganizationId  uuid.UUID
	UserId          uuid.UUID
	CreatorUsername string
}

type EditBidReq struct {
	Name        string
	Description string
	BidId       uuid.UUID
}

type ChangeStatusBidReq struct {
	Status string
	BidId  uuid.UUID
}
