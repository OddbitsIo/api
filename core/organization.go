package core

type OrganizationDto struct {
	Id string
	Name string
	TaxId string
}

func (this * OrganizationDto) IsEmpty() bool {
	return this.Id == ""
}

type IOrganizationRepository interface {
	Get(id string) (*OrganizationDto, error)
	Save(organization *OrganizationDto) error
}