package model

type Tender struct {
	ID              int
	Name            string
	Description     string
	ServiceType     string
	Status          string
	OrganizationId  int
	CreatorUsername string
}
