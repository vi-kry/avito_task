package employee

import (
	"github.com/google/uuid"
	"time"
)

// ----- DTO

type EmployeeRow struct {
	ID        uuid.UUID `db:"id"`
	Username  string    `db:"username"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
