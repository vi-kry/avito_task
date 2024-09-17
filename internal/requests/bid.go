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
