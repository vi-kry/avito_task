package tender

import (
	"time"

	"github.com/google/uuid"
)

//DTO

type tenderRow struct {
	ID             uuid.UUID `db:"id"`
	Name           string    `db:"name"`
	Description    string    `db:"description"`
	ServiceType    string    `db:"service_type"`
	Status         string    `db:"status"`
	OrganizationId uuid.UUID `db:"organization_id"`
	UserId         uuid.UUID `db:"user_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
