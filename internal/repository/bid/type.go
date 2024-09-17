package bid

import (
	"github.com/google/uuid"
	"time"
)

// DTO

type bidRow struct {
	ID              uuid.UUID `db:"id"`
	Name            string    `db:"name"`
	Description     string    `db:"description"`
	Status          string    `db:"status"`
	TenderId        uuid.UUID `db:"tender_id"`
	OrganizationId  uuid.UUID `db:"organization_id"`
	UserId          uuid.UUID `db:"user_id"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
	CreatorUsername string
}
