package tender

//DTO

type TenderRow struct {
	ID              int    `db:"id"`
	Name            string `db:"name"`
	Description     string `db:"description"`
	ServiceType     string `db:"service_type"`
	Status          string `db:"status"`
	OrganizationId  int    `db:"organization_id"`
	CreatorUsername string `db:"creator_username"`
}
