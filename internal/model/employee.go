package model

import (
	"github.com/google/uuid"
	"time"
)

type Employee struct {
	ID        uuid.UUID
	Username  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FetchEmployeeByUsernameReq struct {
	Username string
}

type FetchEmployeeByUsernameResp struct {
	ID        uuid.UUID
	Username  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
