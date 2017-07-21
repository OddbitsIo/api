package core

type OrganizationDto struct {
	Id string
	Name string
	TaxId string
}

type IOrganizationRepository interface {
	Get(id string) (*OrganizationDto, error)
	Save(organization *OrganizationDto) error
}